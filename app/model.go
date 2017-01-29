package app

import (
	"time"

	"github.com/acoshift/ds"
)

const (
	kindUser   = "User"
	kindRole   = "Role"
	kindCourse = "Course"
	kindEnroll = "Enroll"
)

// Course model
type Course struct {
	ds.StringIDModel
	ds.StampModel
	Title            string `datastore:",noindex"`
	ShortDescription string `datastore:",noindex"`
	Description      string `datastore:",noindex"` // Markdown
	Photo            string `datastore:",noindex"` // URL
	Owner            string
	Start            time.Time
	URL              string
	Type             CourseType
	Video            string `datastore:",noindex"` // Cover Video
	Price            float64
	DiscountedPrice  float64
	Options          CourseOption
	Contents         CourseContents `datastore:",noindex"`
	EnrollDetail     string         `datastore:",noindex"`

	EnrollCount int `datastore:"-"`
}

// Courses type
type Courses []*Course

// Len implements Sort interface
func (xs Courses) Len() int {
	return len(xs)
}

// Less implements Sort interface
func (xs Courses) Less(i, l int) bool {
	return xs[i].CreatedAt.Before(xs[l].CreatedAt)
}

// Swap implements Sort interface
func (xs Courses) Swap(i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

// CourseOption type
type CourseOption struct {
	Public     bool
	Enroll     bool `datastore:",noindex"`
	Attend     bool `datastore:",noindex"`
	Assignment bool `datastore:",noindex"`
	Discount   bool
}

// CourseContent type
type CourseContent struct {
	Title       string `datastore:",noindex"`
	Description string `datastore:",noindex"` // Markdown
	Video       string `datastore:",noindex"` // Youtube ID
	DownloadURL string `datastore:",noindex"` // Video download link
}

// CourseContents type
type CourseContents []CourseContent

// CourseType type
type CourseType string

// CourseType
const (
	CourseTypeLive  CourseType = "live"
	CourseTypeVideo CourseType = "video"
	CourseTypeEbook CourseType = "ebook"
)
