package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"websocket_p4/core/entities"
)

//	jpg,jpeg,gif,png,svg

func SetByCurlImage(ctx context.Context, file *multipart.FileHeader) *entities.UploadResponse {

	filePath := file.Filename
	fileExt := strings.ToLower(filepath.Ext(file.Filename))

	acceptedExts := []string{".jpg", ".jpeg", ".gif", ".png", ".svg"}
	accepted := false
	for _, ext := range acceptedExts {
		if fileExt == ext {
			accepted = true
			break
		}
	}

	if !accepted {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    7,
				Message: "Định dạng file không hợp lệ",
			},
		}
	}

	// Tạo một multipart form data
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// Thêm file vào form data
	fileWriter, err := bodyWriter.CreateFormFile("file", filePath)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}
	}
	fh, err := file.Open()
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}
	}
	defer fh.Close()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}
	}

	// Kết thúc form data
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// Gửi yêu cầu POST
	resp, err := http.Post("http://localhost:1234/manager/upload", contentType, bodyBuf)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}
	}
	defer resp.Body.Close()

	// Đọc phản hồi
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}
	}

	// Decode phản hồi thành struct
	var uploadResp entities.UploadResponse
	err = json.Unmarshal(respBody, &uploadResp)
	if err != nil {
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    4,
				Message: "Lỗi server",
			},
		}
	}
	switch uploadResp.Result.Code {
	case 1:
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    1,
				Message: "tệp rỗng 1 ",
			},
		}
	case 2:
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    2,
				Message: "Không có tệp được tải lên 1",
			},
		}
	case 3:
		return &entities.UploadResponse{
			Result: entities.Result{
				Code:    3,
				Message: "Tệp không hợp lệ 1",
			},
		}
	default:
		return &uploadResp
	}
}
