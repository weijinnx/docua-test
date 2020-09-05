package redis

import "log"

// GetMessages from messages hash map in redis
func (rdb *Redis) GetMessages() map[string]string {
	messages, err := rdb.HGetAll(MessagesKeyName).Result()
	if err != nil {
		log.Printf("[redis] fail get messages hash map: %v", err)
		return nil
	}
	return messages
}

// AddMessage to messages hash map in redis
func (rdb *Redis) AddMessage(msg, t string) {
	err := rdb.HSet(MessagesKeyName, msg, t).Err()
	if err != nil {
		log.Printf("[redis] fail set message to messages hash map: %v", err)
	}
}

// RemoveMessage from messages hash map in redis
func (rdb *Redis) RemoveMessage(msg string) {
	err := rdb.HDel(MessagesKeyName, msg).Err()
	if err != nil {
		log.Printf("[redis] fail delete message from hash map: %v", err)
	}
}

// GetPrintedMessagesListLen returns len of printed messages list
func (rdb *Redis) GetPrintedMessagesListLen() int64 {
	len, err := rdb.LLen(PrintedMessagesKeyName).Result()
	if err != nil {
		log.Printf("[redis] fail get len of printed list: %v", err)
		return 0
	}
	return len
}

// GetPrintedMessages from printed messages list in redis
func (rdb *Redis) GetPrintedMessages() []string {
	empty := make([]string, 0)
	len := rdb.GetPrintedMessagesListLen()

	if len > 0 {
		messages, err := rdb.LRange(PrintedMessagesKeyName, 0, len-1).Result()
		if err != nil {
			log.Printf("[redis] fail get printed messages from list: %v", err)
			return empty
		}
		return messages
	}
	return empty
}

// AddToPrintedMessagesList adds messsage to printed messages list in redis
func (rdb *Redis) AddToPrintedMessagesList(msg string) {
	err := rdb.LPush(PrintedMessagesKeyName, msg).Err()
	if err != nil {
		log.Printf("[redis] fail add message to printed messages list: %v", err)
	}
}
