package grammarPractice3

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func handle(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

func IoWriterInterfaceTest() {
	msg := []string{"This", "is", "an", "example", "of", "io.Writer"}

	for _, s := range msg {
		time.Sleep(100 * time.Millisecond)
		handle(os.Stdout, s)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r.URL.Path[1:])
	})

	fmt.Println("start listening on port 4000")
	http.ListenAndServe(":4000", nil)
}
