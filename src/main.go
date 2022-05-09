package main

import "os"

func main() {

	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")

	server := NewServer(":8081")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/app", HandleRoot2)

	//server.Handle("GET", "/app", server.AddMiddleware(HandleRoot2, CheckAuthWebSocket()))
	server.Handle("GET", "/asana/code", server.AddMiddleware(HandleAsanaCode, CheckAuth()))
	server.Handle("GET", "/asana/projects", server.AddMiddleware(HandleAsanaProjects, CheckAuth()))
	server.Handle("GET", "/asana/sections", server.AddMiddleware(HandleAsanaSections, CheckAuth()))
	server.Handle("GET", "/asana/sections/:id", server.AddMiddleware(HandleAsanaSections, CheckAuth()))
	server.Handle("GET", "/asana/tasks", server.AddMiddleware(HandleAsanaTasks, CheckAuth()))
	server.Handle("GET", "/asana/tasks/:id", server.AddMiddleware(HandleAsanaTasksId, CheckAuth()))
	server.Handle("GET", "/asana/tasks/:id/stories", server.AddMiddleware(HandleAsanaTasksIdStories, CheckAuth()))
	server.Handle("GET", "/asana/tasks/:id/dependecies", server.AddMiddleware(HandleAsanaTasksIdDependencies, CheckAuth()))
	server.Handle("GET", "/asana/sections/tasks", server.AddMiddleware(HandleAsanaSectionsTasks, CheckAuth()))
	server.Handle("POST", "/asana/oauth", server.AddMiddleware(HandleAsanaOauth, CheckAuth()))

	server.Handle("POST", "/login", server.AddMiddleware(HandleLoginHQA, CheckAuth()))
	server.Handle("POST", "/todo", server.AddMiddleware(HandleRoot3, CheckAuthToken()))

	server.Handle("POST", "/test/webhook", HandleRoot3)
	server.Handle("POST", "/token/logout", server.AddMiddleware(HandleLogOut, CheckAuthToken()))
	server.Handle("POST", "/token/refresh", server.AddMiddleware(HandleRefresh))

	server.Handle("GET", "/hack/protocol", server.AddMiddleware(HandleProtocol, CheckAuthToken()))

	server.Handle("POST", "/cars", server.AddMiddleware(CarPostRequest, CheckAuth(), CheckBodyCar(), Loggin()))
	server.Handle("GET", "/cars/:id", server.AddMiddleware(CarGetRequest, CheckAuth(), Loggin()))
	server.Handle("DELETE", "/cars/:id", server.AddMiddleware(CarDeleteRequest, CheckAuth(), Loggin()))
	server.Listen()

}
