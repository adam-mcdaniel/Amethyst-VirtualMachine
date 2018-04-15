package main

import (
	"./parser"
	"./virtualmachine"
	"./library"
	"fmt"
)

func PrintList(tokens []string) {
	fmt.Print("[")
	for i, token := range tokens {
		fmt.Print("'")
		fmt.Print(token)
		fmt.Print("'")
		if i < len(tokens)-1 {
			fmt.Println(", ")
		}
	}
	fmt.Println("]")
}


func main() {
	// parser := parser.MakeParser("1 a < a > | '\n' |")
	// parser := parser.MakeParser("9 a < a > 1 + | '\n' |")
	// parser := parser.MakeParser("{param1 < @ this < param1 > this > x , this >} cons < 2 cons > ! MyObject < MyObject > | '\n' |")
	// parser := parser.MakeParser("7 b < b a < a > > | '\n' |")
	// parser := parser.MakeParser("'# ' | $ input < 'you said: ' | input > | '\n' | ")
	// parser := parser.MakeParser("1 'hello world!' 'your mom: ' | | '\n' | | '\n' |")

	// parser := parser.MakeParser("{i < i > factorial < factorial > | ' factorial is: ' | {1 i > - i < i > factorial > * factorial <} {1 i > >>} & factorial > |} factorial < 7 factorial > ! '\n' |")
	
	// parser := parser.MakeParser("{condition < function < function > ! {function > !} {condition > !} &} dowhile < " +
	// 														"{1 |} {0} dowhile > ! " +
	// 														"'\n' |")
	
	// parser := parser.MakeParser("{condition < function < {function > ! {0} condition <} {condition > !} &} if < {1 |} {1} if > ! '\n' |")

	// Making a game is possible using amethyst!
	parser := parser.MakeParser("240 320 'Adam' make > ! 'It works!!!\n' |")
	tokens := parser.Parse()
	// PrintList(tokens)
	VM := virtualmachine.MakeVM(tokens, library.GetLibrary())
	library.InstallLibrary(&VM)
	VM.Run()

}
