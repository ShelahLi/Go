package main

import "testing"
/**
	testing 提供对 Go 包的自动化测试的支持。通过 `go test` 命令，能够自动执行如下形式的任何函数：
	func TestXxx(*testing.T)
	其中 Xxx 可以是任何字母数字字符串（但第一个字母不能是 [a-z]），用于识别测试例程。
	在这些函数中，使用 Error, Fail 或相关方法来发出失败信号。
*/
func TestSubstr(t *testing.T) {
	/**
		表格驱动测试：分离了测试数据与测试逻辑、明确了出错信息
	 */
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是北京邮电大学", 9},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}

}

/**
	Benchmark:基准测试；(性能测试)
 */
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	// 打印日志
	b.Logf("len(s) = %d", len(s))
	ans := 8
	/**
		重置时间；剔除上方准备输入数据的时间；
	 */
	b.ResetTimer()
	/**
		具体做多少遍测试不用我们关心，次数为：b.N
		输出：运行1000000次，每次执行多长时间:BenchmarkSubstr-4   	 1000000	      1291 ns/op
	*/
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d",
				actual, s, ans)
		}
	}
}
