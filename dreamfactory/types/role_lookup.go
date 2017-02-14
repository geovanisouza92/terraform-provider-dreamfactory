package types

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func (r *Role) updateMissingLookups(d *schema.ResourceData) []map[string]interface{} {
	count := d.Get("lookup.#").(int)

	// Find for specific lookup ID
	find := func(id int) /* index */ int {
		for i := 0; i < count; i++ {
			if _id, ok := d.Get(fmt.Sprintf("lookup.%d.id", i)).(int); ok && id == _id {
				return i
			}
		}
		return -1
	}

	// Load existing items from ResourceData
	lookup := []map[string]interface{}{}
	for i := 0; i < count; i++ {
		lookup = append(lookup, map[string]interface{}{
			"id":      d.Get(fmt.Sprintf("lookup.%d.id", i)),
			"name":    d.Get(fmt.Sprintf("lookup.%d.name", i)),
			"value":   d.Get(fmt.Sprintf("lookup.%d.value", i)),
			"private": d.Get(fmt.Sprintf("lookup.%d.private", i)),
		})
	}

	// Loop through remote lookups, marking items that was removed locally
	for _, l := range r.Lookups {
		if i := find(l.ID); i == -1 {
			lookup = append(lookup, map[string]interface{}{
				"id":      l.ID * -1,
				"name":    l.Name,
				"value":   l.Value,
				"private": l.Private,
			})
		}
	}

	return lookup
}
