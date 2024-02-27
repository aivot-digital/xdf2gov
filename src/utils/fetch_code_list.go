package utils

import (
	"encoding/xml"
	"errors"
	"fmt"
	"goxdf2gov/src/models/xrepository_models"
	"io"
	"net/http"
	"strings"
)

func FetchCodeList(urn string) ([]map[string]string, error) {
	url := fmt.Sprintf("https://www.xrepository.de/api/xrepository/%s:technischerBestandteilGenericode", strings.TrimSpace(urn))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", res.StatusCode)
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("could not read response body")
	}

	var codeList xrepository_models.CodeList
	err = xml.Unmarshal(resBody, &codeList)
	if err != nil {
		return nil, err
	}

	var codes []map[string]string

	for _, row := range codeList.SimpleCodeList.Rows {
		keyRef := codeList.ColumnSet.Key.ColumnRef.Ref

		for _, value := range row.Values {
			if value.ColumnRef == keyRef {
				code := make(map[string]string)
				code["value"] = value.SimpleValue
				code["label"] = value.SimpleValue
				codes = append(codes, code)
				break
			}
		}
	}

	if len(codes) == 0 {
		return nil, errors.New("no codes found")
	}

	return codes, nil
}
