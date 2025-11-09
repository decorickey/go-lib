package stream_test

import (
	"strings"
	"testing"

	"github.com/decorickey/go-lib/internal/stream"
	"github.com/stretchr/testify/assert"
)

func TestFilterInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input <-chan int
		fn    func(int) bool
		want  []int
	}{
		{
			name: "偶数のみ抽出",
			input: stream.FromSlice(
				[]int{1, 2, 3, 4, 5},
			),
			fn: func(v int) bool {
				return v%2 == 0
			},
			want: []int{2, 4},
		},
		{
			name: "全要素一致",
			input: stream.FromSlice(
				[]int{1, 2, 3, 4, 5},
			),
			fn: func(v int) bool {
				return v > 0
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "該当なし",
			input: stream.FromSlice(
				[]int{1, 3, 5},
			),
			fn: func(v int) bool {
				return v%2 == 0
			},
			want: []int{},
		},
		{
			name: "空入力",
			input: stream.FromSlice(
				[]int{},
			),
			fn: func(v int) bool {
				return v%2 == 0
			},
			want: []int{},
		},
		{
			name:  "nil入力",
			input: nil,
			fn: func(v int) bool {
				return v%2 == 0
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stream.Filter(tt.input, tt.fn)
			results := make([]int, 0, len(tt.want))
			for v := range got {
				results = append(results, v)
			}
			assert.Equal(t, tt.want, results)
		})
	}
}

func TestFilterString(t *testing.T) {
	t.Parallel()

	tests := []struct {
	name  string
	input <-chan string
	fn    func(string) bool
	want  []string
	}{
		{
			name: "hoge 含有抽出",
			input: stream.FromSlice(
				[]string{"hoge", "fuga", "piyo", "hogehoge"},
			),
			fn: func(v string) bool {
				return strings.Contains(v, "hoge")
			},
			want: []string{"hoge", "hogehoge"},
		},
		{
			name: "接頭辞一致",
			input: stream.FromSlice(
				[]string{"foo", "foobar", "barfoo"},
			),
			fn: func(v string) bool {
				return strings.HasPrefix(v, "foo")
			},
			want: []string{"foo", "foobar"},
		},
		{
			name: "空入力",
			input: stream.FromSlice(
				[]string{},
			),
			fn: func(v string) bool {
				return strings.Contains(v, "dummy")
			},
			want: []string{},
		},
		{
			name:  "nil入力",
			input: nil,
			fn: func(v string) bool {
				return strings.Contains(v, "dummy")
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stream.Filter(tt.input, tt.fn)
			results := make([]string, 0, len(tt.want))
			for v := range got {
				results = append(results, v)
			}
			assert.Equal(t, tt.want, results)
		})
	}
}
