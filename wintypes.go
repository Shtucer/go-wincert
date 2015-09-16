package wincert

import (
	"syscall"
)


const (
	CERT_STORE_PROV_SYSTEM          = 9
	CERT_SYSTEM_STORE_LOCAL_MACHINE = 0x20000
	CERT_NAME_SIMPLE_DISPLAY_TYPE   = 0x4

	AT_SIGNATURE    = 0x2
	AT_KEYEXCHANGE  = 0x1
	OIG_RSA_RSA1SHA = "1.2.840.113549.1.1.5"

	CALG_SHA_256     = 0x0000800c
	CRYPT_EXPORTABLE = 0x1
	CRYPT_NEWKEYSET  = 0x8

	CERT_KEY_PROV_INFO_PROP_ID = 0x2
)

type CertInfo struct {
	DwVersion            uint32
	SerialNumber         CryptoApiBlob
	SignatureAlgorithm   AlgorithmId
	Issuer               CryptoApiBlob
	NotBefore            syscall.Filetime
	NotAfter             syscall.Filetime
	Subject              CryptoApiBlob
	SubjectPublicKeyInfo PublicKeyInfo
	IssuerUniqueId       CryptBitBlob
	SubjectUniqueId      CryptBitBlob
	CExtension           uint32
	RgExtension          *CertExtension
}

type PublicKeyInfo struct {
	Algorithm AlgorithmId
	PublicKey CryptBitBlob
}

type CryptBitBlob struct {
	CbData      uint32
	PbData      *uint8
	CUnusedBits uint32
}

type AlgorithmId struct {
	PszObjId   *int8
	Parameters CryptoApiBlob
}

type SystemTime struct {
	Year         uint16
	Month        uint16
	DayOfWeek    uint16
	Day          uint16
	Hour         uint16
	Minute       uint16
	Second       uint16
	MilliSeconds uint16
}

type CryptoApiBlob struct {
	DataSize uint32
	Data     []byte
}

type CertNameBlob struct {
	CryptoApiBlob
}

type KeyProvInfo struct {
	ContainerName *uint16
	ProvName      *uint16
	ProvType      uint32
	Flags         uint32
	CProvParam    uint32
	ProvParam     *KeyProvParam
	KeySpec       uint32
}

type KeyProvParam struct {
	Param    uint32
	Data     *uint8
	DataSize uint32
	Flags    uint32
}

type AlgorithmIdentifier struct {
	OID        *int8
	Parameters CryptoApiBlob
}

type CertExtension struct {
	OIDSize  *int8
	Critical int32
	Value    CryptoApiBlob
}

type CertExtensions struct {
	Capacity   uint32
	Extensions *CertExtension
}
