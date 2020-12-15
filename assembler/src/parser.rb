require_relative 'ast'
require_relative 'opcode'


class Parser
    def initialize(tokenizer)
      @tokens = tokenizer.each
      @current_tok = nil
      @peek_token = nil
  
      pushToken
    end
  
    def addParseError(desc)
      puts desc
      exit(-1)
    end
  
    def pushToken
      @current_token = @peek_token
      begin
        @peek_token = @tokens.next
      rescue StopIteration => err
        @peek_token = nil
      end
    end
  
    def currentTokenEquals(tok)
      @current_token.type == tok
    end
  
    def peekTokenEquals(tok)
      @peek_token.type == tok
    end
  
    def consumePeek(token_kind)
      if peekTokenEquals(token_kind)
        pushToken
        return true
      end
      puts "Expected '#{token_string(token_kind)}', but got '#{token_string(@peek_token.type)}' instead."
      exit(-1)
    end
  
    def parseProgram
      pn = Ast::ProgramNode.new
      pn.main_function = parseFunctionLiteral(true)
      pn
    end
  
    def parseFunctionLiteral(main=false)
      consumePeek(TokenKind::OP_OPEN_BRACKETS)
      fl = Ast::FunctionNode.new
  
      if main
        consumePeek(TokenKind::KEY_HIRU)
      else
        consumePeek(TokenKind::IDENTIFIER)
      end
      fl.identifier = Ast::Identifier.new(@current_token.value)
  
      consumePeek(TokenKind::OP_CLOSE_BRACKETS)
  
      fl.data_segment = parseDataSegment
  
      fl.name_segment = parseNameSegment
  
      fl.code_segment = parseCodeSegment
      fl
    end
  
    def parseDataSegment
      consumePeek(TokenKind::KEY_SEGMENT)
      consumePeek(TokenKind::OP_DOT)
      consumePeek(TokenKind::KEY_DATA)
  
      ds = Ast::DataSegmentNode.new
  
      while peekTokenEquals(TokenKind::OP_AT)
        pushToken
        consumePeek(TokenKind::LITERAL_INT)
  
        index = Ast::IndexNode.new(@current_token.value)
        ds.add(index, parseLiteral)
      end
      ds
    end
  
    def parseNameSegment
      consumePeek(TokenKind::KEY_SEGMENT)
      consumePeek(TokenKind::OP_DOT)
      consumePeek(TokenKind::KEY_NAMES)
  
      ns = Ast::NameSegmentNode.new
  
      while peekTokenEquals(TokenKind::OP_AT)
        pushToken
        consumePeek(TokenKind::LITERAL_INT)
  
        index = Ast::IndexNode.new(@current_token.value)
        consumePeek(TokenKind::IDENTIFIER)
  
        ns.add(index, Ast::Identifier.new(@current_token.value))
      end
  
      ns
    end
  
  
    def parseLiteral
      case @peek_token.type
      when TokenKind::OP_OPEN_BRACKETS
        operand = parseFunctionLiteral
      when TokenKind::LITERAL_STRING
        pushToken
        operand = Ast::StringLiteral.new(@current_token.value)
      when TokenKind::LITERAL_INT
        pushToken
        operand = Ast::IntLiteral.new(@current_token.value)
      end
      operand
    end
  
    def parseCodeSegment
      consumePeek(TokenKind::KEY_SEGMENT)
      consumePeek(TokenKind::OP_DOT)
      consumePeek(TokenKind::KEY_CODE)
  
      cs = Ast::CodeSegmentNode.new
      while peekTokenEquals(TokenKind::IDENTIFIER)
        pushToken
        opcode = opcode_from_string(@current_token.value)
        if not opcode
          addParseError("Error: '#{@current_token.value}' is not an opcode.")
        end
  
        if not has_arg(opcode)
          cs.add(opcode, Ast::NullLiteral.new)
        elsif peekTokenEquals(TokenKind::OP_AT)
          pushToken
          if not index_arg?(opcode)
            addParseError("Error: Instruction '#{@current_token.value}' does not take and index as an argument.")
          end
  
          consumePeek(TokenKind::LITERAL_INT)
          index = Ast::IndexNode.new(@current_token.value)
          cs.add(opcode, index)
        elsif peekTokenEquals(TokenKind::LITERAL_INT)
          if not literal_arg?(opcode)
            addParseError("Error: Instruction '#{@current_token.value}' does not take a number as an argument.")
          end
          pushToken
          index = Ast::IndexNode.new(@current_token.value)
          cs.add(opcode, index)
        end
      end
      cs
    end
  end