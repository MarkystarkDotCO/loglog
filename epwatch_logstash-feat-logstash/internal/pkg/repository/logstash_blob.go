package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type LogstashBlobRepository struct {
	containerClient *azblob.ContainerClient
}

func NewLogstashBlobRepository(containerClient *azblob.ContainerClient) *LogstashBlobRepository {
	return &LogstashBlobRepository{containerClient: containerClient}
}

func (repo *LogstashBlobRepository) NewBlobClient(blobName string) (*azblob.BlockBlobClient, error) {
	blobClient, err := repo.containerClient.NewBlockBlobClient(blobName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return blobClient, nil
}
func (repo *LogstashBlobRepository) UploadBlob(ctx context.Context, blobName string, data []string) (*azblob.BlockBlobClient, error) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	blobClient, err := repo.NewBlobClient(blobName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	accessTier := azblob.AccessTierCool
	blockOptions := azblob.UploadOption{
		AccessTier: &accessTier,
	}
	_, err = blobClient.UploadBuffer(ctx, b, blockOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return blobClient, nil
}

func (repo *LogstashBlobRepository) DeleteBlob(ctx context.Context, blobClient *azblob.BlockBlobClient) error {
	_, err := blobClient.Delete(ctx, nil)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (repo *LogstashBlobRepository) ListBlob(ctx context.Context) ([]*azblob.BlobItemInternal, error) {
	pager := repo.containerClient.ListBlobsFlat(nil)
	var blobClients []*azblob.BlobItemInternal
	for pager.NextPage(ctx) {
		resp := pager.PageResponse()
		for _, v := range resp.ListBlobsFlatSegmentResponse.Segment.BlobItems {
			blobClients = append(blobClients, v)
		}
	}
	if err := pager.Err(); err != nil {
		log.Fatalf("Failure to list blobs: %+v", err)
	}
	return blobClients, nil
}
