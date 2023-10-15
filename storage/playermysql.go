package storage

import (
	"database/sql"
	"fmt"
)

const (
	MigrationPlayers = `
	CREATE TABLE IF NOT EXISTS players(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(80) NOT NULL,
		position VARCHAR(50) NOT NULL,
		number INT NOT NULL,
		create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)ENGINE = InnoDB;
	`
)

type MySQLPlayers struct {
	db *sql.DB
}

func NewPlayersMySQL(db *sql.DB) *MySQLPlayers {
	return &MySQLPlayers{db}
}

func (m *MySQLPlayers) Migrate() error {
	stm, err := m.db.Prepare(MigrationPlayers)
	if err != nil {
		return err
	}
	defer stm.Close()

	_, err = stm.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migrate of Players was successfully")
	return nil
}
