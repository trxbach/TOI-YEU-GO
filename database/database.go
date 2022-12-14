package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func New(connection *string) (*DB, error) {
	default_conn := os.Getenv("DB_DSN")
	if connection == nil {
		if default_conn == "" {
			return nil, nil
		}
		connection = &default_conn
	}
	db_handle, err := sql.Open("mysql", *connection)
	if err != nil {
		return nil, err
	}

	db := DB{
		db_handle,
	}

	return &db, nil
}

func CreateTables(db *DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS accounts (
					   id INT AUTO_INCREMENT,
					   handle TEXT,
					   password TEXT,
					   
					   PRIMARY KEY (id)
					   )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tests (
					  id INT AUTO_INCREMENT,
					  name TEXT,
					  start BIGINT,
					  end BIGINT,

					  PRIMARY KEY (id)
					  )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS questions (
					  id INT AUTO_INCREMENT,
					  content TEXT,
					  
					  idt INT,
					  
					  PRIMARY KEY (id)
					  )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS choices (
					  id INT AUTO_INCREMENT,
					  content TEXT,
					  idq INT,
					  is_answer BOOLEAN,
					  PRIMARY KEY (id)
					  )`)
	if err != nil {
		return err
	}

	return nil
}

func ResetTables(db *DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS accounts")
	if err != nil {
		return err
	}

	_, err = db.Exec("DROP TABLE IF EXISTS choices")
	if err != nil {
		return err
	}

	_, err = db.Exec("DROP TABLE IF EXISTS questions")
	if err != nil {
		return err
	}

	_, err = db.Exec("DROP TABLE IF EXISTS tests")
	if err != nil {
		return err
	}

	err = CreateTables(db)
	if err != nil {
		return err
	}

	return nil
}
