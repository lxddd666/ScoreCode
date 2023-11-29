import React, { createContext, useEffect, useState } from 'react';
import useWebSocket, { Options } from 'react-use-websocket';
import axios from 'utils/axios';
import envRef from 'environment';
import useAuth from 'hooks/useAuth';
import { useIntl } from 'react-intl';
import { useDispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';
import axiosServices from 'utils/axios';
import { setAdminInfo } from 'store/slices/user';

// Create context to manage WebSocket
const WebSocketContext = createContext<any>(null);

// websocket default option
const socketDefaultOption: Options = {
    share: true,
    // onOpen: (e) => console.log(`WebSocket - Opened: \n============\n`, e, `\n============\nInitial Settings:\n${JSON.stringify(socketDefaultOption)}`),
    // onClose: (e) => console.log(`WebSocket - Closed: \n============\n`, e, `\n============\n`),
    // onMessage: (e) => console.log(`WebSocket - Message: \n============\n`, e, `\n============\n`),
    onError: (e) => console.error(`WebSocket - Error: \n============\n`, e, `\n============\n`),
    // TODO Modify these settings before deployment
    shouldReconnect: () => false,
    retryOnError: false
};

// Web Socket Events
export enum EventOptions {
    Ping = 'ping',
    MonitorTrends = 'admin/monitor/trends',
    MonitorRunInfo = 'admin/monitor/runInfo',
    Kick = 'kick',
    Notice = 'notice'
}

type EventInterval = { event: EventOptions; interval: NodeJS.Timer };

async function getSocketUrl() {
    try {
        const response = await axios.get(`${envRef?.API_URL}admin/site/config`, { headers: {} });
        if (response?.data?.code === 0) {
            return response.data.data.wsAddr;
        } else {
            throw new Error('Invalid response from the server');
        }
    } catch (error) {
        console.error('Error fetching WebSocket URL:', error);
        throw error;
    }
}

export const defaultIntervalTime = 5; //unit: seconds

export const WebSocketProvider = ({ children }: { children: React.ReactElement }) => {
    const intl = useIntl();
    const dispatch = useDispatch();
    const { isLoggedIn, logout, userInfo } = useAuth();
    const [connect, setConnect] = useState(true);
    const [messageHistory, setMessageHistory] = useState<MessageEvent<any>[] | null>([]);
    const [socketUrl, setSocketUrl] = useState<string | null>(null);
    const [heartBeatInterval, setHeartBeatInterval] = useState<EventInterval[]>([]); // to setup & clear interval for heart beat function
    const [notificationData, setNotificationData] = useState<{}>({});

    // Initialization
    async function initWebSocket() {
        try {
            const auth_token: string = axios.defaults.headers.common.Authorization?.toString().replace('Bearer ', '?authorization=') || '';
            const url = await getSocketUrl();
            const socketUrl = `${url}${auth_token}`;
            setSocketUrl(socketUrl);
            setConnect(true);
        } catch (error) {
            console.error('Error initializing WebSocket:', error);
        }
    }

    // Connect WebSocket on login,
    // if url=null, useWebSocket() will not connect
    useEffect(() => {
        if (isLoggedIn) {
            initWebSocket();
        }
    }, [isLoggedIn]);

    // Saving User Info into redux store
    useEffect(() => {
        if (userInfo) {
            dispatch(setAdminInfo(userInfo));
        }
    }, [userInfo]);

    async function handleHeartBeatPing(heartBeatInterval: EventInterval[]) {
        if (heartBeatInterval.length === 0) {
            await startHeartBeat(defaultIntervalTime, [EventOptions.Ping]);
        } else if (
            heartBeatInterval.filter((obj) => obj.event !== EventOptions.Ping).length > 0 &&
            heartBeatInterval.filter((obj) => obj.event === EventOptions.Ping).length !== 0
        ) {
            await stopHeartBeat(EventOptions.Ping);
        }
    }

    useEffect(() => {
        handleHeartBeatPing(heartBeatInterval);
    }, [heartBeatInterval]);

    // Default functions provided by library react-use-websocket
    const { sendMessage, lastMessage, lastJsonMessage, readyState } = useWebSocket(socketUrl, socketDefaultOption, connect);

    // Ready State Locale
    const readyStateString = {
        '-1': intl.formatMessage({ id: 'general.uninstantiated' }),
        0: intl.formatMessage({ id: 'general.connecting' }),
        1: intl.formatMessage({ id: 'general.connected' }),
        2: intl.formatMessage({ id: 'general.disconnecting' }),
        3: intl.formatMessage({ id: 'general.disconnected' })
    }[readyState];

    function handleLastMessage(lastMessage: MessageEvent<any>) {
        const obj = JSON.parse(lastMessage.data);
        switch (obj.event) {
            case EventOptions.Kick:
                kick();
                return;
            case EventOptions.Notice:
                notice(obj.data);
                return;
            default:
                return;
        }
    }

    // Handling message history & last message
    useEffect(() => {
        if (lastMessage) {
            handleLastMessage(lastMessage);
            if (messageHistory) {
                setMessageHistory([...messageHistory, lastMessage]);
            } else {
                setMessageHistory([lastMessage]);
            }
        }
    }, [lastMessage]);

    /**
     * A function to recursively sending selected event to the WebSocket.
     * Multiple events are allowed, but duplicated events will only overwrite
     * the previous event interval time, it will NOT add a new second event
     * with different interval time.
     * @param {number} interval - The interval time (in seconds).
     * @param {EventOptions} event - The selected event.
     */
    async function startHeartBeat(interval: number = defaultIntervalTime, events: EventOptions[]) {
        if (interval > 0) {
            await events.forEach((event) => {
                const isDuped = !!heartBeatInterval?.find((selected) => selected.event === event);
                if (!isDuped) {
                    const myInterval = setInterval(() => {
                        switch (event) {
                            case EventOptions.MonitorRunInfo:
                                monitorRunInfo();
                                break;
                            case EventOptions.MonitorTrends:
                                monitorTrends();
                                break;
                            case EventOptions.Ping:
                                ping();
                                break;
                            default:
                                ping();
                                break;
                        }
                    }, interval * 1000);
                    setHeartBeatInterval((prev) => [...prev, { event: event, interval: myInterval }]);
                }
            });
        }
    }

    async function stopHeartBeat(event?: EventOptions) {
        if (heartBeatInterval.length > 0) {
            // IF event is provided, clear the interval of that event ONLY
            if (event) {
                const selected = heartBeatInterval.filter((obj) => obj.event === event);
                await selected.forEach((obj) => {
                    clearInterval(obj.interval);
                });
                setHeartBeatInterval((prev) => prev.filter((obj) => obj.event !== event));
            }
            // IF event is NOT provided, clear intervals of ALL events
            else {
                await heartBeatInterval.forEach((obj) => {
                    // Keeping Ping event to keep websocket alive
                    // if want to clear Ping event, provide event param
                    clearInterval(obj.interval);
                });
                setHeartBeatInterval([]);
            }
        }
    }

    async function ping() {
        await sendMessage(JSON.stringify({ event: EventOptions.Ping }));
    }

    async function monitorTrends() {
        await sendMessage(JSON.stringify({ event: EventOptions.MonitorTrends }));
    }

    async function monitorRunInfo() {
        await sendMessage(JSON.stringify({ event: EventOptions.MonitorRunInfo }));
    }
    const handleLogout = async () => {
        try {
            await logout();
        } catch (err) {
            console.error(err);
        }
    };

    function kick() {
        unsubscribe();
        handleLogout();
    }

    function notice(data: any) {
        dispatch(
            openSnackbar({
                open: true,
                message: JSON.stringify(data),
                variant: 'notification',
                alert: {
                    color: 'info',
                    severity: 'info'
                },
                close: true,
                anchorOrigin: {
                    vertical: 'top',
                    horizontal: 'right'
                }
            })
        );
        getNotificationData();
    }

    async function getNotificationData() {
        await axiosServices
            .get(`${envRef?.API_URL}admin/notice/pullMessages`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setNotificationData(response.data.data);
                }
            })
            .catch(function (error) {
                console.error(error);
            });
    }

    function getMessageHistory() {
        return messageHistory;
    }

    function clearMessageHistory() {
        setMessageHistory([]);
    }

    async function subscribe() {
        await initWebSocket();
    }

    function unsubscribe() {
        setConnect(false);
        setSocketUrl(null);
    }

    return (
        <WebSocketContext.Provider
            value={{
                defaultIntervalTime,
                heartBeatInterval,
                lastMessage,
                lastJsonMessage,
                notificationData,
                subscribe,
                unsubscribe,
                ping,
                startHeartBeat,
                stopHeartBeat,
                monitorTrends,
                monitorRunInfo,
                kick,
                getMessageHistory,
                clearMessageHistory,
                getNotificationData,
                readyState,
                readyStateString
            }}
        >
            {children}
        </WebSocketContext.Provider>
    );
};

export default WebSocketContext;
