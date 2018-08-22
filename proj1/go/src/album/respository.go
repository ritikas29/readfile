package album

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type Repository struct{}
const SERVER ="localhost:27017"
const DBNAME ="musicstrore"
const DOCNAME = "albums"
func(r Repository) GetAlbums() Albums {
	session,err := mgo.Dial(SERVER)
	if err !=nil {
		fmt.Println("failed to establish",err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Albums{}
	if err :=c.Find(nil).All(&results); err !=nil {
		fmt.Println("failed to use",err)
	}
	return results
}
func (r Repository) AddAlbum(album Album) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
   
   album.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(album)
   
   if err != nil {
	 log.Fatal(err)
	 return false
	}
	return true
   }