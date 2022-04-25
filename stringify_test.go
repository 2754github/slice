package slice

import "testing"

func TestStringify(t *testing.T) {
	type args[T comparable] struct {
		slice  []T
		sep    string
		option *StringifyOption
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want string
	}

	stringTests := []testCase[string]{
		{
			name: "nil",
			args: args[string]{
				slice:  nil,
				sep:    ", ",
				option: nil,
			},
			want: "",
		},
		{
			name: "empty slice",
			args: args[string]{
				slice:  []string{},
				sep:    ", ",
				option: nil,
			},
			want: "",
		},
		{
			name: "1 element/sep",
			args: args[string]{
				slice:  []string{"a"},
				sep:    ", ",
				option: nil,
			},
			want: "a",
		},
		{
			name: "1 element/sep/quotation",
			args: args[string]{
				slice: []string{"a"},
				sep:   ", ",
				option: &StringifyOption{
					Quotation: "`",
				},
			},
			want: "`a`",
		},
		{
			name: "1 element/sep/quotation/conjunction",
			args: args[string]{
				slice: []string{"a"},
				sep:   ", ",
				option: &StringifyOption{
					Quotation:   "`",
					Conjunction: "or",
				},
			},
			want: "`a`",
		},
		{
			name: "2 elements/sep",
			args: args[string]{
				slice:  []string{"a", "b"},
				sep:    ", ",
				option: nil,
			},
			want: "a, b",
		},
		{
			name: "2 elements/sep/quotation",
			args: args[string]{
				slice: []string{"a", "b"},
				sep:   ", ",
				option: &StringifyOption{
					Quotation: "`",
				},
			},
			want: "`a`, `b`",
		},
		{
			name: "2 elements/sep/quotation/conjunction",
			args: args[string]{
				slice: []string{"a", "b"},
				sep:   ", ",
				option: &StringifyOption{
					Quotation:   "`",
					Conjunction: "or",
				},
			},
			want: "`a` or `b`",
		},
		{
			name: "3 elements/sep",
			args: args[string]{
				slice:  []string{"a", "b", "c"},
				sep:    ", ",
				option: nil,
			},
			want: "a, b, c",
		},
		{
			name: "3 elements/sep/quotation",
			args: args[string]{
				slice: []string{"a", "b", "c"},
				sep:   ", ",
				option: &StringifyOption{
					Quotation: "`",
				},
			},
			want: "`a`, `b`, `c`",
		},
		{
			name: "3 elements/sep/quotation/conjunction",
			args: args[string]{
				slice: []string{"a", "b", "c"},
				sep:   ", ",
				option: &StringifyOption{
					Quotation:   "`",
					Conjunction: "or",
				},
			},
			want: "`a`, `b` or `c`",
		},
	}
	for _, tt := range stringTests {
		t.Run("string/"+tt.name, func(t *testing.T) {
			if got := Stringify(tt.args.slice, tt.args.sep, tt.args.option); got != tt.want {
				t.Errorf("Stringify() = %v, want %v", got, tt.want)
			}
		})
	}
}
