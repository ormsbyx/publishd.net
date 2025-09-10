package database

import (
	"fmt"
)

func RunMigrations() error {
	if DB == nil {
		return fmt.Errorf("database connection not initialized")
	}

	// Create stories table
	createStoriesTable := `
		CREATE TABLE IF NOT EXISTS stories (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			excerpt TEXT,
			price DECIMAL(10,2) DEFAULT 0.00,
			published_at TIMESTAMP,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	if _, err := DB.Exec(createStoriesTable); err != nil {
		return fmt.Errorf("failed to create stories table: %w", err)
	}

	// Create users table (for future use)
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			subscription_active BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	if _, err := DB.Exec(createUsersTable); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Create purchases table (for future use)
	createPurchasesTable := `
		CREATE TABLE IF NOT EXISTS purchases (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			story_id INTEGER REFERENCES stories(id) ON DELETE CASCADE,
			purchased_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(user_id, story_id)
		);
	`

	if _, err := DB.Exec(createPurchasesTable); err != nil {
		return fmt.Errorf("failed to create purchases table: %w", err)
	}

	// Create subscriptions table (for future use)
	createSubscriptionsTable := `
		CREATE TABLE IF NOT EXISTS subscriptions (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			stripe_subscription_id VARCHAR(255) UNIQUE,
			active BOOLEAN DEFAULT TRUE,
			expires_at TIMESTAMP,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	if _, err := DB.Exec(createSubscriptionsTable); err != nil {
		return fmt.Errorf("failed to create subscriptions table: %w", err)
	}

	fmt.Println("✅ Database migrations completed successfully")
	return nil
}