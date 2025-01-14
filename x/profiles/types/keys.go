package types

//DONTCOVER

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionSaveProfile               = "save_profile"
	ActionDeleteProfile             = "delete_profile"
	ActionRequestDtag               = "request_dtag"
	ActionAcceptDtagTransfer        = "accept_dtag_request"
	ActionRefuseDTagTransferRequest = "refuse_dtag_request"
	ActionCancelDTagTransferRequest = "cancel_dtag_request"

	QuerierRoute              = ModuleName
	QueryProfile              = "profile"
	QueryIncomingDTagRequests = "incoming-dtag-requests"
	QueryParams               = "params"

	DoNotModify = "[do-not-modify]"
)

var (
	DTagPrefix                 = []byte("dtag")
	DTagTransferRequestsPrefix = []byte("transfer_requests")
)

// DTagStoreKey turns a DTag into the key used to store the address associated with it into the store
func DTagStoreKey(dtag string) []byte {
	return append(DTagPrefix, []byte(dtag)...)
}

// DtagTransferRequestStoreKey turns an address to a key used to store a transfer request into the profiles store
func DtagTransferRequestStoreKey(address string) []byte {
	return append(DTagTransferRequestsPrefix, address...)
}
