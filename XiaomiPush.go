package xiaomi_push

import (
	"github.com/cocotyty/httpclient"
	"code.aliyun.com/app-channel-adapter/api/xiaomi_push/model"
	"code.aliyun.com/app-channel-adapter/api/xiaomi_push/consts"
	"strconv"
	"fmt"
)

type XiaomiPush struct {
	appSecret   string
	packageNmae []string
	host        string
}

func NewXiaomiPush(appSecret string, packageName []string, host string) *XiaomiPush {
	return &XiaomiPush{
		appSecret:   appSecret,
		packageNmae: packageName,
		host:        host,
	}
}
func buildRequest(meq *model.Message, request *httpclient.HttpRequest) *httpclient.HttpRequest {
	request.
	Param("payload", meq.Payload).
		Param("notify_type", strconv.FormatInt(meq.NotifyType, 10)). //DEFAULT_VIBRATE = 2;  // 使用默认震动提示；
		Param("title", meq.Title). //知栏展示的通知的标题
		Param("description", meq.Description). //通知栏展示的通知的描述
		Param("restricted_package_name", meq.PackageName). //目前只是一个
		Param("pass_through", strconv.FormatInt(meq.PassThrough, 10)).
		Param("longtime_to_live", strconv.FormatInt(meq.TimeToLive, 10)).
		Param("time_to_send", strconv.FormatInt(meq.TimeToSend, 10))

	maps := meq.Extra
	if maps != nil {
		for k, v := range maps {
			request.Param("extra."+k, v)
		}
	}
	return request
}
func (xm *XiaomiPush) SendRegids(meq *model.Message, regisIds string) (response string, err error) {
	request := httpclient.
	Post(consts.Host + consts.RegURL).
		Head("Authorization", "key="+xm.appSecret).
		Param("registration_id", regisIds) //,号分开
	request = buildRequest(meq, request)
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (xm *XiaomiPush) SendAlias(meq *model.Message, alias string) (response string, err error) {
	request := httpclient.
	Post(consts.Host + consts.AliasURL).
		Head("Authorization", "key="+xm.appSecret).
		Param("alias", alias) //,号分开
	request = buildRequest(meq, request)
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (xm *XiaomiPush) SendAccounts(meq *model.Message, user_account string) (response string, err error) {
	request := httpclient.
	Post(consts.Host + consts.Account).
		Head("Authorization", "key="+xm.appSecret).
		Param("user_account", user_account) //,号分开
	request = buildRequest(meq, request)
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}
func (xm *XiaomiPush) SendTopic(meq *model.Message, topic string) (response string, err error) {
	fmt.Println(xm.packageNmae[0])
	request := httpclient.
	Post(consts.Host + consts.TopicURL).
		Head("Authorization", "key="+xm.appSecret).
		Param("topic", topic). //,号分开
		Param("restricted_package_name", xm.packageNmae[0]) //,号分开
	request = buildRequest(meq, request)
	result, err := request.Send().
		String()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return result, nil
}
func (xm *XiaomiPush) SendTopics(meq *model.Message, topics string, topicop string) (response string, err error) {

	request := httpclient.
	Post(consts.Host + consts.TopicMultiURL).
		Head("Authorization", "key="+xm.appSecret).
		Param("topics", topics). //;$;分割  不可以超过5
		Param("topic_op", meq.TopicOp)
	request = buildRequest(meq, request)
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (xm *XiaomiPush) SendAll(meq *model.Message) (response string, err error) {

	request := httpclient.
	Post(consts.Host + consts.All).
		Head("Authorization", "key="+xm.appSecret)
	request = buildRequest(meq, request)
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}

/********************************************统计查询******************************************************************/
func (xm *XiaomiPush) GetCounters(startDate string, endDate string) (response string, err error) { //时间格式yyyyMMdd

	request := httpclient.
	Get(consts.Host + consts.CounterURL).
		Head("Authorization", "key="+xm.appSecret).
		Query("start_date", startDate).
		Query("end_date", endDate).
		Query("restricted_package_name", xm.packageNmae[0])
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}

/************************************************Topic订阅或者取消根据 registration_id******************************************************************/
/**
 * 	设备regid列表，逗号分隔，必填。限制：最多1000个regid。
 */
func (xm *XiaomiPush) TopicRegIdSubscrib(topic, registration_id string) (response string, err error) {
	request := httpclient.
	Post(consts.Host + consts.TopicRegIdSubscribURL).
		Head("Authorization", "key="+xm.appSecret).
		Param("registration_id", registration_id).
		Param("topic", topic).
		Param("restricted_package_name", xm.packageNmae[0])
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	fmt.Println(request)
	return result, nil
}

func (xm *XiaomiPush) TopicRegIdUnsubscirbe(topic, registration_id string) (response string, err error) {
	request := httpclient.
	Post(consts.Host + consts.TopicRegIdUnsubscirbeURL).
		Head("Authorization", "key="+xm.appSecret).
		Param("registration_id", registration_id).
		Param("topic", topic).
		Param("restricted_package_name", xm.packageNmae[0])
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}

/************************************************Topic订阅或者取消根据 alias******************************************************************/

/**
 *	各个alias之间用逗号分割，必填。限制：最多1000个alias。
 */
func (xm *XiaomiPush) TopicAliasSubscribe(topic, aliases string) (response string, err error) {
	request := httpclient.
	Post(consts.Host + consts.TopicAliasSubscribeURL).
		Head("Authorization", "key="+xm.appSecret).
		Param("aliases", aliases).
		Param("topic", topic).
		Param("restricted_package_name", xm.packageNmae[0])
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (xm *XiaomiPush) TopicAliasUnsubscirbe(topic, aliases string) (response string, err error) {
	request := httpclient.
	Post(consts.Host + consts.TopicAliasUnsubscirbeURL).
		Head("Authorization", "key="+xm.appSecret).
		Param("aliases", aliases).
		Param("topic", topic).
		Param("restricted_package_name", xm.packageNmae[0])
	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}

/************************************************获取Regid 的所有topic******************************************************************/
func (xm *XiaomiPush) TopicAllByRegid(regid string) (response string, err error) {
	fmt.Println(regid)
	request := httpclient.
	Get(consts.Host + consts.TopicAllByRegidURL).
		Head("Authorization", "key="+xm.appSecret).
		Query("restricted_package_name", xm.packageNmae[0]).
		Query("registration_id", regid)

	result, err := request.Send().
		String()
	if err != nil {
		return "", err
	}
	return result, nil
}

