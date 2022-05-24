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
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS videos(videouid uuid primary key,videolink text,videosizeinmb double,title text,description text,language text,genres list<text>,agegroup text,createddatetime text,useruid uuid,tags list<text>,thumnail text,etag text);").Exec()
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
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS userprofiledetails(profileuid uuid primary key,useruid uuid,firstname text,lastname text,DateOfBirth text,Gender text,Email varchar,Mobile text,CountryCode text,State text,country text);;").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS Likes(useruid uuid primary key,videouid uuid)").Exec()
	if err != nil {
		log.Panicln(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS Views(viewuid uuid primary key,videouid uuid);").Exec()
	if err != nil {
		log.Panicln(err)
		return
	}
	err = cassession.Session.Query("CREATE TABLE IF NOT EXISTS RecentlyWatched(recentlywatcheduid uuid ,userid uuid,videoid uuid,datetime text,PRIMARY KEY (recentlywatcheduid, datetime))WITH CLUSTERING ORDER BY (datetime DESC);").Exec()
	if err != nil {
		log.Panicln(err)
		return
	}
	// err = cassession.Session.Query("create table thumnail(thumnailuid uuid primary key,videouid uuid,thumnail blob,useruid uuid); ;").Exec()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
}
