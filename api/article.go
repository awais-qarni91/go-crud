package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"Content"`
}

var id = 0
var ArticleMap = make(map[int]*Article)

// GeArticle function getting article by id
func GeArticle(i int) *Article {
	return ArticleMap[i]
}

// DelArticle function deleting article by id
func DelArticle(i int) {
	delete(ArticleMap, i)
}

// CreateArticle - API to create article in local memory
func CreateArticle(w http.ResponseWriter, r *http.Request) {
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
	title, ok := requestJSON["title"]
	if !ok {
		responseJSON["message"] = "title is required"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	content, ok := requestJSON["content"]
	if !ok {
		responseJSON["message"] = "Content is required"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}

	id = id + 1
	newArticle := &Article{ID: id, Title: title.(string), Content: content.(string)}
	ArticleMap[id] = newArticle

	responseJSON["message"] = "Article entered successfully."
	responseJSON["entered_id"] = id
	responseJSON["status"] = 200
	json.NewEncoder(w).Encode(responseJSON)

}

// GetAllArticles - get all articles saved in local memory.
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	responseJSON := make(map[string]interface{})

	fmt.Println(ArticleMap)
	articles := make([]interface{}, 0)
	for _, v := range ArticleMap {
		fmt.Println(v)
		articles = append(articles, v)
	}
	responseJSON["message"] = "Success"
	responseJSON["articles"] = articles
	responseJSON["status"] = 200
	json.NewEncoder(w).Encode(responseJSON)
}

// GetArticle - get single article by id
func GetArticle(w http.ResponseWriter, r *http.Request) {

	responseJSON := make(map[string]interface{})

	//check request params
	articleId := r.URL.Query().Get("id")
	convertedId, _ := strconv.ParseInt(articleId, 10, 64)
	if _, ok := ArticleMap[int(convertedId)]; ok {
		responseJSON["message"] = "Success"
		responseJSON["articles"] = ArticleMap[int(convertedId)]
		responseJSON["status"] = 200
		json.NewEncoder(w).Encode(responseJSON)
		return
	}

	responseJSON["message"] = "invalid article id"
	responseJSON["article"] = make([]interface{}, 0)
	responseJSON["status"] = http.StatusBadRequest
	json.NewEncoder(w).Encode(responseJSON)
	return
}

// UpdateArticle - update article by id, stored in local memory.
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
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
	articleId, ok := requestJSON["id"]
	if !ok {
		responseJSON["message"] = "article id is required"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}
	var artId int
	switch articleId.(type) {
	case float64:
		artId = int(int64(articleId.(float64)))
		break
	default:
		responseJSON["message"] = "wrong id type"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}

	temp := GeArticle(artId)
	title, titleBool := requestJSON["title"]
	if titleBool {
		temp.Title = title.(string)
	}
	content, contentBool := requestJSON["content"]
	if contentBool {
		temp.Content = content.(string)
	}

	if !titleBool && !contentBool {
		responseJSON["message"] = "nohting to update"
		responseJSON["status"] = http.StatusNotFound
		json.NewEncoder(w).Encode(responseJSON)
		return
	}

	responseJSON["message"] = "Success"
	responseJSON["articles"] = temp
	responseJSON["status"] = 200
	json.NewEncoder(w).Encode(responseJSON)
	return
}

// DeleteArticle - delete article by id, stored in local memory
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	responseJSON := make(map[string]interface{})

	articleId := r.URL.Query().Get("id")
	convertedId, _ := strconv.ParseInt(articleId, 10, 64)
	DelArticle(int(convertedId))
	responseJSON["message"] = "article deleted successfully"
	responseJSON["status"] = http.StatusOK
	json.NewEncoder(w).Encode(responseJSON)
	return
}
