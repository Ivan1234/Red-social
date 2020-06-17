package main

import(	
	"log"
	"github.com/Ivan1234/Red-social/handlers"
	"github.com/Ivan1234/Red-social/bd"
)

func main(){
	if bd.ChequeoConnection()==0{
		log.Fatal("Sin conexión a la base de datos")
		return
	}

	handlers.Manejadores()
}