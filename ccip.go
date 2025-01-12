package ens

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/wealdtech/go-ens/v3/contracts/offchainresolver"
	"github.com/wealdtech/go-ens/v3/contracts/universalresolver"
)

var Bytes, _ = abi.NewType("bytes", "", nil)

func getCcipReadError(err error) (bool, string) {
	var jsonErr, ok = err.(rpc.DataError)
	if !ok {
		return false, ""
	}
	errData, ok := jsonErr.ErrorData().(string)
	if !ok {
		return false, ""
	}
	return (len(errData) >= 10 && errData[:10] == offchainLookupSignature), errData
}

func ccipRead(backend bind.ContractBackend, resolverAddr common.Address, revertData string) ([]byte, error) {
	hexBytes, err := hex.DecodeString(strings.TrimPrefix(revertData, "0x"))
	if err != nil {
		return nil, err
	}
	uAbi, err := universalresolver.ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	if len(revertData) < 5 {
		return nil, errors.New("invalid revert data")
	}

	// Extracting error details from the revert data
	var sig [4]byte
	copy(sig[:], hexBytes[:4])
	abiErr, err := uAbi.ErrorByID(sig)
	if err != nil {
		return nil, err
	}
	errArgs, err := abiErr.Inputs.Unpack(hexBytes[4:])
	if err != nil {
		return nil, err
	}

	sender := errArgs[0].(common.Address)
	urls := errArgs[1].([]string)
	calldata := common.Bytes2Hex(errArgs[2].([]byte))
	calldataHex := fmt.Sprintf("0x%s", calldata)
	callback := errArgs[3].([4]byte)
	extraData := common.Bytes2Hex(errArgs[4].([]byte))
	extraDataHex := fmt.Sprintf("0x%s", extraData)

	// Fetching data from external source using CCIP
	resp, err := ccipFetch(sender, calldataHex, urls)
	if err != nil || len(resp) == 0 {
		return nil, errors.New("unregistered name")
	}

	oAbi, err := offchainresolver.ContractMetaData.GetAbi()
	if err != nil {
		return nil, errors.New("no address")
	}
	// all callback has the same input/output, even if the method name is different
	m := oAbi.Methods["resolveWithProof"]

	args, err := m.Inputs.Pack(common.FromHex(resp), common.FromHex(extraDataHex))
	if err != nil {
		return nil, errors.New("no address")
	}

	encodedResp, err := backend.CallContract(context.Background(), ethereum.CallMsg{
		To:   &resolverAddr,
		Data: append(callback[:], args...),
	}, nil)

	if err != nil {
		return nil, err
	}

	outputs, err := m.Outputs.Unpack(encodedResp)
	if err != nil || len(outputs) == 0 {
		return nil, err
	}
	return outputs[0].([]byte), nil
}

func ccipFetch(sender common.Address, data string, urls []string) (result string, err error) {
	for _, url := range urls {
		method := "POST"
		if strings.Contains(url, "{data}") {
			method = "GET"
		}

		body := []byte{}
		if method == "POST" {
			body, err = json.Marshal(map[string]interface{}{
				"data":   data,
				"sender": sender,
			})
			if err != nil {
				return "", err
			}
		}

		url = strings.ReplaceAll(url, "{sender}", sender.String())
		url = strings.ReplaceAll(url, "{data}", data)
		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()

		var responseData map[string]interface{}
		var result string
		if strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
			err = json.NewDecoder(resp.Body).Decode(&responseData)
			if err != nil {
				continue
			}
			var ok bool
			result, ok = responseData["data"].(string)
			if !ok {
				continue
			}
		} else {
			responseBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			result = string(responseBytes)
		}

		if resp.StatusCode != http.StatusOK {
			continue
		}

		return result, nil
	}

	return "", err
}
