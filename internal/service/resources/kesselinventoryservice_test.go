package resources_test

import (
	"github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta2"
	"testing"

	service "github.com/project-kessel/inventory-api/internal/service/resources"
)

func TestViewResponseFromAuthzRequestV1beta2(t *testing.T) {
	resp := service.ViewResponseFromAuthzRequestV1beta2(true)
	if resp.Allowed != v1beta2.Allowed_ALLOWED_TRUE {
		t.Errorf("expected Allowed_TRUE, got %v", resp.Allowed)
	}

	resp = service.ViewResponseFromAuthzRequestV1beta2(false)
	if resp.Allowed != v1beta2.Allowed_ALLOWED_FALSE {
		t.Errorf("expected Allowed_FALSE, got %v", resp.Allowed)
	}
}
