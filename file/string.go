package file

import "fmt"

func (f *File) String() string {
	return fmt.Sprintf("[%s:%3d]", f.PathName, f.LineNumber)
}
