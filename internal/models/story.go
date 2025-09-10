package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Story struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Excerpt     string    `json:"excerpt"`
	Price       float64   `json:"price"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type StoryCreate struct {
	Title   string  `json:"title" binding:"required"`
	Content string  `json:"content" binding:"required"`
	Excerpt string  `json:"excerpt"`
	Price   float64 `json:"price"`
}

type StoryUpdate struct {
	Title   *string  `json:"title,omitempty"`
	Content *string  `json:"content,omitempty"`
	Excerpt *string  `json:"excerpt,omitempty"`
	Price   *float64 `json:"price,omitempty"`
}

func GetAllStories(db *sql.DB) ([]Story, error) {
	query := `
		SELECT id, title, content, excerpt, price, published_at, created_at
		FROM stories
		WHERE published_at IS NOT NULL
		ORDER BY published_at DESC
	`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []Story
	for rows.Next() {
		var story Story
		err := rows.Scan(
			&story.ID,
			&story.Title,
			&story.Content,
			&story.Excerpt,
			&story.Price,
			&story.PublishedAt,
			&story.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		stories = append(stories, story)
	}

	return stories, nil
}

func GetStoryByID(db *sql.DB, id int) (*Story, error) {
	query := `
		SELECT id, title, content, excerpt, price, published_at, created_at
		FROM stories
		WHERE id = $1
	`
	
	var story Story
	err := db.QueryRow(query, id).Scan(
		&story.ID,
		&story.Title,
		&story.Content,
		&story.Excerpt,
		&story.Price,
		&story.PublishedAt,
		&story.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	
	return &story, nil
}

func CreateStory(db *sql.DB, story *StoryCreate) (*Story, error) {
	query := `
		INSERT INTO stories (title, content, excerpt, price, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, content, excerpt, price, published_at, created_at
	`
	
	var newStory Story
	err := db.QueryRow(
		query,
		story.Title,
		story.Content,
		story.Excerpt,
		story.Price,
		time.Now(),
	).Scan(
		&newStory.ID,
		&newStory.Title,
		&newStory.Content,
		&newStory.Excerpt,
		&newStory.Price,
		&newStory.PublishedAt,
		&newStory.CreatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &newStory, nil
}

func UpdateStory(db *sql.DB, id int, updates *StoryUpdate) (*Story, error) {
	// First, get the current story
	currentStory, err := GetStoryByID(db, id)
	if err != nil {
		return nil, err
	}
	if currentStory == nil {
		return nil, nil
	}

	// Build dynamic update query
	query := "UPDATE stories SET "
	args := []interface{}{}
	argIndex := 1
	
	if updates.Title != nil {
		query += fmt.Sprintf("title = $%d, ", argIndex)
		args = append(args, *updates.Title)
		argIndex++
	}
	if updates.Content != nil {
		query += fmt.Sprintf("content = $%d, ", argIndex)
		args = append(args, *updates.Content)
		argIndex++
	}
	if updates.Excerpt != nil {
		query += fmt.Sprintf("excerpt = $%d, ", argIndex)
		args = append(args, *updates.Excerpt)
		argIndex++
	}
	if updates.Price != nil {
		query += fmt.Sprintf("price = $%d, ", argIndex)
		args = append(args, *updates.Price)
		argIndex++
	}
	
	if len(args) == 0 {
		return currentStory, nil // No updates to make
	}
	
	// Remove trailing comma and add WHERE clause
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d RETURNING id, title, content, excerpt, price, published_at, created_at", argIndex)
	args = append(args, id)
	
	var updatedStory Story
	err = db.QueryRow(query, args...).Scan(
		&updatedStory.ID,
		&updatedStory.Title,
		&updatedStory.Content,
		&updatedStory.Excerpt,
		&updatedStory.Price,
		&updatedStory.PublishedAt,
		&updatedStory.CreatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &updatedStory, nil
}

func DeleteStory(db *sql.DB, id int) error {
	query := "DELETE FROM stories WHERE id = $1"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	
	return nil
}

func PublishStory(db *sql.DB, id int) (*Story, error) {
	query := `
		UPDATE stories 
		SET published_at = $1 
		WHERE id = $2 AND published_at IS NULL
		RETURNING id, title, content, excerpt, price, published_at, created_at
	`
	
	var story Story
	err := db.QueryRow(query, time.Now(), id).Scan(
		&story.ID,
		&story.Title,
		&story.Content,
		&story.Excerpt,
		&story.Price,
		&story.PublishedAt,
		&story.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	
	return &story, nil
}