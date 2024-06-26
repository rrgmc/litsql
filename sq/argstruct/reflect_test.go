package argstruct

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

	a := Values(value, WithTagName("r"))

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
			expected: &oval,
		},
		{
			name:     "P",
			expected: nil,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			v, ok, err := a.Get(test.name)
			assert.NilError(t, err)
			if test.expectedNotFound {
				assert.Assert(t, !ok)
			} else {
				assert.Assert(t, ok)
				assert.DeepEqual(t, test.expected, v)
			}
		})
	}
}

func TestReflectDeref(t *testing.T) {
	type x struct {
		X string
		H *string
		T *string
		J **int
		L ***int
	}

	hval := "99"
	jval := 11
	jval1 := &jval
	lval := 45
	lval1 := &lval
	lval2 := &lval1

	value := &x{
		X: "889",
		H: &hval,
		J: &jval1,
		L: &lval2,
	}

	a := Values(value, WithDerefPointer(true))

	for _, test := range []struct {
		name             string
		expected         any
		expectedNotFound bool
	}{
		{
			name:     "X",
			expected: "889",
		},
		{
			name:     "H",
			expected: "99",
		},
		{
			name:     "T",
			expected: nil,
		},
		{
			name:     "J",
			expected: 11,
		},
		{
			name:     "L",
			expected: 45,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			v, ok, err := a.Get(test.name)
			assert.NilError(t, err)
			if test.expectedNotFound {
				assert.Assert(t, !ok)
			} else {
				assert.Assert(t, ok)
				assert.DeepEqual(t, test.expected, v)
			}
		})
	}
}

func TestReflectEmptyTagName(t *testing.T) {
	type x struct {
		H string
		J int
		L int `r:"LA"`
		M int `r:"MM,omitempty,x=15"`
	}

	value := &x{
		H: "99",
		J: 11,
		L: 45,
		M: 91,
	}

	a := Values(value)

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
			name:     "L",
			expected: 45,
		},
		{
			name:             "LA",
			expectedNotFound: true,
		},
		{
			name:     "M",
			expected: 91,
		},
		{
			name:             "MM",
			expectedNotFound: true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			v, ok, err := a.Get(test.name)
			assert.NilError(t, err)
			if test.expectedNotFound {
				assert.Assert(t, !ok)
			} else {
				assert.Assert(t, ok)
				assert.DeepEqual(t, test.expected, v)
			}
		})
	}
}

func TestReflectMapperFunc(t *testing.T) {
	type x struct {
		H string
		J int
		L int `r:"LA"`
		M int `r:"MM,omitempty,x=15"`
		N int `r:"NN"`
	}

	value := &x{
		H: "99",
		J: 11,
		L: 45,
		M: 91,
		N: 40,
	}

	a := Values(value, WithTagName("r"), WithMapperFunc(func(s string) string {
		switch s {
		case "H":
			return "hmapped"
		case "L":
			return "lmapped"
		case "NN":
			return "nmapped"
		default:
			return s
		}
	}))

	for _, test := range []struct {
		name             string
		expected         any
		expectedNotFound bool
	}{
		{
			name:             "H",
			expectedNotFound: true,
		},
		{
			name:     "hmapped",
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
			name:             "lmapped",
			expectedNotFound: true,
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
			name:             "N",
			expectedNotFound: true,
		},
		{
			name:             "NN",
			expectedNotFound: true,
		},
		{
			name:     "nmapped",
			expected: 40,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			v, ok, err := a.Get(test.name)
			assert.NilError(t, err)
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

	reflectValuesTest(t, Values(value))
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

	reflectValuesTest(t, Values(value))
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

	a := Values(value)

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
			v, ok, err := a.Get(test.name)
			assert.NilError(t, err)
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
			v, ok, err := a.Get(test.name)
			assert.NilError(t, err)
			assert.Assert(t, ok)
			assert.DeepEqual(t, test.expected, v)
		})
	}
}
