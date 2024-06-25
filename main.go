package main
import ( 
	"net/http"
	"fmt"
)


func main () {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write( []byte("<h1>Hello Docker</h1>"))
	})

	fmt.Println("Running!")
	fmt.Println(http.ListenAndServe(":8080", nil))

}