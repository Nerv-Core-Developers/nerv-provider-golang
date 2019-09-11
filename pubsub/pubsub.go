package pubsub

import (
	"encoding/json"
	"errors"
	"sync"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/schema"
	"github.com/Nerv-Core-Developers/nerv-provider-golang/shared/utils"
)

//only support single level topic
const resultTopic string = "cmdresult"

type PubsubService struct {
	topicSubscriberWatchers map[string][]string
	topics                  map[string]map[string]chan interface{}
	subscribers             map[string]map[string]bool
	lock                    sync.RWMutex
}

func (ps *PubsubService) Init() error {
	ps.topicSubscriberWatchers = make(map[string][]string)
	ps.topics = make(map[string]map[string]chan interface{})
	ps.subscribers = make(map[string]map[string]bool)
	return nil
}

func (ps *PubsubService) ProcessCmd(buf []byte, msgID string, requestorID string, requestorCh chan interface{}) {
	var result []byte
	isSuccess := true
	var msgPubsub schema.MsgPubsubCmd
	msgPubsub.Init(buf, 0)
	topic := string(msgPubsub.Topic())
	data := msgPubsub.DataBytes()
	switch msgPubsub.Cmd() {
	case schema.PubsubCmdSub:
		ps.subTopic(topic, requestorID, requestorCh)
	case schema.PubsubCmdUnsub:
		err := ps.unsubTopic(topic, requestorID)
		if err != nil {
			result = []byte(err.Error())
			isSuccess = false
		}
	case schema.PubsubCmdPubAll:
		err := ps.publish(msgID, topic, data)
		if err != nil {
			result = []byte(err.Error())
			isSuccess = false
		}
	case schema.PubsubCmdPubSpec:
		err := ps.publishSpecific(msgID, topic, data, requestorID)
		if err != nil {
			result = []byte(err.Error())
			isSuccess = false
		}
	case schema.PubsubCmdListTopic:
		result = ps.getListTopic()
	case schema.PubsubCmdListSubInTopic:
		result = ps.getListSubscribersOfTopic(topic)
	case schema.PubsubCmdWatchTopicSub:
		err := ps.addTopicSubWatcher(requestorID, topic)
		if err != nil {
			result = []byte(err.Error())
			isSuccess = false
		}
	default:
		err := errors.New("Unknown cmd")
		if err != nil {
			result = []byte(err.Error())
			isSuccess = false
		}
	}
	sendCmdResult(msgID, result, isSuccess, requestorCh)
}

func sendCmdResult(msgID string, result []byte, success bool, requestorCh chan interface{}) {

	builder := flatbuffers.NewBuilder(0)

	dataBytes := builder.CreateByteVector(result)

	schema.PubsubCmdResultStart(builder)
	schema.PubsubCmdResultAddSuccess(builder, success)
	schema.PubsubCmdResultAddContent(builder, dataBytes)
	responseBytes := schema.PubsubCmdResultEnd(builder)

	builder.Finish(responseBytes)
	cmdResult := builder.FinishedBytes()
	select {
	case requestorCh <- buildMsgPubsub(msgID, resultTopic, cmdResult):
		return
	default:
		return
	}
}

func (ps *PubsubService) subTopic(topic string, subscriber string, ch chan interface{}) error {
	ps.lock.Lock()
	isNewTopic := false
	isNewSubscriberToTopic := false
	if _, ok := ps.topics[topic]; !ok {
		ps.topics[topic] = make(map[string]chan interface{})
		isNewTopic = true
	}
	if _, ok := ps.topics[topic][subscriber]; !ok {
		isNewSubscriberToTopic = true
	}
	ps.topics[topic][subscriber] = ch

	if _, ok := ps.subscribers[subscriber]; !ok {
		ps.subscribers[subscriber] = make(map[string]bool)
	}
	ps.subscribers[subscriber][topic] = true
	ps.lock.Unlock()
	if isNewTopic {
		ps.publish(utils.RandomUUIDGenString(), "topiclist", []byte(topic))
	}
	if isNewSubscriberToTopic {
		go ps.notifyTopicSubWatchers(topic, subscriber)
	}
	return nil
}

func (ps *PubsubService) removeSubscriber(subscriber string) error {
	ps.lock.Lock()
	topics, ok := ps.subscribers[subscriber]
	if !ok {
		ps.lock.Unlock()
		return errors.New("subscriber not exist")
	}
	for topic := range topics {
		delete(ps.topics[topic], subscriber)
	}
	delete(ps.subscribers, subscriber)
	ps.lock.Unlock()
	return nil
}

func (ps *PubsubService) unsubTopic(topic string, subscriber string) error {
	ps.lock.Lock()
	if _, ok := ps.topics[topic]; !ok {
		ps.lock.Unlock()
		return errors.New("topic not exist")
	}
	delete(ps.topics[topic], subscriber)

	if _, ok := ps.subscribers[subscriber]; !ok {
		panic("topic exist but subscriber not exist")
	}
	delete(ps.subscribers[subscriber], topic)
	ps.subscribers[subscriber][topic] = true
	ps.lock.Unlock()
	return nil
}

func (ps *PubsubService) publishSpecific(msgID string, topic string, data []byte, subscriber string) error {
	ps.lock.RLock()
	subsMap, ok := ps.topics[topic]
	if ok {
		ch, ok := subsMap[subscriber]
		if ok {
			go func(ch chan interface{}) {
				ch <- buildMsgPubsub(msgID, topic, data)
			}(ch)
		}
	}
	ps.lock.RUnlock()
	return nil
}

func (ps *PubsubService) publish(msgID string, topic string, data []byte) error {
	ps.lock.RLock()
	subsMap, ok := ps.topics[topic]
	if ok {
		for _, ch := range subsMap {
			go func(ch chan interface{}) {
				ch <- buildMsgPubsub(msgID, topic, data)
			}(ch)
		}
	} else {
		return errors.New("Topic not exist")
	}
	ps.lock.RUnlock()
	return nil
}

func (ps *PubsubService) getListSubscribersOfTopic(topic string) []byte {
	var list []string
	ps.lock.RLock()
	if _, ok := ps.topics[topic]; !ok {
		ps.lock.RUnlock()
		bufList, _ := json.Marshal(list)
		return bufList
	}
	for subscriber := range ps.topics[topic] {
		list = append(list, subscriber)
	}
	ps.lock.RUnlock()
	bufList, _ := json.Marshal(list)
	return bufList
}

func (ps *PubsubService) getListTopic() []byte {
	var list []string
	ps.lock.RLock()
	for topic := range ps.topics {
		list = append(list, topic)
	}
	ps.lock.RUnlock()
	bufList, _ := json.Marshal(list)
	return bufList
}

func (ps *PubsubService) notifyTopicSubWatchers(topic string, subscriber string) error {
	return ps.publish(utils.RandomUUIDGenString(), "wch"+topic, []byte(subscriber))
}

func (ps *PubsubService) addTopicSubWatcher(watcher string, topic string) error {
	ps.lock.RLock()
	subWatcherList, ok := ps.topicSubscriberWatchers[topic]
	if !ok {
		return errors.New("Topic not exist")
	}
	if utils.IndexOfString(watcher, subWatcherList) != -1 {
		return errors.New("Already watch topic")
	}
	ps.topicSubscriberWatchers[topic] = append(ps.topicSubscriberWatchers[topic], watcher)
	ps.lock.RUnlock()
	return nil
}

func buildMsgPubsub(msgID string, topic string, data []byte) []byte {
	builder := flatbuffers.NewBuilder(0)

	msgIDBytes := builder.CreateString(msgID)
	topicBytes := builder.CreateString(topic)
	dataBytes := builder.CreateByteVector(data)

	schema.MsgPubsubStart(builder)
	schema.MsgPubsubAddID(builder, msgIDBytes)
	schema.MsgPubsubAddTopic(builder, topicBytes)
	schema.MsgPubsubAddData(builder, dataBytes)
	responseBytes := schema.MsgPubsubEnd(builder)

	builder.Finish(responseBytes)

	return builder.FinishedBytes()
}
