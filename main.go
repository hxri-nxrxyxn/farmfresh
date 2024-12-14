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
		panic(err)
	}
	fmt.Println("Connected Successfully")
	return db, nil
}

type users struct {
	User_id int `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	User_type string `json:"user_type"`
}




