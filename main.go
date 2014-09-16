package main

import (
	"time"
	"net/http"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/goji/glogrus"
	"github.com/stretchr/graceful"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"github.com/googollee/go-socket.io"
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
	m := web.New()
	m.Use(middleware.RequestID)
	m.Use(glogrus.NewGlogrus(log, "dashboard"))
	m.Use(middleware.Recoverer)
	m.Use(middleware.AutomaticOptions)

	// Static assets
	m.Get("/", ServeAsset("index.html", "text/html"))
	for _, asset := range AssetDescriptors() {
		m.Get("/"+asset.Path, ServeAsset(asset.Path, asset.Mime))
	}

	// API sub-handler.
	api := web.New()
	m.Handle("/api/*", api)

	api.Use(JSONContentType)

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
