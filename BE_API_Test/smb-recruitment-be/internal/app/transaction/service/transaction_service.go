package service

import (
	"context"
	"fmt"
	"github.com/tunaiku/mobilebanking/internal/app/domain"
	"github.com/tunaiku/mobilebanking/internal/app/savings/service/fake"
	"github.com/tunaiku/mobilebanking/internal/app/transaction/dto"
)

type ITransactionService interface {
	CreateTransaction(request *dto.CreateTransactionRequest, ctx context.Context) (*dto.CreateTransactionSuccess, error)
}

type TransactionService struct {
}

func NewTransactionService() ITransactionService {
	return &TransactionService{}
}

func (s *TransactionService) CreateTransaction(request *dto.CreateTransactionRequest, ctx context.Context) (*dto.CreateTransactionSuccess, error) {
	//var userSessionHelper domain.UserSessionHelper
	//userSession, err := userSessionHelper.GetFromContext(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//accountNumber := userSession.AccountReference
	var accountInformationService = fake.NewFakeAccountInformationService()
	//_, err = accountInformationService.GetTransactionPrivileges(accountNumber)
	isExists := accountInformationService.IsAccountExists(request.DestinationAccount)
	var trxInformatonService = fake.NewFakeTransactionInformationService()
	if !isExists {
		return nil, nil
	}

	if detail, err := trxInformatonService.FindTransactionDetailByCode(request.TransactionCode); err != nil {
		return nil, err
	} else {
		minimumAmount := detail.MinimumAmount
		if minimumAmount.Cmp(request.Amount) == 1 {
			return nil, nil
		}
	}

	var user domain.User
	if request.AuthorizationMethod == "PIN" {
		user.ConfiguredTransactionCredential.Pin.Pin = "PIN"
		if user.ConfiguredTransactionCredential.IsPinConfigured() {
			transaction := &dto.CreateTransactionSuccess{}
			return transaction, nil
		}
	}

	if request.AuthorizationMethod == "OTP" {
		user.ConfiguredTransactionCredential.Otp.PhoneNumber = "OTP"
		if user.ConfiguredTransactionCredential.IsOtpConfigured() {
			transaction := &dto.CreateTransactionSuccess{}
			return transaction, nil
		}
	}

	fmt.Println("=====================================>")
	transaction := &dto.CreateTransactionSuccess{}
	transaction.TransactionID = "Create transaction failed!"
	return transaction, nil
}
