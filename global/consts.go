package global

const JwtSecretKey = "the_world_s_most_secret_key_that_no_one_will_ever_hack"

// EmployeeFreedomTypeFired - id состояния, когда сотрудник уволен
const EmployeeFreedomTypeFired = 5

// JwtClaimsCtxKey - ключ для хранения JwtClaims в Context
const JwtClaimsCtxKey = "JwtClaimsCtxKey"

// CurSessionCtxKey - ключ для сохранения сессии в Context
const CurSessionCtxKey = "CurSessionCtxKey"

// HeaderAuthenticationKey - ключь заголовка в котором содержится JWT токен
const HeaderAuthenticationKey = "Authentication"
