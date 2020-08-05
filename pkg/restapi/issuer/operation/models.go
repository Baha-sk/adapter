/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"encoding/json"

	"github.com/hyperledger/aries-framework-go/pkg/client/outofband"

	adaptervc "github.com/trustbloc/edge-adapter/pkg/vc"
)

// ProfileDataRequest req for profile creation.
type ProfileDataRequest struct {
	ID                  string   `json:"id,omitempty"`
	Name                string   `json:"name"`
	SupportedVCContexts []string `json:"supportedVCContexts"`
	URL                 string   `json:"url"`
}

// WalletConnect response from wallet.
type WalletConnect struct {
	Resp json.RawMessage `json:"walletResp,omitempty"`
}

// txnData contains session data.
type txnData struct {
	IssuerID          string                `json:"issuerID,omitempty"`
	State             string                `json:"state,omitempty"`
	DIDCommInvitation *outofband.Invitation `json:"didCommInvitation,omitempty"`
}

// CHAPIRequest wallet chapi request.
type CHAPIRequest struct {
	Query             *CHAPIQuery           `json:"query,omitempty"`
	DIDCommInvitation *outofband.Invitation `json:"invitation,omitempty"`
	Manifest          json.RawMessage       `json:"manifest,omitempty"`
}

// CHAPIQuery chapi query type data.
type CHAPIQuery struct {
	Type string `json:"type,omitempty"`
}

// ValidateConnectResp response from validate connect api.
type ValidateConnectResp struct {
	RedirectURL string `json:"redirectURL,omitempty"`
}

// AuthorizationCredentialReq authorization credential request from wallet.
type AuthorizationCredentialReq struct {
	SubjectDID string            `json:"subjectDID,omitempty"`
	RPDIDDoc   *adaptervc.DIDDoc `json:"requestingPartyDIDDoc,omitempty"`
}

// AuthorizationCredentialHandle stores authorization credential related data.
type AuthorizationCredentialHandle struct {
	ID               string `json:"id,omitempty"`
	IssuerDID        string `json:"issuerDID,omitempty"`
	SubjectDID       string `json:"subjectDID,omitempty"`
	RPDID            string `json:"rpDID,omitempty"`
	UserConnectionID string `json:"userConnectionID,omitempty"`
	RPConnectionID   string `json:"rpConnectionID,omitempty"`
	Token            string `json:"token,omitempty"`
	IssuerID         string `json:"issuerID,omitempty"`
}

// UserConnectionMapping stores mapping between the connectionID and issuer.
type UserConnectionMapping struct {
	ConnectionID string `json:"connectionID,omitempty"`
	IssuerID     string `json:"issuerID,omitempty"`
	Token        string `json:"token,omitempty"`
}

// IssuerVCReq request to issuer for user data.
type IssuerVCReq struct {
	Token string `json:"token,omitempty"`
}
