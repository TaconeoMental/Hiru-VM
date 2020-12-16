package vm

import (
        "fmt"
        "errors"
)

type HiruVM struct {
        mainFile *HiruFile
        debug bool

        //callStack CallStack
}

func NewVm(filepath string, debug bool) (*HiruVM, error) {
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
        vm.debug = debug
        return vm, nil
}

func (vm *HiruVM) Run() (error) {
        return nil
}
