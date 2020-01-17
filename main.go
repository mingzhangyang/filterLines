package main

import (
	"os"
	"fmt"
	"errors"
)

var options = map[string]string{
	"field": "1",
	"delimiter": "\t",
	"reverse": "false",
}

func main() {
	switch (len(os.Args)) {
	case 1:
		os.Stderr.WriteString("not enough arguments to run the command")
		os.Exit(1)
	case 2:
		if os.Args[1] != "-h" && os.Args[1] != "--help" {
			os.Stderr.WriteString("bad arguments")
			os.Exit(1)
		} 
		printHelpInfo()
		os.Exit(0)
	case 3:
		filter(os.Args[1], os.Args[2], options)
	case 4:
		if os.Args[1] != "-r" {
			os.Stderr.WriteString("bad argument")
			os.Exit(1)
		}
		options["reverse"] = "true"
		filter(os.Args[2], os.Args[3], options)
	case 5:
		err := setArgument(options, [2]string{os.Args[1], os.Args[2]})
		if err != nil {
			os.Stderr.WriteString("bad argument")
			os.Exit(1)
		}
		filter(os.Args[3], os.Args[4], options)
	case 6:
		if os.Args[1] != "-r" && os.Args[3] != "-r" {
			os.Stderr.WriteString("bad argument")
			os.Exit(1)
		}
		if os.Args[1] == "-r" {
			options["reverse"] = "true"
			err := setArgument(options, [2]string{os.Args[2], os.Args[3]})
			if err != nil {
				os.Stderr.WriteString("bad argument")
				os.Exit(1)
			}
			filter(os.Args[4], os.Args[5], options)
		}
		if os.Args[3] == "-r" {
			options["reverse"] = "true"
			err := setArgument(options, [2]string{os.Args[1], os.Args[2]})
			if err != nil {
				os.Stderr.WriteString("bad argument")
				os.Exit(1)
			}
			filter(os.Args[4], os.Args[5], options)
		}
	case 7:
		err := setArgument(options, [2]string{os.Args[1], os.Args[2]})
		if err != nil {
			os.Stderr.WriteString("bad argument")
			os.Exit(1)
		}
		err = setArgument(options, [2]string{os.Args[3], os.Args[4]})
		if err != nil {
			os.Stderr.WriteString("bad argument")
			os.Exit(1)
		}
		filter(os.Args[5], os.Args[6], options)
	case 8:
		if os.Args[1] != "-r" && os.Args[3] != "-r" && os.Args[5] != "-r" {
			os.Stderr.WriteString("bad argument")
			os.Exit(1)
		}
		options["reverse"] = "true"
		if os.Args[1] == "-r" {
			err := setArgument(options, [2]string{os.Args[2], os.Args[3]})
			if err != nil {
				os.Stderr.WriteString("bad argument")
				os.Exit(1)
			}
			err = setArgument(options, [2]string{os.Args[4], os.Args[5]})
			if err != nil {
				os.Stderr.WriteString("bad argument")
				os.Exit(1)
			}
			filter(os.Args[6], os.Args[7], options)
		}
		if os.Args[3] == "-r" {
			err := setArgument(options, [2]string{os.Args[1], os.Args[2]})
			if err != nil {
				os.Stderr.WriteString("bad argument")
				os.Exit(1)
			}
			err = setArgument(options, [2]string{os.Args[4], os.Args[5]})
			if err != nil {
				os.Stderr.WriteString("bad argument")
				os.Exit(1)
			}
			filter(os.Args[6], os.Args[7], options)
		}
		if os.Args[5] == "-r" {
			err := setArgument(options, [2]string{os.Args[1], os.Args[2]})
			if err != nil {
				os.Stderr.WriteString("bad argument")
				os.Exit(1)
			}
			err = setArgument(options, [2]string{os.Args[3], os.Args[4]})
			if err != nil {
				os.Stderr.WriteString("bad argument")
				os.Exit(1)
			}
			filter(os.Args[6], os.Args[7], options)
		}
	default:
		os.Stderr.WriteString("bad argument")
		os.Exit(1)
	}

	os.Exit(0)
}

func setArgument(opts map[string]string, v [2]string) error {
	switch v[0] {
	case "-r":
		if v[1][0] != '-' {
			return errors.New("bad arguments: " + v[0] + " " + v[1])
		}
		opts["reverse"] = "true"
	case "-f":
		opts["field"] = v[1]
	case "-d":
		opts["delimiter"] = v[1]
	default:
		return errors.New("bad arguments: " + v[0] + " " + v[1])
	}
	return nil
}

func printHelpInfo() {
	fmt.Println("A utility to filter file2 with the patterns (row as pattern) in file1. Similar to \"grep -f file1 file2\"")
	fmt.Println("\nUsage:")
	fmt.Println("\tfilterLine -h\t\tprint help info")
	fmt.Println("\tfilterLine --help\t\tprint help info")
	fmt.Println("\tfilterLine file1 file2\t\tprint rows in file2 that contain a pattern in file1")
	fmt.Println("\tfilterLine -[options] file1 file2\t\tsee below for options")
	fmt.Println("\nOptions:")
	fmt.Println("\t-r: keep the rows in file2 that are not found in file1")
	fmt.Println("\t-d: specify the delimiter in file2, default: \"\\t\"")
	fmt.Println("\t-n: specify the field in rows in file2 to compare with patterns in file1, starts from 1, defalut: 1 ")
}