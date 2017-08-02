package consts

const (
	TOPIC_OP_UNION        = "UNION"        //并集
	TOPIC_OP_INTERSECTION = "INTERSECTION" // 交集
	TOPIC_OP_EXCEPT       = "EXCEPT"       // 差集
)
const (
	Host      = "https://api.xmpush.xiaomi.com"
	//Host = "https://sandbox.xmpush.xiaomi.com"
)
const (
	RegURL        = "/v3/message/regid"        //向某个regid或一组regid列表推送某条消息（这些regId可以属于不同的包名)
	AliasURL      = "/v3/message/alias"        //向某个alias或一组alias列表推送某条消息（这些alias可以属于不同的包名）
	TopicURL      = "/v2/message/topic"        //向某个topic推送某条消息（可以指定一个或多个包名）
	TopicMultiURL = "/v3/message/multi_topic"  //向多个topic推送单条消息（可以指定一个或多个包名）
	Account       = "/v2/message/user_account" //向某个account或一组account列表推送某条消息
	All           = "/v3/message/all"          //向所有设备推送某条消息（可以指定一个或多个包名）
)

//
const (
	CounterURL = "/v1/stats/message/counters" //获取消息的统计数据正式环境API地址
)

//标签 订阅取消订阅标签
const (
	TopicRegIdSubscribURL    = "/v2/topic/subscribe"          //订阅RegId的标签
	TopicRegIdUnsubscirbeURL = "/v2/topic/unsubscribe"       //取消订阅RegId的标签
	TopicAliasSubscribeURL   = "/v2/topic/subscribe/alias"   //订阅alias的标签
	TopicAliasUnsubscirbeURL = "/v2/topic/unsubscribe/alias" //取消订阅alias的标签
)

//获取标签结果
const (
	TopicAllByRegidURL   = "/v1/topic/all"          //个用户目前订阅的所有Topic
)
