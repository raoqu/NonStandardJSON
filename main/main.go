package main

import (
	"NJson/parser"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// 从文件或其他源读取JSON文本
	if len(os.Args) != 2 {
		println("Usage: main <jsonFile>")
		os.Exit(1)
	}

	input, _ := antlr.NewFileStream(os.Args[1])

	// 创建一个词法分析器和语法分析器
	lexer := parser.NewNJsonLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewNJsonParser(stream)

	// 创建一个监听器并遍历解析树
	tree := p.Json()
	println("------------------------")
	antlr.ParseTreeWalkerDefault.Walk(NewNJsonFormatListener(), tree)
}
