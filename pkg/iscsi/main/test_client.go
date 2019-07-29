/**
 * Copyright 2019 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/IBM/ibm-block-csi-driver-operator/pkg/iscsi/client"
	"github.com/operator-framework/operator-sdk/pkg/log/zap"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("iscsi agent")

func main() {
	logf.SetLogger(zap.Logger())
	c := client.NewIscsiClient("9.115.241.201:10086", log)
	c.Login([]string{"9.115.241.215", "9.115.241.219"})
	c.Logout([]string{"9.115.241.215", "9.115.241.219"})
}
