package controllers

import "github.com/victorsteven/fullstack/api/middlewares"

func (server *Server) initializeRoutes() {

	// Home Route
	route := server.Router.HandleFunc
	setMiddleJSON := middlewares.SetMiddlewareJSON
	auth := middlewares.SetMiddlewareAuthentication

	route("/", setMiddleJSON(server.Home)).Methods("GET")

	// Login Route
	route("/login", setMiddleJSON(server.Login)).Methods("POST")

	//Users routes
	route("/users", setMiddleJSON(server.CreateUser)).Methods("POST")
	route("/users", setMiddleJSON(server.GetUsers)).Methods("GET")
	route("/users/{id}", setMiddleJSON(server.GetUser)).Methods("GET")
	route("/users/{id}", setMiddleJSON(auth(server.UpdateUser))).Methods("PUT")
	route("/users/{id}", auth(server.DeleteUser)).Methods("DELETE")

	//Posts routes
	route("/posts", setMiddleJSON(server.CreatePost)).Methods("POST")
	route("/posts", setMiddleJSON(server.GetPosts)).Methods("GET")
	route("/posts/{id}", setMiddleJSON(server.GetPost)).Methods("GET")
	route("/posts/{id}", setMiddleJSON(auth(server.UpdatePost))).Methods("PUT")
	route("/posts/{id}", auth(server.DeletePost)).Methods("DELETE")
}
