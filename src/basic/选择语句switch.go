package main

func main() {
	a := 1
	switch a {
	case 0:
		println("a = 0")
	case 1:
		println("a = 1")
	default:
		println("a != 1 && a != 0")
	}

	switch {
	case a >= 0:
		println("a >= 0")
		fallthrough
	case a >= 1:
		println("aa >= 1")
		//fallthrough
	default:
		println("a != 1 && a != 0")
	}
}
