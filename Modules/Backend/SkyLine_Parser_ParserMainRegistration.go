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
		IDENT:          parser.parseIdent,
		INT:            parser.parseIntegerLiteral,
		FLOAT:          parser.parseFloatLiteral,
		BANG:           parser.parsePrefixExpression,
		MINUS:          parser.parsePrefixExpression,
		TRUE:           parser.parseBoolean,
		FALSE:          parser.parseBoolean,
		LPAREN:         parser.parseGroupedExpression,
		IF:             parser.parseIfExpression,
		FUNCTION:       parser.parseFunctionLiteral,
		STRING:         parser.parseStringLiteral,
		LBRACKET:       parser.parseArrayLiteral,
		LBRACE:         parser.parseHashLiteral,
		LINE:           parser.ParseGroupImportExpression,
		SWITCH:         parser.ParseSwitchCaseStatement,
		REGISTER:       parser.parseGroupedExpression,
		KEYWORD_ENGINE: parser.parseGroupedExpression,
		IMPORT:         parser.ParseImportExpression,
		STD_FOR:        parser.ParseForLoop,
	}

	parser.InfixParseFns = map[Token_Type]InfixParseFn{
		PLUS:            parser.parseInfixExpression,
		MINUS:           parser.parseInfixExpression,
		ASTARISK:        parser.parseInfixExpression,
		SLASH:           parser.parseInfixExpression,
		EQ:              parser.parseInfixExpression,
		NEQ:             parser.parseInfixExpression,
		LT:              parser.parseInfixExpression,
		GT:              parser.parseInfixExpression,
		LPAREN:          parser.parseCallExpression,
		LBRACKET:        parser.parseIndexExpression,
		GTEQ:            parser.parseInfixExpression,
		LTEQ:            parser.parseInfixExpression,
		PLUS_EQUALS:     parser.parseAssignmentStatement,
		DIVEQ:           parser.parseAssignmentStatement,
		MINUS_EQUALS:    parser.parseAssignmentStatement,
		ASTERISK_EQUALS: parser.parseAssignmentStatement,
		ASSIGN:          parser.parseAssignmentStatement,
		ANDAND:          parser.parseInfixExpression,
		OROR:            parser.parseInfixExpression,
		POWEROF:         parser.parseInfixExpression,
		PERIOD:          parser.parseMethodCallExpression,
		MODULECALL:      parser.parseSelectorExpression,
	}

	// Read two tokens, so curToken and peekToken are both set

	return parser
}
