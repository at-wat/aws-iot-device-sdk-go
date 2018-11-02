// Copyright 2018 SEQSENSE, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awsiotdev

import (
	"time"
)

type TopicPayload struct {
	Topic   string
	Payload string
}

type Options struct {
	KeyPath                  string
	CertPath                 string
	CaPath                   string
	ClientId                 string
	Region                   string
	BaseReconnectTime        time.Duration
	MaximumReconnectTime     time.Duration
	MinimumConnectionTime    time.Duration
	Keepalive                time.Duration
	Protocol                 string
	Host                     string
	Debug                    bool
	Qos                      byte
	Retain                   bool
	Will                     *TopicPayload
	OfflineQueueing          bool
	OfflineQueueMaxSize      int
	OfflineQueueDropBehavior string
	AutoResubscribe          bool
}
