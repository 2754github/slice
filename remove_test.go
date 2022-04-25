package slice

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	type args[T comparable] struct {
		slice   []T
		element T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}

	stringTests := []testCase[string]{
		{
			name: "unremoved/nil",
			args: args[string]{
				slice:   nil,
				element: "c",
			},
			want: []string{},
		},
		{
			name: "unremoved/empty slice",
			args: args[string]{
				slice:   []string{},
				element: "c",
			},
			want: []string{},
		},
		{
			name: "unremoved",
			args: args[string]{
				slice:   []string{"a", "b"},
				element: "c",
			},
			want: []string{"a", "b"},
		},
		{
			name: "removed",
			args: args[string]{
				slice:   []string{"a", "b"},
				element: "a",
			},
			want: []string{"b"},
		},
	}
	for _, tt := range stringTests {
		t.Run("string/"+tt.name, func(t *testing.T) {
			if got := Remove(tt.args.slice, tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}

	intTests := []testCase[int]{
		{
			name: "unremoved/nil",
			args: args[int]{
				slice:   nil,
				element: 3,
			},
			want: []int{},
		},
		{
			name: "unremoved/empty slice",
			args: args[int]{
				slice:   []int{},
				element: 3,
			},
			want: []int{},
		},
		{
			name: "unremoved",
			args: args[int]{
				slice:   []int{1, 2},
				element: 3,
			},
			want: []int{1, 2},
		},
		{
			name: "removed",
			args: args[int]{
				slice:   []int{1, 2},
				element: 1,
			},
			want: []int{2},
		},
	}
	for _, tt := range intTests {
		t.Run("int/"+tt.name, func(t *testing.T) {
			if got := Remove(tt.args.slice, tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
