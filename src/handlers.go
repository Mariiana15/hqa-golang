package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HandleRoot(write_ http.ResponseWriter, req *http.Request) {
	write_.WriteHeader(http.StatusOK)

}

func HandleAsanaCode(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	var mainAsana MainAsana
	mainAsana.GetProperties()
	m, err := GetCode(mainAsana)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
		return
	}
	fmt.Fprintf(w, m)
}

func HandleAsanaOauth(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	body, error := ioutil.ReadAll(req.Body)
	if error != nil {
		fmt.Println(error)
	}
	req.Body.Close()
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	//return result
	//result := GetBodyResponse(req)
	code := result["code"]
	code_verifier := result["code_verifier"]
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", "1201830256646257")
	data.Set("client_secret", "3d19ec3d43e86e52add5d0439bafc054")
	data.Set("redirect_uri", "http://localhost:3000/sync/")
	data.Set("code", code.(string))
	data.Set("code_verifier", code_verifier.(string))

	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodPost, "https://app.asana.com/-/oauth_token", strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	respuesta, err := client.Do(r)
	if err != nil {
		// Maneja el error de acuerdo a tu situación
		fmt.Printf("Error haciendo petición: %v", err)
	}

	// No olvides cerrar el cuerpo al terminar
	defer respuesta.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Printf("Error leyendo respuesta: %v", err)
	}

	respuestaString := string(cuerpoRespuesta)
	var response map[string]interface{}
	json.Unmarshal([]byte(respuestaString), &response)

	fmt.Fprintf(w, "{\"token\":\"%v\"}", response["access_token"])
	fmt.Println(response["access_token"])

}

func HandleAsanaProjects(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/projects", nil) // URL-encoded payload
	r.Header.Add("Authorization", "Bearer "+token)
	respuesta, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error haciendo petición: %v", err)
	}
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Printf("Error leyendo respuesta: %v", err)
	}
	respuestaString := string(cuerpoRespuesta)

	elements := GetGeneral(respuestaString)
	if len(elements) > 0 {
		json.NewEncoder(w).Encode(elements)
	} else {
		fmt.Fprintf(w, "[]")

	}

}

func HandleAsanaSections(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	project := req.Header.Get("projectId")
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/projects/"+project+"/sections", nil) // URL-encoded payload
	r.Header.Add("Authorization", "Bearer "+token)
	respuesta, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error haciendo petición: %v", err)
	}
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Printf("Error leyendo respuesta: %v", err)
	}
	respuestaString := string(cuerpoRespuesta)

	elements := GetGeneral(respuestaString)
	if len(elements) > 0 {
		json.NewEncoder(w).Encode(elements)
	} else {
		fmt.Fprintf(w, "[]")

	}
}

func HandleAsanaSectionsTasks(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	section := req.Header.Get("sectionId")
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/tasks", nil) // URL-encoded payload
	r.Header.Add("Authorization", "Bearer "+token)

	values := r.URL.Query()
	values.Add("section", section)
	r.URL.RawQuery = values.Encode()

	respuesta, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error haciendo petición: %v", err)
	}
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Printf("Error leyendo respuesta: %v", err)
	}
	respuestaString := string(cuerpoRespuesta)
	elements := GetGeneral(respuestaString)

	var task []Task
	for i := len(elements) - 1; i >= 0; i-- {
		r, _ = http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/tasks/"+elements[i].Gid, nil) // URL-encoded payload
		r.Header.Add("Authorization", "Bearer "+token)
		respuesta, err := client.Do(r)
		if err != nil {
			fmt.Printf("Error haciendo petición: %v", err)
		}
		defer respuesta.Body.Close()
		cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
		if err != nil {
			fmt.Printf("Error leyendo respuesta: %v", err)
		}
		respuestaString := string(cuerpoRespuesta)
		task_und := GetTask(respuestaString)

		r, _ = http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/tasks/"+elements[i].Gid+"/stories", nil) // URL-encoded payload
		r.Header.Add("Authorization", "Bearer "+token)
		respuesta, err = client.Do(r)
		if err != nil {
			fmt.Printf("Error haciendo petición: %v", err)
		}
		defer respuesta.Body.Close()
		cuerpoRespuesta, err = ioutil.ReadAll(respuesta.Body)
		if err != nil {
			fmt.Printf("Error leyendo respuesta: %v", err)
		}
		respuestaString = string(cuerpoRespuesta)
		elements_ := GetStoriesFilter(respuestaString, "comment")
		task_und.Story = elements_

		r, _ = http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/tasks/"+elements[i].Gid+"/dependencies", nil) // URL-encoded payload
		r.Header.Add("Authorization", "Bearer "+token)

		respuesta, err = client.Do(r)
		if err != nil {
			fmt.Printf("Error haciendo petición: %v", err)
		}
		defer respuesta.Body.Close()
		cuerpoRespuesta, err = ioutil.ReadAll(respuesta.Body)
		if err != nil {
			fmt.Printf("Error leyendo respuesta: %v", err)
		}
		respuestaString = string(cuerpoRespuesta)
		elements_dep := GetGeneral(respuestaString)
		task_und.Dependecies = elements_dep
		task = append(task, task_und)

	}
	json.NewEncoder(w).Encode(task)

}

func HandleAsanaTasks(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	section := req.Header.Get("sectionId")
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/tasks", nil) // URL-encoded payload
	r.Header.Add("Authorization", "Bearer "+token)

	values := r.URL.Query()
	values.Add("section", section)
	r.URL.RawQuery = values.Encode()

	respuesta, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error haciendo petición: %v", err)
	}
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Printf("Error leyendo respuesta: %v", err)
	}
	respuestaString := string(cuerpoRespuesta)
	elements := GetGeneral(respuestaString)
	if len(elements) > 0 {
		json.NewEncoder(w).Encode(elements)
	} else {
		fmt.Fprintf(w, "[]")

	}

}

func HandleAsanaTasksId(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	task := req.Header.Get("id")
	fmt.Println(task)
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/tasks/"+task, nil) // URL-encoded payload
	r.Header.Add("Authorization", "Bearer "+token)
	respuesta, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error haciendo petición: %v", err)
	}
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Printf("Error leyendo respuesta: %v", err)
	}
	respuestaString := string(cuerpoRespuesta)
	json.NewEncoder(w).Encode(GetTask(respuestaString))
}

func HandleAsanaTasksIdStories(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	task := req.Header.Get("id")
	fmt.Println(task)
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/tasks/"+task+"/stories", nil) // URL-encoded payload
	r.Header.Add("Authorization", "Bearer "+token)
	respuesta, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error haciendo petición: %v", err)
	}
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Printf("Error leyendo respuesta: %v", err)
	}
	respuestaString := string(cuerpoRespuesta)
	elements := GetStoriesFilter(respuestaString, "comment")
	if len(elements) > 0 {
		json.NewEncoder(w).Encode(elements)
	} else {
		fmt.Fprintf(w, "[]")
	}
}

func HandleAsanaTasksIdDependencies(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	token := req.Header.Get("token")
	task := req.Header.Get("id")
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, "https://app.asana.com/api/1.0/tasks/"+task+"/dependecies", nil) // URL-encoded payload
	r.Header.Add("Authorization", "Bearer "+token)

	respuesta, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error haciendo petición: %v", err)
	}
	defer respuesta.Body.Close()
	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Printf("Error leyendo respuesta: %v", err)
	}
	respuestaString := string(cuerpoRespuesta)

	elements := GetGeneral(respuestaString)
	if len(elements) > 0 {
		json.NewEncoder(w).Encode(elements)
	} else {
		fmt.Fprintf(w, "[]")

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
