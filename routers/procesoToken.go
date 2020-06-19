package routers

import(
	"errors"
	//"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/Ivan1234/Red-social/bd"
	"github.com/Ivan1234/Red-social/models"
)
/*Email valor de Email usado en todos los endpoints*/
var Email string
/*IDUsuario es el Id devuelto del modelo , que se usará en todos los endpoints*/
var IDUsuario string

/*ProcesoToken() proceso token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error){
	miClave := []byte("MiClaveSuperSecreta")
	claims := &models.Claim{}

	/*splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2{
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])*/

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error){
		return miClave, nil
	})
	if err == nil{
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true{
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims,encontrado, IDUsuario, nil
	}

	if !tkn.Valid{
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}