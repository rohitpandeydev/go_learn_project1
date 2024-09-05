package cart

import (
	"database/sql"

	"github.com/rohitpandeydev/go_learn_project1/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetCart(userID int) (*types.Cart, error) {
	// TODO: Implement fetching cart from database
	return nil, nil
}

func (s *Store) AddToCart(userID int, productID int, quantity int) error {
	// TODO: Implement adding item to cart in database
	return nil
}

func (s *Store) RemoveFromCart(userID int, productID int) error {
	// TODO: Implement removing item from cart in database
	return nil
}

func (s *Store) UpdateCartItemQuantity(userID int, productID int, quantity int) error {
	// TODO: Implement updating cart item quantity in database
	return nil
}
