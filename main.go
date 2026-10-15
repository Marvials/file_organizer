package main

import "fmt"
import "flag"
import "os"

func main(){
	fmt.Println("Start project");

	defaultPath , err:= os.Getwd();
	if err != nil {
		fmt.Println("An error occurred while getting the path of the current dicrectory");
		return;
	}

	organizer := flag.String("organizer", defaultPath, "Organize files in directories. Default value is the current path");

	flag.Parse();

	fmt.Println("Organizer:", *organizer);
}
