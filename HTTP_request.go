package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	/* 5a
	Output:
	User Id: 1
	Id: 1
	Title: sunt aut facere repellat provident occaecati excepturi optio reprehenderit
	Body: quia et suscipit
	suscipit recusandae consequuntur expedita et cum
	reprehenderit molestiae ut ut quas totam
	nostrum rerum est autem sunt rem eveniet architecto
	*/

	
	url := "https://jsonplaceholder.typicode.com/posts/1"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP GET request failed: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body: ", err)
		return
	}
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("Error parsing the JSON response: ", err)
		return
	}
	fmt.Printf("User Id: %d\n", post.UserId)
	fmt.Printf("Id: %d\n", post.Id)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Body: %s\n", post.Body)
}
