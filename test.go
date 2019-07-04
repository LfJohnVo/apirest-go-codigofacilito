package main

import (
	"bytes"
	"fmt"
//"io/ioutil"
"net/http"
)

func main() {
	url := "http://localhost:8080/user/delete/1" //Esta es la ruta
	var jsonStr= []byte(`{"username": "PATCH1", "first_name": "PATCH2", "last_name": "PATCH3"}`)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))

	fmt.Println("URL:>", url)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}