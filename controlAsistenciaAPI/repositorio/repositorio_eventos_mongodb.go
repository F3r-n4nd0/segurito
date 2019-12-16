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

func NuevoRepositorioEventosMongoDB() controlasistencia.RepositorioEventos {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("chikitania-db").Collection("kardex")
	return &repositorioEventosMongoDB{
		collection: collection,
	}
}

func (r repositorioEventosMongoDB) AlmacenarEntrada(ctx context.Context, usuario modelos.Usuario, date time.Time) error {
	return nil
}

func (r repositorioEventosMongoDB) AlmacenarSalida(ctx context.Context, usuario modelos.Usuario, date time.Time) error {
	return nil
}
