package domain

import (
	"context"
	"mime/multipart"
	"time"
)

// Attachment ...
type Attachment struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	File      multipart.File `json:"-"`
	Type      string         `json:"file_type,omitempty"`
	Filename  string         `json:"-"`
	Size      int64          `json:"file_size,omitempty"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
}

// AttachmentUseCase represents attachments usecase contract
type AttachmentUseCase interface {
	CreateAttachment(ctx context.Context, attachment Attachment) (*Attachment, error)
	DownloadAttachment(ctx context.Context, fileName string) (string, error)
}

// AttachmentRepository represent the attachment's repository contract
type AttachmentRepository interface {
	// CreateAttachment(ctx context.Context, attachment Attachment) (*Attachment, error)
}

// AttachmentStorage represent the attachment's storage contract
type AttachmentStorage interface {
	CreateAttachment(ctx context.Context, attachment Attachment) error
	DownloadAttachment(ctx context.Context, fileName string) (string, error)
}
