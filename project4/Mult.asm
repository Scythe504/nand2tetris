// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// The algorithm is based on repetitive addition.

//// Replace this comment with your code.

// a
@R0
D=M

@sumAB
M=D
// b
@R1
D=M

@sumAB
D=D+M

@BOTH_0
D;JEQ

@i
M=1

@res
M=0

(LOOP)    
    @i
    D=M
    @R1
    D=M-D
    @END
    D;JLT // if i - R1 > 0 goto end

    // Mult with addition
    @R0
    D=M
    @res
    M=D+M // res = r1 + res

    @i
    M=M+1
    @LOOP
    0;JMP
(BOTH_0)
    @res
    D=M-1
    @R2
    M=D

    @END
    0;JMP
(END)
    @res
    D=M
    @R2
    M=D
    @END
    0;JMP
