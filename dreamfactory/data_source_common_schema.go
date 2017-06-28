package dreamfactory

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

var (
	errFilter = errors.New("You need to provide a filter on datasource")
)

func dataSourceFiltersSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func buildDataSourceFilters(set *schema.Set) []string {
	filter := []string{}
	for _, v := range set.List() {
		filter = append(filter, v.(string))
	}
	return filter
}
