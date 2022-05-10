package main

import (
	"fmt"
	"time"
)

func main() {

	Set("achal", "22")
	Expire("achal", 3)

	time.Sleep(1 * time.Second)

	fmt.Println(Ttl("achal"))

	time.Sleep(1 * time.Second)

	fmt.Println(Get("achal"))

}
