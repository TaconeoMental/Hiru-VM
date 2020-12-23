package vm

type Opcode uint8
const (
        POP Opcode = iota
        UPOS
        UNEG
        UNOT
        BPOW
        BMUL
        BDIV
        BMOD
        BSUB
        BADD
        BAND
        BOR
        CMPLT
        CMPLE
        CMPEQ
        CMPNE
        CMPGT
        CMPGE
        RET
        MAKEFN
        MAKEMOD
        EXIT
        PLOOP
        PRINT
        NO_ARGS

        BLIST
        BSTR
        JUMPFWD
        PJMPT
        PJMPF
        JMPABS
        CALLFN
        SLOOP
        LITERAL_ARGS

        LATTR
        IMPORT
        LNAME
        LCONST
        SNAME
)

func HasArgs(op Opcode) bool {
        return op > NO_ARGS
}
