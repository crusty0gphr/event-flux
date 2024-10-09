package domain

import (
	"time"
)

type Event struct {
	ID                 string    `json:"id,omitempty"`
	BrowserFingerprint uint64    `json:"browser-fingerprint,omitempty"`
	CanvasFingerprint  uint64    `json:"canvas-fingerprint,omitempty"`
	CreatedAt          time.Time `json:"created-at"`
	DeviceLanguage     string    `json:"device-language,omitempty"`
	DeviceTimezone     string    `json:"device-timezone,omitempty"`
	EventName          string    `json:"event-name,omitempty"`
	FontFingerprint    uint64    `json:"font-fingerprint,omitempty"`
	Incognito          bool      `json:"incognito,omitempty"`
	IP                 string    `json:"ip,omitempty"`
	PeriodicWave       uint64    `json:"periodic-wave,omitempty"`
	Processed          bool      `json:"processed,omitempty"`
	ScreenResolution   string    `json:"screen-resolution,omitempty"`
	Session            string    `json:"session,omitempty"`
	Status             string    `json:"status,omitempty"`
	Storage            uint64    `json:"storage,omitempty"`
	UpdatedAt          time.Time `json:"updated-at"`
	UserAgent          string    `json:"user-agent,omitempty"`
	UserID             uint64    `json:"user-id,omitempty"`
	Utm                string    `json:"utm,omitempty"`
	WebglFingerprint   uint64    `json:"webgl-fingerprint,omitempty"`
}

type EventList []Event
