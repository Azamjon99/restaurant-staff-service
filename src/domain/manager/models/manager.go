package models

import "time"

type Manager struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"index:idx_manager_by_email"`
	PasswordHash string    `json:"password_hash"`
	PasswordSalt string    `json:"password_salt"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
func (manager *Manager) UpdatePassword(hash,salt string) {
	manager.PasswordHash = hash
	manager.PasswordSalt = salt
	manager.UpdatedAt = time.Now()
}