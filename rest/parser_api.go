package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"cadence-parser/ast"
	"cadence-parser/parser"
)

type benchResult struct {
	// N is the the number of iterations
	Iterations int `json:"iterations"`
	// T is the total time taken
	Time time.Duration `json:"time"`
}

type result struct {
	Path     string       `json:"path,omitempty"`
	Code     string       `json:"-"`
	Bench    *benchResult `json:"bench,omitempty"`
	BenchStr string       `json:"-"`
	Error    error        `json:"error,omitempty"`
	Program  *ast.Program `json:"program"`
}

func encodeResult(res result) (res_str []byte) {
	res_str, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return
}

func encodeResult2(res result) (res_str []byte) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(res)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func Parse(code []byte) (res result, succeeded bool) {
	res = result{
		Code: string(code),
	}
	succeeded = true

	var program *ast.Program
	var err error

	func() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%s", debug.Stack())
				res.Error = err
			}
		}()

		program, err = parser.ParseProgram(string(code), nil)
		res.Program = program
		res.Error = err
	}()

	if err != nil {
		succeeded = false
		return
	}

	return
}

func Parse2Json(code []byte) (res_json []byte) {
	res, succ := Parse(code)
	if !succ {
		return
	}
	res_json = encodeResult(res)
	return
}
