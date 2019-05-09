package slices

import (
	"testing"
)

func TestContains(t *testing.T) {

	in := []string{"a", "b", "c", "zzabyycdxx", "aba", "中文字符"}

	out := Contains(in, "d")
	if out == true {
		t.Error(out)
	}

	out = Contains(in, "a")
	if out != true {
		t.Error(out)
	}

	out = Contains(in, "zzabyycdxx")
	if out != true {
		t.Error(out)
	}

	out = Contains(in, "abac")
	if out == true {
		t.Error(out)
	}

	out = Contains(in, "中文字符")
	if out != true {
		t.Error(out)
	}

	out = Contains(in, "中文")
	if out == true {
		t.Error(out)
	}
}
