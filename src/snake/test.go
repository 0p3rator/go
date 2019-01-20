package snake

// import "fmt"

type options struct {
	x, y, z int
}

func newOption(opt ...serverOptions) *options {
	r := new(options)
	for _, option := range opt {
		option(r)
	}

	return r
}

type serverOptions func(*options)

func writeX(x int) serverOptions {
	return func(o *options) {
		o.x = x
	}
}
func writeY(y int) serverOptions {
	return func(o *options) {
		o.y = y
	}
}
func writeZ(z int)    serverOptions {
	return func(o *options) {
		o.z = z
	}
}
// func main() {
// 	opt1 := writeX(1)
// 	opt2 := writeY(2)
// 	opt3 := writeZ(3)
// 	opt := newOption(opt1, opt2, opt3)
// 	fmt.Printf("%d; %d; %d", opt.x, opt.y, opt.z)
// }