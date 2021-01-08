package main
import (
	"fmt"
	"os"
	"bufio"
	"math/rand"
)

func main(){
	fmt.Println("running")
	fmt.Println(rand.Float64() * 3.0)
	exercise1()
}

func exercise1(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0{
		countLines(os.Stdin, counts, "")	
	}else{
		for _, arg := range files{
			f, err := os.Open(arg)
			if err != nil{
				fmt.Println("error of some sort", err)
				continue
			}
			countLines(f, counts, arg)

			f.Close()
		}
	}
	for _, n := range counts{
		if n > 1{
			fmt.Println("its over 1", n)
		}
	}
}

func countLines(f *os.File, counts map[string]int, arg string){
	localCounts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
		localCounts[input.Text()]++
	}
	for line, in := range localCounts {
		if in > 1{
			fmt.Printf("dupe in file %s\n line is %s \n count is %d \n", arg, line, in)
			fmt.Println("-----")
			
		}

	}
}