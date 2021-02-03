package models

import (
	"time"
)

type AldSupportiveContactHeader struct {
	ID                int64            `json:"id"`
	UserID            int64            `json:"userID"`
	User              *User            `json:"user" pg:"joinFK:id"`
	Name              string           `json:"name"`
	ContactNumber     string           `json:"contactNumber"`
	AldRelationShipID int64            `json:"aldRelationShipId"`
	AldRelationShip   *AldRelationShip `json:"aldRelationShip" pg:"joinFK:id"`
	Version           int64            `json:"version"`
	IsActive          bool             `json:"isActive"`
	CreatedAt         time.Time        `json:"createdAt" sql:",default:now()"`
	UpdatedAt         time.Time        `json:"updatedAt" sql:",default:now()"`
}

type AldSupportiveContactLog struct {
	ID                           int64                       `json:"id"`
	Name                         string                      `json:"name"`
	ContactNumber                string                      `json:"contactNumber"`
	AldRelationShipID            int64                       `json:"aldRelationShipId"`
	AldRelationShip              *AldRelationShip            `json:"aldRelationShip" pg:"joinFK:id"`
	UserID                       int64                       `json:"userID"`
	User                         *User                       `json:"user" pg:"joinFK:id"`
	AldSupportiveContactHeaderID int64                       `json:"aldSupportiveContactHeaderId"`
	AldSupportiveContactHeader   *AldSupportiveContactHeader `json:"aldSupportiveContactHeader" pg:"joinFK:id"`
	Version                      int64                       `json:"version"`
	IsActive                     bool                        `json:"isActive"`
	CreatedAt                    time.Time                   `json:"createdAt" sql:",default:now()"`
	UpdatedAt                    time.Time                   `json:"updatedAt" sql:",default:now()"`
}
