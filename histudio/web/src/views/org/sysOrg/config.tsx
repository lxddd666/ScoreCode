
export const columns = [
    {
        title: '公司ID',
        key: 'id'
    },
    {
        title: '公司名称',
        key: 'name'
    },
    {
        title: '公司编码',
        key: 'code'
    },
    {
        title: '负责人',
        key: 'leader'
    },
    {
        title: '联系电话',
        key: 'phone'
    },
    {
        title: '邮箱',
        key: 'email'
    },
    {
        title: '总端口数',
        key: 'ports'
    },
    {
        title: '公司状态',
        key: 'status'
    },
    {
        title: '创建时间',
        key: 'createdAt'
    },
    {
        title: '更新时间',
        key: 'updatedAt'
    }
];

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