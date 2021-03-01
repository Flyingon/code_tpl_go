package config

type ChannelInfo struct {
	ScheduleQpm int    `json:"schedule_qpm,omitempty"`
	ChannelName string `json:"channel_name,omitempty"`
}

var ScheduleConf struct {
	Channels                []ChannelInfo `json:"channels,omitempty"`
	Default                 ChannelInfo   `json:"default,omitempty"`
	TotalScheduleQpm        int           `json:"total_schedule_qpm,omitempty"`
	ReadyQueuePushQpm       int           `json:"ready_queue_push_qpm,omitempty"`
	ReadyQueuePopBatchNum   int           `json:"ready_queue_pop_batch_num,omitempty"`
	WaitingQueuePopBatchNum int           `json:"waiting_queue_pop_batch_num,omitempty"`

	WaitingQueueName     string `json:"waiting_queue_name"`
	ReadyQueueNamePrefix string `json:"ready_queue_name_prefix"`
}
