package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONType map[string]interface{}

func (j JSONType) Value() (driver.Value, error) {
	v, err := json.Marshal(j)
	return v, err
}

func (j *JSONType) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]bype) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*j, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}
	return nil
}

type WikiPage struct {
	ID        string `json:id,omitempty`
	Title     string `json:title,omitempty`
	Redirect  string `json:redirect,omitempty`
	Text      string `json:text,omitempty`
	CreatedAt string `json:text,omitempty`
}
