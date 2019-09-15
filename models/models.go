package models

import (
 
    // Import because this is a mysql model
    "strings"
 
    // we are using mysql
    _ "github.com/mattn/go-sqlite3"
    "github.com/jmoiron/sqlx"
    "github.com/jmoiron/sqlx/reflectx"
 
    "github.com/ldb358/GRPCMTGInventory/api"
)
 
//InventoryStore the interface that all Inventory models must implement
type InventoryStore interface {
	GetInventory(id int) (*api.Inventory, error)
	CreateInventory(name string)  (*api.Inventory, error)
}

//DB the database object that models are implimented on top of
type DB struct {
    *sqlx.DB
}

var schema = []string{
    `CREATE TABLE inventory (
        id INTEGER PRIMARY KEY,
        name VARCHAR(80) NOT NULL
    );`,
    `CREATE TABLE mtgcard (
        inventory_id INTEGER,
        name VARCHAR(80),
        qty INTEGER,
        PRIMARY KEY (inventory_id, name),
        FOREIGN KEY(inventory_id) REFERENCES inventory(id)
    );`,
}

//NewDB Passed a database connection string create a database connection
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Open("sqlite3", dataSourceName)
    db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}
 
//GetInventory Fetch an inventory 
func (db *DB) GetInventory(id int) (api.Inventory, error) {
    inventory := api.Inventory{}
    err := db.Get(&inventory, "SELECT id, name FROM inventory WHERE id=$1", id)
    return inventory, err
}

//CreateInventory Create an inventory
func (db *DB) CreateInventory(name string) (*api.Inventory, error) {
    inventory := api.Inventory{
        Name: name,
    }
    tx, err := db.Begin()
    if err != nil {
        return nil, err
    }
    res, err := tx.Exec("INSERT INTO inventory (name) VALUES ($1)", name)
    if err != nil {
        return nil, err
    }
    tx.Commit()
    inventory.Id, err = res.LastInsertId()
    if err != nil {
        return nil, err
    }
    return &inventory, err
}