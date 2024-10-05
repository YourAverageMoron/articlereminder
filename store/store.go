package store

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Item struct {
	ID          int
	Author      string
	Title       string
	Favourite   bool
	FeedURL     string
	FeedName    string // added from config if set
	Link        string
	Content     string
	ReadAt      time.Time
	PublishedAt time.Time
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type SQLiteStore struct {
	path string
	db   *sql.DB
}

func NewSQLiteStore(filePath string) (*SQLiteStore, error) {
	dbpath := "file:" + filePath
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, fmt.Errorf("NewSQLiteCache: %w", err)
	}
	return &SQLiteStore{
		path: dbpath,
		db:   db,
	}, nil
}

func (sls *SQLiteStore) MarkRead(id int) error {
	stmt, _ := sls.db.Prepare(`update items set readat = case when readat is null then ? else null end where id = ?`)
	_, err := stmt.Exec(time.Now(), id)
	return err
}

func (sls *SQLiteStore) ReadRandomArticles(n int) ([]Item, error) {
	stmt := `
		select id, feedurl, link, title, content, author, readat, favourite, publishedat, createdat, updatedat from items where readat is null order by RANDOM() limit %d;
	`
	stmt = fmt.Sprintf(stmt, n)

	rows, err := sls.db.Query(stmt)
	if err != nil {
		return []Item{}, fmt.Errorf("store.go: GetAllItems: %w", err)
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		var readAtNull sql.NullTime
		var publishedAtNull sql.NullTime
		var linkNull sql.NullString

		if err := rows.Scan(&item.ID, &item.FeedURL, &linkNull, &item.Title, &item.Content, &item.Author, &readAtNull, &item.Favourite, &publishedAtNull, &item.CreatedAt, &item.UpdatedAt); err != nil {
			fmt.Println("errrerre: ", err)
			continue
		}

		item.Link = linkNull.String
		item.ReadAt = readAtNull.Time
		item.PublishedAt = publishedAtNull.Time

		items = append(items, item)
	}
	return items, nil
}

// TODO: MARK AS READ -> TAKE IN AN ARTICLE AND MARK AS READ
