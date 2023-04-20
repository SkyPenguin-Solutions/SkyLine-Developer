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
// Filename      |  SkyLine_Parser_Subroutine_FileSys_ParserRunner.go
// Project       |  SkyLine programming language
// Line Count    |  50 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines a function to run and parse data that comes in from the REPL ( Read Eval Print Loop)
//
// State         | Working but can be worked on
// Changes?      | Can be organized into a specific "loader" file or a file that revolves around loading or parsing data such as the file system files or set of files
//
//
//
package SkyLine_Backend

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func RunAndParseNoFile(data string) error {
	parser := New_Parser(LexNew(data))
	program := parser.ParseProgram()
	if len(parser.ParserErrors()) > 0 {
		return errors.New(parser.ParserErrors()[0])
	}
	Env := NewEnvironment()
	result := Eval(program, Env)
	if _, ok := result.(*Nil); ok {
		return nil
	}
	defer func() {
		if xy := recover(); xy != nil {
			if strings.Contains(fmt.Sprint(xy), "invalid memory address or nil pointer dereference") {
				if *ErrorsTrace {
					fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_IS_EMPTY]("FILE_SECTOR_1:WAS_NOT_FILE_USED_FLAG_(-e) | data when using flag -e was empty, interpreter will NOT run"))
				}
			}
		}
	}()
	_, x := io.WriteString(os.Stdout, result.SL_InspectObject()+"\n")
	return x
}
