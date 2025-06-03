package ptr

// StringToPointer แปลง string เป็น *string ถ้าเป็น empty string จะคืนค่า nil
func StringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// IntToPointer แปลง int เป็น *int ถ้าเป็น 0 จะคืนค่า nil
func IntToPointer(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}


func StringPtrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func IntPtrToInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

