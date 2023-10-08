package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-staff-service/src/domain/manager/models"
	"github.com/Azamjon99/restaurant-staff-service/src/domain/manager/repositories"

	"gorm.io/gorm"
)

const (
	tableManagers                    = "manager.managers"
	tableManagerProfile              = "manager.manager_profiles"
	tableManagerRestaurantAssignment = "manager.manager_restaurant_assignments"
)

type managerRepoImpl struct {
	db *gorm.DB
}

func NewMangerRepository(db *gorm.DB) repositories.ManagerRepository {
	return &managerRepoImpl{
		db: db,
	}
}

func (r *managerRepoImpl) WithTx(ctx context.Context, f func(r repositories.ManagerRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := managerRepoImpl{
			db: tx,
		}
		if err := f(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
func (r *managerRepoImpl) SaveManager(ctx context.Context, manager *models.Manager) error {

	result := r.db.WithContext(ctx).Table(tableManagers).Create(manager)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *managerRepoImpl) UpdateManager(ctx context.Context, manager *models.Manager) error {

	result := r.db.Table(tableManagers).Save(manager)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *managerRepoImpl) GetManager(ctx context.Context, managerID string) (*models.Manager, error) {

	var manager models.Manager
	result := r.db.WithContext(ctx).Table(tableManagers).First(&manager, "id = ?", managerID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &manager, nil
}
func (r *managerRepoImpl) DeleteManager(ctx context.Context, managerID string) error {

	var manager models.Manager

	result := r.db.WithContext(ctx).Table(tableManagers).Delete(&manager, "id = ?", managerID)

	return result.Error
}
func (r *managerRepoImpl) GetManagerByEmail(ctx context.Context, email string) (*models.Manager, error) {

	var manager models.Manager
	result := r.db.WithContext(ctx).Table(tableManagers).First(&manager, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	} 
	return &manager, nil
}
func (r *managerRepoImpl) SaveManagerProfile(ctx context.Context, profile *models.ManagerProfile) error {

	result := r.db.WithContext(ctx).Table(tableManagerProfile).Create(profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *managerRepoImpl) UpdateManagerProfile(ctx context.Context, profile *models.ManagerProfile) error {

	result := r.db.WithContext(ctx).Table(tableManagerProfile).Save(profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *managerRepoImpl) GetManagerProfile(ctx context.Context, managerID string) (*models.ManagerProfile, error) {

	var profile models.ManagerProfile
	result := r.db.WithContext(ctx).Table(tableManagerProfile).First(&profile, "manager_id = ?", managerID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &profile, nil
}
func (r *managerRepoImpl) SaveManagerRestaurantAssignment(ctx context.Context, assignment *models.ManagerRestaurantAssignment) error {

	result := r.db.WithContext(ctx).Table(tableManagerRestaurantAssignment).Create(assignment)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *managerRepoImpl) DeleteManagerRestaurantAssignment(ctx context.Context, assignmentID int) error {

	var manager_restaurant_assignment models.ManagerRestaurantAssignment
	result := r.db.WithContext(ctx).Table(tableManagerRestaurantAssignment).Delete(&manager_restaurant_assignment, "id = ?", assignmentID)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *managerRepoImpl) GetManagerRestaurant(ctx context.Context, managerID string) (string, error) {

	var manager_restaurant models.ManagerRestaurantAssignment

	result := r.db.WithContext(ctx).Table(tableManagerRestaurantAssignment).First(&manager_restaurant, "manager_id = ?", managerID)

	if result.Error != nil {
		return "", result.Error
	}
	return manager_restaurant.RestaurantID, nil

}
