package hello

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	context := appengine.NewContext(request)
	currentUser := user.Current(context)

	if currentUser == nil {
		url, err := user.LoginURL(context, request.URL.String())
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Location", url)
		writer.WriteHeader(http.StatusFound)
		return
	}

	fmt.Fprintf(writer, "Hello, %v!", currentUser)
}
