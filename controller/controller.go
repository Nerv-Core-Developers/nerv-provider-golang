package controller

import (
	"mime/multipart"
	"net/http"

	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/noid"
)

func (ctrl *Controller) Init() error {
	ctrl.Logger.Init("controller")
	ctrl.Logger.Log("Init controller...")
	return nil
}
func (ctrl *Controller) GetFunctionInfo(ID noid.Identifier) (shared.FunctionInfo, error) {
	var result shared.FunctionInfo
	return result, nil
}
func (ctrl *Controller) GetAssetLink(ID noid.Identifier) (string, error) {
	var result string
	return result, nil
}
func (ctrl *Controller) DeployPkg(pkg *multipart.FileHeader) error {
	return nil
}
func (ctrl *Controller) APIHandler(*http.Request, []string) ([]byte, error) {
	return nil, nil
}
func (ctrl *Controller) Version() int {
	return Version
}
func (ctrl *Controller) GetNodeVersion() shared.NodeVersion {
	return shared.NodeVersion{}
}
func (ctrl *Controller) HandleMessage(msg *shared.NetMsg) error {
	return nil
}
