package main

import (
	"net/http"
	"strings"
)

type Router struct {
	rules     map[string]map[string]http.HandlerFunc
	pathParam map[string]string
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (router_ *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := router_.rules[path]
	exist, path = router_.pathParams(path, exist)
	handler, methodExist := router_.rules[path][method]
	return handler, methodExist, exist
}

func (router_ *Router) pathParams(path string, exist bool) (bool, string) {
	if !exist {
		router_.pathParam = make(map[string]string)
		currentRoute := strings.Split(path, "/")
		for route := range router_.rules {
			exist = router_.validateParams(route, currentRoute)
			if exist {
				return exist, route
			}
		}
	}
	return exist, path
}

func (router_ *Router) validateParams(route string, currentRoute []string) bool {
	if strings.Contains(route, ":") && len(currentRoute) == len(strings.Split(route, "/")) {
		var error bool
		for i, sroute := range strings.Split(route, "/") {
			if strings.Contains(sroute, ":") {
				router_.pathParam[strings.Split(sroute, ":")[1]] = currentRoute[i]
				continue
			}
			if sroute != currentRoute[i] {
				error = true
			}
		}
		if !error {
			return true
		}
	}
	return false
}

func corsRoute(w http.ResponseWriter) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token, token, projectId"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
}

func (router_ *Router) ServeHTTP(write_ http.ResponseWriter, request *http.Request) {

	corsRoute(write_)
	if request.Method == "OPTIONS" {
		write_.WriteHeader(http.StatusOK)
	} else {
		hanlder, methodExist, exist := router_.FindHandler(request.URL.Path, request.Method)
		if !exist {
			write_.WriteHeader(http.StatusNotFound)
			return
		}

		if !methodExist {
			write_.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		for k, v := range router_.pathParam {
			request.Header.Set(k, v)
		}
		hanlder(write_, request)
	}

}
