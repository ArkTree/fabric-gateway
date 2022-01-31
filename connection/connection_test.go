package connection

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	lb "github.com/hyperledger/fabric-protos-go/peer/lifecycle"
	"io"
	"testing"
)

var testConn = APIConfig{
	MSPID:       "test-org-111",
	TLSCert:     "-----BEGIN CERTIFICATE-----\nMIICsjCCAligAwIBAgIUV4F5119omYBg2qTtBA5vITx4w4cwCgYIKoZIzj0EAwIw\nZDEMMAoGA1UEBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMG\nA1UEChMMdGVzdC1vcmctMTExMSEwHwYDVQQDExh0ZXN0LW9yZy0xMTEtcGVlci10\nbHMtY2EwHhcNMjIwMTMxMDkyMTQ2WhcNMjMwMTMxMDkyMjE2WjCBgjEMMAoGA1UE\nBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMGA1UEChMMdGVz\ndC1vcmctMTExMSIwCwYDVQQLEwRwZWVyMBMGA1UECxMMdGVzdC1vcmctMTExMRsw\nGQYDVQQDExJwZWVyMC50ZXN0LW9yZy0xMTEwWTATBgcqhkjOPQIBBggqhkjOPQMB\nBwNCAAQ/B8tT5U9OFG+TqusRfoeSNoO2s/EK5JmKQtCV2uSGDns0nRzYvUBbNyVW\nEYrNC1V2nJYmrC/dVtw/8PJXqnT7o4HIMIHFMA4GA1UdDwEB/wQEAwIDqDAdBgNV\nHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4E\nFgQUOGJ3Fm8sl+UP4kOUnR97mwdYV+UwHwYDVR0jBBgwFoAU1u2p8LzOp9tt2g4i\nl0cirEVaTBswRgYDVR0RBD8wPYIScGVlcjAudGVzdC1vcmctMTExgidwZWVyMC50\nZXN0LW9yZy0xMTEuZnJhY3RhLml2b3J5Y2hpYW4uaW8wCgYIKoZIzj0EAwIDSAAw\nRQIgXHF74mGMoxs/MkwasIPnKUX8shnhOUjHuzXUJaNecEsCIQDUhbWxHgp3Rxd4\naX5uUcuvLSL/b6VYGfMqSdybLEVdrQ==\n-----END CERTIFICATE-----",
	UserCert:    "-----BEGIN CERTIFICATE-----\nMIICZTCCAgugAwIBAgIUTIuLgBR7EaSCwsCXO0sacQdNSNUwCgYIKoZIzj0EAwIw\nYDEMMAoGA1UEBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMG\nA1UEChMMdGVzdC1vcmctMTExMR0wGwYDVQQDExR0ZXN0LW9yZy0xMTEtcGVlci1j\nYTAeFw0yMjAxMzEwOTIxNDdaFw0yMzAxMzEwOTIyMTdaMIGDMQwwCgYDVQQGEwNS\nU0ExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2NjAxMRUwEwYDVQQKEwx0ZXN0LW9y\nZy0xMTExIzAMBgNVBAsTBWFkbWluMBMGA1UECxMMdGVzdC1vcmctMTExMRswGQYD\nVQQDExJhZG1pbi11c2VyLWRlZmF1bHQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC\nAARl0VuFHnMgOIsVXoTul2OORdXuvVb6ff/jAafM96zDHKt0u3df6TttgDynSa+d\nJPd9Ny3h3o4RMtCtjQUY8K7vo38wfTAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/\nBAIwADAdBgNVHQ4EFgQUzapv9sYfDPio+H0Bgs5ZQRXa2FQwHwYDVR0jBBgwFoAU\nmZtmceoNv6VVNA5aQzF/bQhpRRMwHQYDVR0RBBYwFIISYWRtaW4tdXNlci1kZWZh\ndWx0MAoGCCqGSM49BAMCA0gAMEUCIDsVpLZ/K/UZJMP/qBuad3BLeq4spd+hWQg8\nIUwUcObwAiEA05zkqE5TR2ZdeCqjBwedClWTDth6rW2C/OMWwvzxTjo=\n-----END CERTIFICATE-----",
	UserKey:     "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIDnXUMiOrRMyX/VVZUYOnxivIQlmT1EAYx2nln4cWiPRoAoGCCqGSM49\nAwEHoUQDQgAEZdFbhR5zIDiLFV6E7pdjjkXV7r1W+n3/4wGnzPeswxyrdLt3X+k7\nbYA8p0mvnST3fTct4d6OETLQrY0FGPCu7w==\n-----END EC PRIVATE KEY-----",
	RootCACert:  "-----BEGIN CERTIFICATE-----\nMIICVDCCAfqgAwIBAgIUIGFcaGPs9Dfa0qAbPd3wV4r417gwCgYIKoZIzj0EAwIw\nZDEMMAoGA1UEBhMDUlNBMQswCQYDVQQIEwJFQzENMAsGA1UEERMENjYwMTEVMBMG\nA1UEChMMdGVzdC1vcmctMTExMSEwHwYDVQQDExh0ZXN0LW9yZy0xMTEtcGVlci10\nbHMtY2EwHhcNMjIwMTIxMTA0MjQxWhcNMzIwMTE5MTA0MzExWjBkMQwwCgYDVQQG\nEwNSU0ExCzAJBgNVBAgTAkVDMQ0wCwYDVQQREwQ2NjAxMRUwEwYDVQQKEwx0ZXN0\nLW9yZy0xMTExITAfBgNVBAMTGHRlc3Qtb3JnLTExMS1wZWVyLXRscy1jYTBZMBMG\nByqGSM49AgEGCCqGSM49AwEHA0IABDoPKkUsZh7Gqzv9xLNJh7izExC/9PngEqmL\nLsG7HgShOky5ZZ+LYPF8954kCsqYrBHNP4LTHytN3T093y8xRc2jgYkwgYYwDgYD\nVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFNbtqfC8zqfb\nbdoOIpdHIqxFWkwbMB8GA1UdIwQYMBaAFNbtqfC8zqfbbdoOIpdHIqxFWkwbMCMG\nA1UdEQQcMBqCGHRlc3Qtb3JnLTExMS1wZWVyLXRscy1jYTAKBggqhkjOPQQDAgNI\nADBFAiEA6RdJksXpl6ZAGcFQdrZuxOQ6oxMOXVyrLL4xghrCVM0CIARd6nikEgsn\nq/p2sjTMzU/JzemUBwyQvKJFTdSSKHTF\n-----END CERTIFICATE-----",
	PeerAddress: "peer0.test-org-111.fracta.ivorychian.io:8443",
}

const channelName = "first-channel"
const ccName = "test-chaincode"

type ConnectionInfo struct {
	Address            string `json:"address"`
	DialTimeout        string `json:"dial_timeout"`
	TLSRequired        bool   `json:"tls_required"`
	ClientAuthRequired bool   `json:"client_auth_required"`
	ClientKey          string `json:"client_key"`
	ClientCert         string `json:"client_cert"`
	RootCert           string `json:"root_cert"`
}

type MetaDataInfo struct {
	Path  string `json:"path"`
	Type  string `json:"type"`
	Label string `json:"label"`
}

type CCPackageInfo struct {
	*ConnectionInfo
	*MetaDataInfo
}

func TestNewGateway(t *testing.T) {
	con, err := New(&testConn)
	defer con.Close()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Test Lifecycle install cc", func(t *testing.T) {
		packageInfo := CCPackageInfo{
			ConnectionInfo: &ConnectionInfo{
				Address:            fmt.Sprintf("%s.%s:%d", ccName, testConn.MSPID, 1234),
				DialTimeout:        "10s",
				TLSRequired:        false,
				ClientAuthRequired: false,
				ClientKey:          "-----BEGIN EC PRIVATE KEY----- ... -----END EC PRIVATE KEY-----",
				ClientCert:         "-----BEGIN CERTIFICATE----- ... -----END CERTIFICATE-----",
				RootCert:           "-----BEGIN CERTIFICATE---- ... -----END CERTIFICATE-----",
			},
			MetaDataInfo: &MetaDataInfo{
				Path:  "",
				Type:  "external",
				Label: ccName,
			},
		}
		tarPackage, _ := packageCC(packageInfo)
		res, err := con.InstallNewCC(&lb.InstallChaincodeArgs{ChaincodeInstallPackage: tarPackage}, channelName)
		fmt.Println(err)
		fmt.Println(res)
	})

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

func packageCC(ccInfo CCPackageInfo) ([]byte, error) {
	connectionTar := bytes.NewBuffer(nil)
	packageTar := bytes.NewBuffer(nil)

	connectionJson, err := json.Marshal(ccInfo.ConnectionInfo)
	if err != nil {
		return nil, err
	}
	metaJson, err := json.Marshal(ccInfo.MetaDataInfo)
	if err != nil {
		return nil, err
	}

	gw := gzip.NewWriter(connectionTar)
	tw := tar.NewWriter(gw)

	err = tw.WriteHeader(&tar.Header{
		Name: "connection.json",
		Size: int64(len(connectionJson)),
		Mode: 0100644,
	})

	if err != nil {
		return nil, err
	}

	_, err = tw.Write(connectionJson)

	if err != nil {
		return nil, err
	}

	err = tw.Close()
	err = gw.Close()

	gw.Reset(packageTar)
	tw = tar.NewWriter(gw)

	err = tw.WriteHeader(&tar.Header{
		Name: "metadata.json",
		Size: int64(len(metaJson)),
		Mode: 0100644,
	})

	if err != nil {
		return nil, err
	}

	_, err = tw.Write(metaJson)

	if err != nil {
		return nil, err
	}

	err = tw.WriteHeader(&tar.Header{
		Name: "code.tar.gz",
		Size: int64(connectionTar.Len()),
		Mode: 0100644,
	})

	_, err = io.Copy(tw, connectionTar)

	if err != nil {
		return nil, err
	}

	tw.Close()
	gw.Close()

	return packageTar.Bytes(), nil
}
