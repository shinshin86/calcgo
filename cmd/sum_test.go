package cmd

import "testing"

func TestSumActionSuccess(t *testing.T) {
	arr := []string{"1","23"}
	err := sumAction(arr)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestSumActionFailed(t *testing.T) {
	arr := []string{"1","foo"}
	err := sumAction(arr)

	if err == nil {
		t.Fatalf("failed test %#v", err)
	}
}