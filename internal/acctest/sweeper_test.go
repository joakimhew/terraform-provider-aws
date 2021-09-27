package acctest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func TestMain(m *testing.M) {
	sweeperAwsClients = make(map[string]interface{})
	resource.TestMain(m)
}
