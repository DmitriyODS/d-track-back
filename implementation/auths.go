package implementation

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
)

func (bs *BasicService) AuthenticationByLogin(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	if !auth.ValidateFields() {
		return domain.Auth{}, global.IncorrectValidFormErr
	}

	employee, err := bs.rep.SelectUserByLoginPass(ctx, auth)
	if err != nil {
		log.Println("AuthenticationByLogin err:", err)
		return domain.Auth{}, global.InternalServerErr
	}

	if employee.ID == 0 {
		return domain.Auth{}, global.IncorrectLoginOrPassErr
	}

	auth.Password = ""
	auth.UserID = employee.ID

	// выдаём бессрочный токен, это упрощает проверку, но и снижает безопасность
	// в целях учебного проекта - норм, но в бизнес с таким лучше не ходить
	claims := &global.JwtClaims{
		UserID:           employee.ID,
		UserPositionID:   employee.Position.ID,
		UserLogin:        employee.Login,
		LevelAccess:      employee.LevelAccess.Access[0],
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// "секретный" ключ, тоже так себе лишь для учебного проекта
	auth.JWT, err = token.SignedString([]byte(global.JwtSecretKey))
	if err != nil {
		log.Println("AuthenticationByLogin err:", err)
		return auth, global.InternalServerErr
	}

	return auth, nil
}

func (bs *BasicService) CheckLevelAccessByEmployeeID(ctx context.Context, id uint64, levelAccess []byte) (bool, error) {
	return bs.rep.CheckLevelAccessByEmployeeID(ctx, id, levelAccess)
}
