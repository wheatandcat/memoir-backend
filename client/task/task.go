package task

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"

	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

type NotificationRequest struct {
	Token     []string `json:"token"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	URLScheme string   `json:"urlScheme"`
}

type HTTPTaskInterface interface {
	PushNotification(r NotificationRequest) (*taskspb.Task, error)
}

type HTTPTask struct {
	ProjectID  string
	LocationID string
	QueueID    string
	URL        string
}

func NewNotificationTask() HTTPTaskInterface {
	return &HTTPTask{
		ProjectID:  os.Getenv("GCP_PROJECT_ID"),
		LocationID: os.Getenv("GCP_LOCATION_ID"),
		QueueID:    os.Getenv("NOTIFICATION_QUEUE_ID"),
		URL:        os.Getenv("NOTIFICATION_URL"),
	}

}

func (t *HTTPTask) PushNotification(r NotificationRequest) (*taskspb.Task, error) {
	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, ce.CustomErrorWrap(err, "NewClient")
	}

	defer func() {
		err = client.Close()
	}()

	body, err := json.Marshal(r)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	// Build the Task queue path.
	queuePath := fmt.Sprintf("projects/%s/locations/%s/queues/%s", t.ProjectID, t.LocationID, t.QueueID)

	req := &taskspb.CreateTaskRequest{
		Parent: queuePath,
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        t.URL,
					Body:       body,
				},
			},
		},
	}

	createdTask, err := client.CreateTask(ctx, req)
	if err != nil {
		return nil, ce.CustomErrorWrap(err, "cloudtasks.CreateTask")
	}

	return createdTask, nil
}
