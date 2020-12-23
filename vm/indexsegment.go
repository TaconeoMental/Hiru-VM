package vm

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
        numberEntries  int32

        // Entradas individuales. Nombre exportado
        entries []IndexSegmentEntry
}

// Constructor de un segmento de índice.
func NewIndexSegment(entries int32) *IndexSegment {
        is := new(IndexSegment)
        is.numberEntries = entries
        is.entries = make([]IndexSegmentEntry, entries)
        return is
}

func (is *IndexSegment) AddSegment(segment IndexSegmentEntry) {
        is.entries[segment.Type()] = segment
}

type IndexSegmentEntryType int32
const (
        typeIndexSegment IndexSegmentEntryType = iota
        typeConstantSegment
        typeNamesSegment
        typeCodeSegment
)

type IndexSegmentEntry struct {
        // Tipo de entrada
        // 0x00 - Index Table Segment
        // 0x01 - Constant Table Segment
        // 0x02 - Names Segment
        // 0x03 - Bytecode Segment
        entryType     IndexSegmentEntryType

        // El desfase del inicio de la entrada relativo al inicio del archivo
        relativeStart int32

        // El largo de la entrada contando su header
        length        int32
}

func NewIndexSegmentEntry(t int32, start int32, length int32) *IndexSegmentEntry {
        return &IndexSegmentEntry{
                entryType: IndexSegmentEntryType(t),
                relativeStart: start,
                length: length}
}

func (ise IndexSegmentEntry) Type() IndexSegmentEntryType {
        return ise.entryType
}

func (hf *IndexSegment) Entries() []IndexSegmentEntry {
        return hf.entries
}
