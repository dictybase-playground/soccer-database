package main

import (

  "net/http"
  "github.com/julienschmidt/httprouter"
  "context"

)

const params = "params"

//httprouter wrapper
type RouterWrapper struct {
  Router *httprouter.Router
}

//initialization
func NewRouter() *RouterWrapper{
  return &RouterWrapper{Router: httprouter.New()}
}
// Get is a shortcut for router.Handle("GET", path, handle)
func (r *RouterWrapper) Get(path string, fn http.HandlerFunc) {
	r.Router.GET(path, HandlerFunc(fn))
}

// Post is a shortcut for router.Handle("POST", path, handle)
func (r *RouterWrapper) Post(path string, fn http.HandlerFunc) {
	r.Router.POST(path, HandlerFunc(fn))
}
func HandlerFunc(fn http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := context.WithValue(r.Context(), params, p)
		fn(w, r.WithContext(ctx))
	}
}

// Context returns the URL parameters
func Params(r *http.Request) httprouter.Params {
	return r.Context().Value(params).(httprouter.Params)
}
//  r := httprouter.New()
  //r.GET("/", HomeHandler)
  //r.GET("/countries", countryHandler)

  // Posts collection
  //r.GET("/posts", PostsIndexHandler)
  //r.POST("/posts", PostsCreateHandler)

  // Posts singular
  //r.GET("/posts/:id", PostShowHandler)
  //r.PUT("/posts/:id", PostUpdateHandler)
  //r.GET("/posts/:id/edit", PostEditHandler)

  //fmt.Println("Starting server on :8080")
  //http.ListenAndServe(":8080", r)
