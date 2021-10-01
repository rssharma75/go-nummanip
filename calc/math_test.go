package calc

import ("testing")
func TestAdd(t *testing.T) {
	expect := 6

	err, got := Add(1,2,3)
	if err != nil {
		t.Errorf("Test failed unexpected error")
	}
	if got != expect {
		t.Errorf("Test failed: expected:%v, got:%v", expect, got)
	}

}
