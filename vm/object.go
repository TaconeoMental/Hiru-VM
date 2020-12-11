package vm


type IndexSegmentEntry {
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

type CodeObject struct {
        index IndexSegment
        data DataSegment
        names NameSegment
        code CodeSegment
}
