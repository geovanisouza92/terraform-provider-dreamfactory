package types

import (
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
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	IsActive    bool         `json:"is_active"`
	Access      []RoleAccess `json:"role_service_access_by_role_id"`
	Lookups     []RoleLookup `json:"role_lookup_by_role_id"`
}

// RoleAccess represents a DreamFactory role access configuration
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

// RoleFilter represents a DreamFactory role filter configuration
type RoleFilter struct {
	Name     string `json:"name"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// RoleLookup represents a DreamFactory role lookup
type RoleLookup struct {
	ID      int    `json:"id"`
	RoleID  int    `json:"role_id"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Private bool   `json:"private"`
}

const (
	// Get HTTP method
	Get byte = 1 << iota
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
	API byte = 1 << iota
	// Script define that request can came from Event scripts
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
		"api":    API,
		"script": Script,
	}
	roleRequestorByteToString = map[byte]string{
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
	if err != nil {
		return nil, err
	}

	r := Role{
		ID:          id,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	access := []RoleAccess{}
	for _, ra := range d.Get("access").(*schema.Set).List() {
		if ra == nil {
			continue
		}
		access = append(access, roleAccessFromMap(r.ID, ra.(map[string]interface{})))
	}
	r.Access = access

	lookups := []RoleLookup{}
	for _, rl := range d.Get("lookup").(*schema.Set).List() {
		if rl == nil {
			continue
		}
		lookups = append(lookups, roleLookupFromMap(r.ID, rl.(map[string]interface{})))
	}
	r.Lookups = lookups

	if s := d.Get("is_active").(string); s == "true" {
		r.IsActive = true
	}

	return &r, nil
}

// FillResourceData copy information from the Role to Terraform ResourceData
func (r *Role) FillResourceData(d *schema.ResourceData) error {
	access := &schema.Set{F: hashMapStringInterface}
	for _, ra := range r.Access {
		access.Add(ra.toMap())
	}

	lookups := &schema.Set{F: hashMapStringInterface}
	for _, rl := range r.Lookups {
		lookups.Add(rl.toMap())
	}

	d.Set("name", r.Name)
	d.Set("description", r.Description)
	d.Set("is_active", r.IsActive)
	d.Set("access", access)
	d.Set("lookup", lookups)

	return nil
}

func roleAccessFromMap(roleID int, m map[string]interface{}) RoleAccess {
	verbMask := byte(0)
	for _, s := range m["verb_mask"].(*schema.Set).List() {
		if s == nil {
			continue
		}
		verbMask += roleVerbsStringToByte[s.(string)]
	}

	requestorMask := byte(0)
	for _, s := range m["requestor_mask"].(*schema.Set).List() {
		if s == nil {
			continue
		}
		requestorMask += roleRequestorStringToByte[s.(string)]
	}

	filters := []RoleFilter{}
	for _, rf := range m["filters"].(*schema.Set).List() {
		if rf == nil {
			continue
		}
		filters = append(filters, roleFilterFromMap(rf.(map[string]interface{})))
	}

	return RoleAccess{
		ID:            getInt(m, "id"),
		RoleID:        roleID,
		ServiceID:     m["service_id"].(int),
		Component:     m["component"].(string),
		VerbMask:      int(verbMask),
		RequestorMask: int(requestorMask),
		Filters:       filters,
		FilterOp:      m["filter_op"].(string),
	}
}

func (ra *RoleAccess) toMap() map[string]interface{} {
	verbMask := &schema.Set{F: hashString}
	for b, s := range roleVerbsByteToString {
		if byte(ra.VerbMask)&b != 0 {
			verbMask.Add(s)
		}
	}

	requestorMask := &schema.Set{F: hashString}
	for b, s := range roleRequestorByteToString {
		if byte(ra.RequestorMask)&b != 0 {
			requestorMask.Add(s)
		}
	}

	filters := &schema.Set{F: hashMapStringString}
	for _, rf := range ra.Filters {
		filters.Add(rf.toMap())
	}

	return map[string]interface{}{
		"id":             ra.ID,
		"service_id":     ra.ServiceID,
		"component":      ra.Component,
		"verb_mask":      verbMask,
		"requestor_mask": requestorMask,
		"filters":        filters,
		"filter_op":      ra.FilterOp,
	}
}

func roleFilterFromMap(m map[string]interface{}) RoleFilter {
	return RoleFilter{
		Name:     m["name"].(string),
		Operator: m["operator"].(string),
		Value:    m["value"].(string),
	}
}

func (rf *RoleFilter) toMap() map[string]string {
	return map[string]string{
		"name":     rf.Name,
		"operator": rf.Operator,
		"value":    rf.Value,
	}
}

func roleLookupFromMap(roleID int, m map[string]interface{}) RoleLookup {
	return RoleLookup{
		ID:      getInt(m, "id"),
		RoleID:  roleID,
		Name:    m["name"].(string),
		Value:   m["value"].(string),
		Private: m["private"].(bool),
	}
}

func (r *RoleLookup) toMap() map[string]interface{} {
	return map[string]interface{}{
		"id":      r.ID,
		"name":    r.Name,
		"value":   r.Value,
		"private": r.Private,
	}
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
