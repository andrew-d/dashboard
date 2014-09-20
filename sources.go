package main

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/jmoiron/modl"
)

type Source struct {
	Id       int64                  `json:"id"`
	Name     string                 `json:"name"`
	Type     string                 `json:"type"`
	Settings map[string]interface{} `json:"settings"`
}

// "Internal" only, for writing to DB
type dbSource struct {
	Id       int64
	Name     string
	Type     string
	Settings []byte
}

func fromSource(s *Source) (*dbSource, error) {
	var err error

	ret := &dbSource{
		Id:   s.Id,
		Name: s.Name,
		Type: s.Type,
	}

	ret.Settings, err = json.Marshal(s.Settings)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *dbSource) toSource() (*Source, error) {
	ret := &Source{
		Id:       s.Id,
		Name:     s.Name,
		Type:     s.Type,
		Settings: make(map[string]interface{}),
	}

	err := json.Unmarshal(s.Settings, &ret.Settings)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *dbSource) updateFrom(n *Source) (err error) {
	s.Name = n.Name
	s.Type = n.Type
	s.Settings, err = json.Marshal(n.Settings)
	return
}

type SourceApi struct {
	dbm *modl.DbMap
}

func NewSourceApi(dbm *modl.DbMap) *SourceApi {
	ret := &SourceApi{
		dbm: dbm,
	}

	tmap := dbm.AddTable(dbSource{}, "sources").SetKeys(true, "id")
	tmap.ColMap("name").SetUnique(true)

	dbm.AddTable(SourceData{}, "source_data").SetKeys(true, "sourceid")
	return ret
}

func (s *SourceApi) Get(id int64) (*Source, error) {
	var d dbSource

	err := s.dbm.Get(&d, id)
	if err != nil {
		return nil, err
	}

	return d.toSource()
}

func (s *SourceApi) Delete(id int64) error {
	d := dbSource{Id: id}
	i, err := s.dbm.Delete(&d)
	if i == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (s *SourceApi) Add(n *Source) error {
	d, err := fromSource(n)
	if err != nil {
		return err
	}

	err = s.dbm.Insert(d)
	if err != nil {
		return err
	}

	n.Id = d.Id

	// Also insert a blank SourceData entry that can be updated later.
	sd := SourceData{
		SourceId: n.Id,
		Updated:  0,
		Value:    []byte("{}"),
	}
	err = s.dbm.Insert(&sd)
	if err != nil {
		return err
	}

	return nil
}

func (s *SourceApi) Update(u *Source) error {
	var d dbSource

	err := s.dbm.Get(&d, u.Id)
	if err != nil {
		return err
	}

	err = d.updateFrom(u)
	if err != nil {
		return err
	}

	_, err = s.dbm.Update(s)
	return err
}

func (s *SourceApi) List() ([]*Source, error) {
	var sources []*dbSource
	err := s.dbm.Select(&sources, `SELECT * FROM sources`)
	if err != nil {
		return nil, err
	}

	ret := make([]*Source, len(sources))
	for i, s := range sources {
		ret[i], err = s.toSource()
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

type SourceData struct {
	SourceId int64
	Updated  int64
	Value    []byte
}

func (s *SourceApi) AddData(i *Source, data interface{}) error {
	// Validate the data
	err := ValidateTypeData(i.Type, data)
	if err != nil {
		return err
	}

	d := SourceData{
		SourceId: i.Id,
		Updated:  time.Now().UTC().Unix(),
	}

	d.Value, err = json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = s.dbm.Update(&d)
	return err
}

func (s *SourceApi) GetData(i *Source) (data interface{}, updated int64, err error) {
	var d SourceData

	err = s.dbm.Get(&d, i.Id)
	if err != nil {
		return
	}

	updated = d.Updated
	err = json.Unmarshal(d.Value, &data)
	return
}
