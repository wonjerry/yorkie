/*
 * Copyright 2021 The Yorkie Authors. All rights reserved.
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

package memory

import (
	"context"

	"github.com/moby/locker"

	"github.com/yorkie-team/yorkie/internal/log"
)

type internalLocker struct {
	key   string
	locks *locker.Locker
}

// Lock locks the mutex.
func (il *internalLocker) Lock(ctx context.Context) error {
	il.locks.Lock(il.key)

	return nil
}

// TryLock locks the mutex if not already locked by another session.
func (il *internalLocker) TryLock(ctx context.Context) error {
	// TODO(hackerwins): We need to replace Lock with TryLock.
	il.locks.Lock(il.key)

	return nil
}

// Unlock unlocks the mutex.
func (il *internalLocker) Unlock(ctx context.Context) error {
	if err := il.locks.Unlock(il.key); err != nil {
		log.Logger.Error(err)
		return err
	}

	return nil
}
