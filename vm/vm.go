package vm

import (
        "fmt"
        "errors"
        "os"
)

// Enum del modo de ejecución de la VM
type VmMode uint8
const (
        IndexBasedMode VmMode = iota
        ObjectBasedMode
)

// Opciones para la vm
type VmOptions struct {
        debug           bool
        mode          VmMode
}

// Constructor para las opciones
func NewVmOptions(debug bool, mode VmMode) *VmOptions {
    return &VmOptions{debug: debug, mode: mode}
}

type HiruVM struct {
        mainFile     *HiruFile

        vmOptions    VmOptions

        callStack    CallStack

        objectStack  ObjectStack

        // Para OBI
        mainObject   *CodeObject
        tempObject   *CodeObject // Reemplazar esto por un stack o meterlo directamente en cada stack frame.
        isReturn     bool
}

func NewVm(filepath string, options VmOptions) (*HiruVM, error) {
        vm := new(HiruVM)
        hirufile, err := NewHiruFile(filepath)
        if err != nil {
                return nil, fmt.Errorf("Could not open file '%s'.", filepath)
        }

        // Hiru magic number. 0x48495255 == HIRU
        if hirufile.Read4Bytes() != 0x48495255 {
                return nil, errors.New("Wrong magic number, not a Hiru bytecode file.")
        }

        vm.mainFile = hirufile
        vm.vmOptions = options

        return vm, nil
}

func (vm *HiruVM) DebugPrint(s string, params ...interface{}) {
        switch {
        case !vm.vmOptions.debug:
                return
        case len(params) == 0:
                fmt.Print("[DEBUG] ")
                fmt.Fprint(os.Stdout, s)
                fmt.Println()
        default:
                fmt.Print("[DEBUG] ")
                fmt.Fprintf(os.Stdout, s, params...)
                fmt.Println()
        }
}

func (vm *HiruVM) DebugPrintCallStack() {
        vm.callStack.PrettyPrint()
}
func (vm *HiruVM) Run() (error) {
        vm.DebugPrint("VM.Run() called")

        // Empujamos el stackframe global
        sf := NewStackFrame("global")
        vm.callStack.Push(sf)

        if vm.vmOptions.mode == ObjectBasedMode {
                return vm.runObjectBasedVm()
        } else {
                return vm.runIndexBasedVm()
        }
}
// TODO
func (vm *HiruVM) runIndexBasedVm() (error) {
        vm.DebugPrint("Running in IBI Mode")
        return nil
}
