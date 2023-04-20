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
// Filename      | SkyLine_Parser_SubRoutine_FileSys_CarrierFileParse.go
// Project       |  SkyLine programming language
// Line Count    |  110+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file contains fucntions that can read specific files or start new environments
//
// State         | Working but can be worked on
// Changes?      | Can be organized into a specific "loader" file or a file that revolves around loading or parsing data such as the file system files or set of files
//
//
//
package SkyLine_Backend

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func RunAndParseFiles(filename string) error {
	FileCurrent.New(filename)
	f, x := os.Open(filename)
	if x != nil {
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(x)
			}
		}()
		fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_DOES_NOT_EXIST](filename))
		os.Exit(0)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	if line == nil {
		fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_IS_EMPTY](filename))
	}
	data, x := ioutil.ReadFile(filename)
	if x != nil {
		fmt.Println(Map_Parser[ERROR_FILE_INPUT_OUTPUT_BUFFER_FAILED](filename, fmt.Sprint(x)))
		os.Exit(1)
	}
	parser := New_Parser(LexNew(string(data)))
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
					fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_IS_EMPTY](filename))
				}
			}
		}
	}()
	_, x = io.WriteString(os.Stdout, result.SL_InspectObject()+"\n")
	return x
}
