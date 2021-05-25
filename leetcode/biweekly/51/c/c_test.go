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
			`[2,2,1,2,1]`, 
			`2`,
		},
		{
			`[100,1,1000]`, 
			`3`,
		},
		{
			`[1,2,3,4,5]`, 
			`5`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, maximumElementAfterDecrementingAndRearranging, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-51/problems/maximum-element-after-decreasing-and-rearranging/