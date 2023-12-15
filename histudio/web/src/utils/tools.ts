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
