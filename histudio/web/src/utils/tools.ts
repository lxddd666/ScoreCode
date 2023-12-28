import { useState, useEffect } from 'react';

// 获取窗体高度
export const useHeightComponent = (el: any) => {
    const [height, setHeight] = useState(0);
    const [width, setWidth] = useState(0);

    useEffect(() => {
        // 利用ref获取元素的当前高度，并更新状态
        if (el.current) {
            setHeight(el.current.getBoundingClientRect().height);
            setWidth(el.current.getBoundingClientRect().width);
        }

        // 创建一个函数来处理窗口大小变化的事件
        const handleResize = () => {
            if (el.current) {
                setHeight(el.current.getBoundingClientRect().height);
                setWidth(el.current.getBoundingClientRect().width);
            }
        };

        // 添加事件监听器
        window.addEventListener('resize', handleResize);

        // 清理函数：当组件卸载时，移除事件监听器
        return () => {
            window.removeEventListener('resize', handleResize);
        };
    }, [el]); // 空依赖数组意味着这个effect只在组件挂载时运行一次

    return {
        height,
        width
    };
};

export const useScroll = (divRef: any) => {
    // const divRef = useRef(null); // 创建一个ref来引用div元素
    const [scrollInfo, setScrollInfo] = useState({
        scrollTop: 0,
        divHeight: 0,
        scrollHeight: 0,
        offsetHeight: 0
    });

    useEffect(() => {
        const handleScroll = () => {
            // 当滚动事件触发时，更新状态
            setScrollInfo({
                scrollTop: divRef.current.scrollTop,
                divHeight: divRef.current.clientHeight,
                scrollHeight: divRef.current.scrollHeight,
                offsetHeight: divRef.current.height,
            });
        };

        const div = divRef.current;
        // 添加滚动监听
        div.addEventListener('scroll', handleScroll);

        // 组件卸载时移除监听
        return () => div.removeEventListener('scroll', handleScroll);
    }, []); // 空依赖数组意味着effect只在挂载时运行

    return { divRef, scrollInfo };
}
// 根据时间差返回对应的时间描述
export const timeAgo = (input: any) => {
    if (!input) return '从未登录'
    // 将输入转换为Date对象
    const date = (typeof input === 'object' && input instanceof Date) ? input : new Date(input);
    const now = new Date();
    const secondsPast = (now.getTime() - date.getTime()) / 1000;

    // 根据时间差返回对应的时间描述
    if (secondsPast < 60) { // 小于1分钟
        return '刚刚';
    } else if (secondsPast < 3600) { // 小于1小时
        const minutes = Math.round(secondsPast / 60);
        return `${minutes}分钟前`;
    } else if (secondsPast < 86400) { // 小于1天
        const hours = Math.round(secondsPast / 3600);
        return `${hours}小时前`;
    } else if (secondsPast < 604800) { // 小于1周
        const days = Math.round(secondsPast / 86400);
        return `${days}天前`;
    } else if (secondsPast < 2592000) { // 小于1个月
        const weeks = Math.round(secondsPast / 604800);
        return `${weeks}周前`;
    } else if (secondsPast < 31536000) { // 小于1年
        const months = Math.round(secondsPast / 2592000);
        return `${months}个月前`;
    } else { // 超过1年
        const years = Math.round(secondsPast / 31536000);
        return `${years}年前`;
    }
}

// 高阶函数，用于处理 async 函数的错误
export const handleAsync = async (asyncFn: any) => {
    try {
        const res = await asyncFn();
        return { res, error: null };
    } catch (error:any) {
        console.error('执行失败', error);
        return { res: null, error };
    }
}
