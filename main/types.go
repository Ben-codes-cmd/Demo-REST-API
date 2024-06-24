package main

// basic struct definition
type Member struct {
	Id            string `json:"Id"`
	Name          string `json:"Name"`
	Age           string `json:"Age"`
	MembershipFee string `json:"MembershipFee"`
}

type Members []Member
