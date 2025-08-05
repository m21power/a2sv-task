package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"task/domain"
	"task/usecases"
	"task/utils"
	"time"

	"github.com/gorilla/mux"
)

type CMSPHandler struct {
	cmspUseCase usecases.CMSPUseCase
}

func NewCMSPHandler(useCase usecases.CMSPUseCase) *CMSPHandler {
	return &CMSPHandler{
		cmspUseCase: useCase,
	}
}

// GetAllCMSPs godoc
// @Summary      List all CMSPs
// @Description  Get all Capital Market Service Providers, filter by type or license date
// @Tags         cmsp
// @Accept       json
// @Produce      json
// @Param        type            query     string  false  "CMSP type"
// @Param        licensedBefore  query     string  false  "Licensed before (YYYY-MM-DD)"
// @Param        licensedAfter   query     string  false  "Licensed after (YYYY-MM-DD)"
// @Success      200   {array}   domain.CMSP
// @Failure      400   {object}  utils.StandardResponse
// @Failure      500   {object}  utils.StandardResponse
// @Router       /api/v1/cmsp [get]
func (h *CMSPHandler) GetAllCMSPs(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02"

	var licensedBefore, licensedAfter *time.Time

	if str := r.URL.Query().Get("licensedBefore"); str != "" {
		t, err := time.Parse(layout, str)
		if err != nil {
			utils.WriteError(w, fmt.Errorf("invalid licensedBefore format. Use YYYY-MM-DD"), http.StatusBadRequest)
			return
		}
		licensedBefore = &t
	}

	if str := r.URL.Query().Get("licensedAfter"); str != "" {
		t, err := time.Parse(layout, str)
		if err != nil {
			utils.WriteError(w, fmt.Errorf("invalid licensedAfter format. Use YYYY-MM-DD"), http.StatusBadRequest)
			return
		}
		licensedAfter = &t
	}
	filter := domain.FilterModel{
		Type:           r.URL.Query().Get("type"),
		LicensedBefore: licensedBefore,
		LicensedAfter:  licensedAfter,
	}
	cmsp, err := h.cmspUseCase.GetAllCMSPs(filter)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, cmsp)
}

// GetCMSPByID godoc
// @Summary      Get CMSP by ID
// @Description  Get a Capital Market Service Provider by its ID
// @Tags         cmsp
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "CMSP ID"
// @Success      200  {object}  domain.CMSP
// @Failure      400  {object}  utils.StandardResponse
// @Failure      404  {object}  utils.StandardResponse
// @Router       /api/v1/cmsp/{id} [get]
func (h *CMSPHandler) GetCMSPByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	cmsp, err := h.cmspUseCase.GetCMSPByID(id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJSON(w, http.StatusOK, cmsp)
}
