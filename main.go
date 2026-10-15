package main

import "fmt"
import "flag"
import "os"

func main(){
	fmt.Println("Start project");

	// Defines the default path with the current directory
	defaultPath , err:= os.Getwd();
	if err != nil {
		fmt.Println("An error occurred while getting the path of the current directory");
		return;
	}

	// Create the organizer flag
	organizer := flag.String("organizer", defaultPath, "Organize files in directories. Default value is the current path");

	flag.Parse();

	// Check if the path provided exists
	_, err = os.Stat(*organizer);
	if err != nil {
		fmt.Println("The path provided does not exist");
		return;
	}

	fmt.Println("Organizer:", *organizer);
}
