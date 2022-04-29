package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Green = "\u001b[32m"
	Red   = "\u001b[31m"
	Blue  = "\u001b[34m"
	Reset = "\u001b[0m"
)

func intro() {
	fmt.Println(Green, "Welcome...")
	fmt.Println("Welcome to Curs>d")
}

func level1() {
	fmt.Println(Green, "Level 1")
	fmt.Println("You are in a dark room.")
	fmt.Println("There is a door to your right and left.")
	fmt.Println("Which one do you take?")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "left" {
			fmt.Println(Red, "You died.", Reset)
			os.Exit(0)
		} else if text == "right" {
			fmt.Println(Blue, "You got a key!", Reset)
			level2()
		} else {
			fmt.Println(Red, "Please enter a valid direction, left or right.", Reset)
		}
	}
}

func level2() {
	fmt.Println(Green, "Level 2")
	fmt.Println("You have entered a room with a dim light, there is a shadow that isn't yours and one door.")
	fmt.Println("Go back to the previous room or go through the door?")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "back" {
			fmt.Println(Red, "You died.", Reset)
			os.Exit(0)
		} else if text == "door" {
			fmt.Println(Green, "You used the key to open the door.", Reset)
			fmt.Println(Blue, "You lived but the shadow is still creeping behind you!", Reset)
			level3()
		} else {
			fmt.Println(Red, "Please enter a valid direction, back or door.", Reset)
		}
	}
}

func level3() {
	fmt.Println(Green, "Level 3")
	fmt.Println("The shadow has caught up to you.")
	fmt.Println("Turn around and fight the shadow or run away?")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "fight" {
			fmt.Println(Blue, "You escaped, but are now bleeding heavily.", Reset)
			level4()
		} else if text == "run" {
			fmt.Println(Red, "You Died.", Reset)
			os.Exit(0)
		} else {
			fmt.Println(Red, "Please enter a valid direction, fight or run.", Reset)
		}
	}
}

func level4() {
	fmt.Println(Green, "Level 4")
	fmt.Println("You have reached the boss level. Find the key to the door.")
	fmt.Println("You have to find the key to the door.")
	for {
		if yesOrNo("You already used your key, do you search search for another?") {
			fmt.Println(Green, "You search and search but cannot find another key.", Reset)
		} else {
			fmt.Println(Red, "You Died!", Reset)
			os.Exit(0)
		}
		if yesOrNo("You are tired from searching... Do you use the last of your strength to bust down the door?") {
			winner()
			os.Exit(0)
		} else {
			fmt.Println(Red, "You Died!", Reset)
			os.Exit(0)
		}
	}
}

func yesOrNo(s string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/n]: ", s)
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func winner() {
	fmt.Print(Red, "Y", Reset)
	fmt.Print(Green, "o", Reset)
	fmt.Print(Blue, "u", Reset)
	fmt.Print(" ")
	fmt.Print(Red, "W", Reset)
	fmt.Print(Green, "i", Reset)
	fmt.Print(Blue, "n", Reset)
	fmt.Print(" ")
}

func main() {
	intro()
	confirmation := yesOrNo("Are you Worthy?")
	if confirmation {
		fmt.Println(Red, "Welcome in!", Reset)
		level1()
	}
}
