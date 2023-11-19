package util

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"io"
	"ms-go-initial/model"
	"net/http"
)

func CloseConnectionSafely(closable io.Closer) {
	if closable == nil {
		return
	}
	if err := closable.Close(); err != nil {
		log.Warnf("ActionLog.CloseConnectionSafely.warn Failed to close %T connection: %s", closable, err)
	}
}
func SliceContains[V comparable](search V, array []V) bool {
	for idx := range array {
		if search == array[idx] {
			return true
		}
	}
	return false
}

func Unmarshal(logger *log.Entry, cmd *redis.StringCmd, value interface{}) error {
	bytes, err := cmd.Bytes()
	if err != nil {
		logger.Errorf("ActionLog.unmarshalData.error - couldn't unmarshal data %v", err)
		return model.NewInternalServerError("Could not unmarshal model")

	}
	_ = json.Unmarshal(bytes, value)
	return nil
}
func Marshal[T any](logger *log.Entry, res []T) ([]byte, error) {
	marshalled, err := json.Marshal(res)
	if err != nil {
		logger.Errorf("ActionLog.marshal.error - couldn't marshall data %v", err)
		return nil, model.NewInternalServerError("Could not marshal model")
	}
	return marshalled, nil
}

func DecodeBody(r *http.Request, data interface{}) error {
	body, err := io.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	if err != nil {
		log.Error(err)
		return model.NewBadRequestError(err.Error())
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error(err)
		return model.NewBadRequestError(err.Error())
	}

	return nil
}

func HandleResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Add(model.ContentType, model.ApplicationJSON)
	w.WriteHeader(status)
	if res != nil {
		_ = json.NewEncoder(w).Encode(res)
	}
}

func GetHeaderMap(ctx context.Context) map[string]string {
	headerMap := make(map[string]string)
	for key, values := range ctx.Value(model.ContextHeader).(http.Header) {
		headerMap[key] = values[0]
	}
	return headerMap
}
