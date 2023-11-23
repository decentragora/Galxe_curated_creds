// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dagoramembershipsabi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DagoraMembershipsV1Membership is an auto generated low-level Go binding around an user-defined struct.
type DagoraMembershipsV1Membership struct {
	Tier       uint8
	Member     common.Address
	TokenId    *big.Int
	Expiration *big.Int
}

// DagoramembershipsabiMetaData contains all meta data concerning the Dagoramembershipsabi contract.
var DagoramembershipsabiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"DelegateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"DelegateRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldDelegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDelegatee\",\"type\":\"address\"}],\"name\":\"DelegateSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"FreeMembershipClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"MembershipCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"MembershipGifted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"MembershipPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"MembershipRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDagoraMembershipsV1.Tiers\",\"name\":\"newTier\",\"type\":\"uint8\"}],\"name\":\"MembershipUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"newTier\",\"type\":\"uint8\"}],\"name\":\"_getUpgradePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"addDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"addressTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelMembership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dagoraRenewPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dagoraTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dagorianPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecclesiaPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecclesiaRenewPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"experation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getExpiration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getMembership\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"internalType\":\"structDagoraMembershipsV1.Membership\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getMembershipTier\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_durationInMonths\",\"type\":\"uint96\"},{\"internalType\":\"uint8\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"getMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_newDuration\",\"type\":\"uint96\"},{\"internalType\":\"uint8\",\"name\":\"currentTier\",\"type\":\"uint8\"}],\"name\":\"getRenewalPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenDelegates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"durationInMonths\",\"type\":\"uint96\"}],\"name\":\"giftExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"durationInMonths\",\"type\":\"uint96\"}],\"name\":\"giftMembership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"giftUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hoplitePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hopliteRenewPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_dagoraTreasury\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addrs\",\"type\":\"address\"}],\"name\":\"isOwnerOrDelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isOwnerOrDelegate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isValidMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"memberships\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"_durationInMonths\",\"type\":\"uint96\"}],\"name\":\"mintMembership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"percelsiaPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"percelsiaRenewPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"slot\",\"type\":\"uint8\"}],\"name\":\"removeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_durationInMonths\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"renewMembership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dagoraTreasury\",\"type\":\"address\"}],\"name\":\"setDagoraTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setDagorianPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setDagorianRenewPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_discount\",\"type\":\"uint256\"}],\"name\":\"setDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setEcclesiaPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setEcclesiaRenewPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setHoplitePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setHopliteRenewPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPercelsiaPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPercelsiaRenewPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyAddress\",\"type\":\"address\"}],\"name\":\"setProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oldDelegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newDelegate\",\"type\":\"address\"}],\"name\":\"swapDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenDelegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newTier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"oldTier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"upgradeMembership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DagoramembershipsabiABI is the input ABI used to generate the binding from.
// Deprecated: Use DagoramembershipsabiMetaData.ABI instead.
var DagoramembershipsabiABI = DagoramembershipsabiMetaData.ABI

// Dagoramembershipsabi is an auto generated Go binding around an Ethereum contract.
type Dagoramembershipsabi struct {
	DagoramembershipsabiCaller     // Read-only binding to the contract
	DagoramembershipsabiTransactor // Write-only binding to the contract
	DagoramembershipsabiFilterer   // Log filterer for contract events
}

// DagoramembershipsabiCaller is an auto generated read-only Go binding around an Ethereum contract.
type DagoramembershipsabiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagoramembershipsabiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DagoramembershipsabiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagoramembershipsabiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DagoramembershipsabiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DagoramembershipsabiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DagoramembershipsabiSession struct {
	Contract     *Dagoramembershipsabi // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DagoramembershipsabiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DagoramembershipsabiCallerSession struct {
	Contract *DagoramembershipsabiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DagoramembershipsabiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DagoramembershipsabiTransactorSession struct {
	Contract     *DagoramembershipsabiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DagoramembershipsabiRaw is an auto generated low-level Go binding around an Ethereum contract.
type DagoramembershipsabiRaw struct {
	Contract *Dagoramembershipsabi // Generic contract binding to access the raw methods on
}

// DagoramembershipsabiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DagoramembershipsabiCallerRaw struct {
	Contract *DagoramembershipsabiCaller // Generic read-only contract binding to access the raw methods on
}

// DagoramembershipsabiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DagoramembershipsabiTransactorRaw struct {
	Contract *DagoramembershipsabiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDagoramembershipsabi creates a new instance of Dagoramembershipsabi, bound to a specific deployed contract.
func NewDagoramembershipsabi(address common.Address, backend bind.ContractBackend) (*Dagoramembershipsabi, error) {
	contract, err := bindDagoramembershipsabi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dagoramembershipsabi{DagoramembershipsabiCaller: DagoramembershipsabiCaller{contract: contract}, DagoramembershipsabiTransactor: DagoramembershipsabiTransactor{contract: contract}, DagoramembershipsabiFilterer: DagoramembershipsabiFilterer{contract: contract}}, nil
}

// NewDagoramembershipsabiCaller creates a new read-only instance of Dagoramembershipsabi, bound to a specific deployed contract.
func NewDagoramembershipsabiCaller(address common.Address, caller bind.ContractCaller) (*DagoramembershipsabiCaller, error) {
	contract, err := bindDagoramembershipsabi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiCaller{contract: contract}, nil
}

// NewDagoramembershipsabiTransactor creates a new write-only instance of Dagoramembershipsabi, bound to a specific deployed contract.
func NewDagoramembershipsabiTransactor(address common.Address, transactor bind.ContractTransactor) (*DagoramembershipsabiTransactor, error) {
	contract, err := bindDagoramembershipsabi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiTransactor{contract: contract}, nil
}

// NewDagoramembershipsabiFilterer creates a new log filterer instance of Dagoramembershipsabi, bound to a specific deployed contract.
func NewDagoramembershipsabiFilterer(address common.Address, filterer bind.ContractFilterer) (*DagoramembershipsabiFilterer, error) {
	contract, err := bindDagoramembershipsabi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiFilterer{contract: contract}, nil
}

// bindDagoramembershipsabi binds a generic wrapper to an already deployed contract.
func bindDagoramembershipsabi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DagoramembershipsabiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagoramembershipsabi *DagoramembershipsabiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagoramembershipsabi.Contract.DagoramembershipsabiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagoramembershipsabi *DagoramembershipsabiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.DagoramembershipsabiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagoramembershipsabi *DagoramembershipsabiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.DagoramembershipsabiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagoramembershipsabi *DagoramembershipsabiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagoramembershipsabi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.contract.Transact(opts, method, params...)
}

// DAI is a free data retrieval call binding the contract method 0xe0bab4c4.
//
// Solidity: function DAI() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) DAI(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "DAI")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAI is a free data retrieval call binding the contract method 0xe0bab4c4.
//
// Solidity: function DAI() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) DAI() (common.Address, error) {
	return _Dagoramembershipsabi.Contract.DAI(&_Dagoramembershipsabi.CallOpts)
}

// DAI is a free data retrieval call binding the contract method 0xe0bab4c4.
//
// Solidity: function DAI() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) DAI() (common.Address, error) {
	return _Dagoramembershipsabi.Contract.DAI(&_Dagoramembershipsabi.CallOpts)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) GRACEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "GRACE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GRACEPERIOD() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GRACEPERIOD(&_Dagoramembershipsabi.CallOpts)
}

// GRACEPERIOD is a free data retrieval call binding the contract method 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) GRACEPERIOD() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GRACEPERIOD(&_Dagoramembershipsabi.CallOpts)
}

// GetUpgradePrice is a free data retrieval call binding the contract method 0x9b4d7066.
//
// Solidity: function _getUpgradePrice(uint256 tokenId, uint8 oldTier, uint8 newTier) view returns(uint256 _price)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) GetUpgradePrice(opts *bind.CallOpts, tokenId *big.Int, oldTier uint8, newTier uint8) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "_getUpgradePrice", tokenId, oldTier, newTier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpgradePrice is a free data retrieval call binding the contract method 0x9b4d7066.
//
// Solidity: function _getUpgradePrice(uint256 tokenId, uint8 oldTier, uint8 newTier) view returns(uint256 _price)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GetUpgradePrice(tokenId *big.Int, oldTier uint8, newTier uint8) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GetUpgradePrice(&_Dagoramembershipsabi.CallOpts, tokenId, oldTier, newTier)
}

// GetUpgradePrice is a free data retrieval call binding the contract method 0x9b4d7066.
//
// Solidity: function _getUpgradePrice(uint256 tokenId, uint8 oldTier, uint8 newTier) view returns(uint256 _price)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) GetUpgradePrice(tokenId *big.Int, oldTier uint8, newTier uint8) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GetUpgradePrice(&_Dagoramembershipsabi.CallOpts, tokenId, oldTier, newTier)
}

// IsInitialized is a free data retrieval call binding the contract method 0xb893e86e.
//
// Solidity: function _isInitialized() view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "_isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0xb893e86e.
//
// Solidity: function _isInitialized() view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) IsInitialized() (bool, error) {
	return _Dagoramembershipsabi.Contract.IsInitialized(&_Dagoramembershipsabi.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0xb893e86e.
//
// Solidity: function _isInitialized() view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) IsInitialized() (bool, error) {
	return _Dagoramembershipsabi.Contract.IsInitialized(&_Dagoramembershipsabi.CallOpts)
}

// AddressTokenIds is a free data retrieval call binding the contract method 0xcc0843f2.
//
// Solidity: function addressTokenIds(address _owner) view returns(uint256 _tokenId)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) AddressTokenIds(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "addressTokenIds", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressTokenIds is a free data retrieval call binding the contract method 0xcc0843f2.
//
// Solidity: function addressTokenIds(address _owner) view returns(uint256 _tokenId)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) AddressTokenIds(_owner common.Address) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.AddressTokenIds(&_Dagoramembershipsabi.CallOpts, _owner)
}

// AddressTokenIds is a free data retrieval call binding the contract method 0xcc0843f2.
//
// Solidity: function addressTokenIds(address _owner) view returns(uint256 _tokenId)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) AddressTokenIds(_owner common.Address) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.AddressTokenIds(&_Dagoramembershipsabi.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.BalanceOf(&_Dagoramembershipsabi.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.BalanceOf(&_Dagoramembershipsabi.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) BaseURI() (string, error) {
	return _Dagoramembershipsabi.Contract.BaseURI(&_Dagoramembershipsabi.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) BaseURI() (string, error) {
	return _Dagoramembershipsabi.Contract.BaseURI(&_Dagoramembershipsabi.CallOpts)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) Claimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "claimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) Claimed(arg0 common.Address) (bool, error) {
	return _Dagoramembershipsabi.Contract.Claimed(&_Dagoramembershipsabi.CallOpts, arg0)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) Claimed(arg0 common.Address) (bool, error) {
	return _Dagoramembershipsabi.Contract.Claimed(&_Dagoramembershipsabi.CallOpts, arg0)
}

// DagoraRenewPrice is a free data retrieval call binding the contract method 0xc54343db.
//
// Solidity: function dagoraRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) DagoraRenewPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "dagoraRenewPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DagoraRenewPrice is a free data retrieval call binding the contract method 0xc54343db.
//
// Solidity: function dagoraRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) DagoraRenewPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.DagoraRenewPrice(&_Dagoramembershipsabi.CallOpts)
}

// DagoraRenewPrice is a free data retrieval call binding the contract method 0xc54343db.
//
// Solidity: function dagoraRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) DagoraRenewPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.DagoraRenewPrice(&_Dagoramembershipsabi.CallOpts)
}

// DagoraTreasury is a free data retrieval call binding the contract method 0x8733067b.
//
// Solidity: function dagoraTreasury() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) DagoraTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "dagoraTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DagoraTreasury is a free data retrieval call binding the contract method 0x8733067b.
//
// Solidity: function dagoraTreasury() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) DagoraTreasury() (common.Address, error) {
	return _Dagoramembershipsabi.Contract.DagoraTreasury(&_Dagoramembershipsabi.CallOpts)
}

// DagoraTreasury is a free data retrieval call binding the contract method 0x8733067b.
//
// Solidity: function dagoraTreasury() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) DagoraTreasury() (common.Address, error) {
	return _Dagoramembershipsabi.Contract.DagoraTreasury(&_Dagoramembershipsabi.CallOpts)
}

// DagorianPrice is a free data retrieval call binding the contract method 0x9cc2ec01.
//
// Solidity: function dagorianPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) DagorianPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "dagorianPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DagorianPrice is a free data retrieval call binding the contract method 0x9cc2ec01.
//
// Solidity: function dagorianPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) DagorianPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.DagorianPrice(&_Dagoramembershipsabi.CallOpts)
}

// DagorianPrice is a free data retrieval call binding the contract method 0x9cc2ec01.
//
// Solidity: function dagorianPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) DagorianPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.DagorianPrice(&_Dagoramembershipsabi.CallOpts)
}

// Discount is a free data retrieval call binding the contract method 0x6b6f4a9d.
//
// Solidity: function discount() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) Discount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "discount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Discount is a free data retrieval call binding the contract method 0x6b6f4a9d.
//
// Solidity: function discount() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) Discount() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.Discount(&_Dagoramembershipsabi.CallOpts)
}

// Discount is a free data retrieval call binding the contract method 0x6b6f4a9d.
//
// Solidity: function discount() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) Discount() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.Discount(&_Dagoramembershipsabi.CallOpts)
}

// EcclesiaPrice is a free data retrieval call binding the contract method 0x5c9717ec.
//
// Solidity: function ecclesiaPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) EcclesiaPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "ecclesiaPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EcclesiaPrice is a free data retrieval call binding the contract method 0x5c9717ec.
//
// Solidity: function ecclesiaPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) EcclesiaPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.EcclesiaPrice(&_Dagoramembershipsabi.CallOpts)
}

// EcclesiaPrice is a free data retrieval call binding the contract method 0x5c9717ec.
//
// Solidity: function ecclesiaPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) EcclesiaPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.EcclesiaPrice(&_Dagoramembershipsabi.CallOpts)
}

// EcclesiaRenewPrice is a free data retrieval call binding the contract method 0x6ad88158.
//
// Solidity: function ecclesiaRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) EcclesiaRenewPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "ecclesiaRenewPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EcclesiaRenewPrice is a free data retrieval call binding the contract method 0x6ad88158.
//
// Solidity: function ecclesiaRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) EcclesiaRenewPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.EcclesiaRenewPrice(&_Dagoramembershipsabi.CallOpts)
}

// EcclesiaRenewPrice is a free data retrieval call binding the contract method 0x6ad88158.
//
// Solidity: function ecclesiaRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) EcclesiaRenewPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.EcclesiaRenewPrice(&_Dagoramembershipsabi.CallOpts)
}

// Experation is a free data retrieval call binding the contract method 0x90a0bddf.
//
// Solidity: function experation(uint256 ) view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) Experation(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "experation", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Experation is a free data retrieval call binding the contract method 0x90a0bddf.
//
// Solidity: function experation(uint256 ) view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) Experation(arg0 *big.Int) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.Experation(&_Dagoramembershipsabi.CallOpts, arg0)
}

// Experation is a free data retrieval call binding the contract method 0x90a0bddf.
//
// Solidity: function experation(uint256 ) view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) Experation(arg0 *big.Int) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.Experation(&_Dagoramembershipsabi.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Dagoramembershipsabi.Contract.GetApproved(&_Dagoramembershipsabi.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Dagoramembershipsabi.Contract.GetApproved(&_Dagoramembershipsabi.CallOpts, tokenId)
}

// GetExpiration is a free data retrieval call binding the contract method 0xa05b775f.
//
// Solidity: function getExpiration(uint256 _tokenId) view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) GetExpiration(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "getExpiration", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpiration is a free data retrieval call binding the contract method 0xa05b775f.
//
// Solidity: function getExpiration(uint256 _tokenId) view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GetExpiration(_tokenId *big.Int) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GetExpiration(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// GetExpiration is a free data retrieval call binding the contract method 0xa05b775f.
//
// Solidity: function getExpiration(uint256 _tokenId) view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) GetExpiration(_tokenId *big.Int) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GetExpiration(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// GetMembership is a free data retrieval call binding the contract method 0x38fbf2c6.
//
// Solidity: function getMembership(uint256 _tokenId) view returns((uint8,address,uint256,uint256))
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) GetMembership(opts *bind.CallOpts, _tokenId *big.Int) (DagoraMembershipsV1Membership, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "getMembership", _tokenId)

	if err != nil {
		return *new(DagoraMembershipsV1Membership), err
	}

	out0 := *abi.ConvertType(out[0], new(DagoraMembershipsV1Membership)).(*DagoraMembershipsV1Membership)

	return out0, err

}

// GetMembership is a free data retrieval call binding the contract method 0x38fbf2c6.
//
// Solidity: function getMembership(uint256 _tokenId) view returns((uint8,address,uint256,uint256))
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GetMembership(_tokenId *big.Int) (DagoraMembershipsV1Membership, error) {
	return _Dagoramembershipsabi.Contract.GetMembership(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// GetMembership is a free data retrieval call binding the contract method 0x38fbf2c6.
//
// Solidity: function getMembership(uint256 _tokenId) view returns((uint8,address,uint256,uint256))
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) GetMembership(_tokenId *big.Int) (DagoraMembershipsV1Membership, error) {
	return _Dagoramembershipsabi.Contract.GetMembership(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// GetMembershipTier is a free data retrieval call binding the contract method 0x7dd346a9.
//
// Solidity: function getMembershipTier(uint256 _tokenId) view returns(uint8)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) GetMembershipTier(opts *bind.CallOpts, _tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "getMembershipTier", _tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMembershipTier is a free data retrieval call binding the contract method 0x7dd346a9.
//
// Solidity: function getMembershipTier(uint256 _tokenId) view returns(uint8)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GetMembershipTier(_tokenId *big.Int) (uint8, error) {
	return _Dagoramembershipsabi.Contract.GetMembershipTier(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// GetMembershipTier is a free data retrieval call binding the contract method 0x7dd346a9.
//
// Solidity: function getMembershipTier(uint256 _tokenId) view returns(uint8)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) GetMembershipTier(_tokenId *big.Int) (uint8, error) {
	return _Dagoramembershipsabi.Contract.GetMembershipTier(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// GetMintPrice is a free data retrieval call binding the contract method 0x08e6c1a7.
//
// Solidity: function getMintPrice(uint96 _durationInMonths, uint8 _tier) view returns(uint256 _price)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) GetMintPrice(opts *bind.CallOpts, _durationInMonths *big.Int, _tier uint8) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "getMintPrice", _durationInMonths, _tier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintPrice is a free data retrieval call binding the contract method 0x08e6c1a7.
//
// Solidity: function getMintPrice(uint96 _durationInMonths, uint8 _tier) view returns(uint256 _price)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GetMintPrice(_durationInMonths *big.Int, _tier uint8) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GetMintPrice(&_Dagoramembershipsabi.CallOpts, _durationInMonths, _tier)
}

// GetMintPrice is a free data retrieval call binding the contract method 0x08e6c1a7.
//
// Solidity: function getMintPrice(uint96 _durationInMonths, uint8 _tier) view returns(uint256 _price)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) GetMintPrice(_durationInMonths *big.Int, _tier uint8) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GetMintPrice(&_Dagoramembershipsabi.CallOpts, _durationInMonths, _tier)
}

// GetRenewalPrice is a free data retrieval call binding the contract method 0xb9e8e750.
//
// Solidity: function getRenewalPrice(uint96 _newDuration, uint8 currentTier) view returns(uint256 _price)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) GetRenewalPrice(opts *bind.CallOpts, _newDuration *big.Int, currentTier uint8) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "getRenewalPrice", _newDuration, currentTier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRenewalPrice is a free data retrieval call binding the contract method 0xb9e8e750.
//
// Solidity: function getRenewalPrice(uint96 _newDuration, uint8 currentTier) view returns(uint256 _price)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GetRenewalPrice(_newDuration *big.Int, currentTier uint8) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GetRenewalPrice(&_Dagoramembershipsabi.CallOpts, _newDuration, currentTier)
}

// GetRenewalPrice is a free data retrieval call binding the contract method 0xb9e8e750.
//
// Solidity: function getRenewalPrice(uint96 _newDuration, uint8 currentTier) view returns(uint256 _price)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) GetRenewalPrice(_newDuration *big.Int, currentTier uint8) (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.GetRenewalPrice(&_Dagoramembershipsabi.CallOpts, _newDuration, currentTier)
}

// GetTokenDelegates is a free data retrieval call binding the contract method 0xc537bccc.
//
// Solidity: function getTokenDelegates(uint256 _tokenId) view returns(address[])
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) GetTokenDelegates(opts *bind.CallOpts, _tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "getTokenDelegates", _tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenDelegates is a free data retrieval call binding the contract method 0xc537bccc.
//
// Solidity: function getTokenDelegates(uint256 _tokenId) view returns(address[])
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GetTokenDelegates(_tokenId *big.Int) ([]common.Address, error) {
	return _Dagoramembershipsabi.Contract.GetTokenDelegates(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// GetTokenDelegates is a free data retrieval call binding the contract method 0xc537bccc.
//
// Solidity: function getTokenDelegates(uint256 _tokenId) view returns(address[])
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) GetTokenDelegates(_tokenId *big.Int) ([]common.Address, error) {
	return _Dagoramembershipsabi.Contract.GetTokenDelegates(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// HoplitePrice is a free data retrieval call binding the contract method 0x7978b3e0.
//
// Solidity: function hoplitePrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) HoplitePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "hoplitePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HoplitePrice is a free data retrieval call binding the contract method 0x7978b3e0.
//
// Solidity: function hoplitePrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) HoplitePrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.HoplitePrice(&_Dagoramembershipsabi.CallOpts)
}

// HoplitePrice is a free data retrieval call binding the contract method 0x7978b3e0.
//
// Solidity: function hoplitePrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) HoplitePrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.HoplitePrice(&_Dagoramembershipsabi.CallOpts)
}

// HopliteRenewPrice is a free data retrieval call binding the contract method 0xe6d44510.
//
// Solidity: function hopliteRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) HopliteRenewPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "hopliteRenewPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HopliteRenewPrice is a free data retrieval call binding the contract method 0xe6d44510.
//
// Solidity: function hopliteRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) HopliteRenewPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.HopliteRenewPrice(&_Dagoramembershipsabi.CallOpts)
}

// HopliteRenewPrice is a free data retrieval call binding the contract method 0xe6d44510.
//
// Solidity: function hopliteRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) HopliteRenewPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.HopliteRenewPrice(&_Dagoramembershipsabi.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Dagoramembershipsabi.Contract.IsApprovedForAll(&_Dagoramembershipsabi.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Dagoramembershipsabi.Contract.IsApprovedForAll(&_Dagoramembershipsabi.CallOpts, owner, operator)
}

// IsOwnerOrDelegate is a free data retrieval call binding the contract method 0x2fa117e5.
//
// Solidity: function isOwnerOrDelegate(uint256 tokenId, address addrs) view returns(bool _isOwnerOrDelegate)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) IsOwnerOrDelegate(opts *bind.CallOpts, tokenId *big.Int, addrs common.Address) (bool, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "isOwnerOrDelegate", tokenId, addrs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwnerOrDelegate is a free data retrieval call binding the contract method 0x2fa117e5.
//
// Solidity: function isOwnerOrDelegate(uint256 tokenId, address addrs) view returns(bool _isOwnerOrDelegate)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) IsOwnerOrDelegate(tokenId *big.Int, addrs common.Address) (bool, error) {
	return _Dagoramembershipsabi.Contract.IsOwnerOrDelegate(&_Dagoramembershipsabi.CallOpts, tokenId, addrs)
}

// IsOwnerOrDelegate is a free data retrieval call binding the contract method 0x2fa117e5.
//
// Solidity: function isOwnerOrDelegate(uint256 tokenId, address addrs) view returns(bool _isOwnerOrDelegate)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) IsOwnerOrDelegate(tokenId *big.Int, addrs common.Address) (bool, error) {
	return _Dagoramembershipsabi.Contract.IsOwnerOrDelegate(&_Dagoramembershipsabi.CallOpts, tokenId, addrs)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) IsPaused() (bool, error) {
	return _Dagoramembershipsabi.Contract.IsPaused(&_Dagoramembershipsabi.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) IsPaused() (bool, error) {
	return _Dagoramembershipsabi.Contract.IsPaused(&_Dagoramembershipsabi.CallOpts)
}

// IsValidMembership is a free data retrieval call binding the contract method 0xc5a9c2c7.
//
// Solidity: function isValidMembership(uint256 _tokenId) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) IsValidMembership(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "isValidMembership", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidMembership is a free data retrieval call binding the contract method 0xc5a9c2c7.
//
// Solidity: function isValidMembership(uint256 _tokenId) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) IsValidMembership(_tokenId *big.Int) (bool, error) {
	return _Dagoramembershipsabi.Contract.IsValidMembership(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// IsValidMembership is a free data retrieval call binding the contract method 0xc5a9c2c7.
//
// Solidity: function isValidMembership(uint256 _tokenId) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) IsValidMembership(_tokenId *big.Int) (bool, error) {
	return _Dagoramembershipsabi.Contract.IsValidMembership(&_Dagoramembershipsabi.CallOpts, _tokenId)
}

// Memberships is a free data retrieval call binding the contract method 0x321621d7.
//
// Solidity: function memberships(uint256 ) view returns(uint8 tier, address member, uint256 tokenId, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) Memberships(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Tier       uint8
	Member     common.Address
	TokenId    *big.Int
	Expiration *big.Int
}, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "memberships", arg0)

	outstruct := new(struct {
		Tier       uint8
		Member     common.Address
		TokenId    *big.Int
		Expiration *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tier = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Member = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Memberships is a free data retrieval call binding the contract method 0x321621d7.
//
// Solidity: function memberships(uint256 ) view returns(uint8 tier, address member, uint256 tokenId, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) Memberships(arg0 *big.Int) (struct {
	Tier       uint8
	Member     common.Address
	TokenId    *big.Int
	Expiration *big.Int
}, error) {
	return _Dagoramembershipsabi.Contract.Memberships(&_Dagoramembershipsabi.CallOpts, arg0)
}

// Memberships is a free data retrieval call binding the contract method 0x321621d7.
//
// Solidity: function memberships(uint256 ) view returns(uint8 tier, address member, uint256 tokenId, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) Memberships(arg0 *big.Int) (struct {
	Tier       uint8
	Member     common.Address
	TokenId    *big.Int
	Expiration *big.Int
}, error) {
	return _Dagoramembershipsabi.Contract.Memberships(&_Dagoramembershipsabi.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) Name() (string, error) {
	return _Dagoramembershipsabi.Contract.Name(&_Dagoramembershipsabi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) Name() (string, error) {
	return _Dagoramembershipsabi.Contract.Name(&_Dagoramembershipsabi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) Owner() (common.Address, error) {
	return _Dagoramembershipsabi.Contract.Owner(&_Dagoramembershipsabi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) Owner() (common.Address, error) {
	return _Dagoramembershipsabi.Contract.Owner(&_Dagoramembershipsabi.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Dagoramembershipsabi.Contract.OwnerOf(&_Dagoramembershipsabi.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Dagoramembershipsabi.Contract.OwnerOf(&_Dagoramembershipsabi.CallOpts, tokenId)
}

// PercelsiaPrice is a free data retrieval call binding the contract method 0x426c3d43.
//
// Solidity: function percelsiaPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) PercelsiaPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "percelsiaPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PercelsiaPrice is a free data retrieval call binding the contract method 0x426c3d43.
//
// Solidity: function percelsiaPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) PercelsiaPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.PercelsiaPrice(&_Dagoramembershipsabi.CallOpts)
}

// PercelsiaPrice is a free data retrieval call binding the contract method 0x426c3d43.
//
// Solidity: function percelsiaPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) PercelsiaPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.PercelsiaPrice(&_Dagoramembershipsabi.CallOpts)
}

// PercelsiaRenewPrice is a free data retrieval call binding the contract method 0x407ecdd1.
//
// Solidity: function percelsiaRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) PercelsiaRenewPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "percelsiaRenewPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PercelsiaRenewPrice is a free data retrieval call binding the contract method 0x407ecdd1.
//
// Solidity: function percelsiaRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) PercelsiaRenewPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.PercelsiaRenewPrice(&_Dagoramembershipsabi.CallOpts)
}

// PercelsiaRenewPrice is a free data retrieval call binding the contract method 0x407ecdd1.
//
// Solidity: function percelsiaRenewPrice() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) PercelsiaRenewPrice() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.PercelsiaRenewPrice(&_Dagoramembershipsabi.CallOpts)
}

// ProxyImplementation is a free data retrieval call binding the contract method 0x0c870f91.
//
// Solidity: function proxyImplementation() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) ProxyImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "proxyImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyImplementation is a free data retrieval call binding the contract method 0x0c870f91.
//
// Solidity: function proxyImplementation() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) ProxyImplementation() (common.Address, error) {
	return _Dagoramembershipsabi.Contract.ProxyImplementation(&_Dagoramembershipsabi.CallOpts)
}

// ProxyImplementation is a free data retrieval call binding the contract method 0x0c870f91.
//
// Solidity: function proxyImplementation() view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) ProxyImplementation() (common.Address, error) {
	return _Dagoramembershipsabi.Contract.ProxyImplementation(&_Dagoramembershipsabi.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dagoramembershipsabi.Contract.SupportsInterface(&_Dagoramembershipsabi.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dagoramembershipsabi.Contract.SupportsInterface(&_Dagoramembershipsabi.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) Symbol() (string, error) {
	return _Dagoramembershipsabi.Contract.Symbol(&_Dagoramembershipsabi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) Symbol() (string, error) {
	return _Dagoramembershipsabi.Contract.Symbol(&_Dagoramembershipsabi.CallOpts)
}

// TokenDelegates is a free data retrieval call binding the contract method 0xc4c91695.
//
// Solidity: function tokenDelegates(uint256 , uint256 ) view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) TokenDelegates(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "tokenDelegates", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenDelegates is a free data retrieval call binding the contract method 0xc4c91695.
//
// Solidity: function tokenDelegates(uint256 , uint256 ) view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) TokenDelegates(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Dagoramembershipsabi.Contract.TokenDelegates(&_Dagoramembershipsabi.CallOpts, arg0, arg1)
}

// TokenDelegates is a free data retrieval call binding the contract method 0xc4c91695.
//
// Solidity: function tokenDelegates(uint256 , uint256 ) view returns(address)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) TokenDelegates(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Dagoramembershipsabi.Contract.TokenDelegates(&_Dagoramembershipsabi.CallOpts, arg0, arg1)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Dagoramembershipsabi.Contract.TokenURI(&_Dagoramembershipsabi.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Dagoramembershipsabi.Contract.TokenURI(&_Dagoramembershipsabi.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagoramembershipsabi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiSession) TotalSupply() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.TotalSupply(&_Dagoramembershipsabi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dagoramembershipsabi *DagoramembershipsabiCallerSession) TotalSupply() (*big.Int, error) {
	return _Dagoramembershipsabi.Contract.TotalSupply(&_Dagoramembershipsabi.CallOpts)
}

// AddDelegate is a paid mutator transaction binding the contract method 0x5bcfd212.
//
// Solidity: function addDelegate(address _delegatee, uint256 _tokenId) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) AddDelegate(opts *bind.TransactOpts, _delegatee common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "addDelegate", _delegatee, _tokenId)
}

// AddDelegate is a paid mutator transaction binding the contract method 0x5bcfd212.
//
// Solidity: function addDelegate(address _delegatee, uint256 _tokenId) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) AddDelegate(_delegatee common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.AddDelegate(&_Dagoramembershipsabi.TransactOpts, _delegatee, _tokenId)
}

// AddDelegate is a paid mutator transaction binding the contract method 0x5bcfd212.
//
// Solidity: function addDelegate(address _delegatee, uint256 _tokenId) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) AddDelegate(_delegatee common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.AddDelegate(&_Dagoramembershipsabi.TransactOpts, _delegatee, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.Approve(&_Dagoramembershipsabi.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.Approve(&_Dagoramembershipsabi.TransactOpts, to, tokenId)
}

// CancelMembership is a paid mutator transaction binding the contract method 0x89def320.
//
// Solidity: function cancelMembership(uint256 tokenId) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) CancelMembership(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "cancelMembership", tokenId)
}

// CancelMembership is a paid mutator transaction binding the contract method 0x89def320.
//
// Solidity: function cancelMembership(uint256 tokenId) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) CancelMembership(tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.CancelMembership(&_Dagoramembershipsabi.TransactOpts, tokenId)
}

// CancelMembership is a paid mutator transaction binding the contract method 0x89def320.
//
// Solidity: function cancelMembership(uint256 tokenId) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) CancelMembership(tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.CancelMembership(&_Dagoramembershipsabi.TransactOpts, tokenId)
}

// FreeMint is a paid mutator transaction binding the contract method 0x5b70ea9f.
//
// Solidity: function freeMint() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) FreeMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "freeMint")
}

// FreeMint is a paid mutator transaction binding the contract method 0x5b70ea9f.
//
// Solidity: function freeMint() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) FreeMint() (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.FreeMint(&_Dagoramembershipsabi.TransactOpts)
}

// FreeMint is a paid mutator transaction binding the contract method 0x5b70ea9f.
//
// Solidity: function freeMint() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) FreeMint() (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.FreeMint(&_Dagoramembershipsabi.TransactOpts)
}

// GiftExtension is a paid mutator transaction binding the contract method 0x3e8fb84b.
//
// Solidity: function giftExtension(uint256 tokenId, uint96 durationInMonths) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) GiftExtension(opts *bind.TransactOpts, tokenId *big.Int, durationInMonths *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "giftExtension", tokenId, durationInMonths)
}

// GiftExtension is a paid mutator transaction binding the contract method 0x3e8fb84b.
//
// Solidity: function giftExtension(uint256 tokenId, uint96 durationInMonths) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GiftExtension(tokenId *big.Int, durationInMonths *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.GiftExtension(&_Dagoramembershipsabi.TransactOpts, tokenId, durationInMonths)
}

// GiftExtension is a paid mutator transaction binding the contract method 0x3e8fb84b.
//
// Solidity: function giftExtension(uint256 tokenId, uint96 durationInMonths) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) GiftExtension(tokenId *big.Int, durationInMonths *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.GiftExtension(&_Dagoramembershipsabi.TransactOpts, tokenId, durationInMonths)
}

// GiftMembership is a paid mutator transaction binding the contract method 0x5a0e932f.
//
// Solidity: function giftMembership(address to, uint8 tier, uint96 durationInMonths) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) GiftMembership(opts *bind.TransactOpts, to common.Address, tier uint8, durationInMonths *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "giftMembership", to, tier, durationInMonths)
}

// GiftMembership is a paid mutator transaction binding the contract method 0x5a0e932f.
//
// Solidity: function giftMembership(address to, uint8 tier, uint96 durationInMonths) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GiftMembership(to common.Address, tier uint8, durationInMonths *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.GiftMembership(&_Dagoramembershipsabi.TransactOpts, to, tier, durationInMonths)
}

// GiftMembership is a paid mutator transaction binding the contract method 0x5a0e932f.
//
// Solidity: function giftMembership(address to, uint8 tier, uint96 durationInMonths) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) GiftMembership(to common.Address, tier uint8, durationInMonths *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.GiftMembership(&_Dagoramembershipsabi.TransactOpts, to, tier, durationInMonths)
}

// GiftUpgrade is a paid mutator transaction binding the contract method 0x970e4d85.
//
// Solidity: function giftUpgrade(uint256 tokenId, uint8 tier) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) GiftUpgrade(opts *bind.TransactOpts, tokenId *big.Int, tier uint8) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "giftUpgrade", tokenId, tier)
}

// GiftUpgrade is a paid mutator transaction binding the contract method 0x970e4d85.
//
// Solidity: function giftUpgrade(uint256 tokenId, uint8 tier) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) GiftUpgrade(tokenId *big.Int, tier uint8) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.GiftUpgrade(&_Dagoramembershipsabi.TransactOpts, tokenId, tier)
}

// GiftUpgrade is a paid mutator transaction binding the contract method 0x970e4d85.
//
// Solidity: function giftUpgrade(uint256 tokenId, uint8 tier) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) GiftUpgrade(tokenId *big.Int, tier uint8) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.GiftUpgrade(&_Dagoramembershipsabi.TransactOpts, tokenId, tier)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string _name, string _symbol, string baseURI_, address _dagoraTreasury) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, baseURI_ string, _dagoraTreasury common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "initialize", _name, _symbol, baseURI_, _dagoraTreasury)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string _name, string _symbol, string baseURI_, address _dagoraTreasury) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) Initialize(_name string, _symbol string, baseURI_ string, _dagoraTreasury common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.Initialize(&_Dagoramembershipsabi.TransactOpts, _name, _symbol, baseURI_, _dagoraTreasury)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string _name, string _symbol, string baseURI_, address _dagoraTreasury) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) Initialize(_name string, _symbol string, baseURI_ string, _dagoraTreasury common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.Initialize(&_Dagoramembershipsabi.TransactOpts, _name, _symbol, baseURI_, _dagoraTreasury)
}

// MintMembership is a paid mutator transaction binding the contract method 0xcdc34ea2.
//
// Solidity: function mintMembership(uint8 _tier, uint96 _durationInMonths) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) MintMembership(opts *bind.TransactOpts, _tier uint8, _durationInMonths *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "mintMembership", _tier, _durationInMonths)
}

// MintMembership is a paid mutator transaction binding the contract method 0xcdc34ea2.
//
// Solidity: function mintMembership(uint8 _tier, uint96 _durationInMonths) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) MintMembership(_tier uint8, _durationInMonths *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.MintMembership(&_Dagoramembershipsabi.TransactOpts, _tier, _durationInMonths)
}

// MintMembership is a paid mutator transaction binding the contract method 0xcdc34ea2.
//
// Solidity: function mintMembership(uint8 _tier, uint96 _durationInMonths) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) MintMembership(_tier uint8, _durationInMonths *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.MintMembership(&_Dagoramembershipsabi.TransactOpts, _tier, _durationInMonths)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x19cb1662.
//
// Solidity: function removeDelegate(address _delegatee, uint256 _tokenId, uint8 slot) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) RemoveDelegate(opts *bind.TransactOpts, _delegatee common.Address, _tokenId *big.Int, slot uint8) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "removeDelegate", _delegatee, _tokenId, slot)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x19cb1662.
//
// Solidity: function removeDelegate(address _delegatee, uint256 _tokenId, uint8 slot) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) RemoveDelegate(_delegatee common.Address, _tokenId *big.Int, slot uint8) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.RemoveDelegate(&_Dagoramembershipsabi.TransactOpts, _delegatee, _tokenId, slot)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x19cb1662.
//
// Solidity: function removeDelegate(address _delegatee, uint256 _tokenId, uint8 slot) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) RemoveDelegate(_delegatee common.Address, _tokenId *big.Int, slot uint8) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.RemoveDelegate(&_Dagoramembershipsabi.TransactOpts, _delegatee, _tokenId, slot)
}

// RenewMembership is a paid mutator transaction binding the contract method 0x1464bb48.
//
// Solidity: function renewMembership(uint96 _durationInMonths, uint256 _tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) RenewMembership(opts *bind.TransactOpts, _durationInMonths *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "renewMembership", _durationInMonths, _tokenId)
}

// RenewMembership is a paid mutator transaction binding the contract method 0x1464bb48.
//
// Solidity: function renewMembership(uint96 _durationInMonths, uint256 _tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) RenewMembership(_durationInMonths *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.RenewMembership(&_Dagoramembershipsabi.TransactOpts, _durationInMonths, _tokenId)
}

// RenewMembership is a paid mutator transaction binding the contract method 0x1464bb48.
//
// Solidity: function renewMembership(uint96 _durationInMonths, uint256 _tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) RenewMembership(_durationInMonths *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.RenewMembership(&_Dagoramembershipsabi.TransactOpts, _durationInMonths, _tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.RenounceOwnership(&_Dagoramembershipsabi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.RenounceOwnership(&_Dagoramembershipsabi.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SafeTransferFrom(&_Dagoramembershipsabi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SafeTransferFrom(&_Dagoramembershipsabi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SafeTransferFrom0(&_Dagoramembershipsabi.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SafeTransferFrom0(&_Dagoramembershipsabi.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetApprovalForAll(&_Dagoramembershipsabi.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetApprovalForAll(&_Dagoramembershipsabi.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setBaseURI", baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetBaseURI(&_Dagoramembershipsabi.TransactOpts, baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetBaseURI(&_Dagoramembershipsabi.TransactOpts, baseURI_)
}

// SetDagoraTreasury is a paid mutator transaction binding the contract method 0xd08706b4.
//
// Solidity: function setDagoraTreasury(address _dagoraTreasury) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetDagoraTreasury(opts *bind.TransactOpts, _dagoraTreasury common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setDagoraTreasury", _dagoraTreasury)
}

// SetDagoraTreasury is a paid mutator transaction binding the contract method 0xd08706b4.
//
// Solidity: function setDagoraTreasury(address _dagoraTreasury) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetDagoraTreasury(_dagoraTreasury common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetDagoraTreasury(&_Dagoramembershipsabi.TransactOpts, _dagoraTreasury)
}

// SetDagoraTreasury is a paid mutator transaction binding the contract method 0xd08706b4.
//
// Solidity: function setDagoraTreasury(address _dagoraTreasury) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetDagoraTreasury(_dagoraTreasury common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetDagoraTreasury(&_Dagoramembershipsabi.TransactOpts, _dagoraTreasury)
}

// SetDagorianPrice is a paid mutator transaction binding the contract method 0x8e828148.
//
// Solidity: function setDagorianPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetDagorianPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setDagorianPrice", _price)
}

// SetDagorianPrice is a paid mutator transaction binding the contract method 0x8e828148.
//
// Solidity: function setDagorianPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetDagorianPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetDagorianPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetDagorianPrice is a paid mutator transaction binding the contract method 0x8e828148.
//
// Solidity: function setDagorianPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetDagorianPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetDagorianPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetDagorianRenewPrice is a paid mutator transaction binding the contract method 0x77aa5b80.
//
// Solidity: function setDagorianRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetDagorianRenewPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setDagorianRenewPrice", _price)
}

// SetDagorianRenewPrice is a paid mutator transaction binding the contract method 0x77aa5b80.
//
// Solidity: function setDagorianRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetDagorianRenewPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetDagorianRenewPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetDagorianRenewPrice is a paid mutator transaction binding the contract method 0x77aa5b80.
//
// Solidity: function setDagorianRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetDagorianRenewPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetDagorianRenewPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetDiscount is a paid mutator transaction binding the contract method 0xdabd2719.
//
// Solidity: function setDiscount(uint256 _discount) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetDiscount(opts *bind.TransactOpts, _discount *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setDiscount", _discount)
}

// SetDiscount is a paid mutator transaction binding the contract method 0xdabd2719.
//
// Solidity: function setDiscount(uint256 _discount) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetDiscount(_discount *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetDiscount(&_Dagoramembershipsabi.TransactOpts, _discount)
}

// SetDiscount is a paid mutator transaction binding the contract method 0xdabd2719.
//
// Solidity: function setDiscount(uint256 _discount) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetDiscount(_discount *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetDiscount(&_Dagoramembershipsabi.TransactOpts, _discount)
}

// SetEcclesiaPrice is a paid mutator transaction binding the contract method 0xd23e1f01.
//
// Solidity: function setEcclesiaPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetEcclesiaPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setEcclesiaPrice", _price)
}

// SetEcclesiaPrice is a paid mutator transaction binding the contract method 0xd23e1f01.
//
// Solidity: function setEcclesiaPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetEcclesiaPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetEcclesiaPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetEcclesiaPrice is a paid mutator transaction binding the contract method 0xd23e1f01.
//
// Solidity: function setEcclesiaPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetEcclesiaPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetEcclesiaPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetEcclesiaRenewPrice is a paid mutator transaction binding the contract method 0xe6ea07d6.
//
// Solidity: function setEcclesiaRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetEcclesiaRenewPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setEcclesiaRenewPrice", _price)
}

// SetEcclesiaRenewPrice is a paid mutator transaction binding the contract method 0xe6ea07d6.
//
// Solidity: function setEcclesiaRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetEcclesiaRenewPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetEcclesiaRenewPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetEcclesiaRenewPrice is a paid mutator transaction binding the contract method 0xe6ea07d6.
//
// Solidity: function setEcclesiaRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetEcclesiaRenewPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetEcclesiaRenewPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetHoplitePrice is a paid mutator transaction binding the contract method 0x454da668.
//
// Solidity: function setHoplitePrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetHoplitePrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setHoplitePrice", _price)
}

// SetHoplitePrice is a paid mutator transaction binding the contract method 0x454da668.
//
// Solidity: function setHoplitePrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetHoplitePrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetHoplitePrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetHoplitePrice is a paid mutator transaction binding the contract method 0x454da668.
//
// Solidity: function setHoplitePrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetHoplitePrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetHoplitePrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetHopliteRenewPrice is a paid mutator transaction binding the contract method 0x594b8163.
//
// Solidity: function setHopliteRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetHopliteRenewPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setHopliteRenewPrice", _price)
}

// SetHopliteRenewPrice is a paid mutator transaction binding the contract method 0x594b8163.
//
// Solidity: function setHopliteRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetHopliteRenewPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetHopliteRenewPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetHopliteRenewPrice is a paid mutator transaction binding the contract method 0x594b8163.
//
// Solidity: function setHopliteRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetHopliteRenewPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetHopliteRenewPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetPercelsiaPrice is a paid mutator transaction binding the contract method 0x1d9a5b00.
//
// Solidity: function setPercelsiaPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetPercelsiaPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setPercelsiaPrice", _price)
}

// SetPercelsiaPrice is a paid mutator transaction binding the contract method 0x1d9a5b00.
//
// Solidity: function setPercelsiaPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetPercelsiaPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetPercelsiaPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetPercelsiaPrice is a paid mutator transaction binding the contract method 0x1d9a5b00.
//
// Solidity: function setPercelsiaPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetPercelsiaPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetPercelsiaPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetPercelsiaRenewPrice is a paid mutator transaction binding the contract method 0xc79215ec.
//
// Solidity: function setPercelsiaRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetPercelsiaRenewPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setPercelsiaRenewPrice", _price)
}

// SetPercelsiaRenewPrice is a paid mutator transaction binding the contract method 0xc79215ec.
//
// Solidity: function setPercelsiaRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetPercelsiaRenewPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetPercelsiaRenewPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetPercelsiaRenewPrice is a paid mutator transaction binding the contract method 0xc79215ec.
//
// Solidity: function setPercelsiaRenewPrice(uint256 _price) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetPercelsiaRenewPrice(_price *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetPercelsiaRenewPrice(&_Dagoramembershipsabi.TransactOpts, _price)
}

// SetProxyAddress is a paid mutator transaction binding the contract method 0x46a7dadc.
//
// Solidity: function setProxyAddress(address _proxyAddress) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SetProxyAddress(opts *bind.TransactOpts, _proxyAddress common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "setProxyAddress", _proxyAddress)
}

// SetProxyAddress is a paid mutator transaction binding the contract method 0x46a7dadc.
//
// Solidity: function setProxyAddress(address _proxyAddress) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SetProxyAddress(_proxyAddress common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetProxyAddress(&_Dagoramembershipsabi.TransactOpts, _proxyAddress)
}

// SetProxyAddress is a paid mutator transaction binding the contract method 0x46a7dadc.
//
// Solidity: function setProxyAddress(address _proxyAddress) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SetProxyAddress(_proxyAddress common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SetProxyAddress(&_Dagoramembershipsabi.TransactOpts, _proxyAddress)
}

// SwapDelegate is a paid mutator transaction binding the contract method 0x5cf2afaf.
//
// Solidity: function swapDelegate(uint256 _tokenId, address oldDelegate, address newDelegate) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) SwapDelegate(opts *bind.TransactOpts, _tokenId *big.Int, oldDelegate common.Address, newDelegate common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "swapDelegate", _tokenId, oldDelegate, newDelegate)
}

// SwapDelegate is a paid mutator transaction binding the contract method 0x5cf2afaf.
//
// Solidity: function swapDelegate(uint256 _tokenId, address oldDelegate, address newDelegate) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) SwapDelegate(_tokenId *big.Int, oldDelegate common.Address, newDelegate common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SwapDelegate(&_Dagoramembershipsabi.TransactOpts, _tokenId, oldDelegate, newDelegate)
}

// SwapDelegate is a paid mutator transaction binding the contract method 0x5cf2afaf.
//
// Solidity: function swapDelegate(uint256 _tokenId, address oldDelegate, address newDelegate) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) SwapDelegate(_tokenId *big.Int, oldDelegate common.Address, newDelegate common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.SwapDelegate(&_Dagoramembershipsabi.TransactOpts, _tokenId, oldDelegate, newDelegate)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) TogglePaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "togglePaused")
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) TogglePaused() (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.TogglePaused(&_Dagoramembershipsabi.TransactOpts)
}

// TogglePaused is a paid mutator transaction binding the contract method 0x36566f06.
//
// Solidity: function togglePaused() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) TogglePaused() (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.TogglePaused(&_Dagoramembershipsabi.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.TransferFrom(&_Dagoramembershipsabi.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.TransferFrom(&_Dagoramembershipsabi.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.TransferOwnership(&_Dagoramembershipsabi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.TransferOwnership(&_Dagoramembershipsabi.TransactOpts, newOwner)
}

// UpgradeMembership is a paid mutator transaction binding the contract method 0x632b277e.
//
// Solidity: function upgradeMembership(uint8 newTier, uint8 oldTier, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) UpgradeMembership(opts *bind.TransactOpts, newTier uint8, oldTier uint8, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "upgradeMembership", newTier, oldTier, tokenId)
}

// UpgradeMembership is a paid mutator transaction binding the contract method 0x632b277e.
//
// Solidity: function upgradeMembership(uint8 newTier, uint8 oldTier, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) UpgradeMembership(newTier uint8, oldTier uint8, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.UpgradeMembership(&_Dagoramembershipsabi.TransactOpts, newTier, oldTier, tokenId)
}

// UpgradeMembership is a paid mutator transaction binding the contract method 0x632b277e.
//
// Solidity: function upgradeMembership(uint8 newTier, uint8 oldTier, uint256 tokenId) payable returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) UpgradeMembership(newTier uint8, oldTier uint8, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.UpgradeMembership(&_Dagoramembershipsabi.TransactOpts, newTier, oldTier, tokenId)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _token) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "withdrawERC20", _token)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _token) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) WithdrawERC20(_token common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.WithdrawERC20(&_Dagoramembershipsabi.TransactOpts, _token)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _token) returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) WithdrawERC20(_token common.Address) (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.WithdrawERC20(&_Dagoramembershipsabi.TransactOpts, _token)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactor) WithdrawETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagoramembershipsabi.contract.Transact(opts, "withdrawETH")
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiSession) WithdrawETH() (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.WithdrawETH(&_Dagoramembershipsabi.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() returns()
func (_Dagoramembershipsabi *DagoramembershipsabiTransactorSession) WithdrawETH() (*types.Transaction, error) {
	return _Dagoramembershipsabi.Contract.WithdrawETH(&_Dagoramembershipsabi.TransactOpts)
}

// DagoramembershipsabiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiApprovalIterator struct {
	Event *DagoramembershipsabiApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiApproval represents a Approval event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DagoramembershipsabiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiApprovalIterator{contract: _Dagoramembershipsabi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiApproval)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseApproval(log types.Log) (*DagoramembershipsabiApproval, error) {
	event := new(DagoramembershipsabiApproval)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiApprovalForAllIterator struct {
	Event *DagoramembershipsabiApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiApprovalForAll represents a ApprovalForAll event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DagoramembershipsabiApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiApprovalForAllIterator{contract: _Dagoramembershipsabi.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiApprovalForAll)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseApprovalForAll(log types.Log) (*DagoramembershipsabiApprovalForAll, error) {
	event := new(DagoramembershipsabiApprovalForAll)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiConsecutiveTransferIterator struct {
	Event *DagoramembershipsabiConsecutiveTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiConsecutiveTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiConsecutiveTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*DagoramembershipsabiConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiConsecutiveTransferIterator{contract: _Dagoramembershipsabi.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiConsecutiveTransfer)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseConsecutiveTransfer(log types.Log) (*DagoramembershipsabiConsecutiveTransfer, error) {
	event := new(DagoramembershipsabiConsecutiveTransfer)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiDelegateAddedIterator is returned from FilterDelegateAdded and is used to iterate over the raw logs and unpacked data for DelegateAdded events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiDelegateAddedIterator struct {
	Event *DagoramembershipsabiDelegateAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiDelegateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiDelegateAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiDelegateAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiDelegateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiDelegateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiDelegateAdded represents a DelegateAdded event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiDelegateAdded struct {
	Member    common.Address
	TokenId   *big.Int
	Delegatee common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateAdded is a free log retrieval operation binding the contract event 0xb40b86aa864ebe491a5c7e19513b562d8a0351b0ddcc86dedbeaa00dc4a43bd3.
//
// Solidity: event DelegateAdded(address indexed member, uint256 indexed tokenId, address delegatee)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterDelegateAdded(opts *bind.FilterOpts, member []common.Address, tokenId []*big.Int) (*DagoramembershipsabiDelegateAddedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "DelegateAdded", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiDelegateAddedIterator{contract: _Dagoramembershipsabi.contract, event: "DelegateAdded", logs: logs, sub: sub}, nil
}

// WatchDelegateAdded is a free log subscription operation binding the contract event 0xb40b86aa864ebe491a5c7e19513b562d8a0351b0ddcc86dedbeaa00dc4a43bd3.
//
// Solidity: event DelegateAdded(address indexed member, uint256 indexed tokenId, address delegatee)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchDelegateAdded(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiDelegateAdded, member []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "DelegateAdded", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiDelegateAdded)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "DelegateAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateAdded is a log parse operation binding the contract event 0xb40b86aa864ebe491a5c7e19513b562d8a0351b0ddcc86dedbeaa00dc4a43bd3.
//
// Solidity: event DelegateAdded(address indexed member, uint256 indexed tokenId, address delegatee)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseDelegateAdded(log types.Log) (*DagoramembershipsabiDelegateAdded, error) {
	event := new(DagoramembershipsabiDelegateAdded)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "DelegateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiDelegateRemovedIterator is returned from FilterDelegateRemoved and is used to iterate over the raw logs and unpacked data for DelegateRemoved events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiDelegateRemovedIterator struct {
	Event *DagoramembershipsabiDelegateRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiDelegateRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiDelegateRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiDelegateRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiDelegateRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiDelegateRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiDelegateRemoved represents a DelegateRemoved event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiDelegateRemoved struct {
	Member    common.Address
	TokenId   *big.Int
	Delegatee common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateRemoved is a free log retrieval operation binding the contract event 0x8ee86be580995f641440ac23ffd5440a237669582927ad4fe2545f70a6309a3c.
//
// Solidity: event DelegateRemoved(address indexed member, uint256 indexed tokenId, address delegatee)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterDelegateRemoved(opts *bind.FilterOpts, member []common.Address, tokenId []*big.Int) (*DagoramembershipsabiDelegateRemovedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "DelegateRemoved", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiDelegateRemovedIterator{contract: _Dagoramembershipsabi.contract, event: "DelegateRemoved", logs: logs, sub: sub}, nil
}

// WatchDelegateRemoved is a free log subscription operation binding the contract event 0x8ee86be580995f641440ac23ffd5440a237669582927ad4fe2545f70a6309a3c.
//
// Solidity: event DelegateRemoved(address indexed member, uint256 indexed tokenId, address delegatee)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchDelegateRemoved(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiDelegateRemoved, member []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "DelegateRemoved", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiDelegateRemoved)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "DelegateRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateRemoved is a log parse operation binding the contract event 0x8ee86be580995f641440ac23ffd5440a237669582927ad4fe2545f70a6309a3c.
//
// Solidity: event DelegateRemoved(address indexed member, uint256 indexed tokenId, address delegatee)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseDelegateRemoved(log types.Log) (*DagoramembershipsabiDelegateRemoved, error) {
	event := new(DagoramembershipsabiDelegateRemoved)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "DelegateRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiDelegateSwappedIterator is returned from FilterDelegateSwapped and is used to iterate over the raw logs and unpacked data for DelegateSwapped events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiDelegateSwappedIterator struct {
	Event *DagoramembershipsabiDelegateSwapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiDelegateSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiDelegateSwapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiDelegateSwapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiDelegateSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiDelegateSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiDelegateSwapped represents a DelegateSwapped event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiDelegateSwapped struct {
	Member       common.Address
	TokenId      *big.Int
	OldDelegatee common.Address
	NewDelegatee common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateSwapped is a free log retrieval operation binding the contract event 0x668c71dc00d40436644bcd8e7ed9975dc0f37adb7a151aafe3309144af84a5b3.
//
// Solidity: event DelegateSwapped(address indexed member, uint256 indexed tokenId, address oldDelegatee, address newDelegatee)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterDelegateSwapped(opts *bind.FilterOpts, member []common.Address, tokenId []*big.Int) (*DagoramembershipsabiDelegateSwappedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "DelegateSwapped", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiDelegateSwappedIterator{contract: _Dagoramembershipsabi.contract, event: "DelegateSwapped", logs: logs, sub: sub}, nil
}

// WatchDelegateSwapped is a free log subscription operation binding the contract event 0x668c71dc00d40436644bcd8e7ed9975dc0f37adb7a151aafe3309144af84a5b3.
//
// Solidity: event DelegateSwapped(address indexed member, uint256 indexed tokenId, address oldDelegatee, address newDelegatee)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchDelegateSwapped(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiDelegateSwapped, member []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "DelegateSwapped", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiDelegateSwapped)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "DelegateSwapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateSwapped is a log parse operation binding the contract event 0x668c71dc00d40436644bcd8e7ed9975dc0f37adb7a151aafe3309144af84a5b3.
//
// Solidity: event DelegateSwapped(address indexed member, uint256 indexed tokenId, address oldDelegatee, address newDelegatee)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseDelegateSwapped(log types.Log) (*DagoramembershipsabiDelegateSwapped, error) {
	event := new(DagoramembershipsabiDelegateSwapped)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "DelegateSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiFreeMembershipClaimedIterator is returned from FilterFreeMembershipClaimed and is used to iterate over the raw logs and unpacked data for FreeMembershipClaimed events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiFreeMembershipClaimedIterator struct {
	Event *DagoramembershipsabiFreeMembershipClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiFreeMembershipClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiFreeMembershipClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiFreeMembershipClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiFreeMembershipClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiFreeMembershipClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiFreeMembershipClaimed represents a FreeMembershipClaimed event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiFreeMembershipClaimed struct {
	Member     common.Address
	TokenId    *big.Int
	Tier       uint8
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFreeMembershipClaimed is a free log retrieval operation binding the contract event 0x5a861571891bddcf3f13024cd56c52af80b18e8bb05326bf2c90cbf4a3fda7f4.
//
// Solidity: event FreeMembershipClaimed(address indexed member, uint256 indexed tokenId, uint8 tier, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterFreeMembershipClaimed(opts *bind.FilterOpts, member []common.Address, tokenId []*big.Int) (*DagoramembershipsabiFreeMembershipClaimedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "FreeMembershipClaimed", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiFreeMembershipClaimedIterator{contract: _Dagoramembershipsabi.contract, event: "FreeMembershipClaimed", logs: logs, sub: sub}, nil
}

// WatchFreeMembershipClaimed is a free log subscription operation binding the contract event 0x5a861571891bddcf3f13024cd56c52af80b18e8bb05326bf2c90cbf4a3fda7f4.
//
// Solidity: event FreeMembershipClaimed(address indexed member, uint256 indexed tokenId, uint8 tier, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchFreeMembershipClaimed(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiFreeMembershipClaimed, member []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "FreeMembershipClaimed", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiFreeMembershipClaimed)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "FreeMembershipClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreeMembershipClaimed is a log parse operation binding the contract event 0x5a861571891bddcf3f13024cd56c52af80b18e8bb05326bf2c90cbf4a3fda7f4.
//
// Solidity: event FreeMembershipClaimed(address indexed member, uint256 indexed tokenId, uint8 tier, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseFreeMembershipClaimed(log types.Log) (*DagoramembershipsabiFreeMembershipClaimed, error) {
	event := new(DagoramembershipsabiFreeMembershipClaimed)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "FreeMembershipClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiInitializedIterator struct {
	Event *DagoramembershipsabiInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiInitialized represents a Initialized event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterInitialized(opts *bind.FilterOpts) (*DagoramembershipsabiInitializedIterator, error) {

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiInitializedIterator{contract: _Dagoramembershipsabi.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiInitialized) (event.Subscription, error) {

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiInitialized)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseInitialized(log types.Log) (*DagoramembershipsabiInitialized, error) {
	event := new(DagoramembershipsabiInitialized)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiMembershipCanceledIterator is returned from FilterMembershipCanceled and is used to iterate over the raw logs and unpacked data for MembershipCanceled events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipCanceledIterator struct {
	Event *DagoramembershipsabiMembershipCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiMembershipCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiMembershipCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiMembershipCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiMembershipCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiMembershipCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiMembershipCanceled represents a MembershipCanceled event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipCanceled struct {
	Member     common.Address
	TokenId    *big.Int
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMembershipCanceled is a free log retrieval operation binding the contract event 0x15ac7f52d52547617925b6b1a38c797c03f44cde1e133547dfc02100063db001.
//
// Solidity: event MembershipCanceled(address indexed member, uint256 indexed tokenId, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterMembershipCanceled(opts *bind.FilterOpts, member []common.Address, tokenId []*big.Int) (*DagoramembershipsabiMembershipCanceledIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "MembershipCanceled", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiMembershipCanceledIterator{contract: _Dagoramembershipsabi.contract, event: "MembershipCanceled", logs: logs, sub: sub}, nil
}

// WatchMembershipCanceled is a free log subscription operation binding the contract event 0x15ac7f52d52547617925b6b1a38c797c03f44cde1e133547dfc02100063db001.
//
// Solidity: event MembershipCanceled(address indexed member, uint256 indexed tokenId, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchMembershipCanceled(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiMembershipCanceled, member []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "MembershipCanceled", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiMembershipCanceled)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMembershipCanceled is a log parse operation binding the contract event 0x15ac7f52d52547617925b6b1a38c797c03f44cde1e133547dfc02100063db001.
//
// Solidity: event MembershipCanceled(address indexed member, uint256 indexed tokenId, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseMembershipCanceled(log types.Log) (*DagoramembershipsabiMembershipCanceled, error) {
	event := new(DagoramembershipsabiMembershipCanceled)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiMembershipGiftedIterator is returned from FilterMembershipGifted and is used to iterate over the raw logs and unpacked data for MembershipGifted events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipGiftedIterator struct {
	Event *DagoramembershipsabiMembershipGifted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiMembershipGiftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiMembershipGifted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiMembershipGifted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiMembershipGiftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiMembershipGiftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiMembershipGifted represents a MembershipGifted event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipGifted struct {
	Member     common.Address
	TokenId    *big.Int
	Tier       uint8
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMembershipGifted is a free log retrieval operation binding the contract event 0x4b946c1e1d3a0d07d3a6dc6a82d838c57da6dc13704822932424e5ddc55ea7ae.
//
// Solidity: event MembershipGifted(address indexed member, uint256 indexed tokenId, uint8 tier, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterMembershipGifted(opts *bind.FilterOpts, member []common.Address, tokenId []*big.Int) (*DagoramembershipsabiMembershipGiftedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "MembershipGifted", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiMembershipGiftedIterator{contract: _Dagoramembershipsabi.contract, event: "MembershipGifted", logs: logs, sub: sub}, nil
}

// WatchMembershipGifted is a free log subscription operation binding the contract event 0x4b946c1e1d3a0d07d3a6dc6a82d838c57da6dc13704822932424e5ddc55ea7ae.
//
// Solidity: event MembershipGifted(address indexed member, uint256 indexed tokenId, uint8 tier, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchMembershipGifted(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiMembershipGifted, member []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "MembershipGifted", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiMembershipGifted)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipGifted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMembershipGifted is a log parse operation binding the contract event 0x4b946c1e1d3a0d07d3a6dc6a82d838c57da6dc13704822932424e5ddc55ea7ae.
//
// Solidity: event MembershipGifted(address indexed member, uint256 indexed tokenId, uint8 tier, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseMembershipGifted(log types.Log) (*DagoramembershipsabiMembershipGifted, error) {
	event := new(DagoramembershipsabiMembershipGifted)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipGifted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiMembershipPurchasedIterator is returned from FilterMembershipPurchased and is used to iterate over the raw logs and unpacked data for MembershipPurchased events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipPurchasedIterator struct {
	Event *DagoramembershipsabiMembershipPurchased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiMembershipPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiMembershipPurchased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiMembershipPurchased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiMembershipPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiMembershipPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiMembershipPurchased represents a MembershipPurchased event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipPurchased struct {
	Member     common.Address
	TokenId    *big.Int
	Tier       uint8
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMembershipPurchased is a free log retrieval operation binding the contract event 0x3ce34c4daa4cc817ecd03b796bfda58751aa15ba8683a597a19d81add1969cfb.
//
// Solidity: event MembershipPurchased(address indexed member, uint256 indexed tokenId, uint8 tier, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterMembershipPurchased(opts *bind.FilterOpts, member []common.Address, tokenId []*big.Int) (*DagoramembershipsabiMembershipPurchasedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "MembershipPurchased", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiMembershipPurchasedIterator{contract: _Dagoramembershipsabi.contract, event: "MembershipPurchased", logs: logs, sub: sub}, nil
}

// WatchMembershipPurchased is a free log subscription operation binding the contract event 0x3ce34c4daa4cc817ecd03b796bfda58751aa15ba8683a597a19d81add1969cfb.
//
// Solidity: event MembershipPurchased(address indexed member, uint256 indexed tokenId, uint8 tier, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchMembershipPurchased(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiMembershipPurchased, member []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "MembershipPurchased", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiMembershipPurchased)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipPurchased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMembershipPurchased is a log parse operation binding the contract event 0x3ce34c4daa4cc817ecd03b796bfda58751aa15ba8683a597a19d81add1969cfb.
//
// Solidity: event MembershipPurchased(address indexed member, uint256 indexed tokenId, uint8 tier, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseMembershipPurchased(log types.Log) (*DagoramembershipsabiMembershipPurchased, error) {
	event := new(DagoramembershipsabiMembershipPurchased)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiMembershipRenewedIterator is returned from FilterMembershipRenewed and is used to iterate over the raw logs and unpacked data for MembershipRenewed events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipRenewedIterator struct {
	Event *DagoramembershipsabiMembershipRenewed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiMembershipRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiMembershipRenewed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiMembershipRenewed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiMembershipRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiMembershipRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiMembershipRenewed represents a MembershipRenewed event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipRenewed struct {
	Member     common.Address
	TokenId    *big.Int
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMembershipRenewed is a free log retrieval operation binding the contract event 0xc2c37e3ea905ba32871ee4ffd3a7c43bc327abbae7b8852edeaf5bfda7152357.
//
// Solidity: event MembershipRenewed(address indexed member, uint256 indexed tokenId, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterMembershipRenewed(opts *bind.FilterOpts, member []common.Address, tokenId []*big.Int) (*DagoramembershipsabiMembershipRenewedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "MembershipRenewed", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiMembershipRenewedIterator{contract: _Dagoramembershipsabi.contract, event: "MembershipRenewed", logs: logs, sub: sub}, nil
}

// WatchMembershipRenewed is a free log subscription operation binding the contract event 0xc2c37e3ea905ba32871ee4ffd3a7c43bc327abbae7b8852edeaf5bfda7152357.
//
// Solidity: event MembershipRenewed(address indexed member, uint256 indexed tokenId, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchMembershipRenewed(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiMembershipRenewed, member []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "MembershipRenewed", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiMembershipRenewed)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipRenewed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMembershipRenewed is a log parse operation binding the contract event 0xc2c37e3ea905ba32871ee4ffd3a7c43bc327abbae7b8852edeaf5bfda7152357.
//
// Solidity: event MembershipRenewed(address indexed member, uint256 indexed tokenId, uint256 expiration)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseMembershipRenewed(log types.Log) (*DagoramembershipsabiMembershipRenewed, error) {
	event := new(DagoramembershipsabiMembershipRenewed)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiMembershipUpgradedIterator is returned from FilterMembershipUpgraded and is used to iterate over the raw logs and unpacked data for MembershipUpgraded events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipUpgradedIterator struct {
	Event *DagoramembershipsabiMembershipUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiMembershipUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiMembershipUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiMembershipUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiMembershipUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiMembershipUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiMembershipUpgraded represents a MembershipUpgraded event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiMembershipUpgraded struct {
	Member  common.Address
	TokenId *big.Int
	OldTier uint8
	NewTier uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMembershipUpgraded is a free log retrieval operation binding the contract event 0x2d3a0d643826fdc9f11a0ff0e3140800ce61aaba0525fa0fa59af6d445a2fc89.
//
// Solidity: event MembershipUpgraded(address indexed member, uint256 indexed tokenId, uint8 oldTier, uint8 newTier)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterMembershipUpgraded(opts *bind.FilterOpts, member []common.Address, tokenId []*big.Int) (*DagoramembershipsabiMembershipUpgradedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "MembershipUpgraded", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiMembershipUpgradedIterator{contract: _Dagoramembershipsabi.contract, event: "MembershipUpgraded", logs: logs, sub: sub}, nil
}

// WatchMembershipUpgraded is a free log subscription operation binding the contract event 0x2d3a0d643826fdc9f11a0ff0e3140800ce61aaba0525fa0fa59af6d445a2fc89.
//
// Solidity: event MembershipUpgraded(address indexed member, uint256 indexed tokenId, uint8 oldTier, uint8 newTier)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchMembershipUpgraded(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiMembershipUpgraded, member []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "MembershipUpgraded", memberRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiMembershipUpgraded)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMembershipUpgraded is a log parse operation binding the contract event 0x2d3a0d643826fdc9f11a0ff0e3140800ce61aaba0525fa0fa59af6d445a2fc89.
//
// Solidity: event MembershipUpgraded(address indexed member, uint256 indexed tokenId, uint8 oldTier, uint8 newTier)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseMembershipUpgraded(log types.Log) (*DagoramembershipsabiMembershipUpgraded, error) {
	event := new(DagoramembershipsabiMembershipUpgraded)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "MembershipUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiOwnershipTransferredIterator struct {
	Event *DagoramembershipsabiOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiOwnershipTransferred represents a OwnershipTransferred event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DagoramembershipsabiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiOwnershipTransferredIterator{contract: _Dagoramembershipsabi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiOwnershipTransferred)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseOwnershipTransferred(log types.Log) (*DagoramembershipsabiOwnershipTransferred, error) {
	event := new(DagoramembershipsabiOwnershipTransferred)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DagoramembershipsabiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiTransferIterator struct {
	Event *DagoramembershipsabiTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DagoramembershipsabiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DagoramembershipsabiTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DagoramembershipsabiTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DagoramembershipsabiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DagoramembershipsabiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DagoramembershipsabiTransfer represents a Transfer event raised by the Dagoramembershipsabi contract.
type DagoramembershipsabiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DagoramembershipsabiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DagoramembershipsabiTransferIterator{contract: _Dagoramembershipsabi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DagoramembershipsabiTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagoramembershipsabi.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DagoramembershipsabiTransfer)
				if err := _Dagoramembershipsabi.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dagoramembershipsabi *DagoramembershipsabiFilterer) ParseTransfer(log types.Log) (*DagoramembershipsabiTransfer, error) {
	event := new(DagoramembershipsabiTransfer)
	if err := _Dagoramembershipsabi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
