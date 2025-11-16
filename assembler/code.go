package main

import "fmt"

// dest bits
var destMap = map[string]string{
	"DEST_NULL": "000", // The value is not stored
	"M":         "001", // RAM[A]
	"D":         "010", // D Register
	"MD":        "011", // RAM[A] and D register
	"A":         "100", // A register
	"AM":        "101", // A register and RAM[A]
	"AD":        "110", // A register and D register
	"AMD":       "111", // A register, RAM[A] and D register
}

// jump bits
var jumpMap = map[string]string{
	"JUMP_NULL": "000", // no jump
	"JGT":       "001", // if out > 0 jump
	"JEQ":       "010", // if out == 0 jump
	"JGE":       "011", // if out >= 0 jump
	"JLT":       "100", // if out < 0 jump
	"JNE":       "101", // if out != 0 jump
	"JLE":       "110", // if out <= 0 jump
	"JMP":       "111", // Unconditional jump
}

var compMap = map[string]string{
	// bits a|c1|c2|c3|c4|c5|c6
	"0":   "0101010", // a = 0
	"1":   "0111111", // a = 0
	"-1":  "0111010", // a = 0
	"D":   "0001100", // a = 0
	"A":   "0110000", // a = 0
	"M":   "1110000", // a = 1
	"!D":  "0001101", // a = 0
	"!A":  "0110001", // a = 0
	"!M":  "1110001", // a = 1
	"-D":  "0001111", // a = 0
	"-A":  "0110011", // a = 0
	"-M":  "1110011", // a = 1
	"D+1": "0011111", // a = 0
	"A+1": "0110111", // a = 0
	"M+1": "1110111", // a = 1
	"D-1": "0001110", // a = 0
	"A-1": "0110010", // a = 0
	"M-1": "1110010", // a = 1
	"D+A": "0000010", // a = 0
	"D+M": "1000010", // a = 1
	"D-A": "0010011", // a = 0
	"D-M": "1010011", // a = 1
	"A-D": "0000111", // a = 0
	"M-D": "1000111", // a = 1
	"D&A": "0000000", // a = 0
	"D&M": "1000000", // a = 1
	"D|A": "0010101", // a = 0
	"D|M": "1010101", // a = 1
}

// comp bits C instruction[6..12]
func TranslateComp(comp string) string {
	return compMap[comp]
}

// jump bits C instruction [3..5]
func TranslateJump(jump string) string {
	return jumpMap[jump]
}

// jump bits C instruciton [0..2]
func TranslateDest(dest string) string {
	return destMap[dest]
}

func ToBinary15Bit(addr int) string {
	return fmt.Sprintf("%015b", addr)
}

func TranslateCInstruction(dest, comp, jump string) string {
	binaryInstruction := "111"
	
	binaryInstruction += TranslateComp(comp)

	if dest == "" {
		binaryInstruction += TranslateDest("DEST_NULL")
	} else {
		binaryInstruction += TranslateDest(dest)
	}

	if jump == "" {
		binaryInstruction += TranslateJump("JUMP_NULL")
	} else {
		binaryInstruction += TranslateJump(jump)
	}

	return binaryInstruction
}