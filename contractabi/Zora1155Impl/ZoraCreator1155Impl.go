// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zora1155impl

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

// ICreatorRoyaltiesControlRoyaltyConfiguration is an auto generated low-level Go binding around an user-defined struct.
type ICreatorRoyaltiesControlRoyaltyConfiguration struct {
	RoyaltyMintSchedule uint32
	RoyaltyBPS          uint32
	RoyaltyRecipient    common.Address
}

// IZoraCreator1155TypesV1ContractConfig is an auto generated low-level Go binding around an user-defined struct.
type IZoraCreator1155TypesV1ContractConfig struct {
	Owner          common.Address
	Gap1           *big.Int
	FundsRecipient common.Address
	Gap2           *big.Int
	TransferHook   common.Address
	Gap3           *big.Int
}

// IZoraCreator1155TypesV1TokenData is an auto generated low-level Go binding around an user-defined struct.
type IZoraCreator1155TypesV1TokenData struct {
	Uri         string
	MaxSupply   *big.Int
	TotalMinted *big.Int
}

// Zora1155implMetaData contains all meta data concerning the Zora1155impl contract.
var Zora1155implMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mintFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESS_DELEGATECALL_TO_NON_CONTRACT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ADDRESS_LOW_LEVEL_CALL_FAILED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Burn_NotOwnerOrApproved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"CannotMintMoreTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mintFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintFee\",\"type\":\"uint256\"}],\"name\":\"CannotSendMintFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSetMintFeeToZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"Config_TransferHookNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_ACCOUNTS_AND_IDS_LENGTH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_ADDRESS_ZERO_IS_NOT_A_VALID_OWNER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_BURN_AMOUNT_EXCEEDS_BALANCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_BURN_FROM_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_CALLER_IS_NOT_TOKEN_OWNER_OR_APPROVED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_ERC1155RECEIVER_REJECTED_TOKENS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_IDS_AND_AMOUNTS_LENGTH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_INSUFFICIENT_BALANCE_FOR_TRANSFER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_MINT_TO_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_SETTING_APPROVAL_FOR_SELF\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_TRANSFER_TO_NON_ERC1155RECEIVER_IMPLEMENTER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155_TRANSFER_TO_ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967_NEW_IMPL_NOT_CONTRACT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967_NEW_IMPL_NOT_UUPS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967_UNSUPPORTED_PROXIABLEUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHWithdrawFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FUNCTION_MUST_BE_CALLED_THROUGH_ACTIVE_PROXY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FUNCTION_MUST_BE_CALLED_THROUGH_DELEGATECALL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contractValue\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawInsolvent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INITIALIZABLE_CONTRACT_ALREADY_INITIALIZED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INITIALIZABLE_CONTRACT_IS_NOT_INITIALIZING\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMintSchedule\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintFeeBPS\",\"type\":\"uint256\"}],\"name\":\"MintFeeCannotBeMoreThanZeroPointOneETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_InsolventSaleTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_TokenIDMintNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_UnknownCommand\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_ValueTransferFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerNeedsToBeAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NoRendererForToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"RendererNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Renderer_NotValidRendererContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"Sale_CannotCallNonSalesContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"TokenIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPS_UPGRADEABLE_MUST_NOT_BE_CALLED_THROUGH_DELEGATECALL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"UserMissingRoleForToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIZoraCreator1155.ConfigUpdate\",\"name\":\"updateType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap1\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap2\",\"type\":\"uint96\"},{\"internalType\":\"contractITransferHookReceiver\",\"name\":\"transferHook\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap3\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structIZoraCreator1155TypesV1.ContractConfig\",\"name\":\"newConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ContractMetadataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIRenderer1155\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"ContractRendererUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lastOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Purchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"renderer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RendererUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"SetupNewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"permissions\",\"type\":\"uint256\"}],\"name\":\"UpdatedPermissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structICreatorRoyaltiesControl.RoyaltyConfiguration\",\"name\":\"configuration\",\"type\":\"tuple\"}],\"name\":\"UpdatedRoyalties\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIZoraCreator1155TypesV1.TokenData\",\"name\":\"tokenData\",\"type\":\"tuple\"}],\"name\":\"UpdatedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONTRACT_BASE_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_ADMIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_FUNDS_MANAGER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_METADATA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_MINTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_BIT_SALES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissionBits\",\"type\":\"uint256\"}],\"name\":\"addPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"adminMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"adminMintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastTokenId\",\"type\":\"uint256\"}],\"name\":\"assumeLastTokenIdMatches\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"batchBalances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callRenderer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIMinter1155\",\"name\":\"salesConfig\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap1\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap2\",\"type\":\"uint96\"},{\"internalType\":\"contractITransferHookReceiver\",\"name\":\"transferHook\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"__gap3\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"customRenderers\",\"outputs\":[{\"internalType\":\"contractIRenderer1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCustomRenderer\",\"outputs\":[{\"internalType\":\"contractIRenderer1155\",\"name\":\"customRenderer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPermissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyalties\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"internalType\":\"structICreatorRoyaltiesControl.RoyaltyConfiguration\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"}],\"internalType\":\"structIZoraCreator1155TypesV1.TokenData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newContractURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"internalType\":\"structICreatorRoyaltiesControl.RoyaltyConfiguration\",\"name\":\"defaultRoyaltyConfiguration\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"setupActions\",\"type\":\"bytes[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"isAdminOrRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"metadataRendererContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMinter1155\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"minterArguments\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"permissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissionBits\",\"type\":\"uint256\"}],\"name\":\"removePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"royalties\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"name\":\"setFundsRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIRenderer1155\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"setTokenMetadataRenderer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransferHookReceiver\",\"name\":\"transferHook\",\"type\":\"address\"}],\"name\":\"setTransferHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"setupNewToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"supplyRoyaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_newName\",\"type\":\"string\"}],\"name\":\"updateContractMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"royaltyMintSchedule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"royaltyBPS\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"internalType\":\"structICreatorRoyaltiesControl.RoyaltyConfiguration\",\"name\":\"newConfiguration\",\"type\":\"tuple\"}],\"name\":\"updateRoyaltiesForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newURI\",\"type\":\"string\"}],\"name\":\"updateTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Zora1155implABI is the input ABI used to generate the binding from.
// Deprecated: Use Zora1155implMetaData.ABI instead.
var Zora1155implABI = Zora1155implMetaData.ABI

// Zora1155impl is an auto generated Go binding around an Ethereum contract.
type Zora1155impl struct {
	Zora1155implCaller     // Read-only binding to the contract
	Zora1155implTransactor // Write-only binding to the contract
	Zora1155implFilterer   // Log filterer for contract events
}

// Zora1155implCaller is an auto generated read-only Go binding around an Ethereum contract.
type Zora1155implCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Zora1155implTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Zora1155implTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Zora1155implFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Zora1155implFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Zora1155implSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Zora1155implSession struct {
	Contract     *Zora1155impl     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Zora1155implCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Zora1155implCallerSession struct {
	Contract *Zora1155implCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Zora1155implTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Zora1155implTransactorSession struct {
	Contract     *Zora1155implTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Zora1155implRaw is an auto generated low-level Go binding around an Ethereum contract.
type Zora1155implRaw struct {
	Contract *Zora1155impl // Generic contract binding to access the raw methods on
}

// Zora1155implCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Zora1155implCallerRaw struct {
	Contract *Zora1155implCaller // Generic read-only contract binding to access the raw methods on
}

// Zora1155implTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Zora1155implTransactorRaw struct {
	Contract *Zora1155implTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZora1155impl creates a new instance of Zora1155impl, bound to a specific deployed contract.
func NewZora1155impl(address common.Address, backend bind.ContractBackend) (*Zora1155impl, error) {
	contract, err := bindZora1155impl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zora1155impl{Zora1155implCaller: Zora1155implCaller{contract: contract}, Zora1155implTransactor: Zora1155implTransactor{contract: contract}, Zora1155implFilterer: Zora1155implFilterer{contract: contract}}, nil
}

// NewZora1155implCaller creates a new read-only instance of Zora1155impl, bound to a specific deployed contract.
func NewZora1155implCaller(address common.Address, caller bind.ContractCaller) (*Zora1155implCaller, error) {
	contract, err := bindZora1155impl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Zora1155implCaller{contract: contract}, nil
}

// NewZora1155implTransactor creates a new write-only instance of Zora1155impl, bound to a specific deployed contract.
func NewZora1155implTransactor(address common.Address, transactor bind.ContractTransactor) (*Zora1155implTransactor, error) {
	contract, err := bindZora1155impl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Zora1155implTransactor{contract: contract}, nil
}

// NewZora1155implFilterer creates a new log filterer instance of Zora1155impl, bound to a specific deployed contract.
func NewZora1155implFilterer(address common.Address, filterer bind.ContractFilterer) (*Zora1155implFilterer, error) {
	contract, err := bindZora1155impl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Zora1155implFilterer{contract: contract}, nil
}

// bindZora1155impl binds a generic wrapper to an already deployed contract.
func bindZora1155impl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Zora1155implMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zora1155impl *Zora1155implRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zora1155impl.Contract.Zora1155implCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zora1155impl *Zora1155implRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zora1155impl.Contract.Zora1155implTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zora1155impl *Zora1155implRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zora1155impl.Contract.Zora1155implTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zora1155impl *Zora1155implCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zora1155impl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zora1155impl *Zora1155implTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zora1155impl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zora1155impl *Zora1155implTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zora1155impl.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTBASEID is a free data retrieval call binding the contract method 0x01144201.
//
// Solidity: function CONTRACT_BASE_ID() view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) CONTRACTBASEID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "CONTRACT_BASE_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONTRACTBASEID is a free data retrieval call binding the contract method 0x01144201.
//
// Solidity: function CONTRACT_BASE_ID() view returns(uint256)
func (_Zora1155impl *Zora1155implSession) CONTRACTBASEID() (*big.Int, error) {
	return _Zora1155impl.Contract.CONTRACTBASEID(&_Zora1155impl.CallOpts)
}

// CONTRACTBASEID is a free data retrieval call binding the contract method 0x01144201.
//
// Solidity: function CONTRACT_BASE_ID() view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) CONTRACTBASEID() (*big.Int, error) {
	return _Zora1155impl.Contract.CONTRACTBASEID(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITADMIN is a free data retrieval call binding the contract method 0xc0464356.
//
// Solidity: function PERMISSION_BIT_ADMIN() view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) PERMISSIONBITADMIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "PERMISSION_BIT_ADMIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITADMIN is a free data retrieval call binding the contract method 0xc0464356.
//
// Solidity: function PERMISSION_BIT_ADMIN() view returns(uint256)
func (_Zora1155impl *Zora1155implSession) PERMISSIONBITADMIN() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITADMIN(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITADMIN is a free data retrieval call binding the contract method 0xc0464356.
//
// Solidity: function PERMISSION_BIT_ADMIN() view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) PERMISSIONBITADMIN() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITADMIN(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITFUNDSMANAGER is a free data retrieval call binding the contract method 0x929a7128.
//
// Solidity: function PERMISSION_BIT_FUNDS_MANAGER() view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) PERMISSIONBITFUNDSMANAGER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "PERMISSION_BIT_FUNDS_MANAGER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITFUNDSMANAGER is a free data retrieval call binding the contract method 0x929a7128.
//
// Solidity: function PERMISSION_BIT_FUNDS_MANAGER() view returns(uint256)
func (_Zora1155impl *Zora1155implSession) PERMISSIONBITFUNDSMANAGER() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITFUNDSMANAGER(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITFUNDSMANAGER is a free data retrieval call binding the contract method 0x929a7128.
//
// Solidity: function PERMISSION_BIT_FUNDS_MANAGER() view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) PERMISSIONBITFUNDSMANAGER() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITFUNDSMANAGER(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITMETADATA is a free data retrieval call binding the contract method 0xa453eaf0.
//
// Solidity: function PERMISSION_BIT_METADATA() view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) PERMISSIONBITMETADATA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "PERMISSION_BIT_METADATA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITMETADATA is a free data retrieval call binding the contract method 0xa453eaf0.
//
// Solidity: function PERMISSION_BIT_METADATA() view returns(uint256)
func (_Zora1155impl *Zora1155implSession) PERMISSIONBITMETADATA() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITMETADATA(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITMETADATA is a free data retrieval call binding the contract method 0xa453eaf0.
//
// Solidity: function PERMISSION_BIT_METADATA() view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) PERMISSIONBITMETADATA() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITMETADATA(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITMINTER is a free data retrieval call binding the contract method 0xf1b0d6bb.
//
// Solidity: function PERMISSION_BIT_MINTER() view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) PERMISSIONBITMINTER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "PERMISSION_BIT_MINTER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITMINTER is a free data retrieval call binding the contract method 0xf1b0d6bb.
//
// Solidity: function PERMISSION_BIT_MINTER() view returns(uint256)
func (_Zora1155impl *Zora1155implSession) PERMISSIONBITMINTER() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITMINTER(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITMINTER is a free data retrieval call binding the contract method 0xf1b0d6bb.
//
// Solidity: function PERMISSION_BIT_MINTER() view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) PERMISSIONBITMINTER() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITMINTER(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITSALES is a free data retrieval call binding the contract method 0x18711c7d.
//
// Solidity: function PERMISSION_BIT_SALES() view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) PERMISSIONBITSALES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "PERMISSION_BIT_SALES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERMISSIONBITSALES is a free data retrieval call binding the contract method 0x18711c7d.
//
// Solidity: function PERMISSION_BIT_SALES() view returns(uint256)
func (_Zora1155impl *Zora1155implSession) PERMISSIONBITSALES() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITSALES(&_Zora1155impl.CallOpts)
}

// PERMISSIONBITSALES is a free data retrieval call binding the contract method 0x18711c7d.
//
// Solidity: function PERMISSION_BIT_SALES() view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) PERMISSIONBITSALES() (*big.Int, error) {
	return _Zora1155impl.Contract.PERMISSIONBITSALES(&_Zora1155impl.CallOpts)
}

// AssumeLastTokenIdMatches is a free data retrieval call binding the contract method 0xe72878b4.
//
// Solidity: function assumeLastTokenIdMatches(uint256 lastTokenId) view returns()
func (_Zora1155impl *Zora1155implCaller) AssumeLastTokenIdMatches(opts *bind.CallOpts, lastTokenId *big.Int) error {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "assumeLastTokenIdMatches", lastTokenId)

	if err != nil {
		return err
	}

	return err

}

// AssumeLastTokenIdMatches is a free data retrieval call binding the contract method 0xe72878b4.
//
// Solidity: function assumeLastTokenIdMatches(uint256 lastTokenId) view returns()
func (_Zora1155impl *Zora1155implSession) AssumeLastTokenIdMatches(lastTokenId *big.Int) error {
	return _Zora1155impl.Contract.AssumeLastTokenIdMatches(&_Zora1155impl.CallOpts, lastTokenId)
}

// AssumeLastTokenIdMatches is a free data retrieval call binding the contract method 0xe72878b4.
//
// Solidity: function assumeLastTokenIdMatches(uint256 lastTokenId) view returns()
func (_Zora1155impl *Zora1155implCallerSession) AssumeLastTokenIdMatches(lastTokenId *big.Int) error {
	return _Zora1155impl.Contract.AssumeLastTokenIdMatches(&_Zora1155impl.CallOpts, lastTokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Zora1155impl *Zora1155implSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Zora1155impl.Contract.BalanceOf(&_Zora1155impl.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Zora1155impl.Contract.BalanceOf(&_Zora1155impl.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[] batchBalances)
func (_Zora1155impl *Zora1155implCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[] batchBalances)
func (_Zora1155impl *Zora1155implSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Zora1155impl.Contract.BalanceOfBatch(&_Zora1155impl.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[] batchBalances)
func (_Zora1155impl *Zora1155implCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Zora1155impl.Contract.BalanceOfBatch(&_Zora1155impl.CallOpts, accounts, ids)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, uint96 __gap1, address fundsRecipient, uint96 __gap2, address transferHook, uint96 __gap3)
func (_Zora1155impl *Zora1155implCaller) Config(opts *bind.CallOpts) (struct {
	Owner          common.Address
	Gap1           *big.Int
	FundsRecipient common.Address
	Gap2           *big.Int
	TransferHook   common.Address
	Gap3           *big.Int
}, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		Owner          common.Address
		Gap1           *big.Int
		FundsRecipient common.Address
		Gap2           *big.Int
		TransferHook   common.Address
		Gap3           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Gap1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FundsRecipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Gap2 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TransferHook = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Gap3 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, uint96 __gap1, address fundsRecipient, uint96 __gap2, address transferHook, uint96 __gap3)
func (_Zora1155impl *Zora1155implSession) Config() (struct {
	Owner          common.Address
	Gap1           *big.Int
	FundsRecipient common.Address
	Gap2           *big.Int
	TransferHook   common.Address
	Gap3           *big.Int
}, error) {
	return _Zora1155impl.Contract.Config(&_Zora1155impl.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, uint96 __gap1, address fundsRecipient, uint96 __gap2, address transferHook, uint96 __gap3)
func (_Zora1155impl *Zora1155implCallerSession) Config() (struct {
	Owner          common.Address
	Gap1           *big.Int
	FundsRecipient common.Address
	Gap2           *big.Int
	TransferHook   common.Address
	Gap3           *big.Int
}, error) {
	return _Zora1155impl.Contract.Config(&_Zora1155impl.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Zora1155impl *Zora1155implCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Zora1155impl *Zora1155implSession) ContractURI() (string, error) {
	return _Zora1155impl.Contract.ContractURI(&_Zora1155impl.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Zora1155impl *Zora1155implCallerSession) ContractURI() (string, error) {
	return _Zora1155impl.Contract.ContractURI(&_Zora1155impl.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(string)
func (_Zora1155impl *Zora1155implCaller) ContractVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(string)
func (_Zora1155impl *Zora1155implSession) ContractVersion() (string, error) {
	return _Zora1155impl.Contract.ContractVersion(&_Zora1155impl.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(string)
func (_Zora1155impl *Zora1155implCallerSession) ContractVersion() (string, error) {
	return _Zora1155impl.Contract.ContractVersion(&_Zora1155impl.CallOpts)
}

// CustomRenderers is a free data retrieval call binding the contract method 0xdd15e05f.
//
// Solidity: function customRenderers(uint256 ) view returns(address)
func (_Zora1155impl *Zora1155implCaller) CustomRenderers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "customRenderers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustomRenderers is a free data retrieval call binding the contract method 0xdd15e05f.
//
// Solidity: function customRenderers(uint256 ) view returns(address)
func (_Zora1155impl *Zora1155implSession) CustomRenderers(arg0 *big.Int) (common.Address, error) {
	return _Zora1155impl.Contract.CustomRenderers(&_Zora1155impl.CallOpts, arg0)
}

// CustomRenderers is a free data retrieval call binding the contract method 0xdd15e05f.
//
// Solidity: function customRenderers(uint256 ) view returns(address)
func (_Zora1155impl *Zora1155implCallerSession) CustomRenderers(arg0 *big.Int) (common.Address, error) {
	return _Zora1155impl.Contract.CustomRenderers(&_Zora1155impl.CallOpts, arg0)
}

// GetCustomRenderer is a free data retrieval call binding the contract method 0xe74d86c2.
//
// Solidity: function getCustomRenderer(uint256 tokenId) view returns(address customRenderer)
func (_Zora1155impl *Zora1155implCaller) GetCustomRenderer(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "getCustomRenderer", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCustomRenderer is a free data retrieval call binding the contract method 0xe74d86c2.
//
// Solidity: function getCustomRenderer(uint256 tokenId) view returns(address customRenderer)
func (_Zora1155impl *Zora1155implSession) GetCustomRenderer(tokenId *big.Int) (common.Address, error) {
	return _Zora1155impl.Contract.GetCustomRenderer(&_Zora1155impl.CallOpts, tokenId)
}

// GetCustomRenderer is a free data retrieval call binding the contract method 0xe74d86c2.
//
// Solidity: function getCustomRenderer(uint256 tokenId) view returns(address customRenderer)
func (_Zora1155impl *Zora1155implCallerSession) GetCustomRenderer(tokenId *big.Int) (common.Address, error) {
	return _Zora1155impl.Contract.GetCustomRenderer(&_Zora1155impl.CallOpts, tokenId)
}

// GetPermissions is a free data retrieval call binding the contract method 0x4913162d.
//
// Solidity: function getPermissions(uint256 tokenId, address user) view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) GetPermissions(opts *bind.CallOpts, tokenId *big.Int, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "getPermissions", tokenId, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPermissions is a free data retrieval call binding the contract method 0x4913162d.
//
// Solidity: function getPermissions(uint256 tokenId, address user) view returns(uint256)
func (_Zora1155impl *Zora1155implSession) GetPermissions(tokenId *big.Int, user common.Address) (*big.Int, error) {
	return _Zora1155impl.Contract.GetPermissions(&_Zora1155impl.CallOpts, tokenId, user)
}

// GetPermissions is a free data retrieval call binding the contract method 0x4913162d.
//
// Solidity: function getPermissions(uint256 tokenId, address user) view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) GetPermissions(tokenId *big.Int, user common.Address) (*big.Int, error) {
	return _Zora1155impl.Contract.GetPermissions(&_Zora1155impl.CallOpts, tokenId, user)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns((uint32,uint32,address))
func (_Zora1155impl *Zora1155implCaller) GetRoyalties(opts *bind.CallOpts, tokenId *big.Int) (ICreatorRoyaltiesControlRoyaltyConfiguration, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "getRoyalties", tokenId)

	if err != nil {
		return *new(ICreatorRoyaltiesControlRoyaltyConfiguration), err
	}

	out0 := *abi.ConvertType(out[0], new(ICreatorRoyaltiesControlRoyaltyConfiguration)).(*ICreatorRoyaltiesControlRoyaltyConfiguration)

	return out0, err

}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns((uint32,uint32,address))
func (_Zora1155impl *Zora1155implSession) GetRoyalties(tokenId *big.Int) (ICreatorRoyaltiesControlRoyaltyConfiguration, error) {
	return _Zora1155impl.Contract.GetRoyalties(&_Zora1155impl.CallOpts, tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 tokenId) view returns((uint32,uint32,address))
func (_Zora1155impl *Zora1155implCallerSession) GetRoyalties(tokenId *big.Int) (ICreatorRoyaltiesControlRoyaltyConfiguration, error) {
	return _Zora1155impl.Contract.GetRoyalties(&_Zora1155impl.CallOpts, tokenId)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 tokenId) view returns((string,uint256,uint256))
func (_Zora1155impl *Zora1155implCaller) GetTokenInfo(opts *bind.CallOpts, tokenId *big.Int) (IZoraCreator1155TypesV1TokenData, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "getTokenInfo", tokenId)

	if err != nil {
		return *new(IZoraCreator1155TypesV1TokenData), err
	}

	out0 := *abi.ConvertType(out[0], new(IZoraCreator1155TypesV1TokenData)).(*IZoraCreator1155TypesV1TokenData)

	return out0, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 tokenId) view returns((string,uint256,uint256))
func (_Zora1155impl *Zora1155implSession) GetTokenInfo(tokenId *big.Int) (IZoraCreator1155TypesV1TokenData, error) {
	return _Zora1155impl.Contract.GetTokenInfo(&_Zora1155impl.CallOpts, tokenId)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x8c7a63ae.
//
// Solidity: function getTokenInfo(uint256 tokenId) view returns((string,uint256,uint256))
func (_Zora1155impl *Zora1155implCallerSession) GetTokenInfo(tokenId *big.Int) (IZoraCreator1155TypesV1TokenData, error) {
	return _Zora1155impl.Contract.GetTokenInfo(&_Zora1155impl.CallOpts, tokenId)
}

// IsAdminOrRole is a free data retrieval call binding the contract method 0x23bd0386.
//
// Solidity: function isAdminOrRole(address user, uint256 tokenId, uint256 role) view returns(bool)
func (_Zora1155impl *Zora1155implCaller) IsAdminOrRole(opts *bind.CallOpts, user common.Address, tokenId *big.Int, role *big.Int) (bool, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "isAdminOrRole", user, tokenId, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdminOrRole is a free data retrieval call binding the contract method 0x23bd0386.
//
// Solidity: function isAdminOrRole(address user, uint256 tokenId, uint256 role) view returns(bool)
func (_Zora1155impl *Zora1155implSession) IsAdminOrRole(user common.Address, tokenId *big.Int, role *big.Int) (bool, error) {
	return _Zora1155impl.Contract.IsAdminOrRole(&_Zora1155impl.CallOpts, user, tokenId, role)
}

// IsAdminOrRole is a free data retrieval call binding the contract method 0x23bd0386.
//
// Solidity: function isAdminOrRole(address user, uint256 tokenId, uint256 role) view returns(bool)
func (_Zora1155impl *Zora1155implCallerSession) IsAdminOrRole(user common.Address, tokenId *big.Int, role *big.Int) (bool, error) {
	return _Zora1155impl.Contract.IsAdminOrRole(&_Zora1155impl.CallOpts, user, tokenId, role)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Zora1155impl *Zora1155implCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Zora1155impl *Zora1155implSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Zora1155impl.Contract.IsApprovedForAll(&_Zora1155impl.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Zora1155impl *Zora1155implCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Zora1155impl.Contract.IsApprovedForAll(&_Zora1155impl.CallOpts, account, operator)
}

// MetadataRendererContract is a free data retrieval call binding the contract method 0x69a5b302.
//
// Solidity: function metadataRendererContract(uint256 ) view returns(address)
func (_Zora1155impl *Zora1155implCaller) MetadataRendererContract(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "metadataRendererContract", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataRendererContract is a free data retrieval call binding the contract method 0x69a5b302.
//
// Solidity: function metadataRendererContract(uint256 ) view returns(address)
func (_Zora1155impl *Zora1155implSession) MetadataRendererContract(arg0 *big.Int) (common.Address, error) {
	return _Zora1155impl.Contract.MetadataRendererContract(&_Zora1155impl.CallOpts, arg0)
}

// MetadataRendererContract is a free data retrieval call binding the contract method 0x69a5b302.
//
// Solidity: function metadataRendererContract(uint256 ) view returns(address)
func (_Zora1155impl *Zora1155implCallerSession) MetadataRendererContract(arg0 *big.Int) (common.Address, error) {
	return _Zora1155impl.Contract.MetadataRendererContract(&_Zora1155impl.CallOpts, arg0)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) MintFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "mintFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_Zora1155impl *Zora1155implSession) MintFee() (*big.Int, error) {
	return _Zora1155impl.Contract.MintFee(&_Zora1155impl.CallOpts)
}

// MintFee is a free data retrieval call binding the contract method 0x13966db5.
//
// Solidity: function mintFee() view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) MintFee() (*big.Int, error) {
	return _Zora1155impl.Contract.MintFee(&_Zora1155impl.CallOpts)
}

// MintFeeRecipient is a free data retrieval call binding the contract method 0x765b0c36.
//
// Solidity: function mintFeeRecipient() view returns(address)
func (_Zora1155impl *Zora1155implCaller) MintFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "mintFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintFeeRecipient is a free data retrieval call binding the contract method 0x765b0c36.
//
// Solidity: function mintFeeRecipient() view returns(address)
func (_Zora1155impl *Zora1155implSession) MintFeeRecipient() (common.Address, error) {
	return _Zora1155impl.Contract.MintFeeRecipient(&_Zora1155impl.CallOpts)
}

// MintFeeRecipient is a free data retrieval call binding the contract method 0x765b0c36.
//
// Solidity: function mintFeeRecipient() view returns(address)
func (_Zora1155impl *Zora1155implCallerSession) MintFeeRecipient() (common.Address, error) {
	return _Zora1155impl.Contract.MintFeeRecipient(&_Zora1155impl.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zora1155impl *Zora1155implCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zora1155impl *Zora1155implSession) Name() (string, error) {
	return _Zora1155impl.Contract.Name(&_Zora1155impl.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Zora1155impl *Zora1155implCallerSession) Name() (string, error) {
	return _Zora1155impl.Contract.Name(&_Zora1155impl.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Zora1155impl *Zora1155implSession) NextTokenId() (*big.Int, error) {
	return _Zora1155impl.Contract.NextTokenId(&_Zora1155impl.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) NextTokenId() (*big.Int, error) {
	return _Zora1155impl.Contract.NextTokenId(&_Zora1155impl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zora1155impl *Zora1155implCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zora1155impl *Zora1155implSession) Owner() (common.Address, error) {
	return _Zora1155impl.Contract.Owner(&_Zora1155impl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Zora1155impl *Zora1155implCallerSession) Owner() (common.Address, error) {
	return _Zora1155impl.Contract.Owner(&_Zora1155impl.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0x300ecdb9.
//
// Solidity: function permissions(uint256 , address ) view returns(uint256)
func (_Zora1155impl *Zora1155implCaller) Permissions(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "permissions", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Permissions is a free data retrieval call binding the contract method 0x300ecdb9.
//
// Solidity: function permissions(uint256 , address ) view returns(uint256)
func (_Zora1155impl *Zora1155implSession) Permissions(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Zora1155impl.Contract.Permissions(&_Zora1155impl.CallOpts, arg0, arg1)
}

// Permissions is a free data retrieval call binding the contract method 0x300ecdb9.
//
// Solidity: function permissions(uint256 , address ) view returns(uint256)
func (_Zora1155impl *Zora1155implCallerSession) Permissions(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Zora1155impl.Contract.Permissions(&_Zora1155impl.CallOpts, arg0, arg1)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Zora1155impl *Zora1155implCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Zora1155impl *Zora1155implSession) ProxiableUUID() ([32]byte, error) {
	return _Zora1155impl.Contract.ProxiableUUID(&_Zora1155impl.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Zora1155impl *Zora1155implCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Zora1155impl.Contract.ProxiableUUID(&_Zora1155impl.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0x7f77f574.
//
// Solidity: function royalties(uint256 ) view returns(uint32 royaltyMintSchedule, uint32 royaltyBPS, address royaltyRecipient)
func (_Zora1155impl *Zora1155implCaller) Royalties(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoyaltyMintSchedule uint32
	RoyaltyBPS          uint32
	RoyaltyRecipient    common.Address
}, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "royalties", arg0)

	outstruct := new(struct {
		RoyaltyMintSchedule uint32
		RoyaltyBPS          uint32
		RoyaltyRecipient    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoyaltyMintSchedule = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.RoyaltyBPS = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.RoyaltyRecipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Royalties is a free data retrieval call binding the contract method 0x7f77f574.
//
// Solidity: function royalties(uint256 ) view returns(uint32 royaltyMintSchedule, uint32 royaltyBPS, address royaltyRecipient)
func (_Zora1155impl *Zora1155implSession) Royalties(arg0 *big.Int) (struct {
	RoyaltyMintSchedule uint32
	RoyaltyBPS          uint32
	RoyaltyRecipient    common.Address
}, error) {
	return _Zora1155impl.Contract.Royalties(&_Zora1155impl.CallOpts, arg0)
}

// Royalties is a free data retrieval call binding the contract method 0x7f77f574.
//
// Solidity: function royalties(uint256 ) view returns(uint32 royaltyMintSchedule, uint32 royaltyBPS, address royaltyRecipient)
func (_Zora1155impl *Zora1155implCallerSession) Royalties(arg0 *big.Int) (struct {
	RoyaltyMintSchedule uint32
	RoyaltyBPS          uint32
	RoyaltyRecipient    common.Address
}, error) {
	return _Zora1155impl.Contract.Royalties(&_Zora1155impl.CallOpts, arg0)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Zora1155impl *Zora1155implCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Zora1155impl *Zora1155implSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Zora1155impl.Contract.RoyaltyInfo(&_Zora1155impl.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Zora1155impl *Zora1155implCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Zora1155impl.Contract.RoyaltyInfo(&_Zora1155impl.CallOpts, tokenId, salePrice)
}

// SupplyRoyaltyInfo is a free data retrieval call binding the contract method 0x8621ea4b.
//
// Solidity: function supplyRoyaltyInfo(uint256 tokenId, uint256 totalSupply, uint256 mintAmount) view returns(address receiver, uint256 royaltyAmount)
func (_Zora1155impl *Zora1155implCaller) SupplyRoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, totalSupply *big.Int, mintAmount *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "supplyRoyaltyInfo", tokenId, totalSupply, mintAmount)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SupplyRoyaltyInfo is a free data retrieval call binding the contract method 0x8621ea4b.
//
// Solidity: function supplyRoyaltyInfo(uint256 tokenId, uint256 totalSupply, uint256 mintAmount) view returns(address receiver, uint256 royaltyAmount)
func (_Zora1155impl *Zora1155implSession) SupplyRoyaltyInfo(tokenId *big.Int, totalSupply *big.Int, mintAmount *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Zora1155impl.Contract.SupplyRoyaltyInfo(&_Zora1155impl.CallOpts, tokenId, totalSupply, mintAmount)
}

// SupplyRoyaltyInfo is a free data retrieval call binding the contract method 0x8621ea4b.
//
// Solidity: function supplyRoyaltyInfo(uint256 tokenId, uint256 totalSupply, uint256 mintAmount) view returns(address receiver, uint256 royaltyAmount)
func (_Zora1155impl *Zora1155implCallerSession) SupplyRoyaltyInfo(tokenId *big.Int, totalSupply *big.Int, mintAmount *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Zora1155impl.Contract.SupplyRoyaltyInfo(&_Zora1155impl.CallOpts, tokenId, totalSupply, mintAmount)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zora1155impl *Zora1155implCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zora1155impl *Zora1155implSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zora1155impl.Contract.SupportsInterface(&_Zora1155impl.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zora1155impl *Zora1155implCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zora1155impl.Contract.SupportsInterface(&_Zora1155impl.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Zora1155impl *Zora1155implCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Zora1155impl *Zora1155implSession) Symbol() (string, error) {
	return _Zora1155impl.Contract.Symbol(&_Zora1155impl.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Zora1155impl *Zora1155implCallerSession) Symbol() (string, error) {
	return _Zora1155impl.Contract.Symbol(&_Zora1155impl.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Zora1155impl *Zora1155implCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Zora1155impl.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Zora1155impl *Zora1155implSession) Uri(tokenId *big.Int) (string, error) {
	return _Zora1155impl.Contract.Uri(&_Zora1155impl.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Zora1155impl *Zora1155implCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _Zora1155impl.Contract.Uri(&_Zora1155impl.CallOpts, tokenId)
}

// AddPermission is a paid mutator transaction binding the contract method 0x8ec998a0.
//
// Solidity: function addPermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zora1155impl *Zora1155implTransactor) AddPermission(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "addPermission", tokenId, user, permissionBits)
}

// AddPermission is a paid mutator transaction binding the contract method 0x8ec998a0.
//
// Solidity: function addPermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zora1155impl *Zora1155implSession) AddPermission(tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zora1155impl.Contract.AddPermission(&_Zora1155impl.TransactOpts, tokenId, user, permissionBits)
}

// AddPermission is a paid mutator transaction binding the contract method 0x8ec998a0.
//
// Solidity: function addPermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zora1155impl *Zora1155implTransactorSession) AddPermission(tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zora1155impl.Contract.AddPermission(&_Zora1155impl.TransactOpts, tokenId, user, permissionBits)
}

// AdminMint is a paid mutator transaction binding the contract method 0xc238d1ee.
//
// Solidity: function adminMint(address recipient, uint256 tokenId, uint256 quantity, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactor) AdminMint(opts *bind.TransactOpts, recipient common.Address, tokenId *big.Int, quantity *big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "adminMint", recipient, tokenId, quantity, data)
}

// AdminMint is a paid mutator transaction binding the contract method 0xc238d1ee.
//
// Solidity: function adminMint(address recipient, uint256 tokenId, uint256 quantity, bytes data) returns()
func (_Zora1155impl *Zora1155implSession) AdminMint(recipient common.Address, tokenId *big.Int, quantity *big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.AdminMint(&_Zora1155impl.TransactOpts, recipient, tokenId, quantity, data)
}

// AdminMint is a paid mutator transaction binding the contract method 0xc238d1ee.
//
// Solidity: function adminMint(address recipient, uint256 tokenId, uint256 quantity, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactorSession) AdminMint(recipient common.Address, tokenId *big.Int, quantity *big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.AdminMint(&_Zora1155impl.TransactOpts, recipient, tokenId, quantity, data)
}

// AdminMintBatch is a paid mutator transaction binding the contract method 0xd1ad846b.
//
// Solidity: function adminMintBatch(address recipient, uint256[] tokenIds, uint256[] quantities, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactor) AdminMintBatch(opts *bind.TransactOpts, recipient common.Address, tokenIds []*big.Int, quantities []*big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "adminMintBatch", recipient, tokenIds, quantities, data)
}

// AdminMintBatch is a paid mutator transaction binding the contract method 0xd1ad846b.
//
// Solidity: function adminMintBatch(address recipient, uint256[] tokenIds, uint256[] quantities, bytes data) returns()
func (_Zora1155impl *Zora1155implSession) AdminMintBatch(recipient common.Address, tokenIds []*big.Int, quantities []*big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.AdminMintBatch(&_Zora1155impl.TransactOpts, recipient, tokenIds, quantities, data)
}

// AdminMintBatch is a paid mutator transaction binding the contract method 0xd1ad846b.
//
// Solidity: function adminMintBatch(address recipient, uint256[] tokenIds, uint256[] quantities, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactorSession) AdminMintBatch(recipient common.Address, tokenIds []*big.Int, quantities []*big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.AdminMintBatch(&_Zora1155impl.TransactOpts, recipient, tokenIds, quantities, data)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address from, uint256[] tokenIds, uint256[] amounts) returns()
func (_Zora1155impl *Zora1155implTransactor) BurnBatch(opts *bind.TransactOpts, from common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "burnBatch", from, tokenIds, amounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address from, uint256[] tokenIds, uint256[] amounts) returns()
func (_Zora1155impl *Zora1155implSession) BurnBatch(from common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Zora1155impl.Contract.BurnBatch(&_Zora1155impl.TransactOpts, from, tokenIds, amounts)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address from, uint256[] tokenIds, uint256[] amounts) returns()
func (_Zora1155impl *Zora1155implTransactorSession) BurnBatch(from common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Zora1155impl.Contract.BurnBatch(&_Zora1155impl.TransactOpts, from, tokenIds, amounts)
}

// CallRenderer is a paid mutator transaction binding the contract method 0x9c5c63c9.
//
// Solidity: function callRenderer(uint256 tokenId, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactor) CallRenderer(opts *bind.TransactOpts, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "callRenderer", tokenId, data)
}

// CallRenderer is a paid mutator transaction binding the contract method 0x9c5c63c9.
//
// Solidity: function callRenderer(uint256 tokenId, bytes data) returns()
func (_Zora1155impl *Zora1155implSession) CallRenderer(tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.CallRenderer(&_Zora1155impl.TransactOpts, tokenId, data)
}

// CallRenderer is a paid mutator transaction binding the contract method 0x9c5c63c9.
//
// Solidity: function callRenderer(uint256 tokenId, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactorSession) CallRenderer(tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.CallRenderer(&_Zora1155impl.TransactOpts, tokenId, data)
}

// CallSale is a paid mutator transaction binding the contract method 0xd904b94a.
//
// Solidity: function callSale(uint256 tokenId, address salesConfig, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactor) CallSale(opts *bind.TransactOpts, tokenId *big.Int, salesConfig common.Address, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "callSale", tokenId, salesConfig, data)
}

// CallSale is a paid mutator transaction binding the contract method 0xd904b94a.
//
// Solidity: function callSale(uint256 tokenId, address salesConfig, bytes data) returns()
func (_Zora1155impl *Zora1155implSession) CallSale(tokenId *big.Int, salesConfig common.Address, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.CallSale(&_Zora1155impl.TransactOpts, tokenId, salesConfig, data)
}

// CallSale is a paid mutator transaction binding the contract method 0xd904b94a.
//
// Solidity: function callSale(uint256 tokenId, address salesConfig, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactorSession) CallSale(tokenId *big.Int, salesConfig common.Address, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.CallSale(&_Zora1155impl.TransactOpts, tokenId, salesConfig, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a08eb4c.
//
// Solidity: function initialize(string contractName, string newContractURI, (uint32,uint32,address) defaultRoyaltyConfiguration, address defaultAdmin, bytes[] setupActions) returns()
func (_Zora1155impl *Zora1155implTransactor) Initialize(opts *bind.TransactOpts, contractName string, newContractURI string, defaultRoyaltyConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration, defaultAdmin common.Address, setupActions [][]byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "initialize", contractName, newContractURI, defaultRoyaltyConfiguration, defaultAdmin, setupActions)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a08eb4c.
//
// Solidity: function initialize(string contractName, string newContractURI, (uint32,uint32,address) defaultRoyaltyConfiguration, address defaultAdmin, bytes[] setupActions) returns()
func (_Zora1155impl *Zora1155implSession) Initialize(contractName string, newContractURI string, defaultRoyaltyConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration, defaultAdmin common.Address, setupActions [][]byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.Initialize(&_Zora1155impl.TransactOpts, contractName, newContractURI, defaultRoyaltyConfiguration, defaultAdmin, setupActions)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a08eb4c.
//
// Solidity: function initialize(string contractName, string newContractURI, (uint32,uint32,address) defaultRoyaltyConfiguration, address defaultAdmin, bytes[] setupActions) returns()
func (_Zora1155impl *Zora1155implTransactorSession) Initialize(contractName string, newContractURI string, defaultRoyaltyConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration, defaultAdmin common.Address, setupActions [][]byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.Initialize(&_Zora1155impl.TransactOpts, contractName, newContractURI, defaultRoyaltyConfiguration, defaultAdmin, setupActions)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address minter, uint256 tokenId, uint256 quantity, bytes minterArguments) payable returns()
func (_Zora1155impl *Zora1155implTransactor) Mint(opts *bind.TransactOpts, minter common.Address, tokenId *big.Int, quantity *big.Int, minterArguments []byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "mint", minter, tokenId, quantity, minterArguments)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address minter, uint256 tokenId, uint256 quantity, bytes minterArguments) payable returns()
func (_Zora1155impl *Zora1155implSession) Mint(minter common.Address, tokenId *big.Int, quantity *big.Int, minterArguments []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.Mint(&_Zora1155impl.TransactOpts, minter, tokenId, quantity, minterArguments)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address minter, uint256 tokenId, uint256 quantity, bytes minterArguments) payable returns()
func (_Zora1155impl *Zora1155implTransactorSession) Mint(minter common.Address, tokenId *big.Int, quantity *big.Int, minterArguments []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.Mint(&_Zora1155impl.TransactOpts, minter, tokenId, quantity, minterArguments)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Zora1155impl *Zora1155implTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Zora1155impl *Zora1155implSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.Multicall(&_Zora1155impl.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Zora1155impl *Zora1155implTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.Multicall(&_Zora1155impl.TransactOpts, data)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x5d0f6cba.
//
// Solidity: function removePermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zora1155impl *Zora1155implTransactor) RemovePermission(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "removePermission", tokenId, user, permissionBits)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x5d0f6cba.
//
// Solidity: function removePermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zora1155impl *Zora1155implSession) RemovePermission(tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zora1155impl.Contract.RemovePermission(&_Zora1155impl.TransactOpts, tokenId, user, permissionBits)
}

// RemovePermission is a paid mutator transaction binding the contract method 0x5d0f6cba.
//
// Solidity: function removePermission(uint256 tokenId, address user, uint256 permissionBits) returns()
func (_Zora1155impl *Zora1155implTransactorSession) RemovePermission(tokenId *big.Int, user common.Address, permissionBits *big.Int) (*types.Transaction, error) {
	return _Zora1155impl.Contract.RemovePermission(&_Zora1155impl.TransactOpts, tokenId, user, permissionBits)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Zora1155impl *Zora1155implSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SafeBatchTransferFrom(&_Zora1155impl.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SafeBatchTransferFrom(&_Zora1155impl.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Zora1155impl *Zora1155implSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SafeTransferFrom(&_Zora1155impl.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Zora1155impl *Zora1155implTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SafeTransferFrom(&_Zora1155impl.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Zora1155impl *Zora1155implTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Zora1155impl *Zora1155implSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetApprovalForAll(&_Zora1155impl.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Zora1155impl *Zora1155implTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetApprovalForAll(&_Zora1155impl.TransactOpts, operator, approved)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address fundsRecipient) returns()
func (_Zora1155impl *Zora1155implTransactor) SetFundsRecipient(opts *bind.TransactOpts, fundsRecipient common.Address) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "setFundsRecipient", fundsRecipient)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address fundsRecipient) returns()
func (_Zora1155impl *Zora1155implSession) SetFundsRecipient(fundsRecipient common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetFundsRecipient(&_Zora1155impl.TransactOpts, fundsRecipient)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address fundsRecipient) returns()
func (_Zora1155impl *Zora1155implTransactorSession) SetFundsRecipient(fundsRecipient common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetFundsRecipient(&_Zora1155impl.TransactOpts, fundsRecipient)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Zora1155impl *Zora1155implTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Zora1155impl *Zora1155implSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetOwner(&_Zora1155impl.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Zora1155impl *Zora1155implTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetOwner(&_Zora1155impl.TransactOpts, newOwner)
}

// SetTokenMetadataRenderer is a paid mutator transaction binding the contract method 0x6661a9ba.
//
// Solidity: function setTokenMetadataRenderer(uint256 tokenId, address renderer) returns()
func (_Zora1155impl *Zora1155implTransactor) SetTokenMetadataRenderer(opts *bind.TransactOpts, tokenId *big.Int, renderer common.Address) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "setTokenMetadataRenderer", tokenId, renderer)
}

// SetTokenMetadataRenderer is a paid mutator transaction binding the contract method 0x6661a9ba.
//
// Solidity: function setTokenMetadataRenderer(uint256 tokenId, address renderer) returns()
func (_Zora1155impl *Zora1155implSession) SetTokenMetadataRenderer(tokenId *big.Int, renderer common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetTokenMetadataRenderer(&_Zora1155impl.TransactOpts, tokenId, renderer)
}

// SetTokenMetadataRenderer is a paid mutator transaction binding the contract method 0x6661a9ba.
//
// Solidity: function setTokenMetadataRenderer(uint256 tokenId, address renderer) returns()
func (_Zora1155impl *Zora1155implTransactorSession) SetTokenMetadataRenderer(tokenId *big.Int, renderer common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetTokenMetadataRenderer(&_Zora1155impl.TransactOpts, tokenId, renderer)
}

// SetTransferHook is a paid mutator transaction binding the contract method 0x7f2dc61c.
//
// Solidity: function setTransferHook(address transferHook) returns()
func (_Zora1155impl *Zora1155implTransactor) SetTransferHook(opts *bind.TransactOpts, transferHook common.Address) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "setTransferHook", transferHook)
}

// SetTransferHook is a paid mutator transaction binding the contract method 0x7f2dc61c.
//
// Solidity: function setTransferHook(address transferHook) returns()
func (_Zora1155impl *Zora1155implSession) SetTransferHook(transferHook common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetTransferHook(&_Zora1155impl.TransactOpts, transferHook)
}

// SetTransferHook is a paid mutator transaction binding the contract method 0x7f2dc61c.
//
// Solidity: function setTransferHook(address transferHook) returns()
func (_Zora1155impl *Zora1155implTransactorSession) SetTransferHook(transferHook common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetTransferHook(&_Zora1155impl.TransactOpts, transferHook)
}

// SetupNewToken is a paid mutator transaction binding the contract method 0xd258609a.
//
// Solidity: function setupNewToken(string newURI, uint256 maxSupply) returns(uint256)
func (_Zora1155impl *Zora1155implTransactor) SetupNewToken(opts *bind.TransactOpts, newURI string, maxSupply *big.Int) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "setupNewToken", newURI, maxSupply)
}

// SetupNewToken is a paid mutator transaction binding the contract method 0xd258609a.
//
// Solidity: function setupNewToken(string newURI, uint256 maxSupply) returns(uint256)
func (_Zora1155impl *Zora1155implSession) SetupNewToken(newURI string, maxSupply *big.Int) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetupNewToken(&_Zora1155impl.TransactOpts, newURI, maxSupply)
}

// SetupNewToken is a paid mutator transaction binding the contract method 0xd258609a.
//
// Solidity: function setupNewToken(string newURI, uint256 maxSupply) returns(uint256)
func (_Zora1155impl *Zora1155implTransactorSession) SetupNewToken(newURI string, maxSupply *big.Int) (*types.Transaction, error) {
	return _Zora1155impl.Contract.SetupNewToken(&_Zora1155impl.TransactOpts, newURI, maxSupply)
}

// UpdateContractMetadata is a paid mutator transaction binding the contract method 0xef71c82e.
//
// Solidity: function updateContractMetadata(string _newURI, string _newName) returns()
func (_Zora1155impl *Zora1155implTransactor) UpdateContractMetadata(opts *bind.TransactOpts, _newURI string, _newName string) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "updateContractMetadata", _newURI, _newName)
}

// UpdateContractMetadata is a paid mutator transaction binding the contract method 0xef71c82e.
//
// Solidity: function updateContractMetadata(string _newURI, string _newName) returns()
func (_Zora1155impl *Zora1155implSession) UpdateContractMetadata(_newURI string, _newName string) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpdateContractMetadata(&_Zora1155impl.TransactOpts, _newURI, _newName)
}

// UpdateContractMetadata is a paid mutator transaction binding the contract method 0xef71c82e.
//
// Solidity: function updateContractMetadata(string _newURI, string _newName) returns()
func (_Zora1155impl *Zora1155implTransactorSession) UpdateContractMetadata(_newURI string, _newName string) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpdateContractMetadata(&_Zora1155impl.TransactOpts, _newURI, _newName)
}

// UpdateRoyaltiesForToken is a paid mutator transaction binding the contract method 0xafed7e9e.
//
// Solidity: function updateRoyaltiesForToken(uint256 tokenId, (uint32,uint32,address) newConfiguration) returns()
func (_Zora1155impl *Zora1155implTransactor) UpdateRoyaltiesForToken(opts *bind.TransactOpts, tokenId *big.Int, newConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "updateRoyaltiesForToken", tokenId, newConfiguration)
}

// UpdateRoyaltiesForToken is a paid mutator transaction binding the contract method 0xafed7e9e.
//
// Solidity: function updateRoyaltiesForToken(uint256 tokenId, (uint32,uint32,address) newConfiguration) returns()
func (_Zora1155impl *Zora1155implSession) UpdateRoyaltiesForToken(tokenId *big.Int, newConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpdateRoyaltiesForToken(&_Zora1155impl.TransactOpts, tokenId, newConfiguration)
}

// UpdateRoyaltiesForToken is a paid mutator transaction binding the contract method 0xafed7e9e.
//
// Solidity: function updateRoyaltiesForToken(uint256 tokenId, (uint32,uint32,address) newConfiguration) returns()
func (_Zora1155impl *Zora1155implTransactorSession) UpdateRoyaltiesForToken(tokenId *big.Int, newConfiguration ICreatorRoyaltiesControlRoyaltyConfiguration) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpdateRoyaltiesForToken(&_Zora1155impl.TransactOpts, tokenId, newConfiguration)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string _newURI) returns()
func (_Zora1155impl *Zora1155implTransactor) UpdateTokenURI(opts *bind.TransactOpts, tokenId *big.Int, _newURI string) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "updateTokenURI", tokenId, _newURI)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string _newURI) returns()
func (_Zora1155impl *Zora1155implSession) UpdateTokenURI(tokenId *big.Int, _newURI string) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpdateTokenURI(&_Zora1155impl.TransactOpts, tokenId, _newURI)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 tokenId, string _newURI) returns()
func (_Zora1155impl *Zora1155implTransactorSession) UpdateTokenURI(tokenId *big.Int, _newURI string) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpdateTokenURI(&_Zora1155impl.TransactOpts, tokenId, _newURI)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Zora1155impl *Zora1155implTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Zora1155impl *Zora1155implSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpgradeTo(&_Zora1155impl.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Zora1155impl *Zora1155implTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpgradeTo(&_Zora1155impl.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Zora1155impl *Zora1155implTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Zora1155impl *Zora1155implSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpgradeToAndCall(&_Zora1155impl.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Zora1155impl *Zora1155implTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Zora1155impl.Contract.UpgradeToAndCall(&_Zora1155impl.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Zora1155impl *Zora1155implTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zora1155impl.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Zora1155impl *Zora1155implSession) Withdraw() (*types.Transaction, error) {
	return _Zora1155impl.Contract.Withdraw(&_Zora1155impl.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Zora1155impl *Zora1155implTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Zora1155impl.Contract.Withdraw(&_Zora1155impl.TransactOpts)
}

// Zora1155implAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Zora1155impl contract.
type Zora1155implAdminChangedIterator struct {
	Event *Zora1155implAdminChanged // Event containing the contract specifics and raw log

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
func (it *Zora1155implAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implAdminChanged)
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
		it.Event = new(Zora1155implAdminChanged)
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
func (it *Zora1155implAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implAdminChanged represents a AdminChanged event raised by the Zora1155impl contract.
type Zora1155implAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Zora1155impl *Zora1155implFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*Zora1155implAdminChangedIterator, error) {

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &Zora1155implAdminChangedIterator{contract: _Zora1155impl.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Zora1155impl *Zora1155implFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *Zora1155implAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implAdminChanged)
				if err := _Zora1155impl.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Zora1155impl *Zora1155implFilterer) ParseAdminChanged(log types.Log) (*Zora1155implAdminChanged, error) {
	event := new(Zora1155implAdminChanged)
	if err := _Zora1155impl.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Zora1155impl contract.
type Zora1155implApprovalForAllIterator struct {
	Event *Zora1155implApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Zora1155implApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implApprovalForAll)
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
		it.Event = new(Zora1155implApprovalForAll)
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
func (it *Zora1155implApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implApprovalForAll represents a ApprovalForAll event raised by the Zora1155impl contract.
type Zora1155implApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Zora1155impl *Zora1155implFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Zora1155implApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implApprovalForAllIterator{contract: _Zora1155impl.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Zora1155impl *Zora1155implFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Zora1155implApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implApprovalForAll)
				if err := _Zora1155impl.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Zora1155impl *Zora1155implFilterer) ParseApprovalForAll(log types.Log) (*Zora1155implApprovalForAll, error) {
	event := new(Zora1155implApprovalForAll)
	if err := _Zora1155impl.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Zora1155impl contract.
type Zora1155implBeaconUpgradedIterator struct {
	Event *Zora1155implBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *Zora1155implBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implBeaconUpgraded)
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
		it.Event = new(Zora1155implBeaconUpgraded)
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
func (it *Zora1155implBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implBeaconUpgraded represents a BeaconUpgraded event raised by the Zora1155impl contract.
type Zora1155implBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Zora1155impl *Zora1155implFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*Zora1155implBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implBeaconUpgradedIterator{contract: _Zora1155impl.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Zora1155impl *Zora1155implFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *Zora1155implBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implBeaconUpgraded)
				if err := _Zora1155impl.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Zora1155impl *Zora1155implFilterer) ParseBeaconUpgraded(log types.Log) (*Zora1155implBeaconUpgraded, error) {
	event := new(Zora1155implBeaconUpgraded)
	if err := _Zora1155impl.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implConfigUpdatedIterator is returned from FilterConfigUpdated and is used to iterate over the raw logs and unpacked data for ConfigUpdated events raised by the Zora1155impl contract.
type Zora1155implConfigUpdatedIterator struct {
	Event *Zora1155implConfigUpdated // Event containing the contract specifics and raw log

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
func (it *Zora1155implConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implConfigUpdated)
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
		it.Event = new(Zora1155implConfigUpdated)
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
func (it *Zora1155implConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implConfigUpdated represents a ConfigUpdated event raised by the Zora1155impl contract.
type Zora1155implConfigUpdated struct {
	Updater    common.Address
	UpdateType uint8
	NewConfig  IZoraCreator1155TypesV1ContractConfig
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdated is a free log retrieval operation binding the contract event 0x3be6d3a1d957610f7e900c66889b874cdc9f0c22901aa8be6ec3d2d04c14ca0f.
//
// Solidity: event ConfigUpdated(address indexed updater, uint8 indexed updateType, (address,uint96,address,uint96,address,uint96) newConfig)
func (_Zora1155impl *Zora1155implFilterer) FilterConfigUpdated(opts *bind.FilterOpts, updater []common.Address, updateType []uint8) (*Zora1155implConfigUpdatedIterator, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var updateTypeRule []interface{}
	for _, updateTypeItem := range updateType {
		updateTypeRule = append(updateTypeRule, updateTypeItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "ConfigUpdated", updaterRule, updateTypeRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implConfigUpdatedIterator{contract: _Zora1155impl.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigUpdated is a free log subscription operation binding the contract event 0x3be6d3a1d957610f7e900c66889b874cdc9f0c22901aa8be6ec3d2d04c14ca0f.
//
// Solidity: event ConfigUpdated(address indexed updater, uint8 indexed updateType, (address,uint96,address,uint96,address,uint96) newConfig)
func (_Zora1155impl *Zora1155implFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *Zora1155implConfigUpdated, updater []common.Address, updateType []uint8) (event.Subscription, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}
	var updateTypeRule []interface{}
	for _, updateTypeItem := range updateType {
		updateTypeRule = append(updateTypeRule, updateTypeItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "ConfigUpdated", updaterRule, updateTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implConfigUpdated)
				if err := _Zora1155impl.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

// ParseConfigUpdated is a log parse operation binding the contract event 0x3be6d3a1d957610f7e900c66889b874cdc9f0c22901aa8be6ec3d2d04c14ca0f.
//
// Solidity: event ConfigUpdated(address indexed updater, uint8 indexed updateType, (address,uint96,address,uint96,address,uint96) newConfig)
func (_Zora1155impl *Zora1155implFilterer) ParseConfigUpdated(log types.Log) (*Zora1155implConfigUpdated, error) {
	event := new(Zora1155implConfigUpdated)
	if err := _Zora1155impl.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implContractMetadataUpdatedIterator is returned from FilterContractMetadataUpdated and is used to iterate over the raw logs and unpacked data for ContractMetadataUpdated events raised by the Zora1155impl contract.
type Zora1155implContractMetadataUpdatedIterator struct {
	Event *Zora1155implContractMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *Zora1155implContractMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implContractMetadataUpdated)
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
		it.Event = new(Zora1155implContractMetadataUpdated)
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
func (it *Zora1155implContractMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implContractMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implContractMetadataUpdated represents a ContractMetadataUpdated event raised by the Zora1155impl contract.
type Zora1155implContractMetadataUpdated struct {
	Updater common.Address
	Uri     string
	Name    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractMetadataUpdated is a free log retrieval operation binding the contract event 0x74b7c2afa3f89c562b59674a101e2c48bceeb27cdb620afefa14446f1ffa487b.
//
// Solidity: event ContractMetadataUpdated(address indexed updater, string uri, string name)
func (_Zora1155impl *Zora1155implFilterer) FilterContractMetadataUpdated(opts *bind.FilterOpts, updater []common.Address) (*Zora1155implContractMetadataUpdatedIterator, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "ContractMetadataUpdated", updaterRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implContractMetadataUpdatedIterator{contract: _Zora1155impl.contract, event: "ContractMetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchContractMetadataUpdated is a free log subscription operation binding the contract event 0x74b7c2afa3f89c562b59674a101e2c48bceeb27cdb620afefa14446f1ffa487b.
//
// Solidity: event ContractMetadataUpdated(address indexed updater, string uri, string name)
func (_Zora1155impl *Zora1155implFilterer) WatchContractMetadataUpdated(opts *bind.WatchOpts, sink chan<- *Zora1155implContractMetadataUpdated, updater []common.Address) (event.Subscription, error) {

	var updaterRule []interface{}
	for _, updaterItem := range updater {
		updaterRule = append(updaterRule, updaterItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "ContractMetadataUpdated", updaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implContractMetadataUpdated)
				if err := _Zora1155impl.contract.UnpackLog(event, "ContractMetadataUpdated", log); err != nil {
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

// ParseContractMetadataUpdated is a log parse operation binding the contract event 0x74b7c2afa3f89c562b59674a101e2c48bceeb27cdb620afefa14446f1ffa487b.
//
// Solidity: event ContractMetadataUpdated(address indexed updater, string uri, string name)
func (_Zora1155impl *Zora1155implFilterer) ParseContractMetadataUpdated(log types.Log) (*Zora1155implContractMetadataUpdated, error) {
	event := new(Zora1155implContractMetadataUpdated)
	if err := _Zora1155impl.contract.UnpackLog(event, "ContractMetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implContractRendererUpdatedIterator is returned from FilterContractRendererUpdated and is used to iterate over the raw logs and unpacked data for ContractRendererUpdated events raised by the Zora1155impl contract.
type Zora1155implContractRendererUpdatedIterator struct {
	Event *Zora1155implContractRendererUpdated // Event containing the contract specifics and raw log

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
func (it *Zora1155implContractRendererUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implContractRendererUpdated)
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
		it.Event = new(Zora1155implContractRendererUpdated)
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
func (it *Zora1155implContractRendererUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implContractRendererUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implContractRendererUpdated represents a ContractRendererUpdated event raised by the Zora1155impl contract.
type Zora1155implContractRendererUpdated struct {
	Renderer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractRendererUpdated is a free log retrieval operation binding the contract event 0x56e810c8cae84731149f628981d25769a084570b9ba6eebf3c32879e3dce5609.
//
// Solidity: event ContractRendererUpdated(address renderer)
func (_Zora1155impl *Zora1155implFilterer) FilterContractRendererUpdated(opts *bind.FilterOpts) (*Zora1155implContractRendererUpdatedIterator, error) {

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "ContractRendererUpdated")
	if err != nil {
		return nil, err
	}
	return &Zora1155implContractRendererUpdatedIterator{contract: _Zora1155impl.contract, event: "ContractRendererUpdated", logs: logs, sub: sub}, nil
}

// WatchContractRendererUpdated is a free log subscription operation binding the contract event 0x56e810c8cae84731149f628981d25769a084570b9ba6eebf3c32879e3dce5609.
//
// Solidity: event ContractRendererUpdated(address renderer)
func (_Zora1155impl *Zora1155implFilterer) WatchContractRendererUpdated(opts *bind.WatchOpts, sink chan<- *Zora1155implContractRendererUpdated) (event.Subscription, error) {

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "ContractRendererUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implContractRendererUpdated)
				if err := _Zora1155impl.contract.UnpackLog(event, "ContractRendererUpdated", log); err != nil {
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

// ParseContractRendererUpdated is a log parse operation binding the contract event 0x56e810c8cae84731149f628981d25769a084570b9ba6eebf3c32879e3dce5609.
//
// Solidity: event ContractRendererUpdated(address renderer)
func (_Zora1155impl *Zora1155implFilterer) ParseContractRendererUpdated(log types.Log) (*Zora1155implContractRendererUpdated, error) {
	event := new(Zora1155implContractRendererUpdated)
	if err := _Zora1155impl.contract.UnpackLog(event, "ContractRendererUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Zora1155impl contract.
type Zora1155implInitializedIterator struct {
	Event *Zora1155implInitialized // Event containing the contract specifics and raw log

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
func (it *Zora1155implInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implInitialized)
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
		it.Event = new(Zora1155implInitialized)
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
func (it *Zora1155implInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implInitialized represents a Initialized event raised by the Zora1155impl contract.
type Zora1155implInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Zora1155impl *Zora1155implFilterer) FilterInitialized(opts *bind.FilterOpts) (*Zora1155implInitializedIterator, error) {

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Zora1155implInitializedIterator{contract: _Zora1155impl.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Zora1155impl *Zora1155implFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Zora1155implInitialized) (event.Subscription, error) {

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implInitialized)
				if err := _Zora1155impl.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Zora1155impl *Zora1155implFilterer) ParseInitialized(log types.Log) (*Zora1155implInitialized, error) {
	event := new(Zora1155implInitialized)
	if err := _Zora1155impl.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Zora1155impl contract.
type Zora1155implOwnershipTransferredIterator struct {
	Event *Zora1155implOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Zora1155implOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implOwnershipTransferred)
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
		it.Event = new(Zora1155implOwnershipTransferred)
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
func (it *Zora1155implOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implOwnershipTransferred represents a OwnershipTransferred event raised by the Zora1155impl contract.
type Zora1155implOwnershipTransferred struct {
	LastOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address lastOwner, address newOwner)
func (_Zora1155impl *Zora1155implFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*Zora1155implOwnershipTransferredIterator, error) {

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &Zora1155implOwnershipTransferredIterator{contract: _Zora1155impl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address lastOwner, address newOwner)
func (_Zora1155impl *Zora1155implFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Zora1155implOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implOwnershipTransferred)
				if err := _Zora1155impl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address lastOwner, address newOwner)
func (_Zora1155impl *Zora1155implFilterer) ParseOwnershipTransferred(log types.Log) (*Zora1155implOwnershipTransferred, error) {
	event := new(Zora1155implOwnershipTransferred)
	if err := _Zora1155impl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implPurchasedIterator is returned from FilterPurchased and is used to iterate over the raw logs and unpacked data for Purchased events raised by the Zora1155impl contract.
type Zora1155implPurchasedIterator struct {
	Event *Zora1155implPurchased // Event containing the contract specifics and raw log

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
func (it *Zora1155implPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implPurchased)
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
		it.Event = new(Zora1155implPurchased)
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
func (it *Zora1155implPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implPurchased represents a Purchased event raised by the Zora1155impl contract.
type Zora1155implPurchased struct {
	Sender   common.Address
	Minter   common.Address
	TokenId  *big.Int
	Quantity *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPurchased is a free log retrieval operation binding the contract event 0xb362243af1e2070d7d5bf8d713f2e0fab64203f1b71462afbe20572909788c5e.
//
// Solidity: event Purchased(address indexed sender, address indexed minter, uint256 indexed tokenId, uint256 quantity, uint256 value)
func (_Zora1155impl *Zora1155implFilterer) FilterPurchased(opts *bind.FilterOpts, sender []common.Address, minter []common.Address, tokenId []*big.Int) (*Zora1155implPurchasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "Purchased", senderRule, minterRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implPurchasedIterator{contract: _Zora1155impl.contract, event: "Purchased", logs: logs, sub: sub}, nil
}

// WatchPurchased is a free log subscription operation binding the contract event 0xb362243af1e2070d7d5bf8d713f2e0fab64203f1b71462afbe20572909788c5e.
//
// Solidity: event Purchased(address indexed sender, address indexed minter, uint256 indexed tokenId, uint256 quantity, uint256 value)
func (_Zora1155impl *Zora1155implFilterer) WatchPurchased(opts *bind.WatchOpts, sink chan<- *Zora1155implPurchased, sender []common.Address, minter []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "Purchased", senderRule, minterRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implPurchased)
				if err := _Zora1155impl.contract.UnpackLog(event, "Purchased", log); err != nil {
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

// ParsePurchased is a log parse operation binding the contract event 0xb362243af1e2070d7d5bf8d713f2e0fab64203f1b71462afbe20572909788c5e.
//
// Solidity: event Purchased(address indexed sender, address indexed minter, uint256 indexed tokenId, uint256 quantity, uint256 value)
func (_Zora1155impl *Zora1155implFilterer) ParsePurchased(log types.Log) (*Zora1155implPurchased, error) {
	event := new(Zora1155implPurchased)
	if err := _Zora1155impl.contract.UnpackLog(event, "Purchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implRendererUpdatedIterator is returned from FilterRendererUpdated and is used to iterate over the raw logs and unpacked data for RendererUpdated events raised by the Zora1155impl contract.
type Zora1155implRendererUpdatedIterator struct {
	Event *Zora1155implRendererUpdated // Event containing the contract specifics and raw log

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
func (it *Zora1155implRendererUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implRendererUpdated)
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
		it.Event = new(Zora1155implRendererUpdated)
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
func (it *Zora1155implRendererUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implRendererUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implRendererUpdated represents a RendererUpdated event raised by the Zora1155impl contract.
type Zora1155implRendererUpdated struct {
	TokenId  *big.Int
	Renderer common.Address
	User     common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRendererUpdated is a free log retrieval operation binding the contract event 0x5010f780a0de79bcfb9f3d6fec3cfe29758ef5c5800d575af709bc590bd78ade.
//
// Solidity: event RendererUpdated(uint256 indexed tokenId, address indexed renderer, address indexed user)
func (_Zora1155impl *Zora1155implFilterer) FilterRendererUpdated(opts *bind.FilterOpts, tokenId []*big.Int, renderer []common.Address, user []common.Address) (*Zora1155implRendererUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var rendererRule []interface{}
	for _, rendererItem := range renderer {
		rendererRule = append(rendererRule, rendererItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "RendererUpdated", tokenIdRule, rendererRule, userRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implRendererUpdatedIterator{contract: _Zora1155impl.contract, event: "RendererUpdated", logs: logs, sub: sub}, nil
}

// WatchRendererUpdated is a free log subscription operation binding the contract event 0x5010f780a0de79bcfb9f3d6fec3cfe29758ef5c5800d575af709bc590bd78ade.
//
// Solidity: event RendererUpdated(uint256 indexed tokenId, address indexed renderer, address indexed user)
func (_Zora1155impl *Zora1155implFilterer) WatchRendererUpdated(opts *bind.WatchOpts, sink chan<- *Zora1155implRendererUpdated, tokenId []*big.Int, renderer []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var rendererRule []interface{}
	for _, rendererItem := range renderer {
		rendererRule = append(rendererRule, rendererItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "RendererUpdated", tokenIdRule, rendererRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implRendererUpdated)
				if err := _Zora1155impl.contract.UnpackLog(event, "RendererUpdated", log); err != nil {
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

// ParseRendererUpdated is a log parse operation binding the contract event 0x5010f780a0de79bcfb9f3d6fec3cfe29758ef5c5800d575af709bc590bd78ade.
//
// Solidity: event RendererUpdated(uint256 indexed tokenId, address indexed renderer, address indexed user)
func (_Zora1155impl *Zora1155implFilterer) ParseRendererUpdated(log types.Log) (*Zora1155implRendererUpdated, error) {
	event := new(Zora1155implRendererUpdated)
	if err := _Zora1155impl.contract.UnpackLog(event, "RendererUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implSetupNewTokenIterator is returned from FilterSetupNewToken and is used to iterate over the raw logs and unpacked data for SetupNewToken events raised by the Zora1155impl contract.
type Zora1155implSetupNewTokenIterator struct {
	Event *Zora1155implSetupNewToken // Event containing the contract specifics and raw log

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
func (it *Zora1155implSetupNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implSetupNewToken)
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
		it.Event = new(Zora1155implSetupNewToken)
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
func (it *Zora1155implSetupNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implSetupNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implSetupNewToken represents a SetupNewToken event raised by the Zora1155impl contract.
type Zora1155implSetupNewToken struct {
	TokenId   *big.Int
	Sender    common.Address
	NewURI    string
	MaxSupply *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetupNewToken is a free log retrieval operation binding the contract event 0x1b944478023872bf91b25a13fdba3a686fdb1bf4dbb872f850240fad4b8cc068.
//
// Solidity: event SetupNewToken(uint256 indexed tokenId, address indexed sender, string newURI, uint256 maxSupply)
func (_Zora1155impl *Zora1155implFilterer) FilterSetupNewToken(opts *bind.FilterOpts, tokenId []*big.Int, sender []common.Address) (*Zora1155implSetupNewTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "SetupNewToken", tokenIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implSetupNewTokenIterator{contract: _Zora1155impl.contract, event: "SetupNewToken", logs: logs, sub: sub}, nil
}

// WatchSetupNewToken is a free log subscription operation binding the contract event 0x1b944478023872bf91b25a13fdba3a686fdb1bf4dbb872f850240fad4b8cc068.
//
// Solidity: event SetupNewToken(uint256 indexed tokenId, address indexed sender, string newURI, uint256 maxSupply)
func (_Zora1155impl *Zora1155implFilterer) WatchSetupNewToken(opts *bind.WatchOpts, sink chan<- *Zora1155implSetupNewToken, tokenId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "SetupNewToken", tokenIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implSetupNewToken)
				if err := _Zora1155impl.contract.UnpackLog(event, "SetupNewToken", log); err != nil {
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

// ParseSetupNewToken is a log parse operation binding the contract event 0x1b944478023872bf91b25a13fdba3a686fdb1bf4dbb872f850240fad4b8cc068.
//
// Solidity: event SetupNewToken(uint256 indexed tokenId, address indexed sender, string newURI, uint256 maxSupply)
func (_Zora1155impl *Zora1155implFilterer) ParseSetupNewToken(log types.Log) (*Zora1155implSetupNewToken, error) {
	event := new(Zora1155implSetupNewToken)
	if err := _Zora1155impl.contract.UnpackLog(event, "SetupNewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Zora1155impl contract.
type Zora1155implTransferBatchIterator struct {
	Event *Zora1155implTransferBatch // Event containing the contract specifics and raw log

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
func (it *Zora1155implTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implTransferBatch)
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
		it.Event = new(Zora1155implTransferBatch)
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
func (it *Zora1155implTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implTransferBatch represents a TransferBatch event raised by the Zora1155impl contract.
type Zora1155implTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Zora1155impl *Zora1155implFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Zora1155implTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implTransferBatchIterator{contract: _Zora1155impl.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Zora1155impl *Zora1155implFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Zora1155implTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implTransferBatch)
				if err := _Zora1155impl.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Zora1155impl *Zora1155implFilterer) ParseTransferBatch(log types.Log) (*Zora1155implTransferBatch, error) {
	event := new(Zora1155implTransferBatch)
	if err := _Zora1155impl.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Zora1155impl contract.
type Zora1155implTransferSingleIterator struct {
	Event *Zora1155implTransferSingle // Event containing the contract specifics and raw log

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
func (it *Zora1155implTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implTransferSingle)
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
		it.Event = new(Zora1155implTransferSingle)
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
func (it *Zora1155implTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implTransferSingle represents a TransferSingle event raised by the Zora1155impl contract.
type Zora1155implTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Zora1155impl *Zora1155implFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Zora1155implTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implTransferSingleIterator{contract: _Zora1155impl.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Zora1155impl *Zora1155implFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Zora1155implTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implTransferSingle)
				if err := _Zora1155impl.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Zora1155impl *Zora1155implFilterer) ParseTransferSingle(log types.Log) (*Zora1155implTransferSingle, error) {
	event := new(Zora1155implTransferSingle)
	if err := _Zora1155impl.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Zora1155impl contract.
type Zora1155implURIIterator struct {
	Event *Zora1155implURI // Event containing the contract specifics and raw log

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
func (it *Zora1155implURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implURI)
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
		it.Event = new(Zora1155implURI)
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
func (it *Zora1155implURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implURI represents a URI event raised by the Zora1155impl contract.
type Zora1155implURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Zora1155impl *Zora1155implFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Zora1155implURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implURIIterator{contract: _Zora1155impl.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Zora1155impl *Zora1155implFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Zora1155implURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implURI)
				if err := _Zora1155impl.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Zora1155impl *Zora1155implFilterer) ParseURI(log types.Log) (*Zora1155implURI, error) {
	event := new(Zora1155implURI)
	if err := _Zora1155impl.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implUpdatedPermissionsIterator is returned from FilterUpdatedPermissions and is used to iterate over the raw logs and unpacked data for UpdatedPermissions events raised by the Zora1155impl contract.
type Zora1155implUpdatedPermissionsIterator struct {
	Event *Zora1155implUpdatedPermissions // Event containing the contract specifics and raw log

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
func (it *Zora1155implUpdatedPermissionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implUpdatedPermissions)
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
		it.Event = new(Zora1155implUpdatedPermissions)
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
func (it *Zora1155implUpdatedPermissionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implUpdatedPermissionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implUpdatedPermissions represents a UpdatedPermissions event raised by the Zora1155impl contract.
type Zora1155implUpdatedPermissions struct {
	TokenId     *big.Int
	User        common.Address
	Permissions *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPermissions is a free log retrieval operation binding the contract event 0x35fb03d0d293ef5b362761900725ce891f8f766b5a662cdd445372355448e7ca.
//
// Solidity: event UpdatedPermissions(uint256 indexed tokenId, address indexed user, uint256 indexed permissions)
func (_Zora1155impl *Zora1155implFilterer) FilterUpdatedPermissions(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address, permissions []*big.Int) (*Zora1155implUpdatedPermissionsIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var permissionsRule []interface{}
	for _, permissionsItem := range permissions {
		permissionsRule = append(permissionsRule, permissionsItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "UpdatedPermissions", tokenIdRule, userRule, permissionsRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implUpdatedPermissionsIterator{contract: _Zora1155impl.contract, event: "UpdatedPermissions", logs: logs, sub: sub}, nil
}

// WatchUpdatedPermissions is a free log subscription operation binding the contract event 0x35fb03d0d293ef5b362761900725ce891f8f766b5a662cdd445372355448e7ca.
//
// Solidity: event UpdatedPermissions(uint256 indexed tokenId, address indexed user, uint256 indexed permissions)
func (_Zora1155impl *Zora1155implFilterer) WatchUpdatedPermissions(opts *bind.WatchOpts, sink chan<- *Zora1155implUpdatedPermissions, tokenId []*big.Int, user []common.Address, permissions []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var permissionsRule []interface{}
	for _, permissionsItem := range permissions {
		permissionsRule = append(permissionsRule, permissionsItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "UpdatedPermissions", tokenIdRule, userRule, permissionsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implUpdatedPermissions)
				if err := _Zora1155impl.contract.UnpackLog(event, "UpdatedPermissions", log); err != nil {
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

// ParseUpdatedPermissions is a log parse operation binding the contract event 0x35fb03d0d293ef5b362761900725ce891f8f766b5a662cdd445372355448e7ca.
//
// Solidity: event UpdatedPermissions(uint256 indexed tokenId, address indexed user, uint256 indexed permissions)
func (_Zora1155impl *Zora1155implFilterer) ParseUpdatedPermissions(log types.Log) (*Zora1155implUpdatedPermissions, error) {
	event := new(Zora1155implUpdatedPermissions)
	if err := _Zora1155impl.contract.UnpackLog(event, "UpdatedPermissions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implUpdatedRoyaltiesIterator is returned from FilterUpdatedRoyalties and is used to iterate over the raw logs and unpacked data for UpdatedRoyalties events raised by the Zora1155impl contract.
type Zora1155implUpdatedRoyaltiesIterator struct {
	Event *Zora1155implUpdatedRoyalties // Event containing the contract specifics and raw log

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
func (it *Zora1155implUpdatedRoyaltiesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implUpdatedRoyalties)
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
		it.Event = new(Zora1155implUpdatedRoyalties)
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
func (it *Zora1155implUpdatedRoyaltiesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implUpdatedRoyaltiesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implUpdatedRoyalties represents a UpdatedRoyalties event raised by the Zora1155impl contract.
type Zora1155implUpdatedRoyalties struct {
	TokenId       *big.Int
	User          common.Address
	Configuration ICreatorRoyaltiesControlRoyaltyConfiguration
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRoyalties is a free log retrieval operation binding the contract event 0x5837d55897cfc337f160a71d7b63a047abd50a3a8834f1c5d70f338846358c6d.
//
// Solidity: event UpdatedRoyalties(uint256 indexed tokenId, address indexed user, (uint32,uint32,address) configuration)
func (_Zora1155impl *Zora1155implFilterer) FilterUpdatedRoyalties(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address) (*Zora1155implUpdatedRoyaltiesIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "UpdatedRoyalties", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implUpdatedRoyaltiesIterator{contract: _Zora1155impl.contract, event: "UpdatedRoyalties", logs: logs, sub: sub}, nil
}

// WatchUpdatedRoyalties is a free log subscription operation binding the contract event 0x5837d55897cfc337f160a71d7b63a047abd50a3a8834f1c5d70f338846358c6d.
//
// Solidity: event UpdatedRoyalties(uint256 indexed tokenId, address indexed user, (uint32,uint32,address) configuration)
func (_Zora1155impl *Zora1155implFilterer) WatchUpdatedRoyalties(opts *bind.WatchOpts, sink chan<- *Zora1155implUpdatedRoyalties, tokenId []*big.Int, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "UpdatedRoyalties", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implUpdatedRoyalties)
				if err := _Zora1155impl.contract.UnpackLog(event, "UpdatedRoyalties", log); err != nil {
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

// ParseUpdatedRoyalties is a log parse operation binding the contract event 0x5837d55897cfc337f160a71d7b63a047abd50a3a8834f1c5d70f338846358c6d.
//
// Solidity: event UpdatedRoyalties(uint256 indexed tokenId, address indexed user, (uint32,uint32,address) configuration)
func (_Zora1155impl *Zora1155implFilterer) ParseUpdatedRoyalties(log types.Log) (*Zora1155implUpdatedRoyalties, error) {
	event := new(Zora1155implUpdatedRoyalties)
	if err := _Zora1155impl.contract.UnpackLog(event, "UpdatedRoyalties", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implUpdatedTokenIterator is returned from FilterUpdatedToken and is used to iterate over the raw logs and unpacked data for UpdatedToken events raised by the Zora1155impl contract.
type Zora1155implUpdatedTokenIterator struct {
	Event *Zora1155implUpdatedToken // Event containing the contract specifics and raw log

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
func (it *Zora1155implUpdatedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implUpdatedToken)
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
		it.Event = new(Zora1155implUpdatedToken)
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
func (it *Zora1155implUpdatedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implUpdatedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implUpdatedToken represents a UpdatedToken event raised by the Zora1155impl contract.
type Zora1155implUpdatedToken struct {
	From      common.Address
	TokenId   *big.Int
	TokenData IZoraCreator1155TypesV1TokenData
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedToken is a free log retrieval operation binding the contract event 0x5086d1bcea28999da9875111e3592688fbfa821db63214c695ca35768080c2fe.
//
// Solidity: event UpdatedToken(address indexed from, uint256 indexed tokenId, (string,uint256,uint256) tokenData)
func (_Zora1155impl *Zora1155implFilterer) FilterUpdatedToken(opts *bind.FilterOpts, from []common.Address, tokenId []*big.Int) (*Zora1155implUpdatedTokenIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "UpdatedToken", fromRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implUpdatedTokenIterator{contract: _Zora1155impl.contract, event: "UpdatedToken", logs: logs, sub: sub}, nil
}

// WatchUpdatedToken is a free log subscription operation binding the contract event 0x5086d1bcea28999da9875111e3592688fbfa821db63214c695ca35768080c2fe.
//
// Solidity: event UpdatedToken(address indexed from, uint256 indexed tokenId, (string,uint256,uint256) tokenData)
func (_Zora1155impl *Zora1155implFilterer) WatchUpdatedToken(opts *bind.WatchOpts, sink chan<- *Zora1155implUpdatedToken, from []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "UpdatedToken", fromRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implUpdatedToken)
				if err := _Zora1155impl.contract.UnpackLog(event, "UpdatedToken", log); err != nil {
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

// ParseUpdatedToken is a log parse operation binding the contract event 0x5086d1bcea28999da9875111e3592688fbfa821db63214c695ca35768080c2fe.
//
// Solidity: event UpdatedToken(address indexed from, uint256 indexed tokenId, (string,uint256,uint256) tokenData)
func (_Zora1155impl *Zora1155implFilterer) ParseUpdatedToken(log types.Log) (*Zora1155implUpdatedToken, error) {
	event := new(Zora1155implUpdatedToken)
	if err := _Zora1155impl.contract.UnpackLog(event, "UpdatedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Zora1155implUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Zora1155impl contract.
type Zora1155implUpgradedIterator struct {
	Event *Zora1155implUpgraded // Event containing the contract specifics and raw log

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
func (it *Zora1155implUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Zora1155implUpgraded)
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
		it.Event = new(Zora1155implUpgraded)
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
func (it *Zora1155implUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Zora1155implUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Zora1155implUpgraded represents a Upgraded event raised by the Zora1155impl contract.
type Zora1155implUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Zora1155impl *Zora1155implFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Zora1155implUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Zora1155impl.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Zora1155implUpgradedIterator{contract: _Zora1155impl.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Zora1155impl *Zora1155implFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Zora1155implUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Zora1155impl.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Zora1155implUpgraded)
				if err := _Zora1155impl.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Zora1155impl *Zora1155implFilterer) ParseUpgraded(log types.Log) (*Zora1155implUpgraded, error) {
	event := new(Zora1155implUpgraded)
	if err := _Zora1155impl.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
