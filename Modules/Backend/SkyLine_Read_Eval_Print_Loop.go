/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//                              _____ _       __    _
//                             |   __| |_ _ _|  |  |_|___ ___
//                             |__   | '_| | |  |__| |   | -_|
//                             |_____|_,_|_  |_____|_|_|_|___|
//                                       |___|
//
// These sections are to help yopu better understand what each section is or what each file represents within the SkyLine programming language. These sections can also
//
// help seperate specific values so you can better understand what a specific section or specific set of values of functions is doing.
//
// These sections also give information on the file, project and status of the section
//
//
// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// Filename      |  SkyLine_Read_Eval_Print_Loop.go
// Project       |  SkyLine programming language
// Line Count    |  30 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines a REPL starter function
//
// State         | Working but can be worked on
// Changes?      | Can be thrown into another sub function file or file dedicated to REPL
//
//
package SkyLine_Backend

import (
	"bufio"
	"fmt"
	"io"
	"runtime"
)

const prompt = "SkyLine|%s(%s)>> "

func Start(in io.Reader, Out io.Writer) {
	scanner := bufio.NewScanner(in)
	Env := NewEnvironment()

	for {
		fmt.Printf(prompt, runtime.GOOS, runtime.GOARCH)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := LexNew(line)
		parser := New_Parser(l)

		program := parser.ParseProgram()
		if len(parser.ParserErrors()) != 0 {
			printParserErrors(line, Out, parser.ParserErrors())
			continue
		}
		evaluated := Eval(program, Env)
		if evaluated == nil {
			continue
		}
		io.WriteString(Out, evaluated.SL_InspectObject())
		io.WriteString(Out, "\n")
	}
}

func printParserErrors(line string, Out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(Out, msg)
		io.WriteString(Out, "\n")
	}
}
