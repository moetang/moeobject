package file

type PhyFileOpt struct {
	Sv StorageVirt
	//TODO capacity
}

type ObjectId struct {
}

type PutOpt struct {
	NewFileSeg bool
}

type PhyFile2 struct {
}

func NewPhyFile2(opt PhyFileOpt) {
	//TODO
}

func LoadPhyFile2(opt PhyFileOpt) {
	//TODO
}

func (this *PhyFile2) GetObject(id ObjectId) ([]byte, error) {

}

func (this *PhyFile2) PutObject(data []byte, putOpt PutOpt) (ObjectId, error) {
	//TODO only latest FileSeg can be written at the same time
}

func (this *PhyFile2) ModifyObject(data []byte, id ObjectId) ([]byte, error) {

}
