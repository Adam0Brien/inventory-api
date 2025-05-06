package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestK8SClusterFields(t *testing.T) {
	metadata := &Metadata{}
	reporterData := &ReporterData{}
	resourceData := &K8SClusterDetail{}

	cluster := &K8SCluster{
		Metadata:     metadata,
		ReporterData: reporterData,
		ResourceData: resourceData,
	}

	t.Run("GetMetadata returns correct value", func(t *testing.T) {
		assert.Equal(t, metadata, cluster.GetMetadata())
	})

	t.Run("GetReporterData returns correct value", func(t *testing.T) {
		assert.Equal(t, reporterData, cluster.GetReporterData())
	})

	t.Run("GetResourceData returns correct value", func(t *testing.T) {
		assert.Equal(t, resourceData, cluster.GetResourceData())
	})
}

func TestK8SClusterNilSafety(t *testing.T) {
	var cluster *K8SCluster

	t.Run("GetMetadata on nil returns nil", func(t *testing.T) {
		assert.Nil(t, cluster.GetMetadata())
	})

	t.Run("GetReporterData on nil returns nil", func(t *testing.T) {
		assert.Nil(t, cluster.GetReporterData())
	})

	t.Run("GetResourceData on nil returns nil", func(t *testing.T) {
		assert.Nil(t, cluster.GetResourceData())
	})
}
