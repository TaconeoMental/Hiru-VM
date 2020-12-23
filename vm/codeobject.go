package vm

type CodeObject struct {
        indexSegment    *IndexSegment
        dataSegment     *DataSegment
        nameSegment     *NameSegment
        bytecodeSegment *BytecodeSegment
}

func NewCodeObject(indexSeg *IndexSegment, dataSeg *DataSegment,
                nameSeg *NameSegment, bytecodeSeg *BytecodeSegment) *CodeObject {
        return &CodeObject{
                indexSegment: indexSeg,
                dataSegment: dataSeg,
                nameSegment: nameSeg,
                bytecodeSegment: bytecodeSeg}
}

func (vm *HiruVM) CurrentObject() *CodeObject {
        return vm.callStack.GetTopMost().GetObject()
}

// Lo que dice. Lee el segmento de indice del bytecode
func ReadIndexSegment(hf *HiruFile) (*IndexSegment, error) {
        entries := hf.Read4Bytes()

        index_seg := NewIndexSegment(entries)

        for i := int32(1); i <= entries; i++ {
                etype := hf.Read4Bytes()
                start := hf.Read4Bytes()
                length := hf.Read4Bytes()

                entry := NewIndexSegmentEntry(etype, start, length)
                index_seg.AddSegment(*entry)
        }

        return index_seg, nil
}

func ReadDataSegment(hf *HiruFile) (*DataSegment, error) {
        entries := hf.Read4Bytes()

        data_seg := NewDataSegment(entries)

        for i := int32(1); i <= entries; i++ {
                etype := hf.Read4Bytes()

                length := hf.Read4Bytes()

                switch DataSegmentEntryType(etype) {
                case typeNumberConstant:
                        // En el futuro el largo debería ser el indicador de si es un entero o un flotante
                        value := hf.Read4Bytes()
                        number := new(HiruNumber)
                        number.Value = int64(value)
                        data_seg.AddEntry(number)
                        // data_seg.AddEntry(NewNumberConstant(value, length))

                case typeStringConstant:
                        str := string(hf.ReadBytes(int(length)))

                        hstr := new(HiruString)
                        hstr.Value = str
                        data_seg.AddEntry(hstr)
                        // data_seg.AddEntry(NewStringConstant(str, length))

                case typeFunctionConstant:
                        codeObj := ReadObject(hf)
                        hfunc := new(HiruFunction)
                        hfunc.CodeObject = codeObj

                        data_seg.AddEntry(hfunc)
                        // data_seg.AddEntry(NewFunctionConstant(codeObj))
                }
        }

        return data_seg, nil
}

func ReadNameSegment(hf *HiruFile) (*NameSegment, error) {
        entries := hf.Read4Bytes()

        name_seg := NewNameSegment(entries)

        for i := int32(1); i <= entries; i++ {
                length := hf.Read4Bytes()

                name := string(hf.ReadBytes(int(length)))


                name_seg.AddEntry(*NewNameSegmentEntry(name, length))
        }

        return name_seg, nil
}

func ReadBytecodeSegment(hf *HiruFile) (*BytecodeSegment, error) {
        entries := hf.Read4Bytes()

        bytecode_seg := NewBytecodeSegment(entries)

        for i := int32(1); i <= entries; i++ {
                op := hf.Read4Bytes()

                arg := hf.Read4Bytes()

                bytecode_seg.AddEntry(*NewInstruction(Opcode(op), arg))
        }

        return bytecode_seg, nil
}

func ReadObject(hf *HiruFile) *CodeObject {
        is, _ := ReadIndexSegment(hf)
        ds, _ := ReadDataSegment(hf)
        ns, _ := ReadNameSegment(hf)
        bs, _ := ReadBytecodeSegment(hf)
        return NewCodeObject(is, ds, ns, bs)
}


