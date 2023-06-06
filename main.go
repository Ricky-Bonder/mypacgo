package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var maze []string

func loadMaze(file string) error {

	//open file
	f, err := os.Open(file) //

	if err != nil {
		log.Println("error opening maze txt file: ", err)
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

func printScreen() {
	for _, line := range maze {
		fmt.Println(line)
	}
}

func main() {
	// initialize game

	err := loadMaze("/home/rossola/VSCodeProjects/pacgo/step01/maze01.txt")

	if err != nil {
		log.Println("failed to load maze:", err)
		return
	}

	// load resources

	// game loop
	for {
		// update screen
		printScreen()

		// process input

		// process movement

		// process collisions

		// check game over

		// Temp: break infinite loop
		break
		// repeat
	}

}
