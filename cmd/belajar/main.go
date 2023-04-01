package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func decodeJSONToMap() {
	var jsonString = `{"id": 1, "name": "John Doe", "age": 30}`

	var jsonData = map[string]interface{}{}

	err := json.Unmarshal([]byte(jsonString), &jsonData)

	if err != nil {
		fmt.Printf("Error when unmarshal json: %v", err)
		panic(err)
	}

	fmt.Printf("User: %v", jsonData)

	fmt.Printf("age: %v", jsonData["age"])

}

func decodeJSONToStruct() {
	var user User

	var jsonString = `{"id": 1, "name": "John Doe", "age": 30}`

	err := json.Unmarshal([]byte(jsonString), &user)

	if err != nil {
		fmt.Printf("Error when unmarshal json: %v", err)
		panic(err)
	}

	fmt.Printf("User: %v", user)
}

func encodeStructToJSON() {
	var user = User{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}

	var jsonData, err = json.Marshal(user)

	if err != nil {
		fmt.Printf("Error when marshal json: %v", err)
		panic(err)
	}

	var jsonString = string(jsonData)

	fmt.Printf("Json: %v", jsonString)
}

type User struct {
	ID   int
	Name string
	Age  int
}

func parseURLString() {
	var urlString = "https://www.educastudio.com/category/marbel-edu-games?page=1&limit=12&filter=bermain&tags=edukasi&tags=belajar&tags=mainan&tags=anak"

	var baseURL, err = url.Parse(urlString)

	if err != nil {
		fmt.Printf("Error when parse url: %v", err)
		panic(err)
	}

	var path = baseURL.Path
	var category = strings.Split(path, "/")[2]

	fmt.Printf("Category: %v ", category)

	var query = baseURL.Query()
	var page = query.Get("page")
	var limit = query.Get("limit")
	var filter = query.Get("filter")
	var tags = query["tags"]

	fmt.Printf("Page: %v ", page)
	fmt.Printf("Limit: %v ", limit)
	fmt.Printf("Filter: %v ", filter)
	fmt.Printf("Tags: %v ", tags)
}

func main() {
	//decodeJSONToMap()
	//decodeJSONToStruct()
	//encodeStructToJSON()
	parseURLString()
}
