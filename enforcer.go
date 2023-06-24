package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/labstack/echo/v4"
)

func Authenticate(adapter *gormadapter.Adapter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) (err error) {

			ctx := e.Request().Context()

			user, _, _ := e.Request().BasicAuth()
			method := e.Request().Method
			path := e.Request().URL.Path

			key := fmt.Sprintf("%s-%s-%s", user, path, method)

			result := RedisCache.Get(ctx, key)
			val, err := result.Result()
			if err == nil {
				boolValue, err := strconv.ParseBool(val)
				if err != nil {
					log.Fatal(err)
				}

				if !boolValue {
					return &echo.HTTPError{
						Code:    http.StatusForbidden,
						Message: "not allowed",
					}
				}
				return next(e)
			}

			// Casbin enforces policy
			ok, err := enforce(ctx, user, path, method, adapter)
			if err != nil || !ok {

				return &echo.HTTPError{
					Code:    http.StatusForbidden,
					Message: "not allowed",
				}
			}
			if !ok {
				return err
			}
			return next(e)
		}
	}
}

func enforce(ctx context.Context, sub string, obj string, act string, adapter *gormadapter.Adapter) (bool, error) {
	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("./examples/group_model.conf", adapter)
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	// Load policies from DB dynamically
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("error in policy: %w", err)
	}
	// Verify
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, fmt.Errorf("error in policy: %w", err)
	}
	key := fmt.Sprintf("%s-%s-%s", sub, obj, act)
	RedisCache.Set(ctx, key, strconv.FormatBool(ok), time.Hour)
	return ok, nil
}
