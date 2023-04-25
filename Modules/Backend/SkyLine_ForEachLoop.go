package SkyLine_Backend

import (
	"bytes"
	"fmt"
)

const (
	TOKEN_FOREACH = "FOREACH"
	TOKEN_INSIDE  = "IN"
)

type ForeachStatement struct {
	Token Token
	Index string
	Ident string
	Value Expression
	Body  *BlockStatement
}

func (fes *ForeachStatement) EN() {}

// TokenLiteral returns the literal token.
func (fes *ForeachStatement) SL_ExtractNodeValue() string { return fes.Token.Literal }

// String returns this object as a string.
func (fes *ForeachStatement) SL_ExtractStringValue() string {
	var out bytes.Buffer
	out.WriteString("foreach ")
	out.WriteString(fes.Ident)
	out.WriteString(" ")
	out.WriteString(fes.Value.SL_ExtractStringValue())
	out.WriteString(fes.Body.SL_ExtractStringValue())
	return out.String()
}

func (p *Parser) SkyLine_ForEach() Expression {
	expression := &ForeachStatement{Token: p.CurrentToken}
	p.NT()
	expression.Ident = p.CurrentToken.Literal
	if p.PeekTokenIs(COMMA) {
		p.NT()

		if !p.PeekTokenIs(IDENT) {
			p.Errors = append(p.Errors, fmt.Sprintf("second argument to foreach must be ident, got %v", p.PeekToken))
			return nil
		}
		p.NT()

		expression.Index = expression.Ident
		expression.Ident = p.CurrentToken.Literal
	}
	if !p.ExpectPeek(TOKEN_INSIDE) {
		return nil
	}
	p.NT()
	expression.Value = p.SkyLine_Expression(LOWEST)
	if expression.Value == nil {
		return nil
	}
	p.NT()
	expression.Body = p.SkyLine_BlockStatement()
	return expression
}

func SkyLine_Eval_Unit_Foreach(fle *ForeachStatement, env *Environment_of_environment) SLC_Object {
	val := Eval(fle.Value, env)
	helper, ok := val.(Iterable)
	if !ok {
		return NewError("%s object doesn't implement the Iterable interface", val.SL_RetrieveDataType())
	}
	var permit []string
	permit = append(permit, fle.Ident)
	if fle.Index != "" {
		permit = append(permit, fle.Index)
	}
	child := NewTempScop(env, permit)
	helper.Reset()
	ret, idx, ok := helper.Next()
	for ok {
		child.Set(fle.Ident, ret)
		idxName := fle.Index
		if idxName != "" {
			child.Set(fle.Index, idx)
		}
		rt := Eval(fle.Body, child)
		if !isError(rt) && (rt.SL_RetrieveDataType() == ReturnValueType || rt.SL_RetrieveDataType() == ErrorType) {
			return rt
		}
		ret, idx, ok = helper.Next()
	}
	return &Nil{}
}
