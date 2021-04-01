package main

import (
	"fmt"
	"net/http"

	"reflect"
)

var srv http.Server


func main(){
	fmt.Println(reflect.TypeOf(1))

}
