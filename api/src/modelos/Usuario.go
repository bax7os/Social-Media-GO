package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario represents a user
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEM time.Time `json:"criadoEm"`
}

// Preparar goes to check if the user data is valid
func (usuario *Usuario) Preparar(etapa string) error {

	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}
func (usuario *Usuario) validar(etapa string) error {

	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)
}
