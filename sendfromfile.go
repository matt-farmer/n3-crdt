// sendfromfile.go

package crdt

import (
	"io"
	"net/http"
	"os"
	"time"
	"github.com/pkg/errors"
)

//
// Sends a file of json objects through the crdt manager
//
func (crdtm *CRDTManager) SendFromFile(fname string) error {

	defer timeTrack(time.Now(), "SendFromFile() "+fname)

	// open the data file
	f, err := os.Open(fname)
	if err != nil {
		return errors.Wrap(err, "cannot open data file: ")
	}

	return crdtm.SendFromReader(f)

}

//
// sends an htttp request contianing json objects through the
// crdt manager
//
func (crdtm *CRDTManager) SendFromHTTPRequest(r *http.Request) error {

	err := crdtm.SendFromReader(r.Body)
	if err != nil {
		return errors.Wrap(err, "(sendfromfile.SendFromHTTPRequest) unknown:")
	}

	return err;
}

//
// sends the content of the given reader (assumed to be stream of json objects)
// through the crdt manager.
//
func (crdtm *CRDTManager) SendFromReader(r io.Reader) error {

	// err := runSendWithReader(crdtm.sdb, crdtm.swb, crdtm.UserId, crdtm.TopicName, crdtm.sc, r, crdtm.AuditLevel)
	err := runSendWithReader(crdtm.sdb, crdtm.UserId, crdtm.TopicName, crdtm.sc, r, crdtm.AuditLevel)
	if err != nil {
		return errors.Wrap(err, "(sendfromfile.SendFromReader) error ingesting data from reader:")
	}
	// ensure the writer finishes
	// crdtm.swb.Flush()
	// reinstate the writer
	// crdtm.swb = crdtm.sdb.NewWriteBatch()

	return err

}
