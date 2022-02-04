// Copyright 2022 The Howijd.Network Authors
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//   or LICENSE file in repository root.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/howijd/koios-rest-go-client"
	"github.com/urfave/cli/v2"
)

func addAddressCommands(app *cli.App, api *koios.Client) {
	app.Commands = append(app.Commands, []*cli.Command{
		{
			Name:     "address-info",
			Category: "ADDRESS",
			Usage:    "Get address info - balance, associated stake address (if any) and UTxO set.",
		},
		{
			Name:     "address-txs",
			Category: "ADDRESS",
			Usage:    "Get the transaction hash list of input address array, optionally filtering after specified block height (inclusive).",
		},
		{
			Name:     "address-assets",
			Category: "ADDRESS",
			Usage:    "Get the list of all the assets (policy, name and quantity) for a given address.",
		},
		{
			Name:     "credential-txs",
			Category: "ADDRESS",
			Usage:    "Get the transaction hash list of input payment credential array, optionally filtering after specified block height (inclusive).",
		},
	}...)
}
