package main
import(
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"JSMPJ_go_test/models"
)
func handler(){
	newRouter:=mux.NewRouter()
	origins:=handlers.AllowedOrigins([]string{"*"})
	methods:=handlers.AllowedMethods([]string{"GET","PUT","POST","DELETE"})
	newRouter.HandleFunc("/",models.HomePage).Methods("GET")
	newRouter.HandleFunc("/addmovie",models.CreateMovieDetail).Methods("POST")
	newRouter.HandleFunc("/moviedetails",models.SawDetail).Methods("GET")
	newRouter.HandleFunc("/movietitle/{title}",models.MovieByTitle).Methods("GET")
	newRouter.HandleFunc("/updategenre/{id}/{genre}",models.UpdateGenre).Methods("PUT")
	newRouter.HandleFunc("/updaterating/{id}/{rating}",models.UpdateRating).Methods("PUT")
	newRouter.HandleFunc("/searchbyid/{id}",models.SearchById).Methods("GET")
	newRouter.HandleFunc("/searchbyyears/{years}",models.SearchByYear).Methods("GET")
	newRouter.HandleFunc("/searchbyGenre/{genre}",models.SearchByGenre).Methods("GET")
	newRouter.HandleFunc("/searchbyRatings/{ratings}",models.SearchByRatings).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",handlers.CORS(origins,methods)(newRouter)))
}
func main(){
	fmt.Println("Server is Active now")
	models.InitialMigration()
	handler()

}