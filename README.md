# Simple CosmosSDK Token

This is a simple implementation of a token using CosmosSDK. It uses the bank and supply module from CosmosSDK and re-exported the balances, total supply and 1 to 1 transfer features to the `calvincoin` application.

## Pre-requisite

Golang

## Setup

### Build
```
make install
```

### Reset and Init chain
```bash
./unsafe-reset-all.sh
```

p.s. Alice and Bob's address has been created for you in `./appcli`.

### Start chain
```bash
appd start
```

## Interact with chain

To start with, source the environment variables of Alice and Bob addresses:

#### Bash
```bash
. ./env.sh
```

#### Fish
```bash
./env_fish.sh
```

### Query balances
```bash
appcli query calvincoin balances $ALICE
appcli query calvincoin balances $BOB
```

### Query total supply
```bash
appcli query calvincoin totalsupply
```

### Send transfer transaction
```bash
appcli tx calvincoin transfer $BOB 1calvincoin --from alice
```