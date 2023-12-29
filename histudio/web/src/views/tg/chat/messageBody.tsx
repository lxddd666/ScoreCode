import { memo } from 'react';
import Avatar from '@mui/material/Avatar';
import styles from './messageBody.module.scss';
import dayjs from 'dayjs'

// import { useScroll } from 'utils/tools';

const MessageBody = (props: any) => {
    const { messageList } = props;
    // const divRefs: any = useRef(null)
    // let { divRef, scrollInfo } = useScroll(divRefs)
    // console.log(divRef, scrollInfo, divRefs?.current?.scrollHeight);


    return (
        <div className={styles.context}>
            {messageList &&
                messageList.map((item: any, index: any) => {
                    return item.out === false ? (
                        // {item.out === true ? (
                        <div className={styles.other} key={index}>
                            <div className={styles.otherBody}>
                                <div>
                                    <Avatar
                                        variant="rounded"
                                        src="https://gw.alipayobjects.com/zos/antfincdn/x43I27A55%26/photo-1438109491414-7198515b166b.webp"
                                    />
                                </div>
                                <div className={styles.lineFont}>
                                    {item.message}
                                    
                                    <div className={styles.times}>{dayjs(item?.date * 1000).format('YYYY-MM-DD HH:mm:ss')}</div>
                                </div>
                            </div>
                        </div>
                    ) : (
                        <div className={styles.me} key={index}>
                            <div className={styles.meBody}>
                                <div className={styles.lineFont}>
                                    {item.message}

                                    <div className={styles.times}>{dayjs(item?.date * 1000).format('YYYY-MM-DD HH:mm:ss')}</div>
                                </div>
                                <Avatar
                                    variant="rounded"
                                    src="https://gw.alipayobjects.com/zos/antfincdn/x43I27A55%26/photo-1438109491414-7198515b166b.webp"
                                />


                            </div>
                        </div>
                        // )}
                    )
                })}
        </div >
    );
};
export default memo(MessageBody);
