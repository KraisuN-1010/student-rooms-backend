package handlers

import "gofr.dev/pkg/gofr"

// GetRooms returns sample rooms
func GetRooms(c *gofr.Context) (interface{}, error) {
    rooms := []map[string]string{
        {"id": "1", "name": "Math Notes"},
        {"id": "2", "name": "Physics Doubts"},
    }
    return rooms, nil
}
