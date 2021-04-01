package main 
import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"io"
)

func mainz(){
	for _, site := range os.Args[1:]{
		if(!strings.HasPrefix(site, "http://")){
			fmt.Println("appending https")
			site = "http://" + site
			fmt.Println(site)
			
		}else{
		fmt.Println("good to go")
		}
		fmt.Printf("site is %s\n", site)
		resp, err := http.Get(site)
		fmt.Println("status code", resp.Status)
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
		fmt.Println("content ", b)

	}
}