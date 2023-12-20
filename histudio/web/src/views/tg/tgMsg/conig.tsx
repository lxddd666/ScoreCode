// table 表格
export const columns = [
    {
        title: '创建时间',
        key: 'createdAt'
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
        title: '消息内容',
        key: 'sendMsg'
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
        title: '更新时间',
        key: 'updatedAt'
    },
    {
        title: '操作',
        key: 'active'
    }
];

export const readArr = [
    {
        title: '已读',
        key: 1
    },
    {
        title: '未读',
        key: 2
    }
];
export const read = (value: any) => {

    let title = readArr.map((item) => {
        if (item.key === value) {
            return item.title;
        }
    });

    return title;
};

export const sendStatusArr = [
    {
        title: '成功',
        key: 1
    },
    {
        title: '失败',
        key: 2
    }
];
export const sendStatus = (value: any) => {

    let title = sendStatusArr.map((item) => {
        if (item.key === value) {
            return item.title;
        }
    });

    return title;
};