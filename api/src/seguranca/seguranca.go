package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash receives a string and returns a hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compares a password with the hash
func VerificarSenha(hashedSenha, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSenha), []byte(senha))
}
