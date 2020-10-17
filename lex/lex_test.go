package lex_test

import (
	"errors"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/lex"
	"github.com/stretchr/testify/assert"
)

func TestLexCase1(t *testing.T) {
	// Case 1 from problem formulation
	res, _ := lex.NonOrderedColumns([]string{"cba", "daf", "ghi"})
	assert.Equal(t, 1, res)
}

func TestLexCase2(t *testing.T) {
	// Case 2 from problem formulation
	res, _ := lex.NonOrderedColumns([]string{"abcdef"})
	assert.Equal(t, 0, res)
}

func TestLexCase3(t *testing.T) {
	// Case 3 from problem formulation
	res, _ := lex.NonOrderedColumns([]string{"zyx", "wvu", "tsr"})
	assert.Equal(t, 3, res)
}

func TestLexCase4(t *testing.T) {
	// Empty case
	res, _ := lex.NonOrderedColumns([]string{})
	assert.Equal(t, 0, res)
}

func TestLexCase5(t *testing.T) {
	// Rows of variable length
	_, err := lex.NonOrderedColumns([]string{"q", "we"})
	assert.Equal(t, errors.New("strings are of variable length"), err)
}
