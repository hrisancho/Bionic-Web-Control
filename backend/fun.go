package main

import "fmt"

type User struct {
	Age    int    `json:"age"`
	Name   string `json:"name"`
	UserID string `json:"userid"`
}

func main() {
	fmt.Println(hi())
}
func hi() (st string) {
	st = "sadf"
	return st
}
