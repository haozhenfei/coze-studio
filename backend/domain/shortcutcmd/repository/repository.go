/*
 * Copyright 2025 coze-dev Authors
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

package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/coze-dev/coze-studio/backend/domain/shortcutcmd/entity"
	"github.com/coze-dev/coze-studio/backend/domain/shortcutcmd/internal/dal"
	"github.com/coze-dev/coze-studio/backend/infra/contract/idgen"
)

func NewShortCutCmdRepo(db *gorm.DB, idGen idgen.IDGenerator) ShortCutCmdRepo {
	return dal.NewShortCutCmdDAO(db, idGen)
}

type ShortCutCmdRepo interface {
	List(ctx context.Context, lm *entity.ListMeta) ([]*entity.ShortcutCmd, error)
	Create(ctx context.Context, shortcut *entity.ShortcutCmd) (*entity.ShortcutCmd, error)
	Update(ctx context.Context, shortcut *entity.ShortcutCmd) (*entity.ShortcutCmd, error)
	GetByCmdID(ctx context.Context, cmdID int64, isOnline int32) (*entity.ShortcutCmd, error)
	PublishCMDs(ctx context.Context, objID int64, cmdIDs []int64) error
}
