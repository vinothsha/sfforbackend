package main

import (
	"log"
	"net/http"

	"sha/cassession"
	"sha/forgotpassword"
	"sha/signin"
	"sha/signup"
	"sha/userupload"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	cassession.ConnectDB()
	r := mux.NewRouter()
	r.HandleFunc("/check", signin.CheckAPI).Methods("GET")
	r.HandleFunc("/sendotp", signup.OtpSenderInputverifier).Methods("POST")
	r.HandleFunc("/verifyandtopass", signup.CreateAccountOtpVerify).Methods("POST")
	r.HandleFunc("/signup", signup.PasswordEnterSignup).Methods("POST")
	r.HandleFunc("/sendotpforgotpassword", forgotpassword.PasswordResetOtpSender).Methods("POST")
	r.HandleFunc("/otpverifyforgotpassword", forgotpassword.ResetPasswordOtpVerify).Methods("POST")
	r.HandleFunc("/newpassword", forgotpassword.EnterNewPassword).Methods("POST")
	r.HandleFunc("/signin", signin.Signin).Methods("POST")
	r.HandleFunc("/home", signin.IsAuthorized(signin.Home)).Methods("GET")
	r.HandleFunc("/upload", userupload.UploadVideos).Methods("POST")
	r.HandleFunc("/refresh", signin.IsAuthorized(signin.Refresh)).Methods("GET")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))
	http.ListenAndServe(":8080", r)
}
