package type_switch

func test() {
	var x interface{}
	switch y := x.(type) {
	case int:
		println(y)
	}
}
