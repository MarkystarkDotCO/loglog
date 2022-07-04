package main

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/CMKL-PTTEP/epwatch_logstash/internal/pkg/repository"
)

func main() {
	// url := "https://cmklstorage.blob.core.windows.net/logstash-data?sp=racwdl&st=2022-06-23T09:07:24Z&se=2022-06-29T17:07:24Z&spr=https&sv=2021-06-08&sr=c&sig=LZdHQQle9CElSRFCNHsj4qdbF62AsEyE3Id492SgI%2Bg%3D"
	ctx := context.Background()
	url := "DefaultEndpointsProtocol=https;AccountName=cmklstorage;AccountKey=Ic+j8Fhp14qpxPFrOg0uCpuB+GrjI/WX8T4RhaFJq90xVATcKVJTTeKVnWaH1xxpIss2Azigzmuk+ASt1WMsiQ==;EndpointSuffix=core.windows.net"
	// // credential, err := azidentity.NewDefaultAzureCredential(nil)
	// // if err != nil {
	// // 	log.Fatal("Invalid credentials with error: " + err.Error())
	// // }

	serviceClient, err := azblob.NewServiceClientFromConnectionString(url, nil)

	// if err != nil {
	// 	log.Fatal("Invalid credentials with error: " + err.Error())
	// }
	// l := serviceClient.ListContainers(nil)
	// for _, i := range l.PageResponse().ListContainersSegmentResponse.ContainerItems {
	// 	fmt.Println(*i.Name)
	// }

	containerClient, err := serviceClient.NewContainerClient("trendmicro")
	if err != nil {
		log.Println(err)
	}

	// ap, _ := containerClient.NewAppendBlobClient("ap")
	// reader := strings.NewReader("Clear is better than clever")
	// f, err := os.Open("/tmp/123")
	// r := strings.NewReader("hello world.")
	// b := bytes.NewReader("hello world.")
	// r.Close()
	// b, _ := ioutil.ReadAll(r)

	// ap.AppendBlock(ctx, b, nil)
	// io.ReadSeekCloser

	// var a []interface{}
	// a = append(a, "b")
	// a = append(a, 2)
	repo := repository.NewLogstashBlobRepository(containerClient)
	blobs, _ := repo.ListBlob(ctx)
	for _, b := range blobs {
		blob, _ := repo.NewBlobClient(*b.Name)
		repo.DeleteBlob(ctx, blob)
	}
	// blobClient, _ := repo.NewBlobClient("e")
	// blobClient.
	// azblob.BlockBlobCommitBlockListOptions{}
	// repo.UploadBlob(ctx, "blobtest", a)
	// blobs, err := repo.ListBlob(ctx)

	// if err != nil {
	// 	log.Println(err)
	// }
	// // for _, blob := range blobs {
	// 	repo.DeleteBlob(ctx, blob)

	// }
}

// DefaultEndpointsProtocol=https;AccountName=cmklstorage;AccountKey=Ic+j8Fhp14qpxPFrOg0uCpuB+GrjI/WX8T4RhaFJq90xVATcKVJTTeKVnWaH1xxpIss2Azigzmuk+ASt1WMsiQ==;EndpointSuffix=core.windows.net
// Ic+j8Fhp14qpxPFrOg0uCpuB+GrjI/WX8T4RhaFJq90xVATcKVJTTeKVnWaH1xxpIss2Azigzmuk+ASt1WMsiQ==
// sp=r&st=2022-06-23T09:07:24Z&se=2022-06-29T17:07:24Z&spr=https&sv=2021-06-08&sr=c&sig=i7i%2Bf68TK%2Fi3HkpYmuGcI9MJAHSMDBBeetUNQWsEKX4%3D
// https://cmklstorage.blob.core.windows.net/logstash-data?sp=r&st=2022-06-23T09:07:24Z&se=2022-06-29T17:07:24Z&spr=https&sv=2021-06-08&sr=c&sig=i7i%2Bf68TK%2Fi3HkpYmuGcI9MJAHSMDBBeetUNQWsEKX4%3D
