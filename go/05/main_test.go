package main

import "testing"

func TestMoveItemToIndex(t *testing.T) {
	list := []string{"a", "b", "c"}
	movePageToIndex(list, 0, 1)
	if list[0] != "b" || list[1] != "a" || list[2] != "c" {
		t.FailNow()
	}
}
