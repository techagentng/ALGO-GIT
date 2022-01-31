package main

import (
	"github.com/techagentng/testPrac/testPackage"
	"testing"
)
// Testing my average function
func TestMyAverage(t *testing.T){
	s:= []int{1,2,3,4,45,6,7,8,67,4}
	x := testPackage.MyAverage(s...)
 if x != 14 {
 	t.Errorf("Expected... got %v", x)
 }
}

// Testing my average function
func TestMySum(t *testing.T) {
	sl:= []int{1,2,3,4,45,6,7,8,67,4}
	w:= testPackage.MySum(sl...)
	if w != 147 {
		t.Errorf("Expected sum 147, got %v", w)
	}
}