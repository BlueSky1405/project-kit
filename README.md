# project-kit

海岸蓝天kit工具包

## kit项目结构

```
.
├── README.md
├── configcenter
│   ├── configcenter.go         // 配置中心通用接口
│   ├── nacos.go                // nacos配置中心实现
│   └── nacos_test.go           // 测试及事例
├── cron
│   ├── cron.go                 // 定时器封装
│   └── cron_test.go            // 测试及事例
├── database
│   └── mysql
│       ├── mysql.go            // mysql连接
│       └── mysql_test.go       // 测试及事例
├── go.mod
├── go.sum
└── log
    ├── log.go                  // 日志通用接口
    ├── zap.go                  // zap日志实现
    └── zap_test.go             // 测试及事例
```
