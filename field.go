package hl7

import (
	// "errors"
	"strings"
)

func (f Field) query(q *Query) string {
	if len(f) <= q.RepeatedField {
		return f.String()
	}

	if !q.HasComponent {
		if q.HasRepeatedField {
			return f[q.RepeatedField].String()
		}

		return f.String()
	}

	return f[q.RepeatedField].query(q)
}

func (f Field) querySlice(q *Query) []string {
	if !q.HasComponent {
		if q.HasRepeatedField {
			return f[q.RepeatedField].SliceOfStrings()
		}

		return f[0].SliceOfStrings()
	}

	return f[q.RepeatedField].querySlice(q)
}

func (f Field) String() string {
	return strings.Join(f.SliceOfStrings(), repeatingFieldSeperator)
}

func (f Field) SliceOfStrings() []string {
	strs := []string{}

	for _, fi := range f {
		strs = append(strs, fi.String())
	}

	return strs
}

func (f Field) setString(q *Query, value string) (Field, error) {
	var err error

	for len(f) < q.RepeatedField+1 {
		f = append(f, RepeatedField{})
	}

	f[q.RepeatedField], err = f[q.RepeatedField].setString(q, value)

	return f, err
}
