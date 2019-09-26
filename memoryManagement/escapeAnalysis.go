package main

//go:noinline
func f() *int{
	// this should escape to heap
	i := 0
	return &i
}

// build this using go build -gcflags "-m -m"
func main(){
	f()
}