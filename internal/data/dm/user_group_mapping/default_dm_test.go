package usergroupmappingdm

import (
	"fmt"
	"testing"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	usergroupmappingmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/user_group_mapping"
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
				tableName:           usergroupmappingmodel.UserGroupMappingTableName,
				shardingNum:         getShardingNumberByEnv(testCase.env),
				shardingTableFormat: usergroupmappingmodel.UserGroupMappingShardingTableFormat,
			}

			result := mockDM.getShardingTable(testCase.userIdOrGroupId)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
