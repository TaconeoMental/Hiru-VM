package vm

import (
        "math"
)

type CodeObject struct {
        indexSegment    *IndexSegment
        dataSegment     *DataSegment
        nameSegment     *NameSegment
        bytecodeSegment *BytecodeSegment

        ip               uint32
}

func NewCodeObject(indexSeg *IndexSegment, dataSeg *DataSegment,
                nameSeg *NameSegment, bytecodeSeg *BytecodeSegment) *CodeObject {
        return &CodeObject{
                indexSegment: indexSeg,
                dataSegment: dataSeg,
                nameSegment: nameSeg,
                bytecodeSegment: bytecodeSeg}
}

func (vm *HiruVM) getCurrentObject() *CodeObject {
        var object *CodeObject
        if vm.tempObject != nil {
                object = vm.tempObject
        } else {
                object = vm.mainObject
        }
        return object
}

// Busca una constante en el data segment del objecto actual
func (vm *HiruVM) GetConstantAt(index uint32) HiruObject {
        object := vm.getCurrentObject()
        return object.dataSegment.ConstantAt(index)
}

// Busca un nombre en el name segment del objecto actual
func (vm *HiruVM) GetNameAt(index uint32) NameSegmentEntry {
        object := vm.getCurrentObject()
        return object.nameSegment.NameAt(index)
}

// Lo que dice. Lee el segmento de indice del bytecode
func ReadIndexSegment(hf *HiruFile) (*IndexSegment, error) {
        entries := hf.Read4Bytes()

        index_seg := NewIndexSegment(entries)

        for i := uint32(1); i <= entries; i++ {
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

        for i := uint32(1); i <= entries; i++ {
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

        for i := uint32(1); i <= entries; i++ {
                length := hf.Read4Bytes()

                name := string(hf.ReadBytes(int(length)))


                name_seg.AddEntry(*NewNameSegmentEntry(name, length))
        }

        return name_seg, nil
}

func ReadBytecodeSegment(hf *HiruFile) (*BytecodeSegment, error) {
        entries := hf.Read4Bytes()

        bytecode_seg := NewBytecodeSegment(entries)

        for i := uint32(1); i <= entries; i++ {
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


// Corre la máquina virtual en modo OBI
func (vm *HiruVM) runObjectBasedVm() (error) {
        vm.DebugPrint("Running in OBI Mode")
        vm.isReturn = false
        mainObject := ReadObject(vm.mainFile)
        vm.mainObject = mainObject

        vm.RunOBIBytecode(mainObject.bytecodeSegment)

        // Por ahora, imprimir el StackFrame final al terminar la ejcución.
        vm.callStack.PrettyPrint()
        return nil

}

func (vm *HiruVM) RunOBIBytecode(bs *BytecodeSegment) {
        for true {
                instruction := vm.LoadInstruction()
                vm.RunOBInstruction(instruction)

                if vm.isReturn {
                        break
                }

        }
}

func (vm *HiruVM) LoadInstruction() Instruction {
        object := vm.getCurrentObject()
        return object.bytecodeSegment.InstructionAt(object.ip)
}

// Corre una instrucción
func (vm *HiruVM) RunOBInstruction(instruction Instruction) {
        switch instruction.opcode {
        case POP:
                _, err := vm.objectStack.Pop()
                if err != nil {
                        return
                }
        case UPOS:
        case UNEG:
        case UNOT:

        case BPOW:
                operand1, _ := vm.objectStack.Pop()
                operand2, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)

                result := new(HiruNumber)
                result.Value = int64(math.Pow(float64((*num1).Value), float64((*num2).Value)))
                vm.objectStack.Push(result)

        case BMUL:
                operand1, _ := vm.objectStack.Pop()
                operand2, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)

                result := new(HiruNumber)
                result.Value = (*num1).Value * (*num2).Value
                vm.objectStack.Push(result)

        case BDIV:
                operand1, _ := vm.objectStack.Pop()
                operand2, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)

                result := new(HiruNumber)
                result.Value = (*num1).Value / (*num2).Value
                vm.objectStack.Push(result)

        case BMOD:
                operand1, _ := vm.objectStack.Pop()
                operand2, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)

                result := new(HiruNumber)
                result.Value = (*num1).Value % (*num2).Value
                vm.objectStack.Push(result)

        case BSUB:
                operand1, _ := vm.objectStack.Pop()
                operand2, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)

                result := new(HiruNumber)
                result.Value = (*num1).Value - (*num2).Value
                vm.objectStack.Push(result)

        case BADD:
                operand1, _ := vm.objectStack.Pop()
                operand2, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)

                result := new(HiruNumber)
                result.Value = (*num1).Value + (*num2).Value
                vm.objectStack.Push(result)

        case BAND:
        case BOR:
        case CMPEQ:
                vm.DebugPrint("Instruction CMPEQ")
                operand1, _ := vm.objectStack.Pop()
                operand2, _ := vm.objectStack.Pop()

                result := new(HiruBoolean)
                result.Value = operand1.Inspect() == operand2.Inspect()
                vm.objectStack.Push(result)
                vm.getCurrentObject().ip += 1

        case CMPLT:
                vm.DebugPrint("Instruction CMPLT")
                operand1, _ := vm.objectStack.Pop()
                operand2, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)


                result := new(HiruBoolean)
                result.Value = (*num1).Value < (*num2).Value
                vm.objectStack.Push(result)
                vm.getCurrentObject().ip += 1


        case RET:
                // TODO: No sé muy bien qué hacer acá todavía
                vm.DebugPrint("Instruction RET")
                vm.isReturn = true
                vm.getCurrentObject().ip += 1

        case MAKEFN:
                // No se usa en OBI
        case MAKEMOD:
                nameObject, _ := vm.objectStack.Pop()
                nameString := nameObject.Inspect() + ".hbc"
                vm.DebugPrint("Instruction MAKEMOD: module name '%s'", nameString)

                newFile, err := NewHiruFile(nameString)
                if err != nil {
                        vm.DebugPrint(err.Error())
                        return
                }
                magic := newFile.Read4Bytes()
                if magic != 0x48495255 {
                        return
                }
                moduleCodeObject := ReadObject(newFile)
                vm.tempObject = moduleCodeObject

                sf := NewStackFrame(nameString)
                vm.callStack.Push(sf)

                vm.RunOBIBytecode(vm.tempObject.bytecodeSegment)

                moduleObject := new(HiruModule)
                moduleObject.CodeObject = moduleCodeObject
                moduleObject.StackFrame, _ = vm.callStack.Pop()

                vm.objectStack.Push(moduleObject)

                vm.tempObject = nil
                vm.getCurrentObject().ip += 1

        case CALLFN:
                args := make([]HiruObject, instruction.argument)
                for i := 0; uint32(i) < instruction.argument; i++ {
                        pop, _ := vm.objectStack.Pop()
                        args[i] = pop
                }

                hfunc, _ := vm.objectStack.Pop()
                vm.DebugPrint("Instruction CALLFN: %d arguments", instruction.argument)

                hfuncStruct := hfunc.(*HiruFunction) // Type assertion
                codeObject := (*hfuncStruct).RawObject()
                names := codeObject.nameSegment.Entries()

                // Empujamos un nuevo StackFrame al CallStack
                sf := NewStackFrame("TODO")
                sf.MakeLinkTo(vm.callStack.GetTopMost())
                vm.callStack.Push(sf)

                // Ahora asignaremos los valores de args a los primeros n
                // nombres que tenga definida la función, con n =
                // instruction.argument.
                //
                // En términos generales, estamos simulando un:
                // lconst 0
                // sname  0
                // lconst 1
                // sname  1
                // ...
                // lconst n
                // sname  n
                //
                // En donde los valores de const son los del arreglo "args"
                for i := 0; uint32(i) < instruction.argument; i++ {
                        vm.DebugPrint("name %v => %v", names[i], args[i])
                        vm.callStack.Define(names[i].value, args[i])
                }

                func_body := codeObject.bytecodeSegment

                // TODO: Meter todo esto en un método (fn *HiruFunction).Call(args ...interface{})
                vm.tempObject = codeObject

                vm.RunOBIBytecode(func_body)

                vm.callStack.Pop()
                vm.tempObject = nil
                vm.getCurrentObject().ip += 1

        case JMPABS:
                vm.DebugPrint("Instruction JMPABS")
                vm.getCurrentObject().ip += instruction.argument / 8

        case PJMPF:
                vm.DebugPrint("Instruction PJMPF")

                val, _ := vm.objectStack.Pop()
                boolean := val.(*HiruBoolean)

                if !boolean.Value {
                        vm.getCurrentObject().ip += instruction.argument / 8
                } else {
                        vm.getCurrentObject().ip += 1
                }
        case PJMPT:
                vm.DebugPrint("Instruction PJMPT")

                val, _ := vm.objectStack.Pop()
                boolean := val.(*HiruBoolean)

                if boolean.Value {
                        vm.getCurrentObject().ip += instruction.argument / 8
                } else {
                        vm.getCurrentObject().ip += 1
                }
        case JUMPFWD:
                vm.DebugPrint("Instruction JUMPFWD")
        case BSTR:
        case BLIST:
                elems := make([]HiruObject, instruction.argument)
                for i := 0; uint32(i) < instruction.argument; i++ {
                        pop, _ := vm.objectStack.Pop()
                        elems[i] = pop
                }


        case LATTR:
                name := vm.GetNameAt(instruction.argument)
                vm.DebugPrint("Instruction LATTR: name '%v'", name)

                object, _ := vm.objectStack.Pop()

                // Asserteamos (?) que module es tipo HiruModule
                module := object.(*HiruModule)
                attr, _ := module.StackFrame.ResolveName(name.value)

                vm.objectStack.Push(attr)
                vm.getCurrentObject().ip += 1

        case IMPORT:
                // No es necesario para OBI
        case LNAME:
                name := vm.GetNameAt(instruction.argument)
                object, err := vm.callStack.ResolveName(name.value)
                if err != nil {
                        vm.DebugPrint(err.Error())
                }

                vm.DebugPrint("Instruction LNAME: name '%v' resolved to %v", name.value, object.Inspect())
                vm.objectStack.Push(object)
                vm.getCurrentObject().ip += 1

        case LCONST:
                constant := vm.GetConstantAt(instruction.argument)

                vm.DebugPrint("Instruction LCONST: %v", constant.Inspect())
                vm.objectStack.Push(constant)
                vm.getCurrentObject().ip += 1

        case SNAME:
                name := vm.GetNameAt(instruction.argument)
                object, err := vm.objectStack.Pop()
                if err != nil {
                        vm.DebugPrint(err.Error())
                }

                vm.DebugPrint("Instruction SNAME: '%s' => %v", name.value, object.Inspect())

                vm.callStack.Define(name.value, object)
                vm.getCurrentObject().ip += 1
        }
}
