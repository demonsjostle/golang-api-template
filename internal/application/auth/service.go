package auth

import (
    "context"
    "errors"
    "time"

    "cpru-api/internal/domain/user"
    "cpru-api/pkg/security"
    "github.com/golang-jwt/jwt/v5"
)

// AuthService เก็บ userRepo + jwt config
type AuthService struct {
    userRepo     user.Repository
    jwtSecret    string
    jwtExpiresIn time.Duration
}

func NewAuthService(
    userRepo user.Repository,
    jwtSecret string,
    jwtExpiresIn time.Duration,
) *AuthService {
    return &AuthService{
        userRepo:     userRepo,
        jwtSecret:    jwtSecret,
        jwtExpiresIn: jwtExpiresIn,
    }
}

type RegisterInput struct {
    Name         string
    Surname      string
    Username     string
    Email        string
    Phone        *string
    Password     string
    DepartmentID *int
}

type LoginInput struct {
    Username string
    Password string
}

type AuthPayload struct {
    Token string
    User  *user.User
}

func (s *AuthService) Register(ctx context.Context, input RegisterInput) (*user.User, error) {
    // ตรวจซ้ำ username
    if exist, _ := s.userRepo.GetByUsername(ctx, input.Username); exist != nil {
        return nil, errors.New("username already exists")
    }
    // hash password
    hashed, err := security.HashPassword(input.Password)
    if err != nil {
        return nil, errors.New("failed to hash password")
    }
    newUser := &user.User{
        Name:         input.Name,
        Surname:      input.Surname,
        Username:     input.Username,
        Email:        input.Email,
        Phone:        input.Phone,
        Password:     string(hashed),
        DepartmentID: input.DepartmentID,
    }
    created, err := s.userRepo.Create(ctx, newUser)
    if err != nil {
        return nil, err
    }
    return created, nil
}

func (s *AuthService) Login(ctx context.Context, input LoginInput) (string, *user.User, error) {
    u, err := s.userRepo.GetByUsername(ctx, input.Username)
    if err != nil || u == nil {
        return "", nil, errors.New("invalid credentials")
    }
    if !security.ComparePassword(u.Password, input.Password) {
        return "", nil, errors.New("invalid credentials")
    }
    tokenStr, err := s.generateJWT(u.ID)
    if err != nil {
        return "", nil, errors.New("failed to generate token")
    }
    return tokenStr, u, nil
}

func (s *AuthService) generateJWT(userID int) (string, error) {
    claims := jwt.MapClaims{
        "sub": userID,
        "exp": time.Now().Add(s.jwtExpiresIn).Unix(),
        "iat": time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (int, error) {
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(s.jwtSecret), nil
    })
    if err != nil || !token.Valid {
        return 0, errors.New("invalid token")
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("invalid claims")
    }
    // ใน v5 ค่า sub จะ interpret เป็น float64 เสมอ
    sub, ok := claims["sub"].(float64)
    if !ok {
        return 0, errors.New("invalid subject")
    }
    return int(sub), nil
}
