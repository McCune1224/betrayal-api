package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Item struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Description string         `db:"description" json:"description"`
	Cost        int64          `db:"cost" json:"cost"`
	Rarity      string         `db:"rarity" json:"rarity"`
	Categories  pq.StringArray `db:"categories" json:"categories"`
}

type ItemModel struct {
	Db *sqlx.DB
}

func (im *ItemModel) GetItem(id int64) (Item, error) {
	var item Item
	err := im.Db.Get(&item, "SELECT * FROM items WHERE id = $1", id)
	return item, err
}

func (im *ItemModel) GetItemByName(name string) (Item, error) {
	var item Item
	err := im.Db.Get(&item, "SELECT * FROM items WHERE name ILIKE $1", name)
	return item, err
}

// Will return all items in the database within the pagination range
// Hard Set Limit of 100 items per page, and default page size is 10
func (im *ItemModel) PageItems(offset, limit int) ([]Item, error) {
	var items []Item
	err := im.Db.Select(&items, "SELECT * FROM items ORDER BY id ASC LIMIT $1 OFFSET $2", limit, (offset-1)*limit)
	return items, err
}

func (im *ItemModel) QueryItems(query string, page, limit int) ([]Item, error) {
	var items []Item
	err := im.Db.Select(&items, "SELECT * FROM items WHERE name ILIKE $1 ORDER BY id ASC LIMIT $2 OFFSET $3", query, limit, (page-1)*limit)
	return items, err
}

func (im *ItemModel) CreateItem(item Item) (int64, error) {
	var id int64
	err := im.Db.QueryRow("INSERT INTO items DEFAULT VALUES RETURNING id").Scan(&id)
	return id, err
}

func (im *ItemModel) UpdateItem(item Item) error {
	_, err := im.Db.Exec("UPDATE items SET name = $1, description = $2, cost = $3, rarity = $4, categories = $5 WHERE id = $6",
		item.Name, item.Description, item.Cost, item.Rarity, item.Categories, item.ID)
	return err
}
