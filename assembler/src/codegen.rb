require_relative 'ast'

class IndexSegment
  attr_accessor :number_entries
  attr_accessor :index_segment_start, :index_segment_length
  attr_accessor :data_segment_start, :data_segment_length
  attr_accessor :name_segment_start, :name_segment_length
  attr_accessor :code_segment_start, :code_segment_length
  def initialize
    @number_entries = 4

    @index_segment_start = 0
    @index_segment_length = 0

    @data_segment_start = 0
    @data_segment_length = 0

    @name_segment_start = 0
    @name_segment_length = 0

    @code_segment_start = 0
    @code_segment_length = 0
  end

  def set_index_segment(start, len)
    @index_segment_start = start
    @index_segment_length = len
  end

  def set_data_segment(start, len)
    @data_segment_start = start
    @data_segment_length = len
  end

  def set_name_segment(start, len)
    @name_segment_start = start
    @name_segment_length = len
  end

  def set_code_segment(start, len)
    @code_segment_start = start
    @code_segment_length = len
  end
end


class BytecodeFile
  def initialize(filename)
    @hbc_file = File.open(filename, "wb")
  end

  def write_byte(num)
    buffer = [num].pack("C>")
    @hbc_file.write(buffer)
  end

  def write(data)
    @hbc_file.write(data)
  end

  def write_4_bytes(num)
    buffer = [num].pack("L>")
    @hbc_file.write(buffer)
  end

  def close
    @hbc_file.close
  end
end

class CodeGenerator
  def initialize(ast, filename)
    @ast = ast
    @hbc_file = BytecodeFile.new(filename + ".hbc")

    @block = 4 # Bytes

    @current_byte_index = 1

    @prefix_bytes = 0
  end

  def pushIndex(num)
    @current_byte_index += num
  end

  def createIndexSegment(node)
    index_segment = IndexSegment.new
    #self_start = @last_index_end
    self_start =  @prefix_bytes
    self_len = 1 * @block + 12 * @block
    @current_byte_index = @prefix_bytes + self_len
    #pushIndex(self_len)

    data_start = @current_byte_index
    data_len = node.data_segment.byte_length
    pushIndex(data_len)

    name_start = @current_byte_index
    name_len = node.name_segment.byte_length
    pushIndex(name_len)

    code_start = @current_byte_index
    code_len = node.code_segment.byte_length
    pushIndex(code_len)

    index_segment.set_data_segment(data_start, data_len)
    index_segment.set_index_segment(self_start, self_len)
    index_segment.set_name_segment(name_start, name_len)
    index_segment.set_code_segment(code_start, code_len)
    index_segment
  end

  def dataSegmentTypeCode(node)
    case node
    when Ast::IntLiteral
      return 0x6E
    when Ast::StringLiteral
      return 0x73
    when Ast::FunctionNode
      return 0x66
    end
  end

  def emitBytecode(node)
    case node
    when Ast::ProgramNode
      # HIRU Magic number
      @hbc_file.write_4_bytes(0x48495255)
      pushIndex(4)
      @prefix_bytes += 4

      emitBytecode(node.main_function)
      return
    when Ast::FunctionNode
      index_segment = createIndexSegment(node)
      emitBytecode(index_segment)

      emitBytecode(node.data_segment)
      emitBytecode(node.name_segment)
      emitBytecode(node.code_segment)
      return

    when IndexSegment
      @hbc_file.write_4_bytes(node.number_entries)
      pushIndex(4)

      @prefix_bytes += 4

      # Index table
      @hbc_file.write_4_bytes(0x00)
      @hbc_file.write_4_bytes(node.index_segment_start)
      @hbc_file.write_4_bytes(node.index_segment_length)
      pushIndex(12)
      @prefix_bytes += 12

      # Data segment
      @hbc_file.write_4_bytes(0x01)
      @hbc_file.write_4_bytes(node.data_segment_start)
      @hbc_file.write_4_bytes(node.data_segment_length)
      pushIndex(12)
      @prefix_bytes += 12

      # Name Segment
      @hbc_file.write_4_bytes(0x02)
      @hbc_file.write_4_bytes(node.name_segment_start)
      @hbc_file.write_4_bytes(node.name_segment_length)
      pushIndex(12)
      @prefix_bytes += 12

      @hbc_file.write_4_bytes(0x03)
      @hbc_file.write_4_bytes(node.code_segment_start)
      @hbc_file.write_4_bytes(node.code_segment_length)
      pushIndex(12)
      @prefix_bytes += 12
      return

    when Ast::DataSegmentNode
      @hbc_file.write_4_bytes(node.dataLength)
      pushIndex(4)
      @prefix_bytes += 4

      node.data.each do |key, val|
        @hbc_file.write_4_bytes(dataSegmentTypeCode(val))
        pushIndex(4)
        @prefix_bytes += 4

        case val
        when Ast::StringLiteral
          len = val.Value.bytesize
        when Ast::IntLiteral
          len = 4
        when Ast::FunctionNode
          len = val.byte_length
        end

        @hbc_file.write_4_bytes(len)
        pushIndex(4)
        @prefix_bytes += 4

        emitBytecode(val)
      end
      return

    when Ast::NameSegmentNode
      @hbc_file.write_4_bytes(node.namesLength)
      pushIndex(4)

      node.names.each do |key, val|
        bytesize = val.byte_length
        @hbc_file.write_4_bytes(bytesize)
        pushIndex(4)

        @hbc_file.write(val.Value)
        pushIndex(bytesize)
      end
      return

    when Ast::CodeSegmentNode
      @hbc_file.write_4_bytes(node.codeLength)
      pushIndex(4)

      node.code.each do |key, val|
        @hbc_file.write_4_bytes(key)
        pushIndex(4)

        if val.is_a? Ast::NullLiteral
          @hbc_file.write_4_bytes(0x6e) # 0x6E == Null
        else
          @hbc_file.write_4_bytes(val.value.to_i)
        end
        pushIndex(4)
      end
      return

    when Ast::StringLiteral
      @hbc_file.write(node.Value)
      size = node.Value.bytesize
      pushIndex(size)
      @prefix_bytes += size
      return

    when Ast::IntLiteral
      @hbc_file.write_4_bytes(node.Value)
      pushIndex(4)
      @prefix_bytes += 4
      return
    end
    puts "Unreacheable.\nNode: #{node}"
  end

  def run
    emitBytecode(@ast)
    @hbc_file.close
  end
end
