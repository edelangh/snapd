package osutil_test

import (
	"os"
	"path/filepath"
	"gopkg.in/check.v1"
	"github.com/snapcore/snapd/osutil"
)

type ShredSuite struct{}
var _ = check.Suite(&ShredSuite{})

func (s *ShredSuite) TestShred(c *check.C) {
	path := filepath.Join(os.TempDir(), "randomfile")

	/* Check 3 sizes */
	for i := range []int64{0, 1, 1024} {
		f, err := os.Create(path)
		if err != nil {
			c.Fatal("test failed", err)
		}

		f.Truncate(int64(i))
		f.Close()

		err = osutil.Shred(path)
		c.Assert(err, check.IsNil)

		/* Check if file is deleted */
		_, err = os.Open(path)
		if !os.IsNotExist(err) {
			c.Fatal("File should be deleted", err)
		}
	}

	/* Check error case */
	err := osutil.Shred("")
	c.Assert(err, check.NotNil)
}

