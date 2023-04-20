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
// Filename      |  SkyLine_Scanner_Verification_TokenizerAndCategorizer.go
// Project       |  SkyLine programming language
// Line Count    |  50+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines prime categorization functions for the scanner
//
//
package SkyLine_Backend

type TokenizerFunc func(*ScannerStructure) Token
type TokenizerPeeked func(*ScannerStructure) Token

func Categorize_EOF(lex *ScannerStructure) Token {
	var tk Token
	tk.Literal = ""
	tk.Token_Type = EOF
	return tk
}

func Categorize_String(lex *ScannerStructure) Token {
	var tk Token
	tk.Token_Type = STRING
	tk.Literal = lex.ReadString()
	return tk
}

const (
	MODULECALL = "::"
)

func (lex *ScannerStructure) Scan_NT() Token {
	lex.ConsumeWhiteSpace()
	if lex.Char == '/' {
		switch lex.Peek() {
		case '/':
			lex.ConsumeComment()
			return lex.Scan_NT()
		case '*':
			lex.ConsumeMultiLineComment()
			return lex.Scan_NT()
		}
	}
	if lex.Char == '!' && lex.Peek() == '-' {
		lex.ConsumeMultiLineComment()
		return lex.Scan_NT()
	}
	var tok Token

	switch lex.Char {
	case '`':
		tok.Token_Type = STRING
		tok.Literal = lex.ReadBacktick()
	case '.':
		if lex.Peek() == '.' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: DOTDOT,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(PERIOD, lex.Char)
		}
	case '&':
		if lex.Peek() == '&' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: ANDAND,
				Literal:    string(ch) + string(lex.Char),
			}
		}
	case '=':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: EQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(ASSIGN, lex.Char)
		}
	case '!':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: NEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(BANG, lex.Char)
		}
	case ';':
		tok = ScanNewToken(SEMICOLON, lex.Char)
	case ':':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: ASSIGN,
				Literal:    string(ch) + string(lex.Char),
			}
		} else if lex.Peek() == ':' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: MODULECALL,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(COLON, lex.Char)
		}
	case '(':
		tok = ScanNewToken(LPAREN, lex.Char)
	case '|':
		if lex.Peek() == '|' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: OROR,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(LINE, lex.Char)
		}
	case ')':
		tok = ScanNewToken(RPAREN, lex.Char)
	case ',':
		tok = ScanNewToken(COMMA, lex.Char)
	case '+':
		tok = ScanNewToken(PLUS, lex.Char)
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: PLUS_EQUALS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(PLUS, lex.Char)
		}
	case '-':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: MINUS_EQUALS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(MINUS, lex.Char)
		}
	case '*':
		if lex.Peek() == '*' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: POWEROF,
				Literal:    string(ch) + string(lex.Char),
			}
		} else if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: ASTERISK_EQUALS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(ASTARISK, lex.Char)
		}
	case '%':
		tok = ScanNewToken(MODULO, lex.Char)
	case '/':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: DIVEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(SLASH, lex.Char)
		}
	case '<':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: LTEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(LT, lex.Char)
		}
	case '>':
		if lex.Peek() == '=' {
			cha := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: GTEQ,
				Literal:    string(cha) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(GT, lex.Char)
		}
	case '{':
		tok = ScanNewToken(LBRACE, lex.Char)
	case '}':
		tok = ScanNewToken(RBRACE, lex.Char)
	case '[':
		tok = ScanNewToken(LBRACKET, lex.Char)
	case ']':
		tok = ScanNewToken(RBRACKET, lex.Char)
	case '"':
		tok.Token_Type = STRING
		tok.Literal = lex.ReadString()
	case 0:
		tok.Literal = ""
		tok.Token_Type = EOF
	default:
		if CharIsDigit(lex.Char) {
			return lex.ReadIntToken()
		}

		if CharIsLetter(lex.Char) {
			tok.Literal = lex.ReadIdentifier()
			tok.Token_Type = LookupIdentifier(tok.Literal)
			lex.PrevTok = tok
			return tok
		}

		tok = ScanNewToken(ILLEGAL, lex.Char)
	}
	lex.ReadChar()
	return tok
}
