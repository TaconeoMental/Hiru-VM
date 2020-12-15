package vm

type HiruVM struct {
        mainFile HiruFile
        debug bool

        //callStack CallStack
}

func NewVm(filepath string, debug bool) (*HiruVM, error) {
        vm := &HiruVM{}
        return vm, nil
}
