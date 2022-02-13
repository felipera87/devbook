package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications represents a repository of publications
type Publications struct {
	db *sql.DB
}

// NewPublicationRepository creates a new repository of publications
func NewPublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

// Create inserts a publication on database
func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare("insert into publications (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

// SearchByID gets a single publication from the database
func (repository Publications) SearchByID(publicationID uint64) (models.Publication, error) {
	row, err := repository.db.Query(`
		select p.*, u.nick from
		publications p inner join users u
		on u.id = p.author_id where p.id = ?
	`, publicationID)
	if err != nil {
		return models.Publication{}, err
	}
	defer row.Close()

	var publication models.Publication

	if row.Next() {
		if err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}
