package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"monkey/ast"
	"monkey/context"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object/function/environment"
	"monkey/parser"
	"os"
)

func main() {
	replFlag := flag.Bool("i", false, "start interactive mode after processing file")
	jsonFlag := flag.Bool("json", false, "dump AST of file as JSON")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		Start(os.Stdin, os.Stdout, nil)
		os.Exit(0)
	}

	fileName := args[0]
	fileStatus, errStatus := os.Lstat(fileName)
	if errStatus != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Could not obtain status of %s\n", fileName)
		os.Exit(4)
	}
	fileMode := fileStatus.Mode()
	if !fileMode.IsRegular() {
		_, _ = fmt.Fprintf(os.Stderr, "Not a regular file %s\n", fileName)
		os.Exit(5)
	}

	in, errOpen := os.OpenFile(fileName, os.O_RDONLY, 0)
	if errOpen != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to open %s\n", fileName)
		os.Exit(6)
	}

	errCode := ProcessProgram(in, os.Stdout, os.Stderr, *replFlag, *jsonFlag)
	os.Exit(errCode)
}

func ProcessProgram(in io.Reader, out io.Writer, errOut io.Writer, repl bool, json bool) int {
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		var errBuf bytes.Buffer
		_, _ = fmt.Fprintln(&errBuf, "Error reading input")
		_, _ = errOut.Write(errBuf.Bytes())
		return 1
	}
	source := string(buf)

	l := lexer.New(source)
	p := parser.New(l)
	syntaxTree := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(errOut, p.Errors())
		return 2
	}

	if json {
		dumpJSON(syntaxTree, os.Stdout)
	}

	//macros temporarily disabled
	//macroEnv := environment.NewEnvironment()
	//evaluator.DefineMacros(program, macroEnv)
	//expanded := evaluator.ExpandMacros(program, macroEnv)
	ctx := context.New(in, out)
	env := environment.New(ctx)
	evaluated := evaluator.Eval(syntaxTree, env)
	if evaluated != nil {
		_, _ = io.WriteString(out, evaluated.Inspect())
		_, _ = io.WriteString(out, "\n")
	}

	if repl {
		Start(os.Stdin, os.Stdout, env)
	}
	return 0
}

func dumpJSON(syntaxTree ast.Program, out io.Writer) {
	buf, errMarshal := syntaxTree.MarshalJSON()
	if errMarshal != nil {
		_, _ = fmt.Fprintln(out, fmt.Errorf("MarshalJSON() failed: %s", errMarshal))
		return
	}

	m := make(map[string]interface{})
	errUnmarshal := json.Unmarshal(buf, &m)
	if errUnmarshal != nil {
		_, _ = fmt.Fprintln(out, fmt.Errorf("json.Unmarshal failed: %s", errUnmarshal))
		_, _ = fmt.Fprintln(out, string(buf))
	}
	buf, errMarshal = json.MarshalIndent(m, "", "    ")
	_, _ = fmt.Fprintln(out, string(buf))
}

func printParserErrors(out io.Writer, errors []string) {
	_, _ = io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, env *environment.Environment) {
	scanner := bufio.NewScanner(in)
	if env == nil {
		ctx := context.New(in, out)
		env = environment.New(ctx)
	}
	//macros temporarily disabled
	//macroEnv := object.NewEnvironment()

	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		//macros temporarily disabled
		//evaluator.DefineMacros(program, macroEnv)
		//expanded := evaluator.ExpandMacros(program, macroEnv)
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}
