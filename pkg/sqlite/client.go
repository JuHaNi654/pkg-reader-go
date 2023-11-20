package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/JuHaNi654/pkg-reader/pkg/models"
)

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}

func (r *SQLRepository) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS packages(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT
		);
	`

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLRepository) Insert(pkg *models.Pkg) error {
	q := `INSERT INTO packages(name, description) values (?,?);`
	res, err := r.db.Exec(q, pkg.Name, pkg.Description)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	pkg.Id = id
	return nil
}

func (r *SQLRepository) GetItems(page int) ([]models.Pkg, error) {
	q := fmt.Sprintf(`
	SELECT 
		id, 
		name, 
		description 
	FROM packages
	ORDER BY id
	LIMIT %d
	OFFSET %d;`,
		10, (10 * (page)))
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []models.Pkg
	for rows.Next() {
		var pkg models.Pkg
		err := rows.Scan(
			&pkg.Id,
			&pkg.Name,
			&pkg.Description,
		)

		if err != nil {
			return nil, err
		}
		items = append(items, pkg)
	}

	return items, nil
}
