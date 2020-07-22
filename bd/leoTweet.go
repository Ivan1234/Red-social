package bd

import(
	"context"
	"time"
	"log"

	"github.com/Ivan1234/Red-social/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*"pagina" es para paginar, así mismo esta rutina devuelve un slice con todos los tweets del usuario*/
func LeoTweets(ID string, pagina int64)([]*models.DevuelvoTweets, bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("redsocial")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	/*El paquete opciones nos va a dar opciones de filtro para poder darle un comportamiento a 
	nuestra base de datos*/
	opciones := options.Find()
	opciones.SetLimit(20) //Cuantos documentos me tiene que traer límite
	opciones.SetSort(bson.D{{Key:"fecha", Value: -1}}) //Que traiga los ultimos tweets primero
	opciones.SetSkip((pagina -1)*20)//Que vaya salteando los tweets de 20 en 20

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil{
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()){
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)	
		if err != nil{
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}

	return resultados, true
}