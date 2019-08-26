package main
import(
"assign3/router"
"log"
"net/http"
)

 func main() {

	log.Println("Listening...")
	router.Routers()
	//It will listen on local host at port number 8080
	http.ListenAndServe(":8080",router.Mux)

	}