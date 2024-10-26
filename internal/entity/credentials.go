package entity

// Credentials - сущность секрета с парой логин/пароль.
// Создавать лучше через NewCredentials(), чтобы не забыть важные параметры, да и так короче.
type Credentials struct {
	Secret
	Login    string
	Password string
}

func NewCredentials(secretName, login, password string) *Credentials {
	return &Credentials{
		Secret: Secret{
			Name: secretName,
			Kind: KindCredentials,
		},
		Login:    login,
		Password: password,
	}
}
