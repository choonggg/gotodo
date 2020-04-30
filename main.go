package main

import (
  "net/http"
  "log"
  "time"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/choonggg/gotodo/todo"
  c "github.com/choonggg/gotodo/config"
)

func Routes(env *c.Env) *chi.Mux {
  r := chi.NewRouter()
  r.Use(
    middleware.Logger,
    middleware.Recoverer,
    middleware.RealIP,
    middleware.RequestID,
  )

  r.Use(middleware.Timeout(60 * time.Second))

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome"))
  })

  r.Mount("/todos", todo.Routes(env))

  return r
}

func main() {
  env := c.New()
  router := Routes(env)

  defer env.DB.Close()

  walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
    log.Printf("%s %s\n", method, route) // Walk and print all routes
    return nil
  }

  if err := chi.Walk(router, walkFunc); err != nil {
    log.Panicf("Logging err %s\n", err.Error())
  }

  log.Fatal(http.ListenAndServe(":3000", router))
}
