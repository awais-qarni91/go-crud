package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	Id   int
	Name string
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	responseJSON := make(map[string]interface{})
	res, err := DbConnection.Query("SELECT `id`, `name` FROM `products`")
	defer res.Close()
	if err != nil {
		responseJSON["status"] = 500
		responseJSON["message"] = "Internal server error"
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	var products []Product

	for res.Next() {

		var product Product
		err := res.Scan(&product.Id, &product.Name)

		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)

	}

	responseJSON["products"] = products
	responseJSON["message"] = "Success"
	responseJSON["status"] = 200
	json.NewEncoder(w).Encode(responseJSON)

}

func GetProductInfo(w http.ResponseWriter, r *http.Request) {
	responseJSON := make(map[string]interface{})
	responseJSON["status"] = 200
	responseJSON["message"] = "success"

	//check request params
	categoryId := r.URL.Query().Get("id")
	convertedId, _ := strconv.ParseInt(categoryId, 10, 64)

	res, err := DbConnection.Query("SELECT `id`, `name` FROM `products` WHERE `id`=?", int(convertedId))
	defer res.Close()
	if err != nil {
		responseJSON["status"] = 500
		responseJSON["message"] = "Internal server error"
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	var product Product

	if res.Next() {

		err := res.Scan(&product.Id, &product.Name)

		if err != nil {
			log.Fatal(err)
		}
		responseJSON["products"] = product
		json.NewEncoder(w).Encode(responseJSON)
		return

	} else {

		responseJSON["message"] = "no data found for requested id"
		responseJSON["status"] = 501
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	requestJSON := make(map[string]interface{})
	responseJSON := make(map[string]interface{})
	responseJSON["status"] = 200

	//read request body
	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		message := fmt.Sprintf("Error reading request body: %v", err)
		responseJSON["message"] = message
		responseJSON["status"] = http.StatusBadRequest
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	//convert request body into map
	err = json.Unmarshal(requestBytes, &requestJSON)
	if err != nil {
		fmt.Println("Error while unmarshalling request body")
		message := fmt.Sprintf("Error reading request body: %v", err)
		responseJSON["message"] = message
		responseJSON["status"] = http.StatusBadRequest
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	//check request params
	name, ok := requestJSON["name"]
	if !ok {
		responseJSON["message"] = "product name is required"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	_, err = DbConnection.Exec("INSERT INTO `products`(`name`) VALUES (?)", name)
	if err != nil {
		responseJSON["message"] = "server error"
		responseJSON["status"] = 500
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	responseJSON["message"] = "product added successfully"
	json.NewEncoder(w).Encode(responseJSON)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	responseJSON := make(map[string]interface{})
	responseJSON["status"] = 200

	//check request params
	categoryId := r.URL.Query().Get("id")
	convertedId, _ := strconv.ParseInt(categoryId, 10, 64)
	_, err := DbConnection.Exec("DELETE FROM `products` WHERE `id`=?", convertedId)
	if err != nil {
		responseJSON["message"] = "server error"
		responseJSON["status"] = 500
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	responseJSON["message"] = "Product deleted successfully"
	json.NewEncoder(w).Encode(responseJSON)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	requestJSON := make(map[string]interface{})
	responseJSON := make(map[string]interface{})
	responseJSON["status"] = 200

	//read request body
	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		message := fmt.Sprintf("Error reading request body: %v", err)
		responseJSON["message"] = message
		responseJSON["status"] = http.StatusBadRequest
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	//convert request body into map
	err = json.Unmarshal(requestBytes, &requestJSON)
	if err != nil {
		fmt.Println("Error while unmarshalling request body")
		message := fmt.Sprintf("Error reading request body: %v", err)
		responseJSON["message"] = message
		responseJSON["status"] = http.StatusBadRequest
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	//check request params
	id, ok := requestJSON["id"]
	if !ok {
		responseJSON["message"] = "product id is required"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	name, ok := requestJSON["new_name"]
	if !ok {
		responseJSON["message"] = "product name is required"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	_, err = DbConnection.Exec("UPDATE `products` SET `name`=? WHERE `id`=?", name, id)
	if err != nil {
		fmt.Println(err)
		responseJSON["message"] = "server error"
		responseJSON["status"] = 500
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	responseJSON["message"] = "Product info updated successfully"
	json.NewEncoder(w).Encode(responseJSON)

}
