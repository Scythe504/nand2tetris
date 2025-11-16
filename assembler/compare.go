package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

/*
This is simply to check if the assembler implementation is correct or not
*/
func compareHackOutputs(hackDir, cmpDir string) {
	fmt.Println("Running tests...")
	files, err := os.ReadDir(cmpDir)
	checkErr(err)

	allGood := true

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".hack" {
			continue
		}

		cmpPath := filepath.Join(cmpDir, f.Name())
		hackPath := filepath.Join(hackDir, f.Name())

		if _, err := os.Stat(hackPath); err != nil {
			fmt.Printf("❗ Missing: %s\n", hackPath)
			allGood = false
			continue
		}

		ok, err := filesEqual(cmpPath, hackPath)
		checkErr(err)

		if ok {
			fmt.Printf("MATCH: %s\n", f.Name())
		} else {
			fmt.Printf("DIFF : %s\n", f.Name())
			allGood = false
		}
	}

	if allGood {
		fmt.Println("ALL TESTS PASSED!")
	} else {
		fmt.Println("SOME TESTS FAILED.")
	}
}

// Compares two files **line by line**
func filesEqual(a, b string) (bool, error) {
	f1, err := os.Open(a)
	if err != nil {
		return false, err
	}
	defer f1.Close()

	f2, err := os.Open(b)
	if err != nil {
		return false, err
	}
	defer f2.Close()

	s1 := bufio.NewScanner(f1)
	s2 := bufio.NewScanner(f2)

	lineNum := 1

	for {
		line1 := s1.Scan()
		line2 := s2.Scan()

		if !line1 || !line2 {
			// both ended → OK
			if !line1 && !line2 {
				return true, nil
			}
			// one ended early
			return false, nil
		}

		if s1.Text() != s2.Text() {
			fmt.Printf("    Line %d differs:\n", lineNum)
			fmt.Printf("       cmp : %s\n", s1.Text())
			fmt.Printf("       hack: %s\n", s2.Text())
			return false, nil
		}

		lineNum++
	}
}
