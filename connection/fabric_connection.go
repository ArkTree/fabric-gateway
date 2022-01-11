package connection

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/golang/glog"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type FabricConnection struct {
	*client.Gateway
}

type APIConfig struct {
	MSPID       string
	TLSCert     string
	UserCert    string
	UserKey     string
	RootCACert  string
	PeerAddress string
}

func New(config *APIConfig) (*FabricConnection, error) {
	clientConnection, err := grpc.Dial(config.PeerAddress, grpc.WithTransportCredentials(credentials.NewTLS(NewCertPool(config.RootCACert))))

	if err != nil {
		glog.Errorf("Failed to dail connection: %v", err)
		return nil, err
	}

	id, err := NewIdentity(config.UserCert, config.MSPID)
	if err != nil {
		glog.Errorf("Failed to create new Identity: %v", err)
		return nil, err
	}
	sign, err := NewSign(config.UserKey)
	if err != nil {
		glog.Errorf("Failed to create new signer: %v", err)
		return nil, err
	}

	// Create a Gateway connection for a specific client identity.
	gateway, err := client.Connect(id, client.WithSign(sign), client.WithClientConnection(clientConnection))
	if err != nil {
		glog.Errorf("Failed to dail connection: %v", err)
		return nil, err
	}
	return &FabricConnection{gateway}, nil
}

func NewCertPool(rootCert string) *tls.Config {
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM([]byte(rootCert))
	return &tls.Config{
		RootCAs: certPool,
	}
}

func NewIdentity(cert, mspID string) (*identity.X509Identity, error) {
	certificate, err := identity.CertificateFromPEM([]byte(cert))
	if err != nil {
		return nil, err
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		return nil, err
	}

	return id, nil
}

// NewSign creates a function that generates a digital signature from a message digest using a private key.
func NewSign(pvtKey string) (identity.Sign, error) {

	decoded, _ := pem.Decode([]byte(pvtKey))
	pvtKeyEC, err := x509.ParseECPrivateKey(decoded.Bytes)
	if err != nil {
		return nil, err
	}
	sign, err := identity.NewPrivateKeySign(pvtKeyEC)
	if err != nil {
		return nil, err
	}

	return sign, nil
}
