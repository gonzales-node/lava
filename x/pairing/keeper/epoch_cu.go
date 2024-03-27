package keeper

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/x/pairing/types"
)

// UniqueEpochSession is used to detect double spend attacks
// It's kept in a epoch-prefixed store with a unique index: provider, project ID, chain ID and session ID
//
// ProviderEpochCu is used to track the CU of a specific provider in a specific epoch for unresponsiveness
// It's kept in a epoch-prefixed store with a unique index: provider address
//
// ProviderConsumerEpochCu is used to track the CU between a specific provider and
// consumer in a specific epoch for payments
// It's kept in a epoch-prefixed store with a unique index: provider and project ID

/* ########## UniqueEpochSession ############ */

// SetUniqueEpochSession sets a UniqueEpochSession in the store
func (k Keeper) SetUniqueEpochSession(ctx sdk.Context, epoch uint64, provider string, project string, chainID string, sessionID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UniqueEpochSessionKeyPrefix(epoch))
	store.Set(types.UniqueEpochSessionKey(provider, chainID, project, sessionID), []byte{0})
}

// IsUniqueEpochSessionExists checks if a UniqueEpochSession exists
func (k Keeper) IsUniqueEpochSessionExists(ctx sdk.Context, epoch uint64, provider string, project string, chainID string, sessionID uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UniqueEpochSessionKeyPrefix(epoch))
	b := store.Get(types.UniqueEpochSessionKey(provider, chainID, project, sessionID))
	return b != nil
}

// RemoveAllUniqueEpochSession removes all the UniqueEpochSession objects from the store for a specific epoch
func (k Keeper) RemoveAllUniqueEpochSession(ctx sdk.Context, epoch uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UniqueEpochSessionKeyPrefix(epoch))
	iterator := sdk.KVStorePrefixIterator(store, []byte{}) // Get an iterator with no prefix to iterate over all keys
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
}

// GetAllUniqueEpochSessionForEpoch gets all the UniqueEpochSession objects from the store for a specific epoch
func (k Keeper) GetAllUniqueEpochSessionForEpoch(ctx sdk.Context, epoch uint64) []string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UniqueEpochSessionKeyPrefix(epoch))

	iterator := sdk.KVStorePrefixIterator(store, []byte{}) // Get an iterator with no prefix to iterate over all keys
	defer iterator.Close()

	var keys []string
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		// Remove the prefix to get the actual UniqueEpochSession key
		uniqueEpochSessionKey := strings.TrimPrefix(string(key), string(types.UniqueEpochSessionKeyPrefix(epoch)))
		keys = append(keys, uniqueEpochSessionKey)
	}

	return keys
}

// GetAllUniqueEpochSessionStore gets all the UniqueEpochSession objects from the store (used for genesis)
func (k Keeper) GetAllUniqueEpochSessionStore(ctx sdk.Context) []types.UniqueEpochSessionGenesis {
	info := []types.UniqueEpochSessionGenesis{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UniqueEpochSessionPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{}) // Get an iterator with no prefix to iterate over all keys
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		split := strings.Split(key, "/") // key structure: epoch/unique_epoch_session_key
		epoch, err := strconv.ParseUint(split[0], 10, 64)
		if err != nil {
			utils.LavaFormatError("could not decode UniqueEpochSessionKey with epoch", err, utils.LogAttr("key", string(iterator.Key())))
			continue
		}
		provider, chainID, project, sessionID, err := types.DecodeUniqueEpochSessionKey(split[1])
		if err != nil {
			utils.LavaFormatError("could not decode UniqueEpochSessionKey", err, utils.LogAttr("key", split[1]))
		}
		info = append(info, types.UniqueEpochSessionGenesis{
			Epoch:     epoch,
			Provider:  provider,
			Project:   project,
			ChainId:   chainID,
			SessionId: sessionID,
		})
	}

	return info
}

/* ########## ProviderEpochCu ############ */

// SetProviderEpochCu sets a ProviderEpochCu in the store
func (k Keeper) SetProviderEpochCu(ctx sdk.Context, epoch uint64, provider string, chainID string, providerEpochCu types.ProviderEpochCu) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProviderEpochCuKeyPrefix(epoch))
	b := k.cdc.MustMarshal(&providerEpochCu)
	store.Set(types.ProviderEpochCuKey(provider, chainID), b)
}

// GetProviderEpochCu returns a ProviderEpochCu for a specific epoch and provider
func (k Keeper) GetProviderEpochCu(ctx sdk.Context, epoch uint64, provider string, chainID string) (val types.ProviderEpochCu, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProviderEpochCuKeyPrefix(epoch))
	b := store.Get(types.ProviderEpochCuKey(provider, chainID))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProviderEpochCu removes a ProviderEpochCu from the store
func (k Keeper) RemoveAllProviderEpochCu(ctx sdk.Context, epoch uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProviderEpochCuKeyPrefix(epoch))
	iterator := sdk.KVStorePrefixIterator(store, []byte{}) // Get an iterator with no prefix to iterate over all keys
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
}

// GetAllProviderEpochCuStore returns all the ProviderEpochCu from the store (used for genesis)
func (k Keeper) GetAllProviderEpochCuStore(ctx sdk.Context) []types.ProviderEpochCuGenesis {
	info := []types.ProviderEpochCuGenesis{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ProviderEpochCuPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{}) // Get an iterator with no prefix to iterate over all keys
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		split := strings.Split(string(iterator.Key()), "/") // key structure: epoch/provider_epoch_cu_key
		epoch, err := strconv.ParseUint(split[0], 10, 64)
		if err != nil {
			utils.LavaFormatError("could not decode ProviderEpochCuKey with epoch", err, utils.LogAttr("key", string(iterator.Key())))
			continue
		}
		provider, chainID, err := types.DecodeProviderEpochCuKey(split[1])
		if err != nil {
			utils.LavaFormatError("could not decode ProviderEpochCuKey with epoch", err, utils.LogAttr("key", split[1]))
			continue
		}
		var pec types.ProviderEpochCu
		k.cdc.MustUnmarshal(iterator.Value(), &pec)
		info = append(info, types.ProviderEpochCuGenesis{
			Epoch:           epoch,
			Provider:        provider,
			ChainId:         chainID,
			ProviderEpochCu: pec,
		})
	}

	return info
}

/* ########## ProviderConsumerEpochCu ############ */

// SetProviderConsumerEpochCu sets a ProviderConsumerEpochCu in the store
func (k Keeper) SetProviderConsumerEpochCu(ctx sdk.Context, epoch uint64, provider string, project string, chainID string, providerConsumerEpochCu types.ProviderConsumerEpochCu) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProviderConsumerEpochCuKeyPrefix(epoch))
	b := k.cdc.MustMarshal(&providerConsumerEpochCu)
	store.Set(types.ProviderConsumerEpochCuKey(provider, project, chainID), b)
}

// GetProviderConsumerEpochCu returns a ProviderConsumerEpochCu for a specific epoch, provider and project
func (k Keeper) GetProviderConsumerEpochCu(ctx sdk.Context, epoch uint64, provider string, project string, chainID string) (val types.ProviderConsumerEpochCu, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProviderConsumerEpochCuKeyPrefix(epoch))
	b := store.Get(types.ProviderConsumerEpochCuKey(provider, project, chainID))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProviderConsumerEpochCu removes a ProviderConsumerEpochCu from the store
func (k Keeper) RemoveProviderConsumerEpochCu(ctx sdk.Context, epoch uint64, provider string, project string, chainID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProviderConsumerEpochCuKeyPrefix(epoch))
	store.Delete(types.ProviderConsumerEpochCuKey(provider, project, chainID))
}

// GetAllProviderConsumerEpochCuStore returns all the ProviderConsumerEpochCu from the store (used for genesis)
func (k Keeper) GetAllProviderConsumerEpochCu(ctx sdk.Context, epoch uint64) (keys []string, providerConsumerEpochCus []types.ProviderConsumerEpochCu) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProviderConsumerEpochCuKeyPrefix(epoch))
	iterator := sdk.KVStorePrefixIterator(store, []byte{}) // Get an iterator with no prefix to iterate over all keys
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, string(iterator.Key()))
		var pcec types.ProviderConsumerEpochCu
		k.cdc.MustUnmarshal(iterator.Value(), &pcec)
		providerConsumerEpochCus = append(providerConsumerEpochCus, pcec)
	}

	return keys, providerConsumerEpochCus
}

// GetAllProviderConsumerEpochCuStore returns all the ProviderConsumerEpochCu from the store (used for genesis)
func (k Keeper) GetAllProviderConsumerEpochCuStore(ctx sdk.Context) []types.ProviderConsumerEpochCuGenesis {
	info := []types.ProviderConsumerEpochCuGenesis{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ProviderConsumerEpochCuPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{}) // Get an iterator with no prefix to iterate over all keys
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		split := strings.Split(string(iterator.Key()), "/") // key structure: epoch/provider_consumer_epoch_key
		epoch, err := strconv.ParseUint(split[0], 10, 64)
		if err != nil {
			utils.LavaFormatError("could not decode ProviderConsumerEpochCuKey with epoch", err, utils.LogAttr("key", string(iterator.Key())))
			continue
		}
		provider, project, chainID, err := types.DecodeProviderConsumerEpochCuKey(split[1])
		if err != nil {
			utils.LavaFormatError("could not decode ProviderConsumerEpochCuKey", err, utils.LogAttr("key", split[1]))
			continue
		}
		var pcec types.ProviderConsumerEpochCu
		k.cdc.MustUnmarshal(iterator.Value(), &pcec)
		info = append(info, types.ProviderConsumerEpochCuGenesis{
			Epoch:                   epoch,
			Provider:                provider,
			Project:                 project,
			ChainId:                 chainID,
			ProviderConsumerEpochCu: pcec,
		})
	}

	return info
}