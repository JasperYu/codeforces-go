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
			`[2,5,1,3,4,7]`, `3`, 
			`[2,3,5,4,1,7]`,
		},
		{
			`[1,2,3,4,4,3,2,1]`, `4`, 
			`[1,4,2,3,3,2,4,1]`,
		},
		{
			`[1,1,2,2]`, `2`, 
			`[1,2,1,2]`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, shuffle, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-192/problems/shuffle-the-array/