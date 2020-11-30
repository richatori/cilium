// Copyright 2018-2019 Authors of Cilium
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

package cmd

import (
	"net"

	"github.com/cilium/cilium/api/v1/client/daemon"
	"github.com/cilium/cilium/api/v1/models"

	"github.com/spf13/cobra"
)

var nodeNeighRemoveCmd = &cobra.Command{
	Use:     "remove <IP addr>",
	Aliases: []string{"rm"},
	Short:   "Remove node as a neighbor from cluster",
	Example: "cilium node neigh remove 10.10.10.10",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			Usagef(cmd, "Missing node node IP")
		}

		if ip := net.ParseIP(args[1]); ip != nil && ip.To4() != nil {
			_, _, err := client.Daemon.DeleteClusterNodesNeigh(
				daemon.NewDeleteClusterNodesNeighParams().WithAddresses(&models.NodeAddressing{
					IPV4: &models.NodeAddressingElement{
						IP: ip.String(),
					},
				}))
			if err != nil {
				Fatalf("Unable to remove %q from neighbor table: %v\n", ip, err)
			}
		} else if ip != nil && ip.To16() != nil {
			_, _, err := client.Daemon.DeleteClusterNodesNeigh(
				daemon.NewDeleteClusterNodesNeighParams().WithAddresses(&models.NodeAddressing{
					IPV6: &models.NodeAddressingElement{
						IP: ip.String(),
					},
				}))
			if err != nil {
				Fatalf("Unable to remove %q from neighbor table: %v\n", ip, err)
			}
		} else {
			Fatalf("Invalid IP address %q\n", ip)
		}
	},
}

func init() {
	nodeNeighCmd.AddCommand(nodeNeighRemoveCmd)
}
