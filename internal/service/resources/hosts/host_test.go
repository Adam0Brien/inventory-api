package hosts_test

import (
	"testing"

	"github.com/google/uuid"
	pb "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta1/resources"
	"github.com/project-kessel/inventory-api/internal/authn/api"
	"github.com/project-kessel/inventory-api/internal/service/resources/hosts"
	"github.com/stretchr/testify/assert"
)

func TestHostFromCreateRequest(t *testing.T) {
	identity := &api.Identity{Principal: uuid.NewString()}

	req := &pb.CreateRhelHostRequest{
		RhelHost: &pb.RhelHost{
			Metadata: &pb.Metadata{
				Id: "test-host",
			},
			ReporterData: &pb.ReporterData{
				ReporterType:       3,
				ReporterInstanceId: "c2352f40-626a-4415-805b-a1dc9061afa",
			},
		},
	}

	resource, err := hosts.HostFromCreateRequest(req, identity)

	assert.NoError(t, err)
	assert.NotNil(t, resource)
	assert.Equal(t, "rhel_host", resource.ResourceType)
}
