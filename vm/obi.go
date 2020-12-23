package vm

import (
        "math"
        "fmt"
)

// Busca una constante en el data segment del objecto actual
func (vm *HiruVM) GetConstantAt(index int32) HiruObject {
        object := vm.CurrentObject()
        return object.dataSegment.ConstantAt(index)
}

// Busca un nombre en el name segment del objecto actual
func (vm *HiruVM) GetNameAt(index int32) NameSegmentEntry {
        object := vm.CurrentObject()
        return object.nameSegment.NameAt(index)
}


// Corre la máquina virtual en modo OBI
func (vm *HiruVM) runObjectBasedVm() (error) {
        vm.DebugPrint("Running in OBI Mode")

        mainObject := ReadObject(vm.mainFile)

        sf := NewStackFrame("global", mainObject, 0)
        vm.callStack.Push(sf)

        vm.RunOBIBytecode(mainObject.bytecodeSegment)

        return nil

}

func (vm *HiruVM) RunOBIBytecode(bs *BytecodeSegment) {
        for vm.ip < bs.numberEntries {
                instruction := vm.LoadInstruction()
                if instruction.opcode == EXIT {
                        return
                }

                vm.RunOBInstruction(instruction)

        }
}

func (vm *HiruVM) LoadInstruction() Instruction {
        object := vm.CurrentObject()

        return object.bytecodeSegment.InstructionAt(vm.ip)
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

                vm.DebugPrint("Instruction BMOD: %v %% %v = %v", (*num1).Value, (*num2).Value, (*num1).Value % (*num2).Value)
                result := new(HiruNumber)
                result.Value = (*num1).Value % (*num2).Value
                vm.objectStack.Push(result)

        case BSUB:
                operand2, _ := vm.objectStack.Pop()
                operand1, _ := vm.objectStack.Pop()

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
                vm.DebugPrint("%v == %v", operand1.Inspect(), operand2.Inspect())
                result.Value = operand1.Inspect() == operand2.Inspect()
                vm.objectStack.Push(result)

        case CMPLT:
                vm.DebugPrint("Instruction CMPLT")
                operand2, _ := vm.objectStack.Pop()
                operand1, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)

                vm.DebugPrint("Instruction CMPLT: %v < %v = %v", (*num1).Value, (*num2).Value, (*num1).Value < (*num2).Value)

                result := new(HiruBoolean)
                result.Value = (*num1).Value < (*num2).Value
                vm.objectStack.Push(result)

        case CMPLE:
                vm.DebugPrint("Instruction CMPLE")
                operand2, _ := vm.objectStack.Pop()
                operand1, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)

                vm.DebugPrint("Instruction CMPLE: %v <= %v = %v", (*num1).Value, (*num2).Value, (*num1).Value <= (*num2).Value)

                result := new(HiruBoolean)
                result.Value = (*num1).Value <= (*num2).Value
                vm.objectStack.Push(result)


        case RET:
                // TODO: No sé muy bien qué hacer acá todavía
                vm.DebugPrint("Instruction RET")
                retVal, _ := vm.callStack.Pop()
                vm.DebugPrint("%v", retVal.ReturnAddress)
                vm.ip = retVal.ReturnAddress

        case PRINT:
                operand, _ := vm.objectStack.Pop()
                fmt.Println(operand.Inspect())

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

                sf := NewStackFrame(nameString, moduleCodeObject, vm.ip)
                vm.callStack.Push(sf)

                vm.RunOBIBytecode(vm.CurrentObject().bytecodeSegment)

                moduleObject := new(HiruModule)
                moduleObject.CodeObject = moduleCodeObject
                moduleObject.StackFrame, _ = vm.callStack.Pop()

                vm.objectStack.Push(moduleObject)

        case CALLFN:
                args := make([]HiruObject, instruction.argument)
                for i := 0; int32(i) < instruction.argument; i++ {
                        pop, _ := vm.objectStack.Pop()
                        args[i] = pop
                }

                hfunc, _ := vm.objectStack.Pop()
                vm.DebugPrint("Instruction CALLFN: %d arguments", instruction.argument)

                hfuncStruct := hfunc.(*HiruFunction) // Type assertion
                codeObject := (*hfuncStruct).RawObject()
                names := codeObject.nameSegment.Entries()

                // Empujamos un nuevo StackFrame al CallStack
                sf := NewStackFrame("TODO", codeObject, vm.ip)
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
                for i := 0; int32(i) < instruction.argument; i++ {
                        vm.DebugPrint("name %v => %v", names[i], args[instruction.argument - int32(i) - 1])
                        vm.callStack.Define(names[i].value, args[i])
                }

                func_body := codeObject.bytecodeSegment

                vm.ip = 0

                // TODO: Meter todo esto en un método (fn *HiruFunction).Call(args ...interface{})
                vm.RunOBIBytecode(func_body)
                return

        case SLOOP:
                vm.DebugPrint("Instruction SLOOP: New loop block handled at %v", instruction.argument)
                loopBlock := NewBlock("loop", instruction.argument)
                vm.callStack.PushBlock(loopBlock)

        case PLOOP:
                vm.DebugPrint("Instruction PLOOP: Loop block popped")
                //vm.callStack.PopBlock()

        case JMPABS:
                vm.DebugPrint("Instruction JMPABS to %v", instruction.argument)
                vm.ip = instruction.argument / 8
                return

        case PJMPF:
                vm.DebugPrint("Instruction PJMPF")

                val, _ := vm.objectStack.Pop()
                boolean := val.(*HiruBoolean)

                if !boolean.Value {
                        vm.ip = instruction.argument / 8
                        return
                }

        case PJMPT:
                vm.DebugPrint("Instruction PJMPT")

                val, _ := vm.objectStack.Pop()
                boolean := val.(*HiruBoolean)

                if boolean.Value {
                        vm.ip = instruction.argument / 8
                        return
                }

        case JUMPFWD:
                vm.DebugPrint("Instruction JUMPFWD")
        case BSTR:
        case BLIST:
                elems := make([]HiruObject, instruction.argument)
                for i := 0; int32(i) < instruction.argument; i++ {
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

        case LCONST:
                constant := vm.GetConstantAt(instruction.argument)

                vm.DebugPrint("Instruction LCONST: %v", constant.Inspect())
                vm.objectStack.Push(constant)

        case SNAME:
                name := vm.GetNameAt(instruction.argument)
                object, err := vm.objectStack.Pop()
                if err != nil {
                        vm.DebugPrint(err.Error())
                }

                vm.DebugPrint("Instruction SNAME: '%s' => %v", name.value, object.Inspect())

                vm.callStack.Define(name.value, object)

        case EXIT:
        }

        vm.ip += 1
}
