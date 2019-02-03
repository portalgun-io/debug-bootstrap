package mgr

import (
	"fmt"
	"strings"
	"testing"

	"github.com/threefoldtech/0-core/base/pm"

	"github.com/stretchr/testify/assert"
	"github.com/threefoldtech/0-core/base/stream"
)

func TestBuiltInProcess(t *testing.T) {
	work := func(ctx pm.Context) (interface{}, error) {
		return nil, nil
	}

	builtin := NewInternalProcess(work)(nil, nil)

	ch, err := builtin.Run()

	if ok := assert.NoError(t, err); !ok {
		t.Fatal()
	}

	if ok := assert.NotNil(t, ch); !ok {
		t.Fatal()
	}

	var msgs []*stream.Message

	for msg := range ch {
		msgs = append(msgs, msg)
	}

	if ok := assert.Len(t, msgs, 1); !ok {
		t.Fatal()
	}

	msg := msgs[0]

	if ok := assert.Equal(t, pm.LevelResultJSON, msg.Meta.Level()); !ok {
		t.Fatal()
	}

	if ok := assert.Equal(t, "null", msg.Message); !ok {
		t.Fatal()
	}
}

func TestBuiltInProcessData(t *testing.T) {
	work := func(ctx pm.Context) (interface{}, error) {
		return "some data", nil
	}

	builtin := NewInternalProcess(work)(nil, nil)

	ch, err := builtin.Run()

	if ok := assert.NoError(t, err); !ok {
		t.Fatal()
	}

	if ok := assert.NotNil(t, ch); !ok {
		t.Fatal()
	}

	var msgs []*stream.Message

	for msg := range ch {
		msgs = append(msgs, msg)
	}

	if ok := assert.Len(t, msgs, 1); !ok {
		t.Fatal()
	}

	msg := msgs[0]

	if ok := assert.Equal(t, pm.LevelResultJSON, msg.Meta.Level()); !ok {
		t.Fatal()
	}

	if ok := assert.True(t, msg.Meta.Is(stream.ExitSuccessFlag)); !ok {
		t.Fatal()
	}

	if ok := assert.Equal(t, `"some data"`, msg.Message); !ok {
		t.Fatal()
	}
}

func TestBuiltInProcessError(t *testing.T) {
	work := func(ctx pm.Context) (interface{}, error) {
		return nil, fmt.Errorf("error string")
	}

	builtin := NewInternalProcess(work)(nil, nil)

	ch, err := builtin.Run()

	if ok := assert.NoError(t, err); !ok {
		t.Fatal()
	}

	if ok := assert.NotNil(t, ch); !ok {
		t.Fatal()
	}

	var msgs []*stream.Message

	for msg := range ch {
		msgs = append(msgs, msg)
	}

	if ok := assert.Len(t, msgs, 1); !ok {
		t.Fatal()
	}

	msg := msgs[0]

	if ok := assert.Equal(t, pm.LevelResultJSON, msg.Meta.Level()); !ok {
		t.Fatal()
	}

	if ok := assert.True(t, msg.Meta.Is(stream.ExitErrorFlag)); !ok {
		t.Fatal()
	}

	if ok := assert.Equal(t, `"error string"`, msg.Message); !ok {
		t.Fatal()
	}
}

func TestBuiltInProcessRecover(t *testing.T) {
	work := func(ctx pm.Context) (interface{}, error) {
		panic("i paniced")
	}

	builtin := NewInternalProcess(work)(nil, nil)

	ch, err := builtin.Run()

	if ok := assert.NoError(t, err); !ok {
		t.Fatal()
	}

	if ok := assert.NotNil(t, ch); !ok {
		t.Fatal()
	}

	var msgs []*stream.Message

	for msg := range ch {
		msgs = append(msgs, msg)
	}

	if ok := assert.Len(t, msgs, 1); !ok {
		t.Fatal()
	}

	msg := msgs[0]

	if ok := assert.Equal(t, pm.LevelResultJSON, msg.Meta.Level()); !ok {
		t.Fatal()
	}

	if ok := assert.True(t, msg.Meta.Is(stream.ExitErrorFlag)); !ok {
		t.Fatal()
	}

	if ok := assert.True(t, strings.HasPrefix(msg.Message, `"i paniced`)); !ok {
		t.Fatal()
	}
}
