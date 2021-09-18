package main

import (
	"fmt"
	"math/rand"
	"time"
)

//소수 판정 프로그램 v 0.3
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			// count = count + 1
		}
	}

	if isPrime == true {
		fmt.Println(number, "는 소수입니다.")
	} else {
		fmt.Println(number, "는 소수가 아닙니다.")
	}
}
