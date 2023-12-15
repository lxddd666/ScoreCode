import { memo } from 'react';
import Avatar from '@mui/material/Avatar';
import styles from './messageBody.module.scss';

const MessageBody = (props: any) => {
    const { messageList } = props;
    return (
        <div className={styles.context}>
            {messageList &&
                messageList.map((item: any) => {
                    return (
                        <>
                            {item.flag === 'o' ? (
                                <div className={styles.other} key={item.id}>
                                    <div className={styles.otherBody}>
                                        <div>
                                            <Avatar
                                                variant="rounded"
                                                src="https://gw.alipayobjects.com/zos/antfincdn/x43I27A55%26/photo-1438109491414-7198515b166b.webp"
                                            />
                                        </div>
                                        <div className={styles.lineFont}>{item.msg}</div>
                                    </div>
                                </div>
                            ) : (
                                <div className={styles.me} key={item.id}>
                                    <div className={styles.meBody}>
                                        <div className={styles.lineFont}>{item.msg}</div>
                                        <Avatar
                                            variant="rounded"
                                            src="https://gw.alipayobjects.com/zos/antfincdn/x43I27A55%26/photo-1438109491414-7198515b166b.webp"
                                        />
                                    </div>
                                </div>
                            )}
                        </>
                    );
                })}
        </div>
    );
};
export default memo(MessageBody);
