package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Mariiana15/apis"
	"github.com/Mariiana15/dbmanager"
	"github.com/Mariiana15/serverutils"
)

func HandleParamsTech(w http.ResponseWriter, r *http.Request) {

	var m serverutils.ResponseOk
	result, err := serverutils.GetBodyResponse(r)
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

	err = dbmanager.SetInfoTech(result["technologies"].(string), result["architecture"].(string), result["requirement"].(string), result["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	m.Message = serverutils.MsgResponseOk1
	byteData, _ := json.Marshal(m)
	w.Write(byteData)
}
func HandleChangeStateSection(w http.ResponseWriter, r *http.Request) {

	var m serverutils.ResponseOk
	result, err := serverutils.GetBodyResponse(r)
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
	err = dbmanager.SetChangeStateSection(result["state"].(string), result["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	m.Message = serverutils.MsgResponseOk1
	byteData, _ := json.Marshal(m)
	w.Write(byteData)
}

func HandleChangeStateUserStory(w http.ResponseWriter, r *http.Request) {

	var m serverutils.ResponseOk
	result, err := serverutils.GetBodyResponse(r)
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
	err = dbmanager.SetChangeStateUserStory(result["state"].(string), result["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}
	var t dbmanager.Task
	t.Hid = result["id"].(string)
	errTaskR := apis.CreateUserStoryResultHQA(&t)
	if errTaskR != nil {
		log.Println(errTaskR)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", errTaskR.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	m.Message = serverutils.MsgResponseOk1
	byteData, _ := json.Marshal(m)
	w.Write(byteData)
}

func HandleResultUserStory(w http.ResponseWriter, r *http.Request) {

	var t dbmanager.Task
	var res dbmanager.Result
	var m serverutils.ResponseOk

	body, err := serverutils.GetBodyResponse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
		return
	}
	t.Hid = body["id"].(string)
	byteData, _ := json.Marshal(body)
	json.Unmarshal(byteData, &res)
	t.Result = res
	errDB := t.SetUserStoryResult()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", errDB.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	m.Message = serverutils.MsgResponseOk1
	byteData, _ = json.Marshal(m)
	w.Write(byteData)
}

func HandleGetValidateUStory(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	tokenString := serverutils.ExtractToken(r)
	acc, err2 := serverutils.ExtractTokenMetadataWS(tokenString)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err2)
		return

	}
	var s []dbmanager.Section
	errDB := dbmanager.GetSectionDB(acc.UserId, &s)
	if errDB != nil || len(s) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", errDB)
		return
	}
	fmt.Fprintf(w, "{\"message\": \"%v\"}", "It is synchronized")
}
