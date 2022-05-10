package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

type ObjectGenericDB struct {
	Id string `json:"id"`
}

type dbManager interface {
	insert() error
	get(id string) error
	delete() error
}

func newConnect() *sql.DB {
	db, err := sql.Open("mysql", "tfm:tfm@tcp(localhost:3306)/hqa")
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
func insertDB(db dbManager) error {
	err := db.insert()
	return err
}

func getDB(db dbManager, id string) error {
	err := db.get(id)
	return err
}

func deleteDB(db dbManager) error {
	err := db.delete()
	return err
}

func (car *Car) insert() error {
	var db = newConnect()
	car.Id = car.Brand[1:3] + strconv.Itoa(rand.Intn(1000)) + car.Model[1:3]
	_, err := db.Query(fmt.Sprintf("INSERT INTO cars VALUES ( '%s' ,'%s','%s',%d );", car.Id, car.Brand, car.Model, car.Horse_power))
	if err != nil {
		_, err = db.Query(fmt.Sprintf("UPDATE cars SET brand = '%s' , model = '%s', horse_power = %d  WHERE  id = '%s';", car.Brand, car.Model, car.Horse_power, car.Id))
		if err != nil {
			return err
		}
	}
	defer db.Close()
	return err
}

func (car *Car) get(id string) error {
	var db = newConnect()
	err := db.QueryRow("SELECT id,brand,model,horse_power FROM cars WHERE id = ?", id).Scan(&car.Id, &car.Brand, &car.Model, &car.Horse_power)
	if err != nil {
		return err
	}
	defer db.Close()
	return err
}

func (car *Car) delete() error {
	var db = newConnect()
	_, err := db.Query(fmt.Sprintf("DELETE FROM cars WHERE id = '%s'", car.Id))
	if err != nil {
		return err
	}
	defer db.Close()
	return err
}

///////TFM

func (auth *Auth) getUserBasic(id string, email string) error {

	var db = newConnect()
	response, err := db.Query(fmt.Sprintf("SELECT id FROM user WHERE idGoogle = '%v' and email ='%v'", id, email))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for response.Next() {
		response.Scan(&auth.Id)
	}
	response, err = db.Query(fmt.Sprintf("SELECT user, pass FROM user_basic WHERE userId = '%v'", auth.Id))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for response.Next() {
		response.Scan(&auth.User, &auth.Pass)
	}
	defer db.Close()
	return err
}

func (token *Token) insert(access bool, source bool) error {

	var db = newConnect()
	var err error
	token.Id = uuid.NewV4().String()
	token.State = "active"
	table := "user_token"
	if !access {
		table = "user_refresh_token"
	}
	err = setTableToken(db, *token, table, source)
	if err != nil {
		return err
	}
	defer db.Close()
	return err
}

func setTableToken(db *sql.DB, token Token, table string, source bool) error {

	var o ObjectGenericDB

	response, err := db.Query(fmt.Sprintf("SELECT id FROM %v where state in ('%v', '%v') AND userId = '%v';", table, token.State, "refresh", token.UserId))
	if err != nil {
		return err
	}
	for response.Next() {
		response.Scan(&o.Id)
		if o.Id != "" {
			_, err = db.Query(fmt.Sprintf("UPDATE  %v SET state = '%v' WHERE id = '%v';", table, "inactive", o.Id))
			if err != nil {
				return err
			}
		}
	}
	if source {

		token.State = "refresh"
	}

	_, err = db.Query(fmt.Sprintf("INSERT INTO %v VALUES ( '%v', '%v', '%v',%v, '%v' );", table, token.Id, token.UserId, token.Token, token.ExpireAt, token.State))
	if err != nil {
		return err
	}
	return err
}

func (token *Token) getToken(access bool) error {

	var db = newConnect()
	table := "user_token"
	if !access {
		table = "user_refresh_token"
	}
	response, err := db.Query(fmt.Sprintf("SELECT id,token FROM %v where state = '%v' AND userId = '%v';", table, "active", token.UserId))
	if err != nil {
		return err
	}
	for response.Next() {
		response.Scan(&token.Id, &token.Token)
	}
	if token.Id == "" {
		return fmt.Errorf("token no found")
	}
	return nil
}

func (token *Token) deleteToken(access bool) error {

	var db = newConnect()
	table := "user_token"
	if !access {
		table = "user_refresh_token"
	}
	_, err := db.Query(fmt.Sprintf("UPDATE  %v SET state = '%v' WHERE id = '%v';", table, "inactive", token.Id))
	if err != nil {
		return err
	}
	return nil
}
