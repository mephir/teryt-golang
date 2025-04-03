package datastruct

import "github.com/mephir/teryt-golang/internal/dataset/model"

type Datastruct interface {
	ToModel() (model.Model, error)
}
