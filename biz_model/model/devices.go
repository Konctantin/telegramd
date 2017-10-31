/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"github.com/nebulaim/telegramd/base/orm"
)

type Devices struct {
	Id        int32
	AuthId    int64
	UserId    int32
	TokenType int32
	Token     string
	State     int32
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func init() {
	orm.RegisterModel(new(Devices))
}