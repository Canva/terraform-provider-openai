package openai

import (
	"github.com/canva/terraform-provider-openai/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() func() *schema.Provider {
	return func() *schema.Provider {
		p := provider.Provider()
		return p
	}
}
