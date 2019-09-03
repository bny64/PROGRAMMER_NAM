package grammarPractice5

import (
	"fmt"
	"net/http"
	"strconv"
)

func handle(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	a, _ := strconv.Atoi(v.Get("dividend"))
	b, _ := strconv.Atoi(v.Get("divisor"))
	fmt.Fprintf(w, "%d / %d = %d", a, b, a/b)
}

func NetHttpErrorTest() {
	http.HandleFunc("/divide", handle)
	http.ListenAndServe(":8080", nil)
}
