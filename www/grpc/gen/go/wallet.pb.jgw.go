// Code generated by protoc-gen-jrpc-gateway. DO NOT EDIT.
// source: wallet.proto

/*
Package pactus is a reverse proxy.

It translates gRPC into JSON-RPC 2.0
*/
package pactus

import (
	"context"
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
)

type WalletJsonRpcService struct {
	client WalletClient
}

func NewWalletJsonRpcService(client WalletClient) WalletJsonRpcService {
	return WalletJsonRpcService{
		client: client,
	}
}

func (s *WalletJsonRpcService) Methods() map[string]func(ctx context.Context, message json.RawMessage) (any, error) {
	return map[string]func(ctx context.Context, params json.RawMessage) (any, error){

		"pactus.wallet.create_wallet": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(CreateWalletRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.CreateWallet(ctx, req)
		},

		"pactus.wallet.load_wallet": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(LoadWalletRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.LoadWallet(ctx, req)
		},

		"pactus.wallet.unload_wallet": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(UnloadWalletRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.UnloadWallet(ctx, req)
		},

		"pactus.wallet.lock_wallet": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(LockWalletRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.LockWallet(ctx, req)
		},

		"pactus.wallet.unlock_wallet": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(UnlockWalletRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.UnlockWallet(ctx, req)
		},

		"pactus.wallet.get_total_balance": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetTotalBalanceRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetTotalBalance(ctx, req)
		},

		"pactus.wallet.sign_raw_transaction": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(SignRawTransactionRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.SignRawTransaction(ctx, req)
		},

		"pactus.wallet.get_validator_address": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(GetValidatorAddressRequest)
			err := protojson.Unmarshal(data, req)
			if err != nil {
				return nil, err
			}
			return s.client.GetValidatorAddress(ctx, req)
		},
	}
}
