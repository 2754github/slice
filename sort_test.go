package slice

import (
	"reflect"
	"testing"
)

// Sort is just a wrapper for `sort.SliceStable` and `sort.Slice`, so it is not rigorously tested.
func TestSort(t *testing.T) {
	type args[T comparable] struct {
		slice  []T
		less   func(i, j int) bool
		stable bool
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}

	stringTests := []testCase[string]{
		{
			name: "stable=false",
			args: args[string]{
				slice:  []string{"one", "two", "three", "four"},
				stable: false,
			},
			want: []string{"three", "four", "one", "two"},
		},
		{
			name: "stable=true",
			args: args[string]{
				slice:  []string{"one", "two", "three", "four"},
				stable: true,
			},
			want: []string{"three", "four", "one", "two"},
		},
	}
	for _, tt := range stringTests {
		t.Run("string/"+tt.name, func(t *testing.T) {
			// DESC of string length
			tt.args.less = func(i, j int) bool {
				return len(tt.args.slice[i]) > len(tt.args.slice[j])
			}

			// Exec
			Sort(tt.args.slice, tt.args.less, tt.args.stable)

			// Compare
			if !reflect.DeepEqual(tt.args.slice, tt.want) {
				t.Errorf("Sort() = %v, want %v", tt.args.slice, tt.want)
			}
		})
	}

	intTests := []testCase[int]{
		{
			name: "stable=false",
			args: args[int]{
				slice:  []int{1997, 1, 25},
				stable: false,
			},
			want: []int{1997, 25, 1},
		},
		{
			name: "stable=true",
			args: args[int]{
				slice:  []int{1997, 1, 25},
				stable: true,
			},
			want: []int{1997, 25, 1},
		},
	}
	for _, tt := range intTests {
		t.Run("int/"+tt.name, func(t *testing.T) {
			// DESC of number
			tt.args.less = func(i, j int) bool {
				return tt.args.slice[i] > tt.args.slice[j]
			}

			// Exec
			Sort(tt.args.slice, tt.args.less, tt.args.stable)

			// Compare
			if !reflect.DeepEqual(tt.args.slice, tt.want) {
				t.Errorf("Sort() = %v, want %v", tt.args.slice, tt.want)
			}
		})
	}
}
