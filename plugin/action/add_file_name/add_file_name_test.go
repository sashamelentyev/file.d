package add_file_name

import (
	"sync"
	"testing"

	"github.com/ozontech/file.d/pipeline"
	"github.com/ozontech/file.d/test"
	"github.com/stretchr/testify/assert"
)

func TestModify(t *testing.T) {
	config := test.NewConfig(&Config{Field: "file"}, nil)
	p, input, output := test.NewPipelineMock(test.NewActionPluginStaticInfo(factory, config, pipeline.MatchModeAnd, nil, false))
	wg := &sync.WaitGroup{}
	wg.Add(2)

	outEvents := make([]*pipeline.Event, 0)
	output.SetOutFn(func(e *pipeline.Event) {
		outEvents = append(outEvents, e)
		wg.Done()
	})

	input.In(0, "my_file", 0, []byte(`{"error":"info about error"}`))
	input.In(0, "my_file", 0, []byte(`{"file":"not_my_file"}`))

	wg.Wait()
	p.Stop()

	assert.Equal(t, 2, len(outEvents), "wrong out events count")
	assert.Equal(t, "my_file", outEvents[0].Root.Dig("file").AsString(), "wrong field value")
	assert.Equal(t, "my_file", outEvents[1].Root.Dig("file").AsString(), "wrong field value")
}