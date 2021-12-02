package filter

import "testing"

func TestNewSlice(t *testing.T) {
	first := "1.7.6"
	last := "1.8.6"
	versions := []string{first, "1.7.8", "1.12.0", "1.11.4", "1.11.3", "1.10.5", "1.9.9", last}
	s, _ := MakeSlice(versions)
	if len(s) != len(versions) {
		t.Errorf("New slice failed to construct. Length %v, wanted %v", len(s), len(versions))
	}

	if s[0].Original != first {
		t.Errorf("Mismatched first element in slice. %v. Want %v", s[0].Original, first)
	}

	if s[len(s)-1].Original != last {
		t.Errorf("Mismatched last element in slice. %v. Want %v", s[len(s)-1].Original, first)
	}
}

func TestNewSliceError(t *testing.T) {
	values := []string{"wrong", "invalid", "1.2.3"}
	_, err := MakeSlice(values)
	if err == nil {
		t.Errorf("MakeSlice(%v) did not error. It should have", values)
	}
}
