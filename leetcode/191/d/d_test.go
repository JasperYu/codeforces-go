// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		//{
		//	`[2,2]`,
		//	`1`,
		//},
		//{
		//	`[3,3]`,
		//	`1`,
		//},
		//{
		//	`[2,2,2]`,
		//	`1`,
		//},
		//{
		//	`[2,2,2,2]`,
		//	`1`,
		//},
		//{
		//	`[3,3,3,3]`,
		//	`0.7662337662337663`,
		//},
		{
			`[1,1]`, 
			`1`,
		},
		{
			`[2,1,1]`, 
			`0.6666666666666666`,
		},
		{
			`[1,2,1,2]`, 
			`0.6`,
		},
		{
			`[3,2,1]`, 
			`0.3`,
		},
		{
			`[6,6,6,6,6,6]`, 
			`0.90327`,
		},
		// TODO 测试参数的下界和上界

	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, getProbability, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-191/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/