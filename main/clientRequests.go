package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func makeGetRequest() {
	resp, err := http.Get("http://localhost:8081/members")
	body, _ := io.ReadAll(resp.Body)
	if err == nil {
		log.Println(string(body))
	} else {
		log.Fatalln(err)
	}
}

func makeUpdateRequest() {
	id := ""
	fmt.Print("Input an id to modify: ")
	fmt.Scanln(&id)
	opt := 0
	for opt != 5 {
		fmt.Println("Select an attribute to update:\n1)ID\n2)Name\n3)Age\n4)Membership Fee\n5)Return to Menu")
		fmt.Scanln(&opt)
		genericReq := func(attr string, id string) bool {
			input := ""
			fmt.Printf("Please input a new %s: ", attr)
			fmt.Scanln(&input)
			if verify(attr, input) {
				putBody := []byte(fmt.Sprintf("{\"%s\": \"%s\"}", attr, input))
				requestBody := bytes.NewBuffer(putBody)
				req, _ := http.NewRequest("PUT", "http://localhost:8081/member/"+id, requestBody)
				req.Header.Add("Content-Type", "application/json")
				resp, _ := http.DefaultClient.Do(req)
				body, _ := io.ReadAll(resp.Body)
				fmt.Println(string(body))
				return true
			}
			return false
		}
		isIDUpdate := false
		switch opt {
		case (1):
			isIDUpdate = genericReq("ID", id)
		case (2):
			genericReq("Name", id)
		case (3):
			genericReq("Age", id)
		case (4):
			genericReq("MembershipFee", id)
		case (5):
			continue
		default:
			fmt.Println("Invalid Option")
		}
		if isIDUpdate {
			break
		}
	}

}

func verify(attr string, val string) bool {
	fmt.Println("Are you sure that you would like to update the " + attr + " to " + val + "?\n(y/n):")
	opt := ""
	for {
		fmt.Scanln(&opt)
		switch opt {
		case ("y"):
			return true
		case ("n"):
			return false
		default:
			fmt.Println("Please input y/n")
		}
	}
}

func makePostRequest() {
	var newMember Member
	fmt.Println("Fill in the fields below to create a new member: ")
	fmt.Println("ID: ")
	fmt.Scanln(&newMember.Id)
	fmt.Println("Name: ")
	fmt.Scanln(&newMember.Name)
	fmt.Println("Age: ")
	fmt.Scanln(&newMember.Age)
	fmt.Println("Fee: ")
	fmt.Scanln(&newMember.MembershipFee)
	fmt.Printf("\nID: %s\nName: %s\nAge: %s\nMembership Fee: %s\nIS THIS CORRECT? (y/n):", newMember.Id, newMember.Name, newMember.Age, newMember.MembershipFee)
	opt := ""
	isValid := false
	for !isValid {
		fmt.Scanln(&opt)
		switch opt {
		case ("y"):

			jsonString, _ := json.Marshal(&newMember)
			requestBody := bytes.NewBuffer(jsonString)

			req, _ := http.NewRequest("POST", "http://localhost:8081/member", requestBody)
			req.Header.Add("Content-Type", "application/json")
			resp, _ := http.DefaultClient.Do(req)
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
			isValid = true
		case ("n"):
			isValid = true
		default:
			fmt.Println("Please input y/n")
		}
	}

}

func makeDeleteRequest() {
	id := ""
	opt := ""
	isValid := false
	fmt.Print("Please input the ID of the member to be deleted: ")
	fmt.Scan(&id)
	fmt.Printf("Are you sure you want to delete the members with id: %s?\n(y/n):", id)
	for !isValid {
		fmt.Scan(&opt)
		switch opt {
		case "y":
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8081/member/%s", id), nil)
			http.DefaultClient.Do(req)
			isValid = true
		case "n":
			isValid = true
		default:
			fmt.Println("Please input y/n")
		}
	}
}

func makeDeduplicateRequest() {
	id := ""
	opt := ""
	isValid := false
	fmt.Print("Please input the ID of the member to be deduplicated: ")
	fmt.Scan(&id)
	fmt.Printf("Are you sure you want to deduplicate the members with id: %s?\n(y/n):", id)
	for !isValid {
		fmt.Scan(&opt)
		switch opt {
		case "y":
			req, _ := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8081/members/%s", id), nil)
			http.DefaultClient.Do(req)
			isValid = true
		case "n":
			isValid = true
		default:
			fmt.Println("Please input y/n")
		}
	}
}

func makeSecretRequest() {
	req, _ := http.NewRequest("SECRET", "http://localhost:8081/member", nil)
	http.DefaultClient.Do(req)
}

func menu() {
	opt := 0
	for opt != 6 {
		fmt.Println("1)GET\n2)POST(Create)\n3)PUT(Update)\n4)DELETE\n5)Deduplicate ID\n6)EXIT")
		fmt.Scanln(&opt)
		switch opt {
		case (1):
			makeGetRequest()
		case (2):
			makePostRequest()
		case (3):
			makeUpdateRequest()
		case (4):
			makeDeleteRequest()
		case (5):
			makeDeduplicateRequest()
		case (6):
			isRunning <- 1
		case (43):
			makeSecretRequest()
		default:
			fmt.Println("Invalid Operation")
		}
	}
}
