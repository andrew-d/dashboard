package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"
)

func SourcesList(c web.C, w http.ResponseWriter, r *http.Request) {
	api := c.Env["api.sources"].(*SourceApi)
	sources, err := api.List()
	if err != nil {
		writeError(w, 500, err)
		return
	}

	json.NewEncoder(w).Encode(sources)
}

func SourcesGet(c web.C, w http.ResponseWriter, r *http.Request) {
	api := c.Env["api.sources"].(*SourceApi)
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		writeError(w, 400, err)
		return
	}

	db, err := api.Get(id)
	if err != nil {
		writeError(w, 500, err)
		return
	}

	json.NewEncoder(w).Encode(db)
}

func SourcesAdd(c web.C, w http.ResponseWriter, r *http.Request) {
	api := c.Env["api.sources"].(*SourceApi)
	n := Source{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		writeError(w, 400, err)
		return
	}

	// Validation
	if len(n.Name) == 0 {
		writeError(w, 400, fmt.Errorf("'name' cannot be empty"))
		return
	}
	if len(n.Type) == 0 {
		writeError(w, 400, fmt.Errorf("'type' cannot be empty"))
		return
	}
	// TODO: Validate type is correct

	err = api.Add(&n)
	if err != nil {
		writeError(w, 500, err)
		return
	}

	json.NewEncoder(w).Encode(n)
}

func writeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "error",
		"error":  err.Error(),
		"code":   code,
	})
}

func SetupApiRoutes(m *web.Mux) {
	m.Get("/api/sources", SourcesList)
	m.Get("/api/sources/:id", SourcesGet)
	m.Post("/api/sources", SourcesAdd)
}
