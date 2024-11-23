package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/usecase"
)

type ListOrderHandler struct {
	ListOrderUseCase usecase.ListOrderUseCaseInterface
}

func NewListOrderHandler(listOrderUseCase usecase.ListOrderUseCaseInterface) *ListOrderHandler {
	return &ListOrderHandler{
		ListOrderUseCase: listOrderUseCase,
	}
}

func (h *ListOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	// Executa o caso de uso para obter as ordens
	orders, err := h.ListOrderUseCase.Execute()
	if err != nil {
		// Retorna erro 500 caso ocorra algum problema
		http.Error(w, fmt.Sprintf("Erro ao listar ordens: %v", err), http.StatusInternalServerError)
		return
	}

	// Serializa a resposta em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Escreve o JSON no corpo da resposta
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao serializar resposta: %v", err), http.StatusInternalServerError)
		return
	}
}
