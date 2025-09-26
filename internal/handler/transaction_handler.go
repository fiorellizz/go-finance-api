package handler

import (
	"net/http"

	"github.com/fiorellizz/go-finance-api/internal/domain"
	"github.com/fiorellizz/go-finance-api/internal/service"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	svc *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{svc: s}
}

func (h *TransactionHandler) Create(c *gin.Context) {
    var req domain.Transaction
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    uid, ok := c.Get("userID")
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
        return
    }
    req.UserID = uid.(uint)

    if err := h.svc.Create(&req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save transaction"})
        return
    }

    c.JSON(http.StatusCreated, req)
}

func (h *TransactionHandler) List(c *gin.Context) {
	transactions, err := h.svc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transactions"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) ListByUser(c *gin.Context) {
    userID, exists := c.Get("userID") // vem do middleware JWT
    if !exists {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    transactions, err := h.svc.ListByUser(userID.(uint))
    if err != nil {
        c.JSON(500, gin.H{"error": "failed to fetch transactions"})
        return
    }

    c.JSON(200, transactions)
}

func (h *TransactionHandler) Update(c *gin.Context) {
    var req domain.Transaction
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id := c.Param("id")
    if err := h.svc.Update(id, &req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update transaction"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "transaction updated"})
}

func (h *TransactionHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    if err := h.svc.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete transaction"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "transaction deleted"})
}
