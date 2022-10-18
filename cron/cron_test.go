package cron

import (
	"fmt"
	"github.com/BlueSky1405/project-kit/log"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestAsd(t *testing.T) {
	cron := NewCron(log.NewZapLogger(""))
	defer cron.Stop()

	err := cron.AddFunc("我是定时器的测试", "* * * * * *", func() {
		fmt.Println("南哥来广州，但是我没有找他")
	})
	require.Nil(t, err)

	cron.Start()

	time.Sleep(time.Second * 10)
}
