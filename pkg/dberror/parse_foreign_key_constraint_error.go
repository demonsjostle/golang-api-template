package dberror

import (
	"errors"
	"fmt"
	"strings"
	"github.com/lib/pq"
)

// ParseForeignKeyConstraintError พยายามแปลง foreign key constraint error ให้แสดงชื่อ field ที่หายไป
func ParseForeignKeyConstraintError(err error) error {
	if err == nil {
		return nil
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23503" { // foreign_key_violation	
		return fmt.Errorf("related record not found")
	}

	// fallback สำหรับข้อความทั่วไป
	if strings.Contains(err.Error(), "violates foreign key constraint") {
		return fmt.Errorf("foreign key constraint violation")
	}

	return nil
}


