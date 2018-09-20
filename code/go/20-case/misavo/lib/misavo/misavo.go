package misavo

import (
	"fmt"
	// "math/rand"
	"crypto/rand"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Item struct {
	Name string
	Price float64
}

type Location struct {
	Lat float64
	Lng float64
}

type Shop struct {
	Id   string `bson:"_id,omitempty"`
	User string
	Name string
	Tel string
	Country string
	Area string
	Address string
	Items []Item
	Addons []Item
	At Location
	Fulltext string
}
/*
func Seed() {
	rand.Seed(time.Now().UnixNano())
}

func Irand(min, max int) int {
	return rand.Intn(max - min) + min
}
*/
const GSIZE = 16

func Guid() string {
	b := [GSIZE]byte{} // 這種使用 "crypto/rand" 才夠亂，不容易因為 Seed(time) 而衝到。
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	/*
	b := make([]byte, GSIZE)
	for i:=0; i<GSIZE; i++ {
		b[i] = byte(Irand(0, 256))
	}
	*/
	return fmt.Sprintf("%x", b[0:GSIZE])
}

var db *mgo.Database
var session *mgo.Session

func OpenDb() (*mgo.Database, error) {
	var err error
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Timeout:  60 * time.Second,
		Database: "misavo",
		Username: "ccc",
		Password: "12345678",
	}
	// session, err := mgo.Dial("127.0.0.1:27017")
	session, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db = session.DB("misavo")
	return db, err
}

func CloseDb() {
	session.Close()
}

func QueryShop(q bson.M) (shops []Shop, err error) {
	// var shops []Shop
	err = db.C("shop").Find(q).All(&shops)
	return shops, err
}
