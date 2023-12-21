// table 表格
export const columns = [
    {
        title: 'ID',
        key: 'id'
    },
    {
        title: '组织ID',
        key: 'orgId'
    },
    {
        title: '任务名称',
        key: 'taskName'
    },
    {
        title: '操作动作',
        key: 'action'
    },
    {
        title: '任务状态',
        key: 'status'
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
        title: '修改时间',
        key: 'updatedAt'
    }
];
//任务状态
export const statusArr = [
    {
        title: '停用',
        key: 1
    },
    {
        title: '正常',
        key: 2
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
