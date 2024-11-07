package config

type Log struct {
	Level   string   `json:"level" yaml:"level"`
	Channel []string `json:"channel" yaml:"channel"`
	Console Console  `json:"console" yaml:"console"`
	Lark    Lark     `json:"lark" yaml:"lark"`
	File    File     `json:"file" yaml:"file"`
}

// Lark 飞书
type Lark struct {
	Webhook string `json:"webhook"`
	Level   string `json:"level" yaml:"level"`
}

type Console struct {
	Level string `json:"level" yaml:"level"`
}

type File struct {
	// Filename 是要写入日志的文件。备份日志文件将保留在同一目录中。
	// 如果为空，它将使用 <processname>-lumberjack.log 在 os.TempDir() 中。
	Filename string `json:"filename" yaml:"filename"`

	// MaxSize 是日志文件在轮换之前的最大大小（以兆字节为单位）。默认为 100 兆字节。
	MaxSize int `json:"max_size" yaml:"max_size"`

	// MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
	// 注意，一天被定义为 24 小时，可能由于夏令时、闰秒等原因与日历天数不完全对应。
	// 默认情况下，不会根据年龄删除旧日志文件。
	MaxAge int `json:"max_age" yaml:"max_age"`

	// MaxBackups 是要保留的旧日志文件的最大数量。默认情况下保留所有旧日志文件
	// （尽管 MaxAge 仍可能导致它们被删除）。
	MaxBackups int `json:"max_backups" yaml:"max_backups"`

	// LocalTime 决定用于格式化备份文件中时间戳的时间是否为计算机的本地时间。
	// 默认情况下使用 UTC 时间。
	LocalTime bool `json:"local_time" yaml:"local_time"`

	// Compress 决定轮换的日志文件是否应使用 gzip 压缩。默认情况下不进行压缩。
	Compress bool `json:"compress" yaml:"compress"`

	Level string `json:"level" yaml:"level"`
}
