package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Advice struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./advicebank.db")
	if err != nil {
		return err
	}

	advicetable := `CREATE TABLE IF NOT EXISTS advices (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"description" TEXT,
		"author" TEXT);`
	query, err := db.Prepare(advicetable)

	if err != nil {
		return err
	}

	query.Exec()

	DB = db
	return nil
}

func Getalladvices() ([]Advice, error) {
	rows, err := DB.Query("SELECT id, description, author from advices")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	advicelist := make([]Advice, 0)

	for rows.Next() {
		singleadvice := Advice{}
		err = rows.Scan(&singleadvice.Id, &singleadvice.Description, &singleadvice.Author)

		if err != nil {
			return nil, err
		}

		advicelist = append(advicelist, singleadvice)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return advicelist, err
}

func GeteachAdvice(id string) (Advice, error) {
	newId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	stmt, err := DB.Prepare("SELECT id, description, author FROM advices WHERE id = ?")

	if err != nil {
		return Advice{}, err
	}

	advice := Advice{}

	sqlErr := stmt.QueryRow(newId).Scan(&advice.Id, &advice.Description, &advice.Author)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Advice{}, nil
		}
		return Advice{}, sqlErr
	}
	return advice, nil
}

func Addadvcie(newAdvice Advice) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO advices (description, author) VALUES (?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newAdvice.Description, newAdvice.Author)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func Updateanadvice(advice Advice, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE advices SET description = ?, author = ? WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(advice.Description, advice.Author, id)

	if err != nil {

		return false, err
	}

	tx.Commit()

	return true, nil
}

func GetrandomAdvice() (Advice, error) {
	rows, err := DB.Query("SELECT id, description, author from advices")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	i := 0
	for rows.Next() {
		i = i + 1
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	min := 1
	max := i + 1
	fmt.Println(max)
	id := rand.Intn(max-min) + min
	fmt.Println(id)

	stmt, err := DB.Prepare("SELECT id, description, author FROM advices WHERE id = ?")

	if err != nil {
		return Advice{}, err
	}

	advice := Advice{}

	sqlErr := stmt.QueryRow(id).Scan(&advice.Id, &advice.Description, &advice.Author)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Advice{}, nil
		}
		return Advice{}, sqlErr
	}
	return advice, nil
}
