package storage

import (
	"database/sql"
	"fmt"
)

const (
	MigrationTeam = `
	CREATE TABLE IF NOT EXISTS teams(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(80) NOT NULL,
		description VARCHAR(100) NOT NULL,
		manager VARCHAR(80) NOT NULL,
		team INT NOT NULL,
		create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (team) REFERENCES players(id) ON UPDATE CASCADE ON DELETE CASCADE
	)ENGINE = InnoDB;
	`
)

type MySQLTeam struct {
	db *sql.DB
}

func NewTeamMySQL(db *sql.DB) *MySQLTeam {
	return &MySQLTeam{db}
}

func (m *MySQLTeam) Migrate() error {
	stm, err := m.db.Prepare(MigrationTeam)
	if err != nil {
		return err
	}
	defer stm.Close()

	_, err = stm.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migrate of Team was successfully")
	return nil
}
