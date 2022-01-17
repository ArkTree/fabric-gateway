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
	LifecycleQueryApproved(req *lb.QueryApprovedChaincodeDefinitionArgs, channelName string) (*lb.QueryApprovedChaincodeDefinitionResult, error)
	LifecycleCommitCC(req *lb.CommitChaincodeDefinitionArgs, channelName string) ([]byte, error)
	LifecycleQueryCommittedCC(req *lb.QueryChaincodeDefinitionsArgs, channelName string) (*lb.QueryChaincodeDefinitionResult, error)
	LifecycleQueryCommitReadiness(req *lb.CheckCommitReadinessArgs, channelName string) (*lb.CheckCommitReadinessResult, error)
	QueryInstallCC(req *lb.GetInstalledChaincodePackageArgs, channelName string) (*lb.GetInstalledChaincodePackageResult, error)
	InstallNewCC(req *lb.InstallChaincodeArgs, channelName string) ([]byte, error)
}

func (f FabricConnection) LifecycleApproveCC(req *lb.ApproveChaincodeDefinitionForMyOrgArgs, channelName string) ([]byte, error) {
	return f.ccLifecycle(channelName, lifecycleApproveChaincodeFuncName, req)
}

func (f FabricConnection) LifecycleQueryApproved(req *lb.QueryApprovedChaincodeDefinitionArgs, channelName string) (*lb.QueryApprovedChaincodeDefinitionResult, error) {
	res, err := f.ccLifecycle(channelName, lifecycleQueryApprovedCCDefinitionFunc, req)
	if err != nil {
		return nil, err
	}
	qApprovedRes := &lb.QueryApprovedChaincodeDefinitionResult{}
	err = proto.Unmarshal(res, qApprovedRes)
	if err != nil {
		return nil, err
	}
	return qApprovedRes, nil
}

func (f FabricConnection) LifecycleCommitCC(req *lb.CommitChaincodeDefinitionArgs, channelName string) ([]byte, error) {
	return f.ccLifecycle(channelName, lifecycleCommitFuncName, req)
}

func (f FabricConnection) LifecycleQueryCommittedCC(req *lb.QueryChaincodeDefinitionsArgs, channelName string) (*lb.QueryChaincodeDefinitionResult, error) {
	res, err := f.ccLifecycle(channelName, lifecycleQueryChaincodeDefinitionsFunc, req)
	if err != nil {
		return nil, err
	}
	qCommitedCC := &lb.QueryChaincodeDefinitionResult{}
	err = proto.Unmarshal(res, qCommitedCC)
	if err != nil {
		return nil, err
	}
	return qCommitedCC, nil
}

func (f FabricConnection) LifecycleQueryCommitReadiness(req *lb.CheckCommitReadinessArgs, channelName string) (*lb.CheckCommitReadinessResult, error) {
	res, err := f.ccLifecycle(channelName, lifecycleCheckCommitReadinessFuncName, req)
	if err != nil {
		return nil, err
	}
	qCommitedReadiness := &lb.CheckCommitReadinessResult{}
	err = proto.Unmarshal(res, qCommitedReadiness)
	if err != nil {
		return nil, err
	}
	return qCommitedReadiness, nil
}

func (f FabricConnection) QueryInstallCC(req *lb.GetInstalledChaincodePackageArgs, channelName string) (*lb.GetInstalledChaincodePackageResult, error) {
	res, err := f.ccLifecycle(channelName, lifecycleQueryInstalledChaincodesFunc, req)
	if err == nil {
		return nil, err
	}
	qInstalledCC := &lb.GetInstalledChaincodePackageResult{}
	err = proto.Unmarshal(res, qInstalledCC)
	if err != nil {
		return nil, err
	}
	return qInstalledCC, nil
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
