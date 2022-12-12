package usecase

import (
	"crypto/sha1"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"strconv"
)

type UsersUseCaseImpl struct {
	usersGateway      users.Gateway
	robboGroupGateway robboGroup.Gateway
}

func (p *UsersUseCaseImpl) GetStudentsByRobboUnitId(robboUnitId string) (students []*models.StudentCore, err error) {
	return p.usersGateway.GetStudentsByRobboUnitId(robboUnitId)
}

func (p *UsersUseCaseImpl) GetStudentsByTeacherId(teacherId string) (students []*models.StudentCore, err error) {
	relations, getRelationErr := p.usersGateway.GetStudentTeacherRelationsByTeacherId(teacherId)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}
	for _, relation := range relations {
		student, getTeacherErr := p.usersGateway.GetStudentById(relation.StudentId)
		if getTeacherErr != nil {
			err = getTeacherErr
			return
		}
		students = append(students, student)
	}
	return
}

func (p *UsersUseCaseImpl) GetTeacherByRobboGroupId(robboGroupId string) (teachers []*models.TeacherCore, err error) {
	relations, getRelationErr := p.robboGroupGateway.GetRelationByRobboGroupId(robboGroupId)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}
	for _, relation := range relations {
		teacher, getTeacherErr := p.usersGateway.GetTeacherById(relation.TeacherId)
		if getTeacherErr != nil {
			err = getTeacherErr
			return
		}
		teachers = append(teachers, &teacher)
	}
	return
}

func (p *UsersUseCaseImpl) GetStudentsByRobboGroupId(robboGroupId string) (students []*models.StudentCore, err error) {
	return p.usersGateway.GetStudentsByRobboGroupId(robboGroupId)
}

type UsersUseCaseModule struct {
	fx.Out
	users.UseCase
}

func SetupUsersUseCase(usersGateway users.Gateway, robboGroupGateway robboGroup.Gateway) UsersUseCaseModule {
	return UsersUseCaseModule{
		UseCase: &UsersUseCaseImpl{
			usersGateway:      usersGateway,
			robboGroupGateway: robboGroupGateway,
		},
	}
}

//func (p *UsersUseCaseImpl) GetStudent(email, password string) (student *models.StudentCore, err error) {
//	return p.usersGateway.GetStudent(email, password)
//}

func (p *UsersUseCaseImpl) GetStudentById(studentId string) (student *models.StudentCore, err error) {
	return p.usersGateway.GetStudentById(studentId)
}

func (p *UsersUseCaseImpl) SearchStudentByEmail(email string, parentId string) (students []*models.StudentCore, err error) {
	emailCondition := "%" + email + "%"
	return p.usersGateway.SearchStudentsByEmail(emailCondition, parentId)
}

func (p *UsersUseCaseImpl) GetStudentByParentId(parentId string) (students []*models.StudentCore, err error) {
	relations, getRelationErr := p.usersGateway.GetRelationByParentId(parentId)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}
	for _, relation := range relations {
		student, getStudentErr := p.usersGateway.GetStudentById(relation.ChildId)
		if getStudentErr != nil {
			err = getStudentErr
			return
		}
		students = append(students, student)
	}
	return
}

func (p *UsersUseCaseImpl) CreateStudent(student *models.StudentCore, parentId string) (id string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(student.Password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	student.Password = passwordHash
	id, err = p.usersGateway.CreateStudent(student)
	if err != nil {
		return
	}
	relation := &models.ChildrenOfParentCore{
		ChildId:  id,
		ParentId: parentId,
	}
	p.usersGateway.CreateRelation(relation)
	return
}

func (p *UsersUseCaseImpl) DeleteStudent(studentId uint) (err error) {
	if err = p.usersGateway.DeleteStudent(studentId); err != nil {
		return
	}
	if err = p.usersGateway.DeleteRelationByChildrenId(strconv.Itoa(int(studentId))); err != nil {
		return
	}
	if err = p.usersGateway.DeleteStudentTeacherRelationByStudentId(strconv.Itoa(int(studentId))); err != nil {
		return
	}
	return
}

func (p *UsersUseCaseImpl) UpdateStudent(student *models.StudentCore) (err error) {
	err = p.usersGateway.UpdateStudent(student)
	if err != nil {
		log.Println("Error update student")
		return
	}
	return
}

func (p *UsersUseCaseImpl) AddStudentToRobboGroup(studentId string, robboGroupId string, robboUnitId string) (err error) {
	if err = p.usersGateway.AddStudentToRobboGroup(studentId, robboGroupId, robboUnitId); err != nil {
		return
	}
	teachersRobboGroupsRelations, err := p.robboGroupGateway.GetRelationByRobboGroupId(robboGroupId)
	if err != nil {
		return
	}
	for _, relation := range teachersRobboGroupsRelations {
		relationCore := &models.StudentsOfTeacherCore{
			StudentId: studentId,
			TeacherId: relation.TeacherId,
		}
		if err = p.usersGateway.CreateStudentTeacherRelation(relationCore); err != nil {
			return
		}
	}
	return
}

func (p *UsersUseCaseImpl) GetTeacherById(teacherId string) (teacher models.TeacherCore, err error) {
	return p.usersGateway.GetTeacherById(teacherId)
}

func (p *UsersUseCaseImpl) GetTeachersByStudentId(studentId string) (teachers []*models.TeacherCore, err error) {
	relations, getRelationErr := p.usersGateway.GetStudentTeacherRelationsByStudentId(studentId)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}
	for _, relation := range relations {
		teacher, getTeacherErr := p.usersGateway.GetTeacherById(relation.TeacherId)
		if getTeacherErr != nil {
			err = getTeacherErr
			return
		}
		teachers = append(teachers, &teacher)
	}
	return
}

//func (p *UsersUseCaseImpl) GetTeacher(email, password string) (teacher *models.TeacherCore, err error) {
//	return p.usersGateway.GetTeacher(email, password)
//}

func (p *UsersUseCaseImpl) GetAllTeachers() (teachers []models.TeacherCore, err error) {
	return p.usersGateway.GetAllTeachers()
}

func (p *UsersUseCaseImpl) UpdateTeacher(teacher *models.TeacherCore) (err error) {
	err = p.usersGateway.UpdateTeacher(teacher)
	if err != nil {
		log.Println("Error update Teacher")
		return
	}
	return
}

func (p *UsersUseCaseImpl) CreateTeacher(teacher *models.TeacherCore) (id string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(teacher.Password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	teacher.Password = passwordHash
	return p.usersGateway.CreateTeacher(teacher)
}

func (p *UsersUseCaseImpl) DeleteTeacher(teacherId uint) (err error) {
	if err = p.usersGateway.DeleteTeacher(teacherId); err != nil {
		return
	}
	if err = p.usersGateway.DeleteStudentTeacherRelationByTeacherId(fmt.Sprintf("%v", teacherId)); err != nil {
		return
	}
	return
}

//func (p *UsersUseCaseImpl) GetParent(email, password string) (parent *models.ParentCore, err error) {
//	return p.usersGateway.GetParent(email, password)
//}

func (p *UsersUseCaseImpl) GetParentById(parentId string) (parent *models.ParentCore, err error) {
	return p.usersGateway.GetParentById(parentId)
}

func (p *UsersUseCaseImpl) GetAllParent() (parents []*models.ParentCore, err error) {
	parents, err = p.usersGateway.GetAllParent()
	return
}

func (p *UsersUseCaseImpl) CreateParent(parent *models.ParentCore) (id string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(parent.Password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	parent.Password = passwordHash
	return p.usersGateway.CreateParent(parent)
}

func (p *UsersUseCaseImpl) DeleteParent(parentId uint) (err error) {
	relations, getRelationsErr := p.usersGateway.GetRelationByParentId(strconv.Itoa(int(parentId)))
	if getRelationsErr != nil {
		return getRelationsErr
	}

	for _, relation := range relations {
		studentId, _ := strconv.ParseUint(relation.ChildId, 10, 64)
		deleteStudentErr := p.usersGateway.DeleteStudent(uint(studentId))
		if deleteStudentErr != nil {
			return deleteStudentErr
		}
	}
	deleteRelationErr := p.usersGateway.DeleteRelationByParentId(strconv.Itoa(int(parentId)))
	if deleteRelationErr != nil {
		return deleteRelationErr
	}
	return p.usersGateway.DeleteParent(parentId)
}

func (p *UsersUseCaseImpl) UpdateParent(parent *models.ParentCore) (err error) {
	err = p.usersGateway.UpdateParent(parent)
	if err != nil {
		log.Println("Error update Parent")
		return
	}
	return
}

//func (p *UsersUseCaseImpl) GetFreeListener(email, password string) (freeListener *models.FreeListenerCore, err error) {
//	return p.usersGateway.GetFreeListener(email, password)
//}

func (p *UsersUseCaseImpl) GetFreeListenerById(freeListenerId string) (freeListener *models.FreeListenerCore, err error) {
	return p.usersGateway.GetFreeListenerById(freeListenerId)
}

func (p *UsersUseCaseImpl) CreateFreeListener(freeListener *models.FreeListenerCore) (id string, err error) {
	return p.usersGateway.CreateFreeListener(freeListener)
}

func (p *UsersUseCaseImpl) DeleteFreeListener(freeListener uint) (err error) {
	return p.usersGateway.DeleteFreeListener(freeListener)
}

func (p *UsersUseCaseImpl) UpdateFreeListener(freeListener *models.FreeListenerCore) (err error) {
	err = p.usersGateway.UpdateFreeListener(freeListener)
	if err != nil {
		log.Println("Error update Parent")
		return
	}
	return
}

func (p *UsersUseCaseImpl) GetUnitAdminById(unitAdminId string) (unitAdmin *models.UnitAdminCore, err error) {
	return p.usersGateway.GetUnitAdminById(unitAdminId)
}

func (p *UsersUseCaseImpl) GetAllUnitAdmins() (unitAdmins []*models.UnitAdminCore, err error) {
	return p.usersGateway.GetAllUnitAdmins()
}

//func (p *UsersUseCaseImpl) GetUnitAdmin(email, password string) (unitAdmin *models.UnitAdminCore, err error) {
//	return p.usersGateway.GetUnitAdmin(email, password)
//}

func (p *UsersUseCaseImpl) UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (err error) {
	err = p.usersGateway.UpdateUnitAdmin(unitAdmin)
	if err != nil {
		log.Println("Error update Unit Admin")
		return
	}
	return
}

func (p *UsersUseCaseImpl) CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (id string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(unitAdmin.Password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	passwordHash := fmt.Sprintf("%x", pwd.Sum(nil))
	unitAdmin.Password = passwordHash
	return p.usersGateway.CreateUnitAdmin(unitAdmin)
}

func (p *UsersUseCaseImpl) DeleteUnitAdmin(unitAdminId uint) (err error) {
	return p.usersGateway.DeleteUnitAdmin(unitAdminId)
}

func (p *UsersUseCaseImpl) SearchUnitAdminByEmail(email string, robboUnitId string) (unitAdmins []*models.UnitAdminCore, err error) {
	emailCondition := "%" + email + "%"
	return p.usersGateway.SearchUnitAdminByEmail(emailCondition, robboUnitId)
}

func (p *UsersUseCaseImpl) GetUnitAdminByRobboUnitId(robboUnitId string) (unitAdmins []*models.UnitAdminCore, err error) {
	relations, getRelationErr := p.usersGateway.GetRelationByRobboUnitId(robboUnitId)
	if getRelationErr != nil {
		err = getRelationErr
		return
	}

	for _, relation := range relations {
		unitAdmin, getUnitAdminErr := p.usersGateway.GetUnitAdminById(relation.UnitAdminId)
		if getUnitAdminErr != nil {
			err = getRelationErr
			return
		}
		unitAdmins = append(unitAdmins, unitAdmin)
	}
	return
}

func (p *UsersUseCaseImpl) GetSuperAdminById(superAdminId string) (superAdmin *models.SuperAdminCore, err error) {
	return p.usersGateway.GetSuperAdminById(superAdminId)
}

func (p *UsersUseCaseImpl) UpdateSuperAdmin(superAdmin *models.SuperAdminCore) (err error) {
	err = p.usersGateway.UpdateSuperAdmin(superAdmin)
	if err != nil {
		log.Println("Error update Super Admin")
		return
	}
	return
}
func (p *UsersUseCaseImpl) DeleteSuperAdmin(superAdminId uint) (err error) {
	return p.usersGateway.DeleteSuperAdmin(superAdminId)
}

func (p *UsersUseCaseImpl) CreateRelation(parentId, childrenId string) (err error) {
	relationCore := &models.ChildrenOfParentCore{
		ChildId:  childrenId,
		ParentId: parentId,
	}
	return p.usersGateway.CreateRelation(relationCore)
}

func (p *UsersUseCaseImpl) SetNewUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error) {
	relationCore := &models.UnitAdminsRobboUnitsCore{
		UnitAdminId: unitAdminId,
		RobboUnitId: robboUnitId,
	}
	return p.usersGateway.SetUnitAdminForRobboUnit(relationCore)
}

func (p *UsersUseCaseImpl) DeleteUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error) {
	relationCore := &models.UnitAdminsRobboUnitsCore{
		UnitAdminId: unitAdminId,
		RobboUnitId: robboUnitId,
	}
	return p.usersGateway.DeleteUnitAdminForRobboUnit(relationCore)
}

func (p *UsersUseCaseImpl) CreateStudentTeacherRelation(studentId, teacherId string) (student *models.StudentCore, err error) {
	relationCore := &models.StudentsOfTeacherCore{
		StudentId: studentId,
		TeacherId: teacherId,
	}
	if createRelationErr := p.usersGateway.CreateStudentTeacherRelation(relationCore); createRelationErr != nil {
		err = createRelationErr
		return
	}
	student, getStudentErr := p.usersGateway.GetStudentById(studentId)
	if getStudentErr != nil {
		err = getStudentErr
		return
	}
	return
}

func (p *UsersUseCaseImpl) DeleteStudentTeacherRelation(studentId, teacherId string) (student *models.StudentCore, err error) {
	relationCore := &models.StudentsOfTeacherCore{
		StudentId: studentId,
		TeacherId: teacherId,
	}
	if deleteRelationErr := p.usersGateway.DeleteStudentTeacherRelation(relationCore); deleteRelationErr != nil {
		err = deleteRelationErr
		return
	}
	student, getStudentErr := p.usersGateway.GetStudentById(studentId)
	if getStudentErr != nil {
		err = getStudentErr
		return
	}
	return
}
