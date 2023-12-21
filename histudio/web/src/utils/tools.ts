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
