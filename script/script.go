package main

import (
	"fmt"
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
	err = cassession.Session.Query("CREATE TABLE sf.countrystate (id int PRIMARY KEY,countrycode text,countryname text,phonecode int,states list<text>);").Exec()
	if err != nil {
		log.Panic(err)
		return
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(1,'Afghanistan','AF',93,['Badakhshan', 'Badgis', 'Baglan', 'Balkh', 'Bamiyan', 'Farah', 'Faryab', 'Gawr', 'Gazni', 'Herat', 'Hilmand', 'Jawzjan', 'Kabul', 'Kapisa', 'Khawst', 'Kunar', 'Lagman', 'Lawghar', 'Nangarhar', 'Nimruz', 'Nuristan', 'Paktika', 'Paktiya', 'Parwan', 'Qandahar', 'Qunduz', 'Samangan', 'Sar-e Pul', 'Takhar', 'Uruzgan', 'Wardag', 'Zabul']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(2,'Albania','AL',355,['Berat', 'Bulqize', 'Delvine', 'Devoll', 'Dibre', 'Durres', 'Elbasan', 'Fier', 'Gjirokaster', 'Gramsh', 'Has', 'Kavaje', 'Kolonje', 'Korce', 'Kruje', 'Kucove', 'Kukes', 'Kurbin', 'Lezhe', 'Librazhd', 'Lushnje', 'Mallakaster', 'Malsi e Madhe', 'Mat', 'Mirdite', 'Peqin', 'Permet', 'Pogradec', 'Puke', 'Sarande', 'Shkoder', 'Skrapar', 'Tepelene', 'Tirane', 'Tropoje', 'Vlore']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(3,'Algeria','DZ',213,['Ayn Daflah', 'Ayn Tamushanat', 'Adrar', 'Algiers', 'Annabah', 'Bashshar', 'Batnah', 'Bijayah', 'Biskrah', 'Blidah', 'Buirah', 'Bumardas', 'Burj Bu Arririj', 'Ghalizan', 'Ghardayah', 'Ilizi', 'Jijili', 'Jilfah', 'Khanshalah', 'Masilah', 'Midyah', 'Milah', 'Muaskar', 'Mustaghanam', 'Naama', 'Oran', 'Ouargla', 'Qalmah', 'Qustantinah', 'Sakikdah', 'Satif', 'Sayda', 'Sidi ban-al-Abbas', 'Suq Ahras', 'Tamanghasat', 'Tibazah', 'Tibissah', 'Tilimsan', 'Tinduf', 'Tisamsilt', 'Tiyarat', 'Tizi Wazu', 'Umm-al-Bawaghi', 'Wahran', 'Warqla', 'Wilaya d Alger', 'Wilaya de Bejaia', 'Wilaya de Constantine', 'al-Aghwat', 'al-Bayadh', 'al-Jazair', 'al-Wad', 'ash-Shalif', 'at-Tarif']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(4,'American Samoa','AS',1684,['Eastern', 'Manua', 'Swains Island', 'Western']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(5,'Andorra','AD',376,['Andorra la Vella', 'Canillo', 'Encamp', 'La Massana', 'Les Escaldes', 'Ordino', 'Sant Julia de Loria']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(6,'Angola','AO',244,['Bengo', 'Benguela', 'Bie', 'Cabinda', 'Cunene', 'Huambo', 'Huila', 'Kuando-Kubango', 'Kwanza Norte', 'Kwanza Sul', 'Luanda', 'Lunda Norte', 'Lunda Sul', 'Malanje', 'Moxico', 'Namibe', 'Uige', 'Zaire']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(7,'Anguilla','AI',1264,['Other Provinces']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(8,'Antarctica','AQ',0,['Sector claimed by Argentina/Ch', 'Sector claimed by Argentina/UK', 'Sector claimed by Australia', 'Sector claimed by France', 'Sector claimed by New Zealand', 'Sector claimed by Norway', 'Unclaimed Sector']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(9,'Antigua And Barbuda','AG',1268,['Barbuda', 'Saint George', 'Saint John', 'Saint Mary', 'Saint Paul', 'Saint Peter', 'Saint Philip']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(10,'Argentina','AR',54,['Buenos Aires', 'Catamarca', 'Chaco', 'Chubut', 'Cordoba', 'Corrientes', 'Distrito Federal', 'Entre Rios', 'Formosa', 'Jujuy', 'La Pampa', 'La Rioja', 'Mendoza', 'Misiones', 'Neuquen', 'Rio Negro', 'Salta', 'San Juan', 'San Luis', 'Santa Cruz', 'Santa Fe', 'Santiago del Estero', 'Tierra del Fuego', 'Tucuman']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(11,'Armenia','AM',374,['Aragatsotn', 'Ararat', 'Armavir', 'Gegharkunik', 'Kotaik', 'Lori', 'Shirak', 'Stepanakert', 'Syunik', 'Tavush', 'Vayots Dzor', 'Yerevan']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(12,'Aruba','AW',297,['Aruba']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(13,'Australia','AU',61,['Auckland', 'Australian Capital Territory', 'Balgowlah', 'Balmain', 'Bankstown', 'Baulkham Hills', 'Bonnet Bay', 'Camberwell', 'Carole Park', 'Castle Hill', 'Caulfield', 'Chatswood', 'Cheltenham', 'Cherrybrook', 'Clayton', 'Collingwood', 'Frenchs Forest', 'Hawthorn', 'Jannnali', 'Knoxfield', 'Melbourne', 'New South Wales', 'Northern Territory', 'Perth', 'Queensland', 'South Australia', 'Tasmania', 'Templestowe', 'Victoria', 'Werribee south', 'Western Australia', 'Wheeler']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(14,'Austria','AT',43,['Bundesland Salzburg', 'Bundesland Steiermark', 'Bundesland Tirol', 'Burgenland', 'Carinthia', 'Karnten', 'Liezen', 'Lower Austria', 'Niederosterreich', 'Oberosterreich', 'Salzburg', 'Schleswig-Holstein', 'Steiermark', 'Styria', 'Tirol', 'Upper Austria', 'Vorarlberg', 'Wien']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(15,'Azerbaijan','AZ',994,['Abseron', 'Baki Sahari', 'Ganca', 'Ganja', 'Kalbacar', 'Lankaran', 'Mil-Qarabax', 'Mugan-Salyan', 'Nagorni-Qarabax', 'Naxcivan', 'Priaraks', 'Qazax', 'Saki', 'Sirvan', 'Xacmaz']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(16,'Bahamas The','BS',1242,['Abaco', 'Acklins Island', 'Andros', 'Berry Islands', 'Biminis', 'Cat Island', 'Crooked Island', 'Eleuthera', 'Exuma and Cays', 'Grand Bahama', 'Inagua Islands', 'Long Island', 'Mayaguana', 'New Providence', 'Ragged Island', 'Rum Cay', 'San Salvador']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(17,'Bahrain','BH',973,['Isa', 'Badiyah', 'Hidd', 'Jidd Hafs', 'Mahama', 'Manama', 'Sitrah', 'al-Manamah', 'al-Muharraq', 'ar-Rifaa']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(18,'Bangladesh','BD',880,['Bagar Hat', 'Bandarban', 'Barguna', 'Barisal', 'Bhola', 'Bogora', 'Brahman Bariya', 'Chandpur', 'Chattagam', 'Chittagong Division', 'Chuadanga', 'Dhaka', 'Dinajpur', 'Faridpur', 'Feni', 'Gaybanda', 'Gazipur', 'Gopalganj', 'Habiganj', 'Jaipur Hat', 'Jamalpur', 'Jessor', 'Jhalakati', 'Jhanaydah', 'Khagrachhari', 'Khulna', 'Kishorganj', 'Koks Bazar', 'Komilla', 'Kurigram', 'Kushtiya', 'Lakshmipur', 'Lalmanir Hat', 'Madaripur', 'Magura', 'Maimansingh', 'Manikganj', 'Maulvi Bazar', 'Meherpur', 'Munshiganj', 'Naral', 'Narayanganj', 'Narsingdi', 'Nator', 'Naugaon', 'Nawabganj', 'Netrakona', 'Nilphamari', 'Noakhali', 'Pabna', 'Panchagarh', 'Patuakhali', 'Pirojpur', 'Rajbari', 'Rajshahi', 'Rangamati', 'Rangpur', 'Satkhira', 'Shariatpur', 'Sherpur', 'Silhat', 'Sirajganj', 'Sunamganj', 'Tangayal', 'Thakurgaon']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(19,'Barbados','BB',1246,['Christ Church', 'Saint Andrew', 'Saint George', 'Saint James', 'Saint John', 'Saint Joseph', 'Saint Lucy', 'Saint Michael', 'Saint Peter', 'Saint Philip', 'Saint Thomas']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(20,'Belarus','BY',375,['Brest', 'Homjel', 'Hrodna', 'Mahiljow', 'Mahilyowskaya Voblasts', 'Minsk', 'Minskaja Voblasts', 'Petrik', 'Vicebsk']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(21,'Belgium','BE',32,['Antwerpen', 'Berchem', 'Brabant', 'Brabant Wallon', 'Brussel', 'East Flanders', 'Hainaut', 'Liege', 'Limburg', 'Luxembourg', 'Namur', 'Ontario', 'Oost-Vlaanderen', 'Provincie Brabant', 'Vlaams-Brabant', 'Wallonne', 'West-Vlaanderen']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(22,'Belize','BZ',501,['Belize', 'Cayo', 'Corozal', 'Orange Walk', 'Stann Creek', 'Toledo']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(23,'Benin','BJ',229,['Alibori', 'Atacora', 'Atlantique', 'Borgou', 'Collines', 'Couffo', 'Donga', 'Littoral', 'Mono', 'Oueme', 'Plateau', 'Zou']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(24,'Bermuda','BM',1441,['Hamilton', 'Saint George']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(25,'Bhutan','BT',975,['Bumthang', 'Chhukha', 'Chirang', 'Daga', 'Geylegphug', 'Ha', 'Lhuntshi', 'Mongar', 'Pemagatsel', 'Punakha', 'Rinpung', 'Samchi', 'Samdrup Jongkhar', 'Shemgang', 'Tashigang', 'Timphu', 'Tongsa', 'Wangdiphodrang']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(26,'Bolivia','BO',591,['Beni', 'Chuquisaca', 'Cochabamba', 'La Paz', 'Oruro', 'Pando', 'Potosi', 'Santa Cruz', 'Tarija']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(27,'Bosnia and Herzegovina','BA',387,['Federacija Bosna i Hercegovina', 'Republika Srpska']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(28,'Botswana','BW',267,['Central Bobonong', 'Central Boteti', 'Central Mahalapye', 'Central Serowe-Palapye', 'Central Tutume', 'Chobe', 'Francistown', 'Gaborone', 'Ghanzi', 'Jwaneng', 'Kgalagadi North', 'Kgalagadi South', 'Kgatleng', 'Kweneng', 'Lobatse', 'Ngamiland', 'Ngwaketse', 'North East', 'Okavango', 'Orapa', 'Selibe Phikwe', 'South East', 'Sowa']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(29,'Bouvet Island','BV',0,['Bouvet Island']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(30,'Brazil','BR',55,['Acre', 'Alagoas', 'Amapa', 'Amazonas', 'Bahia', 'Ceara', 'Distrito Federal', 'Espirito Santo', 'Estado de Sao Paulo', 'Goias', 'Maranhao', 'Mato Grosso', 'Mato Grosso do Sul', 'Minas Gerais', 'Para', 'Paraiba', 'Parana', 'Pernambuco', 'Piaui', 'Rio Grande do Norte', 'Rio Grande do Sul', 'Rio de Janeiro', 'Rondonia', 'Roraima', 'Santa Catarina', 'Sao Paulo', 'Sergipe', 'Tocantins']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(31,'British Indian Ocean Territory','IO',246,['British Indian Ocean Territory']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(32,'Brunei','BN',673,['Belait', 'Brunei-Muara', 'Temburong', 'Tutong']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(33,'Bulgaria','BG',359,['Blagoevgrad', 'Burgas', 'Dobrich', 'Gabrovo', 'Haskovo', 'Jambol', 'Kardzhali', 'Kjustendil', 'Lovech', 'Montana', 'Oblast Sofiya-Grad', 'Pazardzhik', 'Pernik', 'Pleven', 'Plovdiv', 'Razgrad', 'Ruse', 'Shumen', 'Silistra', 'Sliven', 'Smoljan', 'Sofija grad', 'Sofijska oblast', 'Stara Zagora', 'Targovishte', 'Varna', 'Veliko Tarnovo', 'Vidin', 'Vraca', 'Yablaniza']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(34,'Burkina Faso','BF',226,['Bale', 'Bam', 'Bazega', 'Bougouriba', 'Boulgou', 'Boulkiemde', 'Comoe', 'Ganzourgou', 'Gnagna', 'Gourma', 'Houet', 'Ioba', 'Kadiogo', 'Kenedougou', 'Komandjari', 'Kompienga', 'Kossi', 'Kouritenga', 'Kourweogo', 'Leraba', 'Mouhoun', 'Nahouri', 'Namentenga', 'Noumbiel', 'Oubritenga', 'Oudalan', 'Passore', 'Poni', 'Sanguie', 'Sanmatenga', 'Seno', 'Sissili', 'Soum', 'Sourou', 'Tapoa', 'Tuy', 'Yatenga', 'Zondoma', 'Zoundweogo']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(35,'Burundi','BI',257,['Bubanza', 'Bujumbura', 'Bururi', 'Cankuzo', 'Cibitoke', 'Gitega', 'Karuzi', 'Kayanza', 'Kirundo', 'Makamba', 'Muramvya', 'Muyinga', 'Ngozi', 'Rutana', 'Ruyigi']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(36,'Cambodia','KH',855,['Banteay Mean Chey', 'Bat Dambang', 'Kampong Cham', 'Kampong Chhnang', 'Kampong Spoeu', 'Kampong Thum', 'Kampot', 'Kandal', 'Kaoh Kong', 'Kracheh', 'Krong Kaeb', 'Krong Pailin', 'Krong Preah Sihanouk', 'Mondol Kiri', 'Otdar Mean Chey', 'Phnum Penh', 'Pousat', 'Preah Vihear', 'Prey Veaeng', 'Rotanak Kiri', 'Siem Reab', 'Stueng Traeng', 'Svay Rieng', 'Takaev']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(37,'Cameroon','CM',237,['Adamaoua', 'Centre', 'Est', 'Littoral', 'Nord', 'Nord Extreme', 'Nordouest', 'Ouest', 'Sud', 'Sudouest']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(38,'Canada','CA',1,['Alberta', 'British Columbia', 'Manitoba', 'New Brunswick', 'Newfoundland and Labrador', 'Northwest Territories', 'Nova Scotia', 'Nunavut', 'Ontario', 'Prince Edward Island', 'Quebec', 'Saskatchewan', 'Yukon']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(39,'Cape Verde','CV',238,['Boavista', 'Brava', 'Fogo', 'Maio', 'Sal', 'Santo Antao', 'Sao Nicolau', 'Sao Tiago', 'Sao Vicente']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(40,'Cayman Islands','KY',1345,['Grand Cayman']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(41,'Central African Republic','CF',236,['Bamingui-Bangoran', 'Bangui', 'Basse-Kotto', 'Haut-Mbomou', 'Haute-Kotto', 'Kemo', 'Lobaye', 'Mambere-Kadei', 'Mbomou', 'Nana-Gribizi', 'Nana-Mambere', 'Ombella Mpoko', 'Ouaka', 'Ouham', 'Ouham-Pende', 'Sangha-Mbaere', 'Vakaga']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(42,'Chad','TD',235,['Batha', 'Biltine', 'Bourkou-Ennedi-Tibesti', 'Chari-Baguirmi', 'Guera', 'Kanem', 'Lac', 'Logone Occidental', 'Logone Oriental', 'Mayo-Kebbi', 'Moyen-Chari', 'Ouaddai', 'Salamat', 'Tandjile']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(43,'Chile','CL',56,['Aisen', 'Antofagasta', 'Araucania', 'Atacama', 'Bio Bio', 'Coquimbo', 'Libertador General Bernardo O', 'Los Lagos', 'Magellanes', 'Maule', 'Metropolitana', 'Metropolitana de Santiago', 'Tarapaca', 'Valparaiso']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(44,'China','CN',86,['Anhui', 'Aomen', 'Beijing', 'Beijing Shi', 'Chongqing', 'Fujian', 'Gansu', 'Guangdong', 'Guangxi', 'Guizhou', 'Hainan', 'Hebei', 'Heilongjiang', 'Henan', 'Hubei', 'Hunan', 'Jiangsu', 'Jiangxi', 'Jilin', 'Liaoning', 'Nei Monggol', 'Ningxia Hui', 'Qinghai', 'Shaanxi', 'Shandong', 'Shanghai', 'Shanxi', 'Sichuan', 'Tianjin', 'Xianggang', 'Xinjiang', 'Xizang', 'Yunnan', 'Zhejiang']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(45,'Christmas Island','CX',61,['Christmas Island']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(46,'Cocos (Keeling) Islands','CC',672,['Cocos (Keeling) Islands']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(47,'Colombia','CO',57,['Amazonas', 'Antioquia', 'Arauca', 'Atlantico', 'Bogota', 'Bolivar', 'Boyaca', 'Caldas', 'Caqueta', 'Casanare', 'Cauca', 'Cesar', 'Choco', 'Cordoba', 'Cundinamarca', 'Guainia', 'Guaviare', 'Huila', 'La Guajira', 'Magdalena', 'Meta', 'Narino', 'Norte de Santander', 'Putumayo', 'Quindio', 'Risaralda', 'San Andres y Providencia', 'Santander', 'Sucre', 'Tolima', 'Valle del Cauca', 'Vaupes', 'Vichada']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(48,'Comoros','KM',269,['Mwali', 'Njazidja', 'Nzwani']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(49,'Republic Of The Congo','CG',242,['Bouenza', 'Brazzaville', 'Cuvette', 'Kouilou', 'Lekoumou', 'Likouala', 'Niari', 'Plateaux', 'Pool', 'Sangha']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(50,'Democratic Republic Of The Congo','CD',242,['Bandundu', 'Bas-Congo', 'Equateur', 'Haut-Congo', 'Kasai-Occidental', 'Kasai-Oriental', 'Katanga', 'Kinshasa', 'Maniema', 'Nord-Kivu', 'Sud-Kivu']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(51,'Cook Islands','CK',682,['Aitutaki', 'Atiu', 'Mangaia', 'Manihiki', 'Mauke', 'Mitiaro', 'Nassau', 'Pukapuka', 'Rakahanga', 'Rarotonga', 'Tongareva']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(52,'Costa Rica','CR',506,['Alajuela', 'Cartago', 'Guanacaste', 'Heredia', 'Limon', 'Puntarenas', 'San Jose']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(53,'Cote DIvoire (Ivory Coast)','CI',225,['Abidjan', 'Agneby', 'Bafing', 'Denguele', 'Dix-huit Montagnes', 'Fromager', 'Haut-Sassandra', 'Lacs', 'Lagunes', 'Marahoue', 'Moyen-Cavally', 'Moyen-Comoe', 'Nzi-Comoe', 'Sassandra', 'Savanes', 'Sud-Bandama', 'Sud-Comoe', 'Vallee du Bandama', 'Worodougou', 'Zanzan']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(54,'Croatia (Hrvatska)','HR',385,['Bjelovar-Bilogora', 'Dubrovnik-Neretva', 'Grad Zagreb', 'Istra', 'Karlovac', 'Koprivnica-Krizhevci', 'Krapina-Zagorje', 'Lika-Senj', 'Medhimurje', 'Medimurska Zupanija', 'Osijek-Baranja', 'Osjecko-Baranjska Zupanija', 'Pozhega-Slavonija', 'Primorje-Gorski Kotar', 'Shibenik-Knin', 'Sisak-Moslavina', 'Slavonski Brod-Posavina', 'Split-Dalmacija', 'Varazhdin', 'Virovitica-Podravina', 'Vukovar-Srijem', 'Zadar', 'Zagreb']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(55,'Cuba','CU',53,['Camaguey', 'Ciego de Avila', 'Cienfuegos', 'Ciudad de la Habana', 'Granma', 'Guantanamo', 'Habana', 'Holguin', 'Isla de la Juventud', 'La Habana', 'Las Tunas', 'Matanzas', 'Pinar del Rio', 'Sancti Spiritus', 'Santiago de Cuba', 'Villa Clara']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(56,'Cyprus','CY',357,['Government controlled area', 'Limassol', 'Nicosia District', 'Paphos', 'Turkish controlled area']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(57,'Czech Republic','CZ',420,['Central Bohemian', 'Frycovice', 'Jihocesky Kraj', 'Jihochesky', 'Jihomoravsky', 'Karlovarsky', 'Klecany', 'Kralovehradecky', 'Liberecky', 'Lipov', 'Moravskoslezsky', 'Olomoucky', 'Olomoucky Kraj', 'Pardubicky', 'Plzensky', 'Praha', 'Rajhrad', 'Smirice', 'South Moravian', 'Straz nad Nisou', 'Stredochesky', 'Unicov', 'Ustecky', 'Valletta', 'Velesin', 'Vysochina', 'Zlinsky']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(58,'Denmark','DK',45,['Arhus', 'Bornholm', 'Frederiksborg', 'Fyn', 'Hovedstaden', 'Kobenhavn', 'Kobenhavns Amt', 'Kobenhavns Kommune', 'Nordjylland', 'Ribe', 'Ringkobing', 'Roervig', 'Roskilde', 'Roslev', 'Sjaelland', 'Soeborg', 'Sonderjylland', 'Storstrom', 'Syddanmark', 'Toelloese', 'Vejle', 'Vestsjalland', 'Viborg']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(59,'Djibouti','DJ',253,['Ali Sabih', 'Dikhil', 'Jibuti', 'Tajurah', 'Ubuk']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(60,'Dominica','DM',1767,['Saint Andrew', 'Saint David', 'Saint George', 'Saint John', 'Saint Joseph', 'Saint Luke', 'Saint Mark', 'Saint Patrick', 'Saint Paul', 'Saint Peter']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(61,'Dominican Republic','DO',1809,['Azua', 'Bahoruco', 'Barahona', 'Dajabon', 'Distrito Nacional', 'Duarte', 'El Seybo', 'Elias Pina', 'Espaillat', 'Hato Mayor', 'Independencia', 'La Altagracia', 'La Romana', 'La Vega', 'Maria Trinidad Sanchez', 'Monsenor Nouel', 'Monte Cristi', 'Monte Plata', 'Pedernales', 'Peravia', 'Puerto Plata', 'Salcedo', 'Samana', 'San Cristobal', 'San Juan', 'San Pedro de Macoris', 'Sanchez Ramirez', 'Santiago', 'Santiago Rodriguez', 'Valverde']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(62,'East Timor','TP',670,['Aileu', 'Ainaro', 'Ambeno', 'Baucau', 'Bobonaro', 'Cova Lima', 'Dili', 'Ermera', 'Lautem', 'Liquica', 'Manatuto', 'Manufahi', 'Viqueque']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(63,'Ecuador','EC',593,['Azuay', 'Bolivar', 'Canar', 'Carchi', 'Chimborazo', 'Cotopaxi', 'El Oro', 'Esmeraldas', 'Galapagos', 'Guayas', 'Imbabura', 'Loja', 'Los Rios', 'Manabi', 'Morona Santiago', 'Napo', 'Orellana', 'Pastaza', 'Pichincha', 'Sucumbios', 'Tungurahua', 'Zamora Chinchipe']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(64,'Egypt','EG',20,['Aswan', 'Asyut', 'Bani Suwayf', 'Bur Said', 'Cairo', 'Dumyat', 'Kafr-ash-Shaykh', 'Matruh', 'Muhafazat ad Daqahliyah', 'Muhafazat al Fayyum', 'Muhafazat al Gharbiyah', 'Muhafazat al Iskandariyah', 'Muhafazat al Qahirah', 'Qina', 'Sawhaj', 'Sina al-Janubiyah', 'Sina ash-Shamaliyah', 'ad-Daqahliyah', 'al-Bahr-al-Ahmar', 'al-Buhayrah', 'al-Fayyum', 'al-Gharbiyah', 'al-Iskandariyah', 'al-Ismailiyah', 'al-Jizah', 'al-Minufiyah', 'al-Minya', 'al-Qahira', 'al-Qalyubiyah', 'al-Uqsur', 'al-Wadi al-Jadid', 'as-Suways', 'ash-Sharqiyah']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(65,'El Salvador','SV',503,['Ahuachapan', 'Cabanas', 'Chalatenango', 'Cuscatlan', 'La Libertad', 'La Paz', 'La Union', 'Morazan', 'San Miguel', 'San Salvador', 'San Vicente', 'Santa Ana', 'Sonsonate', 'Usulutan']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(66,'Equatorial Guinea','GQ',240,['Annobon', 'Bioko Norte', 'Bioko Sur', 'Centro Sur', 'Kie-Ntem', 'Litoral', 'Wele-Nzas']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(67,'Eritrea','ER',291,['Anseba', 'Debub', 'Debub-Keih-Bahri', 'Gash-Barka', 'Maekel', 'Semien-Keih-Bahri']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(68,'Estonia','EE',372,['Harju', 'Hiiu', 'Ida-Viru', 'Jarva', 'Jogeva', 'Laane', 'Laane-Viru', 'Parnu', 'Polva', 'Rapla', 'Saare', 'Tartu', 'Valga', 'Viljandi', 'Voru']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(69,'Ethiopia','ET',251,['Addis Abeba', 'Afar', 'Amhara', 'Benishangul', 'Diredawa', 'Gambella', 'Harar', 'Jigjiga', 'Mekele', 'Oromia', 'Somali', 'Southern', 'Tigray']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(70,'External Territories of Australia','XA',61,['Christmas Island', 'Cocos Islands', 'Coral Sea Islands']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(71,'Falkland Islands','FK',500,['Falkland Islands', 'South Georgia']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(72,'Faroe Islands','FO',298,['Klaksvik', 'Nor ara Eysturoy', 'Nor oy', 'Sandoy', 'Streymoy', 'Su uroy', 'Sy ra Eysturoy', 'Torshavn', 'Vaga']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(73,'Fiji Islands','FJ',679,['Central', 'Eastern', 'Northern', 'South Pacific', 'Western']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(74,'Finland','FI',358,['Ahvenanmaa', 'Etela-Karjala', 'Etela-Pohjanmaa', 'Etela-Savo', 'Etela-Suomen Laani', 'Ita-Suomen Laani', 'Ita-Uusimaa', 'Kainuu', 'Kanta-Hame', 'Keski-Pohjanmaa', 'Keski-Suomi', 'Kymenlaakso', 'Lansi-Suomen Laani', 'Lappi', 'Northern Savonia', 'Ostrobothnia', 'Oulun Laani', 'Paijat-Hame', 'Pirkanmaa', 'Pohjanmaa', 'Pohjois-Karjala', 'Pohjois-Pohjanmaa', 'Pohjois-Savo', 'Saarijarvi', 'Satakunta', 'Southern Savonia', 'Tavastia Proper', 'Uleaborgs Lan', 'Uusimaa', 'Varsinais-Suomi']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(75,'France','FR',33,['Ain', 'Aisne', 'Albi Le Sequestre', 'Allier', 'Alpes-Cote dAzur', 'Alpes-Maritimes', 'Alpes-de-Haute-Provence', 'Alsace', 'Aquitaine', 'Ardeche', 'Ardennes', 'Ariege', 'Aube', 'Aude', 'Auvergne', 'Aveyron', 'Bas-Rhin', 'Basse-Normandie', 'Bouches-du-Rhone', 'Bourgogne', 'Bretagne', 'Brittany', 'Burgundy', 'Calvados', 'Cantal', 'Cedex', 'Centre', 'Charente', 'Charente-Maritime', 'Cher', 'Correze', 'Corse-du-Sud', 'Cote-dOr', 'Cotes-dArmor', 'Creuse', 'Crolles', 'Deux-Sevres', 'Dordogne', 'Doubs', 'Drome', 'Essonne', 'Eure', 'Eure-et-Loir', 'Feucherolles', 'Finistere', 'Franche-Comte', 'Gard', 'Gers', 'Gironde', 'Haut-Rhin', 'Haute-Corse', 'Haute-Garonne', 'Haute-Loire', 'Haute-Marne', 'Haute-Saone', 'Haute-Savoie', 'Haute-Vienne', 'Hautes-Alpes', 'Hautes-Pyrenees', 'Hauts-de-Seine', 'Herault', 'Ile-de-France', 'Ille-et-Vilaine', 'Indre', 'Indre-et-Loire', 'Isere', 'Jura', 'Klagenfurt', 'Landes', 'Languedoc-Roussillon', 'Larcay', 'Le Castellet', 'Le Creusot', 'Limousin', 'Loir-et-Cher', 'Loire', 'Loire-Atlantique', 'Loiret', 'Lorraine', 'Lot', 'Lot-et-Garonne', 'Lower Normandy', 'Lozere', 'Maine-et-Loire', 'Manche', 'Marne', 'Mayenne', 'Meurthe-et-Moselle', 'Meuse', 'Midi-Pyrenees', 'Morbihan', 'Moselle', 'Nievre', 'Nord', 'Nord-Pas-de-Calais', 'Oise', 'Orne', 'Paris', 'Pas-de-Calais', 'Pays de la Loire', 'Pays-de-la-Loire', 'Picardy', 'Puy-de-Dome', 'Pyrenees-Atlantiques', 'Pyrenees-Orientales', 'Quelmes', 'Rhone', 'Rhone-Alpes', 'Saint Ouen', 'Saint Viatre', 'Saone-et-Loire', 'Sarthe', 'Savoie', 'Seine-Maritime', 'Seine-Saint-Denis', 'Seine-et-Marne', 'Somme', 'Sophia Antipolis', 'Souvans', 'Tarn', 'Tarn-et-Garonne', 'Territoire de Belfort', 'Treignac', 'Upper Normandy', 'Val-dOise', 'Val-de-Marne', 'Var', 'Vaucluse', 'Vellise', 'Vendee', 'Vienne', 'Vosges', 'Yonne', 'Yvelines']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(76,'French Guiana','GF',594,['Cayenne', 'Saint-Laurent-du-Maroni']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(77,'French Polynesia','PF',689,['Iles du Vent', 'Iles sous le Vent', 'Marquesas', 'Tuamotu', 'Tubuai']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(78,'French Southern Territories','TF',0,['Amsterdam', 'Crozet Islands', 'Kerguelen']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(79,'Gabon','GA',241,['Estuaire', 'Haut-Ogooue', 'Moyen-Ogooue', 'Ngounie', 'Nyanga', 'Ogooue-Ivindo', 'Ogooue-Lolo', 'Ogooue-Maritime', 'Woleu-Ntem']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(80,'Gambia The','GM',220,['Banjul', 'Basse', 'Brikama', 'Janjanbureh', 'Kanifing', 'Kerewan', 'Kuntaur', 'Mansakonko']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(81,'Georgia','GE',995,['Abhasia', 'Ajaria', 'Guria', 'Imereti', 'Kaheti', 'Kvemo Kartli', 'Mcheta-Mtianeti', 'Racha', 'Samagrelo-Zemo Svaneti', 'Samche-Zhavaheti', 'Shida Kartli', 'Tbilisi']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(82,'Germany','DE',49,['Auvergne', 'Baden-Wurttemberg', 'Bavaria', 'Bayern', 'Beilstein Wurtt', 'Berlin', 'Brandenburg', 'Bremen', 'Dreisbach', 'Freistaat Bayern', 'Hamburg', 'Hannover', 'Heroldstatt', 'Hessen', 'Kortenberg', 'Laasdorf', 'Land Baden-Wurttemberg', 'Land Bayern', 'Land Brandenburg', 'Land Hessen', 'Land Mecklenburg-Vorpommern', 'Land Nordrhein-Westfalen', 'Land Rheinland-Pfalz', 'Land Sachsen', 'Land Sachsen-Anhalt', 'Land Thuringen', 'Lower Saxony', 'Mecklenburg-Vorpommern', 'Mulfingen', 'Munich', 'Neubeuern', 'Niedersachsen', 'Noord-Holland', 'Nordrhein-Westfalen', 'North Rhine-Westphalia', 'Osterode', 'Rheinland-Pfalz', 'Rhineland-Palatinate', 'Saarland', 'Sachsen', 'Sachsen-Anhalt', 'Saxony', 'Schleswig-Holstein', 'Thuringia', 'Webling', 'Weinstrabe', 'schlobborn']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(83,'Ghana','GH',233,['Ashanti', 'Brong-Ahafo', 'Central', 'Eastern', 'Greater Accra', 'Northern', 'Upper East', 'Upper West', 'Volta', 'Western']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(84,'Gibraltar','GI',350,['Gibraltar']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(85,'Greece','GR',30,['Acharnes', 'Ahaia', 'Aitolia kai Akarnania', 'Argolis', 'Arkadia', 'Arta', 'Attica', 'Attiki', 'Ayion Oros', 'Crete', 'Dodekanisos', 'Drama', 'Evia', 'Evritania', 'Evros', 'Evvoia', 'Florina', 'Fokis', 'Fthiotis', 'Grevena', 'Halandri', 'Halkidiki', 'Hania', 'Heraklion', 'Hios', 'Ilia', 'Imathia', 'Ioannina', 'Iraklion', 'Karditsa', 'Kastoria', 'Kavala', 'Kefallinia', 'Kerkira', 'Kiklades', 'Kilkis', 'Korinthia', 'Kozani', 'Lakonia', 'Larisa', 'Lasithi', 'Lesvos', 'Levkas', 'Magnisia', 'Messinia', 'Nomos Attikis', 'Nomos Zakynthou', 'Pella', 'Pieria', 'Piraios', 'Preveza', 'Rethimni', 'Rodopi', 'Samos', 'Serrai', 'Thesprotia', 'Thessaloniki', 'Trikala', 'Voiotia', 'West Greece', 'Xanthi', 'Zakinthos']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(86,'Greenland','GL',299,['Aasiaat', 'Ammassalik', 'Illoqqortoormiut', 'Ilulissat', 'Ivittuut', 'Kangaatsiaq', 'Maniitsoq', 'Nanortalik', 'Narsaq', 'Nuuk', 'Paamiut', 'Qaanaaq', 'Qaqortoq', 'Qasigiannguit', 'Qeqertarsuaq', 'Sisimiut', 'Udenfor kommunal inddeling', 'Upernavik', 'Uummannaq']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(87,'Grenada','GD',1473,['Carriacou-Petite Martinique', 'Saint Andrew', 'Saint Davids', 'Saint Georges', 'Saint John', 'Saint Mark', 'Saint Patrick']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(88,'Guadeloupe','GP',590,['Basse-Terre', 'Grande-Terre', 'Iles des Saintes', 'La Desirade', 'Marie-Galante', 'Saint Barthelemy', 'Saint Martin']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(89,'Guam','GU',1671,['Agana Heights', 'Agat', 'Barrigada', 'Chalan-Pago-Ordot', 'Dededo', 'Hagatna', 'Inarajan', 'Mangilao', 'Merizo', 'Mongmong-Toto-Maite', 'Santa Rita', 'Sinajana', 'Talofofo', 'Tamuning', 'Yigo', 'Yona']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(90,'Guatemala','GT',502,['Alta Verapaz', 'Baja Verapaz', 'Chimaltenango', 'Chiquimula', 'El Progreso', 'Escuintla', 'Guatemala', 'Huehuetenango', 'Izabal', 'Jalapa', 'Jutiapa', 'Peten', 'Quezaltenango', 'Quiche', 'Retalhuleu', 'Sacatepequez', 'San Marcos', 'Santa Rosa', 'Solola', 'Suchitepequez', 'Totonicapan', 'Zacapa']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(91,'Guernsey and Alderney','XU',44,['Alderney', 'Castel', 'Forest', 'Saint Andrew', 'Saint Martin', 'Saint Peter Port', 'Saint Pierre du Bois', 'Saint Sampson', 'Saint Saviour', 'Sark', 'Torteval', 'Vale']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(92,'Guinea','GN',224,['Beyla', 'Boffa', 'Boke', 'Conakry', 'Coyah', 'Dabola', 'Dalaba', 'Dinguiraye', 'Faranah', 'Forecariah', 'Fria', 'Gaoual', 'Gueckedou', 'Kankan', 'Kerouane', 'Kindia', 'Kissidougou', 'Koubia', 'Koundara', 'Kouroussa', 'Labe', 'Lola', 'Macenta', 'Mali', 'Mamou', 'Mandiana', 'Nzerekore', 'Pita', 'Siguiri', 'Telimele', 'Tougue', 'Yomou']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(93,'Guinea-Bissau','GW',245,['Bafata', 'Bissau', 'Bolama', 'Cacheu', 'Gabu', 'Oio', 'Quinara', 'Tombali']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(94,'Guyana','GY',592,['Barima-Waini', 'Cuyuni-Mazaruni', 'Demerara-Mahaica', 'East Berbice-Corentyne', 'Essequibo Islands-West Demerar', 'Mahaica-Berbice', 'Pomeroon-Supenaam', 'Potaro-Siparuni', 'Upper Demerara-Berbice', 'Upper Takutu-Upper Essequibo']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(95,'Haiti','HT',509,['Artibonite', 'Centre', 'GrandAnse', 'Nord', 'Nord-Est', 'Nord-Ouest', 'Ouest', 'Sud', 'Sud-Est']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(96,'Heard and McDonald Islands','HM',0,['Heard and McDonald Islands']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(97,'Honduras','HN',504,['Atlantida', 'Choluteca', 'Colon', 'Comayagua', 'Copan', 'Cortes', 'Distrito Central', 'El Paraiso', 'Francisco Morazan', 'Gracias a Dios', 'Intibuca', 'Islas de la Bahia', 'La Paz', 'Lempira', 'Ocotepeque', 'Olancho', 'Santa Barbara', 'Valle', 'Yoro']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(98,'Hong Kong S.A.R.','HK',852,['Hong Kong']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(99,'Hungary','HU',36,['Bacs-Kiskun', 'Baranya', 'Bekes', 'Borsod-Abauj-Zemplen', 'Budapest', 'Csongrad', 'Fejer', 'Gyor-Moson-Sopron', 'Hajdu-Bihar', 'Heves', 'Jasz-Nagykun-Szolnok', 'Komarom-Esztergom', 'Nograd', 'Pest', 'Somogy', 'Szabolcs-Szatmar-Bereg', 'Tolna', 'Vas', 'Veszprem', 'Zala']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(100,'Iceland','IS',354,['Austurland', 'Gullbringusysla', 'Hofu borgarsva i', 'Nor urland eystra', 'Nor urland vestra', 'Su urland', 'Su urnes', 'Vestfir ir', 'Vesturland']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(101,'India','IN',91,['Andaman and Nicobar Islands', 'Andhra Pradesh', 'Arunachal Pradesh', 'Assam', 'Bihar', 'Chandigarh', 'Chhattisgarh', 'Dadra and Nagar Haveli', 'Daman and Diu', 'Delhi', 'Goa', 'Gujarat', 'Haryana', 'Himachal Pradesh', 'Jammu and Kashmir', 'Jharkhand', 'Karnataka', 'Kenmore', 'Kerala', 'Lakshadweep', 'Madhya Pradesh', 'Maharashtra', 'Manipur', 'Meghalaya', 'Mizoram', 'Nagaland', 'Narora', 'Natwar', 'Odisha', 'Paschim Medinipur', 'Pondicherry', 'Punjab', 'Rajasthan', 'Sikkim', 'Tamil Nadu', 'Telangana', 'Tripura', 'Uttar Pradesh', 'Uttarakhand', 'Vaishali', 'West Bengal']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(102,'Indonesia','ID',62,['Aceh', 'Bali', 'Bangka-Belitung', 'Banten', 'Bengkulu', 'Gandaria', 'Gorontalo', 'Jakarta', 'Jambi', 'Jawa Barat', 'Jawa Tengah', 'Jawa Timur', 'Kalimantan Barat', 'Kalimantan Selatan', 'Kalimantan Tengah', 'Kalimantan Timur', 'Kendal', 'Lampung', 'Maluku', 'Maluku Utara', 'Nusa Tenggara Barat', 'Nusa Tenggara Timur', 'Papua', 'Riau', 'Riau Kepulauan', 'Solo', 'Sulawesi Selatan', 'Sulawesi Tengah', 'Sulawesi Tenggara', 'Sulawesi Utara', 'Sumatera Barat', 'Sumatera Selatan', 'Sumatera Utara', 'Yogyakarta']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(103,'Iran','IR',98,['Ardabil', 'Azarbayjan-e Bakhtari', 'Azarbayjan-e Khavari', 'Bushehr', 'Chahar Mahal-e Bakhtiari', 'Esfahan', 'Fars', 'Gilan', 'Golestan', 'Hamadan', 'Hormozgan', 'Ilam', 'Kerman', 'Kermanshah', 'Khorasan', 'Khuzestan', 'Kohgiluyeh-e Boyerahmad', 'Kordestan', 'Lorestan', 'Markazi', 'Mazandaran', 'Ostan-e Esfahan', 'Qazvin', 'Qom', 'Semnan', 'Sistan-e Baluchestan', 'Tehran', 'Yazd', 'Zanjan']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(104,'Iraq','IQ',964,['Babil', 'Baghdad', 'Dahuk', 'Dhi Qar', 'Diyala', 'Erbil', 'Irbil', 'Karbala', 'Kurdistan', 'Maysan', 'Ninawa', 'Salah-ad-Din', 'Wasit', 'al-Anbar', 'al-Basrah', 'al-Muthanna', 'al-Qadisiyah', 'an-Najaf', 'as-Sulaymaniyah', 'at-Tamim']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(105,'Ireland','IE',353,['Armagh', 'Carlow', 'Cavan', 'Clare', 'Cork', 'Donegal', 'Dublin', 'Galway', 'Kerry', 'Kildare', 'Kilkenny', 'Laois', 'Leinster', 'Leitrim', 'Limerick', 'Loch Garman', 'Longford', 'Louth', 'Mayo', 'Meath', 'Monaghan', 'Offaly', 'Roscommon', 'Sligo', 'Tipperary North Riding', 'Tipperary South Riding', 'Ulster', 'Waterford', 'Westmeath', 'Wexford', 'Wicklow']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(106,'Israel','IL',972,['Beit Hanania', 'Ben Gurion Airport', 'Bethlehem', 'Caesarea', 'Centre', 'Gaza', 'Hadaron', 'Haifa District', 'Hamerkaz', 'Hazafon', 'Hebron', 'Jaffa', 'Jerusalem', 'Khefa', 'Kiryat Yam', 'Lower Galilee', 'Qalqilya', 'Talme Elazar', 'Tel Aviv', 'Tsafon', 'Umm El Fahem', 'Yerushalayim']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(107,'Italy','IT',39,['Abruzzi', 'Abruzzo', 'Agrigento', 'Alessandria', 'Ancona', 'Arezzo', 'Ascoli Piceno', 'Asti', 'Avellino', 'Bari', 'Basilicata', 'Belluno', 'Benevento', 'Bergamo', 'Biella', 'Bologna', 'Bolzano', 'Brescia', 'Brindisi', 'Calabria', 'Campania', 'Cartoceto', 'Caserta', 'Catania', 'Chieti', 'Como', 'Cosenza', 'Cremona', 'Cuneo', 'Emilia-Romagna', 'Ferrara', 'Firenze', 'Florence', 'Forli-Cesena ', 'Friuli-Venezia Giulia', 'Frosinone', 'Genoa', 'Gorizia', 'LAquila', 'Lazio', 'Lecce', 'Lecco', 'Liguria', 'Lodi', 'Lombardia', 'Lombardy', 'Macerata', 'Mantova', 'Marche', 'Messina', 'Milan', 'Modena', 'Molise', 'Molteno', 'Montenegro', 'Monza and Brianza', 'Naples', 'Novara', 'Padova', 'Parma', 'Pavia', 'Perugia', 'Pesaro-Urbino', 'Piacenza', 'Piedmont', 'Piemonte', 'Pisa', 'Pordenone', 'Potenza', 'Puglia', 'Reggio Emilia', 'Rimini', 'Roma', 'Salerno', 'Sardegna', 'Sassari', 'Savona', 'Sicilia', 'Siena', 'Sondrio', 'South Tyrol', 'Taranto', 'Teramo', 'Torino', 'Toscana', 'Trapani', 'Trentino-Alto Adige', 'Trento', 'Treviso', 'Udine', 'Umbria', 'Valle dAosta', 'Varese', 'Veneto', 'Venezia', 'Verbano-Cusio-Ossola', 'Vercelli', 'Verona', 'Vicenza', 'Viterbo']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(108,'Jamaica','JM',1876,['Buxoro Viloyati', 'Clarendon', 'Hanover', 'Kingston', 'Manchester', 'Portland', 'Saint Andrews', 'Saint Ann', 'Saint Catherine', 'Saint Elizabeth', 'Saint James', 'Saint Mary', 'Saint Thomas', 'Trelawney', 'Westmoreland']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(109,'Japan','JP',81,['Aichi', 'Akita', 'Aomori', 'Chiba', 'Ehime', 'Fukui', 'Fukuoka', 'Fukushima', 'Gifu', 'Gumma', 'Hiroshima', 'Hokkaido', 'Hyogo', 'Ibaraki', 'Ishikawa', 'Iwate', 'Kagawa', 'Kagoshima', 'Kanagawa', 'Kanto', 'Kochi', 'Kumamoto', 'Kyoto', 'Mie', 'Miyagi', 'Miyazaki', 'Nagano', 'Nagasaki', 'Nara', 'Niigata', 'Oita', 'Okayama', 'Okinawa', 'Osaka', 'Saga', 'Saitama', 'Shiga', 'Shimane', 'Shizuoka', 'Tochigi', 'Tokushima', 'Tokyo', 'Tottori', 'Toyama', 'Wakayama', 'Yamagata', 'Yamaguchi', 'Yamanashi']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(110,'Jersey','XJ',44,['Grouville', 'Saint Brelade', 'Saint Clement', 'Saint Helier', 'Saint John', 'Saint Lawrence', 'Saint Martin', 'Saint Mary', 'Saint Peter', 'Saint Saviour', 'Trinity']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(111,'Jordan','JO',962,['Ajlun', 'Amman', 'Irbid', 'Jarash', 'Maan', 'Madaba', 'al-Aqabah', 'al-Balqa', 'al-Karak', 'al-Mafraq', 'at-Tafilah', 'az-Zarqa']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(112,'Kazakhstan','KZ',7,['Akmecet', 'Akmola', 'Aktobe', 'Almati', 'Atirau', 'Batis Kazakstan', 'Burlinsky Region', 'Karagandi', 'Kostanay', 'Mankistau', 'Ontustik Kazakstan', 'Pavlodar', 'Sigis Kazakstan', 'Soltustik Kazakstan', 'Taraz']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(113,'Kenya','KE',254,['Central', 'Coast', 'Eastern', 'Nairobi', 'North Eastern', 'Nyanza', 'Rift Valley', 'Western']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(114,'Kiribati','KI',686,['Abaiang', 'Abemana', 'Aranuka', 'Arorae', 'Banaba', 'Beru', 'Butaritari', 'Kiritimati', 'Kuria', 'Maiana', 'Makin', 'Marakei', 'Nikunau', 'Nonouti', 'Onotoa', 'Phoenix Islands', 'Tabiteuea North', 'Tabiteuea South', 'Tabuaeran', 'Tamana', 'Tarawa North', 'Tarawa South', 'Teraina']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(115,'Korea North','KP',850,['Chagangdo', 'Hamgyeongbukto', 'Hamgyeongnamdo', 'Hwanghaebukto', 'Hwanghaenamdo', 'Kaeseong', 'Kangweon', 'Nampo', 'Pyeonganbukto', 'Pyeongannamdo', 'Pyeongyang', 'Yanggang']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(116,'Korea South','KR',82,['Busan', 'Cheju', 'Chollabuk', 'Chollanam', 'Chungbuk', 'Chungcheongbuk', 'Chungcheongnam', 'Chungnam', 'Daegu', 'Gangwon-do', 'Goyang-si', 'Gyeonggi-do', 'Gyeongsang ', 'Gyeongsangnam-do', 'Incheon', 'Jeju-Si', 'Jeonbuk', 'Kangweon', 'Kwangju', 'Kyeonggi', 'Kyeongsangbuk', 'Kyeongsangnam', 'Kyonggi-do', 'Kyungbuk-Do', 'Kyunggi-Do', 'Kyunggi-do', 'Pusan', 'Seoul', 'Sudogwon', 'Taegu', 'Taejeon', 'Taejon-gwangyoksi', 'Ulsan', 'Wonju', 'gwangyoksi']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(117,'Kuwait','KW',965,['Al Asimah', 'Hawalli', 'Mishref', 'Qadesiya', 'Safat', 'Salmiya', 'al-Ahmadi', 'al-Farwaniyah', 'al-Jahra', 'al-Kuwayt']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(118,'Kyrgyzstan','KG',996,['Batken', 'Bishkek', 'Chui', 'Issyk-Kul', 'Jalal-Abad', 'Naryn', 'Osh', 'Talas']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(119,'Laos','LA',856,['Attopu', 'Bokeo', 'Bolikhamsay', 'Champasak', 'Houaphanh', 'Khammouane', 'Luang Nam Tha', 'Luang Prabang', 'Oudomxay', 'Phongsaly', 'Saravan', 'Savannakhet', 'Sekong', 'Viangchan Prefecture', 'Viangchan Province', 'Xaignabury', 'Xiang Khuang']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(120,'Latvia','LV',371,['Aizkraukles', 'Aluksnes', 'Balvu', 'Bauskas', 'Cesu', 'Daugavpils', 'Daugavpils City', 'Dobeles', 'Gulbenes', 'Jekabspils', 'Jelgava', 'Jelgavas', 'Jurmala City', 'Kraslavas', 'Kuldigas', 'Liepaja', 'Liepajas', 'Limbazhu', 'Ludzas', 'Madonas', 'Ogres', 'Preilu', 'Rezekne', 'Rezeknes', 'Riga', 'Rigas', 'Saldus', 'Talsu', 'Tukuma', 'Valkas', 'Valmieras', 'Ventspils', 'Ventspils City']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(121,'Lebanon','LB',961,['Beirut', 'Jabal Lubnan', 'Mohafazat Liban-Nord', 'Mohafazat Mont-Liban', 'Sidon', 'al-Biqa', 'al-Janub', 'an-Nabatiyah', 'ash-Shamal']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(122,'Lesotho','LS',266,['Berea', 'Butha-Buthe', 'Leribe', 'Mafeteng', 'Maseru', 'Mohales Hoek', 'Mokhotlong', 'Qachas Nek', 'Quthing', 'Thaba-Tseka']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(123,'Liberia','LR',231,['Bomi', 'Bong', 'Grand Bassa', 'Grand Cape Mount', 'Grand Gedeh', 'Loffa', 'Margibi', 'Maryland and Grand Kru', 'Montserrado', 'Nimba', 'Rivercess', 'Sinoe']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(124,'Libya','LY',218,['Ajdabiya', 'Fezzan', 'Banghazi', 'Darnah', 'Ghadamis', 'Gharyan', 'Misratah', 'Murzuq', 'Sabha', 'Sawfajjin', 'Surt', 'Tarabulus', 'Tarhunah', 'Tripolitania', 'Tubruq', 'Yafran', 'Zlitan', 'al-Aziziyah', 'al-Fatih', 'al-Jabal al Akhdar', 'al-Jufrah', 'al-Khums', 'al-Kufrah', 'an-Nuqat al-Khams', 'ash-Shati', 'az-Zawiyah']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(125,'Liechtenstein','LI',423,['Balzers', 'Eschen', 'Gamprin', 'Mauren', 'Planken', 'Ruggell', 'Schaan', 'Schellenberg', 'Triesen', 'Triesenberg', 'Vaduz']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(126,'Lithuania','LT',370,['Alytaus', 'Anyksciai', 'Kauno', 'Klaipedos', 'Marijampoles', 'Panevezhio', 'Panevezys', 'Shiauliu', 'Taurages', 'Telshiu', 'Telsiai', 'Utenos', 'Vilniaus']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(127,'Luxembourg','LU',352,['Capellen', 'Clervaux', 'Diekirch', 'Echternach', 'Esch-sur-Alzette', 'Grevenmacher', 'Luxembourg', 'Mersch', 'Redange', 'Remich', 'Vianden', 'Wiltz']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(128,'Macau S.A.R.','MO',853,['Macau']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(129,'Macedonia','MK',389,['Berovo', 'Bitola', 'Brod', 'Debar', 'Delchevo', 'Demir Hisar', 'Gevgelija', 'Gostivar', 'Kavadarci', 'Kichevo', 'Kochani', 'Kratovo', 'Kriva Palanka', 'Krushevo', 'Kumanovo', 'Negotino', 'Ohrid', 'Prilep', 'Probishtip', 'Radovish', 'Resen', 'Shtip', 'Skopje', 'Struga', 'Strumica', 'Sveti Nikole', 'Tetovo', 'Valandovo', 'Veles', 'Vinica']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(130,'Madagascar','MG',261,['Antananarivo', 'Antsiranana', 'Fianarantsoa', 'Mahajanga', 'Toamasina', 'Toliary']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(131,'Malawi','MW',265,['Balaka', 'Blantyre City', 'Chikwawa', 'Chiradzulu', 'Chitipa', 'Dedza', 'Dowa', 'Karonga', 'Kasungu', 'Lilongwe City', 'Machinga', 'Mangochi', 'Mchinji', 'Mulanje', 'Mwanza', 'Mzimba', 'Mzuzu City', 'Nkhata Bay', 'Nkhotakota', 'Nsanje', 'Ntcheu', 'Ntchisi', 'Phalombe', 'Rumphi', 'Salima', 'Thyolo', 'Zomba Municipality']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(132,'Malaysia','MY',60,['Johor', 'Kedah', 'Kelantan', 'Kuala Lumpur', 'Labuan', 'Melaka', 'Negeri Johor', 'Negeri Sembilan', 'Pahang', 'Penang', 'Perak', 'Perlis', 'Pulau Pinang', 'Sabah', 'Sarawak', 'Selangor', 'Sembilan', 'Terengganu']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(133,'Maldives','MV',960,['Alif Alif', 'Alif Dhaal', 'Baa', 'Dhaal', 'Faaf', 'Gaaf Alif', 'Gaaf Dhaal', 'Ghaviyani', 'Haa Alif', 'Haa Dhaal', 'Kaaf', 'Laam', 'Lhaviyani', 'Male', 'Miim', 'Nuun', 'Raa', 'Shaviyani', 'Siin', 'Thaa', 'Vaav']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(134,'Mali','ML',223,['Bamako', 'Gao', 'Kayes', 'Kidal', 'Koulikoro', 'Mopti', 'Segou', 'Sikasso', 'Tombouctou']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(135,'Malta','MT',356,['Gozo and Comino', 'Inner Harbour', 'Northern', 'Outer Harbour', 'South Eastern', 'Valletta', 'Western']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(136,'Man (Isle of)','XM',44,['Castletown', 'Douglas', 'Laxey', 'Onchan', 'Peel', 'Port Erin', 'Port Saint Mary', 'Ramsey']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(137,'Marshall Islands','MH',692,['Ailinlaplap', 'Ailuk', 'Arno', 'Aur', 'Bikini', 'Ebon', 'Enewetak', 'Jabat', 'Jaluit', 'Kili', 'Kwajalein', 'Lae', 'Lib', 'Likiep', 'Majuro', 'Maloelap', 'Mejit', 'Mili', 'Namorik', 'Namu', 'Rongelap', 'Ujae', 'Utrik', 'Wotho', 'Wotje']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(138,'Martinique','MQ',596,['Fort-de-France', 'La Trinite', 'Le Marin', 'Saint-Pierre']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(139,'Mauritania','MR',222,['Adrar', 'Assaba', 'Brakna', 'Dhakhlat Nawadibu', 'Hudh-al-Gharbi', 'Hudh-ash-Sharqi', 'Inshiri', 'Nawakshut', 'Qidimagha', 'Qurqul', 'Taqant', 'Tiris Zammur', 'Trarza']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(140,'Mauritius','MU',230,['Black River', 'Eau Coulee', 'Flacq', 'Floreal', 'Grand Port', 'Moka', 'Pamplempousses', 'Plaines Wilhelm', 'Port Louis', 'Riviere du Rempart', 'Rodrigues', 'Rose Hill', 'Savanne']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(141,'Mayotte','YT',269,['Mayotte', 'Pamanzi']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(142,'Mexico','MX',52,['Aguascalientes', 'Baja California', 'Baja California Sur', 'Campeche', 'Chiapas', 'Chihuahua', 'Coahuila', 'Colima', 'Distrito Federal', 'Durango', 'Estado de Mexico', 'Guanajuato', 'Guerrero', 'Hidalgo', 'Jalisco', 'Mexico', 'Michoacan', 'Morelos', 'Nayarit', 'Nuevo Leon', 'Oaxaca', 'Puebla', 'Queretaro', 'Quintana Roo', 'San Luis Potosi', 'Sinaloa', 'Sonora', 'Tabasco', 'Tamaulipas', 'Tlaxcala', 'Veracruz', 'Yucatan', 'Zacatecas']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(143,'Micronesia','FM',691,['Chuuk', 'Kusaie', 'Pohnpei', 'Yap']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(144,'Moldova','MD',373,['Balti', 'Cahul', 'Chisinau', 'Chisinau Oras', 'Edinet', 'Gagauzia', 'Lapusna', 'Orhei', 'Soroca', 'Taraclia', 'Tighina', 'Transnistria', 'Ungheni']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(145,'Monaco','MC',377,['Fontvieille', 'La Condamine', 'Monaco-Ville', 'Monte Carlo']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(146,'Mongolia','MN',976,['Arhangaj', 'Bajan-Olgij', 'Bajanhongor', 'Bulgan', 'Darhan-Uul', 'Dornod', 'Dornogovi', 'Dundgovi', 'Govi-Altaj', 'Govisumber', 'Hentij', 'Hovd', 'Hovsgol', 'Omnogovi', 'Orhon', 'Ovorhangaj', 'Selenge', 'Suhbaatar', 'Tov', 'Ulaanbaatar', 'Uvs', 'Zavhan']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(147,'Montserrat','MS',1664,['Montserrat']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(148,'Morocco','MA',212,['Agadir', 'Casablanca', 'Chaouia-Ouardigha', 'Doukkala-Abda', 'Fes-Boulemane', 'Gharb-Chrarda-Beni Hssen', 'Guelmim', 'Kenitra', 'Marrakech-Tensift-Al Haouz', 'Meknes-Tafilalet', 'Oriental', 'Oujda', 'Province de Tanger', 'Rabat-Sale-Zammour-Zaer', 'Sala Al Jadida', 'Settat', 'Souss Massa-Draa', 'Tadla-Azilal', 'Tangier-Tetouan', 'Taza-Al Hoceima-Taounate', 'Wilaya de Casablanca', 'Wilaya de Rabat-Sale']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(149,'Mozambique','MZ',258,['Cabo Delgado', 'Gaza', 'Inhambane', 'Manica', 'Maputo', 'Maputo Provincia', 'Nampula', 'Niassa', 'Sofala', 'Tete', 'Zambezia']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(150,'Myanmar','MM',95,['Ayeyarwady', 'Bago', 'Chin', 'Kachin', 'Kayah', 'Kayin', 'Magway', 'Mandalay', 'Mon', 'Nay Pyi Taw', 'Rakhine', 'Sagaing', 'Shan', 'Tanintharyi', 'Yangon']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(151,'Namibia','NA',264,['Caprivi', 'Erongo', 'Hardap', 'Karas', 'Kavango', 'Khomas', 'Kunene', 'Ohangwena', 'Omaheke', 'Omusati', 'Oshana', 'Oshikoto', 'Otjozondjupa']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(152,'Nauru','NR',674,['Yaren']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(153,'Nepal','NP',977,['Bagmati', 'Bheri', 'Dhawalagiri', 'Gandaki', 'Janakpur', 'Karnali', 'Koshi', 'Lumbini', 'Mahakali', 'Mechi', 'Narayani', 'Rapti', 'Sagarmatha', 'Seti']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(154,'Netherlands Antilles','AN',599,['Bonaire', 'Curacao', 'Saba', 'Sint Eustatius', 'Sint Maarten']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(155,'Netherlands The','NL',31,['Amsterdam', 'Benelux', 'Drenthe', 'Flevoland', 'Friesland', 'Gelderland', 'Groningen', 'Limburg', 'Noord-Brabant', 'Noord-Holland', 'Overijssel', 'South Holland', 'Utrecht', 'Zeeland', 'Zuid-Holland']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(156,'New Caledonia','NC',687,['Iles', 'Nord', 'Sud']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(157,'New Zealand','NZ',64,['Area Outside Region', 'Auckland', 'Bay of Plenty', 'Canterbury', 'Christchurch', 'Gisborne', 'Hawkes Bay', 'Manawatu-Wanganui', 'Marlborough', 'Nelson', 'Northland', 'Otago', 'Rodney', 'Southland', 'Taranaki', 'Tasman', 'Waikato', 'Wellington', 'West Coast']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(158,'Nicaragua','NI',505,['Atlantico Norte', 'Atlantico Sur', 'Boaco', 'Carazo', 'Chinandega', 'Chontales', 'Esteli', 'Granada', 'Jinotega', 'Leon', 'Madriz', 'Managua', 'Masaya', 'Matagalpa', 'Nueva Segovia', 'Rio San Juan', 'Rivas']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(159,'Niger','NE',227,['Agadez', 'Diffa', 'Dosso', 'Maradi', 'Niamey', 'Tahoua', 'Tillabery', 'Zinder']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(160,'Nigeria','NG',234,['Abia', 'Abuja Federal Capital Territory', 'Adamawa', 'Akwa Ibom', 'Anambra', 'Bauchi', 'Bayelsa', 'Benue', 'Borno', 'Cross River', 'Delta', 'Ebonyi', 'Edo', 'Ekiti', 'Enugu', 'Gombe', 'Imo', 'Jigawa', 'Kaduna', 'Kano', 'Katsina', 'Kebbi', 'Kogi', 'Kwara', 'Lagos', 'Nassarawa', 'Niger', 'Ogun', 'Ondo', 'Osun', 'Oyo', 'Plateau', 'Rivers', 'Sokoto', 'Taraba', 'Yobe', 'Zamfara']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(161,'Niue','NU',683,['Niue']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(162,'Norfolk Island','NF',672,['Norfolk Island']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(163,'Northern Mariana Islands','MP',1670,['Northern Islands', 'Rota', 'Saipan', 'Tinian']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(164,'Norway','NO',47,['Akershus', 'Aust Agder', 'Bergen', 'Buskerud', 'Finnmark', 'Hedmark', 'Hordaland', 'Moere og Romsdal', 'Nord Trondelag', 'Nordland', 'Oestfold', 'Oppland', 'Oslo', 'Rogaland', 'Soer Troendelag', 'Sogn og Fjordane', 'Stavern', 'Sykkylven', 'Telemark', 'Troms', 'Vest Agder', 'Vestfold', 'ÃƒÂ˜stfold']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(165,'Oman','OM',968,['Al Buraimi', 'Dhufar', 'Masqat', 'Musandam', 'Rusayl', 'Wadi Kabir', 'ad-Dakhiliyah', 'adh-Dhahirah', 'al-Batinah', 'ash-Sharqiyah']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(166,'Pakistan','PK',92,['Baluchistan', 'Federal Capital Area', 'Federally administered Tribal ', 'North-West Frontier', 'Northern Areas', 'Punjab', 'Sind']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(167,'Palau','PW',680,['Aimeliik', 'Airai', 'Angaur', 'Hatobohei', 'Kayangel', 'Koror', 'Melekeok', 'Ngaraard', 'Ngardmau', 'Ngaremlengui', 'Ngatpang', 'Ngchesar', 'Ngerchelong', 'Ngiwal', 'Peleliu', 'Sonsorol']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(168,'Palestinian Territory Occupied','PS',970,['Ariha', 'Bayt Lahm', 'Bethlehem', 'Dayr-al-Balah', 'Ghazzah', 'Ghazzah ash-Shamaliyah', 'Janin', 'Khan Yunis', 'Nabulus', 'Qalqilyah', 'Rafah', 'Ram Allah wal-Birah', 'Salfit', 'Tubas', 'Tulkarm', 'al-Khalil', 'al-Quds']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(169,'Panama','PA',507,['Bocas del Toro', 'Chiriqui', 'Cocle', 'Colon', 'Darien', 'Embera', 'Herrera', 'Kuna Yala', 'Los Santos', 'Ngobe Bugle', 'Panama', 'Veraguas']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(170,'Papua new Guinea','PG',675,['East New Britain', 'East Sepik', 'Eastern Highlands', 'Enga', 'Fly River', 'Gulf', 'Madang', 'Manus', 'Milne Bay', 'Morobe', 'National Capital District', 'New Ireland', 'North Solomons', 'Oro', 'Sandaun', 'Simbu', 'Southern Highlands', 'West New Britain', 'Western Highlands']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(171,'Paraguay','PY',595,['Alto Paraguay', 'Alto Parana', 'Amambay', 'Asuncion', 'Boqueron', 'Caaguazu', 'Caazapa', 'Canendiyu', 'Central', 'Concepcion', 'Cordillera', 'Guaira', 'Itapua', 'Misiones', 'Neembucu', 'Paraguari', 'Presidente Hayes', 'San Pedro']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(172,'Peru','PE',51,['Amazonas', 'Ancash', 'Apurimac', 'Arequipa', 'Ayacucho', 'Cajamarca', 'Cusco', 'Huancavelica', 'Huanuco', 'Ica', 'Junin', 'La Libertad', 'Lambayeque', 'Lima y Callao', 'Loreto', 'Madre de Dios', 'Moquegua', 'Pasco', 'Piura', 'Puno', 'San Martin', 'Tacna', 'Tumbes', 'Ucayali']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(173,'Philippines','PH',63,['Batangas', 'Bicol', 'Bulacan', 'Cagayan', 'Caraga', 'Central Luzon', 'Central Mindanao', 'Central Visayas', 'Cordillera', 'Davao', 'Eastern Visayas', 'Greater Metropolitan Area', 'Ilocos', 'Laguna', 'Luzon', 'Mactan', 'Metropolitan Manila Area', 'Muslim Mindanao', 'Northern Mindanao', 'Southern Mindanao', 'Southern Tagalog', 'Western Mindanao', 'Western Visayas']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(174,'Pitcairn Island','PN',0,['Pitcairn Island']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(175,'Poland','PL',48,['Biale Blota', 'Dobroszyce', 'Dolnoslaskie', 'Dziekanow Lesny', 'Hopowo', 'Kartuzy', 'Koscian', 'Krakow', 'Kujawsko-Pomorskie', 'Lodzkie', 'Lubelskie', 'Lubuskie', 'Malomice', 'Malopolskie', 'Mazowieckie', 'Mirkow', 'Opolskie', 'Ostrowiec', 'Podkarpackie', 'Podlaskie', 'Polska', 'Pomorskie', 'Poznan', 'Pruszkow', 'Rymanowska', 'Rzeszow', 'Slaskie', 'Stare Pole', 'Swietokrzyskie', 'Warminsko-Mazurskie', 'Warsaw', 'Wejherowo', 'Wielkopolskie', 'Wroclaw', 'Zachodnio-Pomorskie', 'Zukowo']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(176,'Portugal','PT',351,['Abrantes', 'Acores', 'Alentejo', 'Algarve', 'Braga', 'Centro', 'Distrito de Leiria', 'Distrito de Viana do Castelo', 'Distrito de Vila Real', 'Distrito do Porto', 'Lisboa e Vale do Tejo', 'Madeira', 'Norte', 'Paivas']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(177,'Puerto Rico','PR',1787,['Arecibo', 'Bayamon', 'Carolina', 'Florida', 'Guayama', 'Humacao', 'Mayaguez-Aguadilla', 'Ponce', 'Salinas', 'San Juan']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(178,'Qatar','QA',974,['Doha', 'Jarian-al-Batnah', 'Umm Salal', 'ad-Dawhah', 'al-Ghuwayriyah', 'al-Jumayliyah', 'al-Khawr', 'al-Wakrah', 'ar-Rayyan', 'ash-Shamal']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(179,'Reunion','RE',262,['Saint-Benoit', 'Saint-Denis', 'Saint-Paul', 'Saint-Pierre']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(180,'Romania','RO',40,['Alba', 'Arad', 'Arges', 'Bacau', 'Bihor', 'Bistrita-Nasaud', 'Botosani', 'Braila', 'Brasov', 'Bucuresti', 'Buzau', 'Calarasi', 'Caras-Severin', 'Cluj', 'Constanta', 'Covasna', 'Dambovita', 'Dolj', 'Galati', 'Giurgiu', 'Gorj', 'Harghita', 'Hunedoara', 'Ialomita', 'Iasi', 'Ilfov', 'Maramures', 'Mehedinti', 'Mures', 'Neamt', 'Olt', 'Prahova', 'Salaj', 'Satu Mare', 'Sibiu', 'Sondelor', 'Suceava', 'Teleorman', 'Timis', 'Tulcea', 'Valcea', 'Vaslui', 'Vrancea']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(181,'Russia','RU',70,['Adygeja', 'Aga', 'Alanija', 'Altaj', 'Amur', 'Arhangelsk', 'Astrahan', 'Bashkortostan', 'Belgorod', 'Brjansk', 'Burjatija', 'Chechenija', 'Cheljabinsk', 'Chita', 'Chukotka', 'Chuvashija', 'Dagestan', 'Evenkija', 'Gorno-Altaj', 'Habarovsk', 'Hakasija', 'Hanty-Mansija', 'Ingusetija', 'Irkutsk', 'Ivanovo', 'Jamalo-Nenets', 'Jaroslavl', 'Jevrej', 'Kabardino-Balkarija', 'Kaliningrad', 'Kalmykija', 'Kaluga', 'Kamchatka', 'Karachaj-Cherkessija', 'Karelija', 'Kemerovo', 'Khabarovskiy Kray', 'Kirov', 'Komi', 'Komi-Permjakija', 'Korjakija', 'Kostroma', 'Krasnodar', 'Krasnojarsk', 'Krasnoyarskiy Kray', 'Kurgan', 'Kursk', 'Leningrad', 'Lipeck', 'Magadan', 'Marij El', 'Mordovija', 'Moscow', 'Moskovskaja Oblast', 'Moskovskaya Oblast', 'Moskva', 'Murmansk', 'Nenets', 'Nizhnij Novgorod', 'Novgorod', 'Novokusnezk', 'Novosibirsk', 'Omsk', 'Orenburg', 'Orjol', 'Penza', 'Perm', 'Primorje', 'Pskov', 'Pskovskaya Oblast', 'Rjazan', 'Rostov', 'Saha', 'Sahalin', 'Samara', 'Samarskaya', 'Sankt-Peterburg', 'Saratov', 'Smolensk', 'Stavropol', 'Sverdlovsk', 'Tajmyrija', 'Tambov', 'Tatarstan', 'Tjumen', 'Tomsk', 'Tula', 'Tver', 'Tyva', 'Udmurtija', 'Uljanovsk', 'Ulyanovskaya Oblast', 'Ust-Orda', 'Vladimir', 'Volgograd', 'Vologda', 'Voronezh']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(182,'Rwanda','RW',250,['Butare', 'Byumba', 'Cyangugu', 'Gikongoro', 'Gisenyi', 'Gitarama', 'Kibungo', 'Kibuye', 'Kigali-ngali', 'Ruhengeri']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(183,'Saint Helena','SH',290,['Ascension', 'Gough Island', 'Saint Helena', 'Tristan da Cunha']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(184,'Saint Kitts And Nevis','KN',1869,['Christ Church Nichola Town', 'Saint Anne Sandy Point', 'Saint George Basseterre', 'Saint George Gingerland', 'Saint James Windward', 'Saint John Capesterre', 'Saint John Figtree', 'Saint Mary Cayon', 'Saint Paul Capesterre', 'Saint Paul Charlestown', 'Saint Peter Basseterre', 'Saint Thomas Lowland', 'Saint Thomas Middle Island', 'Trinity Palmetto Point']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(185,'Saint Lucia','LC',1758,['Anse-la-Raye', 'Canaries', 'Castries', 'Choiseul', 'Dennery', 'Gros Inlet', 'Laborie', 'Micoud', 'Soufriere', 'Vieux Fort']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(186,'Saint Pierre and Miquelon','PM',508,['Miquelon-Langlade', 'Saint-Pierre']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(187,'Saint Vincent And The Grenadines','VC',1784,['Charlotte', 'Grenadines', 'Saint Andrew', 'Saint David', 'Saint George', 'Saint Patrick']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(188,'Samoa','WS',684,['Aana', 'Aiga-i-le-Tai', 'Atua', 'Faasaleleaga', 'Gagaemauga', 'Gagaifomauga', 'Palauli', 'Satupaitea', 'Tuamasaga', 'Vaa-o-Fonoti', 'Vaisigano']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(189,'San Marino','SM',378,['Acquaviva', 'Borgo Maggiore', 'Chiesanuova', 'Domagnano', 'Faetano', 'Fiorentino', 'Montegiardino', 'San Marino', 'Serravalle']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(190,'Sao Tome and Principe','ST',239,['Agua Grande', 'Cantagalo', 'Lemba', 'Lobata', 'Me-Zochi', 'Pague']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(191,'Saudi Arabia','SA',966,['Al Khobar', 'Aseer', 'Ash Sharqiyah', 'Asir', 'Central Province', 'Eastern Province', 'Hail', 'Jawf', 'Jizan', 'Makkah', 'Najran', 'Qasim', 'Tabuk', 'Western Province', 'al-Bahah', 'al-Hudud-ash-Shamaliyah', 'al-Madinah', 'ar-Riyad']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(192,'Senegal','SN',221,['Dakar', 'Diourbel', 'Fatick', 'Kaolack', 'Kolda', 'Louga', 'Saint-Louis', 'Tambacounda', 'Thies', 'Ziguinchor']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(193,'Serbia','RS',381,['Central Serbia', 'Kosovo and Metohija', 'Vojvodina']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(194,'Seychelles','SC',248,['Anse Boileau', 'Anse Royale', 'Cascade', 'Takamaka', 'Victoria']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(195,'Sierra Leone','SL',232,['Eastern', 'Northern', 'Southern', 'Western']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(196,'Singapore','SG',65,['Singapore']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(197,'Slovakia','SK',421,['Banskobystricky', 'Bratislavsky', 'Kosicky', 'Nitriansky', 'Presovsky', 'Trenciansky', 'Trnavsky', 'Zilinsky']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(198,'Slovenia','SI',386,['Benedikt', 'Gorenjska', 'Gorishka', 'Jugovzhodna Slovenija', 'Koroshka', 'Notranjsko-krashka', 'Obalno-krashka', 'Obcina Domzale', 'Obcina Vitanje', 'Osrednjeslovenska', 'Podravska', 'Pomurska', 'Savinjska', 'Slovenian Littoral', 'Spodnjeposavska', 'Zasavska']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(199,'Smaller Territories of the UK','XG',44,['Pitcairn']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(200,'Solomon Islands','SB',677,['Central', 'Choiseul', 'Guadalcanal', 'Isabel', 'Makira and Ulawa', 'Malaita', 'Rennell and Bellona', 'Temotu', 'Western']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(201,'Somalia','SO',252,['Awdal', 'Bakol', 'Banadir', 'Bari', 'Bay', 'Galgudug', 'Gedo', 'Hiran', 'Jubbada Hose', 'Jubbadha Dexe', 'Mudug', 'Nugal', 'Sanag', 'Shabellaha Dhexe', 'Shabellaha Hose', 'Togdher', 'Woqoyi Galbed']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(202,'South Africa','ZA',27,['Eastern Cape', 'Free State', 'Gauteng', 'Kempton Park', 'Kramerville', 'KwaZulu Natal', 'Limpopo', 'Mpumalanga', 'North West', 'Northern Cape', 'Parow', 'Table View', 'Umtentweni', 'Western Cape']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(203,'South Georgia','GS',0,['South Georgia']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(204,'South Sudan','SS',211,['Central Equatoria']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(205,'Spain','ES',34,['A Coruna', 'Alacant', 'Alava', 'Albacete', 'Almeria', 'Asturias', 'Avila', 'Badajoz', 'Balears', 'Barcelona', 'Burgos', 'Caceres', 'Cadiz', 'Cantabria', 'Castello', 'Ceuta', 'Ciudad Real', 'Cordoba', 'Cuenca', 'Girona', 'Granada', 'Guadalajara', 'Guipuzcoa', 'Huelva', 'Huesca', 'Jaen', 'La Rioja', 'Las Palmas', 'Leon', 'Lleida', 'Lugo', 'Madrid', 'Malaga', 'Melilla', 'Murcia', 'Navarra', 'Ourense', 'Pais Vasco', 'Palencia', 'Pontevedra', 'Salamanca', 'Segovia', 'Sevilla', 'Soria', 'Tarragona', 'Santa Cruz de Tenerife', 'Teruel', 'Toledo', 'Valencia', 'Valladolid', 'Vizcaya', 'Zamora', 'Zaragoza']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(206,'Sri Lanka','LK',94,['Amparai', 'Anuradhapuraya', 'Badulla', 'Boralesgamuwa', 'Colombo', 'Galla', 'Gampaha', 'Hambantota', 'Kalatura', 'Kegalla', 'Kilinochchi', 'Kurunegala', 'Madakalpuwa', 'Maha Nuwara', 'Malwana', 'Mannarama', 'Matale', 'Matara', 'Monaragala', 'Mullaitivu', 'North Eastern Province', 'North Western Province', 'Nuwara Eliya', 'Polonnaruwa', 'Puttalama', 'Ratnapuraya', 'Southern Province', 'Tirikunamalaya', 'Tuscany', 'Vavuniyawa', 'Western Province', 'Yapanaya', 'kadawatha']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(207,'Sudan','SD',249,['Aali-an-Nil', 'Bahr-al-Jabal', 'Central Equatoria', 'Gharb Bahr-al-Ghazal', 'Gharb Darfur', 'Gharb Kurdufan', 'Gharb-al-Istiwaiyah', 'Janub Darfur', 'Janub Kurdufan', 'Junqali', 'Kassala', 'Nahr-an-Nil', 'Shamal Bahr-al-Ghazal', 'Shamal Darfur', 'Shamal Kurdufan', 'Sharq-al-Istiwaiyah', 'Sinnar', 'Warab', 'Wilayat al Khartum', 'al-Bahr-al-Ahmar', 'al-Buhayrat', 'al-Jazirah', 'al-Khartum', 'al-Qadarif', 'al-Wahdah', 'an-Nil-al-Abyad', 'an-Nil-al-Azraq', 'ash-Shamaliyah']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(208,'Suriname','SR',597,['Brokopondo', 'Commewijne', 'Coronie', 'Marowijne', 'Nickerie', 'Para', 'Paramaribo', 'Saramacca', 'Wanica']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(209,'Svalbard And Jan Mayen Islands','SJ',47,['Svalbard']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(210,'Swaziland','SZ',268,['Hhohho', 'Lubombo', 'Manzini', 'Shiselweni']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(211,'Sweden','SE',46,['Alvsborgs Lan', 'Angermanland', 'Blekinge', 'Bohuslan', 'Dalarna', 'Gavleborg', 'Gaza', 'Gotland', 'Halland', 'Jamtland', 'Jonkoping', 'Kalmar', 'Kristianstads', 'Kronoberg', 'Norrbotten', 'Orebro', 'Ostergotland', 'Saltsjo-Boo', 'Skane', 'Smaland', 'Sodermanland', 'Stockholm', 'Uppsala', 'Varmland', 'Vasterbotten', 'Vastergotland', 'Vasternorrland', 'Vastmanland', 'Vastra Gotaland']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(212,'Switzerland','CH',41,['Aargau', 'Appenzell Inner-Rhoden', 'Appenzell-Ausser Rhoden', 'Basel-Landschaft', 'Basel-Stadt', 'Bern', 'Canton Ticino', 'Fribourg', 'Geneve', 'Glarus', 'Graubunden', 'Heerbrugg', 'Jura', 'Kanton Aargau', 'Luzern', 'Morbio Inferiore', 'Muhen', 'Neuchatel', 'Nidwalden', 'Obwalden', 'Sankt Gallen', 'Schaffhausen', 'Schwyz', 'Solothurn', 'Thurgau', 'Ticino', 'Uri', 'Valais', 'Vaud', 'Vauffelin', 'Zug', 'Zurich']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(213,'Syria','SY',963,['Aleppo', 'Dara', 'Dayr-az-Zawr', 'Dimashq', 'Halab', 'Hamah', 'Hims', 'Idlib', 'Madinat Dimashq', 'Tartus', 'al-Hasakah', 'al-Ladhiqiyah', 'al-Qunaytirah', 'ar-Raqqah', 'as-Suwayda']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(214,'Taiwan','TW',886,['Changhua County', 'Chiayi County', 'Chiayi City', 'Taipei City', 'Hsinchu County', 'Hsinchu City', 'Hualien County', 'Kaohsiung City', 'Keelung City', 'Kinmen County', 'Miaoli County', 'Nantou County', 'Penghu County', 'Pingtung County', 'Taichung City', 'Tainan City', 'New Taipei City', 'Taitung County', 'Taoyuan City', 'Yilan County', 'YunLin County', 'Lienchiang County']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(215,'Tajikistan','TJ',992,['Dushanbe', 'Gorno-Badakhshan', 'Karotegin', 'Khatlon', 'Sughd']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(216,'Tanzania','TZ',255,['Arusha', 'Dar es Salaam', 'Dodoma', 'Iringa', 'Kagera', 'Kigoma', 'Kilimanjaro', 'Lindi', 'Mara', 'Mbeya', 'Morogoro', 'Mtwara', 'Mwanza', 'Pwani', 'Rukwa', 'Ruvuma', 'Shinyanga', 'Singida', 'Tabora', 'Tanga', 'Zanzibar and Pemba']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(217,'Thailand','TH',66,['Amnat Charoen', 'Ang Thong', 'Bangkok', 'Buri Ram', 'Chachoengsao', 'Chai Nat', 'Chaiyaphum', 'Changwat Chaiyaphum', 'Chanthaburi', 'Chiang Mai', 'Chiang Rai', 'Chon Buri', 'Chumphon', 'Kalasin', 'Kamphaeng Phet', 'Kanchanaburi', 'Khon Kaen', 'Krabi', 'Krung Thep', 'Lampang', 'Lamphun', 'Loei', 'Lop Buri', 'Mae Hong Son', 'Maha Sarakham', 'Mukdahan', 'Nakhon Nayok', 'Nakhon Pathom', 'Nakhon Phanom', 'Nakhon Ratchasima', 'Nakhon Sawan', 'Nakhon Si Thammarat', 'Nan', 'Narathiwat', 'Nong Bua Lam Phu', 'Nong Khai', 'Nonthaburi', 'Pathum Thani', 'Pattani', 'Phangnga', 'Phatthalung', 'Phayao', 'Phetchabun', 'Phetchaburi', 'Phichit', 'Phitsanulok', 'Phra Nakhon Si Ayutthaya', 'Phrae', 'Phuket', 'Prachin Buri', 'Prachuap Khiri Khan', 'Ranong', 'Ratchaburi', 'Rayong', 'Roi Et', 'Sa Kaeo', 'Sakon Nakhon', 'Samut Prakan', 'Samut Sakhon', 'Samut Songkhran', 'Saraburi', 'Satun', 'Si Sa Ket', 'Sing Buri', 'Songkhla', 'Sukhothai', 'Suphan Buri', 'Surat Thani', 'Surin', 'Tak', 'Trang', 'Trat', 'Ubon Ratchathani', 'Udon Thani', 'Uthai Thani', 'Uttaradit', 'Yala', 'Yasothon']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(218,'Togo','TG',228,['Centre', 'Kara', 'Maritime', 'Plateaux', 'Savanes']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(219,'Tokelau','TK',690,['Atafu', 'Fakaofo', 'Nukunonu']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(220,'Tonga','TO',676,['Eua', 'Haapai', 'Niuas', 'Tongatapu', 'Vavau']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(221,'Trinidad And Tobago','TT',1868,['Arima-Tunapuna-Piarco', 'Caroni', 'Chaguanas', 'Couva-Tabaquite-Talparo', 'Diego Martin', 'Glencoe', 'Penal Debe', 'Point Fortin', 'Port of Spain', 'Princes Town', 'Saint George', 'San Fernando', 'San Juan', 'Sangre Grande', 'Siparia', 'Tobago']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(222,'Tunisia','TN',216,['Aryanah', 'Bajah', 'Bin Arus', 'Binzart', 'Gouvernorat de Ariana', 'Gouvernorat de Nabeul', 'Gouvernorat de Sousse', 'Hammamet Yasmine', 'Jundubah', 'Madaniyin', 'Manubah', 'Monastir', 'Nabul', 'Qabis', 'Qafsah', 'Qibili', 'Safaqis', 'Sfax', 'Sidi Bu Zayd', 'Silyanah', 'Susah', 'Tatawin', 'Tawzar', 'Tunis', 'Zaghwan', 'al-Kaf', 'al-Mahdiyah', 'al-Munastir', 'al-Qasrayn', 'al-Qayrawan']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(223,'Turkey','TR',90,['Adana', 'Adiyaman', 'Afyon', 'Agri', 'Aksaray', 'Amasya', 'Ankara', 'Antalya', 'Ardahan', 'Artvin', 'Aydin', 'Balikesir', 'Bartin', 'Batman', 'Bayburt', 'Bilecik', 'Bingol', 'Bitlis', 'Bolu', 'Burdur', 'Bursa', 'Canakkale', 'Cankiri', 'Corum', 'Denizli', 'Diyarbakir', 'Duzce', 'Edirne', 'Elazig', 'Erzincan', 'Erzurum', 'Eskisehir', 'Gaziantep', 'Giresun', 'Gumushane', 'Hakkari', 'Hatay', 'Icel', 'Igdir', 'Isparta', 'Istanbul', 'Izmir', 'Kahramanmaras', 'Karabuk', 'Karaman', 'Kars', 'Karsiyaka', 'Kastamonu', 'Kayseri', 'Kilis', 'Kirikkale', 'Kirklareli', 'Kirsehir', 'Kocaeli', 'Konya', 'Kutahya', 'Lefkosa', 'Malatya', 'Manisa', 'Mardin', 'Mugla', 'Mus', 'Nevsehir', 'Nigde', 'Ordu', 'Osmaniye', 'Rize', 'Sakarya', 'Samsun', 'Sanliurfa', 'Siirt', 'Sinop', 'Sirnak', 'Sivas', 'Tekirdag', 'Tokat', 'Trabzon', 'Tunceli', 'Usak', 'Van', 'Yalova', 'Yozgat', 'Zonguldak']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(224,'Turkmenistan','TM',7370,['Ahal', 'Asgabat', 'Balkan', 'Dasoguz', 'Lebap', 'Mari']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(225,'Turks And Caicos Islands','TC',1649,['Grand Turk', 'South Caicos and East Caicos']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(226,'Tuvalu','TV',688,['Funafuti', 'Nanumanga', 'Nanumea', 'Niutao', 'Nui', 'Nukufetau', 'Nukulaelae', 'Vaitupu']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(227,'Uganda','UG',256,['Central', 'Eastern', 'Northern', 'Western']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(228,'Ukraine','UA',380,['Cherkaska', 'Chernihivska', 'Chernivetska', 'Crimea', 'Dnipropetrovska', 'Donetska', 'Ivano-Frankivska', 'Kharkiv', 'Kharkov', 'Khersonska', 'Khmelnytska', 'Kirovohrad', 'Krym', 'Kyyiv', 'Kyyivska', 'Lvivska', 'Luhanska', 'Mykolayivska', 'Odeska', 'Odessa', 'Poltavska', 'Rivnenska', 'Sevastopol', 'Sumska', 'Ternopilska', 'Volynska', 'Vynnytska', 'Zakarpatska', 'Zaporizhia', 'Zhytomyrska']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(229,'United Arab Emirates','AE',971,['Abu Zabi', 'Ajman', 'Dubai', 'Ras al-Khaymah', 'Sharjah', 'Sharjha', 'Umm al Qaywayn', 'al-Fujayrah', 'ash-Shariqah']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(230,'United Kingdom','GB',44,['Aberdeen', 'Aberdeenshire', 'Argyll', 'Armagh', 'Bedfordshire', 'Belfast', 'Berkshire', 'Birmingham', 'Brechin', 'Bridgnorth', 'Bristol', 'Buckinghamshire', 'Cambridge', 'Cambridgeshire', 'Channel Islands', 'Cheshire', 'Cleveland', 'Co Fermanagh', 'Conwy', 'Cornwall', 'Coventry', 'Craven Arms', 'Cumbria', 'Denbighshire', 'Derby', 'Derbyshire', 'Devon', 'Dial Code Dungannon', 'Didcot', 'Dorset', 'Dunbartonshire', 'Durham', 'East Dunbartonshire', 'East Lothian', 'East Midlands', 'East Sussex', 'East Yorkshire', 'England', 'Essex', 'Fermanagh', 'Fife', 'Flintshire', 'Fulham', 'Gainsborough', 'Glocestershire', 'Gwent', 'Hampshire', 'Hants', 'Herefordshire', 'Hertfordshire', 'Ireland', 'Isle Of Man', 'Isle of Wight', 'Kenford', 'Kent', 'Kilmarnock', 'Lanarkshire', 'Lancashire', 'Leicestershire', 'Lincolnshire', 'Llanymynech', 'London', 'Ludlow', 'Manchester', 'Mayfair', 'Merseyside', 'Mid Glamorgan', 'Middlesex', 'Mildenhall', 'Monmouthshire', 'Newton Stewart', 'Norfolk', 'North Humberside', 'North Yorkshire', 'Northamptonshire', 'Northants', 'Northern Ireland', 'Northumberland', 'Nottinghamshire', 'Oxford', 'Powys', 'Roos-shire', 'SUSSEX', 'Sark', 'Scotland', 'Scottish Borders', 'Shropshire', 'Somerset', 'South Glamorgan', 'South Wales', 'South Yorkshire', 'Southwell', 'Staffordshire', 'Strabane', 'Suffolk', 'Surrey', 'Sussex', 'Twickenham', 'Tyne and Wear', 'Tyrone', 'Utah', 'Wales', 'Warwickshire', 'West Lothian', 'West Midlands', 'West Sussex', 'West Yorkshire', 'Whissendine', 'Wiltshire', 'Wokingham', 'Worcestershire', 'Wrexham', 'Wurttemberg', 'Yorkshire']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(231,'United States','US',1,['Alabama', 'Alaska', 'Arizona', 'Arkansas', 'Byram', 'California', 'Cokato', 'Colorado', 'Connecticut', 'Delaware', 'District of Columbia', 'Florida', 'Georgia', 'Hawaii', 'Idaho', 'Illinois', 'Indiana', 'Iowa', 'Kansas', 'Kentucky', 'Louisiana', 'Lowa', 'Maine', 'Maryland', 'Massachusetts', 'Medfield', 'Michigan', 'Minnesota', 'Mississippi', 'Missouri', 'Montana', 'Nebraska', 'Nevada', 'New Hampshire', 'New Jersey', 'New Jersy', 'New Mexico', 'New York', 'North Carolina', 'North Dakota', 'Ohio', 'Oklahoma', 'Ontario', 'Oregon', 'Pennsylvania', 'Ramey', 'Rhode Island', 'South Carolina', 'South Dakota', 'Sublimity', 'Tennessee', 'Texas', 'Trimble', 'Utah', 'Vermont', 'Virginia', 'Washington', 'West Virginia', 'Wisconsin', 'Wyoming']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(232,'United States Minor Outlying Islands','UM',1,['United States Minor Outlying I']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(233,'Uruguay','UY',598,['Artigas', 'Canelones', 'Cerro Largo', 'Colonia', 'Durazno', 'FLorida', 'Flores', 'Lavalleja', 'Maldonado', 'Montevideo', 'Paysandu', 'Rio Negro', 'Rivera', 'Rocha', 'Salto', 'San Jose', 'Soriano', 'Tacuarembo', 'Treinta y Tres']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(234,'Uzbekistan','UZ',998,['Andijon', 'Buhoro', 'Buxoro Viloyati', 'Cizah', 'Fargona', 'Horazm', 'Kaskadar', 'Korakalpogiston', 'Namangan', 'Navoi', 'Samarkand', 'Sirdare', 'Surhondar', 'Toskent']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(235,'Vanuatu','VU',678,['Malampa', 'Penama', 'Sanma', 'Shefa', 'Tafea', 'Torba']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(236,'Vatican City State (Holy See)','VA',39,['Vatican City State (Holy See)']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(237,'Venezuela','VE',58,['Amazonas', 'Anzoategui', 'Apure', 'Aragua', 'Barinas', 'Bolivar', 'Carabobo', 'Cojedes', 'Delta Amacuro', 'Distrito Federal', 'Falcon', 'Guarico', 'Lara', 'Merida', 'Miranda', 'Monagas', 'Nueva Esparta', 'Portuguesa', 'Sucre', 'Tachira', 'Trujillo', 'Vargas', 'Yaracuy', 'Zulia']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(238,'Vietnam','VN',84,['Bac Giang', 'Binh Dinh', 'Binh Duong', 'Da Nang', 'Dong Bang Song Cuu Long', 'Dong Bang Song Hong', 'Dong Nai', 'Dong Nam Bo', 'Duyen Hai Mien Trung', 'Hanoi', 'Hung Yen', 'Khu Bon Cu', 'Long An', 'Mien Nui Va Trung Du', 'Thai Nguyen', 'Thanh Pho Ho Chi Minh', 'Thu Do Ha Noi', 'Tinh Can Tho', 'Tinh Da Nang', 'Tinh Gia Lai']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(239,'Virgin Islands (British)','VG',1284,['Anegada', 'Jost van Dyke', 'Tortola']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(240,'Virgin Islands (US)','VI',1340,['Saint Croix', 'Saint John', 'Saint Thomas']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(241,'Wallis And Futuna Islands','WF',681,['Alo', 'Singave', 'Wallis']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(242,'Western Sahara','EH',212,['Bu Jaydur', 'Wad-adh-Dhahab', 'al-Ayun', 'as-Samarah']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(243,'Yemen','YE',967,['Adan', 'Abyan', 'Dhamar', 'Hadramaut', 'Hajjah', 'Hudaydah', 'Ibb', 'Lahij', 'Marib', 'Madinat Sana', 'Sadah', 'Sana', 'Shabwah', 'Taizz', 'al-Bayda', 'al-Hudaydah', 'al-Jawf', 'al-Mahrah', 'al-Mahwit']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(244,'Yugoslavia','YU',38,['Central Serbia', 'Kosovo and Metohija', 'Montenegro', 'Republic of Serbia', 'Serbia', 'Vojvodina']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(245,'Zambia','ZM',260,['Central', 'Copperbelt', 'Eastern', 'Luapala', 'Lusaka', 'North-Western', 'Northern', 'Southern', 'Western']);").Exec(); err != nil {
		fmt.Println(err)
	}
	if err = cassession.Session.Query("insert into countrystate(id,countryname,countrycode,phonecode,states)Values(246,'Zimbabwe','ZW',263,['Bulawayo', 'Harare', 'Manicaland', 'Mashonaland Central', 'Mashonaland East', 'Mashonaland West', 'Masvingo', 'Matabeleland North', 'Matabeleland South', 'Midlands']);").Exec(); err != nil {
		fmt.Println(err)
	}

	// err = cassession.Session.Query("create table thumnail(thumnailuid uuid primary key,videouid uuid,thumnail blob,useruid uuid); ;").Exec()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
}
