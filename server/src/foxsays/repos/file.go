package repos

import (
	"time"
	"foxsays/models"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type file struct{ gridFile *mgo.GridFile }

func (f *file) Close() error                { return f.gridFile.Close() }
func (f *file) Read(p []byte) (int, error)  { return f.gridFile.Read(p) }
func (f *file) Write(p []byte) (int, error) { return f.gridFile.Write(p) }

func (f *file) Id() models.FileId {
	oid := f.gridFile.Id().(bson.ObjectId)
	return models.FileId(oid.Hex())
}

func (f *file) ContentType() string     { return f.gridFile.ContentType() }
func (f *file) SetContentType(s string) { f.gridFile.SetContentType(s) }

func (f *file) Name() string     { return f.gridFile.Name() }
func (f *file) SetName(s string) { f.gridFile.SetName(s) }

func (f *file) UploadDate() time.Time { return f.gridFile.UploadDate() }

func (f *file) Size() int64 { return f.gridFile.Size() }
