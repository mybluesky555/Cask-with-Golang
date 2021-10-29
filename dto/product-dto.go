package dto

import "mime/multipart"

type ProductImport struct {
	ExcelFile *multipart.FileHeader `form:"excel_file"`
}

type DeleteIDs struct {
	IDs          []string `json:"ids"`
	SearchOption AllDataDTO
}
