// configcenter - 配置中心包

package configcenter

// ConfigCenter 配置中心通用接口
type ConfigCenter interface {
	// Get 根据key获取配置项
	Get(key string) (string, error)
	// GetUnmarshalJSON 根据key获取配置项，并json反序列化至des内，des需要以指针传入且tag需要支持
	GetUnmarshalJSON(key string, des interface{}) error
}
