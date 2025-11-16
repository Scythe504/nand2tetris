package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	outDir := flag.String("o", "", "Output directory for .hack files")
	testDir := flag.String("test", "", "Compare hack output with reference directory")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: hackasm [options] FILES|DIRS")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// if --test mode
	if *testDir != "" {
		compareHackOutputs("hack", *testDir)
		return
	}

	for _, p := range args {
		info, err := os.Stat(p)
		checkErr(err)

		if info.IsDir() {
			assembleDirectory(p, *outDir)
		} else {
			assembleFile(p, *outDir)
		}
	}

	fmt.Println("Done.")
}
func assembleDirectory(dir string, customOut string) {
	entries, err := os.ReadDir(dir)
	checkErr(err)

	for _, e := range entries {
		if filepath.Ext(e.Name()) != ".asm" {
			continue
		}
		asmPath := filepath.Join(dir, e.Name())
		assembleFile(asmPath, customOut)
	}
}

func assembleFile(asmPath string, customOut string) {
	fmt.Println("Assembling:", asmPath)

	resetSymbolTable()

	file, err := os.Open(asmPath)
	checkErr(err)

	code := ParsePhaseOne(file)
	file.Close()

	binary := ParsePhaseTwo(code)

	// Determine output path
	var outPath string
	base := filepath.Base(asmPath)
	hackName := base[:len(base)-4] + ".hack"

	if customOut != "" {
		checkErr(os.MkdirAll(customOut, 0755))
		outPath = filepath.Join(customOut, hackName)
	} else {
		outPath = asmPath[:len(asmPath)-4] + ".hack"
	}

	AssembleBinaryInstructions(binary, outPath)

	fmt.Println(" → Generated:", outPath)
}
