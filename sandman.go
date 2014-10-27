/*
Package sandman implements highlights code using pygments over a
python bridge using github.com/sbinet/go-python.
If go-python doesn't compile correctly try
	$ cd $GOPATH/src/github.com/sevki/sandman/
	$ make
*/
package sandman

import "log"

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func getFunction(moduleName string, functionName string) *python.PyObject {

	Module := python.PyImport_ImportModule(moduleName)
	if Module == nil {
		log.Fatal("Failed to load the " + moduleName + " module")
	}

	var MethodDesired *python.PyObject
	if Module.HasAttrString(functionName) == 1 {
		MethodDesired = Module.GetAttrString(functionName)
	}
	if !MethodDesired.Check_Callable() {
		log.Fatal(moduleName + " is not callable")
	}
	return MethodDesired

}

// Highlight higlights the given code snippet with the given lexer name.
// Adds line numbers if it linenos is true.
// List of available lexers are: http://lea.cx/pygments-lexers
func Highlight(code string, lexer string, linenos bool) string {
	lnos := 0
	if linenos {
		lnos = 1
	}

	GetFormatterByName := getFunction("pygments.formatters", "HtmlFormatter")
	FormatterArgs := python.PyTuple_New(0)

	Formatter := GetFormatterByName.CallObject(FormatterArgs)

	if Formatter == nil {
		log.Fatal("Couldn't get formatter")
	}
	if Formatter.HasAttrString("encoding") == 0 {
		log.Fatal("Wrong formatter")
	}
	if Formatter.HasAttrString("linenos") == 0 {
		log.Fatal("Wrong formatter")
	}

	Formatter.SetAttrString("encoding", python.PyString_FromString("utf-8"))
	Formatter.SetAttrString("linenos", python.PyBool_FromLong(lnos))

	GetLexerByName := getFunction("pygments.lexers", "get_lexer_by_name")
	LexerArgs := python.PyTuple_New(1)
	python.PyTuple_SetItem(LexerArgs, 0, python.PyString_FromString(lexer))
	Lexer := GetLexerByName.CallObject(LexerArgs)
	if Lexer == nil {
		log.Fatal("Couldn't get lexer " + lexer)
	}

	Highlighter := getFunction("pygments", "highlight")
	if Highlighter == nil {
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
