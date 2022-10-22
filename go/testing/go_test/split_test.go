package main

import (
	"reflect"
	"testing"
)

func TestSplitByGoTests(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	var tests []struct {
		name       string
		args       args
		wantResult []string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Split() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// function prefix must be "Test*" and using testing.T as an arg
func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) { // slice 不能直接比較
		t.Errorf("expected:%v, got:%v", want, got) //
	}
}

func TestSplitWithComplexSep(t *testing.T) {
	got := Split("asdawafawda", "awa")
	want := []string{"asd", "fawda"}
	if !reflect.DeepEqual(want, got) { // slice 不能直接比較
		t.Errorf("expected:%v, got:%v", want, got) //
	}
}

// 在一個測試中運行多種子測試
func TestSplitAll(t *testing.T) {
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			t.Parallel() // 將每個測試都表計程能夠彼此並行
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}
