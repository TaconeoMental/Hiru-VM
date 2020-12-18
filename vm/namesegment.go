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
        numberEntries uint32
        entries       []NameSegmentEntry
}

// Constructor de un segmento de índice.
func NewNameSegment(entries uint32) *NameSegment {
        ns := NameSegment{numberEntries: entries}
        ns.entries = make([]NameSegmentEntry, 0)
        return &ns
}

func (ns *NameSegment) AddSegment(segment NameSegmentEntry) {
        ns.entries = append(ns.entries, segment)
}

type NameSegmentEntry struct {
        // El largo de la entrada contando su header
        length uint32
        value  string
}

func (ns *NameSegment) Entries() []NameSegmentEntry {
        return ns.entries
}
