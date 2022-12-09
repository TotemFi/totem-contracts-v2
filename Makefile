CURRENTDIR = $(shell pwd)
ABIS="${CURRENTDIR}/pkg/abi"
CONTRACTS="${CURRENTDIR}/pkg/contracts"
.PHONY: all
mainnet:
	yarn clean
	yarn prepare
	@for f in $(shell ls ${ABIS}); do \
		j=`echo $$f | cut -d . -f 1`; \
		echo "$$j"; \
		abigen --abi ${CURRENTDIR}/pkg/abi/$${f} --pkg=contracts --type=$${j} --out=${CURRENTDIR}/pkg/contracts/$${j}.go --alias getStakingReward=_getStakingRewardA,getIndexedStakingReward=getIndexedStakingRewardA; \
	done