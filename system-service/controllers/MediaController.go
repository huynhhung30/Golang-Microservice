package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sytem_service/models"
	"sytem_service/utils/constants"
	"sytem_service/utils/functions"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

// Upload Files
func UploadFiles(c *gin.Context) {
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	bucketRegion := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("AWS_BUCKET")
	// Declare s3 objects
	creds := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		functions.ShowLog("bad credentials: ", err)
	}
	cfg := aws.NewConfig().WithRegion(bucketRegion).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)
	// Get data from form
	formData, err := c.MultipartForm()
	if err != nil {
		RES_ERROR_MSG(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	files := formData.File[models.AwsFormFiles]
	folder := formData.Value[models.AwsFormFolder]
	arrRes := []string{}
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			functions.ShowLog("err", err)
		}
		defer file.Close()

		fileInfo := files[i]
		size := fileInfo.Size
		buffer := make([]byte, size)
		fileType := http.DetectContentType(buffer)
		file.Read(buffer)

		if len(folder) > 0 {
			s3FilePath := folder[0] + "/" + fileInfo.Filename
			input := &s3.CreateMultipartUploadInput{
				Bucket:      aws.String(bucketName),
				Key:         aws.String(s3FilePath),
				ContentType: aws.String(fileType),
			}

			resp, err := svc.CreateMultipartUpload(input)
			if err != nil {
				RES_ERROR_MSG(c, http.StatusBadRequest, err.Error(), nil)
				return
			}

			var curr, partLength int64
			var remaining = size
			var completedParts []*s3.CompletedPart
			partNumber := 1
			for curr = 0; remaining != 0; curr += partLength {
				if remaining < models.AwsMaxPartSize {
					partLength = remaining
				} else {
					partLength = models.AwsMaxPartSize
				}
				completedPart, err := models.UploadPart(svc, resp, buffer[curr:curr+partLength], partNumber)
				if err != nil {
					fmt.Println(err.Error())
					err := models.AbortMultipartUpload(svc, resp)
					if err != nil {
						functions.ShowLog("err", err.Error())
					}
					RES_ERROR_MSG(c, http.StatusBadRequest, err.Error(), nil)
					return
				}
				remaining -= partLength
				partNumber++
				completedParts = append(completedParts, completedPart)
			}

			completeResponse, err := models.CompleteMultipartUpload(svc, resp, completedParts)
			if err != nil {
				RES_ERROR_MSG(c, http.StatusBadRequest, err.Error(), nil)
				return
			}
			imageUrl, _ := url.QueryUnescape(*completeResponse.Location)
			imageUrl = strings.Replace(imageUrl, constants.IMAGE_BASE_URL, constants.STATIC_WEBSITE_BASE_URL, -1)
			arrRes = append(arrRes, imageUrl)
		}
	}
	RES_SUCCESS(c, arrRes)
	return
}
