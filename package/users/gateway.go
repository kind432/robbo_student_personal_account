package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	GetStudent(email, password string) (student *models.StudentCore, err error)
	SearchStudentByEmail(email string) (students []*models.StudentCore, err error)
	AddStudentToRobboGroup(studentId, robboGroupId, robboUnitId string) (err error)
	CreateStudent(student *models.StudentCore) (id string, err error)
	DeleteStudent(studentId uint) (err error)
	GetStudentById(studentId string) (student *models.StudentCore, err error)
	GetStudentsByRobboGroupId(robboGroupId string) (students []*models.StudentCore, err error)
	UpdateStudent(student *models.StudentCore) (err error)

	GetTeacher(email, password string) (teacher *models.TeacherCore, err error)
	GetAllTeachers() (teachers []*models.TeacherCore, err error)
	CreateTeacher(teacher *models.TeacherCore) (id string, err error)
	DeleteTeacher(teacherId uint) (err error)
	GetTeacherById(teacherId string) (teacher *models.TeacherCore, err error)
	UpdateTeacher(teacher *models.TeacherCore) (err error)

	GetParent(email, password string) (parent *models.ParentCore, err error)
	GetAllParent() (parents []*models.ParentCore, err error)
	GetParentById(parentId string) (parent *models.ParentCore, err error)
	CreateParent(parent *models.ParentCore) (id string, err error)
	UpdateParent(parent *models.ParentCore) (err error)
	DeleteParent(parentId uint) (err error)

	GetFreeListener(email, password string) (freeListener *models.FreeListenerCore, err error)
	GetFreeListenerById(freeListenerId string) (freeListener *models.FreeListenerCore, err error)
	CreateFreeListener(freeListener *models.FreeListenerCore) (id string, err error)
	DeleteFreeListener(freeListenerId uint) (err error)
	UpdateFreeListener(freeListener *models.FreeListenerCore) (err error)

	GetUnitAdmin(email, password string) (unitAdmin *models.UnitAdminCore, err error)
	GetAllUnitAdmins() (unitAdmins []*models.UnitAdminCore, err error)
	GetUnitAdminById(unitAdminId string) (unitAdmin *models.UnitAdminCore, err error)
	CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (id string, err error)
	DeleteUnitAdmin(superAdminId uint) (err error)
	UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (err error)
	SearchUnitAdminByEmail(email string) (unitAdmins []*models.UnitAdminCore, err error)

	GetSuperAdmin(email, password string) (superAdmin *models.SuperAdminCore, err error)
	GetSuperAdminById(superAdminId string) (superAdmin *models.SuperAdminCore, err error)
	UpdateSuperAdmin(superAdmin *models.SuperAdminCore) (err error)
	DeleteSuperAdmin(superAdminId uint) (err error)

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
