package services

import (
	"errors"
	"gopay-clone/config"
	"gopay-clone/models"
	"time"

	"gorm.io/gorm"
)

type QRService struct {
	db *config.Database
}

func NewQRService(db *config.Database) *QRService {
	return &QRService{db: db}
}

func (s *QRService) CreateQR(qr *models.QrCode) error {
	var account models.Account
	if err := s.db.First(&account, qr.ReceiverAccountID).Error; err != nil {
		return errors.New("not found")
	}
	qr.ReceiverAccount = account
	return s.db.Create(qr).Error
}

func (s *QRService) GetQRById(id uint) (*models.QrCode, error) {
	var qr models.QrCode
	result := s.db.
		First(&qr, id)
	return &qr, result.Error
}

func (s *QRService) ScanQR(qr *models.QrCode, senderAccountId uint) error {
	if qr.ExpiresAt.Before(time.Now()) {
		return errors.New("QR code has expired")
	}
	var sender models.Account
	var receiver models.Account
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&sender, senderAccountId).Error; err != nil {
			return err
		}
		if err := tx.First(&receiver, qr.ReceiverAccountID).Error; err != nil {
			return err
		}
		if sender.Balance < qr.Amount {
			return errors.New("insufficient balance")
		}
		sender.Balance -= qr.Amount
		receiver.Balance += qr.Amount
		qr.IsUsed = true

		if err := tx.Save(&sender).Error; err != nil {
			return err
		}
		if err := tx.Save(&receiver).Error; err != nil {
			return err
		}
		createdTransaction := models.Transaction{
			Amount:            qr.Amount,
			SenderAccountId:   sender.ID,
			SenderAccount:     sender,
			ReceiverAccountId: receiver.ID,
			ReceiverAccount:   receiver,
			QrCodeID:          &qr.ID,
			QrCode:            qr,
			Status:            "paid",
		}
		if err := tx.Create(&createdTransaction).Error; err != nil {
			return err
		}
		return nil
	})
}
