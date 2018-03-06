package cron

import "time"

// ConstantDelaySchedule represents a simple recurring duty cycle, e.g. "Every 5 minutes".
// ConstantDelaySchedule 代表一个简单的循环占空比，例如 “每5分钟一次”。
// It does not support jobs more frequent than once a second.
// 它不会比每秒一次更频繁地支持工作。
type ConstantDelaySchedule struct {
	Delay time.Duration
}

// Every returns a crontab Schedule that activates once every duration.
// Every 返回一个定时任务调度，它每隔 duration 时间触发一次。
// Delays of less than a second are not supported (will round up to 1 second).
// 不支持小于1秒的延迟（将会取整到1秒）。
// Any fields less than a Second are truncated.
// 任何小于一秒的字段会被截断。
func Every(duration time.Duration) ConstantDelaySchedule {
	//如果duration小于1秒，设置为1秒
	if duration < time.Second {
		duration = time.Second
	}
	return ConstantDelaySchedule{
		//截断小于秒的部分，保持为整秒
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
	}
}

// Next returns the next time this should be run.
// Next 返回下一次调度运行的时间
// This rounds so that the next activation time will be on the second.
// 这轮，下一个激活时间以秒为单位
func (schedule ConstantDelaySchedule) Next(t time.Time) time.Time {
	return t.Add(schedule.Delay - time.Duration(t.Nanosecond())*time.Nanosecond)
}
