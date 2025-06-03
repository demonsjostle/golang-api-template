package security

// HashPassword รับ raw password แล้วคืน hashed password หรือ error
func HashPassword(raw string) (string, error) {
    return bcryptHash(raw)
}

// ComparePassword ตรวจสอบ raw password กับ hashed password ว่าตรงกันหรือไม่
func ComparePassword(hashed, raw string) bool {
    return bcryptCompare(hashed, raw)
}
