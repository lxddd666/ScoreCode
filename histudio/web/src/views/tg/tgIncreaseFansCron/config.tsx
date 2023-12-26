// table 表格
export const columns = [
    // {
    //     title: 'id',
    //     key: 'id'
    // },
    // {
    //     title: '组织ID',
    //     key: 'orgId'
    // },
    // {
    //     title: '发起任务的用户ID',
    //     key: 'memberId'
    // },
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

    // {
    //     title: '创建时间',
    //     key: 'createdAt'
    // },
    // {
    //     title: '更新时间',
    //     key: 'updatedAt'
    // },
    {
        title: '已执行天数',
        key: 'executedDays'
    }, {
        title: '已添加粉丝数',
        key: 'increasedFans'
    },
    {
        title: '备注',
        key: 'comment'
    },
    {
        title: '操作',
        key: 'active'
    }
];

// 任务状态
export const cronStatusArr = [
    {
        title: '终止',
        key: 0
    },
    {
        title: '正在执行',
        key: 1
    },
    {
        title: '完成',
        key: 2
    }
];
export const cronStatus = (value: any) => {

    let title = cronStatusArr.map((item) => {
        if (item.key === value) {
            return item.title;
        }
    });

    return title;
};
