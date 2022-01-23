package main

import "testing"

const iLikeGo = "I like Go!"
const iLikeGO3 = `I like Go!
I like Go!
I like Go!`

func TestILikeGo(t *testing.T) {
	got := ILikeGo()
	want := iLikeGo
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
func TestRepeatString(t *testing.T) {
	testCases := []struct {
		name   string
		number int
		want   string
	}{
		{"number less than zero", -1, ""},
		{"number equal zero", 0, iLikeGo},
		{"number equal 1", 1, iLikeGo},
		{"number more than zero", 3, iLikeGO3},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := RepeatString(iLikeGo, testCase.number)
			if got != testCase.want {
				t.Errorf("test case |%s| was failed, got=%s, want=%s, number=%d",
					testCase.name, got, testCase.want, testCase.number)
			}
		})
	}
}
