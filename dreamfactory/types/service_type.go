package types

// ServiceTypesResponse represents a bulk response for services
type ServiceTypesResponse struct {
	Resource []ServiceType `json:"resource"`
}

// ServiceType represents an service in DreamFactory
type ServiceType struct {
	Name                 string         `json:"name"`
	Label                string         `json:"label"`
	Description          string         `json:"description"`
	Group                string         `json:"group"`
	Singleton            bool           `json:"singleton"`
	SubscriptionRequired bool           `json:"subscription_required"`
	ConfigSchema         []ConfigSchema `json:"config_schema"`
}

// ConfigSchema represents an service-specific configuration in DreamFactory
type ConfigSchema struct {
	Alias             string        `json:"alias"`
	Name              string        `json:"name"`
	Label             *string       `json:"label,omitempty"`
	Description       string        `json:"description"`
	Native            []interface{} `json:"native"`
	Type              string        `json:"type"` // *
	Length            *int          `json:"length,omitempty"`
	Precision         *int          `json:"precision,omitempty"`
	Scale             *int          `json:"scale,omitempty"`
	Default           *interface{}  `json:"default,omitempty"`
	Required          bool          `json:"required"`   // *
	AllowNull         bool          `json:"allow_null"` // *
	FixedLength       bool          `json:"fixed_length"`
	SupportsMultibyte bool          `json:"supports_multibyte"`
	IsPrimaryKey      bool          `json:"is_primary_key"`
	IsUnique          bool          `json:"is_unique"`
	IsForeignKey      bool          `json:"is_foreign_key"`
	RefTable          *string       `json:"ref_table,omitempty"`
	RefField          *string       `json:"ref_field,omitempty"`
	RefOnUpdate       *string       `json:"ref_on_update,omitempty"`
	RefOnDelete       *string       `json:"ref_on_delete,omitempty"`
	Picklist          *interface{}  `json:"picklist,omitempty"`
	Validation        *interface{}  `json:"validation,omitempty"`
	DbFunction        *interface{}  `json:"db_function,omitempty"`
	IsVirtual         bool          `json:"is_virtual"`
	IsAggregate       bool          `json:"is_aggregate"`
	Items             string        `json:"items,omitempty"`
	DbType            string        `json:"db_type,omitempty"`
	AutoIncrement     bool          `json:"auto_increment,omitempty"`
	IsIndex           bool          `json:"is_index,omitempty"`
	Values            []struct {    // *
		Name    string `json:"name"`
		Label   string `json:"label"`
		Default bool   `json:"default,omitempty"`
	} `json:"values,omitempty"`
}

func (s ServiceType) Validate(config map[string]interface{}) error {
	// TODO: Validar tipo de dado
	// TODO: Validar campos requeridos
	// TODO: Validar nulo
	// TODO: Para s.Type == "picklist", validar Values
	return nil // OK
}
