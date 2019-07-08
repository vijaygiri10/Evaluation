package models
import(
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gorilla/mux"
)
type Movie struct{
	ID string `json:"ID"`
	Title string `json:"Title"`
	ReleasedYear string `json:"ReleasedYear"`
	Genre string `json:"Genre"`
	Ratings string `json:"Ratings"`
}
var db *gorm.DB
var err error
func InitialMigration(){
	db, err = gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("problem to create database")
	}
	defer db.Close()
	db.AutoMigrate(&Movie{})
	fmt.Println("database of movies created successfully")
}
func CreateMovieDetail(w http.ResponseWriter, r *http.Request){
	body,err:=ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("there is error")
		fmt.Fprintf(w,"there is problem to read from body")
	}
	var movies Movie
	err=json.Unmarshal(body,&movies)
	if err != nil{
		fmt.Println("we can not read from body")
		fmt.Fprintf(w,"there is problem to read from body")
	}
	db, err = gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("problem to create database")
	}
	defer db.Close()
	db.Create(&Movie{ID:movies.ID,Title:movies.Title,ReleasedYear:movies.ReleasedYear,Genre:movies.Genre,Ratings:movies.Ratings})
	fmt.Fprintf(w,"New movie detail is posted successfully")
}
func SawDetail(w http.ResponseWriter, r *http.Request){
	db,err= gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("There is problem to read data")
	}
	defer db.Close()
	var Moviee []Movie
	db.Find(&Moviee)
	if len(Moviee) == 0{
		fmt.Fprintf(w,"We can not read from database")
	}else{
       json.NewEncoder(w).Encode(Moviee)
	}
} 
func MovieByTitle(w http.ResponseWriter,r *http.Request){
	db,err= gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("There is problem to read data")
	}
	defer db.Close()
	var moviess []Movie
	vars:=mux.Vars(r)
	title:=vars["title"]
	db.Where("Title = ?",title).Find(&moviess)
	if len(moviess) == 0{
		fmt.Fprintf(w,"Movie is not present in the list")
	}else{
       json.NewEncoder(w).Encode(moviess)
	}
}

func UpdateGenre(w http.ResponseWriter, r *http.Request){
	db,err= gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("There is problem to read data")
	}
	defer db.Close()
	vars:=mux.Vars(r)
	id:=vars["id"]
	genre:=vars["genre"]
	var moviess Movie
	db.Where("ID = ?",id).Find(&moviess)
	moviess.Genre=genre
	db.Save(&moviess)
	fmt.Fprintf(w,"Genre of Movie updated Successfully")
}

func UpdateRating(w http.ResponseWriter, r *http.Request){
	db,err= gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("There is problem to read data")
	}
	defer db.Close()
	vars:=mux.Vars(r)
	id:=vars["id"]
	rating:=vars["rating"]
	var moviess Movie
	db.Where("ID = ?",id).Find(&moviess)
	moviess.Ratings=rating
	db.Save(&moviess)
	fmt.Fprintf(w,"Rating of Movie updated Successfully")
}

func SearchById(w http.ResponseWriter, r *http.Request){
	db,err= gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("There is problem to read data")
	}
	defer db.Close()
	var moviess []Movie
	vars:=mux.Vars(r)
	id:=vars["id"]
	db.Where("ID = ?",id).Find(&moviess)
	if len(moviess) == 0{
		fmt.Fprintf(w,"Movie is not present in the list")
	}else{
       json.NewEncoder(w).Encode(moviess)
	}
}
func SearchByYear(w http.ResponseWriter, r *http.Request){
	db,err= gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("There is problem to read data")
	}
	defer db.Close()
	var moviess []Movie
	vars:=mux.Vars(r)
	years:=vars["years"]
	db.Where("Released_Year = ?",years).Find(&moviess)
	if len(moviess) == 0{
		fmt.Fprintf(w,"Movie is not present in the list")
	}else{
       json.NewEncoder(w).Encode(moviess)
	}
}
func SearchByGenre(w http.ResponseWriter, r *http.Request){
	db,err= gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("There is problem to read data")
	}
	defer db.Close()
	var moviess []Movie
	vars:=mux.Vars(r)
	genre:=vars["genre"]
	db.Where("Genre = ?",genre).Find(&moviess)
	if len(moviess) == 0{
		fmt.Fprintf(w,"Movie is not present in the list")
	}else{
       json.NewEncoder(w).Encode(moviess)
	}
}
func SearchByRatings(w http.ResponseWriter, r *http.Request){
	db,err= gorm.Open("sqlite3","movies.db")
	if err != nil{
		panic("There is problem to read data")
	}
	defer db.Close()
	var moviess []Movie
	vars:=mux.Vars(r)
	ratings:=vars["ratings"]
	db.Where("Ratings = ?",ratings).Find(&moviess)
	if len(moviess) == 0{
		fmt.Fprintf(w,"Movie is not present in the list")
	}else{
       json.NewEncoder(w).Encode(moviess)
	}
}
func HomePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"This is Home page body")
}