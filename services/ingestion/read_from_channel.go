package ingestion

import (
	"fmt"
	"github.com/techrail/bark/constants"

	"github.com/techrail/bark/channels"
	"github.com/techrail/bark/models"
)

// InsertSingle and InsertMultiple functions have a common functionality - pushing the logs into the LogChannel.
// The channel capacities are checked beforehand. In case, the channels are full, that message is conveyed.
// Each log entry's values are validated and then sent to the LogChannel.
func InsertSingle(logEntry models.BarkLog) {
	if len(channels.LogChannel) > constants.ServerLogInsertionChannelCapacity-1 {
		fmt.Printf("E#1KDY0O - Channel is full. Cannot push. Log received: | %v\n", logEntry)
		return
	}
	logEntry, err := logEntry.ValidateForInsert()

	if err == nil {
		channels.LogChannel <- logEntry
	}
}

func InsertMultiple(logEntries []models.BarkLog) {
	var err error
	for _, logEntry := range logEntries {
		if len(channels.LogChannel) > constants.ServerLogInsertionChannelCapacity-1 {
			fmt.Printf("E#1KDZD2 - Channel is full. Cannot push. Log received: | %v\n", logEntry)
			return
		}

		logEntry, err = logEntry.ValidateForInsert()

		if err == nil {
			channels.LogChannel <- logEntry
		}
	}
}
