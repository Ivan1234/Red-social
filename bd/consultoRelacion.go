package bd

import(
	"context"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/Ivan1234/Red-social/models"
)

func ConsultoRelacion(t models.Relacion) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocial")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid": t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil{
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}