package timeuuid

import (
	"MyIM/pkg/mlog"
	"context"
	"github.com/gofrs/uuid"
	"time"
)

func GetTimeFromStr(ctx context.Context, uuidStr string) (t time.Time, err error) {
	parsedUUID, err := uuid.FromString(uuidStr)
	if err != nil {
		mlog.Logger().WithCTX(ctx).SetError("格式转换失败", err)
		return
	}
	_timestamp, err := uuid.TimestampFromV1(parsedUUID)
	if err != nil {
		mlog.Logger().WithCTX(ctx).SetError("格式转换失败", err)
		return
	}
	return _timestamp.Time()

}
