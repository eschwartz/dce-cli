// Created by interfacer; DO NOT EDIT

package util

import (
	"github.com/Optum/dce-cli/client/operations"
	"github.com/go-openapi/runtime"
)

// APIer is an interface generated for "github.com/Optum/dce-cli/client/operations.Client".
type APIer interface {
	DeleteAccountsID(*operations.DeleteAccountsIDParams, runtime.ClientAuthInfoWriter) (*operations.DeleteAccountsIDNoContent, error)
	DeleteLeases(*operations.DeleteLeasesParams, runtime.ClientAuthInfoWriter) (*operations.DeleteLeasesOK, error)
	DeleteLeasesID(*operations.DeleteLeasesIDParams, runtime.ClientAuthInfoWriter) (*operations.DeleteLeasesIDOK, error)
	GetAccounts(*operations.GetAccountsParams, runtime.ClientAuthInfoWriter) (*operations.GetAccountsOK, error)
	GetAccountsID(*operations.GetAccountsIDParams, runtime.ClientAuthInfoWriter) (*operations.GetAccountsIDOK, error)
	GetAuth(*operations.GetAuthParams) (*operations.GetAuthOK, error)
	GetAuthFile(*operations.GetAuthFileParams) (*operations.GetAuthFileOK, error)
	GetLeases(*operations.GetLeasesParams, runtime.ClientAuthInfoWriter) (*operations.GetLeasesOK, error)
	GetLeasesID(*operations.GetLeasesIDParams, runtime.ClientAuthInfoWriter) (*operations.GetLeasesIDOK, error)
	GetUsage(*operations.GetUsageParams, runtime.ClientAuthInfoWriter) (*operations.GetUsageOK, error)
	PostAccounts(*operations.PostAccountsParams, runtime.ClientAuthInfoWriter) (*operations.PostAccountsCreated, error)
	PostLeases(*operations.PostLeasesParams, runtime.ClientAuthInfoWriter) (*operations.PostLeasesCreated, error)
	PostLeasesAuth(*operations.PostLeasesAuthParams, runtime.ClientAuthInfoWriter) (*operations.PostLeasesAuthCreated, error)
	PostLeasesIDAuth(*operations.PostLeasesIDAuthParams, runtime.ClientAuthInfoWriter) (*operations.PostLeasesIDAuthCreated, error)
	PutAccountsID(*operations.PutAccountsIDParams, runtime.ClientAuthInfoWriter) (*operations.PutAccountsIDOK, error)
	SetTransport(runtime.ClientTransport)
}
