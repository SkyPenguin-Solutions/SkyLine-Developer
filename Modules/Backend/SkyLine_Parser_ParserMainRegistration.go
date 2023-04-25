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
// Filename      |  SkyLine_Parser_ParserMainRegistration.go
// Project       |  SkyLine programming language
// Line Count    |  80+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines -> This file defines all of the parser prefix and infix parsing maps which is the way to parse specific tokens. In this file it helps us output and run specific functions
//            based on a token's type
//
//
package SkyLine_Backend

func New_Parser(l *ScannerStructure) *Parser {
	parser := &Parser{
		Lex:    l,
		Errors: []string{},
	}
	parser.NextLoadFaultToken()
	parser.NextLoadFaultToken()

	parser.PrefixParseFns = map[Token_Type]PrefixParseFn{
		IDENT:          parser.SkyLine_Identifier,
		INT:            parser.SkyLine_IntegerLiteral,
		FLOAT:          parser.SkyLine_FloatLiteral,
		BANG:           parser.SkyLine_PrefixExpression,
		MINUS:          parser.SkyLine_PrefixExpression,
		TRUE:           parser.SkyLine_Boolean,
		FALSE:          parser.SkyLine_Boolean,
		LPAREN:         parser.SkyLine_GroupedExpression,
		IF:             parser.SkyLine_ConditionalExpression,
		FUNCTION:       parser.SkyLine_FunctionLiteral,
		STRING:         parser.parseStringLiteral,
		LBRACKET:       parser.parseArrayLiteral,
		LBRACE:         parser.parseHashLiteral,
		LINE:           parser.SkyLine_GroupImportExpression,
		SWITCH:         parser.SkyLine_SwitchCase_Expression,
		REGISTER:       parser.SkyLine_GroupedExpression,
		KEYWORD_ENGINE: parser.SkyLine_GroupedExpression,
		IMPORT:         parser.SkyLine_ImportExpression,
		TOKEN_FOR:      parser.SkyLine_ForLoop,
		TOKEN_FOREACH:  parser.SkyLine_ForEach,
	}

	parser.InfixParseFns = map[Token_Type]InfixParseFn{
		PLUS:            parser.SkyLine_InfixExpression,
		MINUS:           parser.SkyLine_InfixExpression,
		ASTARISK:        parser.SkyLine_InfixExpression,
		SLASH:           parser.SkyLine_InfixExpression,
		EQ:              parser.SkyLine_InfixExpression,
		NEQ:             parser.SkyLine_InfixExpression,
		LT:              parser.SkyLine_InfixExpression,
		GT:              parser.SkyLine_InfixExpression,
		LPAREN:          parser.parseCallExpression,
		LBRACKET:        parser.parseIndexExpression,
		GTEQ:            parser.SkyLine_InfixExpression,
		LTEQ:            parser.SkyLine_InfixExpression,
		PLUS_EQUALS:     parser.SkyLine_Assignment,
		DIVEQ:           parser.SkyLine_Assignment,
		MINUS_EQUALS:    parser.SkyLine_Assignment,
		ASTERISK_EQUALS: parser.SkyLine_Assignment,
		ASSIGN:          parser.SkyLine_Assignment,
		ANDAND:          parser.SkyLine_InfixExpression,
		OROR:            parser.SkyLine_InfixExpression,
		POWEROF:         parser.SkyLine_InfixExpression,
		PERIOD:          parser.parseMethodCallExpression,
		MODULECALL:      parser.SkyLine_SelectorExpression,
	}
	return parser
}
