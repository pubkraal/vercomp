package version

import (
	"sort"
	"strconv"
	"strings"
)

// VersionSlice is a slice of Versions implementing sort.Interface.
type VersionSlice []*Version

func (v VersionSlice) Len() int           { return len(v) }
func (v VersionSlice) Less(i, j int) bool { return v[i].Less(v[j]) }
func (v VersionSlice) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v VersionSlice) Sort()              { sort.Sort(v) }

func NewSlice(versions []string) (VersionSlice, error) {
	v := make(VersionSlice, len(versions))

	for idx, item := range versions {
		ver, err := New(item)
		if err != nil {
			return nil, err
		}

		v[idx] = ver
	}

	return v, nil
}

type Version struct {
	Original string
	Segments []int
}

func New(version string) (*Version, error) {
	expl, err := explodeVersion(version)
	if err != nil {
		return nil, err
	}
	return &Version{
		Original: version,
		Segments: expl,
	}, nil
}

func (v *Version) Less(c *Version) bool {
	sl := len(v.Segments)
	cl := len(c.Segments)
	if cl < sl {
		sl = cl
	}

	for idx := 0; idx < sl; idx++ {
		val := v.Segments[idx]
		cmp := c.Segments[idx]
		if val < cmp {
			// Less!
			return true
		} else if val > cmp {
			// More!
			return false
		}
	}

	// Likely the equal case
	return false
}

func (v *Version) Repr() string {
	return v.Original
}

func explodeVersion(version string) ([]int, error) {
	segments := strings.Split(version, ".")
	ret := make([]int, len(segments))
	for idx, seg := range segments {
		i, err := strconv.ParseInt(seg, 10, 64)
		if err != nil {
			return nil, err
		}
		ret[idx] = int(i)
	}

	return ret, nil
}
