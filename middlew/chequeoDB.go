package middlew

import(
	"net/http"

	"github.com/Ivan1234/Red-social/bd"
)

/*ChequeoBD es el middleware que me permite conocer el estado de la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if bd.ChequeoConnection()==0{
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		
		next.ServeHTTP(w, r)
	}
}