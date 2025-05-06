package model_test

import (
	"github.com/google/uuid"
	"testing"

	"github.com/project-kessel/inventory-api/internal/biz/model"
	"github.com/stretchr/testify/assert"
)

func TestReporterResourceIdv1beta2FromResource(t *testing.T) {
	resource := &model.Resource{
		ReporterResourceId: "resource-123",
		ResourceType:       "type-abc",
		ReporterType:       "reporter-type",
		ReporterInstanceId: "instance-456",
	}

	result := model.ReporterResourceIdv1beta2FromResource(resource)

	assert.Equal(t, resource.ReporterResourceId, result.ReporterResourceId)
	assert.Equal(t, resource.ResourceType, result.ResourceType)
	assert.Equal(t, resource.ReporterType, result.ReporterType)
	assert.Equal(t, resource.ReporterInstanceId, result.ReporterInstanceId)
}

func TestResource_BeforeCreate_AssignsUUID(t *testing.T) {
	resource := &model.Resource{
		ID: uuid.Nil, // no ID yet
	}

	err := resource.BeforeCreate(nil)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, resource.ID)
}

func TestResourceReporter_ValueAndScan(t *testing.T) {
	reporter := model.ResourceReporter{
		LocalResourceId: "local-123",
	}

	// Test Value (marshal)
	val, err := reporter.Value()
	assert.NoError(t, err)

	// Test Scan (unmarshal)
	var newReporter model.ResourceReporter
	err = newReporter.Scan(val)
	assert.NoError(t, err)
	assert.Equal(t, reporter.LocalResourceId, newReporter.LocalResourceId)
}
