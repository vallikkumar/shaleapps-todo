package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sql.DB
}

// NewRepository init repo
func NewRepository() *Repository {
	r := &Repository{}
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Printf("faile to connect to database: %v", err)
	}
	r.db = db
	r.init()
	return r
}

func (r *Repository) init() {
	r.db.Exec(`CREATE TABLE IF NOT EXISTS "todos" ("id" integer primary key autoincrement,"message" varchar(255),"complete" bool )`)
}

// ResolveAll resolve all todo data
func (r *Repository) ResolveAll() ([]Todo, error) {
	// query select all data
	todos := []Todo{}
	rows, err := r.db.Query("SELECT id, message, complete FROM todos")
	if err != nil {
		log.Fatal("could not resolve all: %v", err)
		return todos, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Todo
		// scanning data
		err := rows.Scan(&t.ID, &t.Message, &t.Complete)
		if err != nil {
			log.Fatal("failed to scan data: %v", err)
			return todos, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

// ResolveByID resolve todo by id
func (r *Repository) ResolveByID(id int) (Todo, error) {
	var t Todo
	err := r.db.QueryRow("SELECT id, message, complete FROM todos where id = ?", id).Scan(&t.ID, &t.Message, &t.Complete)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", id)
		return Todo{}, err
	case err != nil:
		log.Fatal("query error: %v\n", err)
		return Todo{}, err
	}

	return t, nil
}

// Store store todo data
func (r *Repository) Store(t Todo) error {
	stmt, err := r.db.Prepare("INSERT INTO todos (message, complete) VALUES (?, ?)")
	if err != nil {
		log.Printf("failed to create statement: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(t.Message, t.Complete)
	if err != nil {
		log.Printf("failed to store: %v", err)
	}
	return err
}

//Update todo data
func (r *Repository) Update(id int, t Todo) error {
	result, err := r.db.Exec("UPDATE todos SET message = ?, complete = ? WHERE id = ?", t.Message, t.Complete, id)
	if err != nil {
		log.Printf("failed to update todo: %v", err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to update todo: %v", err)
	}

	if rows != 1 {
		err = fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}
	return err
}

//Remove todo data
func (r *Repository) Remove(id int) error {
	result, err := r.db.Exec("DELETE from todos WHERE id = ?", id)
	if err != nil {
		log.Printf("failed to remove todo: %v", err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to remove todo: %v", err)
	}

	if rows != 1 {
		err = fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}
	return err
}
