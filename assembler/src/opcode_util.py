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
    "RET",
    "MAKEFN",
]

literal_args = [
    "BLIST",
    "BSTR",
    "JUMPFWD",
    "PJMPT",
    "PJMPF",
    "JMPABS",
    "CALLFN",
]

index_args = [
    "SNAME",
    "LCONST",
    "LNAME",
    "IMPORT",
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
    print(header)

    count = 0
    while count < len(no_args):
        print(op_count(no_args[count], count))
        count += 1

    print(op_count(no_args_separator, count))
    print()
    count += 1

    literal_args_len = len(literal_args)
    max_index = literal_args_len + count
    while count  < max_index:
        print(op_count(literal_args[count % literal_args_len - 1], count))
        count += 1

    print(op_count(literal_args_separator, count))
    print()
    count += 1

    index_args_len = len(index_args)
    max_index = index_args_len + count
    while count  < max_index:
        print(op_count(index_args[max_index - count - 1], count))
        count += 1

    print("end")

    # from_string
    print("def opcode_from_string(str)\n  op = case str")
    for op in no_args:
        print(to_s_format(op))

    for op in literal_args:
        print(to_s_format(op))

    for op in index_args:
        print(to_s_format(op))

    print("       end")
    print("  op\nend")

    print("def opcode_str(op)\n  str = case op")
    for op in no_args:
        print(to_op_format(op))

    for op in literal_args:
        print(to_op_format(op))

    for op in index_args:
        print(to_op_format(op))

    print("       end")
    print("  str\nend")

if __name__ == "__main__":
    main()
