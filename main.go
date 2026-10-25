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
			fmt.Println("Please provide a path to organize.");
		    return
		}
		
		indexOfEqual := strings.Index(string(args[0]), "=")
		pathProvided  := args[0][indexOfEqual+1:]

		if strings.TrimSpace(pathProvided) == "" {
			fmt.Println("Please provide a path to organize.");
			return
		}
	}

	// Process the arguments provided by the user
	flag.Parse();

	file, err := os.Open(*organizer)
	if err != nil {
		fmt.Println("An error occurrend with the provided path:", err)
		return
	}
	defer file.Close()

	//Get information of the file.
	fileInfo, err := file.Stat();
	if err != nil {
		fmt.Println("An error occurred while getting the file information:", err);
		return
	}

	//check whether the path points to a direcotry or a file
	if !fileInfo.IsDir() {
		fmt.Println("The provided path is not a directory");
		return
	}

	directorys, err := os.ReadDir(*organizer);
	if err != nil {
		fmt.Println("An error occurred while reading the directory:", err)
	}

	for index, entry := range directorys {
		if !entry.IsDir() {
			fmt.Println(index,":",entry);
		}
	}

}
