package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var Id int
var FileName = "categories.txt"

func main() {

	AddCategoryInFile(FileName, "Pakistan")
	AddCategoryInFile(FileName, "INDIA")
	AddCategoryInFile(FileName, "CANADA")
	AddCategoryInFile(FileName, "IRAN")
	AddCategoryInFile(FileName, "IRAQ")

	fmt.Println(GetAllCategories(FileName))
	fmt.Println(UpdateCategoryById(FileName, 3, "China"))
	fmt.Println(GetCategoryById(FileName, 3))

}

// CreateCategoryApi	create new category in txt file
func CreateCategoryApi(w http.ResponseWriter, r *http.Request) {
	requestJSON := make(map[string]interface{})
	responseJSON := make(map[string]interface{})
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
		responseJSON["message"] = "category name is required"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	_, err = AddCategoryInFile(FileName, name.(string))
	if err != nil {
		fmt.Println("Error adding category in file")
		responseJSON["message"] = "cant update"
		responseJSON["status"] = 500
		json.NewEncoder(w).Encode(responseJSON)
		return
	}

	responseJSON["message"] = "Category added successfully"
	responseJSON["status"] = http.StatusOK
	json.NewEncoder(w).Encode(responseJSON)
	return

}

// GetAllCategoriesApi	get all categories data from txt file
func GetAllCategoriesApi(w http.ResponseWriter, r *http.Request) {
	responseJSON := make(map[string]interface{})

	categories := GetAllCategories(FileName)

	responseJSON["categories"] = categories
	responseJSON["status"] = http.StatusOK
	json.NewEncoder(w).Encode(responseJSON)
}

// GetCategoryApi	get single category data from txt file
func GetCategoryApi(w http.ResponseWriter, r *http.Request) {
	responseJSON := make(map[string]interface{})

	//check request params
	categoryId := r.URL.Query().Get("id")
	convertedId, _ := strconv.ParseInt(categoryId, 10, 64)

	category := GetCategoryById(FileName, int(convertedId))
	responseJSON["data"] = category
	responseJSON["message"] = "success"
	responseJSON["status"] = http.StatusOK
	json.NewEncoder(w).Encode(responseJSON)
	return
}

// UpdateCategoryApi	Update category data by id, stored in txt file
func UpdateCategoryApi(w http.ResponseWriter, r *http.Request) {
	requestJSON := make(map[string]interface{})
	responseJSON := make(map[string]interface{})
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
	categoryId, ok := requestJSON["id"]
	if !ok {
		responseJSON["message"] = "category id is required"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	updatedName, ok := requestJSON["new_name"]
	if !ok {
		responseJSON["message"] = "new name is required"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	convertedId, err := strconv.ParseInt(categoryId.(string), 10, 64)
	if err != nil {
		responseJSON["message"] = "Internal Server error"
		responseJSON["status"] = "500"
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	data := UpdateCategoryById(FileName, int(convertedId), updatedName.(string))

	responseJSON["data"] = data
	responseJSON["message"] = "Updated successfully"
	responseJSON["status"] = http.StatusOK
	json.NewEncoder(w).Encode(responseJSON)
	return
}

func DeleteCategoryApi(w http.ResponseWriter, r *http.Request) {
	responseJSON := make(map[string]interface{})

	categoryId := r.URL.Query().Get("id")
	convertedId, _ := strconv.ParseInt(categoryId, 10, 64)
	DeleteCategoryById(FileName, int(convertedId))

	responseJSON["message"] = "Category deleted successfully"
	responseJSON["status"] = http.StatusOK
	json.NewEncoder(w).Encode(responseJSON)
	return
}

func LoadFileCategories(fileName string) ([]map[string]interface{}, error) {

	var existingData []map[string]interface{}

	// Read the existing file content
	file := fileName
	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Decode the content into an array of maps
	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &existingData)
		if err != nil {
			fmt.Println("Error decoding data:", err)
			return nil, err
		}
	}

	return existingData, nil
}

func AddCategoryInFile(fileName string, categoryName string) ([]map[string]interface{}, error) {

	//var lastContentId interface{}
	existingData, err := LoadFileCategories(FileName)
	if err != nil {
		log.Println("Unable to Load file ", err.Error())
		return nil, err
	}

	if len(existingData) == 0 {
		Id = 1
	} else {
		lastContent := existingData[len(existingData)-1]
		Id = int(lastContent["id"].(float64)) + 1
	}

	newMap := map[string]interface{}{
		"id":   Id,
		"name": categoryName,
	}

	// Append a new map to the array
	existingData = append(existingData, newMap)
	if err = UpdateFileContent(fileName, existingData); err != nil {
		return existingData, err
	}
	return existingData, nil
}

func GetCategoryById(fileName string, id int) map[string]interface{} {

	existingData, _ := LoadFileCategories(fileName)

	if len(existingData) > 0 {
		for _, content := range existingData {
			if content["id"] == float64(id) {
				return content
			}
		}
	}

	return nil
}
func UpdateCategoryById(fileName string, id int, name string) map[string]interface{} {

	existingData, _ := LoadFileCategories(fileName)

	if len(existingData) > 0 {
		for _, content := range existingData {
			if content["id"] == float64(id) {
				content["name"] = name
			}
		}
	}
	UpdateFileContent(fileName, existingData)

	return nil
}

func GetAllCategories(fileName string) []map[string]interface{} {

	existingData, err := LoadFileCategories(fileName)
	if err != nil {
		return nil
	}
	return existingData

}

func UpdateFileContent(fileName string, existingData []map[string]interface{}) error {

	// Encode the updated array back to JSON
	updatedData, err := json.Marshal(existingData)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return err
	}

	// Write the JSON data back to the file
	err = ioutil.WriteFile(fileName, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return err

}

func DeleteCategoryById(fileName string, id int) []map[string]interface{} {
	existingData, _ := LoadFileCategories(fileName)
	fmt.Println(existingData)

	if len(existingData) > 0 {
		for i, content := range existingData {
			if content["id"] == float64(id) {
				existingData = append(existingData[:i], existingData[i+1:]...)
			}
		}
	}
	fmt.Println(existingData)
	UpdateFileContent(fileName, existingData)
	return existingData
}
