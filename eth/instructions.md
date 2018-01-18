- start node and blockchain syncing in testnet
```
geth --testnet --datadir ~/eth-testnet
```

- create first account (wallet)
```
geth --testnet --datadir ~/eth-testnet account new
```

- mine ethers in testnet
```
geth --testnet --datadir ~/eth-testnet --mine
```

- serve dapp running on localhost:8000 via RPC
```
geth --testnet --datadir ~/eth-testnet --rpc --rpccorsdomain "http://localhost:8000" --rpcapi eth,web3,personal
```


- generate ABI from contract.sol
```
solcjs --abi HelloWorldContract.sol
```

- generate the byte code
```
solcjs --bin HelloWorldContract.sol
```


- deploy contract via geth js console
    - open geth console
    ```
    geth --testnet --datadir ~/eth-testnet console
    ```

    - in the console create the contract using ABI and BYTE CODE result from compilation
    ```js
    var abi = [{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"bytes32"}],"payable":false,"type":"function"}];
    var code = '0x60606040527f48656c6c6f20576f726c64000000000000000000000000000000000000000000600090600019169055341561003957600080fd5b5b609d806100486000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde0314603d575b600080fd5b3415604757600080fd5b604d606b565b60405180826000191660001916815260200191505060405180910390f35b600054815600a165627a7a72305820f4c510e24a238337d5334b5b38a44e88ea53ef40f26aeb96eba4609cb72827cd0029';
    web3.personal.unlockAccount(eth.accounts[0], 'PASSWORD');
    var inputdata = ["pub", "h", "u", "sv", "signerid"];
    var contract = web3.eth.contract(abi).new(inputdata,{ from: eth.accounts[0], data: code, gas: 1000000 });
    web3.personal.lockAccount(eth.accounts[0]);
    ```

    - call the method in the contract
    ```js
    contract.name()
    ```
