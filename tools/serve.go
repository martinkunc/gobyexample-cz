package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	port := "8000"
	publicDir := "docs"
	fs := http.FileServer(http.Dir(publicDir))
	// Like GitHub Pages, serve /arrays from arrays.html.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if path.Ext(r.URL.Path) == "" && r.URL.Path != "/" {
			if _, err := os.Stat(filepath.Join(publicDir, r.URL.Path+".html")); err == nil {
				r.URL.Path += ".html"
			}
		}
		fs.ServeHTTP(w, r)
	})
	fmt.Printf("Serving Go by Example at http://127.0.0.1:%s\n", port)
	http.ListenAndServe(":"+port, handler)
}
