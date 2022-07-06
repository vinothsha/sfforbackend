package main

import (
	"log"
	"sha/cassession"
)

func main() {
	err := cassession.Session.Query("CREATE TABLE IF NOT EXISTS signup(uid uuid primary key,usermail varchar,countrycode text,mobile text,createddatetime text,password text);").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS videos(videouid uuid primary key,videolink text,videosizeinmb double,title text,description text,language text,genres text,agegroup text,createddatetime text,useruid uuid,tags text,thumnail text,etag text);").Exec()
	if err != nil {
		log.Println(err)
		return
	}

	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS addplaylist(addplaylistid uuid primary key,videoid uuid,userid uuid,playlistname text,datetime text)").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS comments(commentid uuid primary key,creatorid uuid,comment text,polarity text,datetime text,videoid uuid);").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS otp(uid uuid primary key,usermail varchar,countrycode text,mobile text,otp text);").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS languagegenres(id int primary key,languages list<text>,genres list<text>);").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = cassession.Session.Query("insert into languagegenres(id,languages,genres)values(1,['English','Hindi','Tamil','Telugu','Malayalam','Kannada','Gujarati','Bengali','Urdu','Marathi','Kashmiri','Odia','Assamese','Chinese','Spanish','Arabic','Portuguese','Russian','Japanese','French'],['FANTASY','ACTION','HORROR','MYSTERY','GENERAL FICTION','ADVENTURE','COMEDY','ROMANCE','THRILLER','NON FICTION','BUSINESS','MYTHOLOGY','LIFE STYLE','INSPIRATION','BIOGRAPHY']);").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS userprofiledetails(profileuid uuid primary key,useruid uuid,firstname text,lastname text,DateOfBirth text,Gender text,Email varchar,countrycode text,Mobile text,Profileimage text,);;").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS followers(Userid uuid primary key,Followers map<uuid,text>)").Exec()
	if err != nil {
		log.Panicln(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS Likes(Videoid uuid primary key,likes map<uuid,text>)").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE  if not exists trending (id uuid PRIMARY KEY ,views varint,up_date text,date varint)").Exec()
	if err != nil {
		log.Panicln(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS views(videoid uuid primary key,userid map<uuid,text> , views int);").Exec()
	if err != nil {
		log.Panicln(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS RecentlyWatched(recentlywatcheduid uuid ,userid uuid,videoid uuid,datetime text,PRIMARY KEY (recentlywatcheduid, datetime))WITH CLUSTERING ORDER BY (datetime DESC);").Exec()
	if err != nil {
		log.Panicln(err)
		return
	}
	// err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS sf.countrystate (id int PRIMARY KEY,countrycode text,countryname text,phonecode int,states list<text>);").Exec()
	// if err != nil {
	// 	log.Panic(err)
	// 	return
	// }

	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS ratings (videoid uuid PRIMARY KEY, rating map<uuid,int> );").Exec()
	if err != nil {
		log.Panic(err)
		return
	}

	// if err = cassession.Session.Query("UPDATE followers SET Followers = Followers + {'sub'}  WHERE Userid = now();").Exec(); err != nil {
	// 	fmt.Println(err)
	// }

	// err = cassession.Session.Query("create table thumnail(thumnailuid uuid primary key,videouid uuid,thumnail blob,useruid uuid); ;").Exec()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
}
