package main

import (
	"net/http"

	// "sha/browsing"
	"sha/commentmoderation"
	"sha/countrystate"
	"sha/followers"
	"sha/forgotpassword"
	"sha/langgenres"
	"sha/middleware"
	"sha/playlist"
	"sha/ratings"
	"sha/recentlywatched"
	"sha/signinout"
	"sha/signup"
	"sha/trending"
	"sha/userprofiledetails"
	"sha/userupload"
	"sha/videolikes"
	"sha/videotofront"
	"sha/viewspervideo"
	"sha/watchedlater"
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
	r.HandleFunc("/upload/", (userupload.UploadVideos)).Methods("POST")
	r.HandleFunc("/refresh", (middleware.Refresh)).Methods("GET")
	r.HandleFunc("/videotohomepage", (videotofront.VideoToHomePage)).Methods("GET")
	r.HandleFunc("/likedbythatuser", videolikes.LikedByThatUser)
	//UserProfile Details
	r.HandleFunc("/profile", userprofiledetails.UserProfileDetails)
	//all liked videos
	r.HandleFunc("/likes", (videolikes.AddLikes)).Methods("POST")
	// r.HandleFunc("/getlikes", signinout.IsAuthorized(videolikes.GetLikes)).Methods("GET")
	//recently watched by user
	r.HandleFunc("/recentlywatchedbyuser", (recentlywatched.RecentlyWatchedVideos)).Methods("POST")
	//get the recently watched videos
	r.HandleFunc("/getrecent", (recentlywatched.GiveRecentlyWatchedVideosToFront)).Methods("GET")
	//get the user profile details
	r.HandleFunc("/getprofile/{id}", (userprofiledetails.GiveprofileDetailsToFrontEnd)).Methods("GET")
	//get the particular your uploaded videos
	r.HandleFunc("/yourvideos/{id}", (yourvideos.GiveYourVideosToFront)).Methods("GET")
	//commentmoderation{{
	r.HandleFunc("/comment", (commentmoderation.CreateEvent)).Methods("POST")
	//add views to videos
	r.HandleFunc("/views", (viewspervideo.Views)).Methods("POST")
	// r.HandleFunc("/getviews", viewspervideo.GetViews).Methods("GET")
	r.HandleFunc("/watched", (watchedlater.Watchedlater)).Methods("POST")
	//addnewplaylist
	r.HandleFunc("/add", (playlist.Addplaylistid)).Methods("POST")
	//browsing
	// r.HandleFunc("/browse", browsing.Browse).Methods("POST")
	//countrystates
	r.HandleFunc("/country", countrystate.CountryToFront).Methods("GET")
	r.HandleFunc("/state", countrystate.StatesToFront).Methods("GET")
	r.HandleFunc("/constate", countrystate.Selectcountrystate)
	//trending
	r.HandleFunc("/event", (trending.CreateEvent)).Methods("POST")
	r.HandleFunc("/com", (trending.Getevent)).Methods("GET")
	//FOllowers
	r.HandleFunc("/followers", (followers.Followers)).Methods("POST")
	r.HandleFunc("/getfollowers", (followers.GetFollowers)).Methods("GET")
	r.HandleFunc("/gettopfollow", (followers.GetTop)).Methods("GET")
	r.HandleFunc("/ratings", (ratings.Rating)).Methods("POST")
	// r.HandleFunc("/getrate", (ratings.Getrate)).Methods("GET")

	//CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	//listen and serve
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r))
	// http.ListenAndServe(":8080", r)
}

//signinout.IsAuthorized
