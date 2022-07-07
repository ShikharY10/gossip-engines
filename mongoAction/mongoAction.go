package mongoAction

import (
	"context"
	"errors"
	"fmt"
	"gbEngine/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Ctx            context.Context
	Client         *mongo.Client
	UserCollection *mongo.Collection
	MsgCollection  *mongo.Collection
}

func (m *Mongo) Init(mongoIP string, username string, password string) {
	var cred options.Credential
	cred.Username = username
	cred.Password = password

	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoIP + ":27017").SetAuth(cred)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	uCollection := client.Database("Users").Collection("UserDatas")
	mCollection := client.Database("messages").Collection("userMsg")
	m.Ctx = ctx
	m.Client = client

	m.UserCollection = uCollection
	m.MsgCollection = mCollection
	fmt.Println("Mongo client connected!")
}

func (m *Mongo) InsertMsg(Mid string, mLoc string, msg []byte) error {
	_id, _ := primitive.ObjectIDFromHex(Mid)
	_, err := m.MsgCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": _id},
		bson.M{"$set": bson.M{"msg." + mLoc: msg}},
	)
	if err != nil {
		log.Println("[MongoUpdateError] : ", err.Error())
		return err
	}
	return nil
}

func (m *Mongo) GetMsg(Mid string, MsgLoc string) ([]byte, error) {
	allData, err := m.GetAllMsg(Mid)
	if err != nil || len(*allData) == 0 {
		return nil, errors.New("no data found")
	}
	qw := *allData
	q := qw[MsgLoc]
	return q, nil
}

func (m *Mongo) UpdateMsgStatus(_id string, status int) (int, error) {
	id, _ := primitive.ObjectIDFromHex(_id)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"sts": status}}
	result, err := m.UserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return -1, err
	}
	i := result.MatchedCount
	return int(i), nil

}

func (m *Mongo) DeleteMsg(Mid string, MsgLoc string) {
	if len(MsgLoc) < 1 || len(Mid) < 1 {
		fmt.Println("[ACK-ERROR] : Bad Acknowledgment Details!")
		return
	}
	_id, _ := primitive.ObjectIDFromHex(Mid)
	r, err := m.MsgCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": _id},
		bson.M{"$unset": bson.M{"msg." + MsgLoc: 1}},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.MatchedCount)
}

func (m *Mongo) GetUserDataByMID(target string) (*utils.UserData, error) {
	cursor, err := m.UserCollection.Find(context.TODO(), bson.M{"msgid": target})
	if err != nil {
		return nil, err
	}
	var ud []utils.UserData
	cursor.All(context.TODO(), &ud)
	return &ud[0], err
}

// func (m *Mongo) GetUserMsgId(target string) string {
// 	uD, err := m.GetUserData(target)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return ""
// 	}
// 	return uD.MsgId
// }

func (m *Mongo) GetAllMsg(TMid string) (*map[string][]byte, error) {
	_id, _ := primitive.ObjectIDFromHex(TMid)
	cursor, err := m.MsgCollection.Find(
		context.TODO(),
		bson.M{"_id": _id},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	var elem utils.Df
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&elem)
		if err != nil {
			fmt.Println("err: ", err.Error())
		}
	}
	if len(elem.Msg) == 0 {
		return nil, errors.New("no data found")
	}
	return &elem.Msg, nil
}

func (m *Mongo) GetMainKey(id string) string {
	cursor, err := m.UserCollection.Find(
		context.TODO(),
		bson.M{"msgid": id},
	)
	if err != nil {
		log.Println("[MONGOGETERROR] : ", err.Error())
	}
	var userd []utils.UserData
	err = cursor.All(context.TODO(), &userd)
	if err != nil {
		log.Println("[MONGOCURSORERROR] : ", err.Error())
	}
	return userd[0].MainKey
}

func (m *Mongo) GetMsgIdByNum(mNum string) string {
	cursor, err := m.UserCollection.Find(
		context.TODO(),
		bson.M{"phone_no": mNum},
	)
	if err != nil {
		log.Println("[MONGOGETERROR] : ", err.Error())
	}
	var userd []utils.UserData
	err = cursor.All(context.TODO(), &userd)
	if err != nil {
		log.Println("[MONGOCURSORERROR] : ", err.Error())
	}
	return userd[0].MsgId
}

func (m *Mongo) InsertIntoConnection(userMID string, connMID string, mainKey string) (int, error) {
	r, err := m.UserCollection.UpdateOne(
		context.TODO(),
		bson.M{"msgid": userMID},
		bson.M{"$set": bson.M{"connections." + connMID: 1}},
	)
	if err != nil {
		return 0, err
	}
	return int(r.MatchedCount), nil
}
