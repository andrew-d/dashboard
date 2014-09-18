package main

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type Source struct {
	Id       int64
	Name     string
	Type     string
	Settings map[string]interface{}
}

type SourceApi struct {
	db *sqlx.DB
}

func NewSourceApi(db *sqlx.DB) *SourceApi {
	return &SourceApi{
		db: db,
	}
}

func (s *SourceApi) CreateTables() error {
	statements := []string{
		`CREATE TABLE sources (
			id INTEGER NOT NULL AUTO_INCREMENT,
			name TEXT NOT NULL,
			type TEXT NOT NULL,
			settings BLOB NOT NULL,

			PRIMARY KEY (id)
		)`,
	}

	tx := s.db.MustBegin()
	for _, stmt := range statements {
		_, err := tx.Exec(stmt)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (s *SourceApi) GetSourceById(id int64) (*Source, error) {
	ret := Source{}
	err := s.db.Get(&ret, s.db.Rebind(`SELECT id, name, type FROM sources WHERE id=?`), id)
	if err != nil {
		return nil, err
	}

	var settingsJson []byte
	err = s.db.Get(&ret, s.db.Rebind(`SELECT settings FROM sources WHERE id=?`), id)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(settingsJson, &ret.Settings)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SourceApi) GetSourceByName(name string) (*Source, error) {
	ret := Source{}
	err := s.db.Get(&ret, s.db.Rebind(`SELECT id, name, type FROM sources WHERE name=?`), name)
	if err != nil {
		return nil, err
	}

	var settingsJson []byte
	err = s.db.Get(&ret, s.db.Rebind(`SELECT settings FROM sources WHERE name=?`), name)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(settingsJson, &ret.Settings)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *SourceApi) AddSource(n *Source) error {
	settingsJson, err := json.Marshal(n.Settings)
	if err != nil {
		return err
	}

	_, err = s.db.NamedExec(""+
		`INSERT INTO sources (id, name, type, settings) `+
		`VALUES (:id, :name, :type, :settings)`,
		map[string]interface{}{
			"id":       n.Id,
			"name":     n.Name,
			"type":     n.Type,
			"settings": settingsJson,
		})

	return err
}

func (s *SourceApi) UpdateSource(u *Source) error {
	settingsJson, err := json.Marshal(u.Settings)
	if err != nil {
		return err
	}

	_, err = s.db.NamedExec(""+
		`UPDATE sources SET name=:name, type=:type, settings=:settings `+
		`WHERE id=:id`,
		map[string]interface{}{
			"id":       u.Id,
			"name":     u.Name,
			"type":     u.Type,
			"settings": settingsJson,
		})

	return err
}
