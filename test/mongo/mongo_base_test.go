package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestConnMongo(t *testing.T) {
	session, e := mgo.Dial("127.0.0.1:27017")
	if e != nil {
		panic(e)
	}
	session.SetMode(mgo.Monotonic, true)
	fmt.Println(session)
	t.Log("conn mongo success")

	//
	names, e := session.DB("mongo_test").CollectionNames()
	t.Log(names)
	for k, v := range names {
		fmt.Println(k, v)
	}

	//
	c := session.DB("mongo_test").C("Log")
	//log := &Log{
	//	Status:  1,
	//	Content: "测试日志",
	//	OpeUser:"kerry",
	//	RespResult:"success",
	//}
	//e = c.Insert(log)
	//if e != nil {
	//	panic(e)
	//	return
	//}
	t.Log("插入mongodb数据成功")

	// mongo 查询
	log1 := &Log{}
	c.FindId(bson.ObjectIdHex("5d4fc9b99993fcb6192dab61")).One(log1)
	t.Logf("log=%v", log1)
}

type Log struct {
	Status     int    `json:"status"`
	Content    string `json:"content"`
	OpeUser    string `json:"ope_user"`
	RespResult string `json:"resp_result"`
}
