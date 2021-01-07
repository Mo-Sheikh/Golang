package main
import ("fmt"
		 "os"
		 "strings"
		)

func main(){
	var s int;
	s++
	fmt.Println("s is",s)
	fmt.Println("exercise 1")
	
	/*Modify the echo program to also print os.Args[0], the name of the command that invoked it.*/
	
	exercise1()
	fmt.Println("exercise 2")

	/* Modify the echo program to print the index and value of each of its arguments, one per line */

	exercise2()
	fmt.Println("exercise 3")
	
	/*Experiment to measure the difference in running time bet ween our potentially
	inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the
	time package, and Section 11.4 shows how to write benchmark tests for systematic performance
	evaluation.*/
	
	exercise3()

}
func exercise1(){
	s, sep := "", ""
	for _, value := range os.Args[1:]{
		s += sep + value
		sep = " "
	}
	print(s)
}

func exercise2(){
	s, sep := "", ""
	for index, value := range os.Args[0:]{
		s += sep + value
		sep = " "
		fmt.Println(index)
		fmt.Println(value)
	}
	fmt.Println(s)
}

func exercise3(){

	fmt.Println(strings.Join(os.Args[1:], " "))
}


func test(){
	var s, sep string = "hello", "mate"
	const firstName string = "mohamed"
	LastName := "abdi"
	counter :=1
	var status bool = true
	fmt.Println("args is ",os.Args)
	for i := 1; i < 10; i++ {
		fmt.Println("args is ",os.Args[2:])
		fmt.Println("hello mate")
		fmt.Println(LastName)
		fmt.Println(s)
		fmt.Println(s)
		fmt.Println(sep)
		fmt.Println(status)
		}
	for {

		counter ++
		fmt.Println("counter is ", counter)
		if counter == 10{
			return
		}
	}
	
}

func test1() {

	var intArray = [5]string{"Aubameyang","Lacazette","Ozil","Chambers","Saka"}
	var s = "hello";
	fmt.Println("args is ",s)
	for _, value := range intArray {
	fmt.Println("trinkster is ")
	fmt.Println("args is ",value)

	}
	fmt.Println("DONE")
	}



