package connection

import (
	"fmt"
	lb "github.com/hyperledger/fabric-protos-go/peer/lifecycle"
	"testing"
)

var testConn = APIConfig{
	MSPID:       "test-org-111",
	TLSCert:     "-----BEGIN CERTIFICATE-----\nMIICszCCAligAwIBAgIUZCleYxUw+qEJ9iHVJXaZL+Vjo5IwCgYIKoZIzj0EAwIw\nZDEMMAoGA1UEBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMG\nA1UEChMMdGVzdC1vcmctMTExMSEwHwYDVQQDExh0ZXN0LW9yZy0xMTEtcGVlci10\nbHMtY2EwHhcNMjIwMTIwMTUzNDM1WhcNMjMwMTIwMTUzNTA1WjCBgjEMMAoGA1UE\nBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMGA1UEChMMdGVz\ndC1vcmctMTExMSIwCwYDVQQLEwRwZWVyMBMGA1UECxMMdGVzdC1vcmctMTExMRsw\nGQYDVQQDExJwZWVyMC50ZXN0LW9yZy0xMTEwWTATBgcqhkjOPQIBBggqhkjOPQMB\nBwNCAAStqo3A1UNqgIVY7US8oSTVMuuCz/48ImWH8iS7KDxr5yuRSM3j9WiN1wl+\neDJWDhKqMNEMbC82Qq5EsGpNVCASo4HIMIHFMA4GA1UdDwEB/wQEAwIDqDAdBgNV\nHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4E\nFgQU07owSmuCgwbZlsdlTzHpL50uPpowHwYDVR0jBBgwFoAUV4JS/iWBR7PALapS\ns55JUSMNi0swRgYDVR0RBD8wPYIScGVlcjAudGVzdC1vcmctMTExgidwZWVyMC50\nZXN0LW9yZy0xMTEuZnJhY3RhLml2b3J5Y2hpYW4uaW8wCgYIKoZIzj0EAwIDSQAw\nRgIhANS/Vm8CrEMVGU/qZFptqvrnkZTzTg2g7EdCdNeis4HJAiEAlQDQ0K1SHf+s\nzl/wxdzNmPr2DsI8oC/1KvvoLNxxTnY=\n-----END CERTIFICATE-----",
	UserCert:    "-----BEGIN CERTIFICATE-----\nMIICZjCCAgugAwIBAgIUYNMPPepgmwgcTUV+EJdk9bloayswCgYIKoZIzj0EAwIw\nYDEMMAoGA1UEBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMG\nA1UEChMMdGVzdC1vcmctMTExMR0wGwYDVQQDExR0ZXN0LW9yZy0xMTEtcGVlci1j\nYTAeFw0yMjAxMjAxNTM0MzVaFw0yMzAxMjAxNTM1MDVaMIGDMQwwCgYDVQQGEwNS\nU0ExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2NjAxMRUwEwYDVQQKEwx0ZXN0LW9y\nZy0xMTExIzAMBgNVBAsTBWFkbWluMBMGA1UECxMMdGVzdC1vcmctMTExMRswGQYD\nVQQDExJhZG1pbi11c2VyLWRlZmF1bHQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC\nAASuAmZ8jp1LorIvm9e4kcs3OhIbV/SjVLw3PyzJEHq3jFJ/my4OWrJWzQ261Pik\nVaw2tu3L8CbUFEwhN6TY2Uf7o38wfTAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/\nBAIwADAdBgNVHQ4EFgQUypSVO9gh+4yAkGTNm5HcZvHcp8YwHwYDVR0jBBgwFoAU\np9tqREsu1jWYlAQbBuCKquin1NYwHQYDVR0RBBYwFIISYWRtaW4tdXNlci1kZWZh\ndWx0MAoGCCqGSM49BAMCA0kAMEYCIQCLoTnN1ZBtKluBzSjidFpi2XlSmD/Qx+YI\nOmiwlHIdCgIhAK1Gc6u9zoqZ+HC51b/T71AP9UaPeW43d73YCYLcPZRf\n-----END CERTIFICATE-----",
	UserKey:     "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIB/SuLL9rITv7uyF9WUmGknCjTOtBwmiV7cXlbTgH20XoAoGCCqGSM49\nAwEHoUQDQgAErgJmfI6dS6KyL5vXuJHLNzoSG1f0o1S8Nz8syRB6t4xSf5suDlqy\nVs0NutT4pFWsNrbty/Am1BRMITek2NlH+w==\n-----END EC PRIVATE KEY-----",
	RootCACert:  "-----BEGIN CERTIFICATE-----\nMIICVDCCAfqgAwIBAgIUGecvSP1uM8zveiKHSClAifNTuZ0wCgYIKoZIzj0EAwIw\nZDEMMAoGA1UEBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMG\nA1UEChMMdGVzdC1vcmctMTExMSEwHwYDVQQDExh0ZXN0LW9yZy0xMTEtcGVlci10\nbHMtY2EwHhcNMjIwMTIwMTExMDA0WhcNMzIwMTE4MTExMDM0WjBkMQwwCgYDVQQG\nEwNSU0ExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2NjAxMRUwEwYDVQQKEwx0ZXN0\nLW9yZy0xMTExITAfBgNVBAMTGHRlc3Qtb3JnLTExMS1wZWVyLXRscy1jYTBZMBMG\nByqGSM49AgEGCCqGSM49AwEHA0IABEcpFN6JpUVRR7ncGZ7LLfzsgcyhOBITxJ1K\nVSIf8LoraUuW+bWi5vsG7CAG9oFepuknaxIT1LGM35YMd6FX7UqjgYkwgYYwDgYD\nVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFFeCUv4lgUez\nwC2qUrOeSVEjDYtLMB8GA1UdIwQYMBaAFFeCUv4lgUezwC2qUrOeSVEjDYtLMCMG\nA1UdEQQcMBqCGHRlc3Qtb3JnLTExMS1wZWVyLXRscy1jYTAKBggqhkjOPQQDAgNI\nADBFAiEAgxohtnqSeh2c+gpKyRjPX3kGJeW8JqEhBYwS0cfZKvACIF2g82EGgXok\nAWmqULWM67wTWOcWXyblsBILy24MjOY6\n-----END CERTIFICATE-----",
	PeerAddress: "peer1.test-org-111.fracta.ivorychian.io:8443",
}

const channelName = "jaie"
const ccName = "test-chaincode"

func TestNewGateway(t *testing.T) {
	con, err := New(&testConn)
	defer con.Close()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Test Lifecycle Approved cc", func(t *testing.T) {
		req := lb.ApproveChaincodeDefinitionForMyOrgArgs{
			Sequence: 1,
			Name:     ccName,
			Version:  "1.0",
		}
		res, err := con.LifecycleApproveCC(&req, channelName)

		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(res))
	})
	t.Run("Test Lifecycle query Approved", func(t *testing.T) {
		req := lb.QueryApprovedChaincodeDefinitionArgs{
			Name:     ccName,
			Sequence: 1,
		}
		res, err := con.LifecycleQueryApproved(&req, channelName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	})
	t.Run("Test Lifecycle Commit CC", func(t *testing.T) {
		req := lb.CommitChaincodeDefinitionArgs{
			Sequence: 1,
			Name:     ccName,
			Version:  "1.0",
		}
		res, err := con.LifecycleCommitCC(&req, channelName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(res))
	})
	t.Run("Test Lifecycle query Commit Readiness", func(t *testing.T) {
		network := con.GetNetwork("test-abc")
		contract := network.GetContract("test-chaincode")
		res, err := contract.SubmitTransaction("", "org.hyperledger.fabric:GetMetadata")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	})

	t.Run("Test Lifecycle submit", func(t *testing.T) {
		req := lb.CheckCommitReadinessArgs{
			Sequence: 1,
			Name:     ccName,
		}
		res, err := con.LifecycleQueryCommitReadiness(&req, channelName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	})

	t.Run("Test Lifecycle query Commited", func(t *testing.T) {
		req := lb.QueryChaincodeDefinitionsArgs{}
		res, err := con.LifecycleQueryCommittedCC(&req, channelName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	})

	t.Run("Test Lifecycle query installed cc", func(t *testing.T) {
		req := lb.GetInstalledChaincodePackageArgs{}
		res, err := con.QueryInstallCC(&req, channelName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	})
}
