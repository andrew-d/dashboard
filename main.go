package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/goji/glogrus"
	"github.com/googollee/go-socket.io"
	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/graceful"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	// Database drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	log *logrus.Logger
)

func init() {
	log = logrus.New()
}

func ServeAsset(name, mime string) http.Handler {
	// Assert that the asset exists.
	_, err := Asset(name)
	if err != nil {
		panic(fmt.Sprintf("asset named '%s' does not exist", name))
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		asset, _ := Asset(name)
		w.Header().Set("Content-Type", mime)
		w.Write(asset)
	}

	return http.HandlerFunc(handler)
}

func main() {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Could not open db")
		return
	}

	// TODO: configurable
	dbm := modl.NewDbMap(db.DB, modl.SqliteDialect{})
	sourceApi := NewSourceApi(dbm)

	err = dbm.CreateTablesIfNotExists()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Could not create DB tables")
	}

	m := web.New()
	m.Use(middleware.RequestID)
	m.Use(glogrus.NewGlogrus(log, "dashboard"))
	m.Use(middleware.Recoverer)
	m.Use(middleware.AutomaticOptions)
	m.Use(func(c *web.C, h http.Handler) http.Handler {
		handler := func(w http.ResponseWriter, r *http.Request) {
			c.Env["dbm"] = dbm
			c.Env["api.sources"] = sourceApi
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(handler)
	})

	// Static assets
	m.Get("/", ServeAsset("index.html", "text/html"))
	for _, asset := range AssetDescriptors() {
		m.Get("/"+asset.Path, ServeAsset(asset.Path, asset.Mime))
	}

	// API sub-handler.
	api := web.New()
	api.Use(JSONContentType)
	m.Handle("/api/*", api)

	SetupApiRoutes(api)

	// Socket.IO handler.
	sio, err := socketio.NewServer(nil)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("error creating Socket.IO server")
	}
	m.Handle("/socket.io/*", sio)

	// Start it all up.
	log.WithFields(logrus.Fields{
		"address": ":8000",
	}).Infof("Server starting")
	graceful.Run(":8000", 10*time.Second, m)
}
