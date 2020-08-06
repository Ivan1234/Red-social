package routers

import(
	"net/http"

	"github.com/Ivan1234/Red-social/bd"
	"github.com/Ivan1234/Red-social/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil{
		http.Error(w, "Ocurrió un error al intentar borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false{
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}