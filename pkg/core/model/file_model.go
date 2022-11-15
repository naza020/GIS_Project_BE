package model

type FileDownloadResponse struct {
	FilePath            string `json:"filePath"`
	DownloadFileName    string `json:"downloadFileName"`
	DeleteAfterDownload bool   `json:"deleteAfterDownload"`
}

type DownloadFile struct {
	URL      string `json:"url"`
	FileName string `json:"fileName"`
}
