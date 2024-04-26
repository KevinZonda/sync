package node

type Master struct {
}

func (m *Master) Start() {

}

type MasterConfig struct {
	SlaveList []string
}
