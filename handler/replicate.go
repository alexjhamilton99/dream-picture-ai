package handler

import (
	"dream-picture-ai/db"
	"dream-picture-ai/types"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const succeeded = "succeeded"

type ReplicateResp struct {
	Status string   `json:"status"`
	Output []string `json:"output"`
}

func HandleReplicateCallback(w http.ResponseWriter, r *http.Request) error {
	var resp ReplicateResp
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return err
	}
	if resp.Status != succeeded {
		return fmt.Errorf("Replicate callback responded with a non 2xx status: %s", resp.Status)
	}

	batchID, err := uuid.Parse(chi.URLParam(r, "batchID"))
	if err != nil {
		return fmt.Errorf("Replicate callback invalid batch ID: %s", err)
	}

	images, err := db.GetImagesByBatchID(batchID)
	if err != nil {
		return fmt.Errorf("Replicate callback failed to find image with batch ID %s: %s", batchID, err)
	}

	if len(images) != len(resp.Output) {
		return fmt.Errorf("Replicate callback unequal images compared to replicate output")
	}

	for i, imageURL := range resp.Output {
		images[i].Status = types.ImageStatusCompleted
		images[i].ImageLocation = imageURL
	}

	return nil
}
