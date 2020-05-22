package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	uuid "github.com/satori/go.uuid"
)

// Server struct server
type Server struct {
	svc *Service
	app *fiber.App
}

// NewServer initialize server
func NewServer() *Server {
	s := &Server{}

	// initiate service
	s.svc = NewService()

	// initiate fiber
	s.app = fiber.New()
	return s
}

// Serve server
func (s *Server) Serve() {
	s.routes()
	// listen server to port: 8080
	log.Fatal(s.app.Listen(8080))
}

// list of routes for the api
func (s *Server) routes() {
	// serve static file in directory ./public
	s.app.Static("/", "./public")

	s.app.Get("/api/todo", s.GetHandler)
	s.app.Get("/api/todo/:id", s.GetByIDHandler)
	s.app.Post("/api/todo", s.PostHandler)
	s.app.Put("/api/todo/:id", s.PutHandler)
	s.app.Delete("/api/todo/:id", s.DeleteHandler)

}

// GetHandler get list of todo
// url /api/todo
// response
// [
//     {
//         "id": 7,
//         "message": "Testing",
//         "complete": true
//     },
//     {
//         "id": 8,
//         "message": "Testing 2",
//         "complete": true
//     }
// ]
func (s *Server) GetHandler(c *fiber.Ctx) {
	result, err := s.svc.Resolve()
	if err != nil {
		Failed(c, http.StatusInternalServerError, err)
		return
	}
	Success(c, http.StatusOK, result)
}

// GetByIDHandler get todo by id
// url /api/todo/7
// response
// {
//         "id": 7,
//         "message": "Testing",
//         "complete": true
// }
func (s *Server) GetByIDHandler(c *fiber.Ctx) {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		Failed(c, http.StatusInternalServerError, fmt.Errorf("id is not valid %s: %v", c.Params("id"), err.Error()))
		return
	}

	result, err := s.svc.ResolveByID(id)
	if err != nil {
		Failed(c, http.StatusInternalServerError, err)
		return
	}
	Success(c, http.StatusOK, result)
}

// PostHandler create todo
// url /todo
// payload
// {
//         "message": "test data",
//         "complete": true
// }
// repsonse
// {
//			"id": 1
//         "message": "test data",
//         "complete": true
// }
func (s *Server) PostHandler(c *fiber.Ctx) {
	var i Input
	if err := c.BodyParser(&i); err != nil {
		Failed(c, http.StatusBadRequest, err)
		return
	}

	result, err := s.svc.Create(i)
	if err != nil {
		Failed(c, http.StatusInternalServerError, err)
		return
	}
	Success(c, http.StatusCreated, result)
}

// PutHandler update todo
// url /todo/7
// payload
// {
//         "message": "test update data",
//         "complete": true
// }
// repsonse
// {
//			"id": 7
//         "message": "test update data",
//         "complete": true
// }
func (s *Server) PutHandler(c *fiber.Ctx) {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		Failed(c, http.StatusInternalServerError, fmt.Errorf("id is not valid %s: %v", c.Params("id"), err.Error()))
		return
	}

	var update Update
	if err := c.BodyParser(&update); err != nil {
		Failed(c, http.StatusBadRequest, err)
		return
	}

	result, err := s.svc.Update(id, update)
	if err != nil {
		Failed(c, http.StatusInternalServerError, err)
		return
	}
	Success(c, http.StatusOK, result)
}

// DeleteHandler delete todo
// url /todo/7
// repsonse
// OK
func (s *Server) DeleteHandler(c *fiber.Ctx) {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		Failed(c, http.StatusInternalServerError, fmt.Errorf("id is not valid %s: %v", c.Params("id"), err.Error()))
		return
	}

	err = s.svc.Remove(id)
	if err != nil {
		Failed(c, http.StatusInternalServerError, err)
		return
	}
	Success(c, http.StatusOK, "OK")
}

// Success wrap success response
func Success(c *fiber.Ctx, status int, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		c.Status(http.StatusInternalServerError).SendString("failed")
	}
	c.Status(status).SendBytes(js)
}

// Failed wrap failed response
func Failed(c *fiber.Ctx, status int, err error) {
	resp := map[string]interface{}{
		"error": fmt.Sprintf("%s", err),
	}
	js, _ := json.Marshal(resp)
	c.Status(status).SendBytes(js)
}
