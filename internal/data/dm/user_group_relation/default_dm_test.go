package usergrouprelationdm

import (
	"fmt"
	"testing"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	usergrouprelationmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_group_relation"
	"github.com/stretchr/testify/assert"
)

func Test_getShardingTable(t *testing.T) {
	testCases := []struct {
		env             config.Env
		userIdOrGroupId uint64
		expected        string
	}{
		{
			env:             config.EnvTest,
			userIdOrGroupId: 1,
			expected:        "user_group_mapping_tab_00000001",
		},
		{
			env:             config.EnvTest,
			userIdOrGroupId: 12,
			expected:        "user_group_mapping_tab_00000000",
		},
		{
			env:             config.EnvLive,
			userIdOrGroupId: 1,
			expected:        "user_group_mapping_tab_00000001",
		},
		{
			env:             config.EnvLive,
			userIdOrGroupId: 12,
			expected:        "user_group_mapping_tab_00000012",
		},
		{
			env:             config.EnvLive,
			userIdOrGroupId: 123,
			expected:        "user_group_mapping_tab_00000123",
		},
		{
			env:             config.EnvLive,
			userIdOrGroupId: 1234,
			expected:        "user_group_mapping_tab_00000234",
		},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("test_%d", idx), func(t *testing.T) {
			mockDM := &defaultDM{
				tableName:           usergrouprelationmodel.UserGroupRelationTableName,
				shardingNum:         getShardingNumberByEnv(testCase.env),
				shardingTableFormat: usergrouprelationmodel.UserGroupRelationShardingTableFormat,
			}

			result := mockDM.getShardingTable(testCase.userIdOrGroupId)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
