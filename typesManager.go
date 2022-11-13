package main

type Industry struct {
	KeyWords string `json:"keywords"`
	Name     string `json:"industry"`
	Context  string `json:"context"`
	Risks    string `json:"risks"`
}

type ContextUserStory struct {
	UserStory      string `json:"userStory"`
	Context        string `json:"context"`
	Keywords       string `json:"keywords"`
	Test           string `json:"test"`
	Contemplations string `json:"contemplations"`
	IdUserStory    string `json:"idUserStory"`
}

type StakeHoldertUserStory struct {
	Industry    string `json:"industry"`
	Job         string `json:"job"`
	Risk        string `json:"risk"`
	Functions   string `json:"functions"`
	Test        string `json:"test"`
	IdUserStory string `json:"idUserStory"`
}

type OperationsUserStory struct {
	CriteriosTech string `json:"criteriosTech"`
	Security      string `json:"security"`
	Technologies  string `json:"technologies"`
	Process       string `json:"process"`
	Database      string `json:"ddbb"`
	Design        string `json:"design"`
	Risk          string `json:"risk"`
	IdUserStory   string `json:"idUserStory"`
	Id            string `json:"id"`
	CreateAt      string `json:"createAt"`
}

type OwaspUserStory struct {
	Technologies string `json:"technologies"`
	Industries   string `json:"industries"`
	BBDD         string `json:"bbdd"`
	IdUserStory  string `json:"idUserStory"`
}

type userst struct {
	Tiltle          string `json:"tiltle"`
	ContextoCliente string `json:"contextoCliente"`
	Contexto        string `json:"contexto"`
	Validaciones    string `json:"validaciones"`
	Cases           string `json:"cases"`
	Tests           string `json:"tests"`
	KeyWords        string `json:"keywords"`
	Risk            string `json:"risk"`
	RiskIndustry    string `json:"riskIndustry"`
}
