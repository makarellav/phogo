package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/makarellav/phogo/rand"
)

const MinSessionTokenBytes = 32

type Session struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB            *sql.DB
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	if ss.BytesPerToken < MinSessionTokenBytes {
		ss.BytesPerToken = MinSessionTokenBytes
	}

	token, err := rand.String(ss.BytesPerToken)

	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}

	row := ss.DB.QueryRow(`
INSERT INTO sessions(user_id, token_hash) 
VALUES ($1, $2) ON CONFLICT (user_id) 
DO UPDATE SET token_hash = $2 
RETURNING id`,
		session.UserID, session.TokenHash)

	err = row.Scan(&session.ID)

	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.hash(token)

	row := ss.DB.QueryRow(`SELECT users.id, users.email, users.password_hash FROM users 
INNER JOIN sessions on users.id = sessions.user_id
WHERE sessions.token_hash = $1`, tokenHash)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	hash := ss.hash(token)

	_, err := ss.DB.Exec(`
DELETE FROM sessions 
WHERE token_hash = $1`,
		hash)

	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func (ss *SessionService) hash(token string) string {
	hash := sha256.Sum256([]byte(token))

	return base64.URLEncoding.EncodeToString(hash[:])
}
