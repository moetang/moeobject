package file

type SvFileSegOpt struct {
	//TODO simple, replica, ec
	Type byte
}

type StorageVirt interface {
	//TODO
	NewFileSeg(opt SvFileSegOpt) (fileSegId []byte, err error)
	LoadFileSeg(fileSegId []byte) error

	Store()
	Get()
	Modify()
}
