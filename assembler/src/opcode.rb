module Opcodes
  POP = 0x00
  UPOS = 0x01
  UNEG = 0x02
  UNOT = 0x03
  BPOW = 0x04
  BMUL = 0x05
  BDIV = 0x06
  BMOD = 0x07
  BSUB = 0x08
  BADD = 0x09
  BAND = 0x0a
  BOR = 0x0b
  CMPLT = 0x0c
  CMPLE = 0x0d
  CMPEQ = 0x0e
  CMPNE = 0x0f
  CMPGT = 0x10
  CMPGE = 0x11
  RET = 0x12
  MAKEFN = 0x13
  MAKEMOD = 0x14
  EXIT = 0x15
  PLOOP = 0x16
  LCTXT = 0x17
  LVARS = 0x18
  LSELF = 0x19
  BUILDS = 0x1A
  PRINT = 0x1b
  NO_ARGS = 0x1c

  BSTR = 0x1d
  JUMPFWD = 0x1e
  PJMPT = 0x1f
  PJMPF = 0x20
  JMPABS = 0x21
  CALLFN = 0x22
  SLOOP = 0x23
  INITS = 0x24
  BLIST = 0x25
  LITERAL_ARGS = 0x26

  SNAME = 0x27
  LCONST = 0x28
  LNAME = 0x29
  IMPORT = 0x2A
  LATTR = 0x2B
  SATTR = 0x2C
end

def opcode_from_string(str)
  op = case str
       when "pop"
         Opcodes::POP
       when "upos"
         Opcodes::UPOS
       when "uneg"
         Opcodes::UNEG
       when "unot"
         Opcodes::UNOT
       when "bpow"
         Opcodes::BPOW
       when "bmul"
         Opcodes::BMUL
       when "bdiv"
         Opcodes::BDIV
       when "bmod"
         Opcodes::BMOD
       when "bsub"
         Opcodes::BSUB
       when "badd"
         Opcodes::BADD
       when "band"
         Opcodes::BAND
       when "bor"
         Opcodes::BOR
       when "cmplt"
         Opcodes::CMPLT
       when "cmple"
         Opcodes::CMPLE
       when "cmpeq"
         Opcodes::CMPEQ
       when "cmpne"
         Opcodes::CMPNE
       when "cmpgt"
         Opcodes::CMPGT
       when "cmpge"
         Opcodes::CMPGE
       when "ret"
         Opcodes::RET
       when "makefn"
         Opcodes::MAKEFN
       when "makemod"
         Opcodes::MAKEMOD
       when "exit"
         Opcodes::EXIT
       when "ploop"
         Opcodes::PLOOP
       when "lctxt"
         Opcodes::LCTXT
       when "lvars"
         Opcodes::LVARS
       when "lself"
         Opcodes::LSELF
       when "builds"
         Opcodes::BUILDS
       when "print"
         Opcodes::PRINT
       when "blist"
         Opcodes::BLIST
       when "bstr"
         Opcodes::BSTR
       when "jumpfwd"
         Opcodes::JUMPFWD
       when "pjmpt"
         Opcodes::PJMPT
       when "pjmpf"
         Opcodes::PJMPF
       when "jmpabs"
         Opcodes::JMPABS
       when "callfn"
         Opcodes::CALLFN
       when "sloop"
         Opcodes::SLOOP
       when "inits"
         Opcodes::INITS
       when "sname"
         Opcodes::SNAME
       when "lconst"
         Opcodes::LCONST
       when "lname"
         Opcodes::LNAME
       when "import"
         Opcodes::IMPORT
       when "lattr"
         Opcodes::LATTR
       when "sattr"
         Opcodes::SATTR
       end
  op
end

def opcode_str(op)
  str = case op
       when Opcodes::POP
         "pop"
       when Opcodes::UPOS
         "upos"
       when Opcodes::UNEG
         "uneg"
       when Opcodes::UNOT
         "unot"
       when Opcodes::BPOW
         "bpow"
       when Opcodes::BMUL
         "bmul"
       when Opcodes::BDIV
         "bdiv"
       when Opcodes::BMOD
         "bmod"
       when Opcodes::BSUB
         "bsub"
       when Opcodes::BADD
         "badd"
       when Opcodes::BAND
         "band"
       when Opcodes::BOR
         "bor"
       when Opcodes::CMPLT
         "cmplt"
       when Opcodes::CMPLE
         "cmple"
       when Opcodes::CMPEQ
         "cmpeq"
       when Opcodes::CMPNE
         "cmpne"
       when Opcodes::CMPGT
         "cmpgt"
       when Opcodes::CMPGE
         "cmpge"
       when Opcodes::RET
         "ret"
       when Opcodes::MAKEFN
         "makefn"
       when Opcodes::MAKEMOD
         "makemod"
       when Opcodes::EXIT
         "exit"
       when Opcodes::PLOOP
         "ploop"
       when Opcodes::LCTXT
         "lctxt"
       when Opcodes::LVARS
         "lvars"
       when Opcodes::LSELF
         "lself"
       when Opcodes::BUILDS
         "builds"
       when Opcodes::PRINT
         "print"
       when Opcodes::BLIST
         "blist"
       when Opcodes::BSTR
         "bstr"
       when Opcodes::JUMPFWD
         "jumpfwd"
       when Opcodes::PJMPT
         "pjmpt"
       when Opcodes::PJMPF
         "pjmpf"
       when Opcodes::JMPABS
         "jmpabs"
       when Opcodes::CALLFN
         "callfn"
       when Opcodes::SLOOP
         "sloop"
       when Opcodes::INITS
         "inits"
       when Opcodes::SNAME
         "sname"
       when Opcodes::LCONST
         "lconst"
       when Opcodes::LNAME
         "lname"
       when Opcodes::IMPORT
         "import"
       when Opcodes::LATTR
         "lattr"
       when Opcodes::SATTR
         "sattr"
       end
  str
end

def has_arg(op)
  return op > Opcodes::NO_ARGS
end

def index_arg?(op)
  return op > Opcodes::LITERAL_ARGS
end

def literal_arg?(op)
  return op < Opcodes::LITERAL_ARGS && op > Opcodes::NO_ARGS
end

$jumps = [Opcodes::PJMPT, Opcodes::PJMPF, Opcodes::JMPABS, Opcodes::JUMPFWD, Opcodes::SLOOP]
