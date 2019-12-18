package repositorio

import (
	"context"
	"controlAsistenciaAPI/controlasistencia"
	"controlAsistenciaAPI/modelos"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type repositorioEventosMongoDB struct {
	collection *mongo.Collection
}

func NuevoRepositorioEventosMongoDB(url string) controlasistencia.RepositorioEventos {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("eventos-db").Collection("control-asistencia")
	return &repositorioEventosMongoDB{
		collection: collection,
	}
}

type EventoMongoDB struct {
	ID            string    `bson:"id" json:"id"`
	NombreUsuario string    `bson:"nombre_usuario" json:nombre_usuario`
	Tipo          string    `bson:"tipo" json:tipo`
	Fecha         time.Time `bson:"fecha" json:fecha`
	FechaRegistro time.Time `bson:"fecha_registro" json:fecha_registro`
}

func (r repositorioEventosMongoDB) AlmacenarEvento(ctx context.Context, evento modelos.Evento) error {
	data := EventoMongoDB{
		ID:            evento.ID,
		NombreUsuario: evento.NombreUsuario,
		Tipo:          string(evento.Tipo),
		Fecha:         evento.Fecha,
		FechaRegistro: time.Now(),
	}
	_, err := r.collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
