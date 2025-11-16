package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParsePhaseOne(file *os.File) []string {

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	code := make([]string, 0)
	romAddr := 0

	for scanner.Scan() {
		line := scanner.Text()

		// strip trailing comment
		if idx := strings.Index(line, "//"); idx != -1 {
			line = line[:idx]
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// LABEL
		if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
			label := line[1 : len(line)-1]
			symbolTable[label] = romAddr
			continue
		}

		// real instruction
		code = append(code, line)
		romAddr++
	}

	return code
}

func ParsePhaseTwo(hackAsmCode []string) []string {
	var binaryInstructions []string = make([]string, 0)
	nextVarAddr := 16

	for _, asmCommand := range hackAsmCode {

		// A-instructions
		if strings.HasPrefix(asmCommand, "@") {
			sym := asmCommand[1:]
			binaryInstruction := "0"
			intAddr := 0

			// numeric literal @123
			if n, err := strconv.Atoi(sym); err == nil {
				intAddr = n

			} else {
				// known symbol (label or predefined)
				if v, exists := symbolTable[sym]; exists {
					intAddr = v

				} else {
					// new variable
					symbolTable[sym] = nextVarAddr
					intAddr = nextVarAddr
					nextVarAddr++
				}
			}

			binaryInstruction += ToBinary15Bit(intAddr)
			binaryInstructions = append(binaryInstructions, binaryInstruction)
			continue
		}

		// C instruction
		eq_idx := strings.Index(asmCommand, "=")
		semicolon_idx := strings.Index(asmCommand, ";")
		var comp string = ""
		var dest string = ""
		var jump string = ""

		// dest
		if eq_idx != -1 {
			dest = asmCommand[:eq_idx]
		}

		// comp and jump
		if semicolon_idx != -1 {
			// comp before ';'
			if eq_idx != -1 {
				comp = asmCommand[eq_idx+1 : semicolon_idx]
			} else {
				comp = asmCommand[:semicolon_idx]
			}

			jump = asmCommand[semicolon_idx+1:]

		} else {
			// no jump
			if eq_idx != -1 {
				comp = asmCommand[eq_idx+1:]
			} else {
				comp = asmCommand
			}
		}

		binaryInstruction := TranslateCInstruction(dest, comp, jump)
		binaryInstructions = append(binaryInstructions, binaryInstruction)
	}

	return binaryInstructions
}
