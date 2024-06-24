package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var members []Member

var isRunning chan int = make(chan int)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage) // register homepage endpoint with function
	myRouter.HandleFunc("/member", createNewMember).Methods("POST")
	myRouter.HandleFunc("/members", allMembers).Methods("GET") // register members endpoint with function
	myRouter.HandleFunc("/member/{id}", returnSingleMember).Methods("GET")
	myRouter.HandleFunc("/member/{id}", deleteMember).Methods("DELETE")
	myRouter.HandleFunc("/member/{id}", editMember).Methods("PUT")
	myRouter.HandleFunc("/members/{id}", deduplicateMembers).Methods("PUT")
	myRouter.HandleFunc("/member", secretAction).Methods("SECRET")
	log.Fatal(http.ListenAndServe(":8081", myRouter)) // create listener from http object (server)
	isRunning <- 1
}

func main() {
	members = Members{
		Member{Id: "1", Name: "Bob", Age: "19", MembershipFee: "19.99"},
		Member{Id: "2", Name: "Joe", Age: "16", MembershipFee: "14.99"},
	}
	go handleRequests() // execute the server instance
	go menu()           // runs client request menu
	<-isRunning
}
