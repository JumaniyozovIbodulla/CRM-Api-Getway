package grpc_client

import (
	at "crmapi/genproto/schedule_service/attendances"
	ls "crmapi/genproto/schedule_service/lessons"
	sc "crmapi/genproto/schedule_service/schedules"
	tk "crmapi/genproto/schedule_service/tasks"
	ad "crmapi/genproto/user_service/administrators"
	br "crmapi/genproto/user_service/branches"
	ev "crmapi/genproto/user_service/events"
	gr "crmapi/genproto/user_service/groups"
	ej "crmapi/genproto/user_service/join_events"
	mn "crmapi/genproto/user_service/managers"
	st "crmapi/genproto/user_service/students"
	sp "crmapi/genproto/user_service/super_admins"
	ss "crmapi/genproto/user_service/support_teachers"
	ts "crmapi/genproto/user_service/teachers"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"crmapi/config"
)

// GrpcClientI ...
type GrpcClientI interface {
	SuperAdmins() sp.SuperAdminServiceClient
	Branches() br.BranchesServiceClient
	Admins() ad.AdministratorsServiceClient
	Events() ev.EventServiceClient
	Groups() gr.GroupsServiceClient
	JoinEvens() ej.JoinEventServiceClient
	Managers() mn.ManagersServiceClient
	Students() st.StudentServiceClient
	SupportTeacher() ss.SupportTeacherServiceClient
	Teachers() ts.TeacherServiceClient
	Schedules() sc.ScheduleServiceClient
	Tasks() tk.TaskServiceClient
	Lessons() ls.LesssonServiceClient
	Attendances() at.AttendanceServiceClient
}

type GrpcClient struct {
	superAdminService sp.SuperAdminServiceClient
	bracnesService    br.BranchesServiceClient
	adminsService     ad.AdministratorsServiceClient // delete, getall
	
	eventsService         ev.EventServiceClient
	groupsService         gr.GroupsServiceClient
	joinEventsService     ej.JoinEventServiceClient
	managersService       mn.ManagersServiceClient
	studentsService       st.StudentServiceClient
	supportTeacherService ss.SupportTeacherServiceClient
	teachersService       ts.TeacherServiceClient
	schedulesService      sc.ScheduleServiceClient
	tasksService          tk.TaskServiceClient
	lessonsService        ls.LesssonServiceClient
	attendacesService     at.AttendanceServiceClient
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	superAdminService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	branchesService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	adminService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	eventsService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	groupsService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	joinEventsService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	managersService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	studentsService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	supportTeachersService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	teachersService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	scheduleService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.ScheduleServiceHost, cfg.ScheduleServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	taskService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.ScheduleServiceHost, cfg.ScheduleServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	lessonsService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.ScheduleServiceHost, cfg.ScheduleServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	attendancesService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.ScheduleServiceHost, cfg.ScheduleServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		superAdminService:     sp.NewSuperAdminServiceClient(superAdminService),
		bracnesService:        br.NewBranchesServiceClient(branchesService),
		adminsService:         ad.NewAdministratorsServiceClient(adminService),
		eventsService:         ev.NewEventServiceClient(eventsService),
		groupsService:         gr.NewGroupsServiceClient(groupsService),
		joinEventsService:     ej.NewJoinEventServiceClient(joinEventsService),
		managersService:       mn.NewManagersServiceClient(managersService),
		studentsService:       st.NewStudentServiceClient(studentsService),
		supportTeacherService: ss.NewSupportTeacherServiceClient(supportTeachersService),
		teachersService:       ts.NewTeacherServiceClient(teachersService),
		schedulesService:      sc.NewScheduleServiceClient(scheduleService),
		tasksService:          tk.NewTaskServiceClient(taskService),
		lessonsService:        ls.NewLesssonServiceClient(lessonsService),
		attendacesService:     at.NewAttendanceServiceClient(attendancesService),
	}, nil
}

func (g *GrpcClient) SuperAdmins() sp.SuperAdminServiceClient {
	return g.superAdminService
}

func (g *GrpcClient) Branches() br.BranchesServiceClient {
	return g.bracnesService
}

func (g *GrpcClient) Admins() ad.AdministratorsServiceClient {
	return g.adminsService
}

func (g *GrpcClient) Events() ev.EventServiceClient {
	return g.eventsService
}

func (g *GrpcClient) Groups() gr.GroupsServiceClient {
	return g.groupsService
}

func (g *GrpcClient) JoinEvens() ej.JoinEventServiceClient {
	return g.joinEventsService
}

func (g *GrpcClient) Managers() mn.ManagersServiceClient {
	return g.managersService
}

func (g *GrpcClient) Students() st.StudentServiceClient {
	return g.studentsService
}

func (g *GrpcClient) SupportTeacher() ss.SupportTeacherServiceClient {
	return g.supportTeacherService
}

func (g *GrpcClient) Teachers() ts.TeacherServiceClient {
	return g.teachersService
}

func (g *GrpcClient) Schedules() sc.ScheduleServiceClient {
	return g.schedulesService
}

func (g *GrpcClient) Tasks() tk.TaskServiceClient {
	return g.tasksService
}

func (g *GrpcClient) LessonsService() ls.LesssonServiceClient {
	return g.lessonsService
}

func (g *GrpcClient) AttendancesService() at.AttendanceServiceClient {
	return g.attendacesService
}
