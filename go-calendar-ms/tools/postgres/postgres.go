package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/stdlib"
	"go.uber.org/zap"
	"go_learning_homework/go-calendar-ms/internal/domain/models"
	"time"
)

// a dummy
type DB struct{}

var d DB

var db *sql.DB
var lg *zap.Logger

func InitDB(ctx context.Context, zlg *zap.Logger, sqlUser, sqlPassword, sqlHost, sqlPort, dbName string) (DB, error) {
	// data source name
	dsn := fmt.Sprintf(`postgres://%v:%v@%v:%v/%v`, sqlUser, sqlPassword, sqlHost, sqlPort, dbName) //?sslmode=verify-full"

	cxn, err := sql.Open("pgx", dsn)
	if err != nil {
		lg.Fatal("failed to load driver: " + err.Error())
	}

	db = cxn
	if err := db.PingContext(ctx); err != nil {
		lg.Fatal("failed to ping db: " + err.Error())
	}
	lg = zlg

	return d, nil
}

// Close the DB connection
func CloseDBCxn() error {
	return db.Close()
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
		lg.Error("failed to load query: " + err.Error())
		return 0, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			lg.Error("failed to close rows: " + err.Error())
		}
	}()

	rows.Next()

	var id int64
	if err := rows.Scan(&id); err != nil {
		lg.Error("failed to scan a row: " + err.Error())
		return 0, err
	}

	return id, nil
}

func (d DB) DeleteEventById(ctx context.Context, id int64) error {
	query := `delete from events where id = $1;` // do we need a returning?

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		lg.Error("failed to load query: " + err.Error())
		return err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			lg.Error("failed to close rows: " + err.Error())
		}
	}()

	return nil
}

func (d DB) GetEventById(ctx context.Context, id int64) (*models.Event, error) {
	query := `select * from events where id = $1;`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		lg.Error("failed to load query: " + err.Error())
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			lg.Error("failed to close rows: " + err.Error())
		}
	}()

	rows.Next()

	// scanning and parsing
	{
		var id, owner int64
		var title, description string
		var startTime, endTime pgtype.Time
		var startDate, endDate pgtype.Date

		if err := rows.Scan(&id, &owner, &title, &description, &startDate, &startTime, &endDate, &endTime); err != nil {
			lg.Error("failed to scan a row: " + err.Error())
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
			lg.Error("failed to get a result after scan a row: " + err.Error())
			return nil, err
		}

		return &event, nil
	}
}

func (d DB) GetEventsByOwnerStartTime(ctx context.Context, owner int64, startTime time.Time) ([]models.Event, error) {
	query := `
		select * from events
		where owner = $1
		and ( cast(start_date as date) >= $2
		or case when cast(start_date as date) = $2 then cast(start_time as time without time zone) >= $3 end );`

	year, month, day := startTime.Date()
	startDate := fmt.Sprintf("%v-%v-%v", year, int(month), day)

	rows, err := db.QueryContext(ctx, query, owner, startDate, startTime.Format("15:04:05"))
	if err != nil {
		lg.Error("failed to load query: " + err.Error())
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			lg.Error("failed to close rows: " + err.Error())
		}
	}()

	events := make([]models.Event, 0)
	for rows.Next() {
		var id, owner int64
		var title, description string
		var startTime, endTime pgtype.Time
		var startDate, endDate pgtype.Date

		if err := rows.Scan(&id, &owner, &title, &description, &startDate, &startTime, &endDate, &endTime); err != nil {
			lg.Error("failed to scan a row: " + err.Error())
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
			lg.Error("failed to get a result after scan a row: " + err.Error())
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil // todo catch null answer
}

func (d DB) GetAllEventsFromNow(ctx context.Context) ([]models.Event, error) {
	query := `
		select * from events
		where ( cast(start_date as date) >= $1
		or case when cast(start_date as date) = $1 then cast(start_time as time without time zone) >= $2 end );`

	startTime := time.Now()
	year, month, day := startTime.Date()
	startDate := fmt.Sprintf("%v-%v-%v", year, int(month), day)

	rows, err := db.QueryContext(ctx, query, startDate, startTime.Format("15:04:05"))
	if err != nil {
		lg.Error("failed to load query: " + err.Error())
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			lg.Error("failed to close rows: " + err.Error())
		}
	}()

	events := make([]models.Event, 0)
	for rows.Next() {
		var id, owner int64
		var title, description string
		var startTime, endTime pgtype.Time
		var startDate, endDate pgtype.Date

		if err := rows.Scan(&id, &owner, &title, &description, &startDate, &startTime, &endDate, &endTime); err != nil {
			lg.Error("failed to scan a row: " + err.Error())
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
			lg.Error("failed to get a result after scan a row: " + err.Error())
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil // todo catch null answer
}
