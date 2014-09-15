package main

import (
	"time"

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

func main() {
	m := web.New()
	m.Use(middleware.RequestID)
	m.Use(glogrus.NewGlogrus(log, "dashboard"))
	m.Use(middleware.Recoverer)
	m.Use(middleware.AutomaticOptions)

	// TODO: serve static assets

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
