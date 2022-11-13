package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

func newConnect() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	db, err := sql.Open("mysql", "tfm:"+os.Getenv("CONNECTION_STRING"))
	if err != nil {
		fmt.Println(err.Error())
	}
	return db

}

//BBDD
func GetIndustryHQA(ind string) (Industry, error) {

	var db = newConnect()
	var r Industry
	i := 0
	response, err := db.Query(fmt.Sprintf("SELECT risks, keywords, context FROM industry_points where industry= '%v'", ind))
	if err != nil {
		return r, err
	}
	for response.Next() {
		response.Scan(&r.Risks, &r.KeyWords, &r.Context)
		i++
	}
	defer db.Close()
	if i == 0 {
		return r, fmt.Errorf("0 datos")
	}
	return r, nil
}

func SetIndustryHQA(ind Industry) error {

	var db = newConnect()
	//atCreate := time.Now().Unix()
	Id := uuid.NewV4().String()
	_, err := db.Query(fmt.Sprintf("INSERT industry_points VALUES ('%v','%v','%v','%v','%v','%v')", Id, ind.Name, ind.Risks, ind.KeyWords, ind.Context, time.Now().Format("2006/1/2")))
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func SetUserStoryContext(c ContextUserStory) error {

	var db = newConnect()
	//atCreate := time.Now().Unix()
	Id := uuid.NewV4().String()
	_, err := db.Query(fmt.Sprintf("INSERT user_story_context VALUES ('%v','%v','%v','%v','%v','%v','%v', '%v')", Id, c.UserStory, c.Context, c.Keywords, c.Test, c.Contemplations, time.Now().Format("2006/1/2"), c.IdUserStory))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	return nil
}

func SetStakeHolder(c *StakeHoldertUserStory) error {

	var db = newConnect()
	//atCreate := time.Now().Unix()
	Id := uuid.NewV4().String()
	_, err := db.Query(fmt.Sprintf("INSERT user_story_stakeholder VALUES ('%v','%v','%v','%v','%v','%v','%v')", Id, c.Industry, c.Job, c.Risk, c.Functions, c.Test, time.Now().Format("2006/1/2")))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	return nil
}

func SetOperationUserStory(c *OperationsUserStory) error {

	var db = newConnect()
	//atCreate := time.Now().Unix()
	Id := uuid.NewV4().String()
	_, err := db.Query(fmt.Sprintf("INSERT user_story_tech_operations VALUES ('%v','%v','%v','%v','%v','%v','%v','%v','%v', '%v')", Id, c.CriteriosTech, c.Security, c.Technologies, c.Process, c.Database, c.Design, c.Risk, time.Now().Format("2006/1/2"), c.IdUserStory))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	return nil
}

func UpdateOperationUserStory(c *OperationsUserStory) error {

	var db = newConnect()
	fmt.Println(c.Database)
	_, err := db.Query(fmt.Sprintf("Update user_story_tech_operations SET security = '%v',technologies ='%v', ddbb ='%v', design = '%v',risk ='%v' where  idUserStory ='%v'", c.Security, c.Technologies, c.Database, c.Design, c.Risk, c.IdUserStory))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	return nil
}

func GetOperationTechId(id string) (OperationsUserStory, error) {

	var db = newConnect()
	var c OperationsUserStory
	fmt.Println(id)
	response, err := db.Query(fmt.Sprintf("Select id, criteriosTech, security, technologies, process, ddbb, design, risk, createAt, idUserStory from user_story_tech_operations where idUserStory = '%v'", id))
	if err != nil {
		return c, err
	}
	i := 0
	for response.Next() {
		response.Scan(&c.Id, &c.CriteriosTech, &c.Security, &c.Technologies, &c.Process, &c.Database, &c.Design, &c.Risk, &c.CreateAt, &c.IdUserStory)
		i++
	}
	defer db.Close()
	if i == 0 {
		return c, fmt.Errorf("0 datos")
	}
	return c, nil
}

func GetUserStorysContext(id string) (ContextUserStory, error) {

	var db = newConnect()
	var c ContextUserStory
	response, err := db.Query(fmt.Sprintf("Select context, keywords, test, contemplations from user_story_context where idUserStory = '%v'", id))
	if err != nil {
		return c, err
	}
	i := 0
	for response.Next() {
		response.Scan(&c.Context, &c.Keywords, &c.Test, &c.Contemplations)
		i++
	}
	defer db.Close()
	if i == 0 {
		return c, fmt.Errorf("0 datos")
	}
	return c, nil
}

func GetUserStorysStakeHolder(industry string) (StakeHoldertUserStory, error) {

	var db = newConnect()
	var c StakeHoldertUserStory
	response, err := db.Query("Select industry, risk, functions, test from hqa.user_story_stakeholder where job like '%" + industry + "%'")
	if err != nil {
		fmt.Println(err.Error())
		return c, err
	}
	i := 0
	for response.Next() {
		response.Scan(&c.Industry, &c.Risk, &c.Functions, &c.Test)
		i++
	}
	defer db.Close()
	if i == 0 {
		fmt.Println("0 datos")
		return c, fmt.Errorf("0 datos")
	}
	return c, nil
}
func GetUserStorysOwasp(id string) (OwaspUserStory, error) {

	var db = newConnect()
	var c OwaspUserStory
	response, err := db.Query(fmt.Sprintf("Select industries, bbdd from hqa.user_story_owasp where idUserStory = '%v'", id))
	if err != nil {
		return c, err
	}
	i := 0
	for response.Next() {
		response.Scan(&c.Industries, &c.BBDD)
		i++
	}
	defer db.Close()
	if i == 0 {
		return c, fmt.Errorf("0 datos")
	}
	return c, nil
}

func SetOperationOwasp(c *OwaspUserStory) error {

	var db = newConnect()
	//atCreate := time.Now().Unix()
	Id := uuid.NewV4().String()
	_, err := db.Query(fmt.Sprintf("INSERT user_story_owasp VALUES ('%v','%v','%v','%v','%v', '%v')", Id, c.Technologies, c.Industries, c.BBDD, time.Now().Format("2006/1/2"), c.IdUserStory))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	return nil
}

//
