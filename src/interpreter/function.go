package interpreter

import (
	//"errors"
	"fmt"
	"log"

	"github.com/NickSavage/glox/src/parser"
	//	"github.com/NickSavage/glox/src/tokens"
)

func (i *Interpreter) FunctionCall(expr *parser.Expression, arguments []interface{}) (interface{}, RuntimeError) {
	statement, err := i.Memory.Get(expr.FunctionName.Lexeme)
	if err != nil {
		return nil, RuntimeError{
			Message:  fmt.Errorf("undefined function: %v", expr.FunctionName.Lexeme),
			HasError: true,
			Token:    expr.Name,
		}
	}
	var s *parser.Statement
	if statement, ok := statement.(*parser.Statement); !ok {
		return nil, RuntimeError{
			Message:  fmt.Errorf("expected a statement but got: %T", statement),
			HasError: true,
			Token:    expr.Name,
		}
	} else {
		s = statement
	}
	log.Printf("statement %v", s)

	return nil, RuntimeError{}
}
