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
// Filename      |  SkyLine_Parser_ParserMainCallFunctions.go
// Project       |  SkyLine programming language
// Line Count    |  1,200+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This entire file is dedicated to every function for the parser or all main functions anyway which return their respective data types. It is important
//                 to mention that these functions are very important and need to be organized.
//
// State         | Working and secure but needs major changes
//
// Resolution    | These functions can be named better, worked better, modified better and even moved into another specific file while also having changes such as
//                 not flooding the file with trees
//
//
package SkyLine_Backend

import (
	"fmt"
	"strconv"
	"strings"
)

var root *TreeNode

// Parser helper functions
func ParserErrorSystem_GetFileName() (fname string) {
	if FileCurrent.Filename == "" && FileCurrent.Get_Name() == "" {
		fname = "REPL_main.skyline"
	} else {
		fname = FileCurrent.Filename
	}
	return fname
}

func (parser *Parser) ParseEngineStatement() *ENGINE {
	statement := &ENGINE{Token: parser.CurrentToken}
	if !parser.PeekTokenIs("(") {
		fmt.Println("Missing ( in import statement")
	}
	parser.NT()
	statement.EngineValue = parser.parseExpression(LOWEST)
	for parser.PeekTokenIs(SEMICOLON) {
		parser.NT()
	}
	return statement
}

func (parser *Parser) parseStatement() Statement {
	switch parser.CurrentToken.Token_Type {
	case KEYWORD_ENGINE:
		return parser.ParseEngineStatement()
	case REGISTER:
		return parser.ParseRegisterStatement()
	case LET:
		return parser.ParserCreateAssignment()
	case RETURN:
		return parser.parseReturnStatement()
	case CONSTANT:
		return parser.ParseConstants()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) ParseSwitchCaseStatement() Expression {
	EXP := &Switch{Token: parser.CurrentToken}
	EXP.Value = parser.ParseArgumentParens()
	if EXP.Value == nil {
		return nil
	}
	if !parser.ExpectPeek(LBRACE) {
		return nil
	}
	parser.NT()
	for !parser.CurrentTokenIs(RBRACE) {
		if parser.CurrentTokenIs(EOF) {
			parser.Errors = append(parser.Errors, "unterminated switch statement")
		}
		exp := &Case{Token: parser.CurrentToken}
		if parser.CurrentTokenIs(DEFAULT) {
			exp.Def = true

		} else if parser.CurrentTokenIs(CASE) {

			parser.NT()
			if parser.CurrentTokenIs(DEFAULT) {
				exp.Def = true
			} else {
				exp.Expr = append(exp.Expr, parser.parseExpression(LOWEST))
				for parser.PeekTokenIs(COMMA) {
					parser.NT()
					parser.NT()
					exp.Expr = append(exp.Expr, parser.parseExpression(LOWEST))
				}
			}
		} else {
			parser.Errors = append(parser.Errors, fmt.Sprintf("expected case|default, got %s >>> %s ", parser.CurrentToken.Token_Type, parser.CurrentToken))
			return nil
		}

		if !parser.ExpectPeek(LBRACE) {

			msg := fmt.Sprintf("expected token to be '{', got %s instead", parser.CurrentToken.Token_Type)
			parser.Errors = append(parser.Errors, msg)
			fmt.Printf("error\n")
			return nil
		}

		// parse the block
		exp.Block = parser.parseBlockStatement()

		if !parser.CurrentTokenIs(RBRACE) {
			msg := fmt.Sprintf("Syntax Error: expected token to be '}', got %s instead", parser.CurrentToken.Token_Type)
			parser.Errors = append(parser.Errors, msg)
			fmt.Printf("error\n")
			return nil

		}
		parser.NT()
		EXP.Choices = append(EXP.Choices, exp)
	}
	count := 0
	for _, c := range EXP.Choices {
		if c.Def {
			count++
		}
	}
	if count > 1 {
		msg := "A switch-statement should only have one default block"
		parser.Errors = append(parser.Errors, msg)
		return nil

	}
	return EXP

}

func (parser *Parser) ParseArgumentParens() Expression {
	if !parser.ExpectPeek(LPAREN) {
		parser.Errors = append(parser.Errors, fmt.Sprintf("Unexpected token | %s | I need %s for this argument list ", parser.CurrentToken.Literal, LPAREN))
		return nil
	}
	parser.NT()
	exp := parser.parseExpression(LOWEST)
	if exp == nil {
		return nil
	}
	if !parser.ExpectPeek(RPAREN) {
		parser.Errors = append(parser.Errors, fmt.Sprintf("Unexpected token | %s|  I need %s for this statement ", parser.CurrentToken.Literal, RPAREN))
		return nil
	}
	return exp
}

func (parser *Parser) ParserCreateAssignment() *LetStatement {
	stmt := &LetStatement{Token: parser.CurrentToken}
	if !parser.ExpectPeek(IDENT) {
		return nil
	}
	stmt.Name = &Ident{Token: parser.CurrentToken, Value: parser.CurrentToken.Literal}
	if !parser.ExpectPeek(ASSIGN) {
		return nil
	}
	parser.NT()
	stmt.Value = parser.parseExpression(LOWEST)
	for !parser.CurrentTokenIs(SEMICOLON) {
		if parser.CurrentTokenIs(EOF) {
			var fname string
			if FileCurrent.Filename == "" {
				fname = "REPL.skyline"
			} else {
				fname = FileCurrent.Filename
			}

			root = &TreeNode{
				Type: SKYLINE_HIGH_DEFRED + "E | " + fname + SKYLINE_RESTORE,
				Children: []*TreeNode{
					{
						Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Code" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprintf("%v", ERROR_MISSING_SEMICOLON_IN_STATEMENT_AT) + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Type" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											"Parser Error (parse create assignment)" +
											SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Message" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											Map_Parser[ERROR_MISSING_SEMICOLON_IN_STATEMENT_AT](stmt.SL_ExtractStringValue(), SKYLINE_HIGH_DEFPURPLE, SKYLINE_RESTORE).Message +
											SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(parser.GetLineCound()) + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated statement   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + strings.Trim(stmt.SL_ExtractStringValue(), ";"),
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_SICK_BLUE + "[S] " + "Consider adding a semicolon to the end of the statement (';')" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Fixed Satatement" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_FIXBLUE + "[F] " + stmt.SL_ExtractStringValue(),
									},
								},
							},
						},
					},
				},
			}
			if ErrorSys.TreeValid() {
				RetTreeSys(root, "", false)
			}
			if ErrorSys.LineValid() {
				var msg string
				msg += SKYLINE_HIGH_DEFRED + "[E] | " + SKYLINE_RESTORE + FileCurrent.GetAbsolute() + "\n"
				msg += SKYLINE_HIGH_DEFRED + "[1] | " + SKYLINE_RESTORE + "Error when trying to parse statement, missing semicolon" + "\n"
				msg += SKYLINE_HIGH_FIXBLUE + "[F] | " + SKYLINE_RESTORE + stmt.SL_ExtractStringValue() + "\n"
				msg += SUNRISE_LIGHT_DEFINITION + "[L] | " + SKYLINE_RESTORE + parser.GetLineCound() + "\n"
				msg += SKYLINE_SICK_BLUE + "[S] | " + SKYLINE_RESTORE + "Suggested to put a semicolon at the end of the statement"
				fmt.Println(msg)
			}
			parser.Errors = append(parser.Errors, "")
			return nil
		}
		parser.NT()
	}
	return stmt

}

// wwork on this error here
func (parser *Parser) parseAssignmentStatement(name Expression) Expression {
	stmt := &AssignmentStatement{Token: parser.CurrentToken}
	if StatementName, ok := name.(*Ident); ok {
		stmt.Name = StatementName
	} else {
		parser.Errors = append(parser.Errors, fmt.Sprintf("Expected assignment token before operator to be an IDENTIFIER not %s", name.SL_ExtractNodeValue()))
	}
	opperand := parser.CurrentToken
	parser.NT()
	switch opperand.Token_Type {
	case PLUS_EQUALS:
		stmt.Operator = "+="
	case MINUS_EQUALS:
		stmt.Operator = "-="
	case DIVEQ:
		stmt.Operator = "/="
	case ASTERISK_EQUALS:
		stmt.Operator = "*="
	default:
		stmt.Operator = "="
	}

	stmt.Value = parser.parseExpression(LOWEST)
	return stmt
}

func (parser *Parser) ParseConstants() *Constant {
	statement := &Constant{Token: parser.CurrentToken}
	if !parser.ExpectPeek(IDENT) {
		return nil
	}
	statement.Name = &Ident{Token: parser.CurrentToken, Value: parser.CurrentToken.Literal}
	if !parser.ExpectPeek(ASSIGN) {
		return nil
	}
	parser.NT()
	statement.Value = parser.parseExpression(LOWEST)
	for !parser.CurrentTokenIs(SEMICOLON) {
		if parser.CurrentTokenIs(EOF) {
			root = &TreeNode{
				Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
				Children: []*TreeNode{
					{
						Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											fmt.Sprintf(
												"%v",
												ERROR_UNTERMINATED_CONSTANT_VALUE,
											) +
											SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse Constant)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Message" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											Map_Parser[ERROR_UNTERMINATED_CONSTANT_VALUE](
												statement.SL_ExtractStringValue(),
												SKYLINE_HIGH_DEFPURPLE,
												SKYLINE_RESTORE,
											).Message,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(parser.GetLineCound()) + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated statement   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + strings.Trim(statement.SL_ExtractStringValue(), ";"),
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											Map_Parser[ERROR_UNTERMINATED_CONSTANT_VALUE](
												statement.SL_ExtractStringValue(),
												SKYLINE_HIGH_DEFPURPLE,
												SKYLINE_RESTORE,
											).Suggestion,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Fixed Statement" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_FIXBLUE + "[F] " + statement.SL_ExtractStringValue(),
									},
								},
							},
						},
					},
				},
			}
			if ErrorSys.TreeValid() {
				RetTreeSys(root, "", false)
			}
			if ErrorSys.LineValid() {
				fmt.Println(parser.Prepare_Base_Error_Message(
					"Missing semicolon at the end of constant assignment",
					strings.Trim(statement.SL_ExtractStringValue(), ";"),
					true,
					statement.SL_ExtractStringValue(),
				))
			}
			parser.Errors = append(parser.Errors, "")
			return nil
		}
		parser.NT()
	}
	return statement
}

func (parser *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{
		Token: parser.CurrentToken,
	}
	parser.NT()

	stmt.ReturnValue = parser.parseExpression(LOWEST)

	for parser.PeekTokenIs(SEMICOLON) {
		parser.NT()
	}

	return stmt
}

func (parser *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{
		Token:      parser.CurrentToken,
		Expression: parser.parseExpression(LOWEST),
	}

	if parser.PeekTokenIs(SEMICOLON) {
		parser.NT()
	}

	return stmt
}

func (parser *Parser) parseExpression(precedence int) Expression {
	prefix := parser.PrefixParseFns[parser.CurrentToken.Token_Type]

	if prefix == nil {
		root = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA +
										fmt.Sprintf(
											"%v",
											ERROR_PREFIX_PARSE_FUNCTION_NOT_LOADED_INTO_ENV,
										) +
										SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse Expression) " + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA +
										Map_Parser[ERROR_PREFIX_PARSE_FUNCTION_NOT_LOADED_INTO_ENV]().Message +
										SKYLINE_RESTORE,
								},
								{
									Type: SKYLINE_HIGH_DEFAQUA + "[Sub Branch] Token " + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SKYLINE_HIGH_DEFAQUA + parser.CurrentToken.Literal + SKYLINE_RESTORE,
										},
									},
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(parser.GetLineCound()),
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute(),
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Token       : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.CurrentToken.Literal,
								},
							},
						},
					},
				},
			},
		}
		if ErrorSys.TreeValid() {
			RetTreeSys(root, "", false)
		}
		if ErrorSys.LineValid() {
			var msg string
			msg += SKYLINE_HIGH_DEFRED + "[E] | " + SKYLINE_RESTORE + FileCurrent.GetAbsolute() + "\n"
			msg += SKYLINE_HIGH_DEFRED + "[1] | " + SKYLINE_RESTORE + "Could not locate a prefix parse function for the parsed token" + "\n"
			msg += SKYLINE_HIGH_FIXBLUE + "[F] | " + SKYLINE_RESTORE + parser.CurrentToken.Literal + "\n"
			msg += SUNRISE_LIGHT_DEFINITION + "[L] | " + SKYLINE_RESTORE + parser.GetLineCound() + "\n"
			msg += SKYLINE_SICK_BLUE + "[S] | " + SKYLINE_RESTORE + "Check your token?"
			fmt.Println(msg)
		}
		parser.Errors = append(parser.Errors, "")
		return nil
	}
	leftExp := prefix()

	for !parser.CurrentTokenIs(SEMICOLON) && precedence < parser.peekPrecedence() {
		infix := parser.InfixParseFns[parser.PeekToken.Token_Type]
		if infix == nil {
			return leftExp
		}

		parser.NT()

		leftExp = infix(leftExp)
	}
	return leftExp
}

func (parser *Parser) parseIdent() Expression {
	return &Ident{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
}

//TODO: Remake this error system for 90%
func (parser *Parser) parseIntegerLiteral() Expression {
	lit := &IntegerLiteral{Token: parser.CurrentToken}
	var value int64
	var x error
	var errtype string
	if strings.HasPrefix(parser.CurrentToken.Literal, "0b") {
		value, x = strconv.ParseInt(parser.CurrentToken.Literal[2:], 2, 64)
		if x != nil {
			errtype = "binary"
		}
	} else if strings.HasPrefix(parser.CurrentToken.Literal, "0x") {
		value, x = strconv.ParseInt(parser.CurrentToken.Literal[2:], 16, 64)
		if x != nil {
			errtype = "hex"
		}
	} else {
		value, x = strconv.ParseInt(parser.CurrentToken.Literal, 0, 64)
		if x != nil {
			errtype = "int"
		}
	}
	// Yes we check again, simple, dont ask
	if x != nil {
		_, x := CheckParseValue(parser.CurrentToken.Literal, errtype)
		root = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_TYPE_INTEGRITY_PARSE_INTEGER_ERROR) + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse Integer Literal)" + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser could not parse the given integer" + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "[Sub Branch] Debug " + SKYLINE_RESTORE,
											Children: []*TreeNode{
												{

													Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Error In Type  ? " + SUNRISE_HIGH_DEFINITION + x.Type + SKYLINE_RESTORE,
												},
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Error In Value ? " + SUNRISE_HIGH_DEFINITION + x.Value + SKYLINE_RESTORE,
												},
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Error In Parse ? " + SUNRISE_HIGH_DEFINITION + x.Err.Error() + SKYLINE_RESTORE,
												},
											},
										},
									},
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number  : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path    : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + lit.SL_ExtractStringValue(),
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Max possible value for int64    ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Max + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Lowest possible value for int64 ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Lowest + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Suggestion " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_SICK_BLUE + "[S] " + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Suggest + SKYLINE_RESTORE,
								},
							},
						},
					},
				},
			},
		}
		if ErrorSys.TreeValid() {
			RetTreeSys(root, "", false)
		}
		if ErrorSys.LineValid() {
			var msg string
			msg += SKYLINE_HIGH_DEFRED + "[E] | " + SKYLINE_RESTORE + FileCurrent.GetAbsolute() + "\n"
			msg += SKYLINE_HIGH_DEFRED + "[1] | " + SKYLINE_RESTORE + "Could not parse the integer/hex/binary number given" + "\n"
			msg += SUNRISE_LIGHT_DEFINITION + "[L] | " + SKYLINE_RESTORE + parser.GetLineCound() + "\n"
			msg += SKYLINE_SICK_BLUE + "[S] | " + SKYLINE_RESTORE + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Suggest + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Max    (int64)  -> " + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Max + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Lowest (int64)  -> " + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Lowest + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Type of (Error) -> " + x.Type + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Error Message   -> " + x.Err.Error() + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Value to parse  -> " + x.Value + SKYLINE_RESTORE
			fmt.Println(msg)
		}
		parser.Errors = append(parser.Errors, "")
	}
	lit.Value = value
	return lit
}

func (parser *Parser) parseFloatLiteral() Expression {
	val, err := strconv.ParseFloat(parser.CurrentToken.Literal, 64)
	if err != nil {
		duringparse := CheckAndVerify(parser.CurrentToken.Literal)
		root = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_COULD_NOT_PARSE_FLOAT_VALUE) + SKYLINE_RESTORE,
						},
					},
				},
				{
					Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFAQUA + "Parser Error ( Parse Float Literal ) " + SKYLINE_RESTORE,
						},
					},
				},
				{
					Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFAQUA + "Parser was not able to process the provided input as a float64",
							Children: []*TreeNode{
								{
									Type: SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "[Sub Branch] Debug " + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Value too large   ? " + SUNRISE_HIGH_DEFINITION + fmt.Sprint(duringparse.TooLong) + SKYLINE_RESTORE,
										},
										{
											Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Value too small   ? " + SUNRISE_HIGH_DEFINITION + fmt.Sprint(duringparse.TooShort) + SKYLINE_RESTORE,
										},
										{
											Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Value parsed      ? " + SUNRISE_HIGH_DEFINITION + fmt.Sprint(duringparse.Parsed) + SKYLINE_RESTORE,
										},
									},
								},
							},
						},
					},
				},
				{
					Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path    ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.CurrentToken.Literal,
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] MAX (Float64)          ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(duringparse.Max) + SKYLINE_RESTORE,
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] MIN (Float64)          ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(duringparse.Low) + SKYLINE_RESTORE,
						},
					},
				},
				{
					Type: SUNRISE_LIGHT_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_SICK_BLUE + "[S] " + duringparse.Recomend + SKYLINE_RESTORE,
						},
					},
				},
			},
		}
		if ErrorSys.TreeValid() {
			RetTreeSys(root, "", false)
		}
		if ErrorSys.LineValid() {
			var msg string
			msg += SKYLINE_HIGH_DEFRED + "[E] | " + SKYLINE_RESTORE + FileCurrent.GetAbsolute() + "\n"
			msg += SKYLINE_HIGH_DEFRED + "[1] | " + SKYLINE_RESTORE + "Could not parse the integer/hex/binary number given" + "\n"
			msg += SUNRISE_LIGHT_DEFINITION + "[L] | " + SKYLINE_RESTORE + parser.GetLineCound() + "\n"
			msg += SKYLINE_SICK_BLUE + "[S] | " + SKYLINE_RESTORE + duringparse.Recomend + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Max    (int64)  -> " + fmt.Sprint(duringparse.Max) + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Lowest (int64)  -> " + fmt.Sprint(duringparse.Low) + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Too long        -> " + fmt.Sprint(duringparse.TooLong) + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Too short       -> " + fmt.Sprint(duringparse.TooShort) + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Value to parse  -> " + parser.CurrentToken.Literal + SKYLINE_RESTORE
			fmt.Println(msg)
		}
		parser.Errors = append(parser.Errors, "")
		return nil
	}

	return &FloatLiteral{
		Token: parser.CurrentToken,
		Value: val,
	}
}

func (parser *Parser) parsePrefixExpression() Expression {
	expr := &PrefixExpression{
		Token:    parser.CurrentToken,
		Operator: parser.CurrentToken.Literal,
	}

	parser.NT()

	expr.Right = parser.parseExpression(PREFIX)
	return expr
}

func (parser *Parser) peekPrecedence() int {
	if p, ok := Precedences[parser.PeekToken.Token_Type]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) curPrecedence() int {
	if p, ok := Precedences[parser.CurrentToken.Token_Type]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) parseInfixExpression(left Expression) Expression {
	expr := &InfixExpression{
		Token:    parser.CurrentToken,
		Operator: parser.CurrentToken.Literal,
		Left:     left,
	}

	prec := parser.curPrecedence()

	parser.NT()

	expr.Right = parser.parseExpression(prec)
	return expr
}

func (parser *Parser) parseBoolean() Expression {
	return &Boolean_AST{
		Token: parser.CurrentToken,
		Value: parser.CurrentTokenIs(TRUE),
	}
}

func (parser *Parser) ParseGroupImportExpression() Expression {
	parser.NT()

	expr := parser.parseExpression(LOWEST)

	if !parser.ExpectPeek(LINE) {
		return nil
	}

	return expr
}

func (parser *Parser) parseGroupedExpression() Expression {
	parser.NT()

	expr := parser.parseExpression(LOWEST)

	if !parser.ExpectPeek(RPAREN) {
		return nil
	}

	return expr
}

func (parser *Parser) parseIfExpression() Expression {
	expr := &ConditionalExpression{Token: parser.CurrentToken}

	parser.NT()
	expr.Condition = parser.parseExpression(LOWEST)
	if !parser.ExpectPeek(LBRACE) {
		return nil
	}

	expr.Consequence = parser.ParseConditionalBlock()
	if expr.Consequence == nil {
		root = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_PARSER_FOUND_NIL_EXPRESSION_UNEXPECTED) + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse If Expression)" + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser found an empty or nil consequence (UNEXPECT:NIL->EXPRESSION)" + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number  : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path    : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + expr.Condition.SL_ExtractStringValue(),
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Suggestion " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_SICK_BLUE + "[S] Make sure the consequence is not empty for the conditional" + SKYLINE_RESTORE,
								},
							},
						},
					},
				},
			},
		}
		RetTreeSys(root, "", true)
		parser.Errors = append(parser.Errors, "")
		return nil
	}
	if parser.PeekTokenIs(ELSE) {
		parser.NT()

		if parser.PeekTokenIs(IF) {

			parser.NT()
			expr.Alternative = &BlockStatement{
				Statements: []Statement{
					&ExpressionStatement{
						Expression: parser.ParseConditionalBlock(),
					},
				},
			}
			return expr
		}

		// parse else

		if !parser.ExpectPeek(LBRACE) {
			root = &TreeNode{
				Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
				Children: []*TreeNode{
					{
						Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_PARSER_EXPECTED_LBRACE_BUT_GOT_SOMETHING) + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse If Expression)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser found an unexpected token but needs '{' (UNEXPECT:Token)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path     : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed  : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + expr.Condition.SL_ExtractStringValue() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Parser Found Unexpected : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.CurrentToken.Literal + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Suggestion " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_SICK_BLUE + "[S] Make sure the conditional includes the proper statement" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Fixed Satatement " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_FIXBLUE + "[F] " + expr.SL_ExtractStringValue() + SKYLINE_RESTORE,
									},
								},
							},
						},
					},
				},
			}
			RetTreeSys(root, "", false)
			parser.Errors = append(parser.Errors, "")
			return nil
		}

		expr.Alternative = parser.ParseConditionalBlock()
		if expr.Alternative == nil {
			root = &TreeNode{
				Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
				Children: []*TreeNode{
					{
						Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_PARSER_FOUND_NIL_EXPRESSION_UNEXPECTED) + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse If Expression)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser found an empty or nil consequence (UNEXPECT:NIL->EXPRESSION)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number  : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path    : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + expr.Condition.SL_ExtractStringValue(),
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Suggestion " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_SICK_BLUE + "[S] Make sure the consequence is not empty for the conditional" + SKYLINE_RESTORE,
									},
								},
							},
						},
					},
				},
			}
			RetTreeSys(root, "", true)
			return nil
		}
	}

	return expr
}

func (parser *Parser) ParseConditionalBlock() *BlockStatement {
	block := &BlockStatement{
		Token:      parser.CurrentToken,
		Statements: []Statement{},
	}
	parser.NT()
	for !parser.CurrentTokenIs(RBRACE) && !parser.CurrentTokenIs(EOF) {
		stmt := parser.parseStatement()
		block.Statements = append(block.Statements, stmt)
		parser.NT()
	}
	return block
}

func (parser *Parser) parseBlockStatement() *BlockStatement {
	block := &BlockStatement{
		Token:      parser.CurrentToken,
		Statements: []Statement{},
	}

	parser.NT()
	for !parser.CurrentTokenIs(RBRACE) && !parser.CurrentTokenIs(EOF) {
		stmt := parser.parseStatement()
		block.Statements = append(block.Statements, stmt)
		parser.NT()
	}
	return block
}

func (parser *Parser) parseFunctionLiteral() Expression {
	lit := &FunctionLiteral{Token: parser.CurrentToken}

	if !parser.ExpectPeek(LPAREN) {
		return nil
	}

	lit.Parameters = parser.parseFunctionParameters()

	if !parser.ExpectPeek(LBRACE) {
		return nil
	}

	lit.Body = parser.parseBlockStatement()

	return lit
}

func (parser *Parser) parseFunctionParameters() []*Ident {
	idents := []*Ident{}

	if parser.PeekTokenIs(RPAREN) {
		parser.NT()
		return idents
	}

	parser.NT()

	ident := &Ident{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
	idents = append(idents, ident)

	for parser.PeekTokenIs(COMMA) || parser.PeekTokenIs(COLON) {
		parser.NT()
		parser.NT()
		ident := &Ident{
			Token: parser.CurrentToken,
			Value: parser.CurrentToken.Literal,
		}
		idents = append(idents, ident)
	}

	if !parser.ExpectPeek(RPAREN) {
		return nil
	}

	return idents
}

func (parser *Parser) parseExpressionList(end Token_Type) []Expression {
	list := make([]Expression, 0)

	if parser.PeekTokenIs(end) {
		parser.NT()
		return list
	}

	parser.NT()
	list = append(list, parser.parseExpression(LOWEST))

	for parser.PeekTokenIs(COMMA) {
		parser.NT()
		parser.NT()
		list = append(list, parser.parseExpression(LOWEST))
	}

	if !parser.ExpectPeek(end) {
		return nil
	}

	return list
}

func (parser *Parser) parseCallExpression(function Expression) Expression {
	return &CallExpression{
		Token:     parser.CurrentToken,
		Function:  function,
		Arguments: parser.parseExpressionList(RPAREN),
	}
}

func (parser *Parser) parseStringLiteral() Expression {
	return &StringLiteral{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
}

func (parser *Parser) parseArrayLiteral() Expression {
	return &ArrayLiteral{
		Token:    parser.CurrentToken,
		Elements: parser.parseExpressionList(RBRACKET),
	}
}

func (parser *Parser) parseIndexExpression(left Expression) Expression {
	expr := &IndexExpression{
		Token: parser.CurrentToken,
		Left:  left,
	}

	parser.NT()
	expr.Index = parser.parseExpression(LOWEST)

	if !parser.ExpectPeek(RBRACKET) {
		return nil
	}

	return expr
}

func (parser *Parser) parseHashLiteral() Expression {
	hash := &HashLiteral{
		Token: parser.CurrentToken,
		Pairs: make(map[Expression]Expression),
	}

	for !parser.PeekTokenIs(RBRACE) {
		parser.NT()
		key := parser.parseExpression(LOWEST)

		if !parser.ExpectPeek(COLON) {
			return nil
		}

		parser.NT()
		value := parser.parseExpression(LOWEST)
		hash.Pairs[key] = value

		if !parser.PeekTokenIs(RBRACE) && !parser.ExpectPeek(COMMA) {
			return nil
		}
	}

	if !parser.ExpectPeek(RBRACE) {
		return nil
	}

	return hash
}

func (parser *Parser) parseMethodCallExpression(obj Expression) Expression {
	methodcall := &ObjectCallExpression{
		Token:      parser.CurrentToken,
		SLC_Object: obj,
	}
	parser.NT()
	name := parser.parseIdent()
	parser.NT()
	methodcall.Call = parser.parseCallExpression(name)
	return methodcall
}

func (parser *Parser) ParseRegisterStatement() *Register {
	statement := &Register{Token: parser.CurrentToken}
	if !parser.PeekTokenIs("(") {
		fmt.Println("Missing ( in import statement")
	}
	parser.NT()
	statement.RegistryValue = parser.parseExpression(LOWEST)
	for parser.PeekTokenIs(SEMICOLON) {
		parser.NT()
	}
	return statement
}

/*

PARSE MODIFY FUNCTIOn - LEGACY STYLE ( modify should not be a keyword )

func (parser *Parser) ParseModifierStatement() *ModifyEnv {
	stmt := &ModifyEnv{Token: parser.CurrentToken}
	if !parser.PeekTokenIs("(") {
		fmt.Println("Missing '(' in modifier statement")
	}
	parser.NT()
	stmt.Modifier = parser.parseExpression(LOWEST)
	for parser.PeekTokenIs(SEMICOLON) {
		parser.NT()
	}
	return stmt
}

	case MODIFY:
		return parser.ParseModifierStatement()

*/
