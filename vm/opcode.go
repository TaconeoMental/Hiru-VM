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
        RET
        MAKEFN
        MAKEMOD
        BLIST
        BSTR
        JUMPFWD
        PJMPT
        PJMPF
        JMPABS
        CALLFN
        SNAME
        LCONST
        LNAME
        IMPORT
        LATTR
)
