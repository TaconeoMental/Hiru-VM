package vm

// Data segment //
// DATA  TABLE SEGMENT
//   +--------+--------+--------------------------------------------------------+
//   | Offset | Length | Description                                            |
//   +--------+--------+--------------------------------------------------------+
//   | 0      | 1      | The number of constants in the table.                  |
//   +--------+--------+--------------------------------------------------------+
//
// DATA TABLE ENTRY
//   +--------+--------+--------------------------------------------------------+
//   | Offset | Length | Description                                            |
//   +--------+--------+--------------------------------------------------------+
//   | 0      | 1      | The type of the constant. Must be one of:              |
//   |        |        |    0x00 - No constant                                  |
//   |        |        |    0x69 - Integer constant (ASCII 'i')                 |
//   |        |        |    0x66 - Floating point number constant (ASCII 'i')   |
//   |        |        |    0x73 - String constant (ASCII 's')                  |
//   |        |        |    0x63 - Code Obj. constant (ASCII 'c')               |
//   +--------+--------+--------------------------------------------------------+
//
//   - Intger Constants:
//     +--------+--------+--------------------------------------------------------+
//     | Offset | Length | Description                                            |
//     +--------+--------+--------------------------------------------------------+
//     | 0      | 1      | Null space.                                            |
//     +--------+--------+--------------------------------------------------------+
//     | 1      | 1      | Integer number. Stored as an Int64.                    |
//     +--------+--------+--------------------------------------------------------+
//
//   - Floating point Constants
//     +--------+--------+--------------------------------------------------------+
//     | Offset | Length | Description                                            |
//     +--------+--------+--------------------------------------------------------+
//     | 0      | 1      | Null space.                                            |
//     +--------+--------+--------------------------------------------------------+
//     | 1      | 1      | Floating point number. Stored as a Float64.            |
//     +--------+--------+--------------------------------------------------------+
//
//   - String Constants:
//     # TODO: Be able to specify encoding type.
//     +--------+--------+--------------------------------------------------------+
//     | Offset | Length | Description                                            |
//     +--------+--------+--------------------------------------------------------+
//     | 0      | 1      | Length of the string data in bytes.                    |
//     +--------+--------+--------------------------------------------------------+
//     | 1      | n      | String data with trailing zero padding as required.    |
//     +--------+--------+--------------------------------------------------------+
//
//   - Code Object Constants
//     +--------+--------+--------------------------------------------------------+
//     | Offset | Length | Description                                            |
//     +--------+--------+--------------------------------------------------------+
//     | 0      | 1      | Code object data length                                |
//     +--------+--------+--------------------------------------------------------+
//     | 1      | n      | Code object data                                       |
//     +--------+--------+--------------------------------------------------------+

type DataSegment struct {
        numberEntries uint32

        entries []DataSegmentEntry
}

func NewDataSegment(entries uint32) *DataSegment {
        ds := DataSegment{numberEntries: entries}
        ds.entries = make([]DataSegmentEntry, 0)
        return &ds
}

func (ds *DataSegment) AddEntry(entry DataSegmentEntry) {
        ds.entries = append(ds.entries, entry)
}

// Tipos de entries para el data segment
type DataSegmentEntryType uint32
const (
        typeNumberConstant DataSegmentEntryType = iota
        typeStringConstant
        typeFunctionConstant
)

// Entrada para el DataSegment
type DataSegmentEntry interface {
        getLength() uint32 // In bytes
        getType() uint8
}

// TIPO: Num
type NumberConstant struct {
        value uint32
}

func (i NumberConstant) getLength() uint32 {
        return 4
}

func (i NumberConstant) getType() DataSegmentEntryType {
    return typeNumberConstant
}

func (i NumberConstant) getValue() uint32 {
        return i.value
}

// TIPO: String
type StringConstant struct {
        value string
        length uint32 // in bytes
}

func (s StringConstant) getLength() uint32 {
        return uint32(len(s.value))
}

func (s StringConstant) getType() DataSegmentEntryType {
    return typeStringConstant
}

func (s StringConstant) getValue() string {
        return s.value
}

// TIPO: Function

type FunctionConstant struct {
        value string
        length uint32 // in bytes
}

func (f FunctionConstant) getLength() uint32 {
        return f.length
}

func (f FunctionConstant) getType() DataSegmentEntryType {
    return typeFunctionConstant
}
