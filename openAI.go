package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)////////////////////////////////
/////////////////////////////////
func GetIndustryV2(req *http.Request, st string) (string, error) {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Industria_v2")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "org-jZhovXfi693JsdyJR781cHOK")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return "", err
	} else {
		elements := GetGeneralOpenaAI(res)
	
		if len(elements.Choices) > 0 {
			return elements.Choices[0].Text, nil
		} else {

			return "", fmt.Errorf("No se ha recuperado los Test de cada Stakeholder")
		}
	}
}

///////////////////////////////////
/////////////////////////////
func GetPorcentural(req *http.Request, lt string, st string) (int, error) {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Afinidad")
	b.Prompt = b.Prompt + lt + "\n" + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return 0, err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			r, _ := regexp.Compile("\\d")
			match := r.FindAllString(elements.Choices[0].Text, 3)
			return strconv.Atoi(strings.Join(match, ""))
		} else {

			return 0, fmt.Errorf("No se ha recuperado las pruebas")
		}
	}
}

func GetUserCasesSecurity(req *http.Request, st string) (string, error) {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Security_Cases")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return "", err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			return elements.Choices[0].Text, nil
		} else {

			return "", fmt.Errorf("No se ha recuperado las pruebas")
		}
	}
}

func GetKeywordsPart(req *http.Request, st string) (string, error) {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_20")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return "", err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			return elements.Choices[0].Text, nil
		} else {

			return "", fmt.Errorf("No se ha recuperado las pruebas")
		}
	}
}

func GetUserStoryTiltle(req *http.Request, st string, id string, us *userst) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Tiltle")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			us.Tiltle = id + " - " + elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado las pruebas")
		}
	}
}

func BuildUserStory(req *http.Request, rop *requestOpenAI) (userst, error) {

	var usJson userst
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	if rop.Id == "" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return usJson, err
		}
		req.Body.Close()
		var rq requestOpenAI
		err = json.Unmarshal([]byte(body), &rq)
		if err != nil {
			return usJson, err
		}
		rop.Text = rq.Text
		rop.Id = rq.Id
		rop.Options = rq.Options
		rop.Auxiliar = rq.Auxiliar
	}
	var usJsom userst
	b := GetOpenAIConfig("config_OpenIA_StakeHolderJob")
	st := rop.Text

	GetUserStoryTiltle(req, st, rop.Id, &usJsom)
	c, _ := GetOperationTechId(rop.Id)
	usCt, _ := GetUserStorysContext(rop.Id)
	ow, _ := GetUserStorysOwasp(rop.Id)
	id, _ := GetIndustryHQA(rop.Auxiliar)
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return usJson, err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			var us StakeHoldertUserStory
			p := 0
			op := strings.Split(elements.Choices[0].Text, "\n")
			lt, _ := GetKeywordsPart(req, usCt.Context)
			for _, opt := range op {
				r, _ := regexp.Compile("\\d.")
				match := r.FindString(opt)
				op1 := strings.Replace(opt, match, "", 2)
				if op1 != "" && r.MatchString(opt) {
					if strings.Contains(op1, " ") {
						if len(strings.Split(op1, " ")[0]) < 4 {
							if len(strings.Split(op1, " ")[len(strings.Split(op1, " "))-1]) > 4 {
								op1 = strings.Split(op1, " ")[len(strings.Split(op1, " "))-1]
							} else {
								op1 = strings.Split(op1, " ")[len(strings.Split(op1, " "))-2]
							}
						}
					}
					us, _ = GetUserStorysStakeHolder(op1)
					p, _ = GetPorcentural(req, lt, st)
					if p > 50 {
						usJsom.ContextoCliente = op1
						break
					}
				}

			}
			p, _ = GetPorcentural(req, usCt.Contemplations, usCt.Keywords)
			if p > 0 {
				usJsom.Contexto = usCt.Contemplations
			} else {
				usJsom.Contexto = usCt.Context
			}
			p, _ = GetPorcentural(req, usCt.Test, usCt.Keywords)
			if p > 0 {
				usJsom.Validaciones = usCt.Test
			}
			usJsom.Cases = c.Risk
			tt, _ := GetUserCasesSecurity(req, c.CriteriosTech)
			usJsom.Tests = tt
			usJsom.KeyWords = ow.Industries
			usJsom.Risk = us.Risk
			usJsom.RiskIndustry = id.Risks

			return usJsom, nil
		} else {

			return usJsom, fmt.Errorf("No se ha recuperado el contexto")
		}
	}
}

func GetUserStory_(req *http.Request, st string) (string, error) {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Operations_Owasp")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return "", err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {

			return elements.Choices[0].Text, nil
		} else {

			return "", fmt.Errorf("No se ha recuperado los Test de cada Stakeholder")
		}
	}
}

func GetOperationOwasp(req *http.Request, st string) (string, error) {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Operations_Owasp")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return "", err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {

			return elements.Choices[0].Text, nil
		} else {

			return "", fmt.Errorf("No se ha recuperado los Test de cada Stakeholder")
		}
	}
}

func GetOperationRisk(req *http.Request, st string, c *OperationsUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Risk")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Risk = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado los Test de cada Stakeholder")
		}
	}
}

func GetOperationDesign(req *http.Request, st string, c *OperationsUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_KeywordsDesign")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Design = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado los Test de cada Stakeholder")
		}
	}
}

func GetOperationBBDD(req *http.Request, st string, c *OperationsUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_KeywordsBBDD")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Database = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado los Test de cada Stakeholder")
		}
	}
}

func GetOperationProcess(req *http.Request, st string, c *OperationsUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Limit")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Process = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado los Test de cada Stakeholder")
		}
	}
}

func GetOperationSecurity(req *http.Request, st string, c *OperationsUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_KeyWords_from_Industry")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Security = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado los Test de cada Stakeholder")
		}
	}
}

func GetStakeholderTest(req *http.Request, st string, c *StakeHoldertUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Stakeholder_UserStory_Test")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Test = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado los Test de cada Stakeholder")
		}
	}
}

func GetStakeholderKeyActivites(req *http.Request, st string, c *StakeHoldertUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Stakeholder_UserStory_KeyActivities")
	fmt.Println(st)
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Functions = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado las actividades de cada stakeholder")
		}
	}
}

func GetStakeholderActivites(req *http.Request, st string, c *StakeHoldertUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Stakeholder_UserStory_Activities")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			GetStakeholderKeyActivites(req, elements.Choices[0].Text, c)
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado las actividades de cada stakeholder")
		}
	}
}

func GetStakeholderRisk(req *http.Request, st string, c *StakeHoldertUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Stakeholder_UserStory_Vulnerabilidades")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Risk = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado los riesgos de cada stakeholder")
		}
	}
}

func GetUserStoryContemplations(req *http.Request, st string, c *ContextUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Context_UserStory_Contemplations")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Contemplations = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado las contemplaciones")
		}
	}
}

func GetUserStoryTest2(req *http.Request, st string, c *ContextUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Context_UserStory_Test2")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Test = elements.Choices[0].Text + c.Test
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado las pruebas")
		}
	}
}

func GetUserStoryTest1(req *http.Request, st string, c *ContextUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Context_UserStory_Test1")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Test = elements.Choices[0].Text + c.Test
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado las pruebas")
		}
	}
}

func GetUserStoryKeywords(req *http.Request, s string, c *ContextUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Context_UserStory_KeyWords")
	b.Prompt = b.Prompt + s
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Keywords = elements.Choices[0].Text
			return nil
		} else {

			return fmt.Errorf("No se ha recuperado las palabras clave")
		}
	}
}

func GetOperationUpdateContext(req *http.Request, rop *requestOpenAI) (OperationsUserStory, error) {

	var c OperationsUserStory
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	if rop.Id == "" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return c, err
		}
		req.Body.Close()
		var rq requestOpenAI
		err = json.Unmarshal([]byte(body), &rq)
		if err != nil {
			return c, err
		}
		rop.Text = rq.Text
		rop.Id = rq.Id
		rop.Options = rq.Options
		rop.Auxiliar = rq.Auxiliar
	}

	c, _ = GetOperationTechId(rop.Id)
	fmt.Println(c)
	b := GetOpenAIConfig("config_OpenIA_Risk")
	st := rop.Text
	c.Technologies = rop.Options
	var owasp OwaspUserStory
	owasp.IdUserStory = rop.Id
	op := strings.Split(rop.Options, ",")
	for _, opt := range op {
		t, _ := GetOperationOwasp(req, opt)
		owasp.Technologies = owasp.Technologies + t
	}
	op = strings.Split(st, "\n")
	for _, opt := range op {
		t, _ := GetOperationOwasp(req, opt)
		owasp.Industries = owasp.Industries + t
	}

	stAuxiliar := rop.Auxiliar
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return c, err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.Risk = elements.Choices[0].Text
			GetOperationBBDD(req, st, &c)
			c.Database = c.Database + " , " + stAuxiliar
			t, _ := GetOperationOwasp(req, c.Database)
			owasp.BBDD = t
			GetOperationDesign(req, st, &c)
			GetOperationSecurity(req, st, &c)
			UpdateOperationUserStory(&c)
			SetOperationOwasp(&owasp)
			return c, nil
		} else {

			return c, fmt.Errorf("No se ha recuperado el contexto")
		}
	}
}

func GetOperationContext(req *http.Request, rop *requestOpenAI) (OperationsUserStory, error) {

	var c OperationsUserStory
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	if rop.Id == "" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return c, err
		}
		req.Body.Close()
		var rq requestOpenAI
		err = json.Unmarshal([]byte(body), &rq)
		if err != nil {
			return c, err
		}
		rop.Text = rq.Text
		rop.Id = rq.Id
	}
	b := GetOpenAIConfig("config_OpenIA_Operations_Tech")
	st := rop.Text
	c.IdUserStory = rop.Id
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return c, err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.CriteriosTech = elements.Choices[0].Text
			GetOperationProcess(req, st, &c)
			SetOperationUserStory(&c)
			return c, nil
		} else {

			return c, fmt.Errorf("No se ha recuperado el contexto")
		}
	}
}

func GetUserStoryContext(req *http.Request, rop *requestOpenAI) (ContextUserStory, error) {

	var c ContextUserStory
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	if rop.Id == "" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return c, err
		}
		req.Body.Close()
		var rq requestOpenAI
		err = json.Unmarshal([]byte(body), &rq)
		if err != nil {
			return c, err
		}
		rop.Text = rq.Text
		rop.Id = rq.Id
	}

	b := GetOpenAIConfig("config_OpenIA_Context_UserStory_Contextos")
	st := rop.Text
	c.IdUserStory = rop.Id
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return c, err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			c.UserStory = st
			c.Context = elements.Choices[0].Text
			err_ct := GetUserStoryKeywords(req, st, &c)
			if err_ct != nil {
				return c, err_ct
			}
			err_ct = GetUserStoryTest1(req, st, &c)
			if err_ct != nil {
				return c, err_ct
			}
			err_ct = GetUserStoryTest2(req, st, &c)
			if err_ct != nil {
				return c, err_ct
			}
			err_ct = GetUserStoryContemplations(req, st, &c)
			if err_ct != nil {
				return c, err_ct
			}
			SetUserStoryContext(c)
			return c, nil
		} else {

			return c, fmt.Errorf("No se ha recuperado el contexto")
		}
	}
}

func GetIndustryOIA(req *http.Request, rop *requestOpenAI) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	var c StakeHoldertUserStory
	err := godotenv.Load()
	if err != nil {

		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}

	if rop.Id == "" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return a, err
		}
		req.Body.Close()
		var rq requestOpenAI
		err = json.Unmarshal([]byte(body), &rq)
		if err != nil {
			return a, err
		}
		rop.Text = rq.Text
		rop.Id = rq.Id
	}
	b := GetOpenAIConfig("config_OpenIA_Industry")
	st := rop.Text
	c.IdUserStory = rop.Id
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return a, err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				s_opt := strings.Split(opt, " ")
				if s_opt[len(s_opt)-1] != "" && len(s_opt[len(s_opt)-1]) > 6 {
					s = append(s, s_opt[len(s_opt)-1])
					_, errdb := GetIndustryHQA(s_opt[len(s_opt)-1])
					c.Industry = s_opt[len(s_opt)-1]
					GetStakeHolder(req, st, &c)
					if errdb != nil {
						GetRiskForIndustry(req, s_opt[len(s_opt)-1])

					} else {
						println(s_opt[len(s_opt)-1])
					}
				}
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}

func GetRiskForIndustry(req *http.Request, ind string) ([]Industry, error) {

	var i []Industry
	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Risk_from_Industry")
	b.Prompt = b.Prompt + ind
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return i, err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			for _, opt := range op {
				if opt != "" && len(opt) > 30 {
					var ids Industry
					ids.Name = ind
					ids.Risks = opt
					keyword, _ := GetKeyWordsForIndustry(req, opt)
					context, _ := GetContextForIndustry(req, opt)
					ids.KeyWords = keyword
					ids.Context = context
					fmt.Println(opt)
					err := SetIndustryHQA(ids)
					if err != nil {
						fmt.Println(err.Error())
					}
					i = append(i, ids)
				}
			}
			return i, nil
		} else {
			return i, nil
		}
	}
}

func GetKeyWordsForIndustry(req *http.Request, o string) (string, error) {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_KeyWords_from_Industry")
	b.Prompt = b.Prompt + o
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return "", err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			return elements.Choices[0].Text, nil
		} else {
			return "", nil
		}
	}
}

func GetContextForIndustry(req *http.Request, o string) (string, error) {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_Context_from_Industry")
	b.Prompt = b.Prompt + o
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		return "", err
	} else {
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s string
			for _, opt := range op {
				s_opt := strings.Split(opt, ":")
				s = s + " , " + s_opt[0]
			}
			return s, nil
		} else {
			return "", nil
		}
	}
}

func GetStakeHolder(req *http.Request, st string, c *StakeHoldertUserStory) error {

	client := &http.Client{}
	b := GetOpenAIConfig("config_OpenIA_StakeHolderJob")
	b.Prompt = b.Prompt + st
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				if opt != "" {
					s = append(s, opt)
				}
			}
			c.Job = strings.Join(s, ",")
			GetStakeholderRisk(req, c.Job, c)
			GetStakeholderActivites(req, c.Job, c)
			GetStakeholderTest(req, c.Job, c)
			SetStakeHolder(c)
			return nil
		} else {

			return nil
		}
	}
}

func GetFunctions(req *http.Request) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return a, err
	}
	req.Body.Close()
	var rq requestOpenAI
	err = json.Unmarshal([]byte(body), &rq)
	if err != nil {
		return a, err
	}

	b := GetOpenAIConfig("config_OpenIA_Functionality")
	b.Prompt = b.Prompt + rq.Text
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return a, err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				//s_opt := strings.Split(opt, " ")
				s = append(s, opt)
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}
func GetProcess(req *http.Request) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return a, err
	}
	req.Body.Close()
	var rq requestOpenAI
	err = json.Unmarshal([]byte(body), &rq)
	if err != nil {
		return a, err
	}

	b := GetOpenAIConfig("config_OpenIA_Process")
	b.Prompt = b.Prompt + rq.Text
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return a, err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				//s_opt := strings.Split(opt, " ")
				s = append(s, opt)
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}

func GetRisk(req *http.Request) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return a, err
	}
	req.Body.Close()
	var rq requestOpenAI
	err = json.Unmarshal([]byte(body), &rq)
	if err != nil {
		return a, err
	}

	b := GetOpenAIConfig("config_OpenIA_Risk")
	b.Prompt = b.Prompt + rq.Text
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return a, err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				//s_opt := strings.Split(opt, " ")
				s = append(s, opt)
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}

func GetKeywords(req *http.Request) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return a, err
	}
	req.Body.Close()
	var rq requestOpenAI
	err = json.Unmarshal([]byte(body), &rq)
	if err != nil {
		return a, err
	}

	b := GetOpenAIConfig("config_OpenIA_Keywords")
	b.Prompt = b.Prompt + rq.Text
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return a, err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				//s_opt := strings.Split(opt, " ")
				s = append(s, opt)
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}

func GetKeywordsAtack(req *http.Request) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return a, err
	}
	req.Body.Close()
	var rq requestOpenAI
	err = json.Unmarshal([]byte(body), &rq)
	if err != nil {
		return a, err
	}

	b := GetOpenAIConfig("config_OpenIA_KeywordsAtack")
	b.Prompt = b.Prompt + rq.Text
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return a, err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				//s_opt := strings.Split(opt, " ")
				s = append(s, opt)
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}

func GetPrograms(req *http.Request) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return a, err
	}
	req.Body.Close()
	var rq requestOpenAI
	err = json.Unmarshal([]byte(body), &rq)
	if err != nil {
		return a, err
	}

	b := GetOpenAIConfig("config_OpenIA_Programs")
	b.Prompt = b.Prompt + rq.Text
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return a, err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				//s_opt := strings.Split(opt, " ")
				s = append(s, opt)
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}

func GetKeywordsBBDD(req *http.Request) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return a, err
	}
	req.Body.Close()
	var rq requestOpenAI
	err = json.Unmarshal([]byte(body), &rq)
	if err != nil {
		return a, err
	}

	b := GetOpenAIConfig("config_OpenIA_KeywordsBBDD")
	b.Prompt = b.Prompt + rq.Text
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return a, err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				//s_opt := strings.Split(opt, " ")
				s = append(s, opt)
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}

func GetKeywordsDesign(req *http.Request) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return a, err
	}
	req.Body.Close()
	var rq requestOpenAI
	err = json.Unmarshal([]byte(body), &rq)
	if err != nil {
		return a, err
	}

	b := GetOpenAIConfig("config_OpenIA_KeywordsDesign")
	b.Prompt = b.Prompt + rq.Text
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return a, err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				//s_opt := strings.Split(opt, " ")
				s = append(s, opt)
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}

func GetLimit(req *http.Request) (ResponseOpenAI, error) {

	var a ResponseOpenAI
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	client := &http.Client{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return a, err
	}
	req.Body.Close()
	var rq requestOpenAI
	err = json.Unmarshal([]byte(body), &rq)
	if err != nil {
		return a, err
	}

	b := GetOpenAIConfig("config_OpenIA_Limit")
	b.Prompt = b.Prompt + rq.Text
	url := fmt.Sprintf("%v", os.Getenv("URL_OPEN_AI"))
	j, _ := json.Marshal(b)
	r, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN_OPEN_AI"))
	r.Header.Set("openai-organization", "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")
	res, err := GetBodyResponseRequest(client, r)
	if err != nil {
		fmt.Println(err.Error())
		return a, err
	} else {
		fmt.Println(res)
		elements := GetGeneralOpenaAI(res)
		if len(elements.Choices) > 0 {
			op := strings.Split(elements.Choices[0].Text, "\n")
			var s []string
			for _, opt := range op {
				//s_opt := strings.Split(opt, " ")
				s = append(s, opt)
			}
			elements.Options = s
			return elements, nil
		} else {

			return a, nil
		}
	}
}
