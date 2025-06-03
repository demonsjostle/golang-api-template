package security

import "golang.org/x/crypto/bcrypt"

// bcryptHash จะทำหน้าที่แฮชรหัสผ่านแบบ bcrypt
func bcryptHash(raw string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashed), nil
}

// bcryptCompare จะเปรียบเทียบว่ารหัสผ่านดิบ (raw) ตรงกับ hashed หรือไม่
func bcryptCompare(hashed, raw string) bool {
    if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw)); err != nil {
        return false
    }
    return true
}
