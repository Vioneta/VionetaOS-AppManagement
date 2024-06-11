package service

import (
	"github.com/Vioneta/VionetaOS-AppManagement/codegen"
	"github.com/Vioneta/VionetaOS-AppManagement/common"
	"github.com/Vioneta/VionetaOS-Common/utils/logger"
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

type App types.ServiceConfig

func (a *App) StoreInfo() (codegen.AppStoreInfo, error) {
	var storeInfo codegen.AppStoreInfo

	ex, ok := a.Extensions[common.ComposeExtensionNameXVionetaOS]
	if !ok {
		logger.Error("extension `x-vionetaos` not found")
		// return storeInfo, ErrComposeExtensionNameXVionetaOSNotFound
	}

	// add image to store info for check stable version function.
	storeInfo.Image = a.Image

	if err := loader.Transform(ex, &storeInfo); err != nil {
		return storeInfo, err
	}

	return storeInfo, nil
}
