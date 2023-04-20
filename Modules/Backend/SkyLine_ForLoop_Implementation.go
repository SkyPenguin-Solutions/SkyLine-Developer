package SkyLine_Backend

import (
	"bytes"
)

/// FOR LOOP IMPLEMENTATIOM

const (
	STD_FOR = "FOR"
)

type ForLoopExpression struct {
	Token       Token
	Condition   Expression
	Consequence *BlockStatement
}

func (Floop *ForLoopExpression) EN()                         {}
func (Floop *ForLoopExpression) SL_ExtractNodeValue() string { return Floop.Token.Literal }

func (Floop *ForLoopExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString("for (")
	Out.WriteString(Floop.Condition.SL_ExtractStringValue())
	Out.WriteString(") {")
	Out.WriteString(Floop.Consequence.SL_ExtractStringValue())
	Out.WriteString("}")
	return Out.String()
}

func EvalUateForLoopFunction(loop *ForLoopExpression, Env *Environment_of_environment) SLC_Object {
	Val := &Boolean_Object{Value: true}
	for {
		condition := Eval(loop.Condition, Env)
		if isError(condition) {
			return condition
		}
		if isTruthy(condition) {
			newval := Eval(loop.Consequence, Env)
			if !isError(newval) && (newval.SL_RetrieveDataType() == ReturnValueType || newval.SL_RetrieveDataType() == ErrorType) {
				return newval
			}
		} else {
			break
		}
	}
	return Val
}

func (parser *Parser) ParseForLoop() Expression {
	expression := &ForLoopExpression{
		Token: parser.CurrentToken,
	}
	if !parser.ExpectPeek(LPAREN) {
		return nil
	}
	parser.NextLoadFaultToken()
	expression.Condition = parser.parseExpression(LOWEST)
	if !parser.ExpectPeek(RPAREN) {
		return nil
	}
	if !parser.ExpectPeek(LBRACE) {
		return nil
	}
	expression.Consequence = parser.parseBlockStatement()
	return expression
}
