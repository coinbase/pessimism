package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/base-org/pessimism/internal/api/models"
	"github.com/base-org/pessimism/internal/logging"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

func renderHeuristicResponse(w http.ResponseWriter, r *http.Request,
	ir *models.InvResponse) {
	w.WriteHeader(ir.Code)
	render.JSON(w, r, ir)
}

// RunHeuristic ... Handle heuristic run request
func (ph *PessimismHandler) RunHeuristic(w http.ResponseWriter, r *http.Request) {
	var body *models.InvRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logging.WithContext(ph.ctx).
			Error("Could not unmarshal request", zap.Error(err))

		renderHeuristicResponse(w, r,
			models.NewInvUnmarshalErrResp())
		return
	}

	sUUID, err := ph.service.ProcessHeuristicRequest(body)
	if err != nil {
		logging.WithContext(ph.ctx).
			Error("Could not process heuristic request", zap.Error(err))

		renderHeuristicResponse(w, r, models.NewInvNoProcessInvResp())
		return
	}

	renderHeuristicResponse(w, r, models.NewInvAcceptedResp(sUUID))
}
