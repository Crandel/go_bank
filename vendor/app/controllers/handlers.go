package controllers

import (
	"net/http"
)

type mainData struct {
	AuthUser uint
}

// MakeHandler - handler wrapper
func MakeHandler(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(h)
}

// mainHandleFunc handler for main page
func mainHandleFunc(w http.ResponseWriter, r *http.Request) {
	// safe type conversion
	// value := r.Context().Value("user")
	// user, check := value.(*m.User)
	// if !check {
	// 	log.Println("User not Found in context")
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// 	return
	// }
	return
}

// MainHandler - for main page
var MainHandler = MakeHandler(mainHandleFunc)

// loginHandleFunc - render template for login page
// func loginHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		login := template.Must(getTemlates("login"))
// 		login.Execute(w, nil)
// 	} else {
// 		r.ParseForm()
// 		// logic part of log in
// 		user := &m.User{
// 			Login:    r.FormValue("login"),
// 			Password: r.FormValue("password"),
// 		}
// 		result, err := valid.ValidateStruct(user)
// 		if err == nil || !result {
// 			err := user.GetUserByLoginPass(user.Login, user.Password)
// 		}
// 	}
// }

// // LoginHandler ...
// var LoginHandler = MakeHandler(loginHandleFunc)

// // signHandleFunc - render template for sign page
// func signHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		templates, err := getTemlates("sign")
// 		sign := template.Must(templates, err)
// 		sign.Execute(w, nil)
// 	} else {
// 		r.ParseForm()
// 		// logic part of sign in
// 		user := &m.User{
// 			Login:    r.FormValue("login"),
// 			Email:    r.FormValue("email"),
// 			Password: r.FormValue("password"),
// 		}
// 		result, err := valid.ValidateStruct(user)
// 		if err == nil || !result {
// 			sess := s.Instance(r)
// 			s.Clear(sess)
// 			err := user.CreateUser()
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusNotAcceptable)
// 				return
// 			}
// 			sess.Values["id"] = user.ID
// 			err = sess.Save(r, w)
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 				return
// 			}
// 			url, err := RedirectFunc("home")
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 				return
// 			}
// 			http.Redirect(w, r, url, http.StatusMovedPermanently)
// 		}
// 	}
// }

// // SignHandler ...
// var SignHandler = MakeHandler(signHandleFunc)

// // signOutHandleFunc - handle func for signout page
// func signOutHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	sess := s.Instance(r)
// 	s.Expire(sess)
// 	sess.Save(r, w)
// 	url, err := RedirectFunc("login")
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}
// 	http.Redirect(w, r, url, http.StatusMovedPermanently)
// }

// // SignOutHandler ...
// var SignOutHandler = MakeHandler(signOutHandleFunc)

// NotFoundHandleFunc ...
// func NotFoundHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Not found handler")
// 	templates, err := getTemlates("404")
// 	notFound := template.Must(templates, err)
// 	w.WriteHeader(404)
// 	notFound.Execute(w, nil)
// }

// // wsHandleFunc ...
// func wsHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	// safe type conversion
// 	user, check := r.Context().Value("user").(*m.User)
// 	if !check {
// 		log.Println("User not Found in context")
// 		return
// 	}
// 	hub := GetHub()
// 	client := &Client{*user, conn, make(chan *SendMessage, 256)}
// 	hub.register <- client
// 	go client.write()
// 	client.read()
// }

// // WsHandler ...
// var WsHandler = MakeHandler(wsHandleFunc)
