package model

const (
	// 任务类型
	HttpTask = "http"    //http请求
	CronTask = "crontab" //linux定时任务

	// 任务状态
	TaskInit      = 0 // 初始化(未生效)
	TaskEffective = 1 // 生效

	GetRequest    = "get"
	PostRequest   = "post"
	DeleteRequest = "delete"
)

// 定时任务
type TimedTask struct {
	TaskName string `json:"task_name"` // 任务名称
	TaskCode string `json:"task_code"` // 任务唯一标识
	Type     string `json:"type"`      // 类型:http或crontab
	// 时间表达式
	CronExpr string `json:"cron_expr"` // 表达式, 如 * 12 * * * *
	// 执行命名
	Command       string    `json:"command"`        // 命令, 如https://www.golang.org/
	RequestMethod string    `json:"request_method"` // 请求方法, 如get post
	Parameters    string    `json:"parameters"`     // 参数
	Status        int       `json:"status"`         // 状态
	Base          BaseModel `json:"base"`
}

// 项目属性
type Project struct {
	ProjectName string `json:"project_name"`
	BaseModel   BaseModel
}

// 任务执行日志
type TaskLog struct {
	ServerHost  string
	Command     string
	RespJsonMsg string
}
