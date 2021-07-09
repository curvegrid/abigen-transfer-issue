# abigen crash on go-ethereum v1.10.4

To reproduce:
1. Clone this repo
2. Edit `main.go` with a suitable blockchain node websockets endpoint. **The endpoint must be websockets to trigger the crash.**
3. Trigger the crash:
```
$ go run .
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x436a575]

goroutine 1 [running]:
github.com/ethereum/go-ethereum/rpc.(*Client).send(0xc000590080, 0x0, 0x0, 0xc000418140, 0x469d6e0, 0xc00042c0e0, 0xc00042c0e0, 0x0)
	/Users/jeff/golang/pkg/mod/github.com/curvegrid/go-ethereum@v1.10.4-curvegrid-1/rpc/client.go:487 +0x55
github.com/ethereum/go-ethereum/rpc.(*Client).CallContext(0xc000590080, 0x0, 0x0, 0x45ddda0, 0xc000416030, 0x46ee509, 0x14, 0xc000412020, 0x2, 0x2, ...)
	/Users/jeff/golang/pkg/mod/github.com/curvegrid/go-ethereum@v1.10.4-curvegrid-1/rpc/client.go:304 +0x3c7
github.com/ethereum/go-ethereum/ethclient.(*Client).HeaderByNumber(0xc000598008, 0x0, 0x0, 0x0, 0x20f1485946613fac, 0xd699ffb, 0x0)
	/Users/jeff/golang/pkg/mod/github.com/curvegrid/go-ethereum@v1.10.4-curvegrid-1/ethclient/ethclient.go:182 +0x165
github.com/ethereum/go-ethereum/accounts/abi/bind.(*BoundContract).transact(0xc0005ea280, 0xc0005cc070, 0xc0005ea280, 0xc0005d40f0, 0x44, 0x50, 0xa, 0x0, 0x4c2af10)
	/Users/jeff/golang/pkg/mod/github.com/curvegrid/go-ethereum@v1.10.4-curvegrid-1/accounts/abi/bind/base.go:232 +0x123
github.com/ethereum/go-ethereum/accounts/abi/bind.(*BoundContract).Transact(0xc0005ea280, 0xc0005cc070, 0x46e1ea6, 0x4, 0xc0003d9f58, 0x2, 0x2, 0x0, 0xc00004e710, 0x40400f6)
	/Users/jeff/golang/pkg/mod/github.com/curvegrid/go-ethereum@v1.10.4-curvegrid-1/accounts/abi/bind/base.go:190 +0x185
main.(*TransferTransactor).Send(...)
	/Users/jeff/projects/curvegrid/abigen-transfer-issue/transfer.go:226
main.main()
	/Users/jeff/projects/curvegrid/abigen-transfer-issue/main.go:52 +0x474
exit status 2
```

Note that we are using a version of go-ethereum with [#23102](https://github.com/ethereum/go-ethereum/pull/23102) included so we can work with ~pre-Berlin~ pre-EIP1559 clients.

Optional:
- Use `generate.sh` to regenerate the abigen'd `transfer.go` from `Transfer.sol`
