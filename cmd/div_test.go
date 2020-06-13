package cmd

import "testing"

func TestDivActionSuccess(t *testing.T) {
	arr := []string{"10","5"}
	err := sumAction(arr)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestDivActionFailed(t *testing.T) {
	arr := []string{"1","foo"}
	err := sumAction(arr)

	if err == nil {
		t.Fatalf("failed test %#v", err)
	}
}