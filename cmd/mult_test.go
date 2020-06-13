package cmd

import "testing"

func TestMultActionSuccess(t *testing.T) {
	arr := []string{"20","2"}
	err := subAction(arr)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestMultActionFailed(t *testing.T) {
	arr := []string{"1","foo"}
	err := subAction(arr)

	if err == nil {
		t.Fatalf("failed test %#v", err)
	}
}