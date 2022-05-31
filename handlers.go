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

func HandleRoot(write_ http.ResponseWriter, req *http.Request) {
	write_.WriteHeader(http.StatusOK)
}

func HandleAsanaCode(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	var asana apis.Asana
	asana.GetProperties()
	m, err := apis.GetCode(asana)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
		return
	} else {
		fmt.Fprintf(w, m)
	}
}

func HandleAsanaCodeDB(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	tokenString := serverutils.ExtractToken(req)
	acc, err2 := serverutils.ExtractTokenMetadataWS(tokenString)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err2)
		return

	}
	code, cv, errDB := dbmanager.GetUserCodeAsana(acc.UserId)
	if errDB != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", errDB)
		return
	}
	fmt.Fprintf(w, "{\"code\": \"%v\", \"code_verifier\":\"%v\"}", code, cv)

}

func HandleAsanaOauth(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	result, _ := serverutils.GetBodyResponse(req)
	code := result["code"].(string)
	code_verifier := result["code_verifier"].(string)

	tokenString := serverutils.ExtractToken(req)
	_, err2 := serverutils.ExtractTokenMetadataWS(tokenString)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err2)
		return
	}

	client := &http.Client{}
	r := apis.OauthAsana(code, code_verifier)
	res, err := serverutils.GetBodyResponseRequest(client, r)
	log.Println(res)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		var response map[string]interface{}
		json.Unmarshal([]byte(res), &response)
		fmt.Fprintf(w, "{\"token\":\"%v\"}", response["access_token"])
		fmt.Println(response["access_token"])
	}
}

func HandleAsanaProjects(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	client := &http.Client{}
	token := req.Header.Get("token")
	r := apis.ProjectsAsana(token)
	res, err := serverutils.GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := apis.GetGeneral(res)
		if len(elements) > 0 {
			json.NewEncoder(w).Encode(elements)
		} else {
			fmt.Fprintf(w, "[]")
		}
	}
}

func HandleAsanaSections(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	client := &http.Client{}
	token := req.Header.Get("token")
	project := req.Header.Get("projectId")
	r := apis.SectionsAsana(token, project)
	res, err := serverutils.GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := apis.GetGeneral(res)
		if len(elements) > 0 {
			json.NewEncoder(w).Encode(elements)
		} else {
			fmt.Fprintf(w, "[]")

		}
	}
}

func HandleAsanaSectionsId(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	client := &http.Client{}
	token := req.Header.Get("token")
	section := req.Header.Get("id")
	r := apis.SectionsAsanaId(token, section)
	res, err := serverutils.GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := apis.GetGeneral(res)
		if len(elements) > 0 {
			json.NewEncoder(w).Encode(elements)
		} else {
			fmt.Fprintf(w, "[]")

		}
	}
}

func HandleAsanaSectionsTasksAsync(w http.ResponseWriter, req *http.Request, elements []dbmanager.General, token string, section string) {

	client := &http.Client{}
	var tasks []dbmanager.Task
	for i := len(elements) - 1; i >= 0; i-- {
		var task dbmanager.Task
		r := make(chan *http.Request)
		r2 := make(chan *http.Request)
		r3 := make(chan *http.Request)

		go apis.GetTaskAsync("task", token, elements[i].Gid, r)
		go apis.GetTaskAsync("stories", token, elements[i].Gid, r2)
		go apis.GetTaskAsync("dependencies", token, elements[i].Gid, r3)

		rr := <-r
		res, err := serverutils.GetBodyResponseRequest(client, rr)
		if err != nil {
			fmt.Fprintf(w, "%v\"%v\"}", res, err)
		}
		task = apis.GetTask(res)

		rr2 := <-r2
		res2, err := serverutils.GetBodyResponseRequest(client, rr2)
		if err != nil {
			fmt.Fprintf(w, "%v\"%v\"}", res, err)
		}
		elements_ := apis.GetStoriesFilter(res2, "comment")
		task.Story = elements_

		rr3 := <-r3
		res3, err := serverutils.GetBodyResponseRequest(client, rr3)
		if err != nil {
			fmt.Fprintf(w, "%v\"%v\"}", res, err)
		}
		elements_dep := apis.GetGeneral(res3)
		task.Dependecies = elements_dep
		tasks = append(tasks, task)
	}
	fmt.Println(tasks)
	//json.NewEncoder(w).Encode(task)
}
func HandleAsanaSectionsTasks(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	section := req.Header.Get("sectionId")
	client := &http.Client{}
	r, t := apis.TaskSectionAsana(token, section)
	res, err := serverutils.GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	}
	elements := apis.GetGeneral(res)
	if len(elements) > 0 {
		fmt.Fprintf(w, "{\"tasks\":\"%v\",\"timeAsync\":\"%v\"}", len(elements), t)
		HandleAsanaSectionsTasksAsync(w, req, elements, token, section)
	} else {
		fmt.Fprintf(w, "[]")
	}
}

func HandleAsanaTasks(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	section := req.Header.Get("sectionId")
	client := &http.Client{}
	r, _ := apis.TaskSectionAsana(token, section)
	res, err := serverutils.GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := apis.GetGeneral(res)
		if len(elements) > 0 {
			json.NewEncoder(w).Encode(elements)
		} else {
			fmt.Fprintf(w, "[]")
		}
	}
}

func HandleAsanaTasksId(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	task := req.Header.Get("id")
	client := &http.Client{}
	r := apis.TaskAsana(token, task)
	res, err := serverutils.GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		json.NewEncoder(w).Encode(apis.GetTask(res))
	}
}

func HandleAsanaTasksIdStories(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	task := req.Header.Get("id")
	client := &http.Client{}
	r := apis.StoriesAsana(token, task)
	res, err := serverutils.GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := apis.GetStoriesFilter(res, "comment")
		if len(elements) > 0 {
			json.NewEncoder(w).Encode(elements)
		} else {
			fmt.Fprintf(w, "[]")
		}
	}
}

func HandleAsanaTasksIdDependencies(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	task := req.Header.Get("id")
	client := &http.Client{}
	r := apis.DependenciesAsana(token, task)
	res, err := serverutils.GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := apis.GetGeneral(res)
		if len(elements) > 0 {
			json.NewEncoder(w).Encode(elements)
		} else {
			fmt.Fprintf(w, "[]")
		}
	}
}

func CarPostRequest(write_ http.ResponseWriter, req *http.Request) {
	var car dbmanager.Car
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&car)
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", err)
		return
	}
	err = dbmanager.InsertDB(&car)
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", serverutils.MsgDatabase)
		return
	}
	responseCarBody(&car, write_)
}

func CarGetRequest(write_ http.ResponseWriter, req *http.Request) {
	var car dbmanager.Car
	err := dbmanager.GetDB(&car, req.Header.Get("id"))
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", serverutils.MsgNotFound)
		return
	}
	write_.WriteHeader(http.StatusOK)
	responseCarBody(&car, write_)
}

func CarDeleteRequest(write_ http.ResponseWriter, req *http.Request) {
	var car dbmanager.Car
	err := dbmanager.GetDB(&car, req.Header.Get("id"))
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", serverutils.MsgNotFound)
		return
	}
	err = dbmanager.DeleteDB(&car)
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", serverutils.MsgDatabase)
		return
	}
	write_.WriteHeader(http.StatusNonAuthoritativeInfo)
	responseCarBody(&car, write_)
}

func responseCarBody(car *dbmanager.Car, write_ http.ResponseWriter) {

	response, err := car.ToJson()
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", serverutils.MsgMalFormat)
		return
	}
	write_.Write(response)
}
