package db

import (
	"database/sql"
	"fmt"
	"techsoc-scrumboard-backend/utils"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Db          *sql.DB
	Initialized bool
}

var dbStructInstance DB

func InitializeDB() {
	db, err := sql.Open("sqlite3", "trello.db")
	utils.HandleError(err)
	sqlInit := `
    CREATE TABLE IF NOT EXISTS users (
      userId INTEGER PRIMARY KEY,
      username TEXT NOT NULL UNIQUE,
      password TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS workspaces (
      workspaceId INTEGER PRIMARY KEY,
      title TEXT NOT NULL,
      description TEXT NOT NULL,
      user INTEGER NOT NULL,
      FOREIGN KEY(user)
        REFERENCES users (userId)
    );
    CREATE TABLE IF NOT EXISTS boards (
      boardId INTEGER PRIMARY KEY,
      title TEXT NOT NULL,
      description TEXT NOT NULL,
      workspace INTEGER NOT NULL,
      FOREIGN KEY(workspace) 
        REFERENCES workspaces (workspaceId)
    );
    CREATE TABLE IF NOT EXISTS cards (
      cardId INTEGER PRIMARY KEY,
      title TEXT NOT NULL,
      description TEXT NOT NULL,
      status TEXT CHECK(status in ('THINGS_TO_DO','DOING','REVIEW','COMPLETED')) NOT NULL,
      board INTEGER NOT NULL,
      FOREIGN KEY(board) 
        REFERENCES boards (boardId)
    );
 `
	result, err := db.Exec(sqlInit)
	fmt.Println(result)
	utils.HandleError(err)
	dbStructInstance.Db = db
	dbStructInstance.Initialized = true
}

func GetDB() DB {
	if !dbStructInstance.Initialized {
		InitializeDB()
	}
	return dbStructInstance
}
