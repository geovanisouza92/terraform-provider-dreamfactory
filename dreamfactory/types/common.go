package types

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
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

func getInt(m map[string]interface{}, k string) int {
	if m[k] != nil {
		return m[k].(int)
	}
	return 0
}

func hashMapStringInterface(i interface{}) int {
	m := i.(map[string]interface{})
	var buf bytes.Buffer
	for k, v := range m {
		buf.WriteString(fmt.Sprintf("%s-", k))
		buf.WriteString(fmt.Sprintf("%v-", v))
	}
	return hashcode.String(buf.String())
}

func hashMapStringString(i interface{}) int {
	m := i.(map[string]string)
	var buf bytes.Buffer
	for k, v := range m {
		buf.WriteString(fmt.Sprintf("%s-", k))
		buf.WriteString(fmt.Sprintf("%s-", v))
	}
	return hashcode.String(buf.String())
}

func hashSliceString(i interface{}) int {
	s := i.([]string)
	var buf bytes.Buffer
	for _, it := range s {
		buf.WriteString(fmt.Sprintf("%s-", it))
	}
	return hashcode.String(buf.String())
}

func hashString(i interface{}) int {
	return hashcode.String(i.(string))
}
