package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"errors"
)
// constants
// rounter information is formated as A.B.C.D=E
// user input information is formated as A.B.C.D
const inputLength int = 4 
const addressSeparator string = "."
const serverSeparator string = "="

// using Trie structure to load and search
var trie *Trie
var configFilePath = "config.txt" // the default configFile path

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		configFilePath = args[0]
	}
	loadConfig()
}

func loadConfig() {
	err, configs := readConfig(configFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	trie = buildTire(configs)
	fmt.Println("configuration loaded")
	// launch service
	runService()
}

func findRoute(input string) string {
	// check if input is valid
	valid, decoded := checkInput(input)
	if valid {
		return trie.search(decoded)
	}
	return "pleaes check input format"
}

func runService() {
	fmt.Println("......service running......type quit() to exit")
	for {
		fmt.Print("enter: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "quit()\n" {
			fmt.Println("............service shut down............")
			return
		}
		fmt.Println(">> ", findRoute(text))
	}
}

// readConfig: read Configuration File
func readConfig(filePath string) (error, [][]string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err) //terminate the program
	}
	defer file.Close()
	var configs [][]string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		valid, decoded := checkRule(line)
		if valid {
			configs = append(configs, decoded)
		}
	}
	if len(configs)== 0{
		return errors.New("fail to load configuration"), configs
	}
	return nil, configs
}

func checkRule(input string) (bool, []string) {
	input = removeSpace(input)
	splited := strings.Split(input, addressSeparator)
	if len(splited) != inputLength {
		return false, nil
	}
	splittedServer := strings.Split(splited[inputLength-1], serverSeparator)
	if len(splittedServer) != 2 {
		return false, nil
	}
	res := append(splited[0:inputLength-1], splittedServer[0])
	res = append(res, splittedServer[1])
	return true, res
	
}

func checkInput(input string) (bool, []string) {
	input = removeSpace(input)
	splited := strings.Split(input, addressSeparator)
	if len(splited) != inputLength {
		return false, nil
	}

	return true, splited
}

func removeSpace(input string) string{
	return strings.Join(strings.Fields(input),"")
}
