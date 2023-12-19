// table 表格
export const columns = [
    {
        title: 'Id',
        key: 'id'
    },
    {
        title: '聊天发起人',
        key: 'initiator'
    },
    {
        title: '发送人',
        key: 'sender'
    },
    {
        title: '接收人',
        key: 'receiver'
    },
    {
        title: '请求id',
        key: 'reqId'
    },
    {
        title: '消息类型',
        key: 'msgType'
    },
    {
        title: '发送时间',
        key: 'sendTime'
    },
    {
        title: '是否已读',
        key: 'read'
    },
    {
        title: '发送状态',
        key: 'sendStatus'
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
