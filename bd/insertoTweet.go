package bd

import(
	"context"
	"time"

	"github.com/Ivan1234/Red-social/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoTweet(t models.GraboTweet) (string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocial")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid" : t.UserID,
		"mensaje":t.Mensaje,
		"fecha": t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil{
		return string(""), false, err 
	}

	/*Extrae el id del último campo insertado*/
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}