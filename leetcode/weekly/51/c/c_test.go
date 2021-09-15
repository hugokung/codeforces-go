// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[[1,2], [1,3], [2,3]]`,
			`[2,3]`,
		},
		{
			`[[1,2], [2,3], [3,4], [1,4], [1,5]]`,
			`[1,4]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, findRedundantDirectedConnection, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// 684 https://leetcode.com/contest/leetcode-weekly-contest-51/problems/redundant-connection/
// 685 https://leetcode-cn.com/problems/redundant-connection-ii/
// 周赛是 684，这里写的是 685