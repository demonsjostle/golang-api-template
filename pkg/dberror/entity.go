package dberror

import (
  "errors"
)

var ErrUniqueViolation = errors.New("unique constraint violation")

// UniqueViolationInfo เก็บข้อมูลว่าคือข้อผิดพลาดชนิด unique violation และชื่อคอลัมน์/constraint ที่ชนกัน
type UniqueViolationInfo struct {
	Column string
}

// NotFoundErrorInfo ช่วยเก็บข้อมูลว่า object ไหน not found
type NotFoundErrorInfo struct {
	Resource string
}
