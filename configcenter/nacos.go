package configcenter

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/pkg/errors"
)

// NacosConfig nacos配置中心相关配置
type NacosConfig struct {
	Port      uint64
	host      string
	Namespace string
	Group     string
}

type NacosConfigCenter struct {
	group  string
	// nacos配置源实例
	cli config_client.IConfigClient
}

func (c *NacosConfigCenter) Get(key string) (string, error) {
	config, err := c.cli.GetConfig(vo.ConfigParam{
		DataId: key,
		Group:  c.group,
	})
	if err != nil {
		return "", errors.Wrapf(err, "NacosConfigCenter Get fail, key:%s", key)
	}
	return config, nil
}

func (c *NacosConfigCenter) GetUnmarshalJSON(key string, des interface{}) error {
	config, err := c.cli.GetConfig(vo.ConfigParam{
		DataId: key,
		Group:  c.group,
	})
	if err != nil {
		return errors.Wrapf(err, "NacosConfigCenter Get fail, key:%s", key)
	}

	if err := json.Unmarshal([]byte(config), des); err != nil {
		return errors.Wrapf(err, "NacosConfigCenter json Unmarshal fail, key:%s", key)
	}
	return nil
}

func NewNacosConfigCenter(cfg *NacosConfig) ConfigCenter {
	cc := constant.NewClientConfig(
		constant.WithNamespaceId(cfg.Namespace),
		constant.WithTimeoutMs(3000),
		constant.WithNotLoadCacheAtStart(true),
	)
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			cfg.host,
			cfg.Port,
		),
	}

	cli, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	})
	if err != nil {
		panic(any(err))
	}

	group := "DEFAULT_GROUP"
	if cfg.Group != "" {
		group = cfg.Group
	}
	return &NacosConfigCenter{
		group:  group,
		cli:    cli,
	}
}
