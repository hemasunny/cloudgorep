package main

import (
	"fmt"
	"net/http"
)


func   Homehandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Home")

}

func   Signuphandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Signup Page")

}

func main() {


    http.HandleFunc("/home",Homehandler)

	http.HandleFunc("/Signup",Signuphandler)

	http.ListenAndServe("9000",nil)
}