package interfaces

import (
	"fmt"
	"net/http"
	"port-domain-service/pkg/utils"
)

func (api *PortDomainAPI) updatePorts(w http.ResponseWriter, r *http.Request) {
	file, header, err := utils.ParseMultipartAndReturnFile(r)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := api.service.Update(r.Context(), file); err != nil {
		utils.Response(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("file:", header.Filename)

	defer file.Close()
}
