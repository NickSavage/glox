package parser

func prettyPrintExpressionTree(input *Expression, result string) string {
	if input.Value.Lexeme != "" {
		result += input.Value.Lexeme
		return result
	}

	result += "("
	if input.Operator.Lexeme != "" {
		result += input.Operator.Lexeme + " "
	} else if input.Type == "Grouping" {
		result += "group" + " "
	}
	if input.Expression != nil {
		result = prettyPrintExpressionTree(input.Expression, result)
	}
	if input.Left != nil && input.Left.Type != "" {
		result = prettyPrintExpressionTree(input.Left, result) + " "
	}
	if input.Right != nil && input.Right.Type != "" {
		result = prettyPrintExpressionTree(input.Right, result)
	}
	result += ")"
	return result
}
