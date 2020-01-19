package controller

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch"
)

// Patch apply json patch
func Patch(rawPatch []byte, dest interface{}) error {
	patch, err := jsonpatch.DecodePatch(rawPatch)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(dest)
	if err != nil {
		return err
	}
	modified, err := patch.Apply(jsonData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(modified, dest)
	if err != nil {
		return err
	}
	return nil
}
