package routes

import (
	"gbEngine/admin"
	"gbEngine/controllers/controller_v1"
	"gbEngine/schema"
	"log"

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

		log.Println("New Job | Type: ", payload.Type)

		switch payload.Type {
		case "000":
			ctrl.Test_API_ENGINE(&payload, &jobs.Body)
		case "001":
			ctrl.AckPayloadDelivery(&payload, &jobs.Body)
		case "002":
			ctrl.DeliverPendingPacket(&payload, &jobs.Body)
		case "011":
			ctrl.MakePartnerRequest(&payload, &jobs.Body)
		case "012":
			ctrl.MakePartnerResponse(&payload, &jobs.Body)
		case "013":
			ctrl.RemovePartnerNotifier(&payload, &jobs.Body)
		}
	}

}
