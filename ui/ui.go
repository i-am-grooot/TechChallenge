// Copyright © 2018 Thomas Winsnes <tom@vibrato.com.au>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ui

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Vibrato/TechTestApp/db"
	"github.com/gorilla/mux"
)

// Config configuration for ui package
type Config struct {
	Assets http.FileSystem
	DB     db.Config
}

// Start - start web server and handle web requets
func Start(cfg Config, listener net.Listener) {

	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	mainRouter := mux.NewRouter().PathPrefix("/").Subrouter()
	mainRouter.PathPrefix("/js/").Handler(assetHandler(cfg))
	mainRouter.Handle("/api/task/", allTasksHandler(cfg))
	mainRouter.Handle("/", indexHandler())
	http.Handle("/", mainRouter)

	go server.Serve(listener)
}

const (
	cdnReact           = "https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"
	cdnReactDom        = "https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"
	cdnBabelStandalone = "https://cdnjs.cloudflare.com/ajax/libs/babel-standalone/6.24.0/babel.min.js"
	cdnAxios           = "https://cdnjs.cloudflare.com/ajax/libs/axios/0.16.1/axios.min.js"
)

const indexHTML = `
<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>Vibrato Tech Test App</title>
  </head>
  <body>
    <div id='root'>Virato Tech Test App</div>
		<script src="` + cdnReact + `"></script>
    <script src="` + cdnReactDom + `"></script>
    <script src="` + cdnBabelStandalone + `"></script>
    <script src="` + cdnAxios + `"></script>
		<script src="/js/app.jsx" type="text/babel"></script>
  </body>
</html>
`

func assetHandler(cfg Config) http.Handler {
	return http.FileServer(cfg.Assets)
}

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, indexHTML)
	})
}

func allTasksHandler(cfg Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		output, _ := db.GetAllTasks(cfg.DB)
		js, _ := json.Marshal(output)
		fmt.Fprintf(w, string(js))
	})
}
