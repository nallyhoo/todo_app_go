package models

import (
    "encoding/json"
    "time"
)

type Todo struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
    StartTime   time.Time `json:"start_time"`
    EndTime     time.Time `json:"end_time"`
    Progress    float64   `json:"progress"`
}

// Custom unmarshaling for Todo to handle datetime-local format
func (t *Todo) UnmarshalJSON(data []byte) error {
    type Alias Todo // Avoid recursion in UnmarshalJSON
    aux := &struct {
        StartTime string `json:"start_time"`
        EndTime   string `json:"end_time"`
        *Alias
    }{
        Alias: (*Alias)(t),
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    // Parse start_time and end_time
    if aux.StartTime != "" {
        parsedTime, err := time.Parse("2006-01-02T15:04", aux.StartTime)
        if err != nil {
            return err
        }
        t.StartTime = parsedTime
    }
    if aux.EndTime != "" {
        parsedTime, err := time.Parse("2006-01-02T15:04", aux.EndTime)
        if err != nil {
            return err
        }
        t.EndTime = parsedTime
    }

    return nil
}