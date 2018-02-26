package simple

import (
	"net/http"
	"fmt"
)

func main() {

	fmt.Println("starting simple webserver")

	http.Handle("/health", SimpleHealthHandler{})
	http.ListenAndServe(":8080", nil)

}
