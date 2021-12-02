package version

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	// Doing a table test, but only because we need to trigger deeper functions
	tables := []struct {
		name    string
		version string
		err     bool
		output  *Version
	}{
		{"OK Constructor", "1.2.3", false, &Version{Original: "1.2.3", Segments: []int{1, 2, 3}}},
		{"Bad Constructor", "1.2-alpha", true, nil},
	}

	for _, tc := range tables {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			v, err := New(tc.version)
			if tc.err && err == nil {
				t.Errorf("Expected error, did not get.")
			} else if !tc.err && !reflect.DeepEqual(v, tc.output) {
				t.Errorf("Got %v, want %v", v, tc.output)
			}
		})
	}
}

func TestVersionLess(t *testing.T) {
	tables := []struct {
		Name   string
		Base   *Version
		Comp   *Version
		Expect bool
	}{
		{"Simple equal", &Version{"1.2.3", []int{1, 2, 3}}, &Version{"1.2.3", []int{1, 2, 3}}, false},
		{"Simple less", &Version{"1.2.3", []int{1, 2, 3}}, &Version{"1.2.4", []int{1, 2, 4}}, true},
		{"Simple less 2", &Version{"1.2.3", []int{1, 2, 3}}, &Version{"1.11.0", []int{1, 11, 0}}, true},
		{"Simple less 3", &Version{"1.2.3", []int{1, 2, 3}}, &Version{"9.3.0", []int{9, 3, 0}}, true},
		{"Simple more", &Version{"1.11.0", []int{1, 11, 0}}, &Version{"1.2.3", []int{1, 2, 3}}, false},
		{"Complex less", &Version{"1.2.3", []int{1, 2, 3}}, &Version{"2.0", []int{2, 0}}, true},
		{"Complex more", &Version{"12.3.4", []int{12, 3, 4}}, &Version{"1.0", []int{1, 0}}, false},
	}

	for _, tc := range tables {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			res := tc.Base.Less(tc.Comp)
			if res != tc.Expect {
				t.Errorf("%v.Less(%v) == %v. Want %v", tc.Base, tc.Comp, res, tc.Expect)
			}
		})
	}
}

func TestExplodeVersion(t *testing.T) {
	tables := []struct {
		name      string
		version   string
		expect    []int
		expectErr bool
	}{
		{"Normal OK", "1.2.3", []int{1, 2, 3}, false},
		{"Short OK", "1.2", []int{1, 2}, false},
		{"Invalid Non-Digits", "1.2-alpha", []int{}, true},
		{"Legit Istio version", "1.7.6", []int{1, 7, 6}, false},
		{"Legit Istio version longer", "1.12.1", []int{1, 12, 1}, false},
	}

	for _, tc := range tables {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := explodeVersion(tc.version)
			if tc.expectErr && err == nil {
				t.Errorf("Expected error, did not get one. %v", res)
			} else if !tc.expectErr && !reflect.DeepEqual(res, tc.expect) {
				t.Errorf("Got %v; wanted %v", res, tc.expect)
			}
		})
	}
}

func TestVersionRepr(t *testing.T) {
	tables := []struct {
		Version *Version
		Expect  string
	}{
		{&Version{"1.2.3", []int{}}, "1.2.3"},
		{&Version{"1.2.3", []int{}}, "1.2.3"},
		{&Version{"1.2.3", []int{}}, "1.2.3"},
		{&Version{"1.2.3", []int{}}, "1.2.3"},
	}

	for _, tc := range tables {
		tc := tc
		t.Run(tc.Expect, func(t *testing.T) {
			r := tc.Version.Repr()
			if r != tc.Expect {
				t.Errorf("v.Repr() = %v. Want %v", r, tc.Expect)
			}
		})
	}

}

func TestSliceLen(t *testing.T) {
	tables := []struct {
		Name  string
		Slice VersionSlice
		Len   int
	}{
		{"Empty Slice", VersionSlice{}, 0},
		{"Filled slice", VersionSlice{&Version{}, &Version{}}, 2},
	}

	for _, tc := range tables {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			n := tc.Slice.Len()
			if n != tc.Len {
				t.Errorf("%v.Len() == %v. Want %v.", tc.Slice, n, tc.Len)
			}
		})
	}
}

func TestSliceLess(t *testing.T) {
	slice := VersionSlice{
		&Version{"1.2.3", []int{1, 2, 3}},
		&Version{"1.2.3", []int{1, 2, 3}},
		&Version{"1.2.4", []int{1, 2, 4}},
		&Version{"1.11.0", []int{1, 11, 0}},
		&Version{"9.3.0", []int{9, 3, 0}},
		&Version{"12.3.4", []int{12, 3, 4}},
		&Version{"2.0", []int{2, 0}},
	}

	tables := []struct {
		Name     string
		Slice    VersionSlice
		Compare1 int
		Compare2 int
		Expect   bool
	}{
		{"Simple equal", slice, 0, 1, false},
		{"Simple equal", slice, 1, 2, true},
		{"Simple equal", slice, 2, 3, true},
		{"Simple equal", slice, 5, 6, false},
		{"Simple equal", slice, 4, 6, false},
	}

	for _, tc := range tables {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			res := tc.Slice.Less(tc.Compare1, tc.Compare2)
			if res != tc.Expect {
				t.Errorf("Less(%v, %v) == %v. Want %v", tc.Compare1, tc.Compare2, res, tc.Expect)
			}
		})
	}
}

func TestSliceSwap(t *testing.T) {
	slice := VersionSlice{
		&Version{"0", []int{0}},
		&Version{"1", []int{1}},
		&Version{"2", []int{2}},
		&Version{"3", []int{3}},
		&Version{"4", []int{4}},
	}

	tables := []struct {
		Name  string
		Slice VersionSlice
		Swap1 int
		Swap2 int
	}{
		{"1 & 2", slice, 1, 2},
		{"3 & 4", slice, 3, 4},
	}

	for _, tc := range tables {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Slice.Swap(tc.Swap1, tc.Swap2)
			if tc.Slice[tc.Swap1].Segments[0] != tc.Swap2 {
				t.Errorf("Swap failed. %v/%v", tc.Slice[tc.Swap1].Original, tc.Slice[tc.Swap2].Original)
			}
		})
	}
}

func TestSliceSort(t *testing.T) {
	slice := VersionSlice{
		&Version{"4", []int{4}},
		&Version{"2", []int{2}},
		&Version{"0", []int{0}},
		&Version{"1", []int{1}},
		&Version{"3", []int{3}},
	}

	slice.Sort()

	for idx := 0; idx < slice.Len(); idx++ {
		if slice[idx].Segments[0] != idx {
			t.Errorf("Slice[%v] = %v. Wanted %v", idx, slice[idx].Segments[0], idx)
		}
	}
}
