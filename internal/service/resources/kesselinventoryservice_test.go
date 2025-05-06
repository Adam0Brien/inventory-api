package resources_test

import (
	"testing"

	"github.com/project-kessel/inventory-api/internal/biz/resources"
	service "github.com/project-kessel/inventory-api/internal/service/resources"
)

func TestNewKesselInventoryServiceV1beta2(t *testing.T) {
	mockUsecase := &resources.Usecase{}
	svc := service.NewKesselInventoryServiceV1beta2(mockUsecase)

	if svc == nil {
		t.Fatal("expected InventoryService instance, got nil")
	}
	if svc.Ctl != mockUsecase {
		t.Fatal("expected Usecase to be assigned")
	}
}
