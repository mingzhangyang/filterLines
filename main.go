package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		log.Fatal("no arguments found")
	case 2:
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			printHelpInfo()
			return
		}
		log.Fatal("bad argument")
	default:
		err := collectArguments()
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(targets, options)
	}
	
	filter(targets[0], targets[1], options)
}



func printHelpInfo() {
	fmt.Println("A utility to filter file2 with the patterns (row as pattern) in file1. Similar to \"grep -f file1 file2\"")
	fmt.Println("\nUsage:")
	fmt.Println("\tfilterLine -h\t\t\t\t\tprint help info")
	fmt.Println("\tfilterLine --help\t\t\t\tprint help info")
	fmt.Println("\tfilterLine file1 file2\t\t\t\tprint rows in file2 that contain a pattern in file1")
	fmt.Println("\tfilterLine [options] file1 file2\t\tsee below for options")
	fmt.Println("\nOptions:")
	fmt.Println("\t-r: keep the rows in file2 that are not found in file1")
	fmt.Println("\t-d: specify the delimiter in file2, default: \"\\t\"")
	fmt.Println("\t-f: specify the field in rows in file2 to compare with patterns in file1, starts from 1, defalut: 1 ")
}