package lambda_handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	adapters_product "github.com/JoseLuis21/skeleton-lambda-ms/internal/core/product/adapters"
	service_product "github.com/JoseLuis21/skeleton-lambda-ms/internal/core/product/service"
	"github.com/aws/aws-lambda-go/events"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (lambdaHandler *Handler) HandlerRequest(ctx context.Context, sqsEvent events.SQSEvent) (sqsBatchResponse map[string]interface{}, err error) {
	//Init Log
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)

	logger.Info("MY ENV ", os.Getenv("TEST_ENV"), nil)

	// Batch Failures
	var batchItemFailures []map[string]interface{}

	for _, record := range sqsEvent.Records {
		var message MessagePath

		// Parse Message
		err := json.Unmarshal([]byte(record.Body), &message)
		if err != nil {
			logger.Error(err.Error(), "RecorId", record.MessageId)
			batchItemFailures = append(batchItemFailures, map[string]interface{}{"itemIdentifier": record.MessageId})
			continue
		}

		switch message.Path {
		case PATH:
			if message.Body != "" {
				// Init Adapter
				productRepo := adapters_product.NewProductRepository("mysql")
				productService := service_product.NewProductService(productRepo)

				// Using the service
				productService.Find("123")
				logger.Info("Test ID", "", message.Body)
			} else {
				batchItemFailures = append(batchItemFailures, map[string]interface{}{"itemIdentifier": record.MessageId})
			}
		}
	}

	// Return tasks with errors
	sqsBatchResponse = map[string]interface{}{
		"batchItemFailures": batchItemFailures,
	}

	logger.Warn("Tasks with errors", "tasks", sqsBatchResponse)
	return sqsBatchResponse, nil
}
