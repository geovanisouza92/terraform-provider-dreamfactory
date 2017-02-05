package types

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

// RolesRequest is a request object for Role CRUD
type RolesRequest struct {
	Resource []Role `json:"resource,omitempty"`
	IDs      []int  `json:"ids,omitempty"`
}

// RolesResponse is a response object for Role CRUD
type RolesResponse struct {
	Resource []Role   `json:"resource,omitempty"`
	Meta     Metadata `json:"meta,omitempty"`
}

// Role represents a DreamFactory role
type Role struct {
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	IsActive    bool         `json:"is_active,omitempty"`
	Access      []RoleAccess `json:"role_service_access_by_role_id"`
	Lookups     []RoleLookup `json:"role_lookup_by_role_id"`
}

// RoleAccess represents a DreamFactory role access configuration
type RoleAccess struct {
	ID            int          `json:"id,omitempty"`
	RoleID        *int         `json:"role_id"`
	ServiceID     int          `json:"service_id"`
	Component     string       `json:"component"`
	VerbMask      int          `json:"verb_mask"`
	RequestorMask int          `json:"requestor_mask"`
	Filters       []RoleFilter `json:"filters"`
	FilterOp      string       `json:"filter_op"`
}

// RoleFilter represents a DreamFactory role filter configuration
type RoleFilter struct {
	Name     string `json:"name"`
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

// RoleLookup represents a DreamFactory role lookup
type RoleLookup struct {
	ID      int    `json:"id,omitempty"`
	RoleID  *int   `json:"role_id"`
	Name    string `json:"name"`
	Value   string `json:"value,omitempty"`
	Private bool   `json:"private,omitempty"`
}

const (
	// Get HTTP method
	Get int = 1 << iota
	// Post HTTP method
	Post
	// Put HTTP method
	Put
	// Patch HTTP method
	Patch
	// Delete HTTP method
	Delete
)

const (
	// API define that request can came from HTTP API
	API int = 1 << iota
	// Script define that request can came from Event scripts
	Script
)

var (
	roleVerbsStringToInt = map[string]int{
		"get":    Get,
		"post":   Post,
		"put":    Put,
		"patch":  Patch,
		"delete": Delete,
	}
	roleVerbsIntToString = map[int]string{
		Get:    "get",
		Post:   "post",
		Put:    "put",
		Patch:  "patch",
		Delete: "delete",
	}
	roleRequestorStringToInt = map[string]int{
		"api":    API,
		"script": Script,
	}
	roleRequestorIntToString = map[int]string{
		API:    "api",
		Script: "script",
	}
	roleFilterOperators = map[string]bool{
		"=":           true,
		"!=":          true,
		">":           true,
		"<":           true,
		">=":          true,
		"<=":          true,
		"in":          true,
		"not in":      true,
		"starts with": true,
		"ends with":   true,
		"contains":    true,
		"is null":     true, // value always empty
		"is not null": true, // value always empty
	}
	roleFilterOp = map[string]bool{"AND": true, "OR": true}
)

// RoleFromResourceData creates a Role object from Terraform ResourceData
func RoleFromResourceData(d *schema.ResourceData) (*Role, error) {
	id, err := strconv.Atoi(d.Id())
	if d.Id() != "" && err != nil {
		return nil, err
	}

	r := Role{
		ID:          id,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		IsActive:    d.Get("is_active").(bool),
	}

	r.Access = make([]RoleAccess, 0)
	/*
		for i := 0; i < d.Get("access.#").(int); i++ {
			prefix := fmt.Sprintf("access.%d.", i)

			var filters []RoleFilter
			for j := 0; j < d.Get(prefix+"filter.#").(int); j++ {
				subprefix := fmt.Sprintf("%s.filter.%d.", prefix, j)
				filters = append(filters, RoleFilter{
					Name:     d.Get(subprefix + "name").(string),
					Operator: d.Get(subprefix + "operator").(string),
					Value:    d.Get(subprefix + "value").(string),
				})
			}

			verbMask := 0
			for _, vm := range d.Get(prefix + "verb_mask").(*schema.Set).List() {
				mask, ok := roleVerbsStringToInt[vm.(string)]
				if !ok {
					panic("Invalid verb mask: " + vm.(string))
				}
				verbMask += mask
			}

			requestorMask := 0
			for _, rm := range d.Get(prefix + "requestor_mask").(*schema.Set).List() {
				mask, ok := roleRequestorStringToInt[rm.(string)]
				if !ok {
					panic("Invalid requestor mask: " + rm.(string))
				}
				requestorMask += mask
			}

			roleID := &r.ID
			accessID := d.Get(prefix + "id").(int)
			if accessID < 0 {
				accessID = accessID * -1
				roleID = nil
			}

			r.Access = append(r.Access, RoleAccess{
				ID:            accessID,
				RoleID:        roleID,
				ServiceID:     d.Get(prefix + "service_id").(int),
				Component:     d.Get(prefix + "component").(string),
				VerbMask:      verbMask,
				RequestorMask: requestorMask,
				Filters:       filters,
				FilterOp:      d.Get(prefix + "filter_op").(string),
			})
		}
	*/

	r.Lookups = make([]RoleLookup, 0)
	for i := 0; i < d.Get("lookup.#").(int); i++ {
		prefix := fmt.Sprintf("lookup.%d.", i)

		roleID := &r.ID
		lookupID := d.Get(prefix + "id").(int)
		if lookupID < 0 {
			lookupID = lookupID * -1
			roleID = nil
		}

		r.Lookups = append(r.Lookups, RoleLookup{
			ID:      lookupID,
			RoleID:  roleID,
			Name:    d.Get(prefix + "name").(string),
			Value:   d.Get(prefix + "value").(string),
			Private: d.Get(prefix + "private").(bool),
		})
	}

	return &r, nil
}

// FillResourceData copy information from the Role to Terraform ResourceData
func (r *Role) FillResourceData(d *schema.ResourceData) error {
	access := []map[string]interface{}{}
	/*
		for _, a := range r.Access {
			access = append(access, map[string]interface{}{
				"id":             a.ID,
				"service_id":     a.ServiceID,
				"component":      a.Component,
				"verb_mask":      verbMaskToSlice(a.VerbMask),
				"requestor_mask": requestorMaskToSlice(a.RequestorMask),
				"filters":        filtersToSlice(a.Filters),
				"filter_op":      a.FilterOp,
			})
		}
	*/

	lookup := []map[string]interface{}{}
	for _, l := range r.Lookups {
		lookup = append(lookup, map[string]interface{}{
			"id":      l.ID,
			"name":    l.Name,
			"value":   l.Value,
			"private": l.Private,
		})
	}

	return firstError([]func() error{
		setOrError(d, "name", r.Name),
		setOrError(d, "description", r.Description),
		setOrError(d, "is_active", r.IsActive),
		setOrError(d, "access", access),
		setOrError(d, "lookup", lookup),
	})
}

// UpdateMissingResourceData checks for remote access/lookup that doesn't exist locally anymore, marking the role with missing ones, to allow remote removal
func (r *Role) UpdateMissingResourceData(d *schema.ResourceData) error {
	/*
		convertAccess := func(i int) map[string]interface{} {
			return map[string]interface{}{
				"id":             d.Get(fmt.Sprintf("access.%d.id", i)),
				"service_id":     d.Get(fmt.Sprintf("access.%d.service_id", i)),
				"component":      d.Get(fmt.Sprintf("access.%d.component", i)),
				"verb_mask":      d.Get(fmt.Sprintf("access.%d.verb_mask", i)),
				"requestor_mask": d.Get(fmt.Sprintf("access.%d.requestor_mask", i)),
				"filters":        d.Get(fmt.Sprintf("access.%d.filters", i)),
				"filter_op":      d.Get(fmt.Sprintf("access.%d.filter_op", i)),
			}
		}
		matchAccess := func(find func(int) int) []map[string]interface{} {
			list := []map[string]interface{}{}
			// Loop through remote items, marking items that was removed locally
			for _, a := range r.Access {
				if i := find(a.ID); i == -1 {
					list = append(list, map[string]interface{}{
						"id":             a.ID * -1,
						"service_id":     a.ServiceID,
						"component":      a.Component,
						"verb_mask":      verbMaskToSlice(a.VerbMask),
						"requestor_mask": requestorMaskToSlice(a.RequestorMask),
						"filters":        filtersToSlice(a.Filters),
						"filter_op":      a.FilterOp,
					})
				}
			}
			return list
		}
		access := walkAndMarkMissing("access", d, convertAccess, matchAccess)

		convertLookup := func(i int) map[string]interface{} {
			return map[string]interface{}{
				"id":      d.Get(fmt.Sprintf("lookup.%d.id", i)),
				"name":    d.Get(fmt.Sprintf("lookup.%d.name", i)),
				"value":   d.Get(fmt.Sprintf("lookup.%d.value", i)),
				"private": d.Get(fmt.Sprintf("lookup.%d.private", i)),
			}
		}
		matchLookup := func(find func(int) int) []map[string]interface{} {
			list := []map[string]interface{}{}
			// Loop through remote items, marking items that was removed locally
			for _, l := range r.Lookups {
				if i := find(l.ID); i == -1 {
					list = append(list, map[string]interface{}{
						"id":      l.ID * -1,
						"name":    l.Name,
						"value":   l.Value,
						"private": l.Private,
					})
				}
			}
			return list
		}
		lookup := walkAndMarkMissing("lookup", d, convertLookup, matchLookup)

		return firstError([]func() error{
			setOrError(d, "access", access),
			setOrError(d, "lookup", lookup),
		})
	*/

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
		prefix := fmt.Sprintf("lookup.%d.", i)
		lookup = append(lookup, map[string]interface{}{
			"id":      d.Get(prefix + "id"),
			"name":    d.Get(prefix + "name"),
			"value":   d.Get(prefix + "value"),
			"private": d.Get(prefix + "private"),
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

	return firstError([]func() error{
		setOrError(d, "lookup", lookup),
	})
}

// IsValidOperator validates if a given operator is valid for lookups
func IsValidOperator(o string) bool {
	_, ok := roleFilterOperators[o]
	return ok
}

// IsValidOp validates if a given logical operator is valid for a combination
// of filters
func IsValidOp(o string) bool {
	_, ok := roleFilterOp[o]
	return ok
}

func verbMaskToSlice(verbMask int) []string {
	var res []string
	for verbInt, verbString := range roleVerbsIntToString {
		if verbMask&verbInt != 0 {
			res = append(res, verbString)
		}
	}
	return res
}

func requestorMaskToSlice(requestorMask int) []string {
	var res []string
	for requestorInt, requestorString := range roleRequestorIntToString {
		if requestorMask&requestorInt != 0 {
			res = append(res, requestorString)
		}
	}
	return res
}

func filtersToSlice(filters []RoleFilter) []map[string]interface{} {
	var res []map[string]interface{}
	for _, f := range filters {
		res = append(res, map[string]interface{}{
			"name":     f.Name,
			"operator": f.Operator,
			"value":    f.Value,
		})
	}
	return res
}
