package services

import (
	"gopay-clone/config"
	apperrors "gopay-clone/errors"
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
		if err == gorm.ErrRecordNotFound {
			return apperrors.ErrAccountNotFound
		}
		return apperrors.NewInternalError("Failed to verify receiver account")
	}

	qr.ReceiverAccount = account
	if err := s.db.Create(qr).Error; err != nil {
		return apperrors.ErrQRCreateFailed
	}

	return nil
}

func (s *QRService) GetQRById(id uint) (*models.QrCode, error) {
	var qr models.QrCode
	err := s.db.First(&qr, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrQRNotFound
		}
		return nil, apperrors.NewInternalError("Failed to fetch QR code")
	}

	return &qr, nil
}

func (s *QRService) ScanQR(qr *models.QrCode, senderAccountId uint) error {
	if qr.ExpiresAt.Before(time.Now()) {
		return apperrors.ErrQRExpired
	}

	if qr.IsUsed {
		return apperrors.ErrQRAlreadyUsed
	}

	var sender models.Account
	var receiver models.Account

	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&sender, senderAccountId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return apperrors.ErrAccountNotFound
			}
			return apperrors.NewInternalError("Failed to fetch sender account")
		}

		if err := tx.First(&receiver, qr.ReceiverAccountID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return apperrors.ErrAccountNotFound
			}
			return apperrors.NewInternalError("Failed to fetch receiver account")
		}

		if sender.Balance < qr.Amount {
			return apperrors.ErrInsufficientBalance
		}

		sender.Balance -= qr.Amount
		receiver.Balance += qr.Amount
		qr.IsUsed = true

		if err := tx.Save(&sender).Error; err != nil {
			return apperrors.ErrTransactionFailed
		}
		if err := tx.Save(&receiver).Error; err != nil {
			return apperrors.ErrTransactionFailed
		}

		createdTransaction := models.Transaction{
			Amount:            qr.Amount,
			SenderAccountID:   sender.ID,
			SenderAccount:     sender,
			ReceiverAccountID: receiver.ID,
			ReceiverAccount:   receiver,
			QrCodeID:          &qr.ID,
			QrCode:            qr,
			Status:            models.TransactionCompleted,
		}

		if err := tx.Create(&createdTransaction).Error; err != nil {
			return apperrors.ErrTransactionFailed
		}

		if err := tx.Save(qr).Error; err != nil {
			return apperrors.ErrTransactionFailed
		}

		return nil
	})
}
