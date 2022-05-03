package main

func main() {
	server := NewServer(":8081")
	server.Handle("GET", "/", HandleRoot)

	server.Handle("GET", "/asana/code", server.AddMiddleware(HandleAsanaCode, CheckAuth()))
	server.Handle("GET", "/asana/projects", server.AddMiddleware(HandleAsanaProjects, CheckAuth()))
	server.Handle("GET", "/asana/sections", server.AddMiddleware(HandleAsanaSections, CheckAuth()))
	server.Handle("GET", "/asana/tasks", server.AddMiddleware(HandleAsanaTasks, CheckAuth()))
	server.Handle("GET", "/asana/tasks/:id", server.AddMiddleware(HandleAsanaTasksId, CheckAuth()))
	server.Handle("GET", "/asana/tasks/:id/stories", server.AddMiddleware(HandleAsanaTasksIdStories, CheckAuth()))
	server.Handle("GET", "/asana/tasks/:id/dependecies", server.AddMiddleware(HandleAsanaTasksIdDependencies, CheckAuth()))
	server.Handle("GET", "/asana/sections/tasks", server.AddMiddleware(HandleAsanaSectionsTasks, CheckAuth()))
	server.Handle("POST", "/asana/oauth", server.AddMiddleware(HandleAsanaOauth, CheckAuth()))

	server.Handle("POST", "/cars", server.AddMiddleware(CarPostRequest, CheckAuth(), CheckBodyCar(), Loggin()))
	server.Handle("GET", "/cars/:id", server.AddMiddleware(CarGetRequest, CheckAuth(), Loggin()))
	server.Handle("DELETE", "/cars/:id", server.AddMiddleware(CarDeleteRequest, CheckAuth(), Loggin()))
	server.Listen()
}
