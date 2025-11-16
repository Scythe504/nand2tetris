// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.

//// Replace this comment with your code.

// Fills the screen black when any key is pressed,
// clears the screen when no key is pressed.

(LOOP)
    // Read keyboard and reduce to 0/1 in variable 'cur'
    @KBD
    D=M
    @PRESS
    D;JNE          // if KBD != 0 jump to PRESS
    // NOT PRESS: cur = 0
    @cur
    M=0
    @CHECK_CHANGE
    0;JMP

(PRESS)
    // PRESS: set cur = 1
    @cur
    M=0
    @cur
    M=M+1

(CHECK_CHANGE)
    // if cur == prev -> no change -> continue looping
    @cur
    D=M
    @prev
    D=D-M
    @NO_CHANGE
    D;JEQ

    // update prev = cur
    @cur
    D=M
    @prev
    M=D

    // branch to draw or clear depending on cur
    @cur
    D=M
    @DRAW
    D;JNE          // if cur != 0 -> DRAW (black)
    @CLEAR
    0;JMP

(NO_CHANGE)
    @LOOP
    0;JMP

// ---------------- DRAW (black) ----------------
(DRAW)
    @SCREEN
    D=A
    @addr
    M=D            // addr = SCREEN base

    @8192
    D=A
    @count
    M=D            // count = 8192 words

(DRAW_LOOP)
    @count
    D=M
    @DONE_DRAW
    D;JEQ

    @addr
    A=M
    M=-1           // write black (all 1s)

    @addr
    M=M+1

    @count
    M=M-1

    @DRAW_LOOP
    0;JMP

(DONE_DRAW)
    @LOOP
    0;JMP

// ---------------- CLEAR (white) ----------------
(CLEAR)
    @SCREEN
    D=A
    @addr
    M=D            // addr = SCREEN base

    @8192
    D=A
    @count
    M=D            // count = 8192 words

(CLEAR_LOOP)
    @count
    D=M
    @DONE_CLEAR
    D;JEQ

    @addr
    A=M
    M=0            // write white (all 0s)

    @addr
    M=M+1

    @count
    M=M-1

    @CLEAR_LOOP
    0;JMP

(DONE_CLEAR)
    @LOOP
    0;JMP

