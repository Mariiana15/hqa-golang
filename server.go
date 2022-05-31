package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (server_ *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := server_.router.rules[path]
	if !exist {
		server_.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	server_.router.rules[path][method] = handler
}

func (server_ *Server) AddMiddleware(hf http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, mw := range middlewares {
		hf = mw(hf)
	}
	return hf
}

func (server_ *Server) Listen() error {
	http.Handle("/", server_.router)
	err := http.ListenAndServe(server_.port, nil)
	if err != nil {
		return err
	}
	return nil
}
