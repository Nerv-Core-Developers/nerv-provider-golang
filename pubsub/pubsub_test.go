package pubsub

import (
	"reflect"
	"sync"
	"testing"
)

func Test_pubsubService_SubTopic(t *testing.T) {
	type fields struct {
		topicSubscriberWatchers map[string][]string
		topics                  map[string]map[string]chan interface{}
		subscribers             map[string]map[string]bool
		RWMutex                 sync.RWMutex
	}
	type args struct {
		topic      string
		subscriber string
		ch         chan interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &pubsubService{
				topicSubscriberWatchers: tt.fields.topicSubscriberWatchers,
				topics:                  tt.fields.topics,
				subscribers:             tt.fields.subscribers,
				RWMutex:                 tt.fields.RWMutex,
			}
			if err := ps.SubTopic(tt.args.topic, tt.args.subscriber, tt.args.ch); (err != nil) != tt.wantErr {
				t.Errorf("pubsubService.SubTopic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pubsubService_RemoveSubscriber(t *testing.T) {
	type fields struct {
		topicSubscriberWatchers map[string][]string
		topics                  map[string]map[string]chan interface{}
		subscribers             map[string]map[string]bool
		RWMutex                 sync.RWMutex
	}
	type args struct {
		subscriber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &pubsubService{
				topicSubscriberWatchers: tt.fields.topicSubscriberWatchers,
				topics:                  tt.fields.topics,
				subscribers:             tt.fields.subscribers,
				RWMutex:                 tt.fields.RWMutex,
			}
			if err := ps.RemoveSubscriber(tt.args.subscriber); (err != nil) != tt.wantErr {
				t.Errorf("pubsubService.RemoveSubscriber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pubsubService_UnsubTopic(t *testing.T) {
	type fields struct {
		topicSubscriberWatchers map[string][]string
		topics                  map[string]map[string]chan interface{}
		subscribers             map[string]map[string]bool
		RWMutex                 sync.RWMutex
	}
	type args struct {
		topic      string
		subscriber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &pubsubService{
				topicSubscriberWatchers: tt.fields.topicSubscriberWatchers,
				topics:                  tt.fields.topics,
				subscribers:             tt.fields.subscribers,
				RWMutex:                 tt.fields.RWMutex,
			}
			if err := ps.UnsubTopic(tt.args.topic, tt.args.subscriber); (err != nil) != tt.wantErr {
				t.Errorf("pubsubService.UnsubTopic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pubsubService_PublishSpecific(t *testing.T) {
	type fields struct {
		topicSubscriberWatchers map[string][]string
		topics                  map[string]map[string]chan interface{}
		subscribers             map[string]map[string]bool
		RWMutex                 sync.RWMutex
	}
	type args struct {
		topic      string
		data       interface{}
		subscriber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &pubsubService{
				topicSubscriberWatchers: tt.fields.topicSubscriberWatchers,
				topics:                  tt.fields.topics,
				subscribers:             tt.fields.subscribers,
				RWMutex:                 tt.fields.RWMutex,
			}
			if err := ps.PublishSpecific(tt.args.topic, tt.args.data, tt.args.subscriber); (err != nil) != tt.wantErr {
				t.Errorf("pubsubService.PublishSpecific() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pubsubService_Publish(t *testing.T) {
	type fields struct {
		topicSubscriberWatchers map[string][]string
		topics                  map[string]map[string]chan interface{}
		subscribers             map[string]map[string]bool
		RWMutex                 sync.RWMutex
	}
	type args struct {
		topic string
		data  interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &pubsubService{
				topicSubscriberWatchers: tt.fields.topicSubscriberWatchers,
				topics:                  tt.fields.topics,
				subscribers:             tt.fields.subscribers,
				RWMutex:                 tt.fields.RWMutex,
			}
			if err := ps.Publish(tt.args.topic, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("pubsubService.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pubsubService_GetListSubscribersOfTopic(t *testing.T) {
	type fields struct {
		topicSubscriberWatchers map[string][]string
		topics                  map[string]map[string]chan interface{}
		subscribers             map[string]map[string]bool
		RWMutex                 sync.RWMutex
	}
	type args struct {
		topic string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &pubsubService{
				topicSubscriberWatchers: tt.fields.topicSubscriberWatchers,
				topics:                  tt.fields.topics,
				subscribers:             tt.fields.subscribers,
				RWMutex:                 tt.fields.RWMutex,
			}
			if got := ps.GetListSubscribersOfTopic(tt.args.topic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pubsubService.GetListSubscribersOfTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pubsubService_GetListTopic(t *testing.T) {
	type fields struct {
		topicSubscriberWatchers map[string][]string
		topics                  map[string]map[string]chan interface{}
		subscribers             map[string]map[string]bool
		RWMutex                 sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &pubsubService{
				topicSubscriberWatchers: tt.fields.topicSubscriberWatchers,
				topics:                  tt.fields.topics,
				subscribers:             tt.fields.subscribers,
				RWMutex:                 tt.fields.RWMutex,
			}
			if got := ps.GetListTopic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pubsubService.GetListTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pubsubService_NotifyTopicSubWatchers(t *testing.T) {
	type fields struct {
		topicSubscriberWatchers map[string][]string
		topics                  map[string]map[string]chan interface{}
		subscribers             map[string]map[string]bool
		RWMutex                 sync.RWMutex
	}
	type args struct {
		topic      string
		subscriber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &pubsubService{
				topicSubscriberWatchers: tt.fields.topicSubscriberWatchers,
				topics:                  tt.fields.topics,
				subscribers:             tt.fields.subscribers,
				RWMutex:                 tt.fields.RWMutex,
			}
			if err := ps.NotifyTopicSubWatchers(tt.args.topic, tt.args.subscriber); (err != nil) != tt.wantErr {
				t.Errorf("pubsubService.NotifyTopicSubWatchers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pubsubService_AddTopicSubWatcher(t *testing.T) {
	type fields struct {
		topicSubscriberWatchers map[string][]string
		topics                  map[string]map[string]chan interface{}
		subscribers             map[string]map[string]bool
		RWMutex                 sync.RWMutex
	}
	type args struct {
		watcher string
		topic   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &pubsubService{
				topicSubscriberWatchers: tt.fields.topicSubscriberWatchers,
				topics:                  tt.fields.topics,
				subscribers:             tt.fields.subscribers,
				RWMutex:                 tt.fields.RWMutex,
			}
			if err := ps.AddTopicSubWatcher(tt.args.watcher, tt.args.topic); (err != nil) != tt.wantErr {
				t.Errorf("pubsubService.AddTopicSubWatcher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
