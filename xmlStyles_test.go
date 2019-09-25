package excelize

import (
	"encoding/xml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestXlsxStyleSheet_MarshalXML(t *testing.T) {
	relationships := xlsxStyleSheet{}

	b, err := xml.Marshal(relationships)
	require.NoError(t, err)

	b2 := replaceStyleRelationshipsNameSpaceBytes(b)
	t.Log(string(b))
	t.Log(string(b2))
	require.Equal(t, string(b), string(b2))
}
