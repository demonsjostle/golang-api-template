package dberror

import (
	"errors"
	"strings" 
	"github.com/lib/pq"
)



func IsUniqueConstraintError(err error) (bool, *UniqueViolationInfo) {
	if err == nil {
		return false, nil
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		if pqErr.Code == "23505" { // unique_violation
			// ตัวอย่าง: constraint name อาจจะเป็น "users_username_key" หรือ "users_email_key"
			constraintName := pqErr.Constraint

			// พยายามดึงชื่อ column จาก constraint name (ตาม naming convention)
			var column string
			if strings.HasSuffix(constraintName, "_key") {
				column = strings.TrimSuffix(constraintName, "_key")
				// ตัวอย่าง: users_username_key → users_username → เอา username (ตัดชื่อ table ออก)
				if idx := strings.LastIndex(column, "_"); idx != -1 {
					column = column[idx+1:]
				}
			} else {
				// กรณีไม่ตรง pattern อาจเอา constraint ทั้งหมดเลย
				column = constraintName
			}

			return true, &UniqueViolationInfo{
				Column: column,
			}
		}
	}

	// ตรวจสอบข้อความ error ทั่วไป (fallback)
	if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
		return true, nil // แต่ไม่มีข้อมูล column
	}

	return false, nil
}


func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
  return strings.Contains(err.Error(), "not found")
}


