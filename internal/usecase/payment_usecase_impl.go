package usecase

import (
	"payment-options/internal/models"
	"payment-options/internal/repository"
	/* "sync" */
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
	/* var wg sync.WaitGroup
	result := make(map[string]models.PaymentMethod)
	mu := sync.Mutex{}

	wg.Add(6)

	go func() {
		defer wg.Done()
		method := u.repo.CallDanamon()
		mu.Lock()
		result["danamon"] = method
		mu.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		method := u.repo.CallBTN()
		mu.Lock()
		result["btn"] = method
		mu.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		method := u.repo.CallBSI()
		mu.Lock()
		result["bsi"] = method
		mu.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		method := u.repo.CallMega()
		mu.Lock()
		result["mega"] = method
		mu.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		method := u.repo.CallOCBC()
		mu.Lock()
		result["ocbc"] = method
		mu.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		method := u.repo.CallMaybank()
		mu.Lock()
		result["maybank"] = method
		mu.Unlock()
	}()

	wg.Wait() */
	result := make(map[string]models.PaymentMethod)

	result["danamon"] = u.repo.CallDanamon()
	result["btn"] = u.repo.CallBTN()
	result["bsi"] = u.repo.CallBSI()
	result["mega"] = u.repo.CallMega()
	result["ocbc"] = u.repo.CallOCBC()
	result["maybank"] = u.repo.CallMaybank()
	return result, nil
}
