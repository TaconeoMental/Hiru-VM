# Hiru-VM
My personal attempt at writing an all purpose virtual machine for interpreted languages.

## Example
The Hiru Virtual machine is not ready yet, but a first version of the assembly language and its compiler is. The following example shows shows the process of compiling and disassembling a progam written in the Hiru assembly language.

### hello_world.asm
This is a simple program written in the Hiru assembly language.

<img width="60%" src="https://raw.githubusercontent.com/TaconeoMental/Hiru-VM/main/assets/assembly_example.png" />

### Compiled program
The previous program can be compiled into bytecode as follows:
```bash
$ ./assembler/assembler.rb hello_world.asm
```
This will produce ```hello_world.hbc``` and it will look something like this:

<img width="60%" src="https://raw.githubusercontent.com/TaconeoMental/Hiru-VM/main/assets/bytecode_example.png" />

### Disassembled bytecode
There is also a disassembler available. This will produce a human friendly version of a very close representation of the bytecode.

<img width="60%" src="https://raw.githubusercontent.com/TaconeoMental/Hiru-VM/main/assets/dis_example.png" />
