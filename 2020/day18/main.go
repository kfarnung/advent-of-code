package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type expressionNumber int64

func newExpressionNumber(str string) expressionNumber {
	value, err := lib.ParseInt32(str)
	if err != nil {
		panic("Couldn't parse number")
	}

	return expressionNumber(value)
}

func (e expressionNumber) String() string {
	return fmt.Sprintf("%d", e)
}

type expressionOperator rune

func (e expressionOperator) String() string {
	val := rune(e)
	return fmt.Sprintf("%c", val)
}

type expressionParenthesis rune

func (e expressionParenthesis) String() string {
	val := rune(e)
	return fmt.Sprintf("%c", val)
}

type isHigherPrecedenceFunction func(expressionOperator, expressionOperator) bool

func isHigherPrecedencePart1(previous, current expressionOperator) bool {
	return true
}

func isHigherPrecedencePart2(previous, current expressionOperator) bool {
	return previous > current
}

func parseExpression(expression string) []interface{} {
	var tokens []interface{}
	var number strings.Builder

	for _, char := range expression {
		if char >= '0' && char <= '9' {
			number.WriteRune(char)
		} else {
			if number.Len() > 0 {
				tokens = append(tokens, newExpressionNumber(number.String()))
				number.Reset()
			}

			if char == '+' || char == '*' {
				operator := expressionOperator(char)
				tokens = append(tokens, operator)
			} else if char == '(' || char == ')' {
				parenthesis := expressionParenthesis(char)
				tokens = append(tokens, parenthesis)
			} else if char != ' ' {
				panic("Unexpected character found")
			}
		}
	}

	if number.Len() > 0 {
		tokens = append(tokens, newExpressionNumber(number.String()))
		number.Reset()
	}

	return tokens
}

func infixToRPN(tokens []interface{}, isHigherPrecedence isHigherPrecedenceFunction) *list.List {
	output := list.New()
	operators := list.New()

	for _, token := range tokens {
		if val, ok := token.(expressionNumber); ok {
			output.PushBack(val)
		} else if val, ok := token.(expressionOperator); ok {
			for operators.Len() > 0 {
				operator := operators.Back()
				if prevVal, ok := operator.Value.(expressionOperator); ok && isHigherPrecedence(prevVal, val) {
					operators.Remove(operator)
					output.PushBack(prevVal)
				} else {
					break
				}
			}

			operators.PushBack(val)
		} else if val, ok := token.(expressionParenthesis); ok {
			if val == '(' {
				operators.PushBack(val)
			} else if val == ')' {
				foundClose := false
				for operators.Len() > 0 {
					operator := operators.Back()
					operators.Remove(operator)

					if val, ok := operator.Value.(expressionParenthesis); ok {
						if val != '(' {
							panic("Unexpected parenthesis")
						}

						foundClose = true
						break
					}

					output.PushBack(operator.Value)
				}

				if !foundClose {
					panic("Couldn't find matching parenthesis")
				}
			} else {
				panic("Unexpected parenthesis value")
			}
		} else {
			panic("Unexpected token type found")
		}
	}

	for operators.Len() > 0 {
		operator := operators.Back()
		operators.Remove(operator)

		output.PushBack(operator.Value)
	}

	return output
}

func doMath(line string, isHigherPrecedence isHigherPrecedenceFunction) int64 {
	tokens := parseExpression(line)
	rpn := infixToRPN(tokens, isHigherPrecedence)
	operands := list.New()

	for current := rpn.Front(); current != nil; current = current.Next() {
		if val, ok := current.Value.(expressionNumber); ok {
			operands.PushBack(val)
		} else if val, ok := current.Value.(expressionOperator); ok {
			operand1 := operands.Back()
			if operand1 == nil {
				panic("Couldn't find operand")
			}
			operands.Remove(operand1)

			operand2 := operands.Back()
			if operand2 == nil {
				panic("Couldn't find operand")
			}
			operands.Remove(operand2)

			switch val {
			case '+':
				operands.PushBack(operand1.Value.(expressionNumber) + operand2.Value.(expressionNumber))

			case '*':
				operands.PushBack(operand1.Value.(expressionNumber) * operand2.Value.(expressionNumber))

			default:
				panic("Unknown operator encountered")
			}
		} else {
			panic("Unexpected token type")
		}
	}

	if operands.Len() != 1 {
		panic("Extra operands left in the stack")
	}

	return int64(operands.Back().Value.(expressionNumber))
}

func part1(lines []string) int64 {
	var total int64
	for _, line := range lines {
		total += doMath(line, isHigherPrecedencePart1)
	}

	return total
}

func part2(lines []string) int64 {
	var total int64
	for _, line := range lines {
		total += doMath(line, isHigherPrecedencePart2)
	}

	return total
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
