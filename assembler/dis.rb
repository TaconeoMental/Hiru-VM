#!/usr/bin/env ruby

$HIRU_MAGIC = 0x48495255

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
  RET = 0x0c
  MAKEFN = 0x0d
  NO_ARGS = 0x0e

  BLIST = 0x0f
  BSTR = 0x10
  JUMPFWD = 0x11
  PJMPT = 0x12
  PJMPF = 0x13
  JMPABS = 0x14
  CALLFN = 0x15
  LITERAL_ARGS = 0x16

  IMPORT = 0x17
  LNAME = 0x18
  LCONST = 0x19
  SNAME = 0x1a
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
       when Opcodes::RET
         "ret"
       when Opcodes::MAKEFN
         "makefn"
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
       when Opcodes::SNAME
         "sname"
       when Opcodes::LCONST
         "lconst"
       when Opcodes::LNAME
         "lname"
       when Opcodes::IMPORT
         "import"
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

class Disassembler
  def initialize(source)
    @source = source
    if not read_compare($HIRU_MAGIC)
      addError("Incorrect magic number.")
    end
    puts "HIRU"
  end

  def read
    @source.read(4)
  end

  def readNum
    read.unpack("N")[0]
  end

  def read_compare(num)
    read == [num].pack("N")
  end

  def compare(num, num2)
    num == [num2].pack("N")
  end

  def addError(desc)
    puts desc
    exit(-1)
  end

  def index_entry_code(code)
    case code
    when 0
      return "index_table"
    when 1
      return "data_table"
    when 2
      return "name_table"
    when 3
      "code_table"
    end
  end

  def indent(level)
    "      " * level
  end

  def dissasembleFunction(level=0)
    # Primero leemos la cantidad de entradas en el index segment.
    entries = readNum
    puts "#{indent(level)}SEGMENT .index #{entries}"

    entries.times do
      type = readNum
      offset = readNum
      length = readNum

      puts "#{indent(level + 1)}#{index_entry_code(type)} #{offset} #{length}"
    end

    n_constants = readNum
    puts "#{indent(level)}SEGMENT .data #{n_constants}"

    n_constants.times do
      type = readNum
      length = readNum
      case type
      when 0x6E
        num = readNum
        puts "#{indent(level + 1)}Number #{length} #{num}"
      when 0x66
        puts "#{indent(level + 1)}Function #{length}"
        dissasembleFunction(level + 2)
      when 0x73
        data_bytecode = @source.read(length)
        puts "#{indent(level + 1)}String #{length} \"#{data_bytecode}\""
      end
    end

    n_names = readNum
    puts "#{indent(level)}SEGMENT .name #{n_names}"

    n_names.times do
      length = readNum
      name = @source.read(length)

      puts "#{indent(level + 1)}#{length} #{name}"
    end

    n_opcodes = readNum
    puts "#{indent(level)}SEGMENT .code #{n_opcodes}"

    n_opcodes.times do
      opcode = readNum
      arg = readNum

      if arg == 0x6E
        arg = ""
      end
      puts "#{indent(level + 1)}#{opcode_str(opcode)} #{arg}"
    end
  end

  def dissasemble
    dissasembleFunction
  end
end

def main
  return if ARGV.length < 1
  filename = ARGV.[](0)

  if not File.exists? filename
    puts "File '#{filename}' does not exist"
    exit(-1)
  end
  source = File.open(filename)
  dis = Disassembler.new(source)
  dis.dissasemble

  source.close
end

main
