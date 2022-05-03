package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(write_ http.ResponseWriter, req *http.Request) {
	write_.WriteHeader(http.StatusOK)

}

func HandleAsanaCode(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	var asana Asana
	asana.GetProperties()
	m, err := GetCode(asana)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
		return
	} else {
		fmt.Fprintf(w, m)
	}
}

func HandleAsanaOauth(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	result := GetBodyResponse(req)
	code := result["code"]
	code_verifier := result["code_verifier"]
	client := &http.Client{}
	r := OauthAsana(code.(string), code_verifier.(string))
	res, err := GetBodyResponseRequest(client, r)
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
	r := ProjectsAsana(token)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := GetGeneral(res)
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
	r := SectionsAsana(token, project)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := GetGeneral(res)
		if len(elements) > 0 {
			json.NewEncoder(w).Encode(elements)
		} else {
			fmt.Fprintf(w, "[]")

		}
	}
}

func HandleAsanaSectionsTasks(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	section := req.Header.Get("sectionId")
	client := &http.Client{}
	r := TaskSectionAsana(token, section)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := GetGeneral(res)
		var task []Task
		for i := len(elements) - 1; i >= 0; i-- {
			r := TaskAsana(token, elements[i].Gid)
			res, err := GetBodyResponseRequest(client, r)
			if err != nil {
				fmt.Fprintf(w, "%v\"%v\"}", res, err)
			} else {
				task_und := GetTask(res)
				r := StoriesAsana(token, elements[i].Gid)
				res2, err := GetBodyResponseRequest(client, r)
				if err != nil {
					fmt.Fprintf(w, "%v\"%v\"}", res, err)
				} else {
					elements_ := GetStoriesFilter(res2, "comment")
					task_und.Story = elements_
					r := DependenciesAsana(token, elements[i].Gid)
					res3, err := GetBodyResponseRequest(client, r)
					if err != nil {
						fmt.Fprintf(w, "%v\"%v\"}", res, err)
					} else {
						elements_dep := GetGeneral(res3)
						task_und.Dependecies = elements_dep
						task = append(task, task_und)
					}
				}
			}
		}
		json.NewEncoder(w).Encode(task)
	}
}

func HandleAsanaTasks(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	section := req.Header.Get("sectionId")
	client := &http.Client{}
	r := TaskSectionAsana(token, section)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := GetGeneral(res)
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
	r := TaskAsana(token, task)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		json.NewEncoder(w).Encode(GetTask(res))
	}
}

func HandleAsanaTasksIdStories(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	task := req.Header.Get("id")
	client := &http.Client{}
	r := StoriesAsana(token, task)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := GetStoriesFilter(res, "comment")
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
	r := DependenciesAsana(token, task)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	} else {
		elements := GetGeneral(res)
		if len(elements) > 0 {
			json.NewEncoder(w).Encode(elements)
		} else {
			fmt.Fprintf(w, "[]")
		}
	}
}

func CarPostRequest(write_ http.ResponseWriter, req *http.Request) {
	var car Car
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&car)
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", err)
		return
	}
	err = insertDB(&car)
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgDatabase)
		return
	}
	responseCarBody(&car, write_)
}

func CarGetRequest(write_ http.ResponseWriter, req *http.Request) {
	var car Car
	err := getDB(&car, req.Header.Get("id"))
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgNotFound)
		return
	}
	write_.WriteHeader(http.StatusOK)
	responseCarBody(&car, write_)
}

func CarDeleteRequest(write_ http.ResponseWriter, req *http.Request) {
	var car Car
	err := getDB(&car, req.Header.Get("id"))
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgNotFound)
		return
	}
	err = deleteDB(&car)
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgDatabase)
		return
	}
	write_.WriteHeader(http.StatusNonAuthoritativeInfo)
	responseCarBody(&car, write_)
}

func responseCarBody(car *Car, write_ http.ResponseWriter) {
	response, err := car.ToJson()
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgMalFormat)
		return
	}
	write_.Write(response)
}
