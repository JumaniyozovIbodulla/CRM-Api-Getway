CURRENT_DIR := $(shell pwd)

gen-proto:
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/super_admins
	mkdir -p ${CURRENT_DIR}/genproto/user_service/super_admins
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/branches
	mkdir -p ${CURRENT_DIR}/genproto/user_service/branches
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/groups
	mkdir -p ${CURRENT_DIR}/genproto/user_service/groups
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/managers
	mkdir -p ${CURRENT_DIR}/genproto/user_service/managers
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/administrators
	mkdir -p ${CURRENT_DIR}/genproto/user_service/administrators
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/support_teachers
	mkdir -p ${CURRENT_DIR}/genproto/user_service/support_teachers
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/teachers
	mkdir -p ${CURRENT_DIR}/genproto/user_service/teachers
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/students
	mkdir -p ${CURRENT_DIR}/genproto/user_service/students
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/events
	mkdir -p ${CURRENT_DIR}/genproto/user_service/events
	sudo rm -rf ${CURRENT_DIR}/genproto/user_service/join_events
	mkdir -p ${CURRENT_DIR}/genproto/user_service/join_events
	sudo rm -rf ${CURRENT_DIR}/genproto/schedule_service/schedules
	mkdir -p ${CURRENT_DIR}/genproto/schedule_service/schedules
	sudo rm -rf ${CURRENT_DIR}/genproto/schedule_service/lessons
	mkdir -p ${CURRENT_DIR}/genproto/schedule_service/lessons
	sudo rm -rf ${CURRENT_DIR}/genproto/schedule_service/tasks
	mkdir -p ${CURRENT_DIR}/genproto/schedule_service/tasks
	sudo rm -rf ${CURRENT_DIR}/genproto/schedule_service/attendances
	mkdir -p ${CURRENT_DIR}/genproto/schedule_service/attendances
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/super_admins --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/super_admins --go-grpc_opt=paths=source_relative protos/user_protos/super_admins.proto
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/branches --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/branches --go-grpc_opt=paths=source_relative protos/user_protos/branches.proto
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/groups --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/groups --go-grpc_opt=paths=source_relative protos/user_protos/groups.proto
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/managers --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/managers --go-grpc_opt=paths=source_relative protos/user_protos/managers.proto
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/administrators --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/administrators --go-grpc_opt=paths=source_relative protos/user_protos/administrators.proto
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/support_teachers --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/support_teachers --go-grpc_opt=paths=source_relative protos/user_protos/support_teachers.proto
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/teachers --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/teachers --go-grpc_opt=paths=source_relative protos/user_protos/teachers.proto
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/students --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/students --go-grpc_opt=paths=source_relative protos/user_protos/students.proto
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/events --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/events --go-grpc_opt=paths=source_relative protos/user_protos/events.proto
	protoc --proto_path=protos/user_protos --go_out=${CURRENT_DIR}/genproto/user_service/join_events --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/user_service/join_events --go-grpc_opt=paths=source_relative protos/user_protos/join_events.proto
	protoc --proto_path=protos/schedule_protos --go_out=${CURRENT_DIR}/genproto/schedule_service/schedules --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/schedule_service/schedules --go-grpc_opt=paths=source_relative protos/schedule_protos/schedules.proto
	protoc --proto_path=protos/schedule_protos --go_out=${CURRENT_DIR}/genproto/schedule_service/lessons --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/schedule_service/lessons --go-grpc_opt=paths=source_relative protos/schedule_protos/lessons.proto
	protoc --proto_path=protos/schedule_protos --go_out=${CURRENT_DIR}/genproto/schedule_service/tasks --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/schedule_service/tasks --go-grpc_opt=paths=source_relative protos/schedule_protos/tasks.proto
	protoc --proto_path=protos/schedule_protos --go_out=${CURRENT_DIR}/genproto/schedule_service/attendances --go_opt=paths=source_relative --go-grpc_out=${CURRENT_DIR}/genproto/schedule_service/attendances --go-grpc_opt=paths=source_relative protos/schedule_protos/attendances.proto

swag_init:
	swag init -g api/router.go -o api/docs

run:
	go run cmd/main.go

git-push:
	git push origin main