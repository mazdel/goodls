package main

const (
	defaultformat = `{
        "application/vnd.google-apps.form": "application/zip",
        "application/vnd.google-apps.document": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        "application/vnd.google-apps.drawing": "application/pdf",
        "application/vnd.google-apps.spreadsheet": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
        "application/vnd.google-apps.script": "application/vnd.google-apps.script+json",
        "application/vnd.google-apps.presentation": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
        "application/vnd.google-apps.site": "text/plain",
        "application/vnd.google-apps.jam": "application/pdf"
    }`

	extVsmime = `{
        "js":    "application/vnd.google-apps.script+json",
        "gs":    "application/vnd.google-apps.script+json",
        "gas":   "application/vnd.google-apps.script+json",
        "csv":   "text/csv",
        "htm":   "text/html",
        "html":  "text/html",
        "xbm":   "text/html",
        "shtml": "text/html",
        "shtm":  "text/html",
        "txt":   "text/plain",
        "text":  "text/plain",
        "json":  "application/json",
        "doc":   "application/msword",
        "xls":   "application/vnd.ms-excel",
        "ppt":   "application/vnd.ms-powerpoint",
        "docx":  "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        "xlsx":  "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
        "pptx":  "application/vnd.openxmlformats-officedocument.presentationml.presentation",
        "pdf":   "application/pdf",
        "ps":    "application/postscript",
        "eps":   "application/postscript",
        "gif":   "image/gif",
        "png":   "image/png",
        "svg":   "image/svg+xml",
        "jpg":   "image/jpeg",
        "jpeg":  "image/jpeg",
        "bmp":   "image/bmp",
        "ico":   "image/x-icon",
        "tif":   "image/tiff",
        "tiff":  "image/tiff",
        "mp3":   "audio/mp3",
        "wav":   "audio/wav",
        "mp4":   "video/mp4",
        "zip":   "application/zip"
    }`

	mimeVsEx = `{
        "application/vnd.google-apps.script": ".gs",
        "text/csv": ".csv",
        "text/html": ".html",
        "text/plain": ".txt",
        "application/json": ".json",
        "application/msword": ".doc",
        "application/vnd.ms-excel": ".xls",
        "application/vnd.ms-powerpoint": ".ppt",
        "application/vnd.google-apps.document": ".docx",
        "application/vnd.google-apps.spreadsheet": ".xlsx",
        "application/vnd.google-apps.presentation": ".pptx",
        "application/vnd.openxmlformats-officedocument.wordprocessingml.document": ".docx",
        "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": ".xlsx",
        "application/vnd.openxmlformats-officedocument.presentationml.presentation": ".pptx",
        "application/pdf": ".pdf",
        "application/postscript": ".ps",
        "image/gif": ".gif",
        "image/png": ".png",
        "image/svg+xml": ".svg",
        "image/jpeg": ".jpg",
        "image/bmp": ".bmp",
        "image/x-icon": ".ico",
        "image/tiff": ".tif",
        "audio/mp3": ".mp3",
        "audio/wav": ".wav",
        "video/mp4": ".mp4",
        "application/zip": ".zip"
    }`
)
