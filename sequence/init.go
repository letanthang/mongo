package sequence

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Counter struct {
	ID  int `json:"id" bson:"id"`
	Seq int `json:"seq" bson:"seq"`
}

func GetNextID(collection *mongo.Collection, sequenceName string) (int, error) {
	// result := bson.M{}
	// if _, err := c.Find(bson.M{idFieldName: name}).Apply(mgo.Change{
	// 	Update:    bson.M{"$set": bson.M{idFieldName: name}, "$inc": bson.M{seqFieldName: 1}},
	// 	Upsert:    true,
	// 	ReturnNew: true,
	// }, &result); err != nil {
	// 	fmt.Println("Autoincrement error(1):", err.Error())
	// }
	// sec, _ := result[seqFieldName].(int)
	findOptions := options.FindOneAndUpdate()
	findOptions.SetUpsert(true)
	var counter Counter
	err := collection.FindOneAndUpdate(context.TODO(),
		bson.D{{"id", sequenceName}},
		bson.D{{"$inc", bson.D{{"seq", 1}}}},
		findOptions).Decode(&counter)

	if err != nil {
		return 0, err
	}
	return counter.Seq, nil
}
