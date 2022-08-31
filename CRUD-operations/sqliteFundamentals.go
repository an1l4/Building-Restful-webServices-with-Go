package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//book placeholder
type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)

	if err != nil {
		log.Println(err)
	}

	//create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books(id INTEGER PRIMARY KEY,isbn INTEGER,author VARCHAR(64),name VARCHAR(64)NULL)")
	if err != nil {
		log.Println("error in creating table")
	} else {
		log.Println("successfully created table books ")
	}
	statement.Exec()

	//create
	statement, _ = db.Prepare("INSERT INTO books(name,author,isbn)VALUES (?,?,?)")
	statement.Exec("A Tale of Two Cities", "Charles Dicken", 14056743)
	log.Println("Inserted the book into database")

	//Read
	rows, _ := db.Query("SELECT id,name,author FROM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID : %d, Book: %s,Author :%s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	//update
	statement, _ = db.Prepare("update books set name=? where id=1")
	statement.Exec("The Tail of Two Cities", 1)
	log.Println("Successfully updated the book in database")

	//Delete
	statement, _ = db.Prepare("delete from books where id=1")
	statement.Exec(1)
	log.Println("successfully deleted the book from database")
}
