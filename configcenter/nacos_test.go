package configcenter

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNacosConfigCenter_Get(t *testing.T) {
	cc := NewNacosConfigCenter(&NacosConfig{
		Port:  8848,
		host:  "43.138.178.12",
		Group: "DEFAULT_GROUP",
	})

	cfg, err := cc.Get("test")
	require.Nil(t, err)
	fmt.Println(cfg)
}
