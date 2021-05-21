package controllers

import (
	context "context"
	"encoding/json"
	"fmt"
	config "gotbotpoc/config"

	userDetails "gotbotpoc/models"

	"github.com/gofiber/fiber"
	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"

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
		fmt.Print("Token Invalid")
	}
}

func AddUpdateUser(c *fiber.Ctx) {
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
