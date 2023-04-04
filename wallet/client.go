package wallet

import (
	"context"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/types/tx"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	blockchainClient  pactus.BlockchainClient
	transactionClient pactus.TransactionClient
}

func newGRPCClient(rpcEndpoint string) (*grpcClient, error) {
	conn, err := grpc.Dial(rpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClient{
		blockchainClient:  pactus.NewBlockchainClient(conn),
		transactionClient: pactus.NewTransactionClient(conn),
	}, nil
}

func (c *grpcClient) getStamp() (hash.Stamp, error) {
	info, err := c.blockchainClient.GetBlockchainInfo(context.Background(),
		&pactus.GetBlockchainInfoRequest{})
	if err != nil {
		return hash.Stamp{}, err
	}
	h, _ := hash.FromBytes(info.LastBlockHash)
	return h.Stamp(), nil
}

func (c *grpcClient) getAccount(addr crypto.Address) (*pactus.AccountInfo, error) {
	res, err := c.blockchainClient.GetAccount(context.Background(),
		&pactus.GetAccountRequest{Address: addr.String()})
	if err != nil {
		return nil, err
	}
	return res.Account, nil
}

func (c *grpcClient) getValidator(addr crypto.Address) (*pactus.ValidatorInfo, error) {
	res, err := c.blockchainClient.GetValidator(context.Background(),
		&pactus.GetValidatorRequest{Address: addr.String()})
	if err != nil {
		return nil, err
	}
	return res.Validator, nil
}

func (c *grpcClient) sendTx(tx *tx.Tx) (tx.ID, error) {
	data, err := tx.Bytes()
	if err != nil {
		return hash.UndefHash, err
	}
	res, err := c.transactionClient.SendRawTransaction(context.Background(), &pactus.SendRawTransactionRequest{
		Data: data,
	})

	if err != nil {
		return hash.UndefHash, err
	}

	return hash.FromBytes(res.Id)
}

func (c *grpcClient) getTransaction(id tx.ID) (*pactus.GetTransactionResponse, error) {
	res, err := c.transactionClient.GetTransaction(context.Background(), &pactus.GetTransactionRequest{
		Id:        id.Bytes(),
		Verbosity: pactus.TransactionVerbosity_TRANSACTION_INFO,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}