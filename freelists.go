// Copyright 2015 Google Inc. All Rights Reserved.
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

package fuse

import (
	"unsafe"

	"github.com/jacobsa/fuse/internal/buffer"
)

// LOCKS_EXCLUDED(c.mu)
func (c *Connection) getInMessage() (m *buffer.InMessage) {
	c.mu.Lock()
	m = (*buffer.InMessage)(c.inMessages.Get())
	if m == nil {
		m = new(buffer.InMessage)
	}
	c.mu.Unlock()
	return
}

// LOCKS_EXCLUDED(c.mu)
func (c *Connection) putInMessage(x *buffer.InMessage) {
	c.mu.Lock()
	c.inMessages.Put(unsafe.Pointer(x))
	c.mu.Unlock()
}
