package main

import (
	"os"

	"github.com/Mariiana15/apis"
	"github.com/Mariiana15/serverutils"
)

func main() {

	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")

	server := NewServer(":8080")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/health", HandleRoot)
	server.Handle("GET", "/app", apis.HandleWebSocket)

	//server.Handle("GET", "/app", server.AddMiddleware(HandleRoot2, CheckAuthWebSocket()))
	server.Handle("GET", "/asana/code", server.AddMiddleware(HandleAsanaCode, CheckAuthToken()))
	server.Handle("GET", "/asana/code/v1", server.AddMiddleware(HandleAsanaCodeDB, CheckAuthToken()))

	server.Handle("GET", "/asana/projects", server.AddMiddleware(HandleAsanaProjects, CheckAuthToken()))
	server.Handle("GET", "/asana/sections", server.AddMiddleware(HandleAsanaSections, CheckAuthToken()))
	server.Handle("GET", "/asana/sections/:id", server.AddMiddleware(HandleAsanaSections, CheckAuthToken()))
	server.Handle("GET", "/asana/tasks", server.AddMiddleware(HandleAsanaTasks, CheckAuthToken()))
	server.Handle("GET", "/asana/tasks/:id", server.AddMiddleware(HandleAsanaTasksId, CheckAuthToken()))
	server.Handle("GET", "/asana/tasks/:id/stories", server.AddMiddleware(HandleAsanaTasksIdStories, CheckAuthToken()))
	server.Handle("GET", "/asana/tasks/:id/dependecies", server.AddMiddleware(HandleAsanaTasksIdDependencies, CheckAuthToken()))
	server.Handle("GET", "/asana/sections/tasks", server.AddMiddleware(HandleAsanaSectionsTasks, CheckAuthToken()))
	server.Handle("POST", "/asana/oauth", server.AddMiddleware(HandleAsanaOauth, CheckAuthToken()))

	server.Handle("POST", "/login", server.AddMiddleware(serverutils.HandleLogin, HandlerResponse()))
	server.Handle("POST", "/token/refresh", server.AddMiddleware(serverutils.HandleRefresh, HandlerResponse()))
	server.Handle("POST", "/token/logout", server.AddMiddleware(serverutils.HandleLogOut, CheckAuthToken()))

	server.Handle("GET", "/hack/protocol", server.AddMiddleware(apis.HandleProtocol, CheckAuthToken()))
	server.Handle("POST", "/hack/us/tech", server.AddMiddleware(HandleParamsTech, CheckAuthToken()))
	server.Handle("POST", "/hack/us/state", server.AddMiddleware(HandleChangeStateUserStory, CheckAuthToken()))
	server.Handle("POST", "/hack/us/section/state", server.AddMiddleware(HandleChangeStateSection, CheckAuthToken()))
	server.Handle("POST", "/hack/us/result", server.AddMiddleware(HandleResultUserStory, CheckAuthToken()))
	server.Handle("GET", "/hack/us/validate", server.AddMiddleware(HandleGetValidateUStory, CheckAuthToken()))

	server.Handle("POST", "/cars", server.AddMiddleware(CarPostRequest, CheckAuth(), CheckBodyCar(), Loggin()))
	server.Handle("GET", "/cars/:id", server.AddMiddleware(CarGetRequest, CheckAuth(), Loggin()))
	server.Handle("DELETE", "/cars/:id", server.AddMiddleware(CarDeleteRequest, CheckAuth(), Loggin()))
	server.Listen()

}
