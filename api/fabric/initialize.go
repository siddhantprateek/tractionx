package fabric

import (
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"os"
	"path"

	handlers "github.com/siddhantprateek/tractionx/api/handlers"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type OrgSetupWrapper struct {
	setup handlers.OrgSetup
}

func Initialize(setup OrgSetupWrapper) (*handlers.OrgSetup, error) {
	log.Printf("Initializing connection for %s...\n", setup.setup.OrgName)
	clientConnection := setup.newGrpcConnection()
	id := setup.newIdentity()
	sign := setup.newSign()

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	setup.setup.Gateway = *gateway
	log.Println("Initialization complete")
	return &setup.setup, nil
}

func (setup OrgSetupWrapper) newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(setup.setup.TLSCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, setup.setup.GatewayPeer)

	connection, err := grpc.Dial(setup.setup.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

func (setup OrgSetupWrapper) newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(setup.setup.CryptoPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(setup.setup.MSPID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func (setup OrgSetupWrapper) newSign() identity.Sign {
	files, err := os.ReadDir(setup.setup.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(setup.setup.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}
