package main

import (
	"testing"
)

func TestWalkFileMap(t *testing.T) {
	dirMap, err := walkFileMap(".")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dirMap)
}
