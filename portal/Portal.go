package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello1(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8081/hello/portal")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, "Server1 say: %s \n ", string(body))
}

func hello2(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:8082/hello/portal")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, "Server2 say: %s \n ", string(body))
}

func main() {
	http.HandleFunc("/helloServer1", hello1)
	http.HandleFunc("/helloServer2", hello2)
	http.ListenAndServe(":8080", nil)
}
