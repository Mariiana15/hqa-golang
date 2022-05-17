package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleParamsTech(w http.ResponseWriter, r *http.Request) {

	var m responseOk
	result, err := GetBodyResponse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}

	if result["technologies"] == nil || result["architecture"] == nil || result["requirement"] == nil || result["id"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", "Rquest no contein field 'technologies', 'architecture','requirement','id'")
		return
	}

	err = setInfoTech(result["technologies"].(string), result["architecture"].(string), result["requirement"].(string), result["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	m.Message = msgResponseOk1
	byteData, _ := json.Marshal(m)
	w.Write(byteData)
}
func HandleChangeStateSection(w http.ResponseWriter, r *http.Request) {

	var m responseOk
	result, err := GetBodyResponse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}
	if result["state"] == nil || result["id"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", "Rquest no contein field 'state', 'id'")
		return
	}
	err = setChangeStateSection(result["state"].(string), result["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	m.Message = msgResponseOk1
	byteData, _ := json.Marshal(m)
	w.Write(byteData)
}

func HandleChangeStateUserStory(w http.ResponseWriter, r *http.Request) {

	var m responseOk
	result, err := GetBodyResponse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}
	if result["state"] == nil || result["id"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", "Rquest no contein field 'state', 'id'")
		return
	}
	err = setChangeStateUserStory(result["state"].(string), result["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}
	var t Task
	t.Hid = result["id"].(string)
	errTaskR := createUserStoryResultHQA(&t)
	if errTaskR != nil {
		log.Println(errTaskR)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", errTaskR.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	m.Message = msgResponseOk1
	byteData, _ := json.Marshal(m)
	w.Write(byteData)
}

func HandleResultUserStory(w http.ResponseWriter, r *http.Request) {

	var t Task
	var res Result
	var m responseOk

	body, err := GetBodyResponse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}
	t.Hid = body["id"].(string)
	byteData, _ := json.Marshal(body)
	json.Unmarshal(byteData, &res)
	t.Result = res
	errDB := t.setUserStoryResult()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", errDB.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	m.Message = msgResponseOk1
	byteData, _ = json.Marshal(m)
	w.Write(byteData)
}

func HandleGetValidateUStory(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	tokenString := ExtractToken(r)
	acc, err2 := ExtractTokenMetadataWS(tokenString)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err2)
		return

	}
	var s []Section
	errDB := getSectionDB(acc.UserId, &s)
	if errDB != nil || len(s) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", errDB)
		return
	}
	fmt.Fprintf(w, "{\"message\": \"%v\"}", "It is synchronized")
}
