package global

const JwtSecretKey = "the_world_s_most_secret_key_that_no_one_will_ever_hack"

// EmployeeFreedomTypeFired - id состояния, когда сотрудник уволен
const EmployeeFreedomTypeFired = 5

// TaskStateClose - id состояния, когда задача закрыта
const TaskStateClose = 4

// ClaimStateClose - id состояния, когда заявка закрыта
const ClaimStateClose = 6

// JwtClaimsCtxKey - ключ для хранения JwtClaims в Context
const JwtClaimsCtxKey = "JwtClaimsCtxKey"

// CurSessionCtxKey - ключ для сохранения сессии в Context
const CurSessionCtxKey = "CurSessionCtxKey"

// HeaderAuthenticationKey - ключ заголовка в котором содержится JWT токен
const HeaderAuthenticationKey = "Authentication"

// EmployeeAdminID - id администратора
const EmployeeAdminID = 1

// EmployeeLevelAccessAdmin - уровень доступа администратора
const EmployeeLevelAccessAdmin = 1

// EmployeePositionAdmin - роль администратора
const EmployeePositionAdmin = 3
