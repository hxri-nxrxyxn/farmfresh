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
	connect = os.Getenv("CONNECT")
	db, err := sql.Query("mysql",connect)
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

type products struct {
	Product_id int `json:"product_id"`
	Farmer_id int `json:"farmer_id"`
	Product_name string `json:"product_name"`
	Product_description string `json:"product_description"`
	Category string `json:"category"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
	Image_url string `json:"image_url"`
	Location string `json:"location"`
	Status string `json:"status"`
}

type orders struct {
	Order_id int `json:"order_id"`
	Buyer_id int `json:"buyer_id"`
	Order_date int `json:"order_date"`
	Total_amount float64 `json:"total_amount"`
	Status string `json:"status"`
}

type order_items struct {
	Order_item_id int `json:"order_item_id"`
	Order_id int `json:"order_id"`
	Product_id int `json:"product_id"`
	Quantity int `json:"quantity"`
	Unit_price float64 `json:"unit_price"`
}

func register(db *sql.DB,w http.ResponseWriter, r *http.Request) {
	var newuser users
	register_query = os.Getenv("REGISTER_QUERY")
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		panic(err)
	}
	_,err = db.Exec(register_query, newuser.Username, newuser.Password,newuser.First_name,newuser.Last_name,newuser.Email,newuser.Phone,newuser.Address,newuser.User_type)
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := connect_database()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/v1/users/register", func(w http.ResponseWriter, r *http.Request) {
		register(db,w,r)
	})
	http.ListenAndServe((":8080",nil))
	
}
