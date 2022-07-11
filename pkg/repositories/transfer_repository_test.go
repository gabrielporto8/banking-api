package repositories_test

import (
	"testing"

	"github.com/gabrielporto8/banking-api/pkg/models"
	"github.com/gabrielporto8/banking-api/pkg/repositories"
)

var transfers = map[int64]*models.Transfer{
	0: {
		ID: 0,
		AccountOriginID: 0,
		AccountDestinationID: 1,
		Amount: 50,
	},
	1: {
		ID: 1,
		AccountOriginID: 1,
		AccountDestinationID: 0,
		Amount: 20,
	},
}

func TestSaveTransfer(t *testing.T) {
	repository := repositories.NewTransferRepository(make(map[int64]*models.Transfer))

	transfer := models.Transfer{
		ID: 5,
		AccountOriginID: 1,
		AccountDestinationID: 2,
		Amount: 90,
	}

	got := repository.SaveTransfer(&transfer)

	if got != nil {
		t.Errorf("want nil, got '%v", got.Error())
	}
}

func TestGetTransfers(t *testing.T) {
	repository := repositories.NewTransferRepository(transfers)

	want := 2
	got := repository.GetTransfers()

	if len(got) != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetTransfersByOriginID(t *testing.T) {
	repository := repositories.NewTransferRepository(transfers)

	t.Run("should return correct transfer", func(t *testing.T) {
		want := accounts[0].ID
		got := repository.GetTransfersByOriginID(0)

		if got[0].ID != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("should return empty slice for not existent transfers", func(t *testing.T) {
		want := 0
		got := repository.GetTransfersByOriginID(3)

		if len(got) != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
