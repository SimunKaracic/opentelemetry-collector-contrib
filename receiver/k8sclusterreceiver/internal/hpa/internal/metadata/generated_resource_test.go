// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetK8sHpaName("k8s.hpa.name-val")
			rb.SetK8sHpaUID("k8s.hpa.uid-val")
			rb.SetK8sNamespaceName("k8s.namespace.name-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return 0

			switch test {
			case "default":
				assert.Equal(t, 3, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 3, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("k8s.hpa.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.hpa.name-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.hpa.uid")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.hpa.uid-val", val.Str())
			}
			val, ok = res.Attributes().Get("k8s.namespace.name")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "k8s.namespace.name-val", val.Str())
			}
		})
	}
}
