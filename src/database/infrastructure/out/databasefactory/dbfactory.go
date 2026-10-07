package databasefactory

import (
	"context"
	"fmt"
	"log"
	"sync"

	"lexdev_api.com/paws/src/database/domain"
	"lexdev_api.com/paws/src/database/infrastructure/config"
	mg "lexdev_api.com/paws/src/database/infrastructure/out/databasefactory/mongo"
)

var (
	instance domain.DriverConnectionBuilder
	once     sync.Once
)

func ConnecToDatabase() domain.DriverConnectionBuilder {
	builderType := config.NewDatabaseCredentials()

	if builderType == nil {
		return nil
	}

	switch builderType.Driver {
	case "MONGO":
		conn := mg.NewMongoConnection()

		err := conn.Connect(context.Background())
		if err != nil {
			log.Fatalf("Error al conectar a %s: %v", builderType.Driver, err)
		}

		fmt.Printf("Error al conectar a %s: %v", builderType.Driver, err)

		return conn
	default:
		return nil
	}
}

func GetInstance() domain.DriverConnectionBuilder {
	once.Do(func() {
		instance = ConnecToDatabase()
	})

	return instance
}
