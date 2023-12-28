import axios from 'utils/axios'

/********************* TG User 账号*****************************/
// tg user 账号列表 请求
export const tgUserList = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.get('tg/tgUser/list', {
            params: {
                ...data
            }
        }).then(res => {
            resolve(res.data)
        }).catch(err => {
            reject(err)
        })

    })
}
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
// tg user 用户编辑 请求
export const tgUserEdit = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`tg/tgUser/edit`, { ...data }).then((res: any) => {
            resolve(res)
        }).catch((err: any) => {
            reject(err)
        })

    })
}

/********************* TG Folders 账号分组*****************************/
// tg Folders 账号分组 请求
export const tgFoldersList = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.get(`tg/tgFolders/list`, {
            params: {
                ...data
            }
        }).then((res: any) => {
            // console.log('tg Folders 账号分组 ', res);
            resolve(res)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
/********************* TG KeepTask 养号任务*****************************/
// 养号动作 请求
export const tgDictDataOptions = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.get(`/admin/dictData/options`, {
            params: {
                types: data
            }
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// 养号任务添加 请求
export const tgtgKeepTaskEdit = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgKeepTask/edit`, {
            ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// 养号任务数据回显 请求
export const tgtgKeepTaskEditEcho = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.get(`/tg/tgKeepTask/view`, {
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
// 养号任务 执行 请求
// 1 执行 2 暂停
export const tgKeepTaskExecute = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgKeepTask/status`, {
            ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// 养号任务 执行一次 请求
export const tgKeepTaskExecuteOne = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgKeepTask/once`, {
            ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// 养号任务 批量删除/删除 请求
export const tgKeepTaskAllDelete = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgKeepTask/delete`, {
            ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
/********************* TG IncreaseFansCron 涨粉任务*****************************/
// 校验频道 请求
export const tgIncreaseFansCronCheckChannel = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgIncreaseFansCron/checkChannel`, {
           ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// 校验频道 请求
export const tgIncreaseFansCronChannelIncreaseFanDetail = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgIncreaseFansCron/channelIncreaseFanDetail`, {
           ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// 涨粉任务添加/修改 请求
export const tgIncreaseFansCronEdit = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgIncreaseFansCron/edit`, {
           ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// 涨粉任务数据回显 请求
export const tgIncreaseFansCronEditEcho = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.get(`/tg/tgIncreaseFansCron/view`, {
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
// 执行/暂停 请求
export const tgIncreaseFansCronUpdateStatus = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgIncreaseFansCron/updateStatus`, {
            ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}
// 涨粉任务 批量删除 请求
export const tgIncreaseFansCronDelete = (data: any) => {
    return new Promise((resolve, reject) => {
        axios.post(`/tg/tgIncreaseFansCron/delete`, {
            ...data
        }).then((res: any) => {
            resolve(res.data)
        }).catch((err: any) => {
            reject(err)
        })

    })
}