package interfaces

import (
	"fmt"
	"net/http"

	"port-domain-service/pkg/utils"
)

func (api *PortDomainAPI) createPorts(w http.ResponseWriter, r *http.Request) {
	file, header, err := utils.ParseMultipartAndReturnFile(r)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	if err := api.service.Create(r.Context(), file); err != nil {
		utils.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Response(w, http.StatusOK, fmt.Sprintf("file: %s saved", header.Filename))
}
