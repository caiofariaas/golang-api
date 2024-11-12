package repository

import (
	"golang-api/configs"
	"golang-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct	{
}

// Construtor para criar uma nova instância de UserRepository.
// Retorna um ponteiro para uma nova estrutura UserRepository.

func NewUserRepository() *UserRepository{
	return &UserRepository{}
}

// Método para criar um novo usuário no banco de dados.

func (r * UserRepository) CreateUser(user *models.User) error {

		// Utiliza o método Create do ORM para inserir o usuário no banco.
		// Retorna um erro, caso ocorra.

	return database.DB.Create(user).Error
}

// Retorna um ponteiro para a entidade User e um erro, caso ocorra.

func (r * UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	// Utiliza o método Where para aplicar uma condição SQL e o método First
	// para pegar a primeira linha correspondente no banco.

	result := database.DB.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (r * UserRepository) DeleteUser(username string) (string, error) {
	var user models.User

	result := database.DB.Where("username = ?", username).Delete(&user)

	// Caso ocorra um erro durante a exclusão, retorna o erro.

	if result.Error != nil {
		return "", result.Error
	}

	// Caso nenhuma linha tenha sido afetada, significa que o usuário não foi encontrado.

	if result.RowsAffected == 0 {
		return "No user found with that username", nil
	}

	return "User deleted successfully", nil
}

// Método para verificar se a senha fornecida é válida.
// Compara a senha armazenada no banco (hash) com a senha fornecida.

func (r * UserRepository) CheckPassword(storedPassword, providedPassword string) bool {
	// Usa o método CompareHashAndPassword do pacote bcrypt para verificar as senhas.

	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	
	return err == nil
}