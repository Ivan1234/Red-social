package routers

import(
	"encoding/json"
	"net/http"

	"github.com/Ivan1234/Red-social/bd"
	"github.com/Ivan1234/Red-social/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request){
	var t models.Usuario

	/*Recibimos a traves del body un json y eso hay que decodificarlo*/
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil{
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	
	var status bool
	/*Se usa la variable global "IDUsuario" de procesoToken*/
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil{
		http.Error(w, "Ocurrió un error al intentar modificar el registro intente nuevamente "+err.Error(),400)
		return
	}

	if status == false{
		http.Error(w, "No se ha logrado modificar el registro de usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}