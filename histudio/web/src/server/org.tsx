import axios from 'utils/axios'


/*********************admin org公司信息*****************************/
// 公司信息 批量删除
export const adminOrgDelete = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/admin/org/delete`, {
            ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// 新增公司信息
export const adminOrgAdd = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/admin/org/add`, {
            ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })
    })
}
// 修改公司信息
export const adminOrgEdit = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.get(`/admin/org/edit`, {
            params: {
                id: data.id
            }
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })
    })
}
//公司信息回显
export const adminOrgView = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.get(`/admin/org/view`, {
            params: {
                id: data.id
            }
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })
    })
}
// 公司信息状态
export const adminOrgStatus = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/admin/org/status`, {
            ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}