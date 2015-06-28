package modalSoul

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	if r.Method == "GET" {
		url := r.URL.Path
		if url == "/" {
			fmt.Fprint(w, "It works!!")
		} else {
			filename := fmt.Sprintf(".%s", url)
			b, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprint(w, "not found.")
			} else {
				fmt.Fprint(w, string(b))
			}
		}
	}
}