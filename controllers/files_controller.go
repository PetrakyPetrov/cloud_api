package controllers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/petrakypetrov/cloud_clients_api/models"
)

// GetFiles ..
func GetFiles(c *gin.Context) {

	var file models.File

	UserID := c.Param("user_id")
	FolderID := c.Param("folder_id")
	file.UserID = UserID

	FID, err := strconv.ParseInt(FolderID, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", FID, FID)
	}
	file.FolderID = FID
	result, err := file.GetAllByUserID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    result,
	})
}

// UploadFiles ...
func UploadFiles(c *gin.Context) {

	var file models.File

	UserID := c.Param("user_id")
	FolderID := c.Param("folder_id")
	files, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	var buff bytes.Buffer
	fileSize, err := buff.ReadFrom(files)
	if _, err := files.Seek(0, 0); err != nil {
		log.Fatalln(err)
	}
	fileSizefloat := float64(fileSize)

	fileSizefloat = fileSizefloat / 1000

	var requestBody bytes.Buffer

	// Create a multipart writer
	multiPartWriter := multipart.NewWriter(&requestBody)

	// Initialize the file field
	fileWriter, err := multiPartWriter.CreateFormFile("file", header.Filename)
	if err != nil {
		log.Fatalln(err)
	}

	// Copy the actual file content to the field field's writer
	_, err = io.Copy(fileWriter, files)
	if err != nil {
		log.Fatalln(err)
	}

	// Populate other fields
	fieldWriter, err := multiPartWriter.CreateFormField("file2")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = fieldWriter.Write([]byte("Value"))
	if err != nil {
		log.Fatalln(err)
	}

	// We completed adding the file and the fields, let's close the multipart writer
	// So it writes the ending boundary
	multiPartWriter.Close()

	req, err := http.NewRequest("POST", "http://localhost:8081/account/"+UserID+"/files", &requestBody)
	if err != nil {
		log.Fatalln(err)
	}
	// We need to set the content type from the writer, it includes necessary boundary as well
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	// Do the request
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	FID, err := strconv.ParseInt(FolderID, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", FID, FID)
	}

	file.Name = header.Filename
	file.UserID = UserID

	file.FolderID = FID
	file.SizeKB = fileSizefloat

	result, err := file.Create()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    result,
	})
}

// DownloadFiles ..
func DownloadFiles(c *gin.Context) {

	UserID := c.Param("user_id")
	Name := c.Param("name")

	response, err := http.Get("http://localhost:8081/account/" + UserID + "/files?name=" + Name)
	if err != nil {
		log.Fatalln(err)
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	cd := response.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(cd)
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename=` + params["filename"],
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

// DeleteFiles ..
func DeleteFiles(c *gin.Context) {
	var file models.File

	UserID := c.Param("user_id")
	FileID := c.Param("file_id")

	FID, err := strconv.ParseInt(FileID, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", FID, FID)
	}

	file.UserID = UserID
	file.ID = FID

	result, err := file.Delete()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    result,
	})
}
