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
        numberEntries int32

        entries []HiruObject
}

func NewDataSegment(entries int32) *DataSegment {
        ds := DataSegment{numberEntries: entries}
        ds.entries = make([]HiruObject, 0)
        return &ds
}

func (ds *DataSegment) AddEntry(entry HiruObject) {
        ds.entries = append(ds.entries, entry)
}

func (ds *DataSegment) ConstantAt(index int32) HiruObject {
        return ds.entries[index]
}

// Tipos de entries para el data segment
type DataSegmentEntryType int32
const (
        typeNumberConstant DataSegmentEntryType = 0x69
        typeStringConstant = 0x73
        typeFunctionConstant = 0x63
)

// Entrada para el DataSegment
type DataSegmentEntry interface {
        getLength() int32 // In bytes
        getType() DataSegmentEntryType
}

// TIPO: Num
type NumberConstant struct {
        value int32
        vType int32
}

func NewNumberConstant(value int32, t int32) *NumberConstant {
        return &NumberConstant{value: value, vType: t}
}

func (i NumberConstant) getLength() int32 {
        return 4
}

func (i NumberConstant) getType() DataSegmentEntryType {
    return typeNumberConstant
}

func (i NumberConstant) getValue() int32 {
        return i.value
}

// TIPO: String
type StringConstant struct {
        value string
        length int32 // in bytes
}

func NewStringConstant(v string, l int32) *StringConstant {
        return &StringConstant{value: v, length: l}
}

func (s StringConstant) getLength() int32 {
        return int32(len(s.value))
}

func (s StringConstant) getType() DataSegmentEntryType {
    return typeStringConstant
}

func (s StringConstant) getValue() string {
        return s.value
}

// TIPO: Function

type FunctionConstant struct {
        value *CodeObject
        length int32
}

func NewFunctionConstant(cobj *CodeObject) *FunctionConstant {
        return &FunctionConstant{value: cobj}
}

func (f FunctionConstant) getLength() int32 {
        return f.length
}

func (f FunctionConstant) getType() DataSegmentEntryType {
    return typeFunctionConstant
}
 /*
func (f *FunctionConstant) getValue() *CodeObject {
        return f.value
}
*/
