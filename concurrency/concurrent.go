package main

import (
	"fmt"
)

func first(a chan struct{}){
	for{
		<-a
		fmt.Println("first")
	}
}

func second(a chan struct{}){
	for{
		<-a
		fmt.Println("second")
	}
}

func third(a chan struct{}){
	for{
		<-a
		fmt.Println("third")
	}
}

type empty struct{}

func main(){
	var input int
	e := empty{}
	f := make(chan struct{})
	s := make(chan struct{})
	t := make(chan struct{})
	go first(f)
	go second(s)
	go third(t)
	fmt.Println("All go routines started")
	fmt.Println("valid input 1,2,3")
	for{
		fmt.Scanf("%d",&input)
		switch(input){
		case 1:
			f <- e
		case 2:
			s <- e
		case 3:
			t <- e
		case 4:
			fmt.Println("invalid input")
			return
		}
	}
}
