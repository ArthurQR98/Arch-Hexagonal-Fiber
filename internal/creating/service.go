package creating

import (
	"context"

	mooc "github.com/ArthurQR98/challenge_fiber/internal"
	"github.com/ArthurQR98/challenge_fiber/kit/event"
)

type CourseService struct {
	courseRepository mooc.CourseRepository
	eventBus         event.Bus
}

func NewCourseService(courseRepository mooc.CourseRepository, eventBus event.Bus) CourseService {
	return CourseService{
		courseRepository: courseRepository,
		eventBus:         eventBus,
	}
}

func (s CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}
	if err := s.courseRepository.Save(ctx, course); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, course.PullEvents())
}
