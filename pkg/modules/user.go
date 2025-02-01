package modules

type User struct {
	Id       int    `db:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"omitempty"`
	Password string `json:"password" binding:"required"`
}

// реализовать напрмую через бд и восстановление пароля через почту
// пароль хэшировать
