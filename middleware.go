package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Mariiana15/apis"
	"github.com/Mariiana15/dbmanager"
	"github.com/Mariiana15/serverutils"
)

func CheckAuthWebSocket() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var ws apis.WebsocketHQA
			ws.NewWebSocketHQA()
			if r.Header.Get("Sec-Websocket-Protocol") != ws.Protocol {
				fmt.Println(401)
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "{\"error\": \"%v\"}", serverutils.MsgUnauthorized)
				return
			}
			hf(w, r)
		}
	}
}

func CheckAuthToken() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			err := serverutils.TokenValid(r)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "{\"error\": \"%v\"}", serverutils.MsgUnauthorized)
				return
			}
			_, err = serverutils.ExtractTokenMetadata(r)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "{\"error\": \"%v\"}", serverutils.MsgUnauthorized)
				return
			}
			hf(w, r)
		}
	}
}

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(write_ http.ResponseWriter, request *http.Request) {

			authStr := strings.SplitN(request.Header.Get("Authorization"), " ", 2)
			write_.Header().Set("Content-Type", "application/json")

			if len(authStr) != 2 || authStr[0] != "Basic" {
				write_.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(write_, "{\"error\": \"%v\"}", serverutils.MsgUnauthorized)
				return
			}
			payload, _ := base64.StdEncoding.DecodeString(authStr[1])
			pair := strings.SplitN(string(payload), ":", 2)
			//if len(pair) != 2 || !validateBasic(write_, pair[0], pair[1]) {
			validate := true
			if len(pair) != 2 || !validate {

				write_.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(write_, "{\"error\": \"%v\"}", serverutils.MsgUnauthorized)
				return
			}
			hf(write_, request)
		}
	}
}

func HandlerResponse() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(write_ http.ResponseWriter, request *http.Request) {
			write_.Header().Set("Content-Type", "application/json")
			hf(write_, request)
		}
	}
}

func validateBasic(write_ http.ResponseWriter, username, password string) bool {
	var auth dbmanager.Auth
	err := auth.GetUserBasic("0", "x")
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", serverutils.MsgDatabase)
		return false
	}
	if username == auth.User && password == auth.Pass {
		return true
	}
	return false
}

func cloneBody(request *http.Request) io.ReadCloser {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("w")
	}
	r2 := request.Clone(request.Context())
	request.Body = ioutil.NopCloser(bytes.NewReader(body))
	r2.Body = ioutil.NopCloser(bytes.NewReader(body))
	return r2.Body
}

func CheckBodyCar() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(write_ http.ResponseWriter, request *http.Request) {
			var car dbmanager.Car
			decoder := json.NewDecoder(cloneBody(request))
			err := decoder.Decode(&car)
			write_.Header().Set("Content-Type", "application/json")
			if err != nil {
				write_.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(write_, "{\"error\": \"%v\"}", err)
				return
			}
			msg, err := car.ValidateStructure()
			if err != nil {
				write_.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(write_, "{\"error\": \"%v\"}", msg)
				return
			}
			hf(write_, request)
		}
	}
}

func Loggin() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(write_ http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(request.URL.Path, time.Since(start))
			}()
			hf(write_, request)
		}

	}
}
