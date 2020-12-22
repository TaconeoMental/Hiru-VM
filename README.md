# Hiru-VM
My personal attempt at writing an all purpose virtual machine for interpreted languages.

## Example
To run a program in the Hiru VM one must first write it in the Hiru assembly language, then it must be compiled into a bytecode file which can then be fed into the VM.

### fib.asm
This is a simple recursive fibonacci program written in the Hiru assembly language.

<img width="60%" src="https://raw.githubusercontent.com/TaconeoMental/Hiru-VM/main/assets/assembly_example.png" />

### Compiled program
The previous program can be compiled into bytecode as follows:
```bash
$ ./assembler/assembler.rb fib.asm
```
This will produce ```fib.hbc``` and it will look something like this:

<img width="60%" src="https://raw.githubusercontent.com/TaconeoMental/Hiru-VM/main/assets/bytecode_example.png" />

### Disassembled bytecode
There is also a disassembler available. This will produce a human friendly version of a very close representation of the bytecode.
To dissasemble the previously compiled bytecode file you can run the following command:
```bash
$ ./assembler/dis.rb fib.hbc
```
Which will output this text into the terminal:

<img width="60%" src="https://raw.githubusercontent.com/TaconeoMental/Hiru-VM/main/assets/dis_example.png" />

### Running
Before anything the VM must be compiled into a binary. This can be easily done like so:
```bash
$ go build -o hiru
```
Then you will be able to run a bytecode file as follows:
```bash
$ ./hiru fib.hbc
6765
```
