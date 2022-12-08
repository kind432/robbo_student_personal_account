package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	GetStudent(email, password string) (student *models.StudentCore, err error)
	SearchStudentsByEmail(email string, parentId string) (students []*models.StudentCore, err error)
	AddStudentToRobboGroup(studentId, robboGroupId, robboUnitId string) (err error)
	CreateStudent(student *models.StudentCore) (newStudent *models.StudentCore, err error)
	UpdateStudent(student *models.StudentCore) (studentUpdated *models.StudentCore, err error)
	DeleteStudent(studentId string) (err error)
	GetStudentById(studentId string) (student *models.StudentCore, err error)
	GetStudentsByRobboGroupId(robboGroupId string) (students []*models.StudentCore, err error)
	GetStudentsByRobboUnitId(robboUnitId string) (students []*models.StudentCore, err error)

	CreateTeacher(teacher *models.TeacherCore) (newTeacher models.TeacherCore, err error)
	UpdateTeacher(teacher *models.TeacherCore) (teacherUpdated models.TeacherCore, err error)
	DeleteTeacher(teacherId string) (err error)
	GetTeacher(email, password string) (teacher models.TeacherCore, err error)
	GetAllTeachers() (teachers []models.TeacherCore, err error)
	GetTeacherById(teacherId string) (teacher models.TeacherCore, err error)

	CreateParent(parent *models.ParentCore) (newParent *models.ParentCore, err error)
	UpdateParent(parent *models.ParentCore) (parentUpdated *models.ParentCore, err error)
	DeleteParent(parentId string) (err error)
	GetParent(email, password string) (parent *models.ParentCore, err error)
	GetAllParent() (parents []*models.ParentCore, err error)
	GetParentById(parentId string) (parent *models.ParentCore, err error)

	CreateFreeListener(freeListener *models.FreeListenerCore) (newFreeListener *models.FreeListenerCore, err error)
	UpdateFreeListener(freeListener *models.FreeListenerCore) (freeListenerUpdated *models.FreeListenerCore, err error)
	DeleteFreeListener(freeListenerId string) (err error)
	GetFreeListener(email, password string) (freeListener *models.FreeListenerCore, err error)
	GetFreeListenerById(freeListenerId string) (freeListener *models.FreeListenerCore, err error)

	CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (newUnitAdmin *models.UnitAdminCore, err error)
	UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (unitAdminUpdated *models.UnitAdminCore, err error)
	DeleteUnitAdmin(superAdminId string) (err error)
	GetUnitAdmin(email, password string) (unitAdmin *models.UnitAdminCore, err error)
	GetAllUnitAdmins() (unitAdmins []*models.UnitAdminCore, err error)
	GetUnitAdminById(unitAdminId string) (unitAdmin *models.UnitAdminCore, err error)
	SearchUnitAdminByEmail(email string, robboUnitId string) (unitAdmins []*models.UnitAdminCore, err error)

	UpdateSuperAdmin(superAdmin *models.SuperAdminCore) (superAdminUpdated *models.SuperAdminCore, err error)
	DeleteSuperAdmin(superAdminId string) (err error)
	GetSuperAdmin(email, password string) (superAdmin *models.SuperAdminCore, err error)
	GetSuperAdminById(superAdminId string) (superAdmin *models.SuperAdminCore, err error)

	CreateRelation(relation *models.ChildrenOfParentCore) (err error)
	DeleteRelationByParentId(parentId string) (err error)
	DeleteRelationByChildrenId(childrenId string) (err error)
	DeleteRelation(relation *models.ChildrenOfParentCore) (err error)
	GetRelationByParentId(parentId string) (relations []*models.ChildrenOfParentCore, err error)
	GetRelationByChildrenId(childrenId string) (relations []*models.ChildrenOfParentCore, err error)

	SetUnitAdminForRobboUnit(relation *models.UnitAdminsRobboUnitsCore) (err error)
	DeleteUnitAdminForRobboUnit(relation *models.UnitAdminsRobboUnitsCore) (err error)
	DeleteRelationByRobboUnitId(robboUnitId string) (err error)
	DeleteRelationByUnitAdminId(unitAdminId string) (err error)
	GetRelationByRobboUnitId(robboUnitId string) (relations []*models.UnitAdminsRobboUnitsCore, err error)
	GetRelationByUnitAdminId(unitAdminId string) (relations []*models.UnitAdminsRobboUnitsCore, err error)
}
