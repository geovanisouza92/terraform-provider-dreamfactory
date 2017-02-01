package types

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func setOrError(d *schema.ResourceData, key string, value interface{}) func() error {
	return func() error {
		return d.Set(key, value)
	}
}

func firstError(ops []func() error) error {
	for _, op := range ops {
		if err := op(); err != nil {
			return err
		}
	}
	return nil
}
