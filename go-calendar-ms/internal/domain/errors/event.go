package errors

type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	ErrDateBusy         = EventError("this time is busy by another owner's event")
	ErrIncorrectEndDate = EventError("end date is incorrect")
)
