class String
  def is_number?
    true if Float(self) rescue false
  end
end

module TokenKind
  OP_AT             = 0x01
  OP_HASH           = 0x02
  OP_OPEN_BRACKETS  = 0x03
  OP_CLOSE_BRACKETS = 0x04
  OP_DOT            = 0x05
  OP_COLON          = 0x06
  OP_NEG            = 0x07

  LITERAL_STRING = 0x08
  LITERAL_FLOAT  = 0x09
  LITERAL_INT    = 0x0A

  IDENTIFIER = 0x0B

  KEY_HIRU = 0x0C
  KEY_SEGMENT = 0x0D
  KEY_DATA    = 0x0E
  KEY_NAMES   = 0x0F
  KEY_CODE    = 0x10

  UNKNOWN = 0x11
  EOF = 0x12
end

def token_string(tok)
  s = case tok
      when TokenKind::EOF
        "EOF"
      when TokenKind::IDENTIFIER
        "IDENTIFIER"
      when TokenKind::LITERAL_INT
        "LITERAL_INT"
      when TokenKind::LITERAL_FLOAT
        "LITERAL_FLOAT"
      when TokenKind::LITERAL_STRING
        "LITERAL_STRING"
      when TokenKind::OP_AT
        "@"
      when TokenKind::OP_DOT
        "."
      when TokenKind::OP_COLON
        ":"
      when TokenKind::OP_OPEN_BRACKETS
        "["
      when TokenKind::OP_CLOSE_BRACKETS
        "]"
      when TokenKind::KEY_HIRU
        "HIRU"
      when TokenKind::KEY_SEGMENT
        "SEGMENT"
      when TokenKind::KEY_DATA
        ".data"
      when TokenKind::KEY_NAMES
        ".names"
      when TokenKind::KEY_CODE
        ".code"
      end
  s # "s" xd
end

Token = Struct.new(:type, :value) do
  def to_str
    "Token(#{:type}, #{:value}"
  end
end

# Tampoco...
def is_whitespace(char)
  char =~ /\s/
end

# Los carácteres válidos para un identificador son [A-Za-z0-9_]
def is_valid_id_char(char)
  !char.match(/\A[a-zA-Z0-9_]*\z/).nil?
end

# Los carácteres válidos para el INICIO de un identificador son [A-Za-z_]
def is_valid_first_id_char(char)
  !char.match(/\A[a-zA-Z_]*\z/).nil?
end

$keywords = {
  "HIRU" => TokenKind::KEY_HIRU,
  "SEGMENT" => TokenKind::KEY_SEGMENT,
  "data" => TokenKind::KEY_DATA,
  "names" => TokenKind::KEY_NAMES,
  "code" =>  TokenKind::KEY_CODE,
}

def is_keyword(word)
  $keywords.key?(word)
end

class Tokenizer
  attr_accessor :tokens
  def initialize(source)
    @source = source

    @tokens = Array.new

    # Número de línea actual
    @line_number = 1

    # Número de columna en la línea actual
    @column_number = 0

    # Linea completa actual
    @current_line = String.new

    # Caracter actual
    @current_char = String.new

    # Caracter siguiente
    @peek_char = @source[@column_number]
  end

  def addToken(tt, value=nil)
    @tokens << Token.new(
      tt,
      value
    )
  end

  def addSyntaxError(desc)
    puts desc
    exit(-1)
  end

  def pushChar
    @current_char = @peek_char

    # El peek char se encuentra en el índice actua +1
    @column_number += 1
    @peek_char = @source[@column_number]

    # Si ya no hay más carácteres, empezamo
    @current_line += @current_char? @current_char : ""

    if @current_char == "\n"
      @line_number += 1
    end
    @current_char
  end

  def readUntil(char)
    res = String.new
    while @current_char != char
      res += @current_char
      pushChar
    end
    res
  end

  def check_operator
    t = case @current_char
        when '@'
          TokenKind::OP_AT
        when '#'
          TokenKind::OP_HASH
        when '.'
          TokenKind::OP_DOT
        when ':'
          TokenKind::OP_COLON
        when '-'
          TokenKind::OP_NEG
        when '['
          TokenKind::OP_OPEN_BRACKETS
        when ']'
          TokenKind::OP_CLOSE_BRACKETS
        else
          return false
        end
    addToken(t)
    true
  end

  def checkComment
    if @current_char == "("
      readUntil(")")
      return true
    else
      return false
    end
  end

  def checkLiteral
    lit = String.new
    if @current_char == '"'
      tt = TokenKind::LITERAL_STRING
      pushChar

      while @current_char != '"'
        lit += @current_char
        if @peek_char == "\n"
          addSyntaxError("Unexpected EOL reading string literal")
          pushChar
          return
        end
        pushChar
      end
      addToken(tt, lit)
      return true
    elsif @current_char.is_number?
      tt = -1
      lit += @current_char
      while @peek_char.is_number?
        pushChar
        lit += @current_char
      end

      if @peek_char == "."
        lit += @peek_char
        pushChar

        if not @peek_char.is_number?
          addSyntaxError("Malformed floating point number.")
          return true
        end

        while @peek_char.is_number?
          pushChar
          lit += @current_char
        end
        tt = TokenKind::LITERAL_FLOAT       
      else
        tt = TokenKind::LITERAL_INT
      end
      addToken(tt, lit)
      return true
    end
    false
  end

  def checkWord
    word = String.new
    if is_valid_first_id_char(@current_char)
      tt = -1
      word += @current_char
      while is_valid_id_char(@peek_char)
        pushChar
        word += @current_char
      end

      if is_keyword(word)
        tt = $keywords[word]
      else
        # En caso contrario es un identificador normal
        tt = TokenKind::IDENTIFIER
      end

      addToken(tt, word)
      return true
    end
    false
  end

  def tokenize!
    while pushChar
      if checkWord or check_operator or checkLiteral or checkComment
        next
      else
        if not is_whitespace(@current_char)
          error = "Unknown char '#{@current_char}'."
          addSyntaxError(error)
        end
        next
      end
    end

    # El último token del arreglo siempre será EOF
    addToken(TokenKind::EOF)
  end
end

