package client



type GrpcClient interface{
	getRichTopList() ([]Address, error)
	getAddrTag() ([]AddrTag, error)
}

type Address struct {
	Address string 
	Balance float64
}

type AddrTag struct {
	Address string 
	Name string
	Link string
}