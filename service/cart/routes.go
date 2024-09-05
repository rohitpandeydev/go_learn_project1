package cart

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohitpandeydev/go_learn_project1/types"
	"github.com/rohitpandeydev/go_learn_project1/utils"
)

type Handler struct {
	store types.CartStore
}

func NewHandler(store types.CartStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/cart", h.handleGetCart).Methods(http.MethodGet)
	router.HandleFunc("/cart/add", h.handleAddToCart).Methods(http.MethodPost)
	router.HandleFunc("/cart/remove", h.handleRemoveFromCart).Methods(http.MethodPost)
	router.HandleFunc("/cart/update", h.handleUpdateCartItemQuantity).Methods(http.MethodPut)
}

func (h *Handler) handleGetCart(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement getting cart
	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleAddToCart(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement adding to cart
	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleRemoveFromCart(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement removing from cart
	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdateCartItemQuantity(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement updating cart item quantity
	utils.WriteJSON(w, http.StatusOK, nil)
}
