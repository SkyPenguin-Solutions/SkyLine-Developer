package main

import (
	"flag"
	"fmt"
	Mod "main/Modules/Backend"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	flag.Parse()
}

func main() {
	if *Mod.CompileWithGo {
		if *Mod.SourceFile != "" {
			if fileext := filepath.Ext(*Mod.SourceFile); fileext == ".csc" {
				if err := Mod.RunAndParseFiles(*Mod.SourceFile); err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(0)
				}
			} else {
				fmt.Println("Error: File is not CSC file")
			}
		}
	}
	if *Mod.Help {
		Mod.Banner()
		fmt.Printf("CSC (Cyber Security Core) | Golang version      ( %s ) \n", runtime.Version())
		fmt.Printf("CSC (Cyber Security Core) | OS Picked Up        ( %s ) \n", Mod.U.OperatingSystem)
		fmt.Printf("CSC (Cyber Security Core) | Architecture        ( %s ) ", Mod.U.OperatingSystemArchitecture)
		fmt.Print("\n\n")
		keys := `
	-e           Run a line of code through the interpreter 
	-source      Specify a source code file to execute
	-build       Will take an input file -source and compile with the interpreter
	-trace       Will trace back any panics or errors that may occure during runtime
		`
		fmt.Println(keys)
		os.Exit(0)
	}
	if *Mod.RunRawC != "" {
		Data := *Mod.RunRawC
		if x := Mod.RunAndParseNoFile(Data); x != nil {
			fmt.Fprintln(os.Stderr, x)
			os.Exit(0)
		}
		os.Exit(0)
	}
	if *Mod.SourceFile != "" {
		if *Mod.Bnn {
			Mod.Banner()
		}
		if fileext := filepath.Ext(*Mod.SourceFile); fileext == ".csc" {
			if err := Mod.RunAndParseFiles(*Mod.SourceFile); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(0)
			}
		} else {
			fmt.Println("Woah there buddy, sorry but that file type is not allowed, files must end in (.csc) not -> ", fileext)
		}
	} else {
		fmt.Println("\x1b[H\x1b[2J\x1b[3J")
		Mod.Banner()
		Mod.Start(os.Stdin, os.Stdout)
		return
	}
}
