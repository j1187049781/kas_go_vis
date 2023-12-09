package client



type GrpcClient interface{
	getRichTopList() ([]Address, error)
	getAddrTag() ([]AddrTag, error)
}

type Address struct {
	Address string 
	Balance float64

	Tags []AddrTag
}

type AddrTag struct {
	Address string 
	Name string
	Link string
}