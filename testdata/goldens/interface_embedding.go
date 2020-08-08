package undefinedtypes

import "testing"

func TestSomeStructDo(t *testing.T) {
	type fields struct {
		Doer some.Doer
	}
	testCases := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range testCases {
		c := &SomeStruct{
			Doer: tt.fields.Doer,
		}
		c.Do()
	}
}
