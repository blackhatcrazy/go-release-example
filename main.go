package main

import (
	"flag"
	"fmt"
	"os"
)

var motd string

func PrintfFix(message string) {
	fmt.Printf("fix: %s\n", message)
}

func PrintfFeature(message string) {
	fmt.Printf("feat: %s\n", message)
}

func PrintfBreakingChange(ty string, message string) {
	fmt.Printf("%s!: %s\n", ty, message)
}

func main() {
	flag.Parse()
	if motd == "" {
		fmt.Println("Error: Message of the day cannot be empty.")
		os.Exit(1)
	}
	fmt.Printf("%s\n", motd)
	PrintfFeature("This is some more text!")
	PrintfFeature("This is the third line of text")
	PrintfFix("Fixed typo in third line of text")
	PrintfBreakingChange("feat", "Add required message of the day")
}

func init() {
	flag.StringVar(&motd, "motd", "", "Message of the day to print at startup")
}
