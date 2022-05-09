package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var protocolBD = "hqaProtocol"
var protocolDBTimer = 60000

func HandleRoot2(w http.ResponseWriter, r *http.Request) {

	NewWebSocket(w, r)

}

func HandleRoot3(w http.ResponseWriter, r *http.Request) {

	b, err := GetBodyResponse(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "{\"error\": \"%v\"}", msgMalFormat)
		return
	}

	byteData, _ := json.Marshal(b)
	w.Write(byteData)
}

func HandleProtocol(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	protocol := req.URL.Query().Get("q")
	if protocol == "tasks" {
		fmt.Fprintf(w, "{\"protocol\": \"%v\",\"timer\":\"%v\"}", protocolBD, protocolDBTimer)
	}
}

func Dispatcher(ws *websocket.Conn, m Message) {
	switch m.Type {
	case "tasks":
		log.Println("Dispatcher")
		Tasks(ws, m)
	default:
		SendMessageWS(ws, "It is connect")
	}
}

func Tasks(ws *websocket.Conn, m Message) {

	client := &http.Client{}

	r, t := TaskSectionAsana(m.Message.Token, m.Message.ObjectId)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		handlerExceptionWS(ws, err.Error(), stateKO)
		log.Printf("Error: %v", err)
	}
	elements := GetGeneral(res)

	r2 := SectionsAsanaId(m.Message.Token, m.Message.ObjectId)
	res2, err2 := GetBodyResponseRequest(client, r2)
	if err2 != nil {
		handlerExceptionWS(ws, err.Error(), stateKO)
		log.Printf("Error: %v", err)
	}
	sectionId := GetGeneralUnd(res2)

	var sections []Section
	if len(elements) > 0 {
		log.Println("task --elments > 0")

		//BD falta realizar la logica
		var section Section
		section.Gid = m.Message.ObjectId
		tasks := HandleAsanaSectionsTasksWS(ws, elements, m.Message.Token, section, t, sections)
		section.StoryUser = tasks
		section.Name = sectionId.Name
		sections = append(sections, section)
		/// BD 2
		section.Gid = m.Message.ObjectId
		section.Name = section.Name + " -project 2"
		tasks = HandleAsanaSectionsTasksWS(ws, elements, m.Message.Token, section, t, sections)
		section.StoryUser = tasks
		sections = append(sections, section)
		ws.WriteJSON(sections)
		ws.Close()
	} else {

		log.Println("task --elments = 0")
		SendMessageJsonWS(ws, fmt.Sprintf("%v", sections))
		ws.Close()
	}
}

func HandleAsanaSectionsTasksWS(ws *websocket.Conn, elements []General, token string, sectionObj Section, timeService int16, sectionList []Section) []Task {

	client := &http.Client{}
	var tasks []Task
	timeCurrent := 3000
	timeCurrentSend := 1
	for i := 0; i <= len(elements)-1; i++ {
		var task Task
		r := make(chan *http.Request)
		r2 := make(chan *http.Request)
		r3 := make(chan *http.Request)

		go getTaskAsync("task", token, elements[i].Gid, r)
		go getTaskAsync("stories", token, elements[i].Gid, r2)
		go getTaskAsync("dependencies", token, elements[i].Gid, r3)

		rr := <-r
		res, err := GetBodyResponseRequest(client, rr)
		if err != nil {
			handlerExceptionWS(ws, err.Error(), stateKO)
			log.Printf("Error: %v", err)
		}
		task = GetTask(res)

		rr2 := <-r2
		res2, err := GetBodyResponseRequest(client, rr2)
		if err != nil {
			handlerExceptionWS(ws, err.Error(), stateKO)
			log.Printf("Error: %v", err)
		}
		elements_ := GetStoriesFilter(res2, "comment")
		task.Story = elements_

		rr3 := <-r3
		res3, err := GetBodyResponseRequest(client, rr3)
		if err != nil {
			handlerExceptionWS(ws, err.Error(), stateKO)
			log.Printf("Error: %v", err)
		}
		elements_dep := GetGeneral(res3)
		task.Dependecies = elements_dep
		task.State = "open"                                // revisar con la base de datos
		task.TypeTest = "TSH001"                           // revisar con la base de datos
		task.TypeUS = "alert"                              // revisar con la base de datos
		task.UserStory = task.Notes                        // revisar con la base de datos
		task.Priority = 45                                 // revisar con la base de datos
		task.Alerts = 2                                    // revisar con la base de datos
		task.Scripts = 1                                   // revisar con la base de datos
		task.UrlAlert = "www.google.com"                   // revisar con la base de datos
		task.UrlScript = "http://localhost:3000/dashboard" // revisar con la base de datos
		task.Date = time.Now().Unix()                      // revisar con la base de datos

		if int(timeService)*(i+1) > timeCurrent {

			task.Date = time.Now().Add(time.Hour * -240).Unix() // revisar con la base de datos
			timeCurrent = timeCurrent * timeCurrentSend
			timeCurrentSend++
			//task.State = "close"
			task.Alerts = 3                                                              // revisar con la base de datos                                                    // revisar con la base de datos
			task.Result.Message = "Succesful"                                            // revisar con la base de datos
			task.Result.Alert = 1                                                        // revisar con la base de
			task.AddInfo = true                                                          // revisar con la base de datos
			task.Result.UrlAlert = "http://localhost:3000/dashboard"                     // revisar con la base de datos
			task.Priority = 75                                                           // revisar con la base de datos
			task.Result.Detail = "Aqui llegara la informacion de las pruebas realizadas" // revisar con la base de datos
			task.Result.Script = "Script 1 generado"                                     // revisar con la base de datos
			task.Result.UrlScript = "http://localhost:3000/dashboard"                    // revisar con la base de datos
			tasks = append(tasks, task)
			sectionObj.StoryUser = tasks
			sectionList = append(sectionList, sectionObj)
			ws.WriteJSON(sectionList)
		} else {
			tasks = append(tasks, task)
		}
	}
	return tasks

}
