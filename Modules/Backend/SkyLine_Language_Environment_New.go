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
// Filename      |  SkyLine_Language_Environment_New.go
// Project       |  SkyLine programming language
// Line Count    |  5 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines a function to create a new environment
//
//
package SkyLine_Backend

func NewEnvironment() *Environment_of_environment {
	return &Environment_of_environment{
		Store: make(map[string]SLC_Object),
		Outer: nil,
		ROM:   make(map[string]bool),
	}
}
