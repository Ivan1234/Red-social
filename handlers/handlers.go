package handlers

import(
	"log"
	"net/http"
	"os"

	"github.com/Ivan1234/Red-social/middlew"
	"github.com/Ivan1234/Red-social/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors" //Permite que se acceda a la API de manera remota
)

/*Manejadores() configura el puerto, el handler y pongo a escuchar al servidor*/
func Manejadores(){
	router := mux.NewRouter()

	router.HandleFunc("/registro",middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login",middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil",middlew.ChequeoBD(middlew.ValidarJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT") //Verifica si en el sistema operativo hay un puerto
	if PORT == ""{
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}