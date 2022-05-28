package main

import (
	"fmt"
	"time"
)

func main() {

	Set("achal", "22")

	Expire("achal", 6)

	time.Sleep(2 * time.Second)

	fmt.Println(Ttl("achal"))

	time.Sleep(5 * time.Second)

	fmt.Println(Get("achal"))

}
