package main

import (
	"net/http"

	"sha/countrystate"
	"sha/forgotpassword"
	"sha/langgenres"
	"sha/recentlywatched"
	"sha/signinout"
	"sha/signup"
	"sha/userprofiledetails"
	"sha/userupload"
	"sha/videolikes"
	"sha/videotofront"
	"sha/viewspervideo"
	"sha/yourvideos"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// cassession.ConnectDB()
	r := mux.NewRouter()
	//SIGNUP
	r.HandleFunc("/check", signinout.CheckAPI).Methods("GET")
	r.HandleFunc("/sendotp", signup.OtpSenderInputverifier).Methods("POST")
	r.HandleFunc("/verifyandtopass", signup.CreateAccountOtpVerify).Methods("POST")
	r.HandleFunc("/signup", signup.PasswordEnterSignup).Methods("POST")
	//FORGOT PASSWORD
	r.HandleFunc("/sendotpforgotpassword", forgotpassword.PasswordResetOtpSender).Methods("POST")
	r.HandleFunc("/otpverifyforgotpassword", forgotpassword.ResetPasswordOtpVerify).Methods("POST")
	r.HandleFunc("/newpassword", forgotpassword.EnterNewPassword).Methods("POST")
	//SIGNIN
	r.HandleFunc("/signin", signinout.Signin).Methods("POST")
	//SIGNOUT
	r.HandleFunc("/signout", signinout.Signout).Methods("POST")
	//GET LANGUAGE AND GENRES
	r.HandleFunc("/langen", langgenres.LanguagesGenresForDropDown).Methods("GET")
	//AUTHORIZED
	//HOME
	r.HandleFunc("/home", (signinout.Home)).Methods("GET")
	//UPLOAD
	r.HandleFunc("/upload", (userupload.UploadVideos)).Methods("POST")
	r.HandleFunc("/refresh", (signinout.Refresh)).Methods("GET")
	r.HandleFunc("/videotohomepage", (videotofront.VideoToHomePage)).Methods("GET")
	//UserProfile Details
	r.HandleFunc("/profile", userprofiledetails.UserProfileDetails)
	//all liked videos
	r.HandleFunc("/likes", videolikes.VideoLikesEndPoint).Methods("POST")
	//recently watched by user
	r.HandleFunc("/recentlywatchedbyuser", recentlywatched.RecentlyWatchedVideos).Methods("POST")
	//get the recently watched videos
	r.HandleFunc("/getrecent", recentlywatched.GiveRecentlyWatchedVideosToFront).Methods("GET")
	//get the user profile details
	r.HandleFunc("/getprofile/{id}", userprofiledetails.GiveprofileDetailsToFrontEnd).Methods("GET")
	//get the particular your uploaded videos
	r.HandleFunc("/yourvideos/{id}", yourvideos.GiveYourVideosToFront).Methods("GET")
	//add views to videos
	r.HandleFunc("/views/{vid}", viewspervideo.ViewsPerVideo)
	//countrystate to front end
	r.HandleFunc("/countrystate", countrystate.CountryStateToFront)
	//CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	//listen and serve
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r))
	// http.ListenAndServe(":8080", r)
}

//signinout.IsAuthorized
