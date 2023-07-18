package link

import (
	"testing"

	"github.com/isu-kim/ebpf-mod/internal/testutils"
)

func TestHaveBPFLinkPerfEvent(t *testing.T) {
	testutils.CheckFeatureTest(t, haveBPFLinkPerfEvent)
}
