package main

import "testing"

func TestToMinut(t *testing.T) {
	var testCases = []struct {
		timee string
		want  string
	}{
		{"0", "0 min"},
		{"1h30m", "90 min"},
		{"", ""},
	}

	for i, tc := range testCases {
		value, err := ToMinut(tc.timee)
		if err != nil {
			t.Errorf("Test case %d was failed, err = %v", i, err)
		}
		if value != tc.want {
			t.Errorf("want = %s not equal returned value = %s", tc.want, value)
		}
	}
}
