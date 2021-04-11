package db

import (
	"context"
	"testing"
)

func TestDb(t *testing.T) {
	var conf Config
	t.Run("InitDb", func(t *testing.T) {
		conf = Config{
			Host:         "http://127.0.0.1:8529",
			Username:     "root",
			Password:     "abc123",
			DatabaseName: "test",
		}
		_, err := conf.Initdb()
		if err != nil {
			t.Fatal(err.Error())
		}
	})
	t.Run("Collection", func(t *testing.T) {
		col, err := conf.Collection("foo")
		if err != nil {
			t.Fatal(err.Error())
		}
		if col.Name() != "foo" {
			t.Fatal("Collection 'foo' not present")
		}
		doc := map[string]interface{}{
			"name":  "John",
			"email": "john@example.com",
			"age":   29,
		}
		meta, err := col.CreateDocument(context.TODO(), doc)
		if err != nil {
			t.Fatal(err.Error())
		}
		_, err = col.RemoveDocument(context.TODO(), meta.Key)
		if err != nil {
			t.Fatal(err.Error())
		}
	})
}
