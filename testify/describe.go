package testify

var describe = DescribeMap{}

func Describe(name string, fnc Callback) int {
	describe[name] = fnc

	return len(describe)
}
