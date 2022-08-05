# how to deploy contract

install truffle

https://trufflesuite.com/docs/truffle/getting-started/compiling-contracts/

```shell
cd $PROJECT
truffle init
truffle compile
cd build/contracts
export N=filename
cat $N.json|jq -c .abi > $N.abi
cat $N.json|jq .bytecode |sed s/\"//g> $N.bin
```

```shell
abigen --bin=./$N.bin --abi=./$N.abi --pkg=$N --out=$N.go
```
