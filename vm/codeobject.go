package vm

type CodeObject struct {
        indexSegment *IndexSegment
        dataSegment  *DataSegment
        nameSegment  *NameSegment
        codeSegment  *CodeSegment
}
