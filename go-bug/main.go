package main


func Func1() int{
	a:=0
	defer func(){
		a++
	}
	return a
}



func main() {
	manyOptionFunc(1, 2, 3, 3, 4)
}
