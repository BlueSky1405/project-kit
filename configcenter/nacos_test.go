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

type testUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestNacosConfigCenter_GetUnmarshalJSON(t *testing.T) {
	cc := NewNacosConfigCenter(&NacosConfig{
		Port:  8848,
		host:  "43.138.178.12",
		Group: "DEFAULT_GROUP",
	})

	var res struct {
		User testUser `json:"user"`
	}

	err := cc.GetUnmarshalJSON("test", &res)
	require.Nil(t, err)
	require.Equal(t, "leslie", res.User.Name)
	require.Equal(t, 18, res.User.Age)
}
