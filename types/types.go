package types

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserbyId(ID int) (*User, error)
	CreateUser(User) error
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	CreateProduct(cp CreateProduct) error
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=120"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type CreateProduct struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type Order struct {
	ID        int       `json:"id" validate:"required"`
	UserID    int       `json:"userid" validate:"required"`
	Total     float64   `json:"total" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
}

type orderItem struct {
	ID        int       `json:"id" validate:"required"`
	OrderId   int       `json:"orderId" validate:"required"`
	ProductId int       `json:"productId" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
}

type CartStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(orderItem) error
}
