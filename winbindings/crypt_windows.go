package wincert

import (
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer
var _ syscall.CertContext

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

var (
//	OIG_RSA_RSA1SHA = syscall.StringBytePtr("1.2.840.113549.1.1.5")
)

//sys	CertGetNameString(ctx *syscall.CertContext, typeName uint32, flags uint32, paraType uintptr, outBuffer *uint16, bufferSize uint32) (outSize uint32, err error) = crypt32.CertGetNameStringW
//sys	CertCreateSelfSignedCertificate(cryptProvOrCryptKey syscall.Handle, subjectIssuer *CertNameBlob, flags uint32, keyProvInfo *KeyProvInfo, signAlgorithm *AlgorithmIdentifier, startTime *SystemTime, endTime *SystemTime, extensions *CertExtensions) (certContext *syscall.CertContext, err error) [failretval == nil] = crypt32.CertCreateSelfSignCertificate
//sys	CertStrToName(encodingType uint32, strX509 *uint16, strType uint32, reserved uintptr, outBuffer *byte, bufferSize *uint32, outError **uint16)(err error) = crypt32.CertStrToNameW
//sys	CryptGenKey(cryptProv syscall.Handle, algID uint32, flags uint32, hKey *syscall.Handle)(err error) = advapi32.CryptGenKey
//sys	CertSetCertificateContextProperty(certCtx *syscall.CertContext, propId uint32, flags uint32, propertyData uintptr)(err error) = crypt32.CertSetCertificateContextProperty

//sys	PfxExportCertStore(storeHandle syscall.Handle, pfxBlob *CryptoApiBlob, password *uint16, data uintptr, flags uint32)(err error) = crypt32.PFXExportCertStoreEx
