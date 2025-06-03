package dberror

import (
	"errors"
	"fmt"
	"strings"
)

// ParseNotFoundError พยายามแยกประเภท not found error จากข้อความ Ent
func ParseNotFoundError(err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()
	if strings.Contains(msg, "not found") {
		// ตัวอย่าง: "ent: role not found"
		parts := strings.Split(msg, " ")
		if len(parts) >= 2 {
			resource := parts[1] // เช่น "role"
			return fmt.Errorf("%s not found", resource)
		}
		return errors.New("resource not found")
	}

	return nil
}
