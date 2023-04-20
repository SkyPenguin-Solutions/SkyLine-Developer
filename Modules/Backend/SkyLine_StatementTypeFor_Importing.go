package SkyLine_Backend

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

var SearchPaths []string

func (p *Parser) ParseImportExpression() Expression {
	expression := &ImportExpression{Token: p.CurrentToken}

	if !p.ExpectPeek(LPAREN) {
		return nil
	}

	p.NT()
	expression.Name = p.parseExpression(LOWEST)
	if !p.ExpectPeek(RPAREN) {
		return nil
	}

	return expression
}

func (iexp *ImportExpression) EN()                         {}
func (iexp *ImportExpression) SL_ExtractNodeValue() string { return iexp.Token.Literal }

func (iexp *ImportExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString(iexp.SL_ExtractNodeValue())
	Out.WriteString("(")
	Out.WriteString(fmt.Sprintf("\"%s\"", iexp.Name))
	return Out.String()
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func FindModule(name string) string {
	basename := fmt.Sprintf("%s.skyline", name)
	for _, p := range SearchPaths {
		filename := filepath.Join(p, basename)
		if Exists(filename) {
			fmt.Println("[=] Found module")
			return filename
		}
	}
	return ""
}

// IMPORTING MUST BE UPPER CASE OR START WITH AN UPPER CASE LETTER
func (e *Environment_of_environment) ExportedHash() *Hash {
	pairs := make(map[HashKey]HashPair)
	for k, v := range e.Store {
		if unicode.IsUpper(rune(k[0])) {
			s := &String{Value: k}
			pairs[s.HashKey()] = HashPair{Key: s, Value: v}
		}
	}
	return &Hash{Pairs: pairs}
}

func (p *Parser) parseSelectorExpression(exp Expression) Expression {
	p.ExpectPeek(IDENT)
	index := &StringLiteral{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
	return &IndexExpression{Left: exp, Index: index}
}

//import("main.skyline")::NewFunction()
//set modname = import("main.skyline");
