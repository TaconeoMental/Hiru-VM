# File used to automatically generate part of the code of the assembler

no_args = [
    "POP",
    "UPOS",
    "UNEG",
    "UNOT",
    "BPOW",
    "BMUL",
    "BDIV",
    "BMOD",
    "BSUB",
    "BADD",
    "BAND",
    "BOR",
    "CMPLT",
    "CMPLE",
    "CMPEQ",
    "CMPNE",
    "CMPGT",
    "CMPGE",
    "RET",
    "MAKEFN",
    "MAKEMOD",
    "EXIT",
    "PLOOP",
    "LCTXT",
    "BUILDS",
    "PRINT",
]

literal_args = [
    "BLIST",
    "BSTR",
    "JUMPFWD",
    "PJMPT",
    "PJMPF",
    "JMPABS",
    "CALLFN",
    "SLOOP",
    "INITS",
]

index_args = [
    "SNAME",
    "LCONST",
    "LNAME",
    "IMPORT",
    "LATTR",
    "SATTR",
]

no_args_separator = "NO_ARGS"
literal_args_separator = "LITERAL_ARGS"

header = "module Opcodes"

def to_s_format(op):
    return f"       when \"{op.lower()}\"\n         Opcodes::{op}"

def to_op_format(string):
    return f"       when Opcodes::{string}\n         \"{string.lower()}\""

def op_count(op_type, count):
    return f"  {op_type} = 0x{count:02x}"

def main():
    print("RUBY CODE")
    print(header)

    last_index = 0
    for index, noarg in enumerate(no_args):
        print(op_count(noarg, index))
    last_index += len(no_args) + 1

    print(op_count(no_args_separator, last_index))
    print()

    for index, litarg in enumerate(literal_args):
        print(op_count(litarg, index + last_index + 1))
    last_index += len(literal_args) + 1

    print(op_count(literal_args_separator, last_index))
    print()

    for index, indexarg in enumerate(index_args):
        print(op_count(indexarg, index + last_index + 1))
    last_index += len(index_args)

    print("end\n")

    # from_string
    print("def opcode_from_string(str)\n  op = case str")
    for op in no_args:
        print(to_s_format(op))

    for op in literal_args:
        print(to_s_format(op))

    for op in index_args:
        print(to_s_format(op))

    print("       end")
    print("  op\nend\n")

    print("def opcode_str(op)\n  str = case op")
    for op in no_args:
        print(to_op_format(op))

    for op in literal_args:
        print(to_op_format(op))

    for op in index_args:
        print(to_op_format(op))

    print("       end")
    print("  str\nend")
    print("END RUBY CODE\n")
    print("START GO CODE")

    print("type Opcode uint8")
    print("const (")
    print(f"        {no_args[0]} Opcode = iota")
    for op in no_args[1:]:
        print(f"        {op}")
    print(f"        {no_args_separator}\n")

    literal_args[:] =  literal_args[-8:] + literal_args[:-8]

    for op in literal_args:
        print(f"        {op}")

    print(f"        {literal_args_separator}\n")

    for op in index_args[::-1]:
        print(f"        {op}")
    print(")")

    print("END GO CODE")

if __name__ == "__main__":
    main()
