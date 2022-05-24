Backend

when run server it sometime shows port already in use---Use below Command
     sudo kill -9 $(sudo lsof -t -i:9001)
     
First time Before Start the Server Run the following command
      sfbeta1.0--->backend
      $ make run_script
For Run the Server
     $make run_server

SIGNUP---All Methods Are "POST"
     --Sendotp to New User
           -->http://ip:8080/sendotp
           Data Passed in JSON Format
           if Email:
           {
  		 "email":"saddam@gmail.com",
	    }
	    
           if Mobile:
           {
  		 "mobile":"+919655373273",
   		 "countrycode":"+91",
  		 "otp":"219433"
	   }
           
     --Verify the OTP 
           -->http://ip:8080/verifyandtopass
           Data Passed in JSON Format
           if Email
           {
  		 "email":"saddam@gmail.com",
  		 "otp":"219433"
	    }
           if Mobile:
           {
   		"mobile":"+919655373273",
   		"countrycode":"+91",
   		"otp":"219433"
	   }
           
     --Signup the New User
           -->http://ip:8080/signup
           Data Passed in JSON Format
            if Email
           {
  		 "email":"saddam@gmail.com",
  		 "password":"saddamAws"
	    }
	    if Mobile:
           {
   		"mobile":"+919655373273",
   		"countrycode":"+91",
   		"password":"saddamAws"
	   }
           
SIGNIN---All Methods Are "POST"
           -->http://ip:8080/signin
           Data Passed in JSON Format
            if Email
           {
  		 "email":"saddam@gmail.com",
  		 "password":"saddamAws"
	    }
           if Mobile:
           {
  		 "mobile":"+919655373273",
  		 "countrycode":"+91",
  		 "password":"123456789"
	    }
           
FORGOT PASSWORD---All Methods Are "POST"
     --Sendotp to ForgotPassword users
           -->http://ip:8080/sendotpforgotpassword
           Data Passed in JSON Format
           if Email:
           {
  		 "email":"saddam@gmail.com",
	    }
	    
           if Mobile:
           {
  		 "mobile":"+919655373273",
   		 "countrycode":"+91",
  		 "otp":"219433"
	   }
     --Verify the OTP
           -->http://ip:8080/otpverifyforgotpassword
           Data Passed in JSON Format
           if Email
           {
  		 "email":"saddam@gmail.com",
  		 "otp":"219433"
	    }
           if Mobile:
           {
   		"mobile":"+919655373273",
   		"countrycode":"+91",
   		"otp":"219433"
	   }
     --Enter New Password
           -->http://ip:8080/newpassword
           if Email
           {
  		 "email":"saddam@gmail.com",
  		 "password":"saddamAws"
	    }
           if Mobile:
           {
  		 "mobile":"+919655373273",
  		 "countrycode":"+91",
  		 "password":"123456789"
	    }

UPLOAD VIDEOS BY USER
      -->http://ip:8080/upload
      Data Passed in Form Format 
     |-----------------------------------------------------------------------| 
     |FormField Name |  values                                               |
     |---------------|-------------------------------------------------------|
     | title         |  "this is a sample title"                             |
     | description   |  "this is a sample Description"                       |
     | language      |  "tamil"                                              |
     | genres        |  "fiction,supernatural,kids,thriller,horror,comedy"   |--->check box data passed as a string and each selected genres is separated by comma(,).
     | agegroup      |  "kids"                                               |
     | tags          |  "#sha,#music,#movie,#drama"                          |--->HashTags are passed as a string and each tag separated by camma(,).
     | myfile        |  Video file                                           |
     | myimage       |  thumnail image                                       |
     |-----------------------------------------------------------------------|
   
