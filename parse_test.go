// package procs_test
// onpar is a bad package, it requires old verison of cobra and makes code elimination not work

// import (
// 	"testing"

// 	"github.com/poy/onpar"
// 	. "github.com/poy/onpar/expect"
// 	. "github.com/poy/onpar/matchers"
// 	"github.com/fy0/procs"
// )

// func matchSplitCommand(t *testing.T, parts, expected []string) {
// 	for i, part := range parts {
// 		Expect(t, part).To(Equal(expected[i]))
// 	}
// }

// func TestSplitCommand(t *testing.T) {
// 	o := onpar.New()
// 	defer o.Run(t)

// 	o.Group("split with pipe", func() {
// 		o.BeforeEach(func(t *testing.T) (*testing.T, []string, []string) {
// 			parts := procs.SplitCommand("echo 'foo' | grep o")
// 			expected := []string{"echo", "foo", "|", "grep", "o"}
// 			return t, parts, expected
// 		})

// 		o.Spec("pass with a pipe", matchSplitCommand)
// 	})

// 	o.Group("replace with specific context", func() {
// 		o.BeforeEach(func(t *testing.T) (*testing.T, []string, []string) {
// 			parts := procs.SplitCommandEnv("echo ${FOO}", func(k string) string {
// 				return "bar"
// 			})

// 			expected := []string{"echo", "bar"}
// 			return t, parts, expected
// 		})

// 		o.Spec("expand values found in provided env", matchSplitCommand)
// 	})
// }
