package wallet

import (
	rpcClient "automato/rpc_client"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

var Wallet EOA

type EOA struct {
	SignerAddress common.Address
	Signer        types.Signer
	PrivateKey    *ecdsa.PrivateKey
	signerMutex   *sync.Mutex
}

func InitializeEOA() {

	wallet := os.Getenv("WALLET_ADDRESS")
	wallet, err := toChecksumAddress(wallet)

	private_key := os.Getenv("PRIVATE_KEY")

	pk := initializePrivateKey(wallet, private_key)

	if err != nil {
		panic("To checksumAddress failed, use a correct PublicKey")
	}

	chainId := os.Getenv("CHAIN_ID")

	newBigInt := new(big.Int)
	chainIdBigInt, ok := newBigInt.SetString(chainId, 0)

	if !ok {
		fmt.Println("Error when converting string to big int during chainId initialization")
		return
	}

	Wallet = EOA{
		SignerAddress: common.HexToAddress(wallet),
		Signer:        types.LatestSignerForChainID(chainIdBigInt),
		PrivateKey:    pk,
		signerMutex:   &sync.Mutex{},
	}

}

//Initialize a new private key and wipe the input after usage
func initializePrivateKey(walletChecksumAddress string, privateKey string) *ecdsa.PrivateKey {
	//Initialize a key variable
	var walletKey *ecdsa.PrivateKey

	if privateKey[:2] == "0x" {
		privateKey = privateKey[2:]
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)

	if err != nil {
		errString := fmt.Sprintf("Incorrect or invalid private key for %s. Please check your wallet address/private key and try again.\n", walletChecksumAddress)
		panic(errString)
	} else {
		//Set user wallet private key
		walletKey = ecdsaPrivateKey
	}
	// Return the wallet key
	return walletKey
}

//Convert a hex address to checksum address
func toChecksumAddress(address string) (string, error) {

	//Check that the address is a valid Ethereum address
	re1 := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re1.MatchString(address) {
		return "", fmt.Errorf("given address '%s' is not a valid Ethereum Address", address)
	}

	//Convert the address to lowercase
	re2 := regexp.MustCompile("^0x")
	address = re2.ReplaceAllString(address, "")
	address = strings.ToLower(address)

	//Convert address to sha3 hash
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(address))
	sum := hasher.Sum(nil)
	addressHash := fmt.Sprintf("%x", sum)
	addressHash = re2.ReplaceAllString(addressHash, "")

	//Compile checksum address
	checksumAddress := "0x"
	for i := 0; i < len(address); i++ {
		indexedValue, err := strconv.ParseInt(string(rune(addressHash[i])), 16, 32)
		if err != nil {
			fmt.Println("Error when parsing addressHash during checksum conversion", err)
			return "", err
		}
		if indexedValue > 7 {
			checksumAddress += strings.ToUpper(string(address[i]))
		} else {
			checksumAddress += string(address[i])
		}
	}

	//Return the checksummed address
	return checksumAddress, nil
}

func (e *EOA) SignAndSendTx(toAddress *common.Address, calldata []byte, msgValue *big.Int, gas uint64, gasTipCap *big.Int, gasFeeCap *big.Int) {

	//lock the mutex so only one tx can be sent at a time. The most recently sent transaction must be confirmed
	//before the next transaction can be sent
	e.signerMutex.Lock()
	//get wallet nonce
	nonce, err := rpcClient.HTTPClient.NonceAt(context.Background(), e.SignerAddress, nil)
	if err != nil {
		fmt.Println(err)
		//TODO: In the future, handle errors gracefully
		os.Exit(1)
	}

	//initialize TXData

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   e.Signer.ChainID(),
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gas,
		To:        toAddress,
		Value:     msgValue,
		Data:      calldata,
	})

	signedTx, err := types.SignTx(tx, e.Signer, e.PrivateKey)
	if err != nil {
		fmt.Println(err)
		//TODO: In the future, handle errors gracefully
		os.Exit(1)
	}

	//send the transaction
	txErr := rpcClient.HTTPClient.SendTransaction(context.Background(), signedTx)
	if txErr != nil {
		fmt.Println(err)
		//TODO: In the future, handle errors gracefully
		os.Exit(1)
	}

	//wait for the tx to complete
	WaitForTransactionToComplete(signedTx.Hash())

	//unlock the signer mutex
	e.signerMutex.Unlock()
}

func WaitForTransactionToComplete(txHash common.Hash) *types.Transaction {
	for {
		confirmedTx, pending, err := rpcClient.HTTPClient.TransactionByHash(context.Background(), txHash)
		if err != nil {
			fmt.Println("Err when getting transaction by hash", err)
			//TODO: In the future, handle errors gracefully
			os.Exit(1)
		}
		if !pending {
			return confirmedTx
		}

		time.Sleep(time.Second * time.Duration(1))
	}
}
