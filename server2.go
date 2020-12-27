package main
//
//import (
//	"encoding/json"
//	"fmt"
//	//"github.com/gorilla/mux"
//	"log"
//	"net/http"
//	"os"
//
//	"./application"
//	"./domain"
//	"./repository"
//	router "./router"
//)
//
//func printDebugf(format string, args ...interface{}) {
//	if env := os.Getenv("GO_SERVER_DEBUG"); len(env) != 0 {
//		log.Printf("[DEBUG] "+format+"\n", args...)
//	}
//}
//
//// ErrorResponse is Error response template
//type ErrorResponse struct {
//	Message string `json:"reason"`
//	Error   error  `json:"-"`
//}
//
//func (e *ErrorResponse) String() string {
//	return fmt.Sprintf("reason: %s, error: %s", e.Message, e.Error.Error())
//}
//
//// Respond is response write to ResponseWriter
//func Respond(w http.ResponseWriter, code int, src interface{}) {
//	var body []byte
//	var err error
//
//	switch s := src.(type) {
//	case []byte:
//		if !json.Valid(s) {
//			Error(w, http.StatusInternalServerError, err, "invalid json")
//			return
//		}
//		body = s
//	case string:
//		body = []byte(s)
//	case *ErrorResponse, ErrorResponse:
//		// avoid infinite loop
//		if body, err = json.Marshal(src); err != nil {
//			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
//			return
//		}
//	default:
//		if body, err = json.Marshal(src); err != nil {
//			Error(w, http.StatusInternalServerError, err, "failed to parse json")
//			return
//		}
//	}
//	w.WriteHeader(code)
//	_, _ = w.Write(body)
//}
//
//// Error is wrapped Respond when error response
//func Error(w http.ResponseWriter, code int, err error, msg string) {
//	e := &ErrorResponse{
//		Message: msg,
//		Error:   err,
//	}
//	printDebugf("%s", e.String())
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	Respond(w, code, e)
//}
//
//// JSON is wrapped Respond when success response
//func JSON(w http.ResponseWriter, code int, src interface{}) {
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	Respond(w, code, src)
//}
//
//// Handler user handler
//type Handler struct {
//	Repository repository.UserRepository
//}
//
//// Routes returns the initialized router
//func (handler Handler) Routes() *router.Router {
//	router := router
//	router.Get("/user/:id", handler.getUser)
//	router.Get("/users", handler.getUsers)
//	router.Post("/user", handler.createUser)
//	return router
//}
//
//// Run start server
//func (handler Handler) Run(port int) error {
//	log.Printf("Server running at http://localhost:%d/", port)
//	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler.Routes())
//}
//
//func (handler Handler) getUser(w http.ResponseWriter, r *http.Request, id int) {
//	ctx := r.Context()
//
//	interactor := application.UserInteractor{
//		Repository: handler.Repository,
//	}
//	user, err := interactor.GetUser(ctx, id)
//	if err != nil {
//		Error(w, http.StatusNotFound, err, "failed to get user")
//		return
//	}
//	JSON(w, http.StatusOK, user)
//}
//
//func (handler Handler) getUsers(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//
//	interactor := application.UserInteractor{
//		Repository: handler.Repository,
//	}
//	users, err := interactor.GetUsers(ctx)
//	if err != nil {
//		Error(w, http.StatusNotFound, err, "failed to get user list")
//		return
//	}
//	type payload struct {
//		Users []*domain.User `json:"users"`
//	}
//	JSON(w, http.StatusOK, payload{Users: users})
//}
//
//func (handler Handler) createUser(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//
//	type payload struct {
//		Name string `json:"name"`
//		Pass string `json:"password"`
//		DOB  string `json:"dob"`
//	}
//
//	var p payload
//	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
//		Error(w, http.StatusBadRequest, err, "failed to parse request")
//		return
//	}
//
//	interactor := application.UserInteractor{
//		Repository: handler.Repository,
//	}
//
//	if err := interactor.AddUser(ctx, p.Name, p.Pass, p.DOB); err != nil {
//		Error(w, http.StatusInternalServerError, err, "failed to create user")
//		return
//	}
//	JSON(w, http.StatusCreated, nil)
//}
