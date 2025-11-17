## HACKASM

`hackasm` is a lightweight, fast, command-line assembler for the Hack computer (from the Nand2Tetris course).
It converts `.asm` programs into `.hack` binary machine code, supporting both individual files and whole directories.

---

## Setup

### 1. Install Go (if not already installed)

Go 1.25.2 required.

Download: [https://go.dev/dl/](https://go.dev/dl/)

### 2. Clone the project

```sh
git clone https://github.com/yourname/hackasm
cd hackasm
```

### 3. Build the binary

```sh
go build -o hackasm
```

This creates an executable named `hackasm` (or `hackasm.exe` on Windows).

---

## Usage

```
hackasm [options] FILES|DIRS
```

### Assemble a single file

```sh
./hackasm Add.asm
```

Generates:

```
Add.hack
```

in the same directory.

---

### Assemble all `.asm` files in a directory

```sh
./hackasm projects/06
```

---

### Specify an output directory

Use `-o` to place compiled `.hack` files elsewhere:

```sh
./hackasm -o out/ Add.asm
```

Output:

```
out/Add.hack
```

---

### Test mode (compare with reference output)

If you have a reference directory (e.g., Nand2Tetris `cmp/` folder):

```sh
./hackasm --test reference/
```

This runs your assembler, compares output, and reports differences.

---

## Example

```sh
./hackasm -o bin/ MyProgram.asm
```

Output:

```
Assembling: MyProgram.asm
 → Generated: bin/MyProgram.hack
Done.
```

---

## Notes

* Non-`.asm` files are ignored when assembling directories.
* Symbol table resets for each file automatically.
* Output directory is created if it does not exist.