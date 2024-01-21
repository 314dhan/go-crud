package main

import (
	"go-web-native/config"
	"go-web-native/controller/categoriescontroller"
	"go-web-native/controller/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)

	log.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
