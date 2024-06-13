package repository

import (
	"Ayala-Crea/ResepBe/model"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *model.Users) error {
	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Jika id_role tidak diisi, atur nilainya ke 2 (atau nilai default yang diinginkan)
	if user.IdRole == 0 {
		user.IdRole = 2 // atau nilai default yang diinginkan
	}

	// Simpan user ke database
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByUsername(db *gorm.DB, username string) (*model.Users, error) {
	var user model.Users
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*model.Users, error) {
	var user model.Users
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserById(db *gorm.DB, userID uint) (*model.Users, error) {
	var user model.Users
	result := db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

// func GetRoleById(db *gorm.DB, roleID int) (*model.Roles, error) {
// 	var role model.Roles
// 	result := db.First(&role, roleID)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, nil
// 		}
// 		return nil, result.Error
// 	}
// 	return &role, nil
// }

// GetUserByRoleId retrieves a user by their role ID. It returns user data or an error.
func GetUserByRoleId(db *gorm.DB, IdRole int) (*model.Users, error) {
	var user model.Users
	// Fetch only the necessary fields from the Users table based on the role ID
	result := db.Select("id_user", "id_role", "nama", "username", "email").Where("id_role = ?", IdRole).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// No user found with the specified role ID, return nil without an error
			return nil, nil
		}
		// Return with an error if it's other than not found
		return nil, result.Error
	}

	// Return user data successfully
	return &user, nil
}

func GenerateToken(user *model.Users) (string, error) {
	claims := &model.JWTClaims{
		IdUser: user.IdUser,
		IdRole: user.IdRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // Token berlaku selama 1 jam
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
