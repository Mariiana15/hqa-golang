package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var oautCodehUrl = "https://app.asana.com/-/oauth_authorize?"

type MainAsana struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecrect"`
	RedirectUri  string `json:"redirect_uri"`
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
}

func (data *MainAsana) GetProperties() {

	path, _ := filepath.Abs("../configuration/config.json")
	file, _ := ioutil.ReadFile(path)
	var result map[string]interface{}
	json.Unmarshal([]byte(file), &result)
	byteData, _ := json.Marshal(result["asana"])
	json.Unmarshal(byteData, &data)
}

func GetCode(data MainAsana) (string, error) {

	v, err := CreateCodeVerifier()
	var message string
	if err != nil {
		return "", err
	}
	code_verifier := v.String()
	code_challenge := v.CodeChallengeS256()
	code_challenge_method := "S256"
	message = fmt.Sprintf("{\"url\": \"%vclient_id=%v&redirect_uri=%v&response_type=code&state=thisIsARandomString&code_challenge_method=%v&code_challenge=%v&scope=default\",\"code_verifier\":\"%v\"}", oautCodehUrl, data.ClientId, data.RedirectUri, code_challenge_method, code_challenge, code_verifier)
	return message, nil
}

func GetGeneral(respuestaString string) []General {
	var response map[string]interface{}
	var projects []General
	json.Unmarshal([]byte(respuestaString), &response)
	byteData, _ := json.Marshal(response["data"])
	json.Unmarshal(byteData, &projects)
	fmt.Println(projects)
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
			fmt.Println(story[i].Type)
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
