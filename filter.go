package main

import (
	"log"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func filter(fp1, fp2 string, opts map[string]string) {
	fb, err := os.Open(fp1)
	if err != nil {
		log.Fatal(err)
	}

	fbScanner := bufio.NewScanner(fb)
	m := make(map[string]bool)

	for fbScanner.Scan() {
		m[fbScanner.Text()] = true
	}

	if err = fbScanner.Err(); err != nil {
		log.Fatal(err)
	}

	if err = fb.Close(); err != nil {
		log.Fatal(err)
	}

	fa, err := os.Open(fp2)
	if err != nil {
		log.Fatal(err)
	}
	defer fa.Close()

	outBuf := bufio.NewWriter(os.Stdout)
	faScanner := bufio.NewScanner(fa)

	d := opts["delimiter"]
	n, err := strconv.Atoi(opts["field"])
	if err != nil {
		log.Fatal("invalid argument for -f")
	}
	n = n - 1
	
	if opts["reverse"] == "false" {
		for faScanner.Scan() {
			if m[strings.Split(faScanner.Text(), d)[n]] {
				_, _ = outBuf.Write(faScanner.Bytes())
				_, _ = outBuf.WriteString("\n")
			}
			continue
		}
	} else {
		for faScanner.Scan() {
			if m[strings.Split(faScanner.Text(), d)[n]] {
				continue
			}
			_, _ = outBuf.Write(faScanner.Bytes())
			_, _ = outBuf.WriteString("\n")
		}
	}

	if err = faScanner.Err(); err != nil {
		log.Fatal(err)
	}

	if err = outBuf.Flush(); err != nil {
		log.Fatal(err)
	}
}

