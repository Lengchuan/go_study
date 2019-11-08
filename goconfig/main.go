package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
)

func main() {

	c, err := goconfig.LoadConfigFile("../go_study/goconfig/config.ini")
	if err != nil {
		log.Fatal(err)
	}

	n1 := c.MustValue("test", "name", "myname")
	n2 := c.MustValue("test1", "name", "myname")
	fmt.Println("n1 : ", n1)
	fmt.Println("n2 : ", n2)

	//
	tint := c.MustInt("testValues", "tInt", -1)
	tBool := c.MustBool("testValues", "tBool", true)
	tArray := c.MustValueArray("testValues", "tArray", ",")
	fmt.Println("tint: ", tint)
	fmt.Println("tBool: ", tBool)
	fmt.Println("tArray: ", tArray)

	_, err = c.Int("testValues", "123")
	fmt.Println(err)
}
