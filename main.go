package main

import (
	"fmt"
	"os"

	"github.com/YourAverageMoron/articlereminder/config"
	"github.com/YourAverageMoron/articlereminder/reminders"
	"github.com/YourAverageMoron/articlereminder/store"
)

func main() {
	configPath := os.Args[1]
	c := config.NewConfig(configPath)
	err := c.Load()
	if err != nil {
		fmt.Println(err)
	}
	db, err := store.NewSQLiteStore(c.DBPath)
	if err != nil {
		fmt.Println(err)
	}
	l := reminders.NewList("Reading")
	articles, err := db.ReadRandomArticles(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, a := range articles {
		err := l.Add(a.Link)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = db.MarkRead(a.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
