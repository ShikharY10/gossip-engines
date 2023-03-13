package routes

// import (
// 	"fmt"
// 	"gbEngine/controllers"
// 	"gbEngine/protobuf"
// 	"log"

// 	"google.golang.org/protobuf/proto"
// )

// func HandleJob(m *controllers.Controller) {
// 	var trans protobuf.Transport

// 	for job := range m.RMQ.Msgs {
// 		fmt.Println("new job")
// 		err := proto.Unmarshal(job.Body, &trans)
// 		if err != nil {
// 			log.Println("[RHJERR1] : ", err.Error())
// 			continue
// 		}
// 		fmt.Println("job type: ", trans.Tp)
// 		if trans.Tp == 1 {
// 			m.TypeOne(&trans)
// 		} else if trans.Tp == 2 {
// 			m.TypeTwo(&trans)
// 		} else if trans.Tp == 3 {
// 			m.TypeThree(&trans)
// 		} else if trans.Tp == 4 {
// 			m.TypeFour(&trans)
// 		} else if trans.Tp == 5 {
// 			m.TypeFive(&trans)
// 		} else if trans.Tp == 7 {
// 			m.TypeSeven(&trans)
// 		} else if trans.Tp == 8 {
// 			m.TypeEight(&trans)
// 		} else if trans.Tp == 9 {
// 			m.TypeNine(&trans)
// 		} else if trans.Tp == 10 {
// 			m.TypeTen(&trans)
// 		} else if trans.Tp == 11 {
// 			m.TypeEleven(&trans)
// 		} else if trans.Tp == 12 {
// 			m.TypeTwelve(&trans)
// 		}
// 	}
// }
