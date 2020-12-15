package vm

// En este archivo se definen las estructuras que representan los distintos
// objetos de código representados en el bytecode. La especificación completa
// está en docs/bytecode.txt.

// Index Table Segment //

// Representa el segmento "Indice" de un objeto de código. Este guarda
// información acerca de la posición absoluta de sus demás partes (Segmento de
// nombres, segmento de código, etc).
// Todo está detallado en docs/bytecode.txt, pero igual pondré pedazos de ese
// texto acá para que se entienda:
//
// INDEX TABLE SEGMENT
// +--------+--------+--------------------------------------------------------+
// | Offset | Length | Description                                            |
// +--------+--------+--------------------------------------------------------+
// | 0      | 1      | The number of entries in the table.                    |
// +--------+--------+--------------------------------------------------------+
//
// INDEX TABLE ENTRY
// +--------+--------+--------------------------------------------------------+
// | Offset | Length | Description                                            |
// +--------+--------+--------------------------------------------------------+
// | 0      | 1      | The type of the segment. Must be one of the following: |
// |        |        |    0x00 - Index Table Segment                          |
// |        |        |    0x01 - Constant Table Segment                       |
// |        |        |    0x02 - Names Table Segment                          |
// |        |        |    0x03 - Bytecode Segment                             |
// +--------+--------+--------------------------------------------------------+
// | 1      | 1      | The offset to the segment, relative to the start of    |
// |        |        | the file.                                              |
// +--------+--------+--------------------------------------------------------+
// | 1      | 1      | The length of the segment, including its header.       |
// +--------+--------+--------------------------------------------------------+
type IndexSegment struct {
        // Cantidad de entradas en el segmento (Siempre debería ser 4)
        entries  int

        // Entradas individuales
        segments []IndexSegmentEntry
}

// Constructor(?) de un segmento de índice.
func NewIndexSegment(entries int) *IndexSegment {
        is := IndexSegment{entries: entries}
        is.segments = make([]IndexSegmentEntry, entries)
        return &is
}

const (
        typeIndexSegment = iota
        typeConstantSegment
        typeNamesSegment
        typeCodeSegment
)

type IndexSegmentEntryType uint8
type IndexSegmentEntry struct {
        // Tipo de entrada
        // 0x00 - Index Table Segment
        // 0x01 - Constant Table Segment
        // 0x02 - Names Segment
        // 0x03 - Bytecode Segment
        entryType     IndexSegmentEntryType

        // El desfase del inicio de la entrada relativo al inicio del archivo
        relativeStart byte

        // El largo de la entrada contando su header
        length        byte
}

// Data segment //
// +--------+--------+--------------------------------------------------------+
// | Offset | Length | Description                                            |
// +--------+--------+--------------------------------------------------------+
// | 0      | 1      | The number of constants in the table.                  |
// +--------+--------+--------------------------------------------------------+
//
// CONSTANTS TABLE ENTRY
// +--------+--------+--------------------------------------------------------+
// | Offset | Length | Description                                            |
// +--------+--------+--------------------------------------------------------+
// | 0      | 1      | The type of the constant. Must be one of:              |
// |        |        |    0x00 - No constant                                  |
// |        |        |    0x6E - Number constant (ASCII 'n')                  |
// |        |        |    0x73 - String constant (ASCII 's')                  |
// |        |        |    0x66 - Code Obj. constant (ASCII 'f')               |
// +--------+--------+--------------------------------------------------------+
//
// - Number Constants:
//   Stored as an Int64 (8 bytes)
//
// - String Constants:
//   # TODO: Be able to specify encoding type.
//   +--------+--------+--------------------------------------------------------+
//   | Offset | Length | Description                                            |
//   +--------+--------+--------------------------------------------------------+
//   | 0      | 1      | Length of the string data in bytes.                    |
//   +--------+--------+--------------------------------------------------------+
//   | 1      | n      | String data with trailing zero padding as required.    |
//   +--------+--------+--------------------------------------------------------+
//
// - Code Object Constants
//   +--------+--------+--------------------------------------------------------+
//   | Offset | Length | Description                                            |
//   +--------+--------+--------------------------------------------------------+
//   | 0      | 1      | Code object data length                                |
//   +--------+--------+--------------------------------------------------------+
//   | 1      | n      | Code object data                                       |
//   +--------+--------+--------------------------------------------------------+

type DataSegment struct {
        numberEntries uint64

        //data []DataSegmentEntry
}

type CodeObject struct {
        index IndexSegment
        data DataSegment
        //names NameSegment
        //code CodeSegment
}
