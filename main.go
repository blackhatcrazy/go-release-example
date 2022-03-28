package main

import "fmt"

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
	fmt.Printf("Hello world!")
	PrintfFeature("This is some more text!")
}
