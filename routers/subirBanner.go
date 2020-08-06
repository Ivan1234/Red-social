package routers

import(
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Ivan1234/Red-social/bd"
	"github.com/Ivan1234/Red-social/models"
)

func SubirBanner(w http.ResponseWriter, r *http.Request){

	file, handler, err := r.FormFile("banner")
	/*Hacemos un split en el punto "Nombre.jpg", haciendo de este un vector,
	sin embargo, solo vamos a guardar el elemento 1*/
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	/*Se crea un espacio en la memoria para almacenar el archvio renombrado*/
	f, err := os.OpenFile(archivo, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil{
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil{
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false{
		http.Error(w, "Error al grabar el banner en la bd! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}