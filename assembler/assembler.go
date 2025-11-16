package main

import (
	"bufio"
	"os"
)


func AssembleBinaryInstructions(binaryInstructions []string, outPath string) bool {
	file, err := os.Create(outPath)
	checkErr(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, instruction := range binaryInstructions {
		_, err := writer.WriteString(instruction + "\n")
		checkErr(err)
	}

	err = writer.Flush()
	checkErr(err)

	return true
}
