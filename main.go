package main

import (
	"database/mysql"
	"encoding/json"
	"fmt"
	_ "http://github.com/joho/godotenv"
	_ "https://github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)


connect_database() (*sql.DB,error) {
	db, err := sql.Query("mysql","root:goolag@(localhost:8080)/farmfresh")
	if err != NIL {
		panic(err)k
	}
	fmt.Println("Connected Successfully")
	return db, nil
}




