package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Azamjon99/restaurant-staff-service/src/infrastructure/crypto"
	"github.com/Azamjon99/restaurant-staff-service/src/infrastructure/rand"

	"github.com/Azamjon99/restaurant-staff-service/src/domain/manager/models"
	"github.com/Azamjon99/restaurant-staff-service/src/domain/manager/repositories"
)

type ManagerService interface {
	RegisterManager(ctx context.Context, entityID, restaurantID, email, password string) (string, error)
	ChangeManagerPassword(ctx context.Context, managerID, currentPassword, newPassword string) error
	LoginManager(ctx context.Context, email, password string) (*models.ManagerProfile, error)
	GetManagerProfile(ctx context.Context, managerID string) (*models.ManagerProfile, error)
	UpdateManagerProfile(ctx context.Context, managerID, name, phoneNumber, email, imageUrl string) (*models.ManagerProfile, error)
	CreateManageRestaurantAssignment(ctx context.Context, managerID, restaurantID string) (int64, error)
	RemoveManagerRestaurantAssignment(ctx context.Context, assignmentID int) error
	GetManagerRestaurant(ctx context.Context, managerID string) (string, error)
}

type managerSvcImpl struct {
	managerRepo repositories.ManagerRepository
}

func NewManagerService(repo repositories.ManagerRepository) ManagerService {
	return &managerSvcImpl{
		managerRepo: repo,
	}
}

func (s *managerSvcImpl) RegisterManager(ctx context.Context, entityID, restaurantID, email, password string) (string, error) {

	manager, err := s.managerRepo.GetManager(ctx, email)
	if err == nil {
		return manager.ID, fmt.Errorf("manager with this eamil already exists: %s", email)
	}
	var (
		managerID   = rand.UUID()
		managerName = fmt.Sprintf("manager-%s", rand.NumericString(5))
		salt        = crypto.GenerateSalt()
		saltedPass  = crypto.Combine(salt, password)
		passHash    = crypto.HashPassword(saltedPass)
		now         = time.Now().UTC()
	)
	manager = &models.Manager{
		ID:           managerID,
		Email:        email,
		PasswordHash: passHash,
		PasswordSalt: salt,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	managerProfile := models.ManagerProfile{
		ManagerID:   managerID,
		EntityID:    entityID,
		Name:        managerName,
		ImageUrl:    "",
		PhoneNumber: "",
		Email:       email,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	managerRestaurant := models.ManagerRestaurantAssignment{
		ManagerID:    managerID,
		RestaurantID: restaurantID,
		CreatedAt:    now,
	}

	err = s.managerRepo.WithTx(ctx, func(r repositories.ManagerRepository) error {
		if err := r.SaveManager(ctx, manager); err != nil {
			return err
		}
		if err := r.SaveManagerProfile(ctx, &managerProfile); err != nil {
			return err
		}
		if err := r.SaveManagerRestaurantAssignment(ctx, &managerRestaurant); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return managerID, nil
}

func (s *managerSvcImpl) ChangeManagerPassword(ctx context.Context, managerID, currentPassword, newPassword string) error {

	manager, err := s.managerRepo.GetManager(ctx, managerID)
	if err != nil {
		return err
	}
	if !crypto.PasswordMatch(currentPassword, manager.PasswordSalt, manager.PasswordHash) {
		return errors.New("password doesn't match")
	}
	var (
		salt       = crypto.GenerateSalt()
		saltedPass = crypto.Combine(salt, newPassword)
		passHash   = crypto.HashPassword(saltedPass)
	)

	if manager.PasswordHash != passHash {
		return errors.New("invalid or missing currentPassword")
	}
	manager.UpdatePassword(passHash, saltedPass)

	err = s.managerRepo.UpdateManager(ctx, manager)

	if err != nil {
		return err
	}

	return nil
}
func (s *managerSvcImpl) LoginManager(ctx context.Context, email, password string) (*models.ManagerProfile, error) {

	manager, err := s.managerRepo.GetManagerByEmail(ctx, email)

	if err != nil {
		return nil, err
	}
	if !crypto.PasswordMatch(password, manager.PasswordSalt, manager.PasswordHash) {
		return nil, errors.New("password doesn't match")
	}

	profile, err := s.GetManagerProfile(ctx, manager.ID)

	if err != nil {
		return nil, err
	}
	return profile, nil
}
func (s *managerSvcImpl) DeleteManager(ctx context.Context, managerID string) error {
	return s.managerRepo.DeleteManager(ctx, managerID)
}
func (s *managerSvcImpl) GetManagerProfile(ctx context.Context, managerID string) (*models.ManagerProfile, error) {

	manager_profile, err := s.managerRepo.GetManagerProfile(ctx, managerID)

	if err != nil {
		return nil, err
	}
	return manager_profile, nil
}
func (s *managerSvcImpl) UpdateManagerProfile(ctx context.Context, managerID, name, phoneNumber, email, imageUrl string) (*models.ManagerProfile, error) {

	profile, err := s.managerRepo.GetManagerProfile(ctx, managerID)

	if err != nil {
		return nil, err
	}

	profile.Name = name
	profile.PhoneNumber = phoneNumber
	profile.Email = email
	profile.ImageUrl = imageUrl
	profile.UpdatedAt = time.Now().UTC()

	err = s.managerRepo.UpdateManagerProfile(ctx, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
func (s *managerSvcImpl) CreateManageRestaurantAssignment(ctx context.Context, managerID, restaurantID string) (int64, error) {

	assigment := models.ManagerRestaurantAssignment{
		ManagerID:    managerID,
		RestaurantID: restaurantID,
		CreatedAt:    time.Now().UTC(),
	}

	err := s.managerRepo.SaveManagerRestaurantAssignment(ctx, &assigment)
	if err != nil {
		return 0, err
	}

	return assigment.ID, nil
}
func (s *managerSvcImpl) RemoveManagerRestaurantAssignment(ctx context.Context, assignmentID int) error {

	err := s.managerRepo.DeleteManagerRestaurantAssignment(ctx, assignmentID)

	if err != nil {
		return err
	}
	return nil
}
func (s *managerSvcImpl) GetManagerRestaurant(ctx context.Context, managerID string) (string, error) {

	restaurantID, err := s.managerRepo.GetManagerRestaurant(ctx, managerID)

	if err != nil {
		return "", err
	}
	return restaurantID, nil
}
