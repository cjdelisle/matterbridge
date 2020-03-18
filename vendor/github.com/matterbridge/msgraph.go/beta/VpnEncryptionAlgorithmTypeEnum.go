// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VpnEncryptionAlgorithmType undocumented
type VpnEncryptionAlgorithmType int

const (
	// VpnEncryptionAlgorithmTypeVAes256 undocumented
	VpnEncryptionAlgorithmTypeVAes256 VpnEncryptionAlgorithmType = 0
	// VpnEncryptionAlgorithmTypeVDes undocumented
	VpnEncryptionAlgorithmTypeVDes VpnEncryptionAlgorithmType = 1
	// VpnEncryptionAlgorithmTypeVTripleDes undocumented
	VpnEncryptionAlgorithmTypeVTripleDes VpnEncryptionAlgorithmType = 2
	// VpnEncryptionAlgorithmTypeVAes128 undocumented
	VpnEncryptionAlgorithmTypeVAes128 VpnEncryptionAlgorithmType = 3
	// VpnEncryptionAlgorithmTypeVAes128Gcm undocumented
	VpnEncryptionAlgorithmTypeVAes128Gcm VpnEncryptionAlgorithmType = 4
	// VpnEncryptionAlgorithmTypeVAes256Gcm undocumented
	VpnEncryptionAlgorithmTypeVAes256Gcm VpnEncryptionAlgorithmType = 5
)

// VpnEncryptionAlgorithmTypePAes256 returns a pointer to VpnEncryptionAlgorithmTypeVAes256
func VpnEncryptionAlgorithmTypePAes256() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVAes256
	return &v
}

// VpnEncryptionAlgorithmTypePDes returns a pointer to VpnEncryptionAlgorithmTypeVDes
func VpnEncryptionAlgorithmTypePDes() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVDes
	return &v
}

// VpnEncryptionAlgorithmTypePTripleDes returns a pointer to VpnEncryptionAlgorithmTypeVTripleDes
func VpnEncryptionAlgorithmTypePTripleDes() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVTripleDes
	return &v
}

// VpnEncryptionAlgorithmTypePAes128 returns a pointer to VpnEncryptionAlgorithmTypeVAes128
func VpnEncryptionAlgorithmTypePAes128() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVAes128
	return &v
}

// VpnEncryptionAlgorithmTypePAes128Gcm returns a pointer to VpnEncryptionAlgorithmTypeVAes128Gcm
func VpnEncryptionAlgorithmTypePAes128Gcm() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVAes128Gcm
	return &v
}

// VpnEncryptionAlgorithmTypePAes256Gcm returns a pointer to VpnEncryptionAlgorithmTypeVAes256Gcm
func VpnEncryptionAlgorithmTypePAes256Gcm() *VpnEncryptionAlgorithmType {
	v := VpnEncryptionAlgorithmTypeVAes256Gcm
	return &v
}