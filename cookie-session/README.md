# Go 中的 Cookie 和 Session 

今天聊聊 Cookie 和 Session 的登录认证怎么搞，以及能用在哪。

## 1. 准备
```go
var sessionStore sync.Map
type User struct {
    Name     string `json:"name" binding:"required"`
    Password string `json:"password" binding:"required"`
}
```
- `sync.Map` 存 Session，线程安全。
- `User` 定义用户名和密码，必填。

## 2. 生成 Session ID

```go
func generateSessionID(user User) (string, error) {
    sessionID := uuid.NewString()
    userData, _ := json.Marshal(user)
    sessionStore.Store(sessionID, userData)
    return sessionID, nil
}
```
- `uuid` 生成随机 ID，用户信息转 JSON 存起来。

## 3. 查 Session
```go
func getUserFromSession(sessionID string) (*User, bool) {
    if data, ok := sessionStore.Load(sessionID); ok {
        var user User
        json.Unmarshal(data.([]byte), &user)
        return &user, true
    }
    return nil, false
}
```
- 用 Session ID 查数据，反序列化成 `User`。

## 4. 登录逻辑
```go
func loginHandler(ctx *gin.Context) {
    var user User
    if ctx.ShouldBindJSON(&user) != nil {
        ctx.JSON(400, gin.H{"error": "bad request"})
        return
    }
    if sessionID, _ := ctx.Cookie("auth"); sessionID != "" {
        if cachedUser, found := getUserFromSession(sessionID); found {
            ctx.JSON(200, cachedUser)
            return
        }
    }
    sessionID, _ := generateSessionID(user)
    ctx.SetCookie("auth", sessionID, 3600, "/", "localhost", false, false)
    ctx.JSON(200, gin.H{"message": "登录成功"})
}
```
- 校验请求，Cookie 命中返回用户信息，没命中生成新 Session ID，设 Cookie。

## 原理
> Cookie 携带 Session ID，服务器通过 ID 查询 Redis 缓存，若命中则返回用户信息，否则生成新会话并缓存后返回客户端。

## 局限

Cookie 未加密，多机不共享。

## 改进
- Redis 存 Session。
- Cookie 加安全标志。

## 总结

小项目好使，大规模用`token`，更安全
