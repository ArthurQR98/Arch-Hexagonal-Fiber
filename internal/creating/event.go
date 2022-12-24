package creating

import (
	"context"
	"errors"

	mooc "github.com/ArthurQR98/challenge_fiber/internal"
	"github.com/ArthurQR98/challenge_fiber/internal/increasing"
	"github.com/ArthurQR98/challenge_fiber/kit/event"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increasingService increasing.CourseCounterService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaserService increasing.CourseCounterService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(mooc.CourseCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}
	return e.increasingService.Increase(courseCreatedEvt.ID())
}
