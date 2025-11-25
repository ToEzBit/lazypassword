package models

type Vault struct {
	ID                   string
	WorkSpaceName        string
	WorkSpaceDescription string
	Credentials          []Credential
}

type Credential struct {
	ID       string
	AcountId string
	Password string
	App      string
	Url      string
	note     string
}

type VaultData struct {
	Vaults []Vault
}

type VaultWithoutCredentails struct {
	ID                   string
	WorkSpaceName        string
	WorkSpaceDescription string
}
