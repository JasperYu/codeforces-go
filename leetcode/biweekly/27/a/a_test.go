// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`[1,2,3,4]`, `[2,4,1,3]`, 
			`true`, 
		},
		{
			`[7]`, `[7]`, 
			`true`, 
		},
		{
			`[1,12]`, `[12,1]`, 
			`true`, 
		},
		{
			`[3,7,9]`, `[3,7,11]`, 
			`false`, 
		},
		{
			`[1,1,1,1,1]`, `[1,1,1,1,1]`, 
			`true`, 
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, canBeEqual, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-27/problems/make-two-arrays-equal-by-reversing-sub-arrays/