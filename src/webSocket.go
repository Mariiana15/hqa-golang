package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

var protocolBD = "hqaProtocol"
var protocolDBTimer = 60000

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	NewWebSocket(w, r)
}

func HandleProtocol(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	protocol := req.URL.Query().Get("q")
	if protocol == "tasks" {
		fmt.Fprintf(w, "{\"protocol\": \"%v\",\"timer\":\"%v\"}", protocolBD, protocolDBTimer)
	}
}

/// dispatcher

func Dispatcher(ws *websocket.Conn, m Message) {
	switch m.Type {
	case "tasks":
		handelCreateSectionTask(ws, m)
		defer handelGeteSectionTask(ws, m)
	case "select":
		log.Println("Get task Dispatcher")
		handelGeteSectionTask(ws, m)
	default:
		SendMessageWS(ws, "It is connect")
	}
}

////

func getTaskSection(client *http.Client, m Message) ([]General, int16, error) {

	r, t := TaskSectionAsana(m.Message.Token, m.Message.ObjectId)
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return nil, 0, err
	}
	elements := GetGeneral(res)
	return elements, t, nil
}

func getSection(client *http.Client, m Message) (Section, error) {

	r := SectionsAsanaId(m.Message.Token, m.Message.ObjectId)
	res, err := GetBodyResponseRequest(client, r)
	var sectionId Section

	if err != nil {
		return sectionId, err
	}
	sectionId = GetSectionId(res)
	return sectionId, nil
}

func handelGeteSectionTask(ws *websocket.Conn, m Message) {

	var sections []Section
	errDB := getUserStoriesComplete(&sections, m.User)
	if errDB != nil {
		handlerExceptionWS(ws, errDB.Error(), stateKO)
		ws.Close()
		return
	}
	ws.WriteJSON(sections)
	ws.Close()
}

func handelCreateSectionTask(ws *websocket.Conn, m Message) {

	var sections []Section
	client := &http.Client{}
	elements, t, _ := getTaskSection(client, m)
	sectionId, _ := getSection(client, m)
	sectionId.ID = uuid.NewV4().String()

	if len(elements) > 0 {
		tasks, err := HandleAsanaSectionsTasksWS(client, ws, elements, m.Message.Token, &sectionId, t, sections, m.User)
		if err != nil {
			log.Println(err)
			handlerExceptionWS(ws, err.Error(), stateKO)
			ws.Close()
			return
		}
		sectionId.StoryUser = tasks
	}
	errDB := sectionId.setSectionProject(m.User)
	if errDB != nil {
		handlerExceptionWS(ws, errDB.Error(), stateKO)
		ws.Close()
		return
	}
	sections = append(sections, sectionId)
	ws.WriteJSON(sections)
	ws.Close()
}

func HandleAsanaSectionsTasksWS(client *http.Client, ws *websocket.Conn, elements []General, token string, sectionObj *Section, timeService int16, sectionList []Section, user string) ([]Task, error) {

	var tasks []Task
	timeCurrent := 3000
	timeCurrentSend := 1

	for i := 0; i <= len(elements)-1; i++ {

		var task Task
		rt := make(chan *http.Request)
		rs := make(chan *http.Request)
		rd := make(chan *http.Request)

		go getTaskAsync("task", token, elements[i].Gid, rt)
		go getTaskAsync("stories", token, elements[i].Gid, rs)
		go getTaskAsync("dependencies", token, elements[i].Gid, rd)

		rst := <-rt
		res, err := GetBodyResponseRequest(client, rst)
		if err != nil {
			return nil, err
		}
		task = GetTask(res)
		errDB := task.setUserStoryAsana(sectionObj.ID)
		if errDB != nil {
			return nil, err
		}
		errDB = task.setUserStoryAsanaCField()
		if errDB != nil {
			return nil, err
		}

		rss := <-rs
		resSt, err := GetBodyResponseRequest(client, rss)
		if err != nil {
			return nil, err
		}
		elements_ := GetStoriesFilter(resSt, "comment")
		task.Story = elements_
		errDB = task.setUserStoryAsanaStories()
		if errDB != nil {
			return nil, err
		}

		rsd := <-rd
		resDep, err := GetBodyResponseRequest(client, rsd)
		if err != nil {
			return nil, err
		}
		elements_dep := GetGeneral(resDep)
		task.Dependecies = elements_dep
		errDB = task.setUserStoryAsanaDependence()
		if errDB != nil {
			return nil, err
		}
		errTask := createUserStoryHQA(&task, user)
		if errTask != nil {
			return nil, errTask
		}

		/*errTaskR := createUserStoryResultHQA(&task)
		if errTaskR != nil {
			return nil, errTaskR
		}*/
		if int(timeService)*(i+1) > timeCurrent {

			timeCurrent = timeCurrent * timeCurrentSend
			timeCurrentSend++
			tasks = append(tasks, task)
			sectionObj.StoryUser = tasks
			sectionList = append(sectionList, *sectionObj)
			ws.WriteJSON(sectionList)
		} else {
			tasks = append(tasks, task)
		}

	}
	return tasks, nil
}

func createUserStoryHQA(task *Task, user string) error {

	task.UserId = user
	task.State = "open" // revisar con la base de datos
	test, errDB := getTestHQA()
	if errDB != nil {
		return errDB
	}
	task.TypeTest = test.Name                          // revisar con la base de datos
	task.TypeTestId = test.Gid                         // revisar con la base de datos
	task.TypeUS = "alert"                              // revisar con la base de datos
	task.UserStory = task.Notes                        // revisar con la base de datos
	task.Priority = 45                                 // revisar con la base de datos
	task.Alerts = 2                                    // revisar con la base de datos
	task.Scripts = 1                                   // revisar con la base de datos
	task.UrlAlert = "www.google.com"                   // revisar con la base de datos
	task.UrlScript = "http://localhost:3000/dashboard" // revisar con la base de datos
	task.Date = time.Now().Unix()                      // revisar con la base de datosu
	task.AddInfo = 1                                   // revisar con la base de datos
	log.Println(task.UserId)
	errDB = task.setUserStory()
	if errDB != nil {
		return errDB
	}
	return nil
}

func createUserStoryResultHQA(task *Task) error {

	//task.State = "close"
	task.Result.Message = "Succesful"                                            // revisar con la base de datos
	task.Result.Alert = 1                                                        // revisar con la base de
	task.Result.UrlAlert = "http://localhost:3000/dashboard"                     // revisar con la base de datos
	task.Result.Detail = "Aqui llegara la informacion de las pruebas realizadas" // revisar con la base de datos
	task.Result.Script = 1                                                       // revisar con la base de datos
	task.Result.UrlScript = "http://localhost:3000/dashboard"                    // revisar con la base de datos

	errDB := task.setUserStoryResult()
	if errDB != nil {
		return errDB
	}
	return nil
}
