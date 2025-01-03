package transport

import (
	"encoding/json"
	"github.com/citixenken/go-api.git/internal/todo"
	"log"
	"net/http"
)

type TodoItem struct {
	//ID   int    `json:"id"`
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {
	//var todos = make([]TodoItem, 0)
	//var nextID = 1
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		todoItems, err := todoSvc.GetAll()

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(todoItems)

		if err != nil {
			log.Println(err)
		}

		_, err = w.Write(b)

		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /todo", func(w http.ResponseWriter, r *http.Request) {
		var t TodoItem
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//t.ID = nextID
		//nextID++
		//todos = append(todos, t)
		err = todoSvc.Add(t.Item)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	})

	//mux.HandleFunc("DELETE /todo", func(w http.ResponseWriter, r *http.Request) {
	//	idStr := r.URL.Query().Get("id")
	//	if idStr == "" {
	//		w.WriteHeader(http.StatusBadRequest)
	//		w.Write([]byte("Missing ID"))
	//		return
	//	}
	//
	//	id, err := strconv.Atoi(idStr)
	//	if err != nil {
	//		log.Println(err)
	//		w.WriteHeader(http.StatusBadRequest)
	//		w.Write([]byte("Invalid ID"))
	//		return
	//	}
	//
	//	found := false
	//	for i, todo := range todos {
	//		if todo.ID == id {
	//			todos = append(todos[:i], todos[i+1:]...)
	//			found = true
	//			break
	//		}
	//	}
	//
	//	if !found {
	//		w.WriteHeader(http.StatusNotFound)
	//		w.Write([]byte("Todo not found"))
	//		return
	//	}
	//
	//	w.WriteHeader(http.StatusOK)
	//	w.Write([]byte("Deleted"))
	//})

	mux.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		results, err := todoSvc.Search(query)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(results)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
			return
		}
	})

	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
