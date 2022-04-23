package auth

import (
	"BitginHomework/config"
	"BitginHomework/database"
	"BitginHomework/model"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// sign user by email and plain password
func SignUser(ctx context.Context, email string, password_plain string) (string, error) {
	user, err := model.UserByEmail(ctx, database.GetDB(), email)
	if err != nil {
		// fetch user error
		log.Println("dsadaeeeeeeeeeeeeeeeeeeeeeeeee")
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password_plain)); err != nil {
		return "", errors.New("miss compare")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
	})

	tokenString, err := token.SignedString([]byte(config.HASH_SECRET))
	if err != nil {
		// sign error
		return "", err
	}

	return tokenString, nil
}

// valid signed token string
func ValidToken(ctx context.Context, tokenString string) (*model.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.HASH_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, ok := claims["email"]; ok {
			user, err := model.UserByEmail(ctx, database.GetDB(), claims["email"].(string))
			if err != nil {
				return nil, err
			}

			return user, nil
		}
	}

	return nil, errors.New("token faild")
}
