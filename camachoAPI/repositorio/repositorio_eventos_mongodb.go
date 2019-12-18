package repositorio

import (
	"camachoAPI/consultaEventos"
	"camachoAPI/modelos"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type repositorioEventosMongoDB struct {
	collection *mongo.Collection
}

func NuevoRepositorioEventosMongoDB(url string) consultaEventos.RepositorioEventos {
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

func (r repositorioEventosMongoDB) ListaDeEventos(ctx context.Context, usuarioID string) ([]modelos.Evento, error) {
	filter := bson.D{{"nombre_usuario", usuarioID}}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"fecha", -1}})
	result, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	var events = make([]modelos.Evento, 0)
	for result.Next(context.TODO()) {
		eventDB := EventoMongoDB{}
		err = result.Decode(eventDB)
		if err != nil {
			return nil, err
		}
		newEvento := modelos.Evento{
			ID:            eventDB.ID,
			NombreUsuario: eventDB.NombreUsuario,
			Fecha:         eventDB.Fecha,
			Tipo:          modelos.TipoEvento(eventDB.Tipo),
		}
		events = append(events, newEvento)
	}
	return events, nil
}

func (r repositorioEventosMongoDB) EstadoUsuario(ctx context.Context, usuarioID string) (modelos.EstadoAsistencia, error) {
	filter := bson.D{{"nombre_usuario", usuarioID}}
	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{"fecha", -1}})
	result := r.collection.FindOne(ctx, filter, findOptions)
	if result == nil {
		return modelos.EstadoAsistenciaNoRegistrado, nil
	}
	evento := EventoMongoDB{}
	err := result.Decode(evento)
	if err != nil {
		return modelos.EstadoAsistenciaNoRegistrado, err
	}
	switch evento.Tipo {
	case string(modelos.EntradaTipoEvento):
		return modelos.EstadoAsistenciaTrabajando, nil
	case string(modelos.SalidaTipoEvento):
		return modelos.EstadoAsistenciaDescanso, nil
	}
	return modelos.EstadoAsistenciaNoRegistrado, nil
}
