package calc

import "testing"

func TestCat(t *testing.T) {
	ret := Add(3, 5)
	if ret != 8 {
		t.Errorf("3 + 5 should = 8, but get %v", ret)
	}
}
