package trips

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/jackc/pgx/v4"
)

func Create(conn *pgx.Conn, req CreateTripRequest) (Trip, error) {

	tripId := fmt.Sprintf("TRP-%06X", rand.Intn(16777215))
	log.Printf("Creating trip with ID: %s\n", tripId)

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return Trip{}, err
	}

	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `INSERT INTO trips (id, traveller_id, destination, date_from, date_to, purpose, total_cost, notes)  VALUES ($1, $2, $3, $4, $5,$6,$7,$8)`,
		tripId, req.TravellerID, req.Destination, req.DateFrom, req.DateTo, req.Purpose, req.TotalCost, req.Notes)

	if err != nil {
		return Trip{}, err
	}

	for _, seg := range req.Segments {
		_, err = tx.Exec(context.Background(), `INSERT INTO trip_segments (trip_id, segment_type, ref_id, details, cost) VALUES ($1, $2, $3, $4, $5)`,
			tripId, seg.SegmentType, seg.RefID, seg.Details, seg.Cost)
		if err != nil {
			return Trip{}, err
		}
	}
	log.Printf("Segments count: %d\n", len(req.Segments))

	tx.Commit(context.Background())

	return Trip{
		ID:     tripId,
		Status: "DRAFT"}, nil

}
