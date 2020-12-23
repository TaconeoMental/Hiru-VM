package vm

// NAME TABLE SEGMENT
//   +--------+--------+--------------------------------------------------------+
//   | Offset | Length | Description                                            |
//   +--------+--------+--------------------------------------------------------+
//   | 0      | 1      | The number of names in the table.                      |
//   +--------+--------+--------------------------------------------------------+
//
// NAME TABLE ENTRY
//   +--------+--------+--------------------------------------------------------+
//   | Offset | Length | Description                                            |
//   +--------+--------+--------------------------------------------------------+
//   | 0      | 1      | The length of the name.                                |
//   +--------+--------+--------------------------------------------------------+
//   | 1      | n      | The name.                                              |
//   +--------+--------+--------------------------------------------------------+

type NameSegment struct {
        numberEntries int32
        entries       []NameSegmentEntry
}

// Constructor de un segmento de índice.
func NewNameSegment(entries int32) *NameSegment {
        ns := NameSegment{numberEntries: entries}
        ns.entries = make([]NameSegmentEntry, 0)
        return &ns
}

func (ns *NameSegment) AddEntry(segment NameSegmentEntry) {
        ns.entries = append(ns.entries, segment)
}

func (ns *NameSegment) NameAt(index int32) NameSegmentEntry {
        return ns.entries[index]
}

type NameSegmentEntry struct {
        // El largo de la entrada contando su header
        length int32
        value  string
}

func NewNameSegmentEntry(value string, length int32) *NameSegmentEntry {
        return &NameSegmentEntry{length: length, value: value}
}

func (ns *NameSegment) Entries() []NameSegmentEntry {
        return ns.entries
}
