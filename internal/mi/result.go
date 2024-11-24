// Copyright 2024 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build windows

package mi

import "errors"

type ResultError uint32

const (
	MI_RESULT_OK ResultError = iota
	MI_RESULT_FAILED
	MI_RESULT_ACCESS_DENIED
	MI_RESULT_INVALID_NAMESPACE
	MI_RESULT_INVALID_PARAMETER
	MI_RESULT_INVALID_CLASS
	MI_RESULT_NOT_FOUND
	MI_RESULT_NOT_SUPPORTED
	MI_RESULT_CLASS_HAS_CHILDREN
	MI_RESULT_CLASS_HAS_INSTANCES
	MI_RESULT_INVALID_SUPERCLASS
	MI_RESULT_ALREADY_EXISTS
	MI_RESULT_NO_SUCH_PROPERTY
	MI_RESULT_TYPE_MISMATCH
	MI_RESULT_QUERY_LANGUAGE_NOT_SUPPORTED
	MI_RESULT_INVALID_QUERY
	MI_RESULT_METHOD_NOT_AVAILABLE
	MI_RESULT_METHOD_NOT_FOUND
	MI_RESULT_NAMESPACE_NOT_EMPTY
	MI_RESULT_INVALID_ENUMERATION_CONTEXT
	MI_RESULT_INVALID_OPERATION_TIMEOUT
	MI_RESULT_PULL_HAS_BEEN_ABANDONED
	MI_RESULT_PULL_CANNOT_BE_ABANDONED
	MI_RESULT_FILTERED_ENUMERATION_NOT_SUPPORTED
	MI_RESULT_CONTINUATION_ON_ERROR_NOT_SUPPORTED
	MI_RESULT_SERVER_LIMITS_EXCEEDED
	MI_RESULT_SERVER_IS_SHUTTING_DOWN
)

func (r ResultError) Error() string {
	return r.String()
}

func (r ResultError) String() string {
	switch {
	case errors.Is(r, MI_RESULT_OK):
		return "MI_RESULT_OK"
	case errors.Is(r, MI_RESULT_FAILED):
		return "MI_RESULT_FAILED"
	case errors.Is(r, MI_RESULT_ACCESS_DENIED):
		return "MI_RESULT_ACCESS_DENIED"
	case errors.Is(r, MI_RESULT_INVALID_NAMESPACE):
		return "MI_RESULT_INVALID_NAMESPACE"
	case errors.Is(r, MI_RESULT_INVALID_PARAMETER):
		return "MI_RESULT_INVALID_PARAMETER"
	case errors.Is(r, MI_RESULT_INVALID_CLASS):
		return "MI_RESULT_INVALID_CLASS"
	case errors.Is(r, MI_RESULT_NOT_FOUND):
		return "MI_RESULT_NOT_FOUND"
	case errors.Is(r, MI_RESULT_NOT_SUPPORTED):
		return "MI_RESULT_NOT_SUPPORTED"
	case errors.Is(r, MI_RESULT_CLASS_HAS_CHILDREN):
		return "MI_RESULT_CLASS_HAS_CHILDREN"
	case errors.Is(r, MI_RESULT_CLASS_HAS_INSTANCES):
		return "MI_RESULT_CLASS_HAS_INSTANCES"
	case errors.Is(r, MI_RESULT_INVALID_SUPERCLASS):
		return "MI_RESULT_INVALID_SUPERCLASS"
	case errors.Is(r, MI_RESULT_ALREADY_EXISTS):
		return "MI_RESULT_ALREADY_EXISTS"
	case errors.Is(r, MI_RESULT_NO_SUCH_PROPERTY):
		return "MI_RESULT_NO_SUCH_PROPERTY"
	case errors.Is(r, MI_RESULT_TYPE_MISMATCH):
		return "MI_RESULT_TYPE_MISMATCH"
	case errors.Is(r, MI_RESULT_QUERY_LANGUAGE_NOT_SUPPORTED):
		return "MI_RESULT_QUERY_LANGUAGE_NOT_SUPPORTED"
	case errors.Is(r, MI_RESULT_INVALID_QUERY):
		return "MI_RESULT_INVALID_QUERY"
	case errors.Is(r, MI_RESULT_METHOD_NOT_AVAILABLE):
		return "MI_RESULT_METHOD_NOT_AVAILABLE"
	case errors.Is(r, MI_RESULT_METHOD_NOT_FOUND):
		return "MI_RESULT_METHOD_NOT_FOUND"
	case errors.Is(r, MI_RESULT_NAMESPACE_NOT_EMPTY):
		return "MI_RESULT_NAMESPACE_NOT_EMPTY"
	case errors.Is(r, MI_RESULT_INVALID_ENUMERATION_CONTEXT):
		return "MI_RESULT_INVALID_ENUMERATION_CONTEXT"
	case errors.Is(r, MI_RESULT_INVALID_OPERATION_TIMEOUT):
		return "MI_RESULT_INVALID_OPERATION_TIMEOUT"
	case errors.Is(r, MI_RESULT_PULL_HAS_BEEN_ABANDONED):
		return "MI_RESULT_PULL_HAS_BEEN_ABANDONED"
	case errors.Is(r, MI_RESULT_PULL_CANNOT_BE_ABANDONED):
		return "MI_RESULT_PULL_CANNOT_BE_ABANDONED"
	case errors.Is(r, MI_RESULT_FILTERED_ENUMERATION_NOT_SUPPORTED):
		return "MI_RESULT_FILTERED_ENUMERATION_NOT_SUPPORTED"
	case errors.Is(r, MI_RESULT_CONTINUATION_ON_ERROR_NOT_SUPPORTED):
		return "MI_RESULT_CONTINUATION_ON_ERROR_NOT_SUPPORTED"
	case errors.Is(r, MI_RESULT_SERVER_LIMITS_EXCEEDED):
		return "MI_RESULT_SERVER_LIMITS_EXCEEDED"
	case errors.Is(r, MI_RESULT_SERVER_IS_SHUTTING_DOWN):
		return "MI_RESULT_SERVER_IS_SHUTTING_DOWN"
	default:
		return "MI_RESULT_UNKNOWN"
	}
}
