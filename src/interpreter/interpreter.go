package interpreter

import (
	"errors"
	"fmt"

	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

func (i *Interpreter) Evaluate(expr *parser.Expression) (interface{}, RuntimeError) {
	switch expr.Type {
	case "Literal":
		return i.evaluateLiteral(expr)
	case "Unary":
		return i.evaluateUnary(expr)
	case "Binary":
		return i.evaluateBinary(expr)
	case "Grouping":
		return i.evaluateGrouping(expr)
	default:
		return nil, RuntimeError{
			Message:  fmt.Errorf("unknown expression type %v", expr.Type),
			HasError: true,
			Token:    expr.Operator,
		}
	}
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
