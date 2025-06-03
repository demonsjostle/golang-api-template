package dberror 

import (
	"errors"
	"strings" 
	"github.com/lib/pq"
	"fmt"
)

// ParseUniqueConstraintError ถ้า err เป็น unique constraint violation จะ return error ข้อความสวยงาม เช่น "username already exists"
// ถ้าไม่ใช่ unique violation จะ return nil
func ParseUniqueConstraintError(err error) error {
	if err == nil {
		return nil
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		if pqErr.Code == "23505" { // unique_violation
			constraintName := pqErr.Constraint

			var column string
			if strings.HasSuffix(constraintName, "_key") {
				column = strings.TrimSuffix(constraintName, "_key")
				if idx := strings.LastIndex(column, "_"); idx != -1 {
					column = column[idx+1:]
				}
			} else {
				column = constraintName
			}

			return fmt.Errorf("%s already exists", column)
		}
	}

	// fallback: ตรวจสอบข้อความทั่วไป
	if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
		return fmt.Errorf("unique constraint violation")
	}

	return nil
}
