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
		if IsNotFound(err) {
			writeError(w, 404, fmt.Errorf("source with id %d not found", id))
			return
		} else {
			writeError(w, 500, err)
			return
		}
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

func SourcesDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	api := c.Env["api.sources"].(*SourceApi)
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		writeError(w, 400, err)
		return
	}

	err = api.Delete(id)
	if err != nil {
		if IsNotFound(err) {
			writeError(w, 404, fmt.Errorf("source with id %d not found", id))
			return
		} else {
			writeError(w, 500, err)
			return
		}
	}

	w.WriteHeader(204)
}

func SourcesGetData(c web.C, w http.ResponseWriter, r *http.Request) {
	api := c.Env["api.sources"].(*SourceApi)
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		writeError(w, 400, err)
		return
	}

	source, err := api.Get(id)
	if err != nil {
		if IsNotFound(err) {
			writeError(w, 404, fmt.Errorf("source with id %d not found", id))
			return
		} else {
			writeError(w, 500, err)
			return
		}
	}

	data, updated, err := api.GetData(source)
	if err != nil {
		if IsNotFound(err) {
			writeError(w, 404, fmt.Errorf("data for source with id %d not found", id))
			return
		} else {
			writeError(w, 500, err)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":    data,
		"updated": updated,
	})
}

func SourcesAddData(c web.C, w http.ResponseWriter, r *http.Request) {
	api := c.Env["api.sources"].(*SourceApi)
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		writeError(w, 400, err)
		return
	}

	var data interface{}

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		writeError(w, 400, err)
		return
	}

	source, err := api.Get(id)
	if err != nil {
		if IsNotFound(err) {
			writeError(w, 404, fmt.Errorf("source with id %d not found", id))
			return
		} else {
			writeError(w, 500, err)
			return
		}
	}

	err = api.AddData(source, data)
	if err != nil {
		writeError(w, 500, err)
		return
	}

	// All good!
	// TODO: what to return?
}

func TypesList(c web.C, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(validTypes)
}

func SetupApiRoutes(m *web.Mux) {
	m.Get("/api/sources", SourcesList)
	m.Post("/api/sources", SourcesAdd)
	m.Get("/api/sources/:id", SourcesGet)
	m.Delete("/api/sources/:id", SourcesDelete)
	m.Get("/api/sources/:id/data", SourcesGetData)
	m.Post("/api/sources/:id/data", SourcesAddData)

	m.Get("/api/types", TypesList)
}

func writeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "error",
		"error":  err.Error(),
		"code":   code,
	})
}
