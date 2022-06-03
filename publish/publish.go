package publish

import "github.com/nats-io/nats.go"

type jsModel struct {
	js nats.JetStreamContext
}

type Models struct {
	JS jsModel
}

//NewModels returns a nats js pool
func NewModels(js nats.JetStreamContext) Models {
	return Models{
		JS: jsModel{
			js: js,
		},
	}
}

func (m *jsModel) GitPublish() {

}
