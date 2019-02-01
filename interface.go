package xnet

//todo  try it's really can use interface.
type Server interface {
	MainLogic() string
}

type Client interface {
}
