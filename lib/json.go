package lib

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2/utils"
)

var JSONMarshal utils.JSONMarshal = json.Marshal

var JSONUnmarshal utils.JSONUnmarshal = json.Unmarshal

func JSONEncoder(engine ...string) utils.JSONMarshal {
	return json.Marshal

}

func JSONDecoder(engine ...string) utils.JSONUnmarshal {
	return json.Unmarshal
}

// SetJSONEngine set default JSON encoder / decoder
func SetJSONEngine(engine ...string) {
	JSONMarshal = JSONEncoder(engine...)
	JSONUnmarshal = JSONDecoder(engine...)
}
