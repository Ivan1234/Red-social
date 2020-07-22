package bd

import(
	"context"
	"time"

	"github.com/Ivan1234/Red-social/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Recibe el ID por separado*/
func ModificoRegistro(u models.Usuario, ID string) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocial")
	col := db.Collection("usuarios")

	/*Se crea un mapa de tipo interface*/
	registro := make(map[string] interface{})
	/*Se haran una validación para que si en algun momento el usuario no modifica algo, esto no 
	sobreescriba con un vacio la base de datos*/
	if len(u.Nombre)>0{
		registro["nombre"] = u.Nombre
	}

	if len(u.Apellidos)>0{
		registro["apellidos"] = u.Apellidos
	}

	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Avatar)>0{
		registro["avatar"] = u.Avatar
	}

	if len(u.Banner)>0{
		registro["banner"] = u.Banner
	}

	if len(u.Biografia)>0{
		registro["biografia"] = u.Biografia
	}

	if len(u.Ubicacion)>0{
		registro["ubicacion"] = u.Ubicacion
	}

	if len(u.SitioWeb)>0{
		registro["sitioWeb"] = u.SitioWeb
	}

	updtString := bson.M{
		"$set":registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	/*El $eq significa que los documentos que queremos actualizar tienen que ser igual al ID 
	proporcionado*/
	filtro := bson.M{"_id":bson.M{"$eq":objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil{
		return false, err
	}

	return true, nil
}