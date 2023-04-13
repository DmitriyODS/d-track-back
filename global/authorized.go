package global

import (
	"fmt"
)

type Section byte

type Method byte

const (
	SectionEmployees  Section = 6
	SectionsCustomers         = 4
	SectionClaims             = 2
	SectionTasks              = 0
	SectionSelectLists
)

const (
	MethodView   Method = 1
	MethodEdit          = 2
	MethodCreate        = 3
)

func CheckLevelAccessWithClaims(levelAccess []byte, section Section, method Method) error {
	if section == SectionSelectLists {
		// списковая секция, списки получать могут все, ничего такого не случится
		return nil
	}

	switch method {
	case MethodView:
		// для просмотра нам нужно проверить один из двух бит, т.к. привелегии идут каскадно
		if (levelAccess[0]>>section)&3 > 0 {
			return nil
		}
	case MethodEdit:
		// для изменения нам нужно проверить только старший бит
		if (levelAccess[0]>>section)&2 > 0 {
			return nil
		}
	case MethodCreate:
		// для создания/удаления нам нужно проверить оба бита
		if (levelAccess[0]>>section)&3 == 3 {
			return nil
		}
	}

	return fmt.Errorf("нет запрашиваемого доступа")
}
