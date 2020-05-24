// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[2,3,1,3,1,null,1]`, 
			`2`, 
		},
		{
			`[2,1,1,1,3,null,null,null,null,null,1]`, 
			`1`, 
		},
		{
			`[9]`, 
			`1`, 
		},
		// TODO 测试参数的下界和上界
		{
			`[1,9,1,null,1,null,1,null,null,7,null,null,4]`,
			`1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, pseudoPalindromicPaths, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-190/problems/pseudo-palindromic-paths-in-a-binary-tree/