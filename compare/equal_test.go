package compare

import "testing"

// unit test
func TestEqual(t *testing.T) {
	testCases := []struct {
		a    []byte
		b    []byte
		want bool
	}{
		{[]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'c'}, true},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := Equal(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("failed to compare %+v and %+b", tc.a, tc.b)
			}
		})
	}
}

// fuzz test
func FuzzEqual(f *testing.F) {
	testCases := []struct {
		a []byte
		b []byte
	}{
		{[]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'c'}},
	}

	for _, tc := range testCases {
		f.Add(tc.a, tc.b)
	}

	f.Fuzz(func(t *testing.T, a, b []byte) {
		// We don't have any assertions here. This test will fail if there is a
		// panic in our code. We can add assertions as well if we want to take
		// this test one step further.
		Equal(a, b)
	})
}
