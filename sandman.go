package sandman

import (
	"github.com/sbinet/go-python"
	"log"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func GetFunction(module_name string, function_name string) *python.PyObject {

	Module := python.PyImport_ImportModule(module_name)
	if Module == nil {
		log.Fatal("Failed to load the "+ module_name+" module")
	}

	var MethodDesired *python.PyObject
	if Module.HasAttrString(function_name) == 1 {
		MethodDesired = Module.GetAttrString(function_name)
	}
	if !MethodDesired.Check_Callable() {
		log.Fatal(module_name+" is not callable")
	}
	return MethodDesired

}
func Highlight(code string, lexer string, formatter string) string {

	// // var Highlighter *python.PyObject
	// // if PygmentsModule.HasAttrString("highlighter") == 1 {
	// //	Highlighter = FormatterModule.GetAttrString("highlighter")
	// // }

	GetFormatterByName := GetFunction("pygments.formatters", "get_formatter_by_name")
	FormatterArgs := python.PyTuple_New(1)
	python.PyTuple_SetItem(FormatterArgs, 0, python.PyString_FromString(formatter))
	Formatter:= GetFormatterByName.CallObject(FormatterArgs)

	if Formatter == nil {
		log.Fatal("Couldn't get formatter")
	}

	GetLexerByName := GetFunction("pygments.lexers", "get_lexer_by_name")
	LexerArgs := python.PyTuple_New(1)
	python.PyTuple_SetItem(LexerArgs, 0, python.PyString_FromString(lexer))
	Lexer:= GetLexerByName.CallObject(LexerArgs)
	if Lexer == nil{
		log.Fatal("Couldn't get lexer")
	}

	Highlighter := GetFunction("pygments", "highlight")
	if Highlighter  == nil{
		log.Fatal("aaasa")
	}
	HighlighterArgs := python.PyTuple_New(3)
	python.PyTuple_SetItem(HighlighterArgs, 0, python.PyString_FromString(code))
	python.PyTuple_SetItem(HighlighterArgs, 1, Lexer)
	python.PyTuple_SetItem(HighlighterArgs, 2, Formatter)

	highlighted := Highlighter.CallObject(HighlighterArgs)
	if highlighted == nil {
		log.Fatal("Couldn't highlight")
	}
	return python.PyString_AsString(highlighted)

}
