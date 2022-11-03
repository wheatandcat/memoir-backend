package mock_task

import (
	"os"

	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"

	"github.com/wheatandcat/memoir-backend/client/task"
)

type HTTPTask struct {
	ProjectID  string
	LocationID string
	QueueID    string
	URL        string
}

func NewNotificationTask() task.HTTPTaskInterface {
	return &HTTPTask{
		ProjectID:  os.Getenv("GCP_PROJECT_ID"),
		LocationID: os.Getenv("GCP_LOCATION_ID"),
		QueueID:    os.Getenv("NOTIFICATION_QUEUE_ID"),
		URL:        os.Getenv("NOTIFICATION_URL"),
	}

}

func (t *HTTPTask) PushNotification(r task.NotificationRequest) (*taskspb.Task, error) {
	return &taskspb.Task{}, nil
}
