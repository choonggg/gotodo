package todo

import (
  "net/http"

  "github.com/go-chi/chi"
  "github.com/go-chi/render"
  c "github.com/choonggg/gotodo/config"
  m "github.com/choonggg/gotodo/models"
)

func Routes(env *c.Env) *chi.Mux{
  router := chi.NewRouter()

  router.Get("/", GetAllTodos(env))
  router.Get("/{todoID}", GetTodo(env))
  return router
}

func GetTodo(env *c.Env) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    todoIDStr := chi.URLParam(r, "todoID")

    todo := env.DB.Find(&m.Todo{}, todoIDStr)


    render.JSON(w, r, todo)
  }
}

func GetAllTodos(env *c.Env) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    todos  := []m.Todo{}
    env.DB.Find(&todos)

    render.JSON(w, r, todos)
  }
}
