package config_test

import (
	"bytes"
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParseOptionsReader(t *testing.T) {
	// Data from https://github.com/VEuPathDB/dataset-handler-biom/blob/master/config.yml
	data := bytes.NewBuffer([]byte(`# For additional information about this config file see:
# https://github.com/VEuPathDB/util-user-dataset-handler-server
service-name: BIOM Format Handler
command:
  # "/app" comes from the working directory configured in the Dockerfile
  # The rest of the path references the directory structure of the
  # VEuPathDB/EuPathGalaxy github repository.
  executable: /opt/handler/bin/exportBiomToEuPathDB
  args:
    # 1. dataset name
    - <<ds-name>>
    # 2. dataset summary
    - <<ds-summary>>
    # 3. dataset description
    - <<ds-description>>
    # 4. dataset uploader id
    - <<ds-user-id>>
    # 5. output file (ignored)
    - html.html
    # 6. Dataset origin
    - <<ds-origin>>
    # 7. dataset file path
    - <<input-files[0]>>
file-types:
  - .biom`))
	log.SetLogger(log.ConfigureLogger())

	convey.Convey("ParseOptionsReader", t, func() {
		conf, err := config.ParseOptionsReader(data)

		convey.So(err, convey.ShouldBeNil)
		convey.So(conf.Commands().Args, convey.ShouldResemble, []string{
			"<<ds-name>>",
			"<<ds-summary>>",
			"<<ds-description>>",
			"<<ds-user-id>>",
			"html.html",
			"<<ds-origin>>",
			"<<input-files[0]>>",
		})
	})
}
