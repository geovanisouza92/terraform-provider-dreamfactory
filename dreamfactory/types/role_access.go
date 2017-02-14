package types

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func (r *Role) updateMissingAccesses(d *schema.ResourceData) []map[string]interface{} {
	count := d.Get("access.#").(int)

	// Find for specific access ID
	find := func(id int) /* index */ int {
		for i := 0; i < count; i++ {
			if _id, ok := d.Get(fmt.Sprintf("access.%d.id", i)).(int); ok && id == _id {
				return i
			}
		}
		return -1
	}

	// Load existing items from ResourceData
	access := []map[string]interface{}{}
	for i := 0; i < count; i++ {
		prefix := fmt.Sprintf("access.%d.", i)

		m := map[string]interface{}{
			"id":             d.Get(prefix + "id"),
			"service_id":     d.Get(prefix + "service_id"),
			"component":      d.Get(prefix + "component"),
			"verb_mask":      d.Get(prefix + "verb_mask").(*schema.Set),
			"requestor_mask": d.Get(prefix + "requestor_mask").(*schema.Set),
			"filter_op":      d.Get(prefix + "filter_op"),
		}
		if f := filtersToSlice(d, prefix); len(f) > 0 {
			m["filter"] = f
		}

		access = append(access, m)
	}

	// Loop through remote accesses, marking items that was removed locally
	for _, a := range r.Access {
		if i := find(a.ID); i == -1 {
			access = append(access, map[string]interface{}{
				"id":             a.ID * -1,
				"service_id":     a.ServiceID,
				"component":      a.Component,
				"verb_mask":      verbMaskToSet(a.VerbMask),
				"requestor_mask": requestorMaskToSet(a.RequestorMask),
				"filter":         filtersToMap(a.Filters),
				"filter_op":      a.FilterOp,
			})
		}
	}

	return access
}

func verbMaskToSet(verbMask int) *schema.Set {
	var res []interface{}
	for verbInt, verbString := range roleVerbsIntToString {
		if verbMask&verbInt != 0 {
			res = append(res, verbString)
		}
	}
	return schema.NewSet(VerbMaskFunc, res)
}

func verbMaskToInt(list *schema.Set) (verbMask int) {
	for _, vm := range list.List() {
		mask, ok := roleVerbsStringToInt[vm.(string)]
		if !ok {
			panic("Invalid verb mask: " + vm.(string))
		}
		verbMask += mask
	}
	return
}

func requestorMaskToSet(requestorMask int) *schema.Set {
	var res []interface{}
	for requestorInt, requestorString := range roleRequestorIntToString {
		if requestorMask&requestorInt != 0 {
			res = append(res, requestorString)
		}
	}
	return schema.NewSet(RequestorMaskFunc, res)
}

func requestorMaskToInt(list *schema.Set) (requestorMask int) {
	for _, rm := range list.List() {
		mask, ok := roleRequestorStringToInt[rm.(string)]
		if !ok {
			panic("Invalid requestor mask: " + rm.(string))
		}
		requestorMask += mask
	}
	return
}

func filtersToMap(filters []RoleFilter) (res []map[string]interface{}) {
	for _, f := range filters {
		res = append(res, map[string]interface{}{
			"name":     f.Name,
			"operator": f.Operator,
			"value":    f.Value,
		})
	}
	return
}

func filtersToSlice(d *schema.ResourceData, prefix string) (filters []RoleFilter) {
	for i := 0; i < d.Get(prefix+"filter.#").(int); i++ {
		subprefix := fmt.Sprintf("%sfilter.%d.", prefix, i)
		filters = append(filters, RoleFilter{
			Name:     d.Get(subprefix + "name").(string),
			Operator: d.Get(subprefix + "operator").(string),
			Value:    d.Get(subprefix + "value").(string),
		})
	}
	return
}
