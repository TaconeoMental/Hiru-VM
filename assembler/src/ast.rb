$AST_INDENTATION = 4
$AST_LAST = "└" + "─" * ($AST_INDENTATION - 1)
$AST_MIDDLE = "├" + "─" * ($AST_INDENTATION - 1)
$AST_LINE = "│" + " " * ($AST_INDENTATION - 1)
$AST_SPACE = " " * $AST_INDENTATION

def indentation(last)
  if last
    print $AST_LAST
    return $AST_SPACE
  else
    print $AST_MIDDLE
    return $AST_LINE
  end
end

module Ast
  class ProgramNode
    attr_accessor :main_function
    def initialize
      @main_function = nil
    end

    def print_tree
      puts "Program"
      @main_function.print_tree("", true)
    end

    def emitBytecode(file)
      @main_function.emitBytecode(file)
    end
  end

  class DataSegmentNode
    def initialize
      # [IndexNode] = Literal
      @data_hash = Hash.new
    end

    def add(key, val)
      @data_hash[key] = val
    end

    def dataLength
      @data_hash.length
    end

    def data
      @data_hash
    end

    # Develve el largo del diccionario expresado en bytecode
    def byte_length
      len = 4 # por el header
      @data_hash.each do |key, val|
        len += 4 # Tipo de dato
        len += 4 # Largo del dato

        case val
        when IntLiteral
          len += 4
        when StringLiteral
          len += val.Value.bytesize
        when FunctionNode
          len += val.byte_length
        end
      end
      len
    end

    def print_tree(indent, last)
      print indent
      indent += indentation(last)

      puts "DataSegment"

      puts indent + $AST_LAST + "Data"
      @data_hash.each do |key, val|
        puts indent + $AST_SPACE + "#{$AST_SPACE}#{key.value}"
        val.print_tree(indent + $AST_SPACE + $AST_SPACE, true)
      end
    end
  end

  class NameSegmentNode
    def initialize
      @names_hash = Hash.new
    end

    def add(key, val)
      @names_hash[key] = val
    end

    def namesLength
      @names_hash.length
    end

    def names
      @names_hash
    end

    def print_tree(indent, last)
      print indent
      indent += indentation(last)

      puts "NameSegment"

      puts indent + $AST_LAST + "Names"
      @names_hash.each do |key, val|
        puts indent + $AST_SPACE + "#{$AST_SPACE}#{key.value}"
        val.print_tree(indent + $AST_SPACE + $AST_SPACE, true)
      end
    end

    def byte_length
      len = 4
      @names_hash.each do |key, val|
        len += 4
        len += val.byte_length
      end
      len
    end
  end

  class CodeSegmentNode
    def initialize
      @code_hash = Array.new
    end

    def add(key, val)
      @code_hash.push([key, val])
    end

    def codeLength
      @code_hash.length
    end

    def code
      @code_hash
    end

    def print_tree(indent, last)
      print indent
      indent += indentation(last)

      puts "CodeSegment"

      puts indent + $AST_LAST + "Opcodes"
      @code_hash.each do |key, val|
        puts indent + $AST_SPACE + "#{$AST_SPACE}#{opcode_str(key)}"
        val.print_tree(indent + $AST_SPACE + $AST_SPACE, true)
      end
    end

    def byte_length
      len = 4
      @code_hash.each do |key, val|
        len += 8
      end
      len
    end
  end

  class FunctionNode
    attr_accessor :identifier, :data_segment, :name_segment, :code_segment
    def initialize
      @identifier = nil
      @data_segment = nil
      @name_segment = nil
      @code_segment = nil
      @index_segment = nil
    end

    def print_tree(indent, last)
      print indent
      indent += indentation(last)

      puts "FunctionNode"
      puts indent + $AST_MIDDLE + "Identifier"
      @identifier.print_tree(indent + $AST_LINE, true)

      puts indent + $AST_MIDDLE + "DataSegment"
      @data_segment.print_tree(indent + $AST_LINE, true)

      puts indent + $AST_MIDDLE + "NameSegment"
      @name_segment.print_tree(indent + $AST_LINE, true)

      puts indent + $AST_LAST + "CodeSegment"
      @code_segment.print_tree(indent + $AST_SPACE, true)
    end

    def byte_length
      len = 13 * 4
      len += @data_segment.byte_length
      len += @name_segment.byte_length
      len += @code_segment.byte_length
      len
    end
  end

  class IntLiteral
    attr_accessor :Value
    def initialize(val=nil)
      @Value = val.to_i
    end

    def print_tree(indent, last)
      print indent
      indent += indentation(last)

      puts "IntLiteral"
      puts indent + $AST_LAST + @Value.to_s
    end
  end

  class StringLiteral
    attr_accessor :Value
    def initialize(val=nil)
      @Value = val
    end

    def emitBytecode(file)
      file.write(@Value, @Value.bytesize)
    end

    def byte_length
      @Value.bytesize
    end

    def print_tree(indent, last)
      print indent
      indent += indentation(last)

      puts "StringLiteral"
      puts indent + $AST_LAST + @Value
    end
  end

  class IndexNode
    attr_accessor :value
    def initialize(num)
      @value = num
    end

    def eval
      @value
    end

    def print_tree(indent, last)
      print indent
      indent += indentation(last)

      puts "Index"
      puts indent + $AST_LAST + @value.to_s
    end
  end

  class Identifier
    attr_accessor :Value
    def initialize(value)
      @Value = value
    end

    def byte_length
      @Value.bytesize
    end

    def print_tree(indent, last)
      print indent
      indent += indentation(last)

      puts "Identifier"
      puts indent + $AST_LAST + @Value
    end
  end

  class NullLiteral
    def initialize
    end

    def print_tree(indent, last)
      print indent
      indent += indentation(last)

      puts "NullLiteral"
      puts indent + $AST_LAST + "Null"
    end
  end
end
