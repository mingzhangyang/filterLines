package main 

import (
	"os"
	"errors"
)

func collectArguments() error {
	arr := os.Args[1:]
	n := len(arr)
	switch {
	case n == 0:
		return errors.New("no arguments provided, run with -h for help")
	case n == 1:
		if arr[0] == "-h" || arr[0] == "--help" {
			printHelpInfo()
			os.Exit(0)
		}
		return errors.New("bad argument")
	case n > 7:
		return errors.New("too many arguments")
	default:
		i := 0
		for i < n {
			switch arr[i] {
			case "-r":
				options["-r"] = "on";
				i++
			case "-f", "-d":
				if i + 1 == n {
					return errors.New("bad argument: " + arr[i])
				}
				if arr[i+1][0] == '-' {
					return errors.New("bad argument: " + arr[i])
				}
				options[arr[i]] = arr[i+1]
				i += 2
			default:
				targets = append(targets, arr[i])
				i++
			}
		}
	}
	if len(targets) != 2 {
		return errors.New("bad arguments for input files")
	}
	return nil
}
