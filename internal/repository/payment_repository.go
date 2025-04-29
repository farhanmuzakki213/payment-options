package repository

import "payment-options/internal/models"

type PaymentRepository interface {
	CallDanamon() models.PaymentMethod
	CallBTN() models.PaymentMethod
	CallBSI() models.PaymentMethod
	CallMega() models.PaymentMethod
	CallOCBC() models.PaymentMethod
	CallMaybank() models.PaymentMethod
}
