package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// function to execute on members endpoint
func allMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\n")
	for _, member := range members {
		json.NewEncoder(w).Encode(member)
	}
}

// returns a specific member
func returnSingleMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, member := range members {
		if member.Id == key {
			json.NewEncoder(w).Encode(member)
		}
	}
}

// function to create a new member
func createNewMember(w http.ResponseWriter, r *http.Request) {
	var m Member
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &m)
	members = append(members, m)
	fmt.Fprintf(w, "Creation Successful.\n")
	fmt.Fprintf(w, "%+v", string(body))
}

// edit a member of an associated id
func editMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	body, _ := io.ReadAll(r.Body)
	for i := 0; i < len(members); i++ {
		if members[i].Id == id {
			// overwrite the existing object fields
			json.Unmarshal(body, &members[i])
			s, _ := json.Marshal(members[i])
			fmt.Fprintf(w, "%+v", string(s))
		}
	}
	fmt.Fprintf(w, "\nEdit Successful.")

}

// delete a member of the associated id
func deleteMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	count := 0
	for i := 0; i < len(members); i++ {
		if members[i].Id == id {
			members = append(members[:i], members[i+1:]...)
			i--
			count += 1
		}
	}
	fmt.Printf("%d deletions occured.\n", count)
}

// testing other verbs/request types
func secretAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Secret Action")
}

// function to execute on homepage endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
}

func deduplicateMembers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	isDuplicate := false
	removed := 0
	for i := 0; i < len(members); i++ {
		if members[i].Id == id {
			if !isDuplicate {
				isDuplicate = true
			} else {
				members = append(members[:i], members[i+1:]...)
				removed++
				i--
			}
		}
	}
	fmt.Printf("%d duplicates removed\n", removed)
}
