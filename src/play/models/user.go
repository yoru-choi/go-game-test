// models/user.go
package models

// User 사용자 모델 구조체
type User struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

// // PushBlockUser 블록을 했을때 블록한 유저의 정보 추가
// func (r *UserBlockRepository) PushBlockUser(ctx context.Context, myUid uint64, blockedUid uint64, txTime time.Time) (int64, error) {
// 	query := bson.M{
// 		dbfields.ID: myUid,
// 	}
// 	blockInfo, _ := ccnmongo.FindOne[*entity.USER_BLOCK](ctx, r.mongoClients, query)
// 	if blockInfo != nil && len(blockInfo.BlockedUid) > 0 {
// 		for _, uid := range blockInfo.BlockedUid {
// 			if uid == blockedUid {
// 				return 0, nil
// 			}
// 		}
// 	}

// 	update := bson.M{
// 		"$push": bson.M{
// 			dbfields.BlockUid: blockedUid,
// 		},
// 		"$set": bson.M{
// 			dbfields.UpdateTime: txTime,
// 		},
// 		"$setOnInsert": bson.M{
// 			dbfields.CreateTime: txTime,
// 		},
// 	}
// 	upsertCount, err := clabmongo.UpsertOne[*entity.USER_BLOCK](ctx, r.mongoClients, query, update)
// 	if err != nil {
// 		return upsertCount, err
// 	}
// 	return upsertCount, nil
// }
