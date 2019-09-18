# Пример подключения к mongoDB

Запуск mongoDB
    
    sudo service mongod start
    mongo

Запуск приложения из его директории
 
    go build && ./test_1

Добавление адреса

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

Подключение к БД

    client, err := mongo.Connect(context.TODO(), clientOptions)

Проверка подключения

	err = client.Ping(context.TODO(), nil)

Обращение к коллекции

	collection := client.Database("test").Collection("trainers")

Вставка объекта в коллекцию

    insertResult, err := collection.InsertOne(context.TODO(), ilya)

Обновление полей объекта

    filter := bson.D{{"name", "Ilya"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 5},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	
Удаление объектов из коллекции

	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})

Вывод на консоль при успешной компиляции
    
    MongoDB EXAMPLE START
    Connected to MongoDB!
    Inserted a single document:  ObjectID("5d7cd203b352dd6d6ac75f69")
    Inserted multiple documents:  [ObjectID("5d7cd203b352dd6d6ac75f6a") ObjectID("5d7cd203b352dd6d6ac75f6b")]
    Matched 1 documents and updated 1 documents.
    Found a single document: {Name:Ilya Age:6 City:Bugulma}
    Found multiple documents (array of pointers): [0xc000142870 0xc0001428a0]
    Deleted 3 documents in the trainers collection
    Connection to MongoDB closed.

