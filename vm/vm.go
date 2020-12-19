package vm

import (
        "fmt"
        "errors"
)

type VmMode uint8
const (
        IndexBasedMode VmMode = iota
        ObjectBasedMode
)

type VmOptions struct {
        debug           bool
        mode          VmMode
}

func NewVmOptions(debug bool, mode VmMode) *VmOptions {
    return &VmOptions{debug: debug, mode: mode}
}

type HiruVM struct {
        mainFile     *HiruFile

        vmOptions    VmOptions

        callStack    CallStack
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

func (vm *HiruVM) Run() (error) {
        if vm.vmOptions.mode == ObjectBasedMode {
                return vm.runObjectBasedVm()
        } else {
                return vm.runIndexBasedVm()
        }
}

// TODO
func (vm *HiruVM) runObjectBasedVm() (error) {
        return nil
}

// TODO
func (vm *HiruVM) runIndexBasedVm() (error) {
        return nil
}
