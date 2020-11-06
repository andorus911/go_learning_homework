package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/stdlib"
	"go_learning_homework/go-calendar-ms/internal/domain/models"
	"log"
	"time"
)

type DB struct {}
var d DB

var db *sql.DB

func InitDB(ctx context.Context) DB {
	// connection to DB
	dsn := "postgres://postgres:postgres@localhost:5432/gocalendar"//?sslmode=verify-full"
	// todo config
	connection, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to load driver: %v", err)
	}
	db = connection
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	return d
}

// Event

func (d DB) SaveEvent(ctx context.Context, event models.Event) (int64, error) {
	query := `insert into events(owner, title, descr, start_date, start_time, end_date, end_time) values ($1, $2, $3, $4, $5, $6, $7) returning id;`

	year, month, day := event.StartTime.Date()
	startDate := fmt.Sprintf("%v-%v-%v", year, month, day)
	year, month, day = event.EndTime.Date()
	endDate := fmt.Sprintf("%v-%v-%v", year, month, day)

	// is it ok send Time as pgTime?
	rows, err := db.QueryContext(ctx, query, event.Owner, event.Title, event.Description, startDate, event.StartTime, endDate, event.EndTime)
	if err != nil {
		log.Printf("failed to load query: %v", err)
		return 0, err
	}
	defer rows.Close() // TODO add error handle

	rows.Next()

	var id int64
	if err := rows.Scan(&id); err != nil {
		log.Printf("failed to scan a row: %v", err)
		return 0, nil
	}

	return id, nil
}

func (d DB) DeleteEventById(ctx context.Context, id int64) error {
	query := `delete from events where id = $1` // do we need a returning?

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		log.Printf("failed to load query: %v", err)
		return err
	}
	defer rows.Close() // TODO add error handle

	return nil
}

func (d DB) GetEventById(ctx context.Context, id int64) (*models.Event, error) {
	query := `select * from events where id = $1`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		log.Printf("failed to load query: %v", err)
		return nil, err
	}
	defer rows.Close() // TODO add error handle
	rows.Next()

	// scanning and parsing
	{
		var id, owner int64
		var title, description string
		var startTime, endTime pgtype.Time
		var startDate, endDate pgtype.Date

		if err := rows.Scan(&id, &owner, &title, &description, &startDate, &startTime, &endDate, &endTime); err != nil {
			log.Printf("failed to scan a row: %v", err)
			return nil, err
		}

		event := models.Event{
			Id:          id,
			Title:       title,
			Description: description,
			Owner:       owner,
			StartTime:   startDate.Time.Add(time.Duration(startTime.Microseconds) * time.Microsecond),
			EndTime:     endDate.Time.Add(time.Duration(endTime.Microseconds) * time.Microsecond),
		}

		if err := rows.Err(); err != nil {
			log.Printf("failed to get a result after scan a row: %v", err)
			return nil, err
		}

		return &event, nil
	}
}

func (d DB) GetEventsByOwnerStartTime(ctx context.Context, owner int64, startTime time.Time) ([]models.Event, error) {
	query := `select * from events where owner = $1, start_date = $2, start_time = $3`

	year, month, day := startTime.Date()
	startDate := fmt.Sprintf("%v-%v-%v", year, month, day)

	rows, err := db.QueryContext(ctx, query, owner, startDate, startTime)
	if err != nil {
		log.Printf("failed to load query: %v", err)
		return nil, err
	}
	defer rows.Close() // TODO add error handle

	events := make([]models.Event, 1)
	for rows.Next() {
		var id, owner int64
		var title, description string
		var startTime, endTime pgtype.Time
		var startDate, endDate pgtype.Date

		if err := rows.Scan(&id, &owner, &title, &description, &startDate, &startTime, &endDate, &endTime); err != nil {
			log.Printf("failed to scan a row: %v", err)
			return nil, err
		}

		fmt.Printf("%v %v %v %v %v %v \n", id, owner, title, description, startDate.Time.Add(time.Duration(startTime.Microseconds) * time.Microsecond), endDate.Time.Add(time.Duration(endTime.Microseconds) * time.Microsecond))

		event := models.Event{
			Id:          id,
			Title:       title,
			Description: description,
			Owner:       owner,
			StartTime:   startDate.Time.Add(time.Duration(startTime.Microseconds) * time.Microsecond),
			EndTime:     endDate.Time.Add(time.Duration(endTime.Microseconds) * time.Microsecond),
		}

		if err := rows.Err(); err != nil {
			log.Printf("failed to get a result after scan a row: %v", err)
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}