package message

// I want to use rabbitmq to receive messages from a queue setting up the consumer in this file

// MessageReader is a struct that contains the connection and channel to the rabbitmq server
type MessageReader struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// NewMessageReader returns a new MessageReader struct
func NewMessageReader() *MessageReader {
	return &MessageReader{}
}

// Connect connects to the rabbitmq server
func (mr *MessageReader) Connect() error {
	// Connect to the rabbitmq server passing username, password, and host from environment variables to the Dial function
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	mr.Conn = conn

	return nil
}

// OpenChannel opens a channel to the rabbitmq server
func (mr *MessageReader) OpenChannel() error {
	// Open a channel to the rabbitmq server
	ch, err := mr.Conn.Channel()
	if err != nil {
		return err
	}

	mr.Channel = ch

	return nil
}

// ConsumeMessages consumes messages from the rabbitmq server
func (mr *MessageReader) ConsumeMessages() (<-chan amqp.Delivery, error) {
	// Declare a queue to consume from
	q, err := mr.Channel.QueueDeclare(
		"test_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, err
	}

	// Consume messages from the queue
	msgs, err := mr.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

// Close closes the connection to the rabbitmq server
func (mr *MessageReader) Close() error {
	// Close the connection to the rabbitmq server
	err := mr.Conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// CloseChannel closes the channel to the rabbitmq server
func (mr *MessageReader) CloseChannel() error {
	// Close the channel to the rabbitmq server
	err := mr.Channel.Close()
	if err != nil {
		return err
	}

	return nil
}

// ReadMessages reads messages from the rabbitmq server
func (mr *MessageReader) ReadMessages() (<-chan amqp.Delivery, error) {
	// Connect to the rabbitmq server
	err := mr.Connect()
	if err != nil {
		return nil, err
	}

	// Open a channel to the rabbitmq server
	err = mr.OpenChannel()
	if err != nil {
		return nil, err
	}

	// Consume messages from the rabbitmq server
	msgs, err := mr.ConsumeMessages()
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
