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
	router.HandleFunc("/modificarPerfil",middlew.ChequeoBD(middlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet",middlew.ChequeoBD(middlew.ValidarJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet",middlew.ChequeoBD(middlew.ValidarJWT(routers.LeoTweets))).Methods("GET")
	
	router.HandleFunc("/eliminarTweet",middlew.ChequeoBD(middlew.ValidarJWT(routers.EliminarTweet))).Methods("delete")
	router.HandleFunc("/subirAvatar",middlew.ChequeoBD(middlew.ValidarJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar",middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner",middlew.ChequeoBD(middlew.ValidarJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner",middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")

	router.HandleFunc("/altaRelacion",middlew.ChequeoBD(middlew.ValidarJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion",middlew.ChequeoBD(middlew.ValidarJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion",middlew.ChequeoBD(middlew.ValidarJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios",middlew.ChequeoBD(middlew.ValidarJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores",middlew.ChequeoBD(middlew.ValidarJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT") //Verifica si en el sistema operativo hay un puerto
	if PORT == ""{
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}