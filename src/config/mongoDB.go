package config

import (
	"fmt"
	"os"
  "gopkg.in/mgo.v2"
)

func ConnectMongoDB() *mgo.Session {
	uri := "127.0.0.10:28017"
	if uri == "" {
		fmt.Printf("No URL found")
		os.Exit(1)
	}
	con, err := mgo.Dial(uri)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	con.SetMode(mgo.Monotonic, true)

	return con.Clone()
}
