package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID                 uuid.UUID `json:"id,omitempty"`
	BrowserFingerprint uint32    `json:"browser-fingerprint,omitempty"`
	CanvasFingerprint  uint32    `json:"canvas-fingerprint,omitempty"`
	CreatedAt          time.Time `json:"created-at"`
	DeviceLanguage     string    `json:"device-language,omitempty"`
	DeviceTimezone     string    `json:"device-timezone,omitempty"`
	EventName          string    `json:"event-name,omitempty"`
	FontFingerprint    uint32    `json:"font-fingerprint,omitempty"`
	Incognito          bool      `json:"incognito,omitempty"`
	IP                 string    `json:"ip,omitempty"`
	PeriodicWave       uint32    `json:"periodic-wave,omitempty"`
	Processed          bool      `json:"processed,omitempty"`
	ScreenResolution   string    `json:"screen-resolution,omitempty"`
	Session            string    `json:"session,omitempty"`
	Status             string    `json:"status,omitempty"`
	Storage            uint32    `json:"storage,omitempty"`
	UpdatedAt          time.Time `json:"updated-at"`
	UserAgent          string    `json:"user-agent,omitempty"`
	UserID             uint32    `json:"user-id,omitempty"`
	Utm                string    `json:"utm,omitempty"`
	WebglFingerprint   uint32    `json:"webgl-fingerprint,omitempty"`
}

type EventList []Event
