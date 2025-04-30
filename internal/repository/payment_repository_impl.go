package repository

import (
	"payment-options/internal/models"
	"time"
)

type paymentRepo struct{}

func NewPaymentRepo() PaymentRepository {
	return &paymentRepo{}
}

func (r *paymentRepo) CallDanamon() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "7890123456",
		Status:  "Active",
		Balance: "700000",
		Icon:    "https://sampleurl.com/danamon.jpg",
	}
}

func (r *paymentRepo) CallBTN() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "8901234567",
		Status:  "Active",
		Balance: "800000",
		Icon:    "https://sampleurl.com/btn.jpg",
	}
}

func (r *paymentRepo) CallBSI() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "9012345678",
		Status:  "Active",
		Balance: "900000",
		Icon:    "https://sampleurl.com/bsi.jpg",
	}
}

func (r *paymentRepo) CallMega() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "0123456789",
		Status:  "Active",
		Balance: "1000000",
		Icon:    "https://sampleurl.com/mega.jpg",
	}
}

func (r *paymentRepo) CallOCBC() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "1123456789",
		Status:  "Active",
		Balance: "1100000",
		Icon:    "https://sampleurl.com/ocbc.jpg",
	}
}

func (r *paymentRepo) CallMaybank() models.PaymentMethod {
	time.Sleep(1 * time.Second)
	return models.PaymentMethod{
		Account: "1223456789",
		Status:  "Active",
		Balance: "1200000",
		Icon:    "https://sampleurl.com/maybank.jpg",
	}
}
