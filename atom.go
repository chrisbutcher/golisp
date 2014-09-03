package main

import (
	"fmt"
	"strconv"
)

type atomEnumType int

const (
	atomSymbol atomEnumType = iota
	atomBuiltin
	atomFloat
	atomQuote
	atomLambda
	atomBoolean
)

type Atom struct {
	typ        atomEnumType
	val        T
	valNum     float64
	lambdaName string
	lambdaArgs []Atom
	lambdaFn   T
}

func atomTypeToString(t atomEnumType) string {
	mapping := map[atomEnumType]string{atomSymbol: "s", atomBuiltin: "fn", atomFloat: "fl", atomQuote: "qt", atomLambda: "lm"}

	return mapping[t]
}

func (a Atom) String() string {
	if a.typ == atomLambda {
		return fmt.Sprintf("<%s:args(%s):fn(%s)>", atomTypeToString(a.typ), a.lambdaArgs, a.lambdaFn)
	} else if a.typ == atomQuote {
		return fmt.Sprintf("%v", a.val)
	} else {
		return fmt.Sprintf("<%s:%s>", atomTypeToString(a.typ), a.val)
	}
}

func (a Atom) BooleanValue() bool {
	return a.val.(bool)
}

func floatToString(f float64) string {
	return fmt.Sprintf("%f", f)
}

func createAtom(val T) Atom {
	switch val.(type) {
	case Atom:
		return val.(Atom)
	default:
		value := val.(string)
		new_atom := Atom{}

		float_value, err := strconv.ParseFloat(value, 64)

		new_atom.val = value
		if err != nil {
			if value == "lambda" {
				new_atom.typ = atomLambda
			} else if value == "true" {
				new_atom.typ = atomBoolean
				new_atom.val = true
			} else if value == "false" {
				new_atom.typ = atomBoolean
				new_atom.val = false
			} else {
				new_atom.typ = atomSymbol
			}
		} else {
			new_atom.typ = atomFloat
			new_atom.valNum = float_value
		}

		return new_atom
	}
}

func genericToAtomSlice(input T) []Atom {
	result := make([]Atom, 0)

	switch input.(type) {
	case Expr:
		slice := input.(Expr)
		for _, item := range slice {
			result = append(result, item.(Atom))
		}
	case Atom:
		result = append(result, input.(Atom))
	}

	return result
}

func genericToExpr(input T) Expr {
	result := make(Expr, 0)

	switch input.(type) {
	case Expr:
		slice := input.(Expr)
		for _, item := range slice {
			result = append(result, item.(Atom))
		}
	case Atom:
		result = append(result, input.(Atom))
	}

	return result
}

func populateLambda(input Expr) Atom {
	return Atom{typ: atomLambda, val: "lambda", lambdaArgs: genericToAtomSlice(input[1]), lambdaFn: input[2].(Expr)}
}

func exprIsConditional(input T) bool {
	switch input.(type) {
	case Atom:
		input_atom := input.(Atom)
		if input_atom.typ == atomBuiltin {
			if input_atom.val == "if" {
				return true
			} else {
				return false
			}
		}
	default:
		return false
	}

	return false
}

func exprIsBoolean(input T) bool {
	switch input.(type) {
	case Atom:
		return input.(Atom).typ == atomBoolean
	default:
		return false
	}
}

func exprIsLambda(input T) bool {
	switch input.(type) {
	case Atom:
		return input.(Atom).typ == atomLambda
	default:
		return false
	}
}
