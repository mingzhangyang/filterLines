package main 

import (
	"os"
	"errors"
)

var targets []string = make([]string, 0, 16)

var options = map[string]string{
	"-f": "1",
	"-d": "\t",
	"-r": "off",
}

func collectArguments() error {
	arr := os.Args[1:]
	n := len(arr)
	switch {
	case n < 2:
		return errors.New("missing arguments")
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
