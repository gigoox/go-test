package main

import "encoding/json"

func main(){

	json.Unmarshal()

	quit := make(chan int)
	go func() {
		quit <- 10
	}()
	select {
	case a := <-quit:
		println(a)
		return
	}
}