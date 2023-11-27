package arr

import "testing"

func Test_Array_Assignment(t *testing.T) {
	var a = [...]int{1, 2, 3}
	var b = a
	b[1] = 5

	t.Log(a) // [1 2 3]   // 数组是值类型，赋值的时候会拷贝一份。所以，可以看到 a 并没有被修改
	t.Log(b) // [1 5 3]   // 可以看得到只有 b[1] 被修改了
}
