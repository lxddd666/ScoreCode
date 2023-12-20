// table 表格
export const columns = [
    {
        title: 'id',
        key: 'id'
    },
    {
        title: '组织ID',
        key: 'orgId'
    },
    {
        title: '发起任务的用户ID',
        key: 'memberId'
    },
    {
        title: '任务名称',
        key: 'taskName'
    },
    {
        title: '频道地址',
        key: 'channel'
    },
    {
        title: '频道地址Id',
        key: 'channelId'
    },
    {
        title: '执行计划',
        key: 'executedPlan'
    },
    {
        title: '持续天数',
        key: 'dayCount'
    },
    {
        title: '涨粉数量',
        key: 'fansCount'
    },
    {
        title: '分组ID',
        key: 'folderId'
    },
    {
        title: '任务状态',
        key: 'cronStatus'
    },
    {
        title: '备注',
        key: 'comment'
    },
    {
        title: '创建时间',
        key: 'createdAt'
    },
    {
        title: '更新时间',
        key: 'updatedAt'
    }, {
        title: '已执行天数',
        key: 'executedDays'
    }, {
        title: '已添加粉丝数',
        key: 'increasedFans'
    },
    {
        title: '操作',
        key: 'active'
    }
];
export const accountStatusArr = [
    {
        title: '正常',
        key: 0
    },
    {
        title: '登陆失败',
        key: 1
    },
    {
        title: '未知',
        key: 2
    },
    {
        title: '不存在',
        key: 3
    },
    {
        title: '已封号',
        key: 4
    },
    {
        title: '认证失败',
        key: 5
    }
];
// 账号状态
export const accountStatus = (value: any) => {
    let title = accountStatusArr.filter((item) => {
        if (item.key === value) {
            return item.title;
        }
    });

    return title.length > 0 ? title[0].title : value;
};

export const isOnlineArr = [
    {
        title: '在线',
        key: 1
    },
    {
        title: '离线',
        key: 2
    }
];
export const isOnline = (value: any) => {

    let title = isOnlineArr.map((item) => {
        if (item.key === value) {
            return item.title;
        }
    });

    return title;
};
