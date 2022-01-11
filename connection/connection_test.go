package connection

import (
	"fmt"
	lb "github.com/hyperledger/fabric-protos-go/peer/lifecycle"
	"testing"
)

var testConn = APIConfig{
	MSPID:       "test-org1",
	TLSCert:     "-----BEGIN CERTIFICATE-----\nMIICmzCCAkCgAwIBAgIUWxvh6s5EKY2K3JlgLxwFPKLu4cMwCgYIKoZIzj0EAwIw\nXTELMAkGA1UEBhMCWkExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2MDAxMRIwEAYD\nVQQKEwl0ZXN0LW9yZzExHjAcBgNVBAMTFXRlc3Qtb3JnMS1wZWVyLXRscy1jYTAe\nFw0yMTEyMTQxMjQ5MjlaFw0yMjEyMTQxMjQ5NTlaMHgxCzAJBgNVBAYTAlpBMQsw\nCQYDVQQIEwJFQzENMAsGA1UEERMENjAwMTESMBAGA1UEChMJdGVzdC1vcmcxMR8w\nCwYDVQQLEwRwZWVyMBAGA1UECxMJdGVzdC1vcmcxMRgwFgYDVQQDEw9wZWVyMC50\nZXN0LW9yZzEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATTwYy+Gj3FGlUgkKSj\nAVB69D9uADuW8EWSnKb4hebHiScNCEY3/UVhVAg68Kq3YLk1AUUBQqmkMwZs9ISO\nuLuOo4HCMIG/MA4GA1UdDwEB/wQEAwIDqDAdBgNVHSUEFjAUBggrBgEFBQcDAQYI\nKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUvVklUt5UsB0Xh5QuAHJS\nzmissMkwHwYDVR0jBBgwFoAUndlgKPL2woomYerkogEwSDYLGx0wQAYDVR0RBDkw\nN4IPcGVlcjAudGVzdC1vcmcxgiRwZWVyMC50ZXN0LW9yZzEuZnJhY3RhLml2b3J5\nY2hpYW4uaW8wCgYIKoZIzj0EAwIDSQAwRgIhAKdaCEisg5fD9FqVDbJM4H+3sXa2\n1RpcjalHimBU/u0HAiEAmNiFMpnVnL2448hjE28RF+xSFajvYrTm8jdZLxf/C+o=\n-----END CERTIFICATE-----",
	UserCert:    "-----BEGIN CERTIFICATE-----\nMIICVzCCAfygAwIBAgIUdSEPq1zve13a07TeFMtmFZCklRIwCgYIKoZIzj0EAwIw\nWTELMAkGA1UEBhMCWkExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2MDAxMRIwEAYD\nVQQKEwl0ZXN0LW9yZzExGjAYBgNVBAMTEXRlc3Qtb3JnMS1wZWVyLWNhMB4XDTIx\nMTIxNDEyNDkzMFoXDTIyMTIxNDEyNTAwMFowfDELMAkGA1UEBhMCWkExCzAJBgNV\nBAgTAkVDMQ0wCwYDVQQREwQ2MDAxMRIwEAYDVQQKEwl0ZXN0LW9yZzExIDAMBgNV\nBAsTBWFkbWluMBAGA1UECxMJdGVzdC1vcmcxMRswGQYDVQQDExJhZG1pbi11c2Vy\nLWRlZmF1bHQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASZP3x1n3EDn0FvZa9z\nkmRoh8XRkTcNOSLFPiAlnaPYrcJqqk3TcG5M0sIW/jhiNsbQDlvv1TGShibw50Nu\nZK9xo38wfTAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQU\n9XLzZeYzVa4m+YM1Az9s/aeKZIowHwYDVR0jBBgwFoAUaFLLNpksUh5dbSGX/4ZL\nBkUWiMkwHQYDVR0RBBYwFIISYWRtaW4tdXNlci1kZWZhdWx0MAoGCCqGSM49BAMC\nA0kAMEYCIQC69OjthoP3g7EVHJCxCQRkQzxW/bRqOTbeeV3Xtk2zLQIhAN5S3tms\n0IOdLij+VN+Y3PGiuMtyEjKMElF2wkwopjt4\n-----END CERTIFICATE-----",
	UserKey:     "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIFYX/gIRD1zrIct50qJdPaoV3BILYGfnnGK59yvSycRZoAoGCCqGSM49\nAwEHoUQDQgAEmT98dZ9xA59Bb2Wvc5JkaIfF0ZE3DTkixT4gJZ2j2K3CaqpN03Bu\nTNLCFv44YjbG0A5b79UxkoYm8OdDbmSvcQ==\n-----END EC PRIVATE KEY-----",
	RootCACert:  "-----BEGIN CERTIFICATE-----\nMIICQzCCAemgAwIBAgIUWjiE+zylhCUgJHDidi3HC2SpFkIwCgYIKoZIzj0EAwIw\nXTELMAkGA1UEBhMCWkExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2MDAxMRIwEAYD\nVQQKEwl0ZXN0LW9yZzExHjAcBgNVBAMTFXRlc3Qtb3JnMS1wZWVyLXRscy1jYTAe\nFw0yMTEyMTQxMjQ5MjlaFw0zMTEyMTIxMjQ5NTlaMF0xCzAJBgNVBAYTAlpBMQsw\nCQYDVQQIEwJFQzENMAsGA1UEERMENjAwMTESMBAGA1UEChMJdGVzdC1vcmcxMR4w\nHAYDVQQDExV0ZXN0LW9yZzEtcGVlci10bHMtY2EwWTATBgcqhkjOPQIBBggqhkjO\nPQMBBwNCAARv4Kdz9IeSO0GghSxVKHEW2SR67SNiqHLM0+RgkClffZxk8fg/Q4JW\nkgaqcxRf+LgJcxvHelFPM5vNpI8yW64no4GGMIGDMA4GA1UdDwEB/wQEAwIBBjAP\nBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBSd2WAo8vbCiiZh6uSiATBINgsbHTAf\nBgNVHSMEGDAWgBSd2WAo8vbCiiZh6uSiATBINgsbHTAgBgNVHREEGTAXghV0ZXN0\nLW9yZzEtcGVlci10bHMtY2EwCgYIKoZIzj0EAwIDSAAwRQIgNNyikL1asjElCH5H\nf+z89bFlpvQh5QSN2DOEVDTRi3sCIQD5HTvHrBnkBjZ9/+MGoinbLH6Uy9TC9+Zi\nOaZrUww97g==\n-----END CERTIFICATE-----",
	PeerAddress: "peer0.test-org1.fracta.ivorychian.io:8443",
}

const channelName = "main-channel-2"
const ccName = "test-abc"

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
		fmt.Println(string(res))
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
		req := lb.CheckCommitReadinessArgs{
			Sequence: 1,
			Name:     ccName,
		}
		res, err := con.LifecycleQueryCommitReadiness(&req, channelName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(res))
	})
	t.Run("Test Lifecycle query Commited", func(t *testing.T) {
		req := lb.QueryChaincodeDefinitionsArgs{}
		res, err := con.LifecycleQueryCommittedCC(&req, channelName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(res))
	})

	t.Run("Test Lifecycle query installed cc", func(t *testing.T) {
		req := lb.GetInstalledChaincodePackageArgs{}
		res, err := con.QueryInstallCC(&req, channelName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(res))
	})
}
