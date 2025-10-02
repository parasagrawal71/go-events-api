package apiclient

import (
	"fmt"
	"time"
)

func Run() {
	client := NewClient(10*time.Second, map[string]string{
		"Authorization": "Bearer YOUR_TOKEN",
	})

	// --- GET Example ---
	var getResult map[string]interface{}
	err := client.Get("https://jsonplaceholder.typicode.com/posts/1", &getResult)
	if err != nil {
		fmt.Println("GET error: ", err)
		return
	}

	// --- POST Example ---
	payload := map[string]interface{}{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	}
	var postResult map[string]interface{}
	err = client.Post("https://jsonplaceholder.typicode.com/posts", payload, &postResult)
	if err != nil {
		fmt.Println("POST error: ", err)
		return
	}
}
