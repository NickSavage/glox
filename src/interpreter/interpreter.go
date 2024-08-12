package interpreter

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

func (i *Interpreter) executePrint(text interface{}) RuntimeError {
	print(fmt.Sprintf("%v", text))
	print("\n")
	return RuntimeError{}
}

func (i *Interpreter) executeBreak() RuntimeError {
	if !i.InLoop {
		return RuntimeError{
			Message:  errors.New("Cannot call break outside a for loop"),
			HasError: true,
		}
	}
	i.BreakTriggered = true
	return RuntimeError{}
}
func (i *Interpreter) executeContinue() RuntimeError {
	if !i.InLoop {
		return RuntimeError{
			Message:  errors.New("Cannot call continue outside a for loop"),
			HasError: true,
		}
	}
	i.ContinueTriggered = true
	return RuntimeError{}
}

func (i *Interpreter) executeVariable(statement *parser.Statement) RuntimeError {
	value, err := i.Evaluate(statement.Initializer)
	if err.HasError {
		log.Printf("err %v", err)
		return err
	}
	i.Memory.Define(statement.VariableName.Lexeme, value)
	return RuntimeError{HasError: false}
}

func (i *Interpreter) executeBlock(statement *parser.Statement) RuntimeError {
	for _, statement := range statement.Statements {
		rerr := i.Execute(statement)
		if rerr.HasError || rerr.Return {
			return rerr
		}
	}
	return RuntimeError{}

}

func (i *Interpreter) Execute(statement *parser.Statement) RuntimeError {
	switch statement.Type.Type {
	case "Block":
		previous := i.Memory
		i.Memory = &Storage{
			Memory:       make(map[string]interface{}),
			Enclosing:    previous,
			HasEnclosing: true,
		}
		i.executeBlock(statement)
		i.Memory = previous
	case "Break":
		return i.executeBreak()
	case "Continue":
		return i.executeContinue()
	case "If":
		result, rerr := i.Evaluate(statement.Condition)
		if rerr.HasError {
			return rerr
		}
		if i.isTruthy(result) {
			for _, statement := range statement.Statements {
				rerr := i.Execute(statement)
				if rerr.HasError || rerr.Return {
					return rerr

				}

			}
		} else {
			for _, statement := range statement.ElseStatements {
				rerr := i.Execute(statement)
				if rerr.HasError || rerr.Return {
					return rerr

				}

			}
		}
	case "For":
		statements := statement.Statements
		i.InLoop = true
		for {
			for _, statement = range statements {
				rerr := i.Execute(statement)
				if i.BreakTriggered {
					i.InLoop = false
					i.BreakTriggered = false
					break
				}
				if i.ContinueTriggered {
					i.ContinueTriggered = false
					break
				}
				if rerr.HasError {
					return rerr
				}
			}
			if !i.InLoop {
				break
			}
		}
	case "Function":
		i.Memory.Define(statement.FunctionName.Lexeme, statement)
	case "Print":
		result, rerr := i.Evaluate(statement.Expression)
		if rerr.HasError {
			log.Printf("rerr %v", rerr)
			log.Printf("? %v", rerr.Message.Error())
			return rerr
		}
		switch v := result.(type) {
		case string:
			return i.executePrint(result)
		default:
			return i.executePrint(fmt.Sprintf("%v", v))
		}
	case "Return":
		return i.executeReturn(statement)
	case "Variable":
		return i.executeVariable(statement)

	case "Expression":
		value, rerr := i.Evaluate(statement.Expression)

		if rerr.HasError {
			log.Printf("rerr %v", rerr)
			log.Printf("? %v", rerr.Message.Error())
			return rerr
		}
		if statement.Expression.Type == "Assignment" {
			err := i.Memory.Assign(statement.Expression.Name.Lexeme, value)
			if err != nil {
				return RuntimeError{
					Message:  err,
					HasError: true,
					Token:    statement.Expression.Name,
				}
			}
		}
	}

	return RuntimeError{}
}

func (i *Interpreter) Evaluate(expr *parser.Expression) (interface{}, RuntimeError) {
	switch expr.Type {
	case "Literal":
		return i.evaluateLiteral(expr)
	case "Unary":
		return i.evaluateUnary(expr)
	case "Binary":
		return i.evaluateBinary(expr)
	case "Function":
		return i.evaluateFunctionCall(expr)
	case "Grouping":
		return i.evaluateGrouping(expr)
	case "Identifier":
		return i.evaluateIdentifier(expr)
	case "Assignment":
		return i.evaluateAssignment(expr)
	case "Variable":
		return i.evaluateVariable(expr)
	case "Logical":
		return i.evaluateLogical(expr)
	case "Lambda":
		return i.evaluateLambda(expr)
	case "Array":
		return i.evaluateArray(expr)
	case "Element":
		return i.evaluateElement(expr)
	default:
		return nil, RuntimeError{
			Message:  fmt.Errorf("unknown expression type %v", expr.Type),
			HasError: true,
			Token:    expr.Operator,
		}
	}
}

func (i *Interpreter) evaluateArray(expr *parser.Expression) (interface{}, RuntimeError) {
	return expr.Array, RuntimeError{HasError: false}

}

func (i *Interpreter) evaluateElement(expr *parser.Expression) (interface{}, RuntimeError) {
	result, err := i.Memory.Get(expr.Name.Lexeme)
	if err != nil {
		return nil, RuntimeError{
			Message:  err,
			HasError: true,
			Token:    expr.Name,
		}

	}
	var array *parser.Array
	array, ok := result.(*parser.Array)
	if !ok {
		return nil, RuntimeError{
			Message:  fmt.Errorf("tried to index a non-array, %v", expr.Name.Lexeme),
			HasError: true,
			Token:    expr.Name,
		}
	}

	element := array.Elements[expr.Index].Literal
	return element, RuntimeError{}
}

func (i *Interpreter) evaluateVariable(expr *parser.Expression) (interface{}, RuntimeError) {
	result, err := i.Memory.Get(expr.Name.Lexeme)
	if err != nil {
		return nil, RuntimeError{
			Message:  err,
			HasError: true,
			Token:    expr.Name,
		}
	}
	return result, RuntimeError{}
}

func (i *Interpreter) evaluateIdentifier(expr *parser.Expression) (interface{}, RuntimeError) {
	return expr.Value, RuntimeError{HasError: false}
}

func (i *Interpreter) evaluateAssignment(expr *parser.Expression) (interface{}, RuntimeError) {
	return i.Evaluate(expr.AssignValue)
}

func (i *Interpreter) evaluateGrouping(expr *parser.Expression) (interface{}, RuntimeError) {
	return i.Evaluate(expr.Expression)
}

func (i *Interpreter) evaluateLiteral(expr *parser.Expression) (interface{}, RuntimeError) {
	if expr.Type != "Literal" {
		return nil, RuntimeError{
			Message:  fmt.Errorf("tried to evaluate a %v as a literal", expr.Type),
			HasError: true,
			Token:    expr.Value,
		}
	}
	return expr.Value.Literal, RuntimeError{}
}

func (i *Interpreter) isTruthy(object interface{}) bool {
	if object == nil {
		return false
	}

	switch v := object.(type) {
	case bool:
		return v
	case int, int64, float64, string:
		// Non-zero numbers and non-empty strings are truthy
		return true
	default:
		// All other non-nil values are considered truthy
		return true
	}
}

func (i *Interpreter) isEqual(a, b interface{}) bool {
	if a == nil {
		return false
	}
	if a == nil && b == nil {
		return true
	}
	return a == b
}

func (i *Interpreter) evaluateLogical(expr *parser.Expression) (interface{}, RuntimeError) {
	left, rerr := i.Evaluate(expr.Left)
	if rerr.HasError {
		return nil, rerr
	}
	if expr.Operator.Type.Type == "Or" {
		if i.isTruthy(left) {
			return left, RuntimeError{}
		}
	} else {
		if !i.isTruthy(left) {
			return left, RuntimeError{}
		}

	}
	return i.Evaluate(expr.Right)
}

func (i *Interpreter) evaluateUnary(expr *parser.Expression) (interface{}, RuntimeError) {
	var rerr RuntimeError
	right, rerr := i.Evaluate(expr.Right)
	if rerr.HasError {
		return nil, rerr
	}

	rerr = RuntimeError{}
	if expr.Operator.Type == tokens.BangToken(0).Type {
		return !i.isTruthy(right), rerr
	}
	if expr.Operator.Type == tokens.MinusToken(0).Type {
		switch v := right.(type) {
		case int:
			return -v, rerr
		case int64:
			return -v, rerr
		case float64:
			return -v, rerr
		default:
			return nil, RuntimeError{
				Message:  fmt.Errorf("unsuported operand type for unary -: %T", right),
				HasError: true,
				Token:    expr.Operator,
			}
		}

	}
	return nil, RuntimeError{
		Message:  errors.New("supposedly unreachable code"),
		HasError: true,
		Token:    expr.Operator,
	}
}

func (i *Interpreter) convertInterfaceNumber(literal interface{}) (float64, string, error) {
	switch v := literal.(type) {
	case int:
		return float64(v), "int", nil
	case int64:
		return float64(v), "int64", nil
	case float64:
		return v, "float64", nil
	default:
		return 0, "", fmt.Errorf("unsupported operand type for number conversion: %T", literal)
	}
}

func (i *Interpreter) evaluateConcat(expr *parser.Expression, left, right interface{}) (interface{}, RuntimeError) {
	a, ok := left.(string)
	if !ok {
		return nil, RuntimeError{
			Message: fmt.Errorf("can only concatenate strings together, got %v", left),
			HasError: true,
			Token: expr.Left.Value,
		}
	}
	b, ok := right.(string)
	if !ok {
		return nil, RuntimeError{
			Message: fmt.Errorf("can only concatenate stings together, got %v", right),
			HasError: true,
			Token: expr.Right.Value,
		}
	}
	var sb strings.Builder
	sb.WriteString(a)
	sb.WriteString(b)
	log.Printf("%q, %q, %v",a, b, sb.String())
	return sb.String(), RuntimeError{}
}

func (i *Interpreter) evaluateBinary(expr *parser.Expression) (interface{}, RuntimeError) {
	var rerr RuntimeError
	var err error
	left, rerr := i.Evaluate(expr.Left)
	if rerr.HasError {
		return nil, rerr
	}
	right, rerr := i.Evaluate(expr.Right)
	if rerr.HasError {
		return nil, rerr
	}

	rerr = RuntimeError{
		HasError: false,
	}
	if expr.Operator.Type == tokens.TildeToken(0).Type {
		return i.evaluateConcat(expr, left, right)
	}
	leftNumber, leftType, err := i.convertInterfaceNumber(left)
	if err != nil {
		return nil, RuntimeError{
			Message:  err,
			HasError: true,
			Token:    expr.Left.Operator,
		}
	}
	rightNumber, rightType, err := i.convertInterfaceNumber(right)
	if err != nil {
		return nil, RuntimeError{
			Message:  err,
			HasError: true,
			Token:    expr.Right.Operator,
		}
	}

	switch expr.Operator.Type {
	case tokens.BangEqualToken(0).Type:
		return !i.isEqual(leftNumber, rightNumber), rerr
	case tokens.EqualEqualToken(0).Type:
		return i.isEqual(leftNumber, rightNumber), rerr

	case tokens.GreaterToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) > int(rightNumber), rerr
		}
		return leftNumber > rightNumber, rerr
	case tokens.GreaterEqualToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) >= int(rightNumber), rerr
		}
		return leftNumber >= rightNumber, rerr
	case tokens.LessToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) < int(rightNumber), rerr
		}
		return leftNumber > rightNumber, rerr
	case tokens.LessEqualToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) <= int(rightNumber), rerr
		}
		return leftNumber <= rightNumber, rerr

	case tokens.PlusToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) + int(rightNumber), rerr
		}
		return leftNumber + rightNumber, rerr
	case tokens.MinusToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) - int(rightNumber), rerr
		}
		return leftNumber - rightNumber, rerr
	case tokens.SlashToken(0).Type:
		if int(rightNumber) == 0 {
			rerr.HasError = true
			rerr.Message = errors.New("dividing by zero")
			rerr.Token = expr.Right.Operator
			return nil, rerr
		}
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) / int(rightNumber), rerr
		}
		return leftNumber / rightNumber, rerr
	case tokens.StarToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) * int(rightNumber), rerr
		}
		return leftNumber * rightNumber, rerr
	}
	rerr.HasError = true
	rerr.Message = errors.New("supposedly unreachable code")
	rerr.Token = expr.Left.Operator
	return nil, rerr
}
