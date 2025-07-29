package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"task/domain"
	"task/usecases"
	"task/utils"
	"time"
)

type BondHandler struct {
	bondUseCase usecases.BondUseCase
}

func NewBondHandler(useCase usecases.BondUseCase) *BondHandler {
	return &BondHandler{
		bondUseCase: useCase,
	}
}

func (h *BondHandler) GetBonds(w http.ResponseWriter, r *http.Request) {
	minCouponStr := r.URL.Query().Get("minCoupon")
	maxCouponStr := r.URL.Query().Get("maxCoupon")

	var minCoupon, maxCoupon *float64

	if minCouponStr != "" {
		val, err := strconv.ParseFloat(minCouponStr, 64)
		if err != nil {
			utils.WriteError(w, err, http.StatusBadRequest)
			return
		}
		minCoupon = &val
	}

	if maxCouponStr != "" {
		val, err := strconv.ParseFloat(maxCouponStr, 64)
		if err != nil {
			utils.WriteError(w, err, http.StatusBadRequest)
			return
		}
		maxCoupon = &val
	}
	layout := "2006-01-02"

	var maturityBefore, maturityAfter *time.Time

	if str := r.URL.Query().Get("maturityBefore"); str != "" {
		t, err := time.Parse(layout, str)
		if err != nil {
			utils.WriteError(w, fmt.Errorf("Invalid maturityBefore format. Use YYYY-MM-DD."), http.StatusBadRequest)
			return
		}
		maturityBefore = &t
	}

	if str := r.URL.Query().Get("maturityAfter"); str != "" {
		t, err := time.Parse(layout, str)
		if err != nil {
			utils.WriteError(w, fmt.Errorf("Invalid maturityAfter format. Use YYYY-MM-DD."), http.StatusBadRequest)
			return
		}
		maturityAfter = &t
	}
	search := domain.BondSearchModel{
		MinCoupon:      minCoupon,
		MaturityBefore: maturityBefore,
		MaturityAfter:  maturityAfter,
		MaxCoupon:      maxCoupon,
	}
	bonds, err := h.bondUseCase.GetBonds(search)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, bonds)
}
