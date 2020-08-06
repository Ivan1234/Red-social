package bd

import(
	"context"
	"fmt"
	"time"

	"github.com/Ivan1234/Red-social/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocial")
	col := db.Collection("usuarios")

	/*Se va a envíar al http un slice de usuarios*/
	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page -1)*20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex":`(?i)`+search},
	}

	/*Cuando la sentencia no es un findOne, lo que retorna es un cursor*/
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil{
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx){
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil{
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false{
			incluir = true
		}

		if tipo == "follow" && encontrado == true{
			incluir = true
		}

		if r.UsuarioRelacionID == ID{
			incluir = false
		}

		if incluir == true{
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil{
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}