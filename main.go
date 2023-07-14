package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-crud/api"
	"log"
	"net/http"
)

func main() {
	//connect database
	var err error
	err = api.InitializeDB()
	if err != nil {
		log.Println("Error connecting database", err.Error())
		return
	} else {
		fmt.Println("DB connected")
	}
	defer api.CloseDB()

	fmt.Println("Server started")

	router := mux.NewRouter()

	//apis for articles stored in local memory
	router.HandleFunc("/CreateArticle", api.CreateArticle).Methods("POST")
	router.HandleFunc("/GetAllArticles", api.GetAllArticles).Methods("GET")
	router.HandleFunc("/GetArticle", api.GetArticle).Methods("GET")
	router.HandleFunc("/UpdateArticle", api.UpdateArticle).Methods("PATCH")
	router.HandleFunc("/DeleteArticle", api.DeleteArticle).Methods("DELETE")

	//apis for Categories stored in txt file
	router.HandleFunc("/CreateCategory", api.CreateCategoryApi).Methods("POST")
	router.HandleFunc("/GetAllCategories", api.GetAllCategoriesApi).Methods("GET")
	router.HandleFunc("/GetCategory", api.GetCategoryApi).Methods("GET")
	router.HandleFunc("/UpdateCategory", api.UpdateCategoryApi).Methods("PATCH")
	router.HandleFunc("/DeleteCategory", api.DeleteCategoryApi).Methods("DELETE")

	//apis for products stored in mysql database
	router.HandleFunc("/GetAllProducts", api.GetAllProducts).Methods("GET")
	router.HandleFunc("/GetProductInfo", api.GetProductInfo).Methods("GET")
	router.HandleFunc("/CreateProduct", api.CreateProduct).Methods("POST")
	router.HandleFunc("/DeleteProduct", api.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/UpdateProduct", api.UpdateProduct).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8080", router))
}
