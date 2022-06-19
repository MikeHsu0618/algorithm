package example1

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}

	// slice 不能直接被比較
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

// 執行 go test -short 將會被 testing.Short 接收，而 t.Skip 可以略過測試
func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}
}

func TestSplitAll(t *testing.T) {
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"test 1", "abc", "b", []string{"a", "c"}},
		{"test 2", "a:b:c", ":", []string{"a", "b", "c"}},
		{"test 3", "abc", "ㄎ", []string{"abc"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Split(tt.input, tt.sep)
			//if !reflect.DeepEqual(tt.want, got) {
			//	t.Errorf("expected: %v, got: %v", tt.want, got)
			//}
			assert.Equal(t, got, tt.want)
		})
	}
}
