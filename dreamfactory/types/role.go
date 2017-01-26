package types

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type RolesRequest struct {
	Resource []Role `json:"resource,omitempty"`
	IDs      []int  `json:"ids,omitempty"`
}

type RolesResponse struct {
	Resource []Role   `json:"resource,omitempty"`
	Meta     Metadata `json:"meta,omitempty"`
}

type Role struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	IsActive    bool         `json:"is_active"`
	Access      []RoleAccess `json:"role_service_access_by_role_id"`
	Lookups     []RoleLookup `json:"role_lookup_by_role_id"`
}

type RoleAccess struct {
	ID            int          `json:"id"`
	RoleID        int          `json:"role_id"`
	ServiceID     int          `json:"service_id"`
	Component     string       `json:"component"`
	VerbMask      int          `json:"verb_mask"`
	RequestorMask int          `json:"requestor_mask"`
	Filters       []RoleFilter `json:"filters"`
	FilterOp      string       `json:"filter_op"`
}

type RoleFilter struct {
	Name     string `json:"name"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type RoleLookup struct {
	ID      int    `json:"id"`
	RoleID  int    `json:"role_id"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Private bool   `json:"private"`
}

const (
	Get byte = 1 << iota
	Post
	Put
	Patch
	Delete
)

const (
	Api byte = 1 << iota
	Script
)

var (
	roleVerbsStringToByte = map[string]byte{
		"get":    Get,
		"post":   Post,
		"put":    Put,
		"patch":  Patch,
		"delete": Delete,
	}
	roleVerbsByteToString = map[byte]string{
		Get:    "get",
		Post:   "post",
		Put:    "put",
		Patch:  "patch",
		Delete: "delete",
	}
	roleRequestorStringToByte = map[string]byte{
		"api":    Api,
		"script": Script,
	}
	roleRequestorByteToString = map[byte]string{
		Api:    "api",
		Script: "script",
	}
	roleFilterOperators = []string{
		"=",
		"!=",
		">",
		"<",
		">=",
		"<=",
		"in",
		"not in",
		"starts with",
		"ends with",
		"contains",
		"is null",     // value always empty
		"is not null", // value always empty
	}
	roleFilterOp = []string{"AND", "OR"}
)

func RoleFromResourceData(d *schema.ResourceData) Role {
	accesses := []RoleAccess{}
	for _, ra := range d.Get("access").(*schema.Set).List() {
		accesses = append(accesses, roleAccessFromMap(ra.(map[string]interface{})))
	}

	lookups := []RoleLookup{}
	for _, rl := range d.Get("lookups").(*schema.Set).List() {
		lookups = append(lookups, roleLookupFromMap(rl.(map[string]interface{})))
	}

	return Role{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		IsActive:    d.Get("is_active").(bool),
		Access:      accesses,
		Lookups:     lookups,
	}
}

func (r *Role) FillResourceData(d *schema.ResourceData) error {
	accesses := []map[string]interface{}{}
	for _, ra := range r.Access {
		accesses = append(accesses, ra.toMap())
	}

	lookups := []map[string]interface{}{}
	for _, rl := range r.Lookups {
		lookups = append(lookups, rl.toMap())
	}

	d.Set("name", r.Name)
	d.Set("description", r.Description)
	d.Set("is_active", r.IsActive)
	d.Set("access", accesses)
	d.Set("lookup", lookups)

	return nil
}

func roleAccessFromMap(m map[string]interface{}) RoleAccess {
	verb_mask := byte(0)
	for _, s := range m["verb_mask"].(*schema.Set).List() {
		verb_mask &= roleVerbsStringToByte[s.(string)]
	}

	requestor_mask := byte(0)
	for _, s := range m["requestor_mask"].(*schema.Set).List() {
		requestor_mask &= roleRequestorStringToByte[s.(string)]
	}

	filters := []RoleFilter{}
	for _, rf := range m["filters"].(*schema.Set).List() {
		filters = append(filters, roleFilterFromMap(rf.(map[string]string)))
	}

	return RoleAccess{
		ID:            m["id"].(int),
		RoleID:        m["role_id"].(int),
		ServiceID:     m["service_id"].(int),
		Component:     m["component"].(string),
		VerbMask:      int(verb_mask),
		RequestorMask: int(requestor_mask),
		Filters:       filters,
		FilterOp:      m["filter_op"].(string),
	}
}

func (ra *RoleAccess) toMap() map[string]interface{} {
	verb_mask := []string{}
	for b, s := range roleVerbsByteToString {
		if byte(ra.VerbMask)&b != 0 {
			verb_mask = append(verb_mask, s)
		}
	}

	requestor_mask := []string{}
	for b, s := range roleRequestorByteToString {
		if byte(ra.RequestorMask)&b != 0 {
			requestor_mask = append(requestor_mask, s)
		}
	}

	filters := []map[string]string{}
	for _, rf := range ra.Filters {
		filters = append(filters, rf.toMap())
	}

	return map[string]interface{}{
		"id":             ra.ID,
		"role_id":        ra.RoleID,
		"service_id":     ra.ServiceID,
		"component":      ra.Component,
		"verb_mask":      verb_mask,
		"requestor_mask": requestor_mask,
		"filters":        filters,
		"filter_op":      ra.FilterOp,
	}
}

func roleFilterFromMap(m map[string]string) RoleFilter {
	return RoleFilter{
		Name:     m["name"],
		Operator: m["operator"],
		Value:    m["value"],
	}
}

func (rf *RoleFilter) toMap() map[string]string {
	return map[string]string{
		"name":     rf.Name,
		"operator": rf.Operator,
		"value":    rf.Value,
	}
}

func roleLookupFromMap(m map[string]interface{}) RoleLookup {
	return RoleLookup{
		ID:      m["id"].(int),
		RoleID:  m["role_id"].(int),
		Name:    m["name"].(string),
		Value:   m["value"].(string),
		Private: m["private"].(bool),
	}
}

func (r *RoleLookup) toMap() map[string]interface{} {
	return map[string]interface{}{
		"id":      r.ID,
		"role_id": r.RoleID,
		"name":    r.Name,
		"value":   r.Value,
		"private": r.Private,
	}
}
