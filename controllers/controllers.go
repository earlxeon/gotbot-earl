package controllers

import (
	context "context"
	"encoding/json"
	config "gotbotpoc/config"

	userDetails "gotbotpoc/models"

	"github.com/gofiber/fiber"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	handlers "gotbotpoc/handlers"

	"gotbotpoc/db"
)

func GetUserList(c *fiber.Ctx) {

	TokenValueFromHeader := c.Get("Token")

	flag := handlers.VerifyToken(TokenValueFromHeader)

	if flag {
		collection, err := db.GetMongoDbCollection(config.Config("DB_NAME"), config.Config("COLLECTION_NAME"))
		if err != nil {
			c.Status(500).Send(err)
			return
		}

		var filter bson.M = bson.M{}

		if c.Params("id") != "" {
			id := c.Params("id")
			objID, _ := primitive.ObjectIDFromHex(id)
			filter = bson.M{"_id": objID}
		}

		var results []bson.M
		cur, err := collection.Find(context.Background(), filter)

		if err != nil {
			defer cur.Close(context.Background())
			c.Status(500).Send(err)
			return
		}

		cur.All(context.Background(), &results)

		if results == nil {
			c.SendStatus(404)
			return
		}

		json, _ := json.Marshal(results)
		c.Send(json)
	} else {

		in := []byte(`{"ErrorMessage":"Invalid or Expired Token"}`)

		var iot userDetails.ErrorMessage
		err := json.Unmarshal(in, &iot)
		if err != nil {
			panic(err)
		}

		// Marshal back to json (as original)
		json, _ := json.Marshal(&iot)
		c.Status(401).Send([]byte(string(json)))
	}
}

func AddUser(c *fiber.Ctx) {
	collection, err := db.GetMongoDbCollection(config.Config("DB_NAME"), config.Config("COLLECTION_NAME"))
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	user := userDetails.User{}
	json.Unmarshal([]byte(c.Body()), &user)

	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)

}

func DeleteUser(c *fiber.Ctx) {
	collection, err := db.GetMongoDbCollection(config.Config("DB_NAME"), config.Config("COLLECTION_NAME"))

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	jsonResponse, _ := json.Marshal(res)
	c.Send(jsonResponse)
}

func Login(c *fiber.Ctx) {
	collection, err := db.GetMongoDbCollection(config.Config("DB_NAME"), config.Config("COLLECTION_NAME"))
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	user := userDetails.LoginDetails{}
	userData := userDetails.User{}

	json.Unmarshal([]byte(c.Body()), &user)

	err = collection.FindOne(context.Background(), user).Decode(&userData)

	tokenincoming := handlers.GenerateNewAccessToken(user.Email, userData.Roleid)
	userData.Token = tokenincoming

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	json, _ := json.Marshal(userData)
	c.Send(json)
}

// // // func FindOneAndUpdate() {

// // // 	collection, err := db.GetMongoDbCollection(config.Config("DB_NAME"), config.Config("COLLECTION_NAME"))

// // // 	if err != nil {
// // // 		//c.Status(500).Send(err)
// // // 		//return
// // // 	}

// // // 	// 5) Create the search filter
// // // 	//take the incoming email and place it here
// // // 	filter := bson.M{"email": "chetanm@winjit.com1"}

// // // 	// 6) Create the update
// // // 	update := bson.M{
// // // 		"$set": bson.M{"age": 5234234, "title": "ajax developer"},
// // // 	}

// // // 	// 7) Create an instance of an options and set the desired options
// // // 	upsert := true
// // // 	after := options.After
// // // 	opt := options.FindOneAndUpdateOptions{
// // // 		ReturnDocument: &after,
// // // 		Upsert:         &upsert,
// // // 	}

// // // 	// 8) Find one result and update it
// // // 	result := collection.FindOneAndUpdate(context.Background(), filter, update, &opt)
// // // 	if result.Err() != nil {
// // // 		//return nil, result.Err()
// // // 	}

// // // 	// 9) Decode the result

// // // }

func UpdateUser(c *fiber.Ctx) {

	collection, err := db.GetMongoDbCollection(config.Config("DB_NAME"), config.Config("COLLECTION_NAME"))

	if err != nil {
		//fmt.Print(errors)
	}

	//to get chetans email in a string
	//user := userDetails.LoginDetails{}
	userData := userDetails.User{}
	json.Unmarshal([]byte(c.Body()), &userData)
	//jamesA := user.Email
	//fmt.Println("jamesA =" + jamesA)
	//err = collection.FindOne(context.Background(), user).Decode(&userData)
	userNameLocal := userData.Name
	userEmailLocal := userData.Email
	userPaswordLocal := userData.Password
	userTitleLocal := userData.Title
	userBirthdateLocal := userData.Birthdate
	//fmt.Println("jamesB =" + userNameLocal)

	//jamesC := userData.Email
	//fmt.Println("jamesC =" + jamesC)
	//jamesD := userData.Email
	//fmt.Println("jamesD =" + jamesD)

	// if err != nil {
	// 	c.Status(500).Send(err)
	// 	return
	// }

	// 5) Create the search filter
	//take the incoming email and place it here
	filter := bson.M{"email": " " + userEmailLocal + ""}

	// 6) Create the update
	update := bson.M{
		"$set": bson.M{
			"name":      userNameLocal,
			"password":  userPaswordLocal,
			"birthdate": userBirthdateLocal,
			"title":     userTitleLocal},
	}

	// 7) Create an instance of an options and set the desired options
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	// 8) Find one result and update it
	result := collection.FindOneAndUpdate(context.Background(), filter, update, &opt)
	if result.Err() != nil {
		//return nil, result.Err()
	}

	// 9) Decode the result

}
