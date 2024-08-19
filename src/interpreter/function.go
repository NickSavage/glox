package interpreter

import (
	//"errors"
	"fmt"
	"log"

	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
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
func (i *Interpreter) evaluateFunctionCall(expr *parser.Expression) (interface{}, RuntimeError) {
	arguments := make([]interface{}, 0)
	for _, a := range expr.Arguments {
		argument, rerr := i.Evaluate(a)
		if rerr.HasError {
			return nil, rerr
		}
		arguments = append(arguments, argument)
	}
	results, rerr := i.FunctionCall(expr, arguments)
	if rerr.HasError {
		return nil, rerr
	}
	return results, RuntimeError{}

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
	s, ok := statement.(*parser.Statement)
	if !ok {
		return nil, RuntimeError{
			Message:  fmt.Errorf("expected a statement but got: %T", statement),
			HasError: true,
			Token:    expr.Name,
		}
	}
	if len(s.Parameters) != len(arguments) {
		return nil, RuntimeError{
			Message:  fmt.Errorf("wrong number of arguments for %v, got %v want %v", expr.FunctionName.Lexeme, len(s.Parameters), len(arguments)),
			HasError: true,
			Token:    expr.FunctionName,
		}
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
	var rerr RuntimeError
	if s.Type.Type == "NativeFunction" {
		log.Printf("s %v", s)
		rerr = i.executeNativeFunction(s.NativeFunction)
	} else {
		rerr = i.executeBlock(s)
	}
	i.Memory = previous
	if rerr.HasError {
		return nil, rerr
	}
	if rerr.Return {
		return rerr.ReturnValue, RuntimeError{}
	}

	return nil, RuntimeError{}
}

func (i *Interpreter) executeNativeFunction(nativeFunction func(i interface{}) (result interface{}, err error)) RuntimeError {
	result, err := nativeFunction(i)
	if err != nil {
		return RuntimeError{
			HasError: true,
			Token:    tokens.Token{},
			Message:  err,
		}
	}
	return RuntimeError{
		HasError:    false,
		ReturnValue: result,
		Return:      true,
	}

}

func (i *Interpreter) evaluateLambda(expr *parser.Expression) (interface{}, RuntimeError) {
	return expr.Lambda, RuntimeError{}
}
