package vm

// BYTECODE SEGMENT
//   Each opcode has a length of 1 byte.
//   +--------+--------+--------------------------------------------------------+
//   | Offset | Length | Description                                            |
//   +--------+--------+--------------------------------------------------------+
//   | 0      | 1      | A Hiru-VM opcode.                                      |
//   +--------+--------+--------------------------------------------------------+
//
//   And each instruction takes exactly one argument.
//   +--------+--------+--------------------------------------------------------+
//   | Offset | Length | Description                                            |
//   +--------+--------+--------------------------------------------------------+
//   | 0      | 1      | An integer (represented as 1 byte).                    |
//   +--------+--------+--------------------------------------------------------+

type CodeSegment struct {
        numberEntries uint32
        entries       []Instruction
}

type Instruction struct {
        opcode   Opcode
        argument uint32
}

func NewInstruction(op Opcode, arg uint32) *Instruction {
        ins := Instruction{
                opcode: op,
                argument: arg}
        return &ins
}
