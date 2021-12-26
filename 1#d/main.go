package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
func main(){
	wg.Add(10)
	go f1()
	go f2()
	go f3()
	go f4()
	go f5()
	go f6()
	go f7()
	go f8()
	go f9()
	go f10()
	wg.Wait()
	time.Sleep(99999999999999)
}
func f1(){
	var i int32
	var j int32

	for i = 2; i < 1000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}
func f2(){
	var i int32
	var j int32

	for i = 1000000; i < 2000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}
func f3(){
	var i int32
	var j int32

	for i = 2000000; i < 3000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}
func f4(){
	var i int32
	var j int32

	for i = 3000000; i < 4000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}
func f5(){
	var i int32
	var j int32

	for i = 4000000; i < 5000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}
func f6(){
	var i int32
	var j int32

	for i = 5000000; i < 6000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}
func f7(){
	var i int32
	var j int32

	for i = 6000000; i < 7000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}
func f8(){
	var i int32
	var j int32

	for i = 7000000; i < 8000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}
func f9(){
	var i int32
	var j int32

	for i = 8000000; i < 9000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}
func f10(){
	var i int32
	var j int32

	for i = 9000000; i < 10000000; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	wg.Done()
}