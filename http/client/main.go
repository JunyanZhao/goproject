package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wesure.com/modifysrc/json"
)
type contact struct {
	Name  string
	Addr  string
	Phone string
}

func main() {
	response, _ := http.Get("http://localhost:8080/list")
	defer response.Body.Close()
	fmt.Printf("response:%+v\n", response)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(body))
	var c []contact
	json.Unmarshal(body, &c)

	for k, v := range c {
		fmt.Printf("%v: %+v\n", k, v)
	}
	fmt.Println("len:",len(c))
}
