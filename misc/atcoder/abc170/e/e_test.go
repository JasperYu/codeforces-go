// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`6 3
8 1
6 2
9 3
1 1
2 2
1 3
4 3
2 1
1 2`,
			`6
2
6`,
		},
		{
			`2 2
4208 1234
3056 5678
1 2020
2 2020`,
			`3056
4208`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc170/tasks/abc170_e
// https://atcoder.jp/contests/abc170/submit?taskScreenName=abc170_e