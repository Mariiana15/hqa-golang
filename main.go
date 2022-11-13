package main

import (
	"os"
)

func main() {

	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")

	server := NewServer(":8080")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/health", HandleRoot)
	server.Handle("GET", "/app", HandleWebSocket)

	server.Handle("POST", "/openai/build", server.AddMiddleware(HandleOpenBuildHQA, CheckAuthToken()))
	server.Handle("POST", "/openai/industry", server.AddMiddleware(HandleOpenAI, CheckAuthToken()))
	server.Handle("POST", "/openai/context-user-story", server.AddMiddleware(HandleOpenAI_US_Context, CheckAuthToken()))
	server.Handle("POST", "/openai/context-tech", server.AddMiddleware(HandleOpenAI_US_Operation, CheckAuthToken()))
	server.Handle("POST", "/openai/context-tech-up", server.AddMiddleware(HandleOpenAI_US_OperationUpdate, CheckAuthToken()))
	server.Handle("POST", "/openai/functions", server.AddMiddleware(HandleOpenAIFunctions, CheckAuthToken()))
	server.Handle("POST", "/openai/process", server.AddMiddleware(HandleOpenAIProcess, CheckAuthToken()))
	server.Handle("POST", "/openai/risk", server.AddMiddleware(HandleOpenAIRisk, CheckAuthToken()))
	server.Handle("POST", "/openai/keywords", server.AddMiddleware(HandleOpenAIKeywords, CheckAuthToken()))
	server.Handle("POST", "/openai/keywordsAtack", server.AddMiddleware(HandleOpenAIKeyWordsAtack, CheckAuthToken()))
	server.Handle("POST", "/openai/programs", server.AddMiddleware(HandleOpenAIPrograms, CheckAuthToken()))
	server.Handle("POST", "/openai/keywordsBBDD", server.AddMiddleware(HandleOpenAIKeywordsBBDD, CheckAuthToken()))
	server.Handle("POST", "/openai/keywordsDesign", server.AddMiddleware(HandleOpenAIKeywordsDesign, CheckAuthToken()))
	server.Handle("POST", "/openai/limit", server.AddMiddleware(HandleOpenAILimit, CheckAuthToken()))
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

	server.Handle("POST", "/login", server.AddMiddleware(HandleLogin, HandlerResponse()))
	server.Handle("POST", "/token/refresh", server.AddMiddleware(HandleRefresh, HandlerResponse()))
	server.Handle("POST", "/token/logout", server.AddMiddleware(HandleLogOut, CheckAuthToken()))

	server.Handle("GET", "/hack/protocol", server.AddMiddleware(HandleProtocol, CheckAuthToken()))
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
