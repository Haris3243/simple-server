package main
/*
*	Purpose: Simple server which get value from Environment Variable and returns to the client in repsonse
*	Developer: Muhammad Haris Shafiq (harisshafiq08@gmail.com)
*	Date: 28-01-2020
*/

import(
	"fmt"
	"os"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SayHello API triggered!")
	
	//write value of environment variable MYNAME to response
	fmt.Fprintf(w,"Hello "+os.Getenv("MYNAME"))
}

func main(){
	
	fmt.Println("Name is:"+os.Getenv("MYNAME"))
	
	http.HandleFunc("/hello", sayHello)
	
	//server is listening on port 8083
	http.ListenAndServe(":8083", nil)

}