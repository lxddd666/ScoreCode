
export const columns = [
    {
        title: 'id',
        key: 'id'
    },
    {
        title: '代理地址',
        key: 'address'
    },
    {
        title: '代理类型',
        key: 'type'
    },
    {
        title: '最大连接数',
        key: 'maxConnections'
    },
    {
        title: '已连接数',
        key: 'connectedCount'
    },
    {
        title: '已分配账号数量',
        key: 'assignedCount'
    },
    {
        title: '长期未登录数量',
        key: 'longTermCount'
    },
    {
        title: '地区',
        key: 'region'
    },
    {
        title: '延迟',
        key: 'delay'
    },
    {
        title: '备注',
        key: 'comment'
    },
    {
        title: '状态',
        key: 'status'
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
        title: '操作',
        key: 'active'
    }
];
//代理类型
export const typeArr = [
    {
        title: 'http',
        key: 0
    },
    {
        title: 'https',
        key: 1
    },
    {
        title: 'socks5',
        key: 2
    }
];
export const type = (value: any) => {
    let title = typeArr.map((item) => {
        if (item.key === value) {
            return item.title;
        }
    });

    return title;
};
//状态
export const statusArr = [
    {
        title: '停用',
        key: 0
    },
    {
        title: '正常',
        key: 1
    }
];
export const status = (value: any) => {
    let title = statusArr.map((item) => {
        if (item.key === value) {
            return item.title;
        }
    });

    return title;
};