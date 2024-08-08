package interpreter

import (
	//"errors"
	"fmt"
	"log"

	"github.com/NickSavage/glox/src/parser"
	//	"github.com/NickSavage/glox/src/tokens"
)

func (i *Interpreter) executeReturn(statement *parser.Statement) RuntimeError {

	result, rerr := i.Evaluate(statement.Expression)
	if rerr.HasError {
		log.Printf("rerr %v", rerr.Message.Error())
		return rerr
	}
	return RuntimeError{
		HasError:    false,
		Return:      true,
		ReturnValue: result,
	}

}

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
	previous := i.Memory
	i.Memory = &Storage{
		Memory:       make(map[string]interface{}),
		Enclosing:    previous,
		HasEnclosing: true,
	}
	for index, _ := range arguments {
		i.Memory.Define(s.Parameters[index].Lexeme, arguments[index])
	}
	rerr := i.executeBlock(s)
	i.Memory = previous
	if rerr.HasError {
		return nil, rerr
	}
	if rerr.Return {
		return rerr.ReturnValue, RuntimeError{}
	}

	return nil, RuntimeError{}
}
