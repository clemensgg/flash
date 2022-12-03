 ![CryptoCrew Validators](ccvalidators_logo.png)
# Flash

**flash** is a blockchain built using Cosmos SDK and Tendermint.  

It uses a forked version of the Cosmos SDK (`Flash SDK`) with updated default consensus timeout parameters, to allow a target block production rate of ~1.7s / block

## Get started

Install [go](https://go.dev/dl/)

## Build and install to go bin path

```
make install
```

## Initialize config

Come up with a moniker for your node, then run:

```
flashd init $MONIKER
```
 
## Launch with genesis file or run as standalone chain

To launch as a consumer chain, download and save shared genesis file to `~/.flash/config/genesis.json`. Additionally add peering information (`persistent_peers` or `seeds`) to `~/.flash/config/config.toml`

To instead launch as a standalone, single node chain, run:

```
flashd add-consumer-section
```

## Launch node

```
flashd start
```
