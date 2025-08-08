package handlers

import (
	"net/http"
	"task/usecases"
	"task/utils"
)

type EconomicIndicatorHandler struct {
	economicIndicatorUseCase usecases.EconomicIndicatorUseCase
}

func NewEconomicIndicatorHandler(useCase usecases.EconomicIndicatorUseCase) *EconomicIndicatorHandler {
	return &EconomicIndicatorHandler{
		economicIndicatorUseCase: useCase,
	}
}

// GetEconomicIndicators godoc
// @Summary      Get economic indicators
// @Description  Retrieve key economic indicators for Ethiopia
// @Tags         economic-indicators
// @Accept       json
// @Produce      json
// @Success      200   {object}  domain.APIResponse
// @Failure      500   {object}  utils.StandardResponse
// @Router       /api/v1/economic-indicators [get]
func (h *EconomicIndicatorHandler) GetEconomicIndicators(w http.ResponseWriter, r *http.Request) {
	response, err := h.economicIndicatorUseCase.GetEconomicIndicators()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, response)
}
