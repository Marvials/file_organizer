package main

import (
	"fmt"
	"os"
	"flag"
)

func main(){
	// Defines the default path as the current directory
	defaultPath , err:= os.Getwd();
	if err != nil {
		fmt.Println("An error occurred while getting the path of the current directory");
		return;
	}

	// Create the organizer flag
	organizer := flag.String("organizer", defaultPath, "Organize files in the provided directory.");

	args := os.Args[1:]

	// check if a path is provided
	if len(args) == 1 {
		fmt.Println("Please provide a path to organizer.");
		return
	}

	// Process the arguments provided by the user
	flag.Parse();

	_, err = os.Stat(*organizer);
	if err != nil {
		fmt.Println("The provided path does not exist");
		return;
	}

	fmt.Println("Organizer:", *organizer);
}
