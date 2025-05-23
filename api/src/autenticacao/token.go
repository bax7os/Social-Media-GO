package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken creates a JWT token
func CriarToken(id uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey)) //secret
}

// ValidarToken verifies if the token is valid
func ValidarToken(r *http.Request) error {
	token := extrairToken(r)
	tokenValido, erro := jwt.Parse(token, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := tokenValido.Claims.(jwt.MapClaims); ok || tokenValido.Valid {
		return nil
	}
	return errors.New("token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado %v", token.Header["alg"])
	}

	return []byte(config.SecretKey), nil
}
