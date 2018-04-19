package main

func main() {
	//LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				println(i)
				//break //跳出内层循环
				//break LABEL1 //break到LABEL1
				goto LABEL1 //LABEL1在前导致死循环
			}
		}
	}
LABEL1:
	println("LABEL1")

LABEL2:

	for i := 0; i < 10; i++ {
		for {
			println("aaaaaaaaaa")
			continue LABEL2 //跳出内层无线循环
		}
	}
}
