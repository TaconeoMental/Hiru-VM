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

type BytecodeSegment struct {
        numberEntries int32
        entries       []Instruction
}

func NewBytecodeSegment(entries int32) *BytecodeSegment {
        return &BytecodeSegment{
                numberEntries: entries,
                entries: make([]Instruction, 0)}
}

func (bs *BytecodeSegment) AddEntry(entry Instruction) {
        bs.entries = append(bs.entries, entry)
}

func (bs *BytecodeSegment) Instructions() []Instruction {
        return bs.entries
}

func (bs *BytecodeSegment) InstructionAt(index int32) Instruction {
        return bs.entries[index]
}


type Instruction struct {
        opcode   Opcode
        argument int32
}

func NewInstruction(op Opcode, arg int32) *Instruction {
        ins := Instruction{
                opcode: op,
                argument: arg}
        return &ins
}
