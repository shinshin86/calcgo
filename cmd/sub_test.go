package cmd

import "testing"

func TestSubActionSuccess(t *testing.T) {
	arr := []string{"20","1"}
	err := subAction(arr)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestSubActionFailed(t *testing.T) {
	arr := []string{"1","foo"}
	err := subAction(arr)

	if err == nil {
		t.Fatalf("failed test %#v", err)
	}
}