package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10000; i++ {
		monkeys.M0throw()
		monkeys.M1throw()
		monkeys.M2throw()
		monkeys.M3throw()
		monkeys.M4throw()
		monkeys.M5throw()
		monkeys.M6throw()
		monkeys.M7throw()
	}
	fmt.Println(monkeys)
	fmt.Println(monkeys.count)

}
