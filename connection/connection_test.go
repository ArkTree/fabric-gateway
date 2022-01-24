package connection

import (
	"fmt"
	lb "github.com/hyperledger/fabric-protos-go/peer/lifecycle"
	"testing"
)

var testConn = APIConfig{
	MSPID:       "test-org-111",
	TLSCert:     "-----BEGIN CERTIFICATE-----\nMIICsTCCAlegAwIBAgITfzi6/5OMlJCLo8CxMZUA82vfzzAKBggqhkjOPQQDAjBk\nMQwwCgYDVQQGEwNSU0ExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2NjAxMRUwEwYD\nVQQKEwx0ZXN0LW9yZy0xMTExITAfBgNVBAMTGHRlc3Qtb3JnLTExMS1wZWVyLXRs\ncy1jYTAeFw0yMjAxMjExMDQyNDFaFw0yMzAxMjExMDQzMTFaMIGCMQwwCgYDVQQG\nEwNSU0ExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2NjAxMRUwEwYDVQQKEwx0ZXN0\nLW9yZy0xMTExIjALBgNVBAsTBHBlZXIwEwYDVQQLEwx0ZXN0LW9yZy0xMTExGzAZ\nBgNVBAMTEnBlZXIwLnRlc3Qtb3JnLTExMTBZMBMGByqGSM49AgEGCCqGSM49AwEH\nA0IABD71Ya5JJDL9i0Bi903wL+wsd8lPiVts12Gj+36m9sA02ZjPiEziUVlNtLBg\njq+tMY6EvxeBlSnAsz86sknmM2KjgcgwgcUwDgYDVR0PAQH/BAQDAgOoMB0GA1Ud\nJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQW\nBBQ169f3KRto9rdSvjOqO8OTrlmaLjAfBgNVHSMEGDAWgBTW7anwvM6n223aDiKX\nRyKsRVpMGzBGBgNVHREEPzA9ghJwZWVyMC50ZXN0LW9yZy0xMTGCJ3BlZXIwLnRl\nc3Qtb3JnLTExMS5mcmFjdGEuaXZvcnljaGlhbi5pbzAKBggqhkjOPQQDAgNIADBF\nAiAWQXckkzcO4mouyQgov7haD29q0eOLL85ruePUgV4jKQIhAPVByosWtlDqB0xp\nKbPUQC5O5y6GLk//vtr3beYCwaqV\n-----END CERTIFICATE-----",
	UserCert:    "-----BEGIN CERTIFICATE-----\nMIICZDCCAgugAwIBAgIUBlAh+NXot6pcwOjnXBGqf8eMztAwCgYIKoZIzj0EAwIw\nYDEMMAoGA1UEBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMG\nA1UEChMMdGVzdC1vcmctMTExMR0wGwYDVQQDExR0ZXN0LW9yZy0xMTEtcGVlci1j\nYTAeFw0yMjAxMjExMDQyNDJaFw0yMzAxMjExMDQzMTJaMIGDMQwwCgYDVQQGEwNS\nU0ExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2NjAxMRUwEwYDVQQKEwx0ZXN0LW9y\nZy0xMTExIzAMBgNVBAsTBWFkbWluMBMGA1UECxMMdGVzdC1vcmctMTExMRswGQYD\nVQQDExJhZG1pbi11c2VyLWRlZmF1bHQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC\nAAQZXXgOXQJS29X2S6Cy/I16ZuvdVkO9BcGFgwSiLS9WfX//z/CwIYcfWNmQXEwP\nkb0I1ORabGBURN52JvGgU8ZXo38wfTAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/\nBAIwADAdBgNVHQ4EFgQUqRvcHyVlzIa7hMaOVbXQoTfE5j8wHwYDVR0jBBgwFoAU\nmZtmceoNv6VVNA5aQzF/bQhpRRMwHQYDVR0RBBYwFIISYWRtaW4tdXNlci1kZWZh\ndWx0MAoGCCqGSM49BAMCA0cAMEQCICqKyU6l4TySopiU/jBNShXTPkvs37f1E3t7\ndUWldnP4AiAM9YAmy1PYP0NXn1eRVtwA/s0dKmZpCeMf6Orin3CvLw==\n-----END CERTIFICATE-----",
	UserKey:     "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIHj0mO8rIne7bQaWFAnflJfOIu8TPi8i2Q6wAdX5aWinoAoGCCqGSM49\nAwEHoUQDQgAEGV14Dl0CUtvV9kugsvyNembr3VZDvQXBhYMEoi0vVn1//8/wsCGH\nH1jZkFxMD5G9CNTkWmxgVETedibxoFPGVw==\n-----END EC PRIVATE KEY-----",
	RootCACert:  "-----BEGIN CERTIFICATE-----\nMIICVDCCAfqgAwIBAgIUIGFcaGPs9Dfa0qAbPd3wV4r417gwCgYIKoZIzj0EAwIw\nZDEMMAoGA1UEBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMG\nA1UEChMMdGVzdC1vcmctMTExMSEwHwYDVQQDExh0ZXN0LW9yZy0xMTEtcGVlci10\nbHMtY2EwHhcNMjIwMTIxMTA0MjQxWhcNMzIwMTE5MTA0MzExWjBkMQwwCgYDVQQG\nEwNSU0ExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2NjAxMRUwEwYDVQQKEwx0ZXN0\nLW9yZy0xMTExITAfBgNVBAMTGHRlc3Qtb3JnLTExMS1wZWVyLXRscy1jYTBZMBMG\nByqGSM49AgEGCCqGSM49AwEHA0IABDoPKkUsZh7Gqzv9xLNJh7izExC/9PngEqmL\nLsG7HgShOky5ZZ+LYPF8954kCsqYrBHNP4LTHytN3T093y8xRc2jgYkwgYYwDgYD\nVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFNbtqfC8zqfb\nbdoOIpdHIqxFWkwbMB8GA1UdIwQYMBaAFNbtqfC8zqfbbdoOIpdHIqxFWkwbMCMG\nA1UdEQQcMBqCGHRlc3Qtb3JnLTExMS1wZWVyLXRscy1jYTAKBggqhkjOPQQDAgNI\nADBFAiEA6RdJksXpl6ZAGcFQdrZuxOQ6oxMOXVyrLL4xghrCVM0CIARd6nikEgsn\nq/p2sjTMzU/JzemUBwyQvKJFTdSSKHTF\n-----END CERTIFICATE-----",
	PeerAddress: "peer0.test-org-111.fracta.ivorychian.io:8443",
}

const channelName = "sanie"
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
		req := lb.QueryInstalledChaincodesArgs{}
		res, err := con.QueryInstallCC(&req, channelName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	})
}
