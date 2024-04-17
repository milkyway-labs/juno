## Configuration
The default `config.yaml` file should look like the following:

<details>

<summary>Default config.yaml file</summary>

```yaml
chain:
  prefix: "cosmos"
  modules:

node:
  type: remote
  config:
    rpc:
      client_name: juno
      address: http://localhost:26657
      max_connections: 20
    grpc:
      address: localhost:9090

parsing:
  workers: 1
  listen_new_blocks: true
  parse_old_blocks: true
  start_height: 1
  parse_genesis: true
  fast_sync: true
  re_enqueue_when_failed: false
  max_retries: 10

database:
  url: postgres://user:password@localhost:5432/juno?sslmode=disable
  partition_size: 100000

logging:
  format: "text"
  level: "debug"
```

</details>

Let's see what each section refers to:

- [`chain`](#chain)
- [`node`](#node)
- [`parsing`](#parsing)
- [`database`](#database)
- [`pruning`](#pruning)
- [`logging`](#logging)
- [`telemetry`](#telemetry)

## `chain`
This section contains the details of the chain configuration regarding the Cosmos SDK.

| Attribute |   Type   | Description                            | Example                              |
|:---------:|:--------:|:---------------------------------------|:-------------------------------------|
| `modules` | `array`  | List of modules that should be enabled | `[ "auth", "bank", "distribution" ]` |
| `prefix`  | `string` | Bech 32 prefix of the addresses        | `cosmos`                             | 

### Supported modules
Currently, we support the following modules:

- `pruning` to periodically prune the old database data
- `telemetry` to support a telemetry service

## `node`
This section contains the details of the node to which Juno will connect.

| Attribute |   Type   | Description                                                  | Example  |
|:---------:|:--------:|:-------------------------------------------------------------|:---------|
|  `type`   | `string` | Tells which type of node to use (either `local` or `remote`) | `remote` |
| `config`  | `object` | Contains the configuration data for the node                 |          | 

### Remote node
A remote node is the default implementation of a node. It relies on both an RPC and gRPC connections to get the data. If
you want to use this kind of node, you need to set the [`node`](#node) type to `remote` and then set the following
attributes of the configuration.

| Attribute |   Type   | Description                                                                                                       | Example |
|:---------:|:--------:|:------------------------------------------------------------------------------------------------------------------|:--------|
|   `rpc`   | `object` | Contains the RPC configuration data                                                                               |         | 
|  `grpc`   | `object` | (Optional) Contains the gRPC configuration data. If not provided, a gRPC-over-RPC connection will be used instead |         | 

#### `rpc`

|     Attribute     |   Type   | Description                                                                                                                           | Example                  |
|:-----------------:|:--------:|:--------------------------------------------------------------------------------------------------------------------------------------|:-------------------------|
|     `address`     | `string` | Address of the RPC endpoint                                                                                                           | `http://localhost:26657` |
|   `client_name`   | `string` | Client name used when subscribing to the Tendermint websocket                                                                         | `juno`                   |
| `max_connections` |  `int`   | Max number of connections that can created towards the RPC node (any value less or equal to `0` means to use the default one instead) | `20`                     | 

#### `grpc`

| Attribute |   Type   | Description                  | Example          |
|:---------:|:--------:|:-----------------------------|:-----------------|
| `address` | `string` | Address of the gRPC endpoint | `localhost:9090` |

### Local node
A local node reads the data to be parsed from a local directory referred to as `home`. If you want to use this kind of
node, you need to set the [`node`](#node) type to `local` and then set the following attributes of the configuration.

| Attribute |   Type   | Description                         | Example             |
|:---------:|:--------:|:------------------------------------|:--------------------|
|  `home`   | `string` | Path to the home folder of the node | `/home/user/.gaiad` |

## `parsing`

|        Attribute         |   Type    | Description                                                                                                                                                                                                                                                               | Example                                  |
|:------------------------:|:---------:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------|
|        `workers`         | `integer` | Number of workers that will be used to fetch the data and store it inside the database                                                                                                                                                                                    | `5`                                      |
|       `fast_sync`        | `boolean` | Whether Juno should use the fast sync abilities of different modules when enabled                                                                                                                                                                                         | `false`                                  |
|   `listen_new_blocks`    | `boolean` | Whether Juno should parse new blocks as soon as they get created                                                                                                                                                                                                          | `true`                                   | 
|     `parse_genesis`      | `boolean` | Whether Juno needs to parse the genesis state or not                                                                                                                                                                                                                      | `true`                                   |
|    `parse_old_blocks`    | `boolean` | Whether Juno should parse old chain blocks or not                                                                                                                                                                                                                         | `true`                                   | 
|      `start_height`      | `integer` | Height at which Juno should start parsing old blocks                                                                                                                                                                                                                      | `250000`                                 | 
|   `genesis_file_path`    | `string`  | Path of the genesis file to be parsed                                                                                                                                                                                                                                     | `'/bdjuno/.bdjuno/genesis/genesis.json'` |
| `re_enqueue_when_failed` | `boolean` | Whether blocks should be re-enqueued if the parsing inside various modules fails. (Default: `false`)                                                                                                                                                                      | `true`                                   |
|      `max_retries`       | `integer` | Number of retries that should be done before giving up on a block parsing. If set to `-1`, the retries will be unlimited. <br/> **Note**: The retry method is going to be based on exponential-backoff, with a wait time of `retryCount * avgBlockTime` per each attempt. | `-1`                                     |

## `database`
This section contains all the different configuration related to the PostgreSQL database where Juno will write the data.

|    Attribute     |   Type    | Description                                                                                                                                               | Example                                                                             |
|:----------------:|:---------:|:----------------------------------------------------------------------------------------------------------------------------------------------------------|:------------------------------------------------------------------------------------|
|      `url`       | `string`  | URI used to connect to the database. (e.g. [PostgreSQL Connection URI](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING-URIS). | `postgresql://user:password@localhost:5432/juno?sslmode=disable&search_path=public` |
| `partition_size` | `integer` | Number of rows that should be present inside each partition (used for transactions and messages).                                                         | `10000`                                                                             |

## `logging`
This section allows to configure the logging details of Juno.

| Attribute |   Type   | Description                                                             | Example |
|:---------:|:--------:|:------------------------------------------------------------------------|:--------|
| `format`  | `string` | Format in which the logs should be output (either `json` or `text`)     | `json`  | 
|  `level`  | `string` | Level of the log (either `verbose`, `debug`, `info`, `warn` or `error`) | `error` | 

## `pruning`
This section contains the configuration about the pruning options of the database. Note that this will have effect only
if you add the `"pruning"` entry to the `modules` field of the [`chain` config](#chain).

|   Attribute   |   Type    | Description                                                                                            | Example |
|:-------------:|:---------:|:-------------------------------------------------------------------------------------------------------|:--------|
|  `interval`   | `integer` | Number of blocks that should pass between one pruning and the other (default: prune every `10` blocks) | `100`   | 
| `keep_every`  | `integer` | Keep the state every `nth` block, even if it should have been pruned                                   | `500`   | 
| `keep_recent` | `integer` | Do not prune this amount of recent states                                                              | `100`   |

## `telemetry`
This section allows to configure the telemetry details of Juno. Note that this will have effect only if you add
the `"telemetry"` entry to the `modules` field of the [`chain` config](#chain).

| Attribute |  Type  | Description                                    | Example |
|:---------:|:------:|:-----------------------------------------------|:--------| 
|  `port`   | `uint` | Port on which the telemetry server will listen | `8000`  | 

**Note**  
If the telemetry server is enabled, a new endpoint at the provided port and path `/metrics` will
expose [Prometheus](https://prometheus.io/) data.
