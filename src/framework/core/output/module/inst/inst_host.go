package inst

import (
	"configcenter/src/framework/core/output/module/model"
	"configcenter/src/framework/core/types"
)

type host struct {
	target model.Model
	datas  types.MapStr
}

func (cli *host) GetModel() model.Model {
	return cli.target
}

func (cli *host) IsMainLine() bool {
	// TODO：判断当前实例是否为主线实例
	return true
}

func (cli *host) GetAssociationModels() ([]model.Model, error) {
	// TODO:需要读取此实例关联的实例，所对应的所有模型
	return nil, nil
}

func (cli *host) GetInstID() int {
	return 0
}
func (cli *host) GetInstName() string {
	return ""
}

func (cli *host) GetValues() (types.MapStr, error) {
	return nil, nil
}

func (cli *host) GetAssociationsByModleID(modleID string) ([]Inst, error) {
	// TODO:获取当前实例所关联的特定模型的所有已关联的实例
	return nil, nil
}

func (cli *host) GetAllAssociations() (map[model.Model][]Inst, error) {
	// TODO:获取所有已关联的模型及对应的实例
	return nil, nil
}

func (cli *host) SetParent(parentInstID int) error {
	return nil
}

func (cli *host) GetParent() ([]Topo, error) {
	return nil, nil
}

func (cli *host) GetChildren() ([]Topo, error) {
	return nil, nil
}

func (cli *host) SetValue(key string, value interface{}) error {

	// TODO:需要根据model 的定义对输入的key 及value 进行校验

	cli.datas[key] = value

	return nil
}

func (cli *host) Save() error {
	return nil
}
