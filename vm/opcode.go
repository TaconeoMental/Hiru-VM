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
        LCTXT
        LVARS
        LSELF
        BUILDS
        PRINT
        NO_ARGS

        BSTR
        JUMPFWD
        PJMPT
        PJMPF
        JMPABS
        CALLFN
        SLOOP
        INITS
        BLIST
        LITERAL_ARGS

        SNAME
        LCONST
        LNAME
        IMPORT
        LATTR
        SATTR
)

func opcodeString(op Opcode) string {
        switch op {
        case POP:
                return "pop"
        case UPOS:
                return "upos"
        case UNEG:
                return "uneg"
        case UNOT:
                return "unot"
        case BPOW:
                return "bpow"
        case BMUL:
                return "bmul"
        case BDIV:
                return "bdiv"
        case BMOD:
                return "bmod"
        case BSUB:
                return "bsub"
        case BADD:
                return "badd"
        case BAND:
                return "band"
        case BOR:
                return "bor"
        case CMPLT:
                return "cmplt"
        case CMPLE:
                return "cmple"
        case CMPEQ:
                return "cmpeq"
        case CMPNE:
                return "cmpne"
        case CMPGT:
                return "cmpgt"
        case CMPGE:
                return "cmpge"
        case RET:
                return "ret"
        case MAKEFN:
                return "makefn"
        case MAKEMOD:
                return "makemod"
        case EXIT:
                return "exit"
        case PLOOP:
                return "ploop"
        case LCTXT:
                return "lctxt"
        case BUILDS:
                return "builds"
        case PRINT:
                return "print"
        case NO_ARGS:
                return "NO ARGS"
        case BSTR:
                return "bstr"
        case JUMPFWD:
                return "jumpfwd"
        case PJMPT:
                return "pjmpt"
        case PJMPF:
                return "pjmpf"
        case JMPABS:
                return "jmpabs"
        case CALLFN:
                return "callfn"
        case SLOOP:
                return "sloop"
        case INITS:
                return "inits"
        case BLIST:
                return "blist"
        case LITERAL_ARGS:
                return "LITERAL ARGS"
        case SNAME:
                return "sname"
        case LCONST:
                return "lconst"
        case LNAME:
                return "lname"
        case IMPORT:
                return "import"
        case LATTR:
                return "lattr"
        case SATTR:
                return "sattr"
        }
        return "never here"
}

func HasArgs(op Opcode) bool {
        return op > NO_ARGS
}
