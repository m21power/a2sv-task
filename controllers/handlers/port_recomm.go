package handlers

import (
	"encoding/json"
	"net/http"
	"task/domain"
	"task/usecases"
	"task/utils"
)

type PortRecommHandler struct {
	portRecommUseCase usecases.PortfolioRecommendationUseCase
}

func NewPortRecommHandler(useCase usecases.PortfolioRecommendationUseCase) *PortRecommHandler {
	return &PortRecommHandler{
		portRecommUseCase: useCase,
	}
}

// RecommendPortfolio godoc
// @Summary      Recommend a portfolio
// @Description  Get a recommended portfolio allocation based on user profile
// @Tags         portfolio
// @Accept       json
// @Produce      json
// @Param        profile  body      domain.PortfolioRequest  true  "Portfolio request body"
// @Success      200      {object}  domain.PortfolioResponse
// @Failure      500      {object}  utils.StandardResponse
// @Router       /api/v1/portfolio/recommendation [post]
func (h *PortRecommHandler) RecommendPortfolio(
	w http.ResponseWriter, r *http.Request) {
	profile := domain.PortfolioRequest{}
	json.NewDecoder(r.Body).Decode(&profile)
	response, err := h.portRecommUseCase.RecommendPortfolio(profile)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, response)
}
