package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("going to make the GET request")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("STATUS CODE :", response.StatusCode)
	fmt.Println("CONTENT LENGTH :", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostJSONRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	requestBody := strings.NewReader(`
		{
			"course": "golang",
			"price" : 100,
			"platform": "learn code online"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}
