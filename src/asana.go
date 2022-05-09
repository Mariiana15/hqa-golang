package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

var oautCodehUrl = "https://app.asana.com/-/oauth_authorize?"
var oautUrl = "https://app.asana.com/-/oauth_token"
var projects = "https://app.asana.com/api/1.0/projects"
var tasks = "https://app.asana.com/api/1.0/tasks"
var sections = "https://app.asana.com/api/1.0/sections"

type Asana struct {
	ClientId      string `json:"clientId"`
	ClientSecret  string `json:"clientSecrect"`
	RedirectUri   string `json:"redirect_uri"`
	TimeAsyncTask int16  `json:"timeAsyncTask"`
}

type General struct {
	Gid  string `json:"gid"`
	Name string `json:"name"`
}

type Story struct {
	Gid  string `json:"gid"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type CustomField struct {
	Gid   string `json:"gid"`
	Name  string `json:"name"`
	Value string `json:"display_value"`
}

type Task struct {
	Gid         string        `json:"gid"`
	Name        string        `json:"name"`
	Notes       string        `json:"notes"`
	Project     []General     `json:"projects"`
	CustomField []CustomField `json:"custom_fields"`
	Link        string        `json:"permalink_url"`
	Story       []Story       `json:"stories"`
	Dependecies []General     `json:"dependencies"`
	State       string        `json:"state"`
	TypeTest    string        `json:"typeTest"`
	TypeUS      string        `json:"typeUS"`
	UserStory   string        `json:"userStory"`
	Priority    int           `json:"priority"`
	Alerts      int           `json:"alerts"`
	Scripts     int           `json:"scripts"`
	Date        int64         `json:"date"`
	UrlAlert    string        `json:"urlAlert"`
	UrlScript   string        `json:"urlScript"`
	AddInfo     bool          `json:"addInfo"`
	Result      Result        `json:"result"`
}
type Result struct {
	Message   string `json:"message"`
	Alert     int    `json:"alert"`
	UrlAlert  string `json:"urlAlert"`
	Detail    string `json:"detail"`
	Script    string `json:"script"`
	UrlScript string `json:"urlScript"`
}

type Section struct {
	Name      string `json:"name"`
	Gid       string `json:"gid"`
	StoryUser []Task `json:"storyUser"`
}

func (asana *Asana) GetProperties() {

	path, _ := filepath.Abs("../configuration/config.json")
	file, _ := ioutil.ReadFile(path)
	var result map[string]interface{}
	json.Unmarshal([]byte(file), &result)
	byteData, _ := json.Marshal(result["asana"])
	json.Unmarshal(byteData, &asana)
}

func GetCode(asana Asana) (string, error) {

	v, err := CreateCodeVerifier()
	var message string
	if err != nil {
		return "", err
	}
	code_verifier := v.String()
	code_challenge := v.CodeChallengeS256()
	code_challenge_method := "S256"
	message = fmt.Sprintf("{\"url\": \"%vclient_id=%v&redirect_uri=%v&response_type=code&state=thisIsARandomString&code_challenge_method=%v&code_challenge=%v&scope=default\",\"code_verifier\":\"%v\"}", oautCodehUrl, asana.ClientId, asana.RedirectUri, code_challenge_method, code_challenge, code_verifier)
	return message, nil
}

func GetParamsOauth(code string, codeVerifier string, asana Asana) *strings.Reader {

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", asana.ClientId)
	data.Set("client_secret", asana.ClientSecret)
	data.Set("redirect_uri", asana.RedirectUri)
	data.Set("code", code)
	data.Set("code_verifier", codeVerifier)
	return strings.NewReader(data.Encode())
}

func OauthAsana(code string, codeVerifier string) *http.Request {

	var asana Asana
	asana.GetProperties()
	r, _ := http.NewRequest(http.MethodPost, oautUrl, GetParamsOauth(code, codeVerifier, asana))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return r
}

func ProjectsAsana(token string) *http.Request {

	r, _ := http.NewRequest(http.MethodGet, projects, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	return r
}

func SectionsAsana(token string, project string) *http.Request {

	url := fmt.Sprintf("%v/%v/sections", projects, project)
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	return r
}

func SectionsAsanaId(token string, sectionId string) *http.Request {

	url := fmt.Sprintf("%v/%v", sections, sectionId)
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	return r
}

func TaskSectionAsana(token string, section string) (*http.Request, int16) {

	r, _ := http.NewRequest(http.MethodGet, tasks, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	values := r.URL.Query()
	values.Add("section", section)
	r.URL.RawQuery = values.Encode()
	var asana Asana
	asana.GetProperties()
	return r, asana.TimeAsyncTask
}

func TaskAsana(token string, task string) *http.Request {

	url := fmt.Sprintf("%v/%v", tasks, task)
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	return r
}

func StoriesAsana(token string, task string) *http.Request {

	url := fmt.Sprintf("%v/%v/stories", tasks, task)
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	return r
}

func DependenciesAsana(token string, task string) *http.Request {

	url := fmt.Sprintf("%v/%v/dependecies", tasks, task)
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	return r
}

func GetGeneral(respuestaString string) []General {
	var response map[string]interface{}
	var projects []General
	json.Unmarshal([]byte(respuestaString), &response)
	byteData, _ := json.Marshal(response["data"])
	json.Unmarshal(byteData, &projects)
	return projects
}

func GetGeneralUnd(respuestaString string) General {
	var response map[string]interface{}
	var projects General
	json.Unmarshal([]byte(respuestaString), &response)
	byteData, _ := json.Marshal(response["data"])
	json.Unmarshal(byteData, &projects)
	return projects
}

func GetStories(respuestaString string) []Story {
	var response map[string]interface{}
	var story []Story
	json.Unmarshal([]byte(respuestaString), &response)
	byteData, _ := json.Marshal(response["data"])
	json.Unmarshal(byteData, &story)
	return story
}

func GetStoriesFilter(respuestaString string, value string) []Story {
	var response map[string]interface{}
	var story []Story
	var storyResponse []Story
	json.Unmarshal([]byte(respuestaString), &response)
	byteData, _ := json.Marshal(response["data"])

	json.Unmarshal(byteData, &story)
	for i := len(story) - 1; i >= 0; i-- {
		if story[i].Type == value {
			storyResponse = append(storyResponse, story[i])
		}
	}
	return storyResponse
}

func GetTask(respuestaString string) Task {
	var response map[string]interface{}
	var tasks Task
	json.Unmarshal([]byte(respuestaString), &response)
	byteData, _ := json.Marshal(response["data"])
	json.Unmarshal(byteData, &tasks)
	return tasks
}

func getTaskAsync(t string, token string, task string, rc chan *http.Request) {

	var r *http.Request
	if t == "stories" {
		r = StoriesAsana(token, task)
	} else if t == "dependencies" {
		r = DependenciesAsana(token, task)

	} else {
		r = TaskAsana(token, task)
	}
	rc <- r
}
