package main

// this code is closely following https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/protobuf_producer_example/protobuf_producer_example.go, except for enabling CacheSchemas

import (
    "fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
	"google.golang.org/protobuf/proto"

    pb "go_example/protos"
)

func serializeWithSchemaregistry(nMessages int) {
	registry, err := schemaregistry.NewClient(schemaregistry.NewConfig("http://schema-registry:8081"))
	if err != nil {
		fmt.Printf("Error creating schema registry client...\n")
		panic(err)
	}
	defer registry.Close()

	serializer, err := protobuf.NewSerializer(
		registry,
		serde.ValueSerde,
		&protobuf.SerializerConfig{
			SerializerConfig: *serde.NewSerializerConfig(),
			CacheSchemas: true,  // big difference for performance
		},
	)
	if err != nil {
		fmt.Printf("Error creating serializer...\n")
		panic(err)
	}
	defer serializer.Close()

	start := time.Now()
	topic := "mytopic"
    for i := 0; i < nMessages; i++ {
		message := pb.MyMessage{
			A: int32(i),
			B: int32(i),
		}

		payload, err := serializer.Serialize(topic, &message)
		if err != nil {
			fmt.Printf("Error serializing message...\n")
			panic(err)
		}
		// Pseudo-use of payload to prevent compiler optimizations
		_ = copy(payload, payload)
    }
	fmt.Printf("Protobuf+schemaregistry: Serialized %d messages in %v\n", nMessages, time.Since(start))
}

func serializeWithProtobuf(nMessages int) {
	start := time.Now()

    for i := 0; i < nMessages; i++ {
		message := pb.MyMessage{
			A: int32(i),
			B: int32(i),
		}

        payload, err := proto.Marshal(&message)
        if err != nil {
            fmt.Println("Error serializing protobuf message:", err)
            return
        }
		// Pseudo-use of payload to prevent compiler optimizations
		_ = copy(payload, payload)
    }

	fmt.Printf("Pure protobuf: Serialized %d messages in %v\n", nMessages, time.Since(start))
}

func main() {
	nMessages := 1_000_000
	serializeWithSchemaregistry(nMessages)
	serializeWithProtobuf(nMessages)
}
