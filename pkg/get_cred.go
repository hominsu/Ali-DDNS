package pkg

import (
	"crypto/tls"
	"crypto/x509"
	terrors "github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

func GetServerCreds() (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair("/cert/server.pem", "/cert/server.key")
	if err != nil {
		return nil, terrors.Wrap(err, "load server X509 key pair failed")
	}

	ca, err := ioutil.ReadFile("/cert/ca.crt")
	if err != nil {
		return nil, terrors.Wrap(err, "load ca failed")
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},  // 加载服务端证书
		ClientAuth:   tls.RequireAnyClientCert, // 需要认证客户端证书
		ClientCAs:    certPool,
	})

	return creds, nil
}

func GetClientCreds() (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair("/cert/client.pem", "/cert/client.key")
	if err != nil {
		return nil, terrors.Wrap(err, "load client X509 key pair failed")
	}

	ca, _ := ioutil.ReadFile("/cert/ca.crt")
	if err != nil {
		return nil, terrors.Wrap(err, "load ca failed")
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, // 加载服务端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	return creds, nil
}
