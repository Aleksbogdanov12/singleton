/*
 * Copyright 2021 VMware, Inc.
 * SPDX-License-Identifier: EPL-2.0
 */

package translation

import (
	jsoniter "github.com/json-iterator/go"
)

// Request
type (
	ReleaseID struct {
		ProductName string `uri:"productName" form:"productName" binding:"alphanum"`
		Version     string `uri:"version" form:"version" binding:"version"`
	}

	BundleID struct {
		ReleaseID
		Locale    string `uri:"locale" binding:"locale"`
		Component string `uri:"component" binding:"component"`
	}

	StringID struct {
		BundleID
		Key string `uri:"key" binding:"key"`
	}

	GetStringReq struct {
		StringID
		Source string `form:"source"`
	}

	GetStringByPostReq struct {
		StringID
		Source                 string `form:"source"`
		CheckTranslationStatus bool   `form:"checkTranslationStatus"`
	}

	ProductReq struct {
		ReleaseID
		Locales    string `form:"locales" binding:"omitempty,locales"`
		Components string `form:"components" binding:"omitempty,components"`
	}

	UpdateTranslationDTO struct {
		Data      *UpdateBundle `json:"data" binding:"required"`
		Requester string        `json:"requester"`
	}
	UpdateBundle struct {
		ReleaseID
		Translation []*BundleData `json:"translation" binding:"required"`
		DataOrigin  string        `json:"dataOrigin"`
		Creation    struct {
			OperationID string `json:"operationid"`
		} `json:"creation"`
	}
	BundleData struct {
		Component string       `json:"component" binding:"component"`
		Locale    string       `json:"locale" binding:"locale"`
		Messages  jsoniter.Any `json:"messages" binding:"required"`
	}

	GetBundleReq struct {
		BundleID
		CheckTranslationStatus bool `form:"checkTranslationStatus" default:"false"`
	}
)

// Response
type (
	SingleBundleData struct {
		ProductName string       `json:"productName"`
		Version     string       `json:"version"`
		Locale      string       `json:"locale"`
		Component   string       `json:"component"`
		ID          int          `json:"id,omitempty"`
		Messages    jsoniter.Any `json:"messages"`

		Status map[string]interface{} `json:"status,omitempty"`
	}

	ReleaseData struct {
		ProductName string        `json:"productName"`
		Version     string        `json:"version"`
		Locales     []interface{} `json:"locales"`
		Components  []interface{} `json:"components"`
		Bundles     []BundleData  `json:"bundles,omitempty"`
		URL         string        `json:"url,omitempty"`
		ID          int           `json:"id,omitempty"`
	}
)