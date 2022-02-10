package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// This is the basic struct for our code. All gets and transfers happen to it.
type player struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Position    string `json:"position"`
	TransferFee int    `json:"transferfee"`
}

// This function will act as a handler to the transfers argument
func transferHandler(transferCmd *flag.FlagSet, in *string, out *string) {
	fmt.Println("You are in the transfer player function")
}

// This function will act as a handler to the get argument
func getHandler(getCmd *flag.FlagSet, squad *bool, budget *bool) {
	getCmd.Parse(os.Args[2:]) //parse the flags
	//can't pass in both flags squad and budget
	if *squad == true && *budget == true {
		fmt.Errorf("Only one of --squad or --budget can be used as a flag")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	//can't leave passing in both flags either
	if *squad == false && *budget == false {
		fmt.Errorf("The command get must be followed by a --squad or --budget")
		getCmd.PrintDefaults()
		os.Exit(1)
	}
	//we will try to get the current squad from the current_squad json.
	if *squad == true {
		var current_squad []player
		fileBytes, err := ioutil.ReadFile("./current_squad.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(fileBytes, &current_squad)
		if err != nil {
			panic(err)
		}
		fmt.Printf("AGE \t NAME \tPOSITION \t VALUE\n")
		for _, player := range current_squad {
			fmt.Printf("%v\t%v\t%v\t%v\n", player.Age, player.Name, player.Position, player.TransferFee)
		}
	}
}

func main() {
	fmt.Println("Hello Man Utd fans!")
	// Get Argument Definition - Can help get the squad or get the current budget
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getSquad := getCmd.Bool("squad", false, "Get the current squad of Manchester United")
	getBudget := getCmd.Bool("budget", false, "Get the current budget of Manchester United")
	//getExpand := getCmd.Bool("expand", false,"Do you want to expand the current squad displayed")

	transferCmd := flag.NewFlagSet("transfer", flag.ExitOnError)
	transferIn := transferCmd.String("in", "", "Sign a player for Man Utd")
	transferOut := transferCmd.String("out", "", "Sell a Man Utd player")

	// if *in == "" && *out == "" {
	// 	fmt.Errorf("Please specify a player to buy or sell using the --in or --out flag")
	// 	getCmd.PrintDefaults()
	// 	os.Exit(1)
	// }

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'transfer' commands")
	}

	switch os.Args[1] {
	case "get":
		{
			getHandler(getCmd, getSquad, getBudget)
		}

	case "transfer":
		{
			transferHandler(transferCmd, transferIn, transferOut)
		}
	}
}
