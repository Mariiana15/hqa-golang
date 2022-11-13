package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Mariiana15/dbmanager"
)

func HandleRoot(write_ http.ResponseWriter, req *http.Request) {
	write_.WriteHeader(http.StatusOK)
}
func HandleOpenBuildHQA(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	var rop requestOpenAI
	elements, err := BuildUserStory(req, &rop)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}
func HandleOpenAI(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	var rop requestOpenAI
	elements, err := GetIndustryOIA(req, &rop)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAI_US_Context(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusAccepted)
	var rop requestOpenAI
	elements, err := GetUserStoryContext(req, &rop)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAI_US_Operation(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusAccepted)
	var rop requestOpenAI
	elements, err := GetOperationContext(req, &rop)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}
func HandleOpenAI_US_OperationUpdate(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusAccepted)
	var rop requestOpenAI
	elements, err := GetOperationUpdateContext(req, &rop)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAIProcess(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	elements, err := GetProcess(req)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAIFunctions(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	elements, err := GetFunctions(req)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAIRisk(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	elements, err := GetRisk(req)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAIKeywords(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	elements, err := GetKeywords(req)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAIKeyWordsAtack(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	elements, err := GetKeywordsAtack(req)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAIPrograms(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	elements, err := GetPrograms(req)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAIKeywordsBBDD(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	elements, err := GetKeywordsBBDD(req)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAIKeywordsDesign(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	elements, err := GetKeywordsDesign(req)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

}

func HandleOpenAILimit(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	elements, err := GetLimit(req)
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err.Error())
	}
	json.NewEncoder(w).Encode(elements)

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

func HandleAsanaCodeDB(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	tokenString := ExtractToken(req)
	acc, err2 := ExtractTokenMetadataWS(tokenString)
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
	result, _ := GetBodyResponse(req)
	code := result["code"].(string)
	code_verifier := result["code_verifier"].(string)

	tokenString := ExtractToken(req)
	_, err2 := ExtractTokenMetadataWS(tokenString)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", err2)
		return
	}

	client := &http.Client{}
	r := OauthAsana(code, code_verifier)
	res, err := GetBodyResponseRequest(client, r)
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

func HandleAsanaSectionsId(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	client := &http.Client{}
	token := req.Header.Get("token")
	section := req.Header.Get("id")
	r := SectionsAsanaId(token, section)
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

func HandleAsanaSectionsTasksAsync(w http.ResponseWriter, req *http.Request, elements []dbmanager.General, token string, section string) {

	client := &http.Client{}
	var tasks []dbmanager.Task
	for i := len(elements) - 1; i >= 0; i-- {
		var task dbmanager.Task
		r := make(chan *http.Request)
		r2 := make(chan *http.Request)
		r3 := make(chan *http.Request)

		go GetTaskAsync("task", token, elements[i].Gid, r)
		go GetTaskAsync("stories", token, elements[i].Gid, r2)
		go GetTaskAsync("dependencies", token, elements[i].Gid, r3)

		rr := <-r
		res, err := GetBodyResponseRequest(client, rr)
		if err != nil {
			fmt.Fprintf(w, "%v\"%v\"}", res, err)
		}
		task = GetTask(res)

		rr2 := <-r2
		res2, err := GetBodyResponseRequest(client, rr2)
		if err != nil {
			fmt.Fprintf(w, "%v\"%v\"}", res, err)
		}
		elements_ := GetStoriesFilter(res2, "comment")
		task.Story = elements_

		rr3 := <-r3
		res3, err := GetBodyResponseRequest(client, rr3)
		if err != nil {
			fmt.Fprintf(w, "%v\"%v\"}", res, err)
		}
		elements_dep := GetGeneral(res3)
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
	r, t := TaskSectionAsana(token, section)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Fprintf(w, "%v\"%v\"}", res, err)
	}
	elements := GetGeneral(res)
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
	r, _ := TaskSectionAsana(token, section)
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
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", MsgDatabase)
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
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", MsgNotFound)
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
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", MsgNotFound)
		return
	}
	err = dbmanager.DeleteDB(&car)
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", MsgDatabase)
		return
	}
	write_.WriteHeader(http.StatusNonAuthoritativeInfo)
	responseCarBody(&car, write_)
}

func responseCarBody(car *dbmanager.Car, write_ http.ResponseWriter) {

	response, err := car.ToJson()
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", MsgMalFormat)
		return
	}
	write_.Write(response)
}
