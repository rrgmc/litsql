package structarg

import (
	"testing"

	"github.com/rrgmc/litsql"
	"gotest.tools/v3/assert"
)

func TestReflect(t *testing.T) {
	type x struct {
		H string
		J int
		L int `r:"LA"`
		M int `r:"MM,omitempty,x=15"`
		O *int
		P *string
	}

	oval := 45

	value := &x{
		H: "99",
		J: 11,
		L: 45,
		M: 91,
		O: &oval,
	}

	a := New(value, WithTagName("r"))

	for _, test := range []struct {
		name             string
		expected         any
		expectedNotFound bool
	}{
		{
			name:     "H",
			expected: "99",
		},
		{
			name:     "J",
			expected: 11,
		},
		{
			name:             "L",
			expectedNotFound: true,
		},
		{
			name:     "LA",
			expected: 45,
		},
		{
			name:             "M",
			expectedNotFound: true,
		},
		{
			name:     "MM",
			expected: 91,
		},
		{
			name:     "O",
			expected: 45,
		},
		{
			name:     "P",
			expected: nil,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			v, ok := a.Get(test.name)
			if test.expectedNotFound {
				assert.Assert(t, !ok)
			} else {
				assert.Assert(t, ok)
				assert.DeepEqual(t, test.expected, v)
			}
		})
	}
}

func TestReflectEmbed(t *testing.T) {
	type Xembed struct {
		A string
		B int
	}

	type x struct {
		Xembed
		H string
		J int
	}

	value := &x{
		Xembed: Xembed{
			A: "77",
			B: 88,
		},
		H: "99",
		J: 11,
	}

	reflectValuesTest(t, New(value))
}

func TestReflectEmbedPtr(t *testing.T) {
	type Xembed struct {
		A string
		B int
	}

	type x struct {
		*Xembed
		H string
		J int
	}

	value := &x{
		Xembed: &Xembed{
			A: "77",
			B: 88,
		},
		H: "99",
		J: 11,
	}

	reflectValuesTest(t, New(value))
}

func TestReflectEmbedPtrNil(t *testing.T) {
	type Xembed struct {
		A string
		B int
	}

	type x struct {
		*Xembed
		H string
		J int
	}

	value := &x{
		Xembed: nil,
		H:      "99",
		J:      11,
	}

	a := New(value)

	for _, test := range []struct {
		name     string
		expected any
	}{
		{
			name:     "H",
			expected: "99",
		},
		{
			name:     "J",
			expected: 11,
		},
		{
			name: "A",
		},
		{
			name: "B",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			v, ok := a.Get(test.name)
			if test.expected == nil {
				assert.Assert(t, !ok)
			} else {
				assert.Assert(t, ok)
				assert.DeepEqual(t, test.expected, v)
			}
		})
	}
}

func reflectValuesTest(t *testing.T, a litsql.ArgValues) {
	for _, test := range []struct {
		name     string
		expected any
	}{
		{
			name:     "H",
			expected: "99",
		},
		{
			name:     "J",
			expected: 11,
		},
		{
			name:     "A",
			expected: "77",
		},
		{
			name:     "B",
			expected: 88,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			v, ok := a.Get(test.name)
			assert.Assert(t, ok)
			assert.DeepEqual(t, test.expected, v)
		})
	}
}
