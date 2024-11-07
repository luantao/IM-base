package limit

import (
	"MyIM/pkg/config"
	"MyIM/pkg/mlog"
	"errors"
	"os"
	"strings"
	"sync"

	sentinel "github.com/alibaba/sentinel-golang/api"
	sentinelbase "github.com/alibaba/sentinel-golang/core/base"
	sentinelconfig "github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/logging"
	"go.uber.org/zap"
)

var defaultOnce sync.Once
var defaultRules *Rules

// Init 初始化默认 Limiter 配置
func Init(section string, logger *Logger) {
	defaultOnce.Do(func() {
		if logger == nil {
			panic("limit init failed: logger invalid")
		}

		logging.ResetGlobalLogger(logger)

		// 环境变量优先
		configPath := os.Getenv(sentinelconfig.ConfFilePathEnvKey)
		if configPath != "" {
			err := sentinelconfig.InitConfigWithYaml(configPath)
			if err == nil {
				return
			}
			mlog.Logger().Named("limit").Panic("new sentinel error", zap.Error(err))
		}

		// 使用 Config 配置
		cfg, err := NewConfigWithACM(section)
		if err != nil {
			mlog.Logger().Named("limit").Panic("new sentinel error", zap.Error(err))
		}

		config := &sentinelconfig.Entity{
			Version: "v1",
			Sentinel: sentinelconfig.SentinelConfig{
				App: struct {
					Name string
					Type int32
				}{
					Name: cfg.AppName,
					Type: AppTypeDefault,
				},
				Log: sentinelconfig.LogConfig{
					Logger: logger,
					Dir:    cfg.LogPath,
					UsePid: false,
					Metric: sentinelconfig.MetricLogConfig{
						SingleFileMaxSize: cfg.MetricMaxFileSize,
						MaxFileCount:      cfg.MetricMaxFileCount,
						FlushIntervalSec:  cfg.MetricFlushInterval,
					},
				},
				Stat: sentinelconfig.StatConfig{
					GlobalStatisticSampleCountTotal: sentinelbase.DefaultSampleCountTotal,
					GlobalStatisticIntervalMsTotal:  sentinelbase.DefaultIntervalMsTotal,
					MetricStatisticSampleCount:      sentinelbase.DefaultSampleCount,
					MetricStatisticIntervalMs:       sentinelbase.DefaultIntervalMs,
					System: sentinelconfig.SystemStatConfig{
						CollectIntervalMs:       cfg.SystemCollectInterval,
						CollectLoadIntervalMs:   sentinelconfig.DefaultLoadStatCollectIntervalMs,
						CollectCpuIntervalMs:    sentinelconfig.DefaultCpuStatCollectIntervalMs,
						CollectMemoryIntervalMs: sentinelconfig.DefaultMemoryStatCollectIntervalMs,
					},
				},
				UseCacheTime: true,
			},
		}

		err = sentinel.InitWithConfig(config)
		if err != nil {
			mlog.Logger().Named("limit").Panic("init sentinel error", zap.Error(err))
		}

		defaultRules = NewRules()
	})
}

// Default 默认规则集合
func Default() *Rules {
	return defaultRules
}

// NewConfigWithACM 从 Config 构造配置信息
func NewConfigWithACM(section string) (conf *Config, err error) {
	buildKey := func(key string) string {
		return strings.Join([]string{section, key}, ".")
	}

	appName := config.GetString(buildKey("app_name"))
	if appName == "" {
		err = errors.New("limit app name not set")
		return
	}

	logPath := config.GetString(buildKey("log_path"))
	if logPath == "" {
		err = errors.New("limit log path not set")
		return
	}

	logLevel := config.GetString(buildKey("log_level"))
	if logLevel == "" {
		logLevel = mlog.InfoLevel
	}

	metricMaxFileCount := config.GetUint32(buildKey("metric_max_file_count"))
	if metricMaxFileCount == 0 {
		metricMaxFileCount = DefaultMetricMaxFileCount
	}

	metricMaxFileSize := config.GetUint64(buildKey("metric_max_file_size"))
	if metricMaxFileSize == 0 {
		metricMaxFileSize = DefaultMetricMaxFileSize
	}

	metricFlushInterval := config.GetUint32(buildKey("metric_flush_interval"))
	if metricFlushInterval == 0 {
		metricFlushInterval = DefaultMetricFlushInterval
	}

	systemCollectInterval := config.GetUint32(buildKey("system_collect_interval"))
	if systemCollectInterval == 0 {
		systemCollectInterval = DefaultSystemCollectInterval
	}

	conf = &Config{
		AppName:               appName,
		LogPath:               logPath,
		LogLevel:              logLevel,
		MetricMaxFileCount:    metricMaxFileCount,
		MetricMaxFileSize:     metricMaxFileSize,
		MetricFlushInterval:   metricFlushInterval,
		SystemCollectInterval: systemCollectInterval,
	}
	return
}
