package main

import (
	"fmt"
	"encoding/json"
	"flag"
)

type player struct {
	Name string `json:"name"`
	Age int  `json:"age"`
	Position string `json:"position"`
	TransferFee int `json:"transferfee"`
}


func buyPlayer() {

}

func sellPlayer() {

}


func getCurrentSquad() {

}

func main() {
	fmt.Println("Hello Man Utd fans")
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "Get all Man United Players")
	getName:= getCmd.String("name", "", "Get all Man United Players")


	buyCmd := flag.NewFlagSet("add", flag.ExitOnError)
	buyName := addCmd.String("name", "", "Get all Man United Players")
	buyAge := addCmd.Int("age", 0, "Age of the United Player")

	sellCmd := flag.NewFlagSet("add", flag.ExitOnError)
	sellName := addCmd.String("name", "", "Get all Man United Players")
	sellAge := addCmd.Int("age", 0, "Age of the United Player")

}