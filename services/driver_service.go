package services

import (
	"gopay-clone/config"
	apperrors "gopay-clone/errors"
	"gopay-clone/models"

	"gorm.io/gorm"
)

type DriverService struct {
	db *config.Database
}

func NewDriverService(db *config.Database) *DriverService {
	return &DriverService{db: db}
}

func (s *DriverService) CreateDriverProfile(driver *models.DriverProfile) error {
	// Check if driver already exists for this user
	var existingDriver models.DriverProfile
	if err := s.db.Where("user_id = ?", driver.UserId).First(&existingDriver).Error; err == nil {
		return apperrors.ErrDriverExists
	}

	if err := s.db.Create(driver).Error; err != nil {
		return apperrors.ErrDriverCreation
	}
	return nil
}

func (s *DriverService) GetDriverByID(id uint) (*models.DriverProfile, error) {
	var driver models.DriverProfile
	if err := s.db.Preload("User").First(&driver, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrDriverNotFound
		}
		return nil, apperrors.ErrDatabaseError
	}
	return &driver, nil
}

func (s *DriverService) GetDriverByUserID(userID uint) (*models.DriverProfile, error) {
	var driver models.DriverProfile
	if err := s.db.Preload("User").Where("user_id = ?", userID).First(&driver).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrDriverNotFound
		}
		return nil, apperrors.ErrDatabaseError
	}
	return &driver, nil
}

func (s *DriverService) GetAllDrivers() ([]models.DriverProfile, error) {
	var drivers []models.DriverProfile
	if err := s.db.Preload("User").Find(&drivers).Error; err != nil {
		return nil, apperrors.ErrDatabaseError
	}
	return drivers, nil
}

func (s *DriverService) GetAvailableDrivers() ([]models.DriverProfile, error) {
	var drivers []models.DriverProfile
	if err := s.db.Preload("User").Where("status = ? AND is_verified = ?", "online", true).Find(&drivers).Error; err != nil {
		return nil, apperrors.ErrDatabaseError
	}

	if len(drivers) == 0 {
		return nil, apperrors.ErrDriverUnavailable
	}

	return drivers, nil
}

func (s *DriverService) GetDriversByVehicleType(vehicleType models.VehicleType) ([]models.DriverProfile, error) {
	var drivers []models.DriverProfile
	if err := s.db.Preload("User").Where("vehicle_type = ? AND status = ? AND is_verified = ?", vehicleType, "online", true).Find(&drivers).Error; err != nil {
		return nil, apperrors.ErrDatabaseError
	}

	if len(drivers) == 0 {
		return nil, apperrors.ErrDriverUnavailable
	}

	return drivers, nil
}

func (s *DriverService) UpdateDriver(id uint, updates map[string]any) error {
	result := s.db.Model(&models.DriverProfile{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return apperrors.ErrDriverProfileUpdateFailed
	}

	if result.RowsAffected == 0 {
		return apperrors.ErrDriverNotFound
	}

	return nil
}

func (s *DriverService) UpdateDriverStatus(driverID uint, status string) error {
	result := s.db.Model(&models.DriverProfile{}).Where("id = ?", driverID).Update("status", status)
	if result.Error != nil {
		return apperrors.ErrDriverStatusUpdateFailed
	}

	if result.RowsAffected == 0 {
		return apperrors.ErrDriverNotFound
	}

	return nil
}

func (s *DriverService) UpdateDriverLocation(userID uint, location string) error {
	result := s.db.Model(&models.DriverProfile{}).Where("user_id = ?", userID).Update("current_location", location)
	if result.Error != nil {
		return apperrors.ErrDriverProfileUpdateFailed
	}

	if result.RowsAffected == 0 {
		return apperrors.ErrDriverNotFound
	}

	return nil
}

func (s *DriverService) DeleteDriver(id uint) error {
	result := s.db.Delete(&models.DriverProfile{}, id)
	if result.Error != nil {
		return apperrors.ErrDriverDeleteFailed
	}

	if result.RowsAffected == 0 {
		return apperrors.ErrDriverNotFound
	}

	return nil
}

func (s *DriverService) VerifyDriver(id uint) error {
	result := s.db.Model(&models.DriverProfile{}).Where("id = ?", id).Update("is_verified", true)
	if result.Error != nil {
		return apperrors.ErrDriverProfileUpdateFailed
	}

	if result.RowsAffected == 0 {
		return apperrors.ErrDriverNotFound
	}

	return nil
}
