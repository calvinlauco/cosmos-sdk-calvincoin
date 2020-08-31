#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

rm -rf ./.appd
appd init node1 --chain-id calvinchain

appd add-genesis-account cosmos136wu3xmj7a5lz5699fcyjatr9k5puqd9g05kdt 1000calvincoin,100000000stake
appd add-genesis-account cosmos1zse0pwwmhj05y0u8c9hpc4m3jj7tyh8rcjchas 1000calvincoin,100000000stake

appd gentx --name alice --keyring-backend test

appd collect-gentxs
appd validate-genesis
