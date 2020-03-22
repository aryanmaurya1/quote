package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var user = os.Getenv("USER")
	var filename = "/home/" + user + "/quote.txt"
	var quotes, err = os.Open(filename)
	if err != nil {
		fmt.Println("Error in Quotes file !!!")
		os.Exit(1)
	}
	defer quotes.Close()

	var listOfQuotes []string

	var scanner = bufio.NewScanner(quotes)
	for scanner.Scan() {
		line := scanner.Text()
		listOfQuotes = append(listOfQuotes, string(line))
	}

	rand.Seed(int64(time.Now().UnixNano()))
	var finalQuote = listOfQuotes[rand.Int()%len(listOfQuotes)]
	fmt.Println(finalQuote)

}
