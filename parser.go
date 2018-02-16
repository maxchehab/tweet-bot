// This program simply parses the text of tweets.csv and outputs to tweet-bot.txt

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("./tweets.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	output, err := os.Create("tweet-bot.txt")
	if err != nil {
		log.Fatal("Cannot create output", err)
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	fmt.Println("parsing.")
	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), ",")
		if len(columns) > 7 {
			text := strings.Split(columns[7], "\"")
			if len(text) > 1 && len(text[1]) > 1 {
				fmt.Fprintf(output, text[1]+"\n")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("done.")
}
