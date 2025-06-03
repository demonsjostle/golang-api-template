// cmd/seed/main.go
package main

import (
	"context"
	"fmt"
	"log"

	// "cpru-api/ent"	
	"cpru-api/ent/permission"
	"cpru-api/ent/role"
	// "cpru-api/ent/rolepermission"
	"cpru-api/ent/user"
	// "cpru-api/ent/userpermission"
	// "cpru-api/ent/userrole"

	"cpru-api/internal/infrastructure/config"
	"cpru-api/internal/infrastructure/db"

	"cpru-api/cmd/seed/data" // import seed data package

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// โหลด .env และ config
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	// สร้าง Ent client
	client, err := db.NewEntClient(dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// --- ลบข้อมูลทั้งหมด ก่อน seed ใหม่ ---
	_, _ = client.UserPermission.Delete().Exec(ctx)
	_, _ = client.RolePermission.Delete().Exec(ctx)
	_, _ = client.UserRole.Delete().Exec(ctx)
	_, _ = client.User.Delete().Exec(ctx)
	_, _ = client.Role.Delete().Exec(ctx)
	_, _ = client.Permission.Delete().Exec(ctx)
	

	// --- 2) Seed Roles ---
	for _, rn := range data.RoleSeed {
		_, err := client.Role.
			Create().
			SetName(rn).
			Save(ctx)
		if err != nil {
			log.Printf("⚠️ failed to create role (%s): %v", rn, err)
		}
	}

	// --- 3) Seed Permissions ---
	for _, rec := range data.PermissionSeed {
		_, err := client.Permission.
			Create().
			SetCode(rec.Code).
			SetNillableDescription(&rec.Description).
			Save(ctx)
		if err != nil {
			log.Printf("⚠️ failed to create permission (%s): %v", rec.Code, err)
		}
	}



	// ฟังก์ชันช่วย “หา User ID” จาก Username
	getUserID := func(username string) *int {
		u, err := client.User.Query().
			Where(user.UsernameEQ(username)).
			Only(ctx)
		if err != nil {
			return nil
		}
		return &u.ID
	}

	// ฟังก์ชันช่วย “หา Role ID” จาก RoleName
	getRoleID := func(roleName string) *int {
		r, err := client.Role.Query().
			Where(role.NameEQ(roleName)).
			Only(ctx)
		if err != nil {
			return nil
		}
		return &r.ID
	}

	// ฟังก์ชันช่วย “หา Permission ID” จาก PermissionCode
	getPermID := func(code string) *int {
		p, err := client.Permission.Query().
			Where(permission.CodeEQ(code)).
			Only(ctx)
		if err != nil {
			return nil
		}
		return &p.ID
	}

	// --- 4) Seed Users ---
	for _, rec := range data.UserSeed {
		hashed, err := bcrypt.GenerateFromPassword([]byte(rec.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("failed to hash password (%s): %v", rec.Username, err)
		}

		builder := client.User.
			Create().
			SetName(rec.Name).
			SetSurname(rec.Surname).
			SetUsername(rec.Username).
			SetEmail(rec.Email).
			SetPassword(string(hashed))

		// ถ้ามี phone
		if rec.Phone != nil {
			builder.SetNillablePhone(rec.Phone)
		}
		

		_, err = builder.Save(ctx)
		if err != nil {
			log.Printf("⚠️ failed to create user (%s): %v", rec.Username, err)
		}
	}

	// --- 5) Seed RolePermission ---
	for _, rec := range data.RolePermissionSeed {
		roleID := getRoleID(rec.RoleName)
		permID := getPermID(rec.PermissionCode)
		if roleID != nil && permID != nil {
			_, err := client.RolePermission.
				Create().
				SetRoleID(*roleID).
				SetPermissionID(*permID).
				Save(ctx)
			if err != nil {
				log.Printf("⚠️ failed to create role_permission (%s,%s): %v", rec.RoleName, rec.PermissionCode, err)
			}
		}
	}

	// --- 6) Seed UserRole ---
	for _, rec := range data.UserRoleSeed {
		userID := getUserID(rec.Username)
		roleID := getRoleID(rec.RoleName)
		if userID != nil && roleID != nil {
			_, err := client.UserRole.
				Create().
				SetUserID(*userID).
				SetRoleID(*roleID).
				Save(ctx)
			if err != nil {
				log.Printf("⚠️ failed to create user_role (%s,%s): %v", rec.Username, rec.RoleName, err)
			}
		}
	}

	// --- 7) Seed UserPermission ---
	for _, rec := range data.UserPermissionSeed {
		userID := getUserID(rec.Username)
		permID := getPermID(rec.PermissionCode)
		if userID != nil && permID != nil {
			_, err := client.UserPermission.
				Create().
				SetUserID(*userID).
				SetPermissionID(*permID).
				Save(ctx)
			if err != nil {
				log.Printf("⚠️ failed to create user_permission (%s,%s): %v", rec.Username, rec.PermissionCode, err)
			}
		}
	}

	log.Println("✅ Seeding completed")
}
