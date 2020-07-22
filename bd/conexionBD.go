package bd

import(
	"context"
	"log"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BD*/
var MongoCN = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://Sormar1:Hackerdross1@cluster0-wxssh.mongodb.net/Cluster0?retryWrites=true&w=majority")

/*ConectarBD() es la función que me permite conectar la BD*/
func ConectarBD() *mongo.Client{ //Función que devuelve un objeto mongo client (devuelve la conexión)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}

	fmt.Println("La conexión con la Base de Datos fue exitosa")
	return client
}

/*ChequeoConnection() es la funcion que hace un ping a la BD*/
func ChequeoConnection() int{
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil{
		return 0
	}
	return 1
}
