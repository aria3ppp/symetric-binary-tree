package symtree_test

import (
	symtree "symmetric-binary-tree"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name string
		root *symtree.Node
		want bool
	}{
		{
			name: "nil root",
			root: nil,
			want: true,
		},
		{
			name: "non symmetric with 2 node on the left",
			root: &symtree.Node{
				Left: &symtree.Node{
					Value: 10,
				},
			},
			want: false,
		},
		{
			name: "non symmetric with 2 node on the right",
			root: &symtree.Node{
				Right: &symtree.Node{
					Value: 10,
				},
			},
			want: false,
		},
		{
			name: "symmetric with 2 node",
			root: &symtree.Node{
				Left:  &symtree.Node{Value: 12},
				Right: &symtree.Node{Value: 12},
			},
			want: true,
		},
		{
			name: "non symmetric with 2 node",
			root: &symtree.Node{
				Left:  &symtree.Node{Value: 12},
				Right: &symtree.Node{Value: 14},
			},
			want: false,
		},
		{
			name: "symmetric with 6 nodes",
			root: &symtree.Node{
				Left: &symtree.Node{
					Left:  &symtree.Node{Value: 10},
					Right: &symtree.Node{Value: 99},
				},
				Right: &symtree.Node{
					Left:  &symtree.Node{Value: 99},
					Right: &symtree.Node{Value: 10},
				},
			},
			want: true,
		},
		{
			name: "symmetric with 6 nodes",
			root: &symtree.Node{
				Left: &symtree.Node{
					Left:  &symtree.Node{Value: 10},
					Right: &symtree.Node{Value: 99},
				},
				Right: &symtree.Node{
					Left:  &symtree.Node{Value: 10},
					Right: &symtree.Node{Value: 99},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if !cmp.Equal(tt.want, symtree.IsSymmetric(tt.root)) {
				if tt.root != nil {
					t.Error(cmp.Diff(tt.root.Left, tt.root.Right))
				} else {
					t.Error(cmp.Diff(tt.want, symtree.IsSymmetric(tt.root)))
				}
			}
		})
	}
}
