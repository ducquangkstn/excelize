package excelize

import (
	"encoding/xml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestXmlWorkBook_MarshalXML(t *testing.T) {
	workBook := xlsxWorkbook{
		FileVersion: &xlsxFileVersion{
			AppName: "aaaa",
		},
		Sheets: xlsxSheets{
			Sheet: []xlsxSheet{
				{
					Name: "name",
					ID:   "id",
				},
			},
		},
		PivotCaches: &xlsxPivotCaches{
			PivotCache: []xlsxPivotCache{
				{
					CacheID: 5,
					RID:     "6",
				},
			},
		},
	}

	b, err := xml.Marshal(workBook)
	require.NoError(t, err)

	b2 := replaceRelationshipsBytes(replaceRelationshipsNameSpaceBytes(b))
	t.Log(string(b2))
	require.Equal(t, string(b), string(b2))
}

func TestXlsxRelationships_MarshalXML(t *testing.T) {
	relationships := xlsxRelationships{
		Relationships: []xlsxRelationship{
			{
				ID:         "ID",
				Type:       "Type",
				Target:     "Target",
				TargetMode: "TargetMode",
			},
		},
	}

	b, err := xml.Marshal(relationships)
	require.NoError(t, err)

	b2 := replaceRelationshipsBytes(replaceWorkSheetsRelationshipsNameSpaceBytes(b))
	t.Log(string(b))
	require.Equal(t, string(b), string(b2))
}
