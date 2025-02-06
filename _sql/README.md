```go
type Machine struct {
	Driver    string `json:"driver"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Database  string `json:"database"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	Collation string `json:"collation"`
}
type Master struct {
	Count       int        `json:"count"`
	MachineList []*Machine `json:"machine_list"`
}
type Slave struct {
	Count       int        `json:"count"`
	MachineList []*Machine `json:"machine_list"`
}
type Cluster struct {
	Master *Master
	Slave  *Slave
}
type Business struct {
	Count       int        `json:"count"`
	ClusterList []*Cluster `json:"cluster_list"`
}

var Conf map[string]*Business

```
_sql组件仅上线底层数据库操作，所需要的信息，比如配置machine，表table，均由上层计算后传入 \
比如分库分表等，需要在上层进行计算，最终的到machine和table，传入，执行

