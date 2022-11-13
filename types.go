package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type InputOpenAI struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Temperature float32 `json:"temperature"`
	MaxTokens   int16   `json:"max_tokens"`
	Topp        int16   `json:"top_p"`
	Frequency   float32 `json:"frequency_penalty"`
	Presence    float32 `json:"presence_penalty"`
}

type requestOpenAI struct {
	Text     string `json:"text"`
	Options  string `json:"options"`
	Id       string `json:"id"`
	Auxiliar string `json:"auxiliar"`
}

const OpenIA_Industry = "Clasifica las industrias en esto texto:\n"

func GetOpenAIConfig(t string) InputOpenAI {

	var i InputOpenAI
	path, _ := filepath.Abs("./configuration/config.json")
	file, _ := ioutil.ReadFile(path)
	var result map[string]interface{}
	json.Unmarshal([]byte(file), &result)
	byteData, _ := json.Marshal(result[t])
	json.Unmarshal(byteData, &i)
	return i
}
