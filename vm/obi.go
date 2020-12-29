package vm

import (
        "math"
        "fmt"
        "os"
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
        vm.DebugPrint("%v", mainObject)

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
                vm.DebugPrint("Instruction BPOW: %v ^ %v = %v", (*num1).Value, (*num2).Value, int64(math.Pow(float64((*num1).Value), float64((*num2).Value))))
                result.Value = int64(math.Pow(float64((*num1).Value), float64((*num2).Value)))
                vm.objectStack.Push(result)

        case BMUL:
                operand1, _ := vm.objectStack.Pop()
                operand2, _ := vm.objectStack.Pop()

                num1 := operand1.(*HiruNumber)
                num2 := operand2.(*HiruNumber)

                vm.DebugPrint("Instruction BMUL: %v * %v = %v", (*num1).Value, (*num2).Value, (*num1).Value * (*num2).Value)
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
                retVal, _ := vm.callStack.Pop()
                vm.DebugPrint("Instruction RET, returning to %v", retVal.ReturnAddress)
                vm.ip = retVal.ReturnAddress

        case PRINT:
                vm.DebugPrint("Instruction PRINT")
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

                vm.ip = 0
                vm.RunOBIBytecode(vm.CurrentObject().bytecodeSegment)

                moduleObject := new(HiruModule)
                moduleObject.CodeObject = moduleCodeObject
                moduleObject.StackFrame, _ = vm.callStack.Pop()

                vm.ip = moduleObject.StackFrame.ReturnAddress

                vm.objectStack.Push(moduleObject)

        case CALLFN:
                args := make([]HiruObject, instruction.argument)
                for i := 0; int32(i) < instruction.argument; i++ {
                        pop, _ := vm.objectStack.Pop()
                        args[i] = pop
                }

                hfunc, _ := vm.objectStack.Pop()
                vm.DebugPrint("Instruction CALLFN: %d arguments, IP: %v", instruction.argument, vm.ip)
                vm.DebugPrint("BLOCKTYPE: %v", vm.callStack.GetCurrentBlockType())


                hfuncStruct := hfunc.(*HiruFunction) // Type assertion
                codeObject := (*hfuncStruct).RawObject()
                names := codeObject.nameSegment.Entries()

                // Empujamos un nuevo StackFrame al CallStack
                sf := NewStackFrame("Function", codeObject, vm.ip)
                sf.MakeLinkTo(vm.callStack.GetTopMost())

                if vm.callStack.GetCurrentBlockType() == "METHOD" {
                        vm.DebugPrint("yay")
                        object, _ := vm.objectStack.Pop()
                        sf.Define("__self__", object)

                }
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
                vm.callStack.PopBlock()
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
                object, _ := vm.objectStack.Pop()
                name := vm.GetNameAt(instruction.argument)
                vm.DebugPrint("Instruction LATTR: name '%v' of object %v", name, object.Inspect())

                switch object.(type) {
                case *HiruModule:
                        // Asserteamos (?) que module es tipo HiruModule
                        module := object.(*HiruModule)
                        attr, _ := module.StackFrame.ResolveName(name.value)
                        vm.objectStack.Push(attr)

                case *HiruInstance:
                        vm.DebugPrint("Is instance")
                        instance := object.(*HiruInstance)

                        attr, _ := instance.StackFrame.StackFrame().ResolveName(name.value)
                        vm.DebugPrint("%v resolved to %v", name, attr)

                        if attr.Type() == "Function" {
                                vm.callStack.PushBlock(&Block{"METHOD", vm.ip})
                                vm.objectStack.Push(instance)
                        }
                        vm.objectStack.Push(attr)
                }

        case LSELF:
                vm.DebugPrint("Instruction LSELF")

                object, _ := vm.callStack.ResolveName("__self__")
                instance := object.(*HiruInstance)
                vm.objectStack.Push(instance)
                vm.DebugPrint("%v", instance.StackFrame.StackFrame().GetEnv())
                vm.DebugPrint("END LSELF")

        case SATTR:
                hiruInstance, _ := vm.objectStack.Pop()
                hiruValue, _ := vm.objectStack.Pop()

                instance := hiruInstance.(*HiruInstance)

                name := vm.GetNameAt(instruction.argument)

                vm.DebugPrint("Instruction SATTR: name '%v' to %v", name, hiruValue.Inspect())

                instance.StackFrame.StackFrame().Define(name.value, hiruValue)
                //vm.objectStack.Push(hiruInstance)


        case LVARS:
                vm.DebugPrint("Instruction LVARS")

                sf, _ := vm.callStack.Pop()
                vm.objectStack.Push(&HiruStructureVars{sf})
                vm.callStack.Push(sf)
                vm.DebugPrint("LVARS: %v", sf)

        case BUILDS:
                vm.DebugPrint("Instruction BUILDS")

                hobject,  _ := vm.objectStack.Pop()
                hvars := hobject.(*HiruStructureVars)

                structureObject := new(HiruStructure)
                structureObject.StackFrame = hvars

                vm.objectStack.Push(structureObject)

        case INITS:
                vm.DebugPrint("Instruction INITS: %v args", instruction.argument)
                args := make([]HiruObject, instruction.argument)
                for i := 0; int32(i) < instruction.argument; i++ {
                        pop, _ := vm.objectStack.Pop()
                        args[i] = pop
                }

                // Este objeto debería ser una estructura
                hiruObject, _ := vm.objectStack.Pop()

                // Asserteamos que es una estructura y sacamos su StackFrame
                structureObject := hiruObject.(*HiruStructure)
                structureStackFrame := structureObject.StackFrame

                // Creamos el objeto instancia. Este quedará en TOS al final de
                // esta función.
                hiruInstance := new(HiruInstance)

                // Buscamos la función __new__ en el StackFrame de la instancia
                __new__Object, _ := structureStackFrame.StackFrame().ResolveName("__new__")
                __new__Function := __new__Object.(*HiruFunction)
                vm.DebugPrint("Found __new__. Resolved to %v", __new__Function.Inspect())

                // codeObject es el Objeto función correspondiente a __new__
                codeObject := (*__new__Function).RawObject()
                names := codeObject.nameSegment.Entries()
                vm.DebugPrint("__new__ entries: %v", names)

                sf := NewStackFrame("__new__", codeObject, vm.ip)
                sf.MakeLinkTo(structureStackFrame.StackFrame())
                vm.callStack.Push(sf)

                i := int32(0)
                for i < instruction.argument {
                        vm.DebugPrint("name %v => %v", names[i], args[instruction.argument - i - 1])
                        vm.callStack.Define(names[i].value, args[instruction.argument - int32(i) - 1])
                        i++
                }

                func_body := codeObject.bytecodeSegment

                sf, _ = vm.callStack.Pop()
                sf.MakeLinkTo(vm.callStack.GetBottomMost())
                hiruInstance.StackFrame = &HiruStructureVars{sf}

                // TODO: Meter todo esto en un método (fn *HiruFunction).Call(args ...interface{})
                vm.callStack.Push(hiruInstance.StackFrame.StackFrame())

                vm.callStack.Define("__self__", hiruInstance)

                vm.ip = 0
                vm.RunOBIBytecode(func_body)
                return

        case IMPORT:
                // No es necesario para OBI
        case LNAME:
                name := vm.GetNameAt(instruction.argument)
                object, err := vm.callStack.ResolveName(name.value)
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Hiru: Variable '%v' not defined", name.value)
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
