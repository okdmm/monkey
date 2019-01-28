package evaluator

import (
	"github.com/okdmm/monkey/ast"
	"github.com/okdmm/monkey/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	defitinitions := []int{}

	for i, statement := range program.Statements {
		if isMacroDefitinition(statement) {
			addMacro(statement, env)
			defitinitions = append(defitinitions, i)
		}
	}

	for i := len(defitinitions) - 1; i >= 0; i = i - 1 {
		defitinitionIndex := defitinitions[i]
		program.Statements = append(
			program.Statements[:defitinitionIndex],
			program.Statements[defitinitionIndex+1:]...,
		)
	}
}

func isMacroDefitinition(node ast.Statement) bool {
	letStatement, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStatement.Value.(*ast.MacroLiteral)
	if !ok {
		return false
	}

	return true
}

func addMacro(stmt ast.Statement, env *object.Environment) {
	letStatement, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStatement.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Env:        env,
		Body:       macroLiteral.Body,
	}

	env.Set(letStatement.Name.Value, macro)
}
