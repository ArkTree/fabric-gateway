package connection

import (
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	lb "github.com/hyperledger/fabric-protos-go/peer/lifecycle"
)

const (
	lifecycleCC                               = "_lifecycle"
	lifecycleInstallFuncName                  = "InstallChaincode"
	lifecycleQueryInstalledChaincodesFunc     = "QueryInstalledChaincodes"
	lifecycleGetInstalledChaincodePackageFunc = "GetInstalledChaincodePackage"
	lifecycleApproveChaincodeFuncName         = "ApproveChaincodeDefinitionForMyOrg"
	lifecycleQueryApprovedCCDefinitionFunc    = "QueryApprovedChaincodeDefinition"
	lifecycleCheckCommitReadinessFuncName     = "CheckCommitReadiness"
	lifecycleCommitFuncName                   = "CommitChaincodeDefinition"
	lifecycleQueryChaincodeDefinitionFunc     = "QueryChaincodeDefinition"
	lifecycleQueryChaincodeDefinitionsFunc    = "QueryChaincodeDefinitions"
)

type CCLifecycleActions interface {
	LifecycleApproveCC(req *lb.ApproveChaincodeDefinitionForMyOrgArgs, channelName string) ([]byte, error)
	LifecycleQueryApproved(req *lb.QueryApprovedChaincodeDefinitionArgs, channelName string) ([]byte, error)
	LifecycleCommitCC(req *lb.CommitChaincodeDefinitionArgs, channelName string) ([]byte, error)
	LifecycleQueryCommittedCC(req *lb.QueryChaincodeDefinitionsArgs, channelName string) ([]byte, error)
	LifecycleQueryCommitReadiness(req *lb.CheckCommitReadinessArgs, channelName string) ([]byte, error)
	QueryInstallCC(req *lb.GetInstalledChaincodePackageArgs, channelName string) ([]byte, error)
	InstallNewCC(req *lb.InstallChaincodeArgs, channelName string) ([]byte, error)
}

func (f FabricConnection) LifecycleApproveCC(req *lb.ApproveChaincodeDefinitionForMyOrgArgs, channelName string) ([]byte, error) {
	return f.ccLifecycle(channelName, lifecycleApproveChaincodeFuncName, req)
}

func (f FabricConnection) LifecycleQueryApproved(req *lb.QueryApprovedChaincodeDefinitionArgs, channelName string) ([]byte, error) {
	return f.ccLifecycle(channelName, lifecycleQueryApprovedCCDefinitionFunc, req)
}

func (f FabricConnection) LifecycleCommitCC(req *lb.CommitChaincodeDefinitionArgs, channelName string) ([]byte, error) {
	return f.ccLifecycle(channelName, lifecycleCommitFuncName, req)
}

func (f FabricConnection) LifecycleQueryCommittedCC(req *lb.QueryChaincodeDefinitionsArgs, channelName string) ([]byte, error) {
	return f.ccLifecycle(channelName, lifecycleQueryChaincodeDefinitionsFunc, req)
}

func (f FabricConnection) LifecycleQueryCommitReadiness(req *lb.CheckCommitReadinessArgs, channelName string) ([]byte, error) {
	return f.ccLifecycle(channelName, lifecycleCheckCommitReadinessFuncName, req)
}

func (f FabricConnection) QueryInstallCC(req *lb.GetInstalledChaincodePackageArgs, channelName string) ([]byte, error) {
	return f.ccLifecycle(channelName, lifecycleQueryInstalledChaincodesFunc, req)
}

func (f FabricConnection) InstallNewCC(req *lb.InstallChaincodeArgs, channelName string) ([]byte, error) {
	return f.ccLifecycle(channelName, lifecycleInstallFuncName, req)
}

func (f FabricConnection) ccLifecycle(channelName, ccAction string, req proto.Message) ([]byte, error) {
	argsBytes, err := proto.Marshal(req)
	if err != nil {
		glog.Errorf("Failed to marshal %v", err)
		return nil, err
	}
	network := f.GetNetwork(channelName)
	contract := network.GetContract(lifecycleCC)

	submitResult, err := contract.Submit(ccAction, client.WithBytesArguments(argsBytes))
	if err != nil {
		return nil, err
	}

	return submitResult, nil
}
