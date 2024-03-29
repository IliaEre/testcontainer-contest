package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	s "testcontainer-contest/app/usecase/portfolio"
	mapper "testcontainer-contest/pkg"
)

func HandleGetPortfolio(service s.PortfolioService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("ID")
		if name == "" {
			http.Error(w, "Missing ID parameter", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		portfolio, err := service.FindByID(ctx, name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error finding portfolio: %v", err), http.StatusInternalServerError)
			return
		}

		res := mapper.MapPortfolioToResult(portfolio)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
