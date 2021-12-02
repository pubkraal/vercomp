package filter

import "github.com/pubkraal/vercomp/internal/version"

func MakeSlice(versions []string) (version.VersionSlice, error) {
	v := make(version.VersionSlice, len(versions))

	for idx, item := range versions {
		ver, err := version.New(item)
		if err != nil {
			return nil, err
		}

		v[idx] = ver
	}

	return v, nil
}
