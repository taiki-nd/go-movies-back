package main

import (
	"encoding/json"
	"net/http"
)

//データをjsonにするための関数
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data
	js, err := json.Marshal(wrapper)
	//js, err := json.MarshalIndent(wrapper, "", "\t")で見やすくなる
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
