// Copyright 2019, 2022 Weald Technology Trading
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ens

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	string2eth "github.com/wealdtech/go-string2eth"
)

func TestName(t *testing.T) {
	dsRegistrant := common.HexToAddress("a303ddc620aa7d1390baccc8a495508b183fab59")
	dsController := common.HexToAddress("a303ddc620aa7d1390baccc8a495508b183fab59")
	dsResolver := common.HexToAddress("DaaF96c344f63131acadD0Ea35170E7892d3dfBA")
	dsExpiry := time.Unix(4741286688, 0)
	dsRegistrationInterval := 60 * time.Second

	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	name, err := NewName(client, "resolver.eth")
	require.Nil(t, err, "Failed to create name")

	registrant, err := name.Registrant()
	require.Nil(t, err, "Failed to obtain registrant")
	assert.Equal(t, dsRegistrant, registrant, "Failed to obtain correct registrant")

	controller, err := name.Controller()
	require.Nil(t, err, "Failed to obtain controller")
	assert.Equal(t, dsController, controller, "Failed to obtain correct controller")

	expiry, err := name.Expires()
	require.Nil(t, err, "Failed to obtain expiry")

	assert.Equal(t, dsExpiry, expiry, "Failed to obtain correct expiry")

	registrationInterval, err := name.RegistrationInterval()
	require.Nil(t, err, "Failed to obtain registration interval")
	assert.Equal(t, dsRegistrationInterval, registrationInterval, "Failed to obtain correct registration interval")

	resolverAddress, err := name.ResolverAddress()
	require.Nil(t, err, "Failed to obtain resolver address")
	assert.Equal(t, dsResolver, resolverAddress, "Failed to obtain correct resolver address")
}

func TestNameExpiry(t *testing.T) {
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	domain := unregisteredDomain(client)
	name, err := NewName(client, domain)
	require.Nil(t, err, "Failed to create name")
	_, err = name.Expires()
	assert.Equal(t, err.Error(), "not registered")
}

func TestNameReRegistration(t *testing.T) {
	registrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(registrant) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	name, err := NewName(client, "resolver.eth")
	require.Nil(t, err, "Failed to create name")

	// Register stage 1 - should fail as already registered
	opts, err := generateTxOpts(client, registrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	_, _, err = name.RegisterStageOne(registrant, opts)
	require.EqualError(t, err, "name is already registered")
}

func TestInvalidName(t *testing.T) {
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	_, err := NewName(client, "ab.eth")
	require.Equal(t, err.Error(), "name is not valid according to the rules of the registrar (too short, invalid characters, etc.)")
}

func TestNameRegistration(t *testing.T) {
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	registrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(registrant) {
		t.Skip()
	}
	domain := unregisteredDomain(client)
	name, err := NewName(client, domain)
	require.Nil(t, err, "Failed to create name")

	// Register stage 1
	opts, err := generateTxOpts(client, registrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, secret, err := name.RegisterStageOne(registrant, opts)
	require.Nil(t, err, "Failed to send stage one transaction")
	// Wait until mined
	waitForTransaction(client, tx.Hash())

	// Wait until ready to submit stage 2
	interval, err := name.RegistrationInterval()
	require.Nil(t, err, "Failed to obtain registration interval")
	time.Sleep(interval)
	// Sleep for 1 more minute to ensure we are over the interval
	time.Sleep(60 * time.Second)

	// Register stage 2
	opts, err = generateTxOpts(client, registrant, "0.1 Ether")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, err = name.RegisterStageTwo(registrant, secret, opts)
	require.Nil(t, err, "Failed to send stage two transaction")
	// Wait until mined
	waitForTransaction(client, tx.Hash())

	// Confirm registered
	registeredRegistrant, err := name.Registrant()
	require.Nil(t, err, "Failed to obtain registrant")
	assert.Equal(t, registrant, registeredRegistrant, "failed to register name")
}

func TestNameRegistrationStageTwoNoStageOne(t *testing.T) {
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	registrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(registrant) {
		t.Skip()
	}
	domain := unregisteredDomain(client)

	name, err := NewName(client, domain)
	require.Nil(t, err, "Failed to create name")

	// Register stage 2
	opts, err := generateTxOpts(client, registrant, "0.1 Ether")
	require.Nil(t, err, "Failed to generate transaction options")
	var secret [32]byte
	_, err = name.RegisterStageTwo(registrant, secret, opts)
	require.Equal(t, err.Error(), "stage 2 attempted prior to successful stage 1 transaction")
}

func TestNameRegistrationNoValue(t *testing.T) {
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	registrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(registrant) {
		t.Skip()
	}
	domain := unregisteredDomain(client)

	name, err := NewName(client, domain)
	require.Nil(t, err, "Failed to create name")

	// Register stage 1
	opts, err := generateTxOpts(client, registrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, secret, err := name.RegisterStageOne(registrant, opts)
	require.Nil(t, err, "Failed to send stage one transaction")
	// Wait until mined
	waitForTransaction(client, tx.Hash())

	// Wait until ready to submit stage 2
	interval, err := name.RegistrationInterval()
	require.Nil(t, err, "Failed to obtain registration interval")
	time.Sleep(interval)
	// Sleep for 1 more minute to ensure we are over the interval
	time.Sleep(60 * time.Second)

	// Register stage 2 - no value
	opts, err = generateTxOpts(client, registrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	_, err = name.RegisterStageTwo(registrant, secret, opts)
	assert.Equal(t, err.Error(), "not enough funds to cover minimum duration of 672h0m0s")
}

func TestNameRegistrationNoInterval(t *testing.T) {
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	registrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(registrant) {
		t.Skip()
	}
	domain := unregisteredDomain(client)

	name, err := NewName(client, domain)
	require.Nil(t, err, "Failed to create name")

	// Register stage 1
	opts, err := generateTxOpts(client, registrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, secret, err := name.RegisterStageOne(registrant, opts)
	require.Nil(t, err, "Failed to send stage one transaction")
	// Wait until mined
	waitForTransaction(client, tx.Hash())

	// Register stage 2 immediately - should fail
	opts, err = generateTxOpts(client, registrant, "0.1 Ether")
	require.Nil(t, err, "Failed to generate transaction options")
	_, err = name.RegisterStageTwo(registrant, secret, opts)
	require.NotNil(t, err, "No error when trying to register stage 2 immediately")
	assert.Equal(t, err.Error(), "too early to send second transaction")
}

func TestNameExtension(t *testing.T) {
	registrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(registrant) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	oldExpires, err := name.Expires()
	require.Nil(t, err, "Failed to obtain old expires")

	opts, err := generateTxOpts(client, registrant, "0.001Ether")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, err := name.ExtendRegistration(opts)
	require.Nil(t, err, "Failed to send transaction")
	// Wait until mined
	waitForTransaction(client, tx.Hash())
	// Confirm expiry has increased
	newExpires, err := name.Expires()
	require.Nil(t, err, "Failed to obtain new expires")
	assert.True(t, newExpires.After(oldExpires), "Failed to increase registration period")
}

func TestNameExtensionLowValue(t *testing.T) {
	registrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(registrant) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	opts, err := generateTxOpts(client, registrant, "1 wei")
	require.Nil(t, err, "Failed to generate transaction options")
	_, err = name.ExtendRegistration(opts)
	assert.Equal(t, err.Error(), "not enough funds to extend the registration")
}

func TestNameExtensionNotRegistered(t *testing.T) {
	registrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(registrant) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	domain := unregisteredDomain(client)
	name, err := NewName(client, domain)
	require.Nil(t, err, "Failed to create name")

	opts, err := generateTxOpts(client, registrant, "0.001Ether")
	require.Nil(t, err, "Failed to generate transaction options")
	_, err = name.ExtendRegistration(opts)
	assert.Equal(t, err.Error(), "name is not registered")
}

func TestNameSubdomainCreate(t *testing.T) {
	dsRegistrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(dsRegistrant) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")

	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	sub := unregisteredDomain(client)
	sub = strings.TrimSuffix(sub, ".eth")

	opts, err := generateTxOpts(client, dsRegistrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")

	tx, err := name.CreateSubdomain(sub, dsRegistrant, opts)
	require.Nil(t, err, "Failed to send transaction")
	// Wait until mined
	waitForTransaction(client, tx.Hash())

	// Confirm registrantship of the subdomain
	subdomain := fmt.Sprintf("%s.foobar5.eth", sub)

	registry, err := NewRegistry(client)
	require.Nil(t, err, "Failed to create registry")
	controller, err := registry.Owner(subdomain)
	require.Nil(t, err, "Failed to obtain subname's controller")
	assert.Equal(t, dsRegistrant, controller, "Unexpected controller for %s", subdomain)
}

func TestNameSubdomainCreateAlreadyExists(t *testing.T) {
	dsRegistrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(dsRegistrant) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")

	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	sub := "go-ens-test-1331354196"

	opts, err := generateTxOpts(client, dsRegistrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")

	_, err = name.CreateSubdomain(sub, dsRegistrant, opts)
	require.NotNil(t, err, "Failed to error when it should")
	assert.Equal(t, "that subdomain already exists", err.Error())
}

func TestSetController(t *testing.T) {
	dsRegistrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(dsRegistrant) {
		t.Skip()
	}
	dsController := common.HexToAddress("E195c59BCF26fD36c82d1C720860127A5c1c4040")
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")

	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	// Ensure that the registrant starts out as the controller
	controller, err := name.Controller()
	require.Nil(t, err, "Failed to obtain controller")
	require.Equal(t, dsRegistrant, controller, "Initial controller incorrect")

	// Set the controller
	opts, err := generateTxOpts(client, dsRegistrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, err := name.SetController(dsController, opts)
	require.Nil(t, err, "Failed to generate transaction")
	waitForTransaction(client, tx.Hash())

	// Confirm the controller is set
	controller, err = name.Controller()
	require.Nil(t, err, "Failed to obtain controller (2)")
	require.Equal(t, dsController, controller, "Updated controller incorrect")

	// Reset the controller role
	opts, err = generateTxOpts(client, dsRegistrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, err = name.SetController(dsRegistrant, opts)
	require.Nil(t, err, "Failed to generate transaction (2)")
	waitForTransaction(client, tx.Hash())

	// Confirm the controller is reset
	controller, err = name.Controller()
	require.Nil(t, err, "Failed to obtain controller (3)")
	require.Equal(t, dsRegistrant, controller, "Reset controller incorrect")
}

func TestSetControllerUnauthorised(t *testing.T) {
	dsRegistrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(dsRegistrant) {
		t.Skip()
	}
	dsThief := common.HexToAddress("E195c59BCF26fD36c82d1C720860127A5c1c4040")
	if !hasPrivateKey(dsThief) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")

	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	// Ensure that the registrant starts out as the controller
	controller, err := name.Controller()
	require.Nil(t, err, "Failed to obtain controller")
	require.Equal(t, dsRegistrant, controller, "Initial controller incorrect")

	// Try to set the controller
	opts, err := generateTxOpts(client, dsThief, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	_, err = name.SetController(dsRegistrant, opts)
	require.NotNil(t, err, "Failed to error when it should")
	assert.Equal(t, "not authorised to change the controller", err.Error())
}

func TestReclaim(t *testing.T) {
	dsRegistrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(dsRegistrant) {
		t.Skip()
	}
	dsController := common.HexToAddress("E195c59BCF26fD36c82d1C720860127A5c1c4040")
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")

	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	// Ensure that the registrant starts out as the controller
	controller, err := name.Controller()
	require.Nil(t, err, "Failed to obtain controller")
	require.Equal(t, dsRegistrant, controller, "Initial controller incorrect")

	// Set the controller
	opts, err := generateTxOpts(client, dsRegistrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, err := name.SetController(dsController, opts)
	require.Nil(t, err, "Failed to generate transaction")
	waitForTransaction(client, tx.Hash())

	// Confirm the controller is set
	controller, err = name.Controller()
	require.Nil(t, err, "Failed to obtain controller (2)")
	require.Equal(t, dsController, controller, "Updated controller incorrect")

	// Set the controller role
	opts, err = generateTxOpts(client, dsRegistrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, err = name.Reclaim(opts)
	require.Nil(t, err, "Failed to generate transaction (2)")
	waitForTransaction(client, tx.Hash())

	// Confirm the controller is reset
	controller, err = name.Controller()
	require.Nil(t, err, "Failed to obtain controller (3)")
	require.Equal(t, dsRegistrant, controller, "Reset controller incorrect")
}

func TestReclaimUnauthorised(t *testing.T) {
	dsRegistrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	dsThief := common.HexToAddress("E195c59BCF26fD36c82d1C720860127A5c1c4040")
	if !hasPrivateKey(dsThief) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")

	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	// Ensure that the registrant starts out as the controller
	controller, err := name.Controller()
	require.Nil(t, err, "Failed to obtain controller")
	require.Equal(t, dsRegistrant, controller, "Initial controller incorrect")

	// Try to reclaim
	opts, err := generateTxOpts(client, dsThief, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	_, err = name.Reclaim(opts)
	require.NotNil(t, err, "Failed to error when it should")
	assert.Equal(t, "not the registrant", err.Error())
}

func TestTransfer(t *testing.T) {
	dsRegistrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	if !hasPrivateKey(dsRegistrant) {
		t.Skip()
	}
	dsNewRegistrant := common.HexToAddress("E195c59BCF26fD36c82d1C720860127A5c1c4040")
	if !hasPrivateKey(dsNewRegistrant) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")

	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	// Ensure the existing registrant is correct
	registrant, err := name.Registrant()
	require.Nil(t, err, "Failed to obtain registrant")
	require.Equal(t, dsRegistrant, registrant, "Initial registrant incorrect")

	// Set the registrant
	opts, err := generateTxOpts(client, dsRegistrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, err := name.Transfer(dsNewRegistrant, opts)
	require.Nil(t, err, "Failed to generate transaction")
	waitForTransaction(client, tx.Hash())

	// Confirm the new registrant is set
	registrant, err = name.Registrant()
	require.Nil(t, err, "Failed to obtain registrant (2)")
	require.Equal(t, dsNewRegistrant, registrant, "Updated registrant incorrect")

	// Reset the registrant
	opts, err = generateTxOpts(client, dsNewRegistrant, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	tx, err = name.Transfer(dsRegistrant, opts)
	require.Nil(t, err, "Failed to generate transaction (2)")
	waitForTransaction(client, tx.Hash())

	// Confirm the registrant is reset
	registrant, err = name.Registrant()
	require.Nil(t, err, "Failed to obtain registrant (3)")
	require.Equal(t, dsRegistrant, registrant, "Reset registrant incorrect")
}

func TestTransferUnauthorised(t *testing.T) {
	dsRegistrant := common.HexToAddress("388Ea662EF2c223eC0B047D41Bf3c0f362142ad5")
	dsThief := common.HexToAddress("E195c59BCF26fD36c82d1C720860127A5c1c4040")
	if !hasPrivateKey(dsThief) {
		t.Skip()
	}
	client, _ := ethclient.Dial("https://ropsten.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")

	name, err := NewName(client, "foobar5.eth")
	require.Nil(t, err, "Failed to create name")

	// Ensure that the registrant starts out as the controller
	registrant, err := name.Registrant()
	require.Nil(t, err, "Failed to obtain registrant")
	require.Equal(t, dsRegistrant, registrant, "Initial registrant incorrect")

	// Try to steal
	opts, err := generateTxOpts(client, dsThief, "0")
	require.Nil(t, err, "Failed to generate transaction options")
	_, err = name.Transfer(dsThief, opts)
	require.NotNil(t, err, "Failed to error when it should")
	assert.Equal(t, "not the current registrant", err.Error())
}

func generateTxOpts(client *ethclient.Client, sender common.Address, valueStr string) (*bind.TransactOpts, error) {
	key, err := crypto.HexToECDSA(os.Getenv(fmt.Sprintf("PRIVATE_KEY_%x", sender)))
	if err != nil {
		return nil, fmt.Errorf("Failed to obtain private key for %x", sender)
	}
	signer := keySigner(big.NewInt(3), key)
	if signer == nil {
		return nil, errors.New("no signer")
	}

	value, err := string2eth.StringToWei(valueStr)
	if err != nil {
		return nil, err
	}

	curNonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		return nil, err
	}
	nonce := int64(curNonce)

	opts := &bind.TransactOpts{
		From:     sender,
		Signer:   signer,
		GasPrice: big.NewInt(10000000000),
		Value:    value,
		Nonce:    big.NewInt(0).SetInt64(nonce),
	}

	return opts, nil
}

func keySigner(chainID *big.Int, key *ecdsa.PrivateKey) bind.SignerFn {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		keyAddr := crypto.PubkeyToAddress(key.PublicKey)
		if address != keyAddr {
			return nil, errors.New("not authorized to sign this account")
		}
		return types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	}
}

func waitForTransaction(client *ethclient.Client, txHash common.Hash) {
	for {
		_, pending, err := client.TransactionByHash(context.Background(), txHash)
		if err == nil && !pending {
			return
		}
		time.Sleep(5 * time.Second)
	}
}

func unregisteredDomain(client *ethclient.Client) string {
	rand.Seed(time.Now().UTC().UnixNano())
	registry, _ := NewRegistry(client)
	for {
		// #nosec G404
		domain := fmt.Sprintf("go-ens-test-%d.eth", rand.Int31())
		controller, _ := registry.Owner(domain)
		if controller == UnknownAddress {
			return domain
		}
	}
}

func hasPrivateKey(address common.Address) bool {
	_, err := crypto.HexToECDSA(os.Getenv(fmt.Sprintf("PRIVATE_KEY_%x", address)))
	return err == nil
}
