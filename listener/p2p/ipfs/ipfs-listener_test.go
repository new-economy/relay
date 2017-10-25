/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package ipfs_test

import (
	"encoding/json"
	"github.com/Loopring/ringminer/log"
	"github.com/Loopring/ringminer/test"
	"github.com/Loopring/ringminer/types"
	"github.com/ipfs/go-ipfs-api"
	"math/big"
	"testing"
)

var testParams *test.TestParams

func init() {
	testParams = test.LoadConfigAndGenerateTestParams()
	log.Infof("contract address:%s", testParams.ImplAddress.Hex())
	log.Infof("delegate address:%s", testParams.DelegateAddress.Hex())
	log.Infof("register address:%s", testParams.TokenRegistryAddress.Hex())
}

func TestPrepareAccount(t *testing.T) {
	testParams.PrepareTestData()
}

func TestCheckAllowance(t *testing.T) {
	//testParams.CheckAllowance(test.TokenAddressA, "0xb5fab0b11776aad5ce60588c16bd59dcfd61a1c2")
	//testParams.CheckAllowance(test.TokenAddressA, "0x48ff2269e58a373120FFdBBdEE3FBceA854AC30A")
	//testParams.CheckAllowance(test.TokenAddressB, "0xb5fab0b11776aad5ce60588c16bd59dcfd61a1c2")
	//testParams.CheckAllowance(test.TokenAddressB, "0x48ff2269e58a373120FFdBBdEE3FBceA854AC30A")
	testParams.IsTestDataReady()
}

/*
fullhash=0x521fe41301c4b567b22212fc8b464396a719786ca559bd659f6f146cef59a546 recipient=0x937Ff659c8a9D85aAC39dfA84c4b49Bb7C9b226E
INFO [10-23|16:31:24] Submitted transaction                    fullhash=0x866d5a7c44a0a12088681734f1724334c4a0c152ed760bc995f59fe1676af310 recipient=0x937Ff659c8a9D85aAC39dfA84c4b49Bb7C9b226E
INFO [10-23|16:31:24] Submitted transaction                    fullhash=0xcd172bda3ab2680e70761d1bd70711694501a3ea463bb514f326aeeacb647ee3 recipient=0x8711aC984e6ce2169a2a6bd83eC15332c366Ee4F
INFO [10-23|16:31:24] Submitted transaction                    fullhash=0x967ec954d366a6f8200747d18576834bbc30bbcfa37d63943948520d18c0e189 recipient=0x8711aC984e6ce2169a2a6bd83eC15332c366Ee4F
*/
func TestOrdersOfRing(t *testing.T) {
	sh := shell.NewLocalShell()

	suffix := "0"

	//scheme 1:MarginSplitPercentage = 0
	amountS1, _ := new(big.Int).SetString("1"+suffix, 0)
	amountB1, _ := new(big.Int).SetString("10"+suffix, 0)
	order1 := test.CreateOrder(
		types.HexToAddress(test.TokenAddressA),
		types.HexToAddress(test.TokenAddressB),
		testParams.ImplAddress,
		amountS1,
		amountB1,
		types.Hex2Bytes("11293da8fdfe3898eae7637e429e7e93d17d0d8293a4d1b58819ac0ca102b446"),
		types.HexToAddress("0xb5fab0b11776aad5ce60588c16bd59dcfd61a1c2"),
	)
	data1, _ := json.Marshal(order1)
	pubMessage(sh, string(data1))

	amountS2, _ := new(big.Int).SetString("20"+suffix, 0)
	amountB2, _ := new(big.Int).SetString("1"+suffix, 0)
	order2 := test.CreateOrder(
		types.HexToAddress(test.TokenAddressB),
		types.HexToAddress(test.TokenAddressA),
		testParams.ImplAddress,
		amountS2,
		amountB2,
		types.Hex2Bytes("07ae9ee56203d29171ce3de536d7742e0af4df5b7f62d298a0445d11e466bf9e"),
		types.HexToAddress("0x48ff2269e58a373120FFdBBdEE3FBceA854AC30A"),
	)
	data2, _ := json.Marshal(order2)
	pubMessage(sh, string(data2))
}

func pubMessage(sh *shell.Shell, data string) {
	topic := testParams.Config.Ipfs.Topic
	err := sh.PubSubPublish(topic, data)
	if err != nil {
		panic(err.Error())
	}
}

func TestIsTestDataReady(t *testing.T) {
	testParams.IsTestDataReady()
}

func TestPrepareTestData(t *testing.T) {
	//testParams.PrepareTestData()
	//var res string
	//testParams.Registry.CanSubmit.Call(&res, "pending", types.HexToHash("0x6114f07229a87ae4cd4c61d1888c34f3c45266b25f2136f3a0068e6b0ccd52d4"), types.HexToAddress("0x48ff2269e58a373120FFdBBdEE3FBceA854AC30A"))
	//println(res)
	sh := shell.NewLocalShell()

	order1 := `{"protocol":"0x29d4178372d890e3127d35c3f49ee5ee215d6fe8","tokenS":"0x8711ac984e6ce2169a2a6bd83ec15332c366ee4f","tokenB":"0x937ff659c8a9d85aac39dfa84c4b49bb7c9b226e","amountS":"0xc8","amountB":"0xa","timestamp":"0x59ef0cc8","ttl":"0x2710","salt":"0x3e8","lrcFee":"0x64","buyNoMoreThanAmountB":false,"marginSplitPercentage":0,"v":27,"r":"0xecdfe5d96346e1a4fffce7a63fe0c8ff6111b13c3c387a296cdc6d9a10599fb0","s":"0x18640bbb9ccc6b667a05abcd349531b58211084b33fbb73270f1eb1861d6559a","owner":"0x48ff2269e58a373120ffdbbdee3fbcea854ac30a","hash":"0x9b7857b006236a148e70e8b07adf6347610a7d1beb88328810528d98f20496e8"}`
	order2 := `{"protocol":"0x29d4178372d890e3127d35c3f49ee5ee215d6fe8","tokenS":"0x937ff659c8a9d85aac39dfa84c4b49bb7c9b226e","tokenB":"0x8711ac984e6ce2169a2a6bd83ec15332c366ee4f","amountS":"0xa","amountB":"0x64","timestamp":"0x59ef0cc8","ttl":"0x2710","salt":"0x3e8","lrcFee":"0x64","buyNoMoreThanAmountB":false,"marginSplitPercentage":0,"v":27,"r":"0xe4c79971b1949223b185101e3bd890f5b5d236d0f6c067bb1e3f36fa3784e79c","s":"0x0512ac6bd868bb92a5e4270ba3cfd3856d21ed5092d4b23b54f5b854fb0777df","owner":"0xb5fab0b11776aad5ce60588c16bd59dcfd61a1c2","hash":"0xdf2d02d7533d4c2faea71588df3e41f8199853cd97ea5eeb29ee0a5c643dd953"}`

	pubMessage(sh, order1)
	pubMessage(sh, order2)

}