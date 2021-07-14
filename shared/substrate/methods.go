// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

// An available method on the substrate chain
type Method string

var AddRelayerMethod Method = BridgePalletName + ".add_relayer"
var SetResourceMethod Method = BridgePalletName + ".set_resource"
var SetThresholdMethod Method = BridgePalletName + ".set_threshold"
var WhitelistChainMethod Method = BridgePalletName + ".whitelist_chain"
var ExampleTransferNativeMethod Method = "ChainBridgeHandler.transfer_native"
var ExampleTransferErc721Method Method = "ChainBridgeHandler.transfer_erc721"
var ExampleTransferHashMethod Method = "ChainBridgeHandler.transfer_hash"
var ExampleMintErc721Method Method = "ChainBridgeHandler.mint_erc721"
var ExampleTransferMethod Method = "ChainBridgeHandler.transfer"
var ExampleRemarkMethod Method = "ChainBridgeHandler.remark"
var Erc721MintMethod Method = "ChainBridgeErc721.mint"
var SudoMethod Method = "Sudo.sudo"
