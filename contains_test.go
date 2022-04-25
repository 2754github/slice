package slice

import (
	"testing"
)

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		slice   []T
		element T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}

	stringTests := []testCase[string]{
		{
			name: "uncontained/nil",
			args: args[string]{
				slice:   nil,
				element: "c",
			},
			want: false,
		},
		{
			name: "uncontained/empty slice",
			args: args[string]{
				slice:   []string{},
				element: "c",
			},
			want: false,
		},
		{
			name: "uncontained",
			args: args[string]{
				slice:   []string{"a", "b"},
				element: "c",
			},
			want: false,
		},
		{
			name: "contained",
			args: args[string]{
				slice:   []string{"a", "b"},
				element: "a",
			},
			want: true,
		},
	}
	for _, tt := range stringTests {
		t.Run("string/"+tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}

	intTests := []testCase[int]{
		{
			name: "uncontained/nil",
			args: args[int]{
				slice:   nil,
				element: 3,
			},
			want: false,
		},
		{
			name: "uncontained/empty slice",
			args: args[int]{
				slice:   []int{},
				element: 3,
			},
			want: false,
		},
		{
			name: "uncontained",
			args: args[int]{
				slice:   []int{1, 2},
				element: 3,
			},
			want: false,
		},
		{
			name: "contained",
			args: args[int]{
				slice:   []int{1, 2},
				element: 1,
			},
			want: true,
		},
	}
	for _, tt := range intTests {
		t.Run("int/"+tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
