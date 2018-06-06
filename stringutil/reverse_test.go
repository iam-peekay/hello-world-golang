package stringutil

import "testing"

func TestReverse(t *testing.T) {
	if Reverse("hello") != "olleh" {
		t.Log("Reverse of \" hello \" is -> olleh ")
		t.Fail()
	}
}
