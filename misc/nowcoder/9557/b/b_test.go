// Code generated by copypasta/template/nowcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`"1#1#+"`,
			`2`,
		},
		{
			`"12#3#+15#*"`,
			`225`,
		},
		// TODO 测试参数的下界和上界
		{
			`"3#4#-"`,
			`-1`,
		},
		{
			`"1#"`,
			`1`,
		},
		{
			`"3#4#+5#*6#-"`,
			`29`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, solve, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/9557/b