<img align="left" src="./.github/resources/goomy.png" width="75">
<h1>Goomy the Blob Tool</h1>

goomy-blob is a simple tool that can be used to generate random blob transactions for ethereum testnets.

Goomy can be used for stress testing (flooding the network with thousands of blobs) or to have a continuous amount of blobs (2-4 blobs per block) over long time for testing purposes.

Goomy provides two commands:
* `blob-sender`: Simple utility to send a single blob transaction with specified parameters.
* `blob-spammer`: Tool for mass blob spamming

## Build

You can use this tool via pre-build docker images: [ethpandaops/goomy-blob](https://hub.docker.com/r/ethpandaops/goomy-blob)

Or build it yourself:

```
git clone https://github.com/ethpandaops/goomy-blob.git
cd goomy-blob
go build ./cmd/blob-spammer
go build ./cmd/blob-sender  # if needed
```



## Usage

### `blob-spammer`
`blob-spammer` is a tool for sending mass blob transactions.

```
Usage of blob-spammer:
Required:
  -p, --privkey string        The private key of the wallet to send funds from.
  
  -h, --rpchost string        The RPC host to send transactions to (multiple allowed).
      --rpchost-file string   File with a list of RPC hosts to send transactions to.
      
Optional:
  -s, --seed string           The child wallet seed.
  -v, --verbose               Run the tool with verbose output.
```

The tool provides multiple scenarios, that focus on different aspects of blob transactions. One of the scenarios must be selected to run the tool:

#### `blob-spammer combined`

For general testing, there is a `combined` scenario, which combines parts of all other scenarios in a randomized way.

```
Usage of blob-spammer combined:
Required (at least one of):
  -c, --count uint            Total number of blob transactions to send
  -t, --throughput uint       Number of blob transactions to send per slot
  
Optional:
      --basefee uint          Max fee per gas to use in blob transactions (in gwei) (default 20)
      --blobfee uint          Max blob fee to use in blob transactions (in gwei) (default 20)
      --max-pending uint      Maximum number of pending transactions
      --max-replace uint      Maximum number of replacement transactions (default 4)
      --max-wallets uint      Maximum number of child wallets to use
      --rebroadcast uint      Number of seconds to wait before re-broadcasting a transaction (default 30)
      --replace uint          Number of seconds to wait before replace a transaction (default 30)
  -b, --sidecars uint         Maximum number of blob sidecars per blob transactions (default 3)
      --tipfee uint           Max tip per gas to use in blob transactions (in gwei) (default 2)
```

##### Examples:

Continuous random blob spamming (~2-4 sidecars per block):
```
blob-spammer combined -p "<PRIVKEY>" -h http://rpc-host1:8545 -b 2 -t 3 --max-pending 3
```

flood the network with 1000 blobs (+some replacements) via 2 rpc hosts:
```
blob-spammer combined -p "<PRIVKEY>" -h http://rpc-host1:8545 -h http://rpc-host2:8545 -c 1000
```

#### `blob-spammer normal`

The `normal` scenario sends out normal blobs with random data only.\
No replacement or cancellation transactions are being send.

#### `blob-spammer replacements`

The `replacements` scenario sends out blobs and always tries to replace these blobs with replacement blob transactions a few seconds later, further replacement transactions are being sent until inclusion in a block.


#### `blob-spammer wallets`

The `wallets` scenario prepares & prints the list of child wallets that are used to send blob transactions from.\
It's more intended for debugging. The tool takes care of these wallets internally, so there is nothing to do with them ;)

### `blob-sender`

`blob-sender` is a simple utility to send a single blob transaction with specified parameters.

```
Usage of blob-sender:
      --addnonce int           Nonce offset to use for transactions (useful for replacement transactions)
  -b, --blobs stringArray      The blobs to reference in the transaction (in hex format or special placeholders).
      --chainid uint           ChainID of the network (For offline mode in combination with --output or to override transactions)
  -n, --count uint             The number of transactions to send. (default 1)
  -d, --data string            The transaction calldata.
      --gaslimit uint          The gas limit for transactions. (default 500000)
      --maxblobfee float32     The maximum blob fee per chunk in gwei. (default 10)
      --maxfeepergas float32   The gas limit for transactions. (default 20)
      --maxpriofee float32     The maximum priority fee per gas in gwei. (default 1.2)
      --nonce uint             Current nonce of the wallet (For offline mode in combination with --output)
  -o, --output                 Output signed transactions to stdout instead of broadcasting them (offline mode).
  -p, --privkey string         The private key of the wallet to send funds from.
                               (Special: "env" to read from FUNDINGTOOL_PRIVKEY environment variable)
      --random-privkey         Use random private key if no privkey supplied
  -r, --rpchost string         The RPC host to send transactions to. (default "http://127.0.0.1:8545")
  -t, --to string              The transaction to address.
  -a, --value uint             The transaction value.
  -v, --verbose                Run the script with verbose output
```
