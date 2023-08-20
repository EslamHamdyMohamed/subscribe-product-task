package test

import (
	"testing"
	"transfer_money/app/services"
	"transfer_money/helpers"
)

func TestFromAccountExist(t *testing.T) {
	request := setRequest("11", "2", 50)
	err := services.TransferService(request)
	if !helpers.CheckError(err) {
		t.Errorf("Error FromAccount Required: %s", err)
	}
}

func TestToAccountExist(t *testing.T) {
	request := setRequest("1", "22", 200)
	err := services.TransferService(request)
	if !helpers.CheckError(err) {
		t.Errorf("Error ToAccount Required: %s", err)
	}

}

func TestBalanceExist(t *testing.T) {
	request := setRequest("1", "2", 0)
	err := services.TransferService(request)
	if !helpers.CheckError(err) {
		t.Errorf("Error Balance Required: %s", err)
	}

}

func TestBalanceEnough(t *testing.T) {
	request := setRequest("1", "2", 200)
	err := services.TransferService(request)
	if !helpers.CheckError(err) {
		t.Errorf("Error check enough balance: %s", err)
	}
}

func TestNegativeBalance(t *testing.T) {
	request := setRequest("1", "2", -1)
	err := services.TransferService(request)
	if !helpers.CheckError(err) {
		t.Errorf("Error balance can not be negative: %s", err)
	}
}

func TestValidCase(t *testing.T) {
	request := setRequest("1", "2", 50)
	err := services.TransferService(request)
	if !helpers.CheckError(err) {
		t.Errorf("Error balance can not be negative: %s", err)
	}
}
