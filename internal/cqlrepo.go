package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"

	"github.com/event-flux/domain"
)

const (
	filterEventName = "event_name"
	filterStartDate = "start_date"
	filterEndDate   = "end_date"
)

type Repo struct {
	session *gocql.Session
}

func (c Repo) GetByID(ctx context.Context, id string) (*domain.Event, error) {
	query := `
	SELECT id, browser_fingerprint, canvas_fingerprint, created_at, device_language,
		device_timezone, event_name, font_fingerprint, incognito, ip, periodic_wave, processed,
		screen_resolution, session, status, storage, updated_at, user_agent, user_id, utm, webgl_fingerprint
	FROM eventflux.events WHERE id = ? LIMIT 1
	`

	event := &domain.Event{}
	if err := c.session.Query(query, id).WithContext(ctx).Consistency(gocql.One).Scan(
		&event.ID, &event.BrowserFingerprint, &event.CanvasFingerprint, &event.CreatedAt,
		&event.DeviceLanguage, &event.DeviceTimezone, &event.EventName, &event.FontFingerprint,
		&event.Incognito, &event.IP, &event.PeriodicWave, &event.Processed, &event.ScreenResolution,
		&event.Session, &event.Status, &event.Storage, &event.UpdatedAt, &event.UserAgent,
		&event.UserID, &event.Utm, &event.WebglFingerprint,
	); err != nil {
		return nil, fmt.Errorf("repository: error fetching event by ID: %w", err)
	}

	return event, nil
}

func (c Repo) GetAll(ctx context.Context) ([]domain.Event, error) {
	query := `
	SELECT id, browser_fingerprint, canvas_fingerprint, created_at, device_language,
		device_timezone, event_name, font_fingerprint, incognito, ip, periodic_wave, processed,
		screen_resolution, session, status, storage, updated_at, user_agent, user_id, utm, webgl_fingerprint
	FROM eventflux.events
	`

	iter := c.session.Query(query).WithContext(ctx).Iter()
	var events domain.EventList
	event := domain.Event{}
	for iter.Scan(
		&event.ID, &event.BrowserFingerprint, &event.CanvasFingerprint, &event.CreatedAt,
		&event.DeviceLanguage, &event.DeviceTimezone, &event.EventName, &event.FontFingerprint,
		&event.Incognito, &event.IP, &event.PeriodicWave, &event.Processed, &event.ScreenResolution,
		&event.Session, &event.Status, &event.Storage, &event.UpdatedAt, &event.UserAgent,
		&event.UserID, &event.Utm, &event.WebglFingerprint,
	) {
		events = append(events, event)
		event = domain.Event{}
	}

	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("repository: error fetching all events: %w", err)
	}

	return events, nil
}

func (c Repo) GetByFilter(ctx context.Context, filters map[string]string) ([]domain.Event, error) {
	query := `SELECT id, browser_fingerprint, canvas_fingerprint, created_at, device_language,
				device_timezone, event_name, font_fingerprint, incognito, ip, periodic_wave, processed,
				screen_resolution, session, status, storage, updated_at, user_agent, user_id, utm, webgl_fingerprint
			FROM eventflux.events WHERE event_name = ?`

	values := make([]interface{}, 0, len(filters))
	eventName, eventNameExists := filters[filterEventName]

	if !eventNameExists {
		return nil, fmt.Errorf("repository: partitioning key %s is not provided", filterEventName)
	}
	values = append(values, eventName)

	startDate, startExists := filters[filterStartDate]
	endDate, endExists := filters[filterEndDate]
	if startExists && endExists {
		query += " AND created_at >= ? AND created_at <= ?"

		startDateTs, err := time.Parse("2006-01-02 15:04:05", startDate)
		if err != nil {
			return nil, fmt.Errorf("repository: error parsing timestamp: %w", err)
		}

		endDateTs, err := time.Parse("2006-01-02 15:04:05", endDate)
		if err != nil {
			return nil, fmt.Errorf("repository: error parsing timestamp: %w", err)
		}
		values = append(values, startDateTs, endDateTs)
	}

	iter := c.session.Query(query, values...).WithContext(ctx).Iter()
	var events domain.EventList
	event := domain.Event{}

	for iter.Scan(
		&event.ID, &event.BrowserFingerprint, &event.CanvasFingerprint, &event.CreatedAt,
		&event.DeviceLanguage, &event.DeviceTimezone, &event.EventName, &event.FontFingerprint,
		&event.Incognito, &event.IP, &event.PeriodicWave, &event.Processed, &event.ScreenResolution,
		&event.Session, &event.Status, &event.Storage, &event.UpdatedAt, &event.UserAgent,
		&event.UserID, &event.Utm, &event.WebglFingerprint,
	) {
		events = append(events, event)
		event = domain.Event{}
	}

	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("repository: error fetching events by filter: %w: query: %s", err, query)
	}

	return events, nil
}

func NewCQLRepo(session *gocql.Session) *Repo {
	return &Repo{session: session}
}
