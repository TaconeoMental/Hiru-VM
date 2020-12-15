require_relative 'src/codegen'
require_relative 'src/tokenizer'
require_relative 'src/parser'

def main
  return if ARGV.length < 1
  filename = ARGV.[](0)

  if not File.exists? filename
    puts "File '#{filename}' does not exist"
    exit(-1)
  end
  source = File.read(filename)

  tokenizer = Tokenizer.new(source)
  tokenizer.tokenize!

  parser = Parser.new(tokenizer.tokens)
  ast = parser.parseProgram

  code_generator = CodeGenerator.new(ast, File.basename(filename, ".asm"))
  code_generator.run
end

main

