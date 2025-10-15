package repository

import (
	"context"
	"database/sql"
	"errors"

	"01_exercise/internal/entity"
)

// Interface repository — biar gampang di-mock saat testing
type IUserRepository interface {
	GetByEmail(ctx context.Context, email string) (entity.Users, error)
	RegisterUser(ctx context.Context, name, email, hashed string) error
	CreateStaff(ctx context.Context, name, email, hashed string) error
}

// Struct repository implementasi interface di atas
type UserRepository struct {
	DB *sql.DB
}

// Constructor
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Get user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (entity.Users, error) {
	var user entity.Users

	query := `
    SELECT 
        u.id, u.name, u.email, u.password,
        r.role_id, r.name AS role_name
    FROM users u
    JOIN roles r ON u.role_id = r.role_id
    WHERE u.email = $1
    `

	err := r.DB.QueryRowContext(ctx, query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.RoleName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

// Register new user
func (r *UserRepository) RegisterUser(ctx context.Context, name, email, hashed string) error {
	query := `INSERT INTO users (name, email, password, role_id) VALUES ($1, $2, $3, 2)`
	_, err := r.DB.ExecContext(ctx, query, name, email, hashed)
	return err
}

// Create staff user
func (r *UserRepository) CreateStaff(ctx context.Context, name, email, hashed string) error {
	query := `INSERT INTO users (name, email, password, role_id) VALUES ($1, $2, $3, 3)`
	_, err := r.DB.ExecContext(ctx, query, name, email, hashed)
	return err
}
