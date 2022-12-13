# tictactoe

**tictactoe** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Transactions

```
Usage:
  tic-tac-toed tx tictactoe [flags]
  tic-tac-toed tx tictactoe [command]

Available Commands:
  accept-game  Broadcast message AcceptGame
  create-game  Broadcast message create-game
  perform-move Broadcast message PerformMove

Flags:
  -h, --help   help for tictactoe

Global Flags:
      --chain-id string     The network chain ID (default "tictactoe")
      --home string         directory for config and data (default "/Users/user/.tic-tac-toe")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

## Queries

```
Usage:
  tic-tac-toed query tictactoe [flags]
  tic-tac-toed query tictactoe [command]

Available Commands:
  games       Query games
  open-games  Query OpenGames
  params      shows the parameters of the module

Flags:
  -h, --help   help for tictactoe

Global Flags:
      --chain-id string     The network chain ID (default "tictactoe")
      --home string         directory for config and data (default "/Users/hendrikoosenbrug/.tic-tac-toe")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

## How to Play

1. Create a game with the create-game command or query open games with the open-games query.
2. Accept an invitation with the accept-game command (must pass in the game id)
3. Starting turn is determine by cryptographic method.
4. Query games and check who's turn is next.
5. Perform a move with the perform-move command and pass in a valid index on the tictactoe board (0-8)
6. Game ends if there's a winner or draw.

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/tic-tac-toe@latest! | sudo bash
```

`username/tic-tac-toe` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
