package slice

import (
	"reflect"
	"testing"
)

func TestPositions(t *testing.T) {
	type args[T comparable] struct {
		slice   []T
		element T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []int
	}

	stringTests := []testCase[string]{
		{
			name: "nil",
			args: args[string]{
				slice:   nil,
				element: "c",
			},
			want: []int{},
		},
		{
			name: "empty slice",
			args: args[string]{
				slice:   []string{},
				element: "c",
			},
			want: []int{},
		},
		{
			name: "uncontained",
			args: args[string]{
				slice:   []string{"a", "b"},
				element: "c",
			},
			want: []int{},
		},
		{
			name: "contained",
			args: args[string]{
				slice:   []string{"a", "b"},
				element: "a",
			},
			want: []int{0},
		},
		{
			name: "multiple contained",
			args: args[string]{
				slice:   []string{"a", "b", "a"},
				element: "a",
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range stringTests {
		t.Run("string/"+tt.name, func(t *testing.T) {
			if got := Positions(tt.args.slice, tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Positions() = %v, want %v", got, tt.want)
			}
		})
	}

	intTests := []testCase[int]{
		{
			name: "nil",
			args: args[int]{
				slice:   nil,
				element: 3,
			},
			want: []int{},
		},
		{
			name: "empty slice",
			args: args[int]{
				slice:   []int{},
				element: 3,
			},
			want: []int{},
		},
		{
			name: "uncontained",
			args: args[int]{
				slice:   []int{1, 2},
				element: 3,
			},
			want: []int{},
		},
		{
			name: "contained",
			args: args[int]{
				slice:   []int{1, 2},
				element: 1,
			},
			want: []int{0},
		},
		{
			name: "multiple contained",
			args: args[int]{
				slice:   []int{1, 2, 1},
				element: 1,
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range intTests {
		t.Run("int/"+tt.name, func(t *testing.T) {
			if got := Positions(tt.args.slice, tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Positions() = %v, want %v", got, tt.want)
			}
		})
	}
}
