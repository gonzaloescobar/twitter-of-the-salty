package db

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:Testing!@twitter-cluster.onspz.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConnectDB es la funcion que me permite conectar a la base de datos*/
func ConnectDB() *mongo.Client {
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err.Error())
        return client
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err.Error())
        return client
    }
    log.Println("DB Connection Success")
    return client
}

/*CheckConnection revisa la conexion a la base de datos*/
func CheckConnection() int {
    err := MongoCN.Ping(context.TODO(), nil)
    if err != nil {
        return 0
    }
    return 1
}