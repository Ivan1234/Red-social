package bd

import(
	"context"
	"time"

	"github.com/Ivan1234/Red-social/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("redsocial")
	col := db.Collection("relacion")

	skip := (pagina -1)*20

	condiciones := make([]bson.M,0)
	/*El comando match lo que hace es buscar el usuarioid de la relacion*/
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid":ID}})
	/*Lookup me permite unir 2 tablas*/
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from": "tweet",
			"localField": "usuariorelacionid",
			"foreignField": "userid",
			"as": "tweet",
		},
	})

	/*Si no usamos el unwind vendria la consulta con un registro maestro con el registro de relación
	y debajo un subregistro con un subdocumento donde estan todos los tweets de esa relación*/
	condiciones = append(condiciones, bson.M{"$unwind":"$tweet"})
	/*Con eso indicamos que entregue los documentos ordenaos por su fecha de menor a mayor de manera
	descendente, es decir que vengan primero los mas recientes y hasta el último los mas viejos*/
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	/*Si no ponemos el limit va a saltar 20 registros y va a mostrar desde ahí hasta el final de la
	base de datos*/
	condiciones = append(condiciones, bson.M{"$limit": 20})
	/*Primero se tiene que configurar el skip y luego el limit*/

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsSeguidores
	/*Se procesan todos los registros de un solo tiron, si hubiera un error lo va a almacenar en la
	variable error*/
	err = cursor.All(ctx, &result)
	if err != nil{
		return result, false
	}

	return result, true
}