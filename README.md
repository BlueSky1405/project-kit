# project-kit
海岸蓝天kit工具包

## kit项目结构

```
.
├── README.md               // kit包介绍
├── configcenter
│   ├── configcenter.go     // 配置中心通用接口
│   ├── nacos.go            // 配置中心基于nacos的实现
│   └── nacos_test.go       // 测试
├── go.mod
├── go.sum
└── log
    ├── log.go              // 日志包通用接口
    ├── zap.go              // 日志包基于zap的实现
    └── zap_test.go         // 测试

```
