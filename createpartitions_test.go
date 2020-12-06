package kafka

import (
	"context"
	"testing"
)

func TestClientCreatePartitions(t *testing.T) {

	client, shutdown := newLocalClient()
	defer shutdown()

	topic := makeTopic()
	createTopic(t, topic, 1)
	defer deleteTopic(t, topic)

	res, err := client.CreatePartitions(context.Background(), &CreatePartitionsRequest{
		Topics: []TopicPartitionsConfig{
			TopicPartitionsConfig{
				Name:  topic,
				Count: 2,
				TopicPartitionAssignments: []TopicPartitionAssignment{
					TopicPartitionAssignment{
						BrokerIDs: []int32{1},
					},
				},
			},
		},
		ValidateOnly: false,
	})

	if err != nil {
		t.Fatal(err)
	}

	if err := res.Errors[topic]; err != nil {
		t.Error(err)
	}
}
