package provisioning

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomPassword(t *testing.T) {
	pwd1, err := generateRandomPassword()
	if err != nil {
		t.Errorf("Unexpected error while generating random password: %s", err)
	}
	if pwd1 == "" {
		t.Errorf("Expected a valid string but got null")
	}
	pwd2, err := generateRandomPassword()
	if err != nil {
		t.Errorf("Unexpected error while re-generating random password: %s", err)
	} else {
		assert.False(t, pwd1 == pwd2, "regenerated random password should not match pervious one")
	}
}

func TestGenerateSerialNumber(t *testing.T) {
	sn, err := generateSerialNumber()
	if err != nil {
		t.Errorf("Unexpected error while generating a serial number: %s", err)
	} else {
		assert.GreaterOrEqual(t, sn.Cmp(big.NewInt(0)), 0, "serial number is negative")
	}
}

func TestGenerateTlsCertificate(t *testing.T) {
	cert, err := generateTlsCertificate()
	if err != nil {
		t.Errorf("Unexpected error while generating a certificate: %s", err)
	} else {
		assert.NotEqual(t, cert.certificate, "", "empty certificate")
		assert.NotEqual(t, cert.privateKey, "", "empty private key")
	}
}
