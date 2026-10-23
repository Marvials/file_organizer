package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
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
		fmt.Println();
		if !strings.ContainsAny( string(args[0]), "=") {
			fmt.Println("Please provide a path to organizer.");
		    return
		}
		
		indexOfEqual := strings.Index(string(args[0]), "=")
		pathProvided  := args[0][indexOfEqual+1:]

		if strings.TrimSpace(pathProvided) == "" {
			fmt.Println("Please provide a path to organizer.");
			return
		}
	}

	// Process the arguments provided by the user
	flag.Parse();

	file, err := os.Open(*organizer)
	if err != nil {
		fmt.Println("An error occurrend with the path provided:", err)
		return
	}
	defer file.Close()

	//Get information of the file.
	fileInfo, err := file.Stat();
	if err != nil {
		fmt.Println("An error occurred while getting the file stat.", err);
		return
	}

	//check whether the path points to a direcotry or a file
	if !fileInfo.IsDir() {
		fmt.Println("The path provided is not a directory");
		return
	}


	fmt.Println("Organizer:", *organizer);
}
