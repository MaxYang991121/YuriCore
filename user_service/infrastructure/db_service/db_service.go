package db_service

import (
	"context"
	"fmt"

	"github.com/KouKouChan/YuriCore/main_service/model/user"
	. "github.com/KouKouChan/YuriCore/verbose"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBImpl struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewDBImpl(username, password, ip, port string) *DBImpl {
	url := ""
	if username != "" &&
		password != "" {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%s/", username, password, ip, port)
	} else {
		url = fmt.Sprintf("mongodb://%s:%s", ip, port)
	}
	fmt.Println("mongo uri=", url)
	clientOptions := options.
		Client().
		ApplyURI(url)
	client, err := mongo.Connect(
		context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	if client == nil {
		panic("db client == null!")
	}

	database := client.Database("yuricore")
	if database == nil {
		panic("db database == null!")
	}

	return &DBImpl{
		client:   client,
		database: database,
	}
}

func (d *DBImpl) GetUser(ctx context.Context, username string) (*user.UserInfo, error) {
	user := &user.UserInfo{}
	collection := d.database.Collection("user_table")
	err := collection.FindOne(
		ctx,
		bson.M{
			"username": username,
		},
	).Decode(user)
	if err != nil {
		return nil, err
	}

	DebugPrintf(2, "db got user %+v", *user)
	return user, nil
}

func (d *DBImpl) GetUserByNickName(ctx context.Context, nickname string) (*user.UserInfo, error) {
	user := &user.UserInfo{}
	collection := d.database.Collection("user_table")
	err := collection.FindOne(
		ctx,
		bson.M{
			"nickname": nickname,
		},
	).Decode(user)
	if err != nil {
		return nil, err
	}

	DebugPrintf(2, "db got user %+v", *user)
	return user, nil
}

func (d *DBImpl) UpdateUser(ctx context.Context, user *user.UserInfo) error {
	collection := d.database.Collection("user_table")
	filter := bson.M{"username": user.UserName}
	update := bson.M{
		"$set": user,
	}
	options := new(options.UpdateOptions)
	options.SetUpsert(true)
	_, err := collection.UpdateOne(
		ctx,
		filter,
		update,
		options,
	)
	DebugPrintf(2, "db update user %+v", *user)
	return err
}
