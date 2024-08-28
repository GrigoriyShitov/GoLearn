package main

import (
	"RestApi/controllers"
	"RestApi/controllers/handlers"
	"log"
	"net/http"
)

func main() {

	h := handlers.HandlerInit()
	h.Router.HandleFunc("/wallets/users", h.HomeHandler)                                                                                                             // home page
	h.Router.Handle("/wallets/users/{idUser:[0-9]+}", controllers.JWTmiddleware(http.HandlerFunc(h.GetUserInfoHandler))).Methods(http.MethodGet)                     // get user info
	h.Router.Handle("/wallets/users/{idUser:[0-9]+}/{name}", controllers.JWTmiddleware(http.HandlerFunc(h.CreateNewUserHandler))).Methods(http.MethodPost)           // create new user
	h.Router.Handle("/wallets/users/{idUser:[0-9]+}", controllers.JWTmiddleware(http.HandlerFunc(h.CreateNewWalletHandler))).Methods(http.MethodPost)                // create new wallet
	h.Router.Handle("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}", controllers.JWTmiddleware(http.HandlerFunc(h.OperationListHandler))).Methods(http.MethodGet) //get operations list
	h.Router.Handle("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}/{idOperation:[0-9]+}", controllers.JWTmiddleware(http.HandlerFunc(h.OperationDeleteHandler))).Methods(http.MethodDelete)
	h.Router.Handle("/wallets/users/{whoChange:[0-9]+}/{idUser:[0-9]+}/{role:admin|user}", controllers.JWTmiddleware(http.HandlerFunc(h.ADMIN_EditRoleHandler))).Methods(http.MethodPatch)
	h.Router.Handle("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}/{type:deposit|withdraw}/{amount:[0-9]+}/{category}", controllers.JWTmiddleware(http.HandlerFunc(h.MakeOperationHandler))).Methods(http.MethodPost)
	log.Println("API is running!")
	http.ListenAndServe(":8080", h.Router)
}
