import axios from 'utils/axios'

// tg user 绑定员工 请求
export const tgUserBindUser = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post('/tg/tgUser/bindMember', { ...data }).then(res => {
            resolve(res)
        }).catch(err => {
            reject(err)
        })

    })
}
// // tg user 解绑员工 请求
export const tgUnUserBindUser = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post('/tg/tgUser/unBindMember', { ...data }).then(res => {
            resolve(res)
        }).catch(err => {
            reject(err)
        })

    })
}
// tg user 绑定代理 请求
export const tgUserBindProxy = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post('/tg/tgUser/bindProxy', { ...data }).then((res: any) => {
            resolve(res)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// tg user 解绑代理 请求
export const tgUserBindUnProxy = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post('/tg/tgUser/unBindProxy', { ...data }).then((res: any) => {
            resolve(res)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// tg user 批量上线下线 请求
// 上线 batchLogin 下线 batchLogout
export const tgUserBatchLoginOut = (data: any, type: String) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/arts/${type}`, { ...data }).then((res: any) => {
            resolve(res)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// tg user 批量删除 请求
export const tgUserAllDelete = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgUser/delete`, { ...data }).then((res: any) => {
            resolve(res)
        }).catch((err: any) => {
            reject(err)
        })

    })
}