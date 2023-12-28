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
        title: '创建时间',
        key: 'createdAt'
    },
    {
        title: '操作',
        key: 'active'
    },

];

// 任务状态
export const cronStatusArr = [
    {
        title: '正在执行',
        color: 'secondary',
        key: 0
    },
    {
        title: '完成',
        color: 'success',
        key: 1
    },
    {
        title: '失败',
        color: 'error',
        key: 2
    },
    {
        title: '暂停',
        color: 'info',
        key: 3
    }
];
export const cronStatus = (value: any) => {

    let title: any = cronStatusArr.filter((item) => {
        if (item.key === value) {
            return item;
        }
    });

    return title[0] ? title[0] : { title: '未定义状态', color: 'default' };
};
