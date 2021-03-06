package bd

import(
	"context"
	"time"

	"github.com/Ivan1234/Red-social/models"
)

func InsertoRelacion(t models.Relacion) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocial")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil{
		return false, err
	}

	return true, nil
}