package core

import (
	"fmt"
	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/pkg/cloud/model/apprun"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/webaccel-api-go"
	"reflect"
	"regexp"
	"strings"
)

// MatchResourceWithName matches a sakura internet resource with its name.
func MatchResourceWithName[T any](resources []T, name string) (T, error) {
	var null T
	v := reflect.ValueOf(resources)
	if v.Kind() != reflect.Slice {
		return null, fmt.Errorf("expected slice, got %T", resources)
	}

	for i := 0; i < v.Len(); i++ {
		switch resource := v.Index(i).Interface().(type) {
		case *sakuravps.Server:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case sakuravps.Server:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case *sakuravps.Switch:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case sakuravps.Switch:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case *sakuravps.NfsServer:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case sakuravps.NfsServer:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case *sakuravps.Permission:
			if resource.GetCode() == name {
				return v.Index(i).Interface().(T), nil
			}
		case sakuravps.Permission:
			if resource.GetCode() == name {
				return v.Index(i).Interface().(T), nil
			}
		case *sakuravps.Role:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case sakuravps.Role:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case *sakuravps.ApiKey:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case sakuravps.ApiKey:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case *iaas.DNS:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case iaas.DNS:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case *webaccel.Site:
			if resource.Name == name {
				return v.Index(i).Interface().(T), nil
			}
		case webaccel.Site:
			if resource.Name == name {
				return v.Index(i).Interface().(T), nil
			}
		case *apprun.HandlerListApplicationsData:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		case apprun.HandlerListApplicationsData:
			if resource.GetName() == name {
				return v.Index(i).Interface().(T), nil
			}
		}
	}
	return null, fmt.Errorf("no such resource: %q", name)
}

// SearchResourceWithName searches resource from a resource slice, with target resource name with its substring.
func SearchResourceWithName[T any](resources []T, pat string) (interface{}, error) {
	var null T
	v := reflect.ValueOf(resources)
	if v.Kind() != reflect.Slice {
		return null, fmt.Errorf("expected slice, got %T", resources)
	}
	switch list := interface{}(resources).(type) {
	case []*sakuravps.Server:
		var ret []*sakuravps.Server
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.Server:
		var ret []sakuravps.Server
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.Switch:
		var ret []*sakuravps.Switch
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.Switch:
		var ret []sakuravps.Switch
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.NfsServer:
		var ret []*sakuravps.NfsServer
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.NfsServer:
		var ret []sakuravps.NfsServer
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.Permission:
		var ret []*sakuravps.Permission
		for _, r := range list {
			if strings.Contains(r.GetCode(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.Permission:
		var ret []sakuravps.Permission
		for _, r := range list {
			if strings.Contains(r.GetCode(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.Role:
		var ret []*sakuravps.Role
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.Role:
		var ret []sakuravps.Role
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.ApiKey:
		var ret []*sakuravps.ApiKey
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.ApiKey:
		var ret []sakuravps.ApiKey
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*iaas.DNS:
		var ret []*iaas.DNS
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []iaas.DNS:
		var ret []iaas.DNS
		for _, r := range list {
			if strings.Contains(r.GetName(), pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*webaccel.Site:
		var ret []*webaccel.Site
		for _, r := range list {
			if strings.Contains(r.Name, pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []webaccel.Site:
		var ret []webaccel.Site
		for _, r := range list {
			if strings.Contains(r.Name, pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*apprun.HandlerListApplicationsData:
		var ret []*apprun.HandlerListApplicationsData
		for _, r := range list {
			if strings.Contains(r.Name, pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []apprun.HandlerListApplicationsData:
		var ret []apprun.HandlerListApplicationsData
		for _, r := range list {
			if strings.Contains(r.Name, pat) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	}
	return null, fmt.Errorf("no resources found by patterns: %q", pat)
}

// SearchResourceWithRegex searches resource from a resource slice, with target resource name with regex pattern.
func SearchResourceWithRegex[T any](resources []T, pat string) (interface{}, error) {
	var null T
	v := reflect.ValueOf(resources)
	expr, err := regexp.CompilePOSIX(pat)
	if err != nil {
		return null, err
	}
	if v.Kind() != reflect.Slice {
		return null, fmt.Errorf("expected slice, got %T", resources)
	}

	switch list := interface{}(resources).(type) {
	case []*sakuravps.Server:
		var ret []*sakuravps.Server
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.Server:
		var ret []sakuravps.Server
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.Switch:
		var ret []*sakuravps.Switch
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.Switch:
		var ret []sakuravps.Switch
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.NfsServer:
		var ret []*sakuravps.NfsServer
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.NfsServer:
		var ret []sakuravps.NfsServer
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.Permission:
		var ret []*sakuravps.Permission
		for _, r := range list {
			if expr.MatchString(r.GetCode()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.Permission:
		var ret []sakuravps.Permission
		for _, r := range list {
			if expr.MatchString(r.GetCode()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.Role:
		var ret []*sakuravps.Role
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.Role:
		var ret []sakuravps.Role
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*sakuravps.ApiKey:
		var ret []*sakuravps.ApiKey
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []sakuravps.ApiKey:
		var ret []sakuravps.ApiKey
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*iaas.DNS:
		var ret []*iaas.DNS
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []iaas.DNS:
		var ret []iaas.DNS
		for _, r := range list {
			if expr.MatchString(r.GetName()) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*webaccel.Site:
		var ret []*webaccel.Site
		for _, r := range list {
			if expr.MatchString(r.Name) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []webaccel.Site:
		var ret []webaccel.Site
		for _, r := range list {
			if expr.MatchString(r.Name) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []*apprun.HandlerListApplicationsData:
		var ret []*apprun.HandlerListApplicationsData
		for _, r := range list {
			if expr.MatchString(r.Name) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	case []apprun.HandlerListApplicationsData:
		var ret []apprun.HandlerListApplicationsData
		for _, r := range list {
			if expr.MatchString(r.Name) {
				ret = append(ret, r)
			}
		}
		return ret, nil
	}
	return nil, fmt.Errorf("no such resource: %q", pat)
}
