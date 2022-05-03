package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBodyResponse(req *http.Request) map[string]interface{} {

	body, error := ioutil.ReadAll(req.Body)
	if error != nil {
		fmt.Println(error)
	}
	req.Body.Close()
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	return result
}

func GetBodyResponseRequest(client *http.Client, r *http.Request) (string, error) {

	respuesta, err := client.Do(r)
	if err != nil {
		return "Error haciendo petición: ", err
	}
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		return "Error leyendo respuesta: ", err
	}
	return string(cuerpoRespuesta), nil
}
