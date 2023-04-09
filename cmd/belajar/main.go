package main

import (
	"belajar-go-orm/cmd/belajar/handlers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func curlGetExample() []map[string]interface{} {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", res.Body)

	var result []map[string]interface{}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	json.Unmarshal(body, &result)

	sb := string(body)
	fmt.Println(sb)

	fmt.Printf("%v", result)

	return result
}

func curlPostExample() {
	data := map[string]interface{}{
		"userId": 1,
		"title":  "Belajar Golang",
		"body":   "Belajar Golang itu mudah",
	}

	payload, err := json.Marshal(data)

	client := &http.Client{}

	//bufferString := bytes.NewBuffer(payload)

	readerString := strings.NewReader(string(payload))

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", readerString)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func initRouting() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/students", handlers.ActionStudent)

	http.HandleFunc("/posts", func(writer http.ResponseWriter, request *http.Request) {
		result := curlGetExample()
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(result)
	})

	server := new(http.Server)
	server.Addr = ":9000"

	server.ListenAndServe()

}

func main() {
	//curlPostExample()

	initRouting()
	//decodeJSONToMap()
	//decodeJSONToStruct()
	//encodeStructToJSON()
	//parseURLString()
}
