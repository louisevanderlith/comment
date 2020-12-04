package api

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"io/ioutil"
	"net/http"
	"strings"
)

func FetchCommentFor(web *http.Client, host string, ct commenttype.Enum, nodeKey hsk.Key) (hsk.Record, error) {
	typenum := strings.ToLower(commenttype.StringEnum(ct))
	url := fmt.Sprintf("%s/%s/%s", host, typenum, nodeKey.String())
	resp, err := web.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	result := records.NewRecord(&core.Message{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	return result, err
}
