package sdk

import (
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
)

type PublicKey = eddsa.PublicKey

func ConvertTransferNftTxInfo(tx *TransferNftTxInfo) *legendTxTypes.TransferNftTxInfo {
	return &legendTxTypes.TransferNftTxInfo{
		FromAccountIndex:  tx.FromAccountIndex,
		ToAccountIndex:    tx.ToAccountIndex,
		ToAccountNameHash: tx.ToAccountNameHash,
		NftIndex:          tx.NftIndex,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		CallData:          tx.CallData,
		CallDataHash:      tx.CallDataHash,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func ConvertWithdrawNftTxInfo(tx *WithdrawNftTxInfo) *legendTxTypes.WithdrawNftTxInfo {
	return &legendTxTypes.WithdrawNftTxInfo{
		AccountIndex:           tx.AccountIndex,
		CreatorAccountIndex:    tx.CreatorAccountIndex,
		CreatorAccountNameHash: tx.CreatorAccountNameHash,
		CreatorTreasuryRate:    tx.CreatorTreasuryRate,
		NftIndex:               tx.NftIndex,
		NftContentHash:         tx.NftContentHash,
		NftL1Address:           tx.NftL1Address,
		NftL1TokenId:           tx.NftL1TokenId,
		CollectionId:           tx.CollectionId,
		ToAddress:              tx.ToAddress,
		GasAccountIndex:        tx.GasAccountIndex,
		GasFeeAssetId:          tx.GasFeeAssetId,
		GasFeeAssetAmount:      tx.GasFeeAssetAmount,
		ExpiredAt:              tx.ExpiredAt,
		Nonce:                  tx.Nonce,
		Sig:                    tx.Sig,
	}
}

func ConvertOfferTxInfo(tx *OfferTxInfo) *legendTxTypes.OfferTxInfo {
	return &legendTxTypes.OfferTxInfo{
		Type:         tx.Type,
		OfferId:      tx.OfferId,
		AccountIndex: tx.AccountIndex,
		NftIndex:     tx.NftIndex,
		AssetId:      tx.AssetId,
		AssetAmount:  tx.AssetAmount,
		ListedAt:     tx.ListedAt,
		ExpiredAt:    tx.ExpiredAt,
		TreasuryRate: tx.TreasuryRate,
		Sig:          tx.Sig,
	}
}

func ConvertMintNftTxInfo(tx *MintNftTxInfo) *legendTxTypes.MintNftTxInfo {
	return &legendTxTypes.MintNftTxInfo{
		CreatorAccountIndex: tx.CreatorAccountIndex,
		ToAccountIndex:      tx.ToAccountIndex,
		ToAccountNameHash:   tx.ToAccountNameHash,
		NftIndex:            tx.NftIndex,
		NftContentHash:      tx.NftContentHash,
		NftCollectionId:     tx.NftCollectionId,
		CreatorTreasuryRate: tx.CreatorTreasuryRate,
		GasAccountIndex:     tx.GasAccountIndex,
		GasFeeAssetId:       tx.GasFeeAssetId,
		GasFeeAssetAmount:   tx.GasFeeAssetAmount,
		ExpiredAt:           tx.ExpiredAt,
		Nonce:               tx.Nonce,
		Sig:                 tx.Sig,
	}
}

func ConvertCreateCollectionTxInfo(tx *CreateCollectionTxInfo) *legendTxTypes.CreateCollectionTxInfo {
	return &legendTxTypes.CreateCollectionTxInfo{
		AccountIndex:      tx.AccountIndex,
		CollectionId:      tx.CollectionId,
		Name:              tx.Name,
		Introduction:      tx.Introduction,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}

func ConvertAtomicMatchTxInfo(tx *AtomicMatchTxInfo) *legendTxTypes.AtomicMatchTxInfo {
	return &legendTxTypes.AtomicMatchTxInfo{
		AccountIndex: tx.AccountIndex,
		BuyOffer: &legendTxTypes.OfferTxInfo{
			Type:         tx.BuyOffer.Type,
			OfferId:      tx.BuyOffer.OfferId,
			AccountIndex: tx.BuyOffer.AccountIndex,
			NftIndex:     tx.BuyOffer.NftIndex,
			AssetId:      tx.BuyOffer.AssetId,
			AssetAmount:  tx.BuyOffer.AssetAmount,
			ListedAt:     tx.BuyOffer.ListedAt,
			ExpiredAt:    tx.BuyOffer.ExpiredAt,
			TreasuryRate: tx.BuyOffer.TreasuryRate,
			Sig:          tx.BuyOffer.Sig,
		},
		SellOffer: &legendTxTypes.OfferTxInfo{
			Type:         tx.SellOffer.Type,
			OfferId:      tx.SellOffer.OfferId,
			AccountIndex: tx.SellOffer.AccountIndex,
			NftIndex:     tx.SellOffer.NftIndex,
			AssetId:      tx.SellOffer.AssetId,
			AssetAmount:  tx.SellOffer.AssetAmount,
			ListedAt:     tx.SellOffer.ListedAt,
			ExpiredAt:    tx.SellOffer.ExpiredAt,
			TreasuryRate: tx.SellOffer.TreasuryRate,
			Sig:          tx.SellOffer.Sig,
		},
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		CreatorAmount:     tx.CreatorAmount,
		TreasuryAmount:    tx.TreasuryAmount,
		Nonce:             tx.Nonce,
		ExpiredAt:         tx.ExpiredAt,
		Sig:               tx.Sig,
	}
}
