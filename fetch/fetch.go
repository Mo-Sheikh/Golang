package main 
import (
	"fmt"
	"net/http"
	"os"

	"io"
)

func main(){
	for _, site := range os.Args[1:]{
		resp, err := http.Get(site)
		if err != nil{
			fmt.Println("you have a problem of some sort brother!", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		
		resp.Body.Close()
		if err != nil {
			fmt.Println("error of some sort bro",err)
			os.Exit(1)
		}
		fmt.Printf("content %s", b)

	}
}