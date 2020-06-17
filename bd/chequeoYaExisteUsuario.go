package bd

import(
	"context"
	"time"

	"github.com/Ivan1234/Red-social/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario() recibe un email de parámetro y chequea si ya esta en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocial")
	col := db.Collection("usuarios")

	condicion := bson.M{"mail":email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil{
		return resultado, false, ID
	}

	return resultado, true, ID
}