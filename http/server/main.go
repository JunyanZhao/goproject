package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"wesure.com/modifysrc/json"
)

type MyHandler []contact

type contact struct {
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Phone string `json:"phone"`
}

//var handler = make(MyHandler, 0)
//var handler =&MyHandler{}

func (this *MyHandler) list(w http.ResponseWriter, req *http.Request) {
	fmt.Println(len(*this))
	for k, v := range *this {
		fmt.Fprintf(w, "%v: %+v\n", k, v)
	}
}

func (this *MyHandler) GetContact(w http.ResponseWriter, req *http.Request) {
	//fmt.Printf("req:%+v\n", req)
	//fmt.Println(len(*this))
	//c := (*this)[0]
	//fmt.Fprintf(w, "GetContact: %+v\n", c)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		if err != nil {
			fmt.Fprintf(w, "ReadAll error\n")
		}
		var c contact
		err = json.Unmarshal(body, c)
		if err != nil {
			fmt.Fprintf(w, "Unmarshal error\n")
		}
		for _,v := range *this {
			if v.Name == c.Name {
				fmt.Fprintf(w, "GetContact: %+v\n", v)
				return
			}
		}
	}
	fmt.Fprintf(w, "not found")
}

func (this *MyHandler) SetContact(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("req:%+v\n", req)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "ReadAll error\n")
	}
	var c = &contact{}
	fmt.Println("body:", string(body))
	err = json.Unmarshal(body, c)
	if err != nil {
		fmt.Fprintf(w, "Unmarshal error\n")
		return
	}
	//*this = append(*this, contact{"John", "China", "10000"})
	*this = append(*this, *c)
	fmt.Fprintf(w, "set new contact to %+v\n", this)
	return
}

func main() {
	handler := make(MyHandler, 0)
	handler = append(handler, contact{"jack", "China", "123455"})

	http.HandleFunc("/list", handler.list)
	http.HandleFunc("/GetContact", handler.GetContact)
	http.HandleFunc("/SetContact", handler.SetContact)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
