# How to use a testnet snapshot in localnet?

```
make install
go run ./scripts/upgrade-assure/... https://snapshots.polkachu.com/testnet-snapshots/elys/elys_5724942.tar.lz4 ~/go/bin/elysd ~/go/bin/elysd --skip-proposal
```

# How can I perform a test with a version upgrade that involves extensive changes to data structures?

```
git checkout v0.28.1
make install
cp -a ~/go/bin/elysd /tmp/elysd-v0.28.1
```

```
go run ./scripts/upgrade-assure/... --home /tmp/elys https://snapshots.polkachu.com/testnet-snapshots/elys/elys_5511381.tar.lz4 /tmp/elysd-v0.28.1 /tmp/elysd-v0.29.0 --skip-node-start
```

```
git checkout v0.29.0
make install
cp -a ~/go/bin/elysd /tmp/elysd-v0.29.0
```

```
go run ./scripts/upgrade-assure/... --home /tmp/elys https://snapshots.polkachu.com/testnet-snapshots/elys/elys_5511381.tar.lz4 /tmp/elysd-v0.28.1 /tmp/elysd-v0.29.0 --skip-snapshot --skip-chain-init
```
