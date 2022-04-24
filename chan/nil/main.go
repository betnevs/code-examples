package main

func main() {
	var c chan int
	<-c
	//c <- 1
	close(c)
}
