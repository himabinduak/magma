/*
Copyright 2022 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package message

import (
	"context"
	"fmt"

	"magma/dp/cloud/go/services/dp/active_mode_controller/protos/active_mode"
)

func NewDeleteMessage(id int64) *deleteMessage {
	return &deleteMessage{id: id}
}

type deleteMessage struct {
	id int64
}

func (d *deleteMessage) Send(ctx context.Context, client active_mode.ActiveModeControllerClient) error {
	req := &active_mode.DeleteCbsdRequest{Id: d.id}
	_, err := client.DeleteCbsd(ctx, req)
	return err
}

func (d *deleteMessage) String() string {
	return fmt.Sprintf("delete: %d", d.id)
}