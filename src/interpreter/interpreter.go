package interpreter

import (
	"errors"
	"fmt"

	"github.com/NickSavage/glox/src/parser"
	"github.com/NickSavage/glox/src/tokens"
)

func (i *Interpreter) Evaluate(expr *parser.Expression) (interface{}, error) {
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
		return nil, fmt.Errorf("unknown expression type %v", expr.Type)
	}
}

func (i *Interpreter) evaluateGrouping(expr *parser.Expression) (interface{}, error) {
	return i.Evaluate(expr.Expression)
}

func (i *Interpreter) evaluateLiteral(expr *parser.Expression) (interface{}, error) {
	if expr.Type != "Literal" {
		return nil, fmt.Errorf("tried to evaluate a %v as a literal", expr.Type)
	}
	return expr.Value.Literal, nil
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

func (i *Interpreter) evaluateUnary(expr *parser.Expression) (interface{}, error) {
	right, err := i.Evaluate(expr.Right)
	if err != nil {
		return nil, err
	}

	if expr.Operator.Type == tokens.BangToken(0).Type {
		return !i.isTruthy(right), nil
	}
	if expr.Operator.Type == tokens.MinusToken(0).Type {
		switch v := right.(type) {
		case int:
			return -v, nil
		case int64:
			return -v, nil
		case float64:
			return -v, nil
		default:
			return nil, fmt.Errorf("unsuported operand type for unary -: %T", right)
		}

	}
	return nil, errors.New("supposedly unreachable code")
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

func (i *Interpreter) evaluateBinary(expr *parser.Expression) (interface{}, error) {
	left, err := i.Evaluate(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := i.Evaluate(expr.Right)
	if err != nil {
		return nil, err
	}

	leftNumber, leftType, err := i.convertInterfaceNumber(left)
	if err != nil {
		return nil, err
	}
	rightNumber, rightType, err := i.convertInterfaceNumber(right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case tokens.PlusToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) + int(rightNumber), nil
		}
		return leftNumber + rightNumber, nil
	case tokens.MinusToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) - int(rightNumber), nil
		}
		return leftNumber - rightNumber, nil
	case tokens.SlashToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) / int(rightNumber), nil
		}
		return leftNumber / rightNumber, nil
	case tokens.StarToken(0).Type:
		if leftType == "int" && rightType == "int" {
			return int(leftNumber) * int(rightNumber), nil
		}
		return leftNumber * rightNumber, nil
	}
	return nil, errors.New("supposedly unreachable code")
}
