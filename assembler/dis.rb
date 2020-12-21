#!/usr/bin/env ruby

require_relative "src/opcode"

$HIRU_MAGIC = 0x48495255

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
      when 0x69
        num = readNum
        puts "#{indent(level + 1)}Number #{length} #{num}"
      when 0x63
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

    instructions = Array.new
    jump_positions = Array.new
    n_opcodes.times do
      opcode = readNum
      arg = readNum

      if $jumps.include? opcode
        jump_positions.push(arg / (2 * 4))
      end

      if arg == 0x6E
        arg = ""
      end
      instructions.push([opcode, arg])
    end

    count = 0
    instructions.each do |opcode, arg|
      if jump_positions.include? count
        puts "\n#{indent(level)}   >> #{opcode_str(opcode)} #{arg}"
      else
        puts "#{indent(level + 1)}#{opcode_str(opcode)} #{arg}"
      end
      count += 1
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
