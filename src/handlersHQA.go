package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleParamsTech(w http.ResponseWriter, r *http.Request) {

	result, err := GetBodyResponse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}
	log.Println("--->-->")
	setInfoTech(result["technologies"].(string), result["architecture"].(string), result["requirement"].(string), result["id"].(string))

	w.WriteHeader(http.StatusOK)
	byteData, _ := json.Marshal(msgResponseOk)
	w.Write(byteData)
}
