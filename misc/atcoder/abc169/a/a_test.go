// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`2 5`,
			`10`,
		},
		{
			`100 100`,
			`10000`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc169/tasks/abc169_a
// https://atcoder.jp/contests/abc169/submit?taskScreenName=abc169_a