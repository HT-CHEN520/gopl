package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *LinkNode
	}
	tests := []struct {
		name string
		args args
		want *LinkNode
	}{
		{
			name: "test",
			args: args{BuildListWithNoHead([]int{1, 2, 3, 4, 5})},
			want: BuildListWithNoHead([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

/* func Test_reverseList(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{5, 4, 3, 2, 1}), reverseList(BuildListWithNoHead([]int{1, 2, 3, 4, 5})))
}
*/
