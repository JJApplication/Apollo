/*
   Create: 2024/5/15
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package cert_manager

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"os"
	"path/filepath"
)

const (
	CertManager = "[Alarm Manager]"
)

type Cert struct {
	CA        string `json:"ca"`
	Algorithm string `json:"algorithm"`
	Name      string `json:"name"`
	Start     string `json:"start"`
	End       string `json:"end"`
}

func CertInfo() Cert {
	certCerFile, err := os.ReadFile(filepath.Join(config.ApolloConf.SSLRoot, config.ApolloConf.SSLCert))
	if err != nil {
		logger.LoggerSugar.Errorf("read certificate error: %s", err.Error())
	}

	certBlock, _ := pem.Decode(certCerFile)
	if certBlock == nil {
		logger.LoggerSugar.Error("decode certificate error")
	}

	// 裁剪字节到实际长度
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		logger.LoggerSugar.Errorf("parse certificate error: %s", err.Error())
		return Cert{}
	}

	return Cert{
		CA:        cert.Issuer.CommonName,
		Algorithm: cert.SignatureAlgorithm.String(),
		Name:      cert.Subject.CommonName,
		Start:     utils.GetTimeByFormat(utils.TimeFormat, cert.NotBefore.Unix()),
		End:       utils.GetTimeByFormat(utils.TimeFormat, cert.NotAfter.Unix()),
	}
}
