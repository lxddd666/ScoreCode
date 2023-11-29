import { useContext } from 'react';

// auth provider
import WebSocketContext from 'contexts/WebSocketContext';

// ==============================|| AUTH HOOKS ||============================== //

const useWebSocket = () => {
    const context = useContext(WebSocketContext);

    if (!context) throw new Error('context must be use inside provider');

    return context;
};

export default useWebSocket;
