package main

type Connection struct{}

func (c Connection) SendTX(t Transaction) (string, error) {
	return "", nil
}
func ConnectBitcoinNetwork() (Connection, error) {
	return Connection{}, nil
}
