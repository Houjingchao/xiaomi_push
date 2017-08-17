package model

import (
	"time"
	"strings"
	"encoding/json"
	"github.com/Houjingchao/xiaomi_push/consts"
)

// Create by LittleM 17/07/20
const (
	MaxTimeToSend = time.Hour * 24 * 7
	MaxTimeToLive = time.Hour * 24 * 7 * 2
)

type Message struct {
	PackageName string            //App的包名
	Payload     string            //消息的内容
	Title       string            // 通知的标题
	Description string            //通知的描述
	PassThrough int64             // 0 表示通知栏消息  1 表示透传消息
	NotifyType  int64             //DEFAULT_ALL = -1; DEFAULT_SOUND  = 1;  // 使用默认提示音提示； DEFAULT_VIBRATE = 2;  // 使用默认震动提示； DEFAULT_LIGHTS = 4;   // 使用默认led灯光提示；
	TimeToLive  int64             // 可选项。如果用户离线，设置消息在服务器保存的时间，单位：ms。服务器默认最长保留两周。
	TimeToSend  int64             // 可选项。定时发送消息。用自1970年1月1日以来00:00:00.0 UTC时间表示（以毫秒为单位的时间）。注：仅支持七天内的定时消息。
	NotifyID    int64             // 可选项。默认情况下，通知栏只显示一条推送消息。如果通知栏要显示多条推送消息，需要针对不同的消息设置不同的notify_id（相同notify_id的通知栏消息会覆盖之前的）。
	Extra       map[string]string // 可选项，对app提供一些扩展的功能，请参考2.2。除了这些扩展功能，开发者还可以定义一些key和value来控制客户端的行为。注：key和value的字符数不能超过1024，至多可以设置10个key-value键值对。
	TopicOp     string
}

func (m *Message) SetPackageName(packageNames []string) *Message {
	m.PackageName = strings.Join(packageNames, ",")
	return m
}
func (m *Message) SetPayload(payload string) *Message {
	m.Payload = payload
	return m
}
func (m *Message) SetTitle(title string) *Message {
	m.Title = title
	return m
}
func (m *Message) SetDescription(description string) *Message {
	m.Description = description
	return m
}

func (m *Message) SetPassThrough(passThrough int64) *Message {
	m.PassThrough = passThrough
	return m
}

func (m *Message) SetNotifyType(notifyType int64) *Message {
	m.NotifyType = notifyType
	return m
}
func (m *Message) SetTimeToSend(tts int64) *Message {
	if time.Since(time.Unix(0, tts*int64(time.Millisecond))) > MaxTimeToSend {
		m.TimeToSend = time.Now().Add(MaxTimeToSend).UnixNano() / 1e6
	} else {
		m.TimeToSend = tts
	}
	return m
}

func (m *Message) SetTimeToLive(ttl int64) *Message {
	if time.Since(time.Unix(0, ttl*int64(time.Millisecond))) > MaxTimeToLive {
		m.TimeToLive = time.Now().Add(MaxTimeToLive).UnixNano() / 1e6
	} else {
		m.TimeToLive = ttl
	}
	return m
}

func (m *Message) SetNotifyID(notifyID int64) *Message {
	m.NotifyID = notifyID
	return m
}

func (m *Message) SetExtra(extra map[string]string) *Message {
	m.Extra = extra
	return m
}

func (m *Message) AddExtra(key, value string) *Message {
	m.Extra[key] = value
	return m
}

//小米推送服务器每隔1s将已送达或已点击的消息ID和对应设备的regid或alias通过调用第三方http接口传给开发者。（每次调用后，小米推送服务器会清空这些数据，下次传给开发者将是新一拨数据。）
//注：消息的送达回执只支持向regId或alias发送的消息。

func (m *Message) JSON() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return bytes
}
func (m *Message) SetCallback(callBackUrl string) *Message {
	m.Extra["callback"] = callBackUrl
	m.Extra["callback.type"] = "3" // 1:送达回执, 2:点击回执, 3:送达和点击回执
	return m
}

func NewAndroidMessage(payload, title, description string,extra map[string]string) *Message {
	return &Message{
		Payload:     payload,//消息的内容
		Title:      title,
		Description: description,
		PassThrough: 0, // 1表示透传消息，0表示通知栏消息。
		NotifyType:  1, // 使用默认提示音提示
		TimeToLive:  0,
		TimeToSend:  0,
		NotifyID:    time.Now().Unix(), //可选项。默认情况下，通知栏只显示一条推送消息。如果通知栏要显示多条推送消息，需要针对不同的消息设置不同的notify_id（相同notify_id的通知栏消息会覆盖之前的）。
		Extra:       extra,
		TopicOp:     consts.TOPIC_OP_UNION,
	}
}

type TargetType int32

const (
	Regids   TargetType = 1
	Alias    TargetType = 2
	Accounts TargetType = 3
	Topic    TargetType = 4
	Topics   TargetType = 5
	All      TargetType = 6
)

type TargetedMessage struct {
	message    *Message
	targetType TargetType
	target     string
}

func NewTargetedMessage(m *Message, target string, targetType TargetType) *TargetedMessage {
	return &TargetedMessage{
		message:    m,
		targetType: targetType,
		target:     target,
	}
}

func (tm *TargetedMessage) SetTargetType(targetType TargetType) *TargetedMessage {
	tm.targetType = targetType
	return tm
}

func (tm *TargetedMessage) SetTarget(target string) *TargetedMessage {
	tm.target = target
	return tm
}

func (tm *TargetedMessage) JSON() []byte {
	bytes, err := json.Marshal(tm)
	if err != nil {
		panic(err)
	}
	return bytes
}
