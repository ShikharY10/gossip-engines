package routes

import (
	"gbEngine/admin"
	"gbEngine/controllers/controller_v1"
	"gbEngine/schema"

	"google.golang.org/protobuf/proto"
)

func HandleJob(ctrl *controller_v1.Controller, logger *admin.Logger) {
	var payload schema.Payload

	for jobs := range ctrl.Handler.Queue.Queue.Jobs {
		err := proto.Unmarshal(jobs.Body, &payload)
		if err != nil {
			logger.LogError(err)
			continue
		}

		switch payload.Type {
		case "011":
			ctrl.MakePartnerRequest(&payload, &jobs.Body)
		case "021":
			ctrl.MakePartnerResponse(&payload, &jobs.Body)
		case "031":
			ctrl.RemovePartnerNotifier(&payload, &jobs.Body)
		}
	}

}
